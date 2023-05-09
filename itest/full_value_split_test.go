package itest

import (
	"context"

	"github.com/lightninglabs/taro/tarorpc"
	"github.com/lightninglabs/taro/tarorpc/mintrpc"
	"github.com/stretchr/testify/require"
)

// testFullValueSend tests that we can properly send the full value of a
// normal asset.
func testFullValueSend(t *harnessTest) {
	// First, we'll make an normal assets with enough units to allow us to
	// send it around a few times.
	rpcAssets := mintAssetsConfirmBatch(
		t, t.tarod, []*mintrpc.MintAssetRequest{
			simpleAssets[0], issuableAssets[0],
		},
	)

	mintedAsset := rpcAssets[0]
	mintedGroupAsset := rpcAssets[1]
	genInfo := mintedAsset.AssetGenesis
	groupGenInfo := mintedGroupAsset.AssetGenesis

	ctxb := context.Background()
	ctxt, cancel := context.WithTimeout(ctxb, defaultWaitTimeout)
	defer cancel()

	// Now that we have the asset created, we'll make a new node that'll
	// serve as the node which'll receive the assets.
	secondTarod := setupTarodHarness(
		t.t, t, t.lndHarness.Bob, t.universeServer,
		func(params *tarodHarnessParams) {
			params.startupSyncNode = t.tarod
			params.startupSyncNumAssets = len(rpcAssets)
		},
	)
	defer func() {
		require.NoError(t.t, secondTarod.stop(true))
	}()

	runFullValueSendTests(
		ctxt, t, t.tarod, secondTarod, genInfo, mintedAsset, 0, 1,
	)
	runFullValueSendTests(
		ctxt, t, t.tarod, secondTarod, groupGenInfo, mintedGroupAsset,
		1, 2,
	)
}

// runFullValueSendTests runs the full value send tests for a single asset.
func runFullValueSendTests(ctxt context.Context, t *harnessTest, alice,
	bob *tarodHarness, genInfo *tarorpc.GenesisInfo,
	mintedAsset *tarorpc.Asset, runIdx, numRuns int) {

	// Next, we'll attempt to complete three transfers of the full value of
	// the asset between our main node and Bob.
	var (
		numSends            = 3
		senderTransferIdx   = runIdx * 2
		receiverTransferIdx = runIdx * 1
		fullAmount          = mintedAsset.Amount
		receiverAddr        *tarorpc.Addr
		err                 error
	)

	for i := 0; i < numSends; i++ {
		// Create an address for the receiver and send the asset. We
		// start with Bob receiving the asset, then sending it back
		// to the main node, and so on.
		if i%2 == 0 {
			receiverAddr, err = bob.NewAddr(
				ctxt, &tarorpc.NewAddrRequest{
					AssetId: genInfo.AssetId,
					Amt:     fullAmount,
				},
			)
			require.NoError(t.t, err)

			assertAddrCreated(t.t, bob, mintedAsset, receiverAddr)
			sendResp := sendAssetsToAddr(t, alice, receiverAddr)
			confirmAndAssertOutboundTransfer(
				t, alice, sendResp, genInfo.AssetId,
				[]uint64{0, fullAmount}, senderTransferIdx,
				senderTransferIdx+1,
			)
			_ = sendProof(
				t, alice, bob, receiverAddr.ScriptKey, genInfo,
			)
			senderTransferIdx++
		} else {
			receiverAddr, err = alice.NewAddr(
				ctxt, &tarorpc.NewAddrRequest{
					AssetId: genInfo.AssetId,
					Amt:     fullAmount,
				},
			)
			require.NoError(t.t, err)

			assertAddrCreated(t.t, alice, mintedAsset, receiverAddr)
			sendResp := sendAssetsToAddr(t, bob, receiverAddr)
			confirmAndAssertOutboundTransfer(
				t, bob, sendResp, genInfo.AssetId,
				[]uint64{0, fullAmount}, receiverTransferIdx,
				receiverTransferIdx+1,
			)
			_ = sendProof(
				t, bob, alice, receiverAddr.ScriptKey, genInfo,
			)
			receiverTransferIdx++
		}
	}

	// Check the final state of both nodes. The main node should list 2
	// zero-value transfers. and Bob should have 1. The main node should
	// show a balance of zero, and Bob should hold the total asset supply.
	assertTransfer(t.t, alice, runIdx*2, numRuns*2, []uint64{0, fullAmount})
	assertTransfer(t.t, alice, runIdx*2+1, numRuns*2, []uint64{0, fullAmount})
	assertBalanceByID(t.t, alice, genInfo.AssetId, 0)

	assertTransfer(t.t, bob, runIdx, numRuns, []uint64{0, fullAmount})
	assertBalanceByID(t.t, bob, genInfo.AssetId, fullAmount)
}
