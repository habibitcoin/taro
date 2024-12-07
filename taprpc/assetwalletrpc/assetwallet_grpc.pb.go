// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package assetwalletrpc

import (
	context "context"
	taprpc "github.com/lightninglabs/taproot-assets/taprpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AssetWalletClient is the client API for AssetWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetWalletClient interface {
	// Creates a virtual PSBT template for an interactive asset send.
	CreateInteractiveSendTemplate(ctx context.Context, in *CreateInteractiveSendTemplateRequest, opts ...grpc.CallOption) (*CreateInteractiveSendTemplateResponse, error)
	// PrepareAnchoringTemplate prepares a Bitcoin PSBT for anchoring the given
	// signed virtual PSBT.
	PrepareAnchoringTemplate(ctx context.Context, in *PrepareAnchoringTemplateRequest, opts ...grpc.CallOption) (*PrepareAnchoringTemplateResponse, error)
	// UpdateVirtualPSBT updates the script key of the output in the given virtual
	// PSBT and optionally the linked Bitcoin PSBT.
	UpdateVirtualPsbt(ctx context.Context, in *UpdateVirtualPsbtRequest, opts ...grpc.CallOption) (*UpdateVirtualPsbtResponse, error)
	// FundVirtualPsbt selects inputs from the available asset commitments to fund
	// a virtual transaction matching the template.
	FundVirtualPsbt(ctx context.Context, in *FundVirtualPsbtRequest, opts ...grpc.CallOption) (*FundVirtualPsbtResponse, error)
	// SignVirtualPsbt signs the inputs of a virtual transaction and prepares the
	// commitments of the inputs and outputs.
	SignVirtualPsbt(ctx context.Context, in *SignVirtualPsbtRequest, opts ...grpc.CallOption) (*SignVirtualPsbtResponse, error)
	// AnchorVirtualPsbts merges and then commits multiple virtual transactions in
	// a single BTC level anchor transaction. This RPC should be used if the BTC
	// level anchor transaction of the assets to be spent are encumbered by a
	// normal key and don't require any special spending conditions. For any custom
	// spending conditions on the BTC level, the two RPCs CommitVirtualPsbts and
	// PublishAndLogTransfer should be used instead (which in combination do the
	// same as this RPC but allow for more flexibility).
	AnchorVirtualPsbts(ctx context.Context, in *AnchorVirtualPsbtsRequest, opts ...grpc.CallOption) (*taprpc.SendAssetResponse, error)
	// CommitVirtualPsbts creates the output commitments and proofs for the given
	// virtual transactions by committing them to the BTC level anchor transaction.
	// In addition, the BTC level anchor transaction is funded and prepared up to
	// the point where it is ready to be signed.
	CommitVirtualPsbts(ctx context.Context, in *CommitVirtualPsbtsRequest, opts ...grpc.CallOption) (*CommitVirtualPsbtsResponse, error)
	// PublishAndLogTransfer accepts a fully committed and signed anchor
	// transaction and publishes it to the Bitcoin network. It also logs the
	// transfer of the given active and passive assets in the database and ships
	// any outgoing proofs to the counterparties.
	PublishAndLogTransfer(ctx context.Context, in *PublishAndLogRequest, opts ...grpc.CallOption) (*taprpc.SendAssetResponse, error)
	// NextInternalKey derives the next internal key for the given key family and
	// stores it as an internal key in the database to make sure it is identified
	// as a local key later on when importing proofs. While an internal key can
	// also be used as the internal key of a script key, it is recommended to use
	// the NextScriptKey RPC instead, to make sure the tweaked Taproot output key
	// is also recognized as a local key.
	NextInternalKey(ctx context.Context, in *NextInternalKeyRequest, opts ...grpc.CallOption) (*NextInternalKeyResponse, error)
	// NextScriptKey derives the next script key (and its corresponding internal
	// key) and stores them both in the database to make sure they are identified
	// as local keys later on when importing proofs.
	NextScriptKey(ctx context.Context, in *NextScriptKeyRequest, opts ...grpc.CallOption) (*NextScriptKeyResponse, error)
	// QueryInternalKey returns the key descriptor for the given internal key.
	QueryInternalKey(ctx context.Context, in *QueryInternalKeyRequest, opts ...grpc.CallOption) (*QueryInternalKeyResponse, error)
	// QueryScriptKey returns the full script key descriptor for the given tweaked
	// script key.
	QueryScriptKey(ctx context.Context, in *QueryScriptKeyRequest, opts ...grpc.CallOption) (*QueryScriptKeyResponse, error)
	// tapcli: `proofs proveownership`
	// ProveAssetOwnership creates an ownership proof embedded in an asset
	// transition proof. That ownership proof is a signed virtual transaction
	// spending the asset with a valid witness to prove the prover owns the keys
	// that can spend the asset.
	ProveAssetOwnership(ctx context.Context, in *ProveAssetOwnershipRequest, opts ...grpc.CallOption) (*ProveAssetOwnershipResponse, error)
	// tapcli: `proofs verifyownership`
	// VerifyAssetOwnership verifies the asset ownership proof embedded in the
	// given transition proof of an asset and returns true if the proof is valid.
	VerifyAssetOwnership(ctx context.Context, in *VerifyAssetOwnershipRequest, opts ...grpc.CallOption) (*VerifyAssetOwnershipResponse, error)
	// RemoveUTXOLease removes the lease/lock/reservation of the given managed
	// UTXO.
	RemoveUTXOLease(ctx context.Context, in *RemoveUTXOLeaseRequest, opts ...grpc.CallOption) (*RemoveUTXOLeaseResponse, error)
	// DeclareScriptKey declares a new script key to the wallet. This is useful
	// when the script key contains scripts, which would mean it wouldn't be
	// recognized by the wallet automatically. Declaring a script key will make any
	// assets sent to the script key be recognized as being local assets.
	DeclareScriptKey(ctx context.Context, in *DeclareScriptKeyRequest, opts ...grpc.CallOption) (*DeclareScriptKeyResponse, error)
}

type assetWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetWalletClient(cc grpc.ClientConnInterface) AssetWalletClient {
	return &assetWalletClient{cc}
}

func (c *assetWalletClient) CreateInteractiveSendTemplate(ctx context.Context, in *CreateInteractiveSendTemplateRequest, opts ...grpc.CallOption) (*CreateInteractiveSendTemplateResponse, error) {
	out := new(CreateInteractiveSendTemplateResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/CreateInteractiveSendTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) PrepareAnchoringTemplate(ctx context.Context, in *PrepareAnchoringTemplateRequest, opts ...grpc.CallOption) (*PrepareAnchoringTemplateResponse, error) {
	out := new(PrepareAnchoringTemplateResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/PrepareAnchoringTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) UpdateVirtualPsbt(ctx context.Context, in *UpdateVirtualPsbtRequest, opts ...grpc.CallOption) (*UpdateVirtualPsbtResponse, error) {
	out := new(UpdateVirtualPsbtResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/UpdateVirtualPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) FundVirtualPsbt(ctx context.Context, in *FundVirtualPsbtRequest, opts ...grpc.CallOption) (*FundVirtualPsbtResponse, error) {
	out := new(FundVirtualPsbtResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/FundVirtualPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) SignVirtualPsbt(ctx context.Context, in *SignVirtualPsbtRequest, opts ...grpc.CallOption) (*SignVirtualPsbtResponse, error) {
	out := new(SignVirtualPsbtResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/SignVirtualPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) AnchorVirtualPsbts(ctx context.Context, in *AnchorVirtualPsbtsRequest, opts ...grpc.CallOption) (*taprpc.SendAssetResponse, error) {
	out := new(taprpc.SendAssetResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/AnchorVirtualPsbts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) CommitVirtualPsbts(ctx context.Context, in *CommitVirtualPsbtsRequest, opts ...grpc.CallOption) (*CommitVirtualPsbtsResponse, error) {
	out := new(CommitVirtualPsbtsResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/CommitVirtualPsbts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) PublishAndLogTransfer(ctx context.Context, in *PublishAndLogRequest, opts ...grpc.CallOption) (*taprpc.SendAssetResponse, error) {
	out := new(taprpc.SendAssetResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/PublishAndLogTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) NextInternalKey(ctx context.Context, in *NextInternalKeyRequest, opts ...grpc.CallOption) (*NextInternalKeyResponse, error) {
	out := new(NextInternalKeyResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/NextInternalKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) NextScriptKey(ctx context.Context, in *NextScriptKeyRequest, opts ...grpc.CallOption) (*NextScriptKeyResponse, error) {
	out := new(NextScriptKeyResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/NextScriptKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) QueryInternalKey(ctx context.Context, in *QueryInternalKeyRequest, opts ...grpc.CallOption) (*QueryInternalKeyResponse, error) {
	out := new(QueryInternalKeyResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/QueryInternalKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) QueryScriptKey(ctx context.Context, in *QueryScriptKeyRequest, opts ...grpc.CallOption) (*QueryScriptKeyResponse, error) {
	out := new(QueryScriptKeyResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/QueryScriptKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) ProveAssetOwnership(ctx context.Context, in *ProveAssetOwnershipRequest, opts ...grpc.CallOption) (*ProveAssetOwnershipResponse, error) {
	out := new(ProveAssetOwnershipResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/ProveAssetOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) VerifyAssetOwnership(ctx context.Context, in *VerifyAssetOwnershipRequest, opts ...grpc.CallOption) (*VerifyAssetOwnershipResponse, error) {
	out := new(VerifyAssetOwnershipResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/VerifyAssetOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) RemoveUTXOLease(ctx context.Context, in *RemoveUTXOLeaseRequest, opts ...grpc.CallOption) (*RemoveUTXOLeaseResponse, error) {
	out := new(RemoveUTXOLeaseResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/RemoveUTXOLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetWalletClient) DeclareScriptKey(ctx context.Context, in *DeclareScriptKeyRequest, opts ...grpc.CallOption) (*DeclareScriptKeyResponse, error) {
	out := new(DeclareScriptKeyResponse)
	err := c.cc.Invoke(ctx, "/assetwalletrpc.AssetWallet/DeclareScriptKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetWalletServer is the server API for AssetWallet service.
// All implementations must embed UnimplementedAssetWalletServer
// for forward compatibility
type AssetWalletServer interface {
	// Creates a virtual PSBT template for an interactive asset send.
	CreateInteractiveSendTemplate(context.Context, *CreateInteractiveSendTemplateRequest) (*CreateInteractiveSendTemplateResponse, error)
	// PrepareAnchoringTemplate prepares a Bitcoin PSBT for anchoring the given
	// signed virtual PSBT.
	PrepareAnchoringTemplate(context.Context, *PrepareAnchoringTemplateRequest) (*PrepareAnchoringTemplateResponse, error)
	// UpdateVirtualPSBT updates the script key of the output in the given virtual
	// PSBT and optionally the linked Bitcoin PSBT.
	UpdateVirtualPsbt(context.Context, *UpdateVirtualPsbtRequest) (*UpdateVirtualPsbtResponse, error)
	// FundVirtualPsbt selects inputs from the available asset commitments to fund
	// a virtual transaction matching the template.
	FundVirtualPsbt(context.Context, *FundVirtualPsbtRequest) (*FundVirtualPsbtResponse, error)
	// SignVirtualPsbt signs the inputs of a virtual transaction and prepares the
	// commitments of the inputs and outputs.
	SignVirtualPsbt(context.Context, *SignVirtualPsbtRequest) (*SignVirtualPsbtResponse, error)
	// AnchorVirtualPsbts merges and then commits multiple virtual transactions in
	// a single BTC level anchor transaction. This RPC should be used if the BTC
	// level anchor transaction of the assets to be spent are encumbered by a
	// normal key and don't require any special spending conditions. For any custom
	// spending conditions on the BTC level, the two RPCs CommitVirtualPsbts and
	// PublishAndLogTransfer should be used instead (which in combination do the
	// same as this RPC but allow for more flexibility).
	AnchorVirtualPsbts(context.Context, *AnchorVirtualPsbtsRequest) (*taprpc.SendAssetResponse, error)
	// CommitVirtualPsbts creates the output commitments and proofs for the given
	// virtual transactions by committing them to the BTC level anchor transaction.
	// In addition, the BTC level anchor transaction is funded and prepared up to
	// the point where it is ready to be signed.
	CommitVirtualPsbts(context.Context, *CommitVirtualPsbtsRequest) (*CommitVirtualPsbtsResponse, error)
	// PublishAndLogTransfer accepts a fully committed and signed anchor
	// transaction and publishes it to the Bitcoin network. It also logs the
	// transfer of the given active and passive assets in the database and ships
	// any outgoing proofs to the counterparties.
	PublishAndLogTransfer(context.Context, *PublishAndLogRequest) (*taprpc.SendAssetResponse, error)
	// NextInternalKey derives the next internal key for the given key family and
	// stores it as an internal key in the database to make sure it is identified
	// as a local key later on when importing proofs. While an internal key can
	// also be used as the internal key of a script key, it is recommended to use
	// the NextScriptKey RPC instead, to make sure the tweaked Taproot output key
	// is also recognized as a local key.
	NextInternalKey(context.Context, *NextInternalKeyRequest) (*NextInternalKeyResponse, error)
	// NextScriptKey derives the next script key (and its corresponding internal
	// key) and stores them both in the database to make sure they are identified
	// as local keys later on when importing proofs.
	NextScriptKey(context.Context, *NextScriptKeyRequest) (*NextScriptKeyResponse, error)
	// QueryInternalKey returns the key descriptor for the given internal key.
	QueryInternalKey(context.Context, *QueryInternalKeyRequest) (*QueryInternalKeyResponse, error)
	// QueryScriptKey returns the full script key descriptor for the given tweaked
	// script key.
	QueryScriptKey(context.Context, *QueryScriptKeyRequest) (*QueryScriptKeyResponse, error)
	// tapcli: `proofs proveownership`
	// ProveAssetOwnership creates an ownership proof embedded in an asset
	// transition proof. That ownership proof is a signed virtual transaction
	// spending the asset with a valid witness to prove the prover owns the keys
	// that can spend the asset.
	ProveAssetOwnership(context.Context, *ProveAssetOwnershipRequest) (*ProveAssetOwnershipResponse, error)
	// tapcli: `proofs verifyownership`
	// VerifyAssetOwnership verifies the asset ownership proof embedded in the
	// given transition proof of an asset and returns true if the proof is valid.
	VerifyAssetOwnership(context.Context, *VerifyAssetOwnershipRequest) (*VerifyAssetOwnershipResponse, error)
	// RemoveUTXOLease removes the lease/lock/reservation of the given managed
	// UTXO.
	RemoveUTXOLease(context.Context, *RemoveUTXOLeaseRequest) (*RemoveUTXOLeaseResponse, error)
	// DeclareScriptKey declares a new script key to the wallet. This is useful
	// when the script key contains scripts, which would mean it wouldn't be
	// recognized by the wallet automatically. Declaring a script key will make any
	// assets sent to the script key be recognized as being local assets.
	DeclareScriptKey(context.Context, *DeclareScriptKeyRequest) (*DeclareScriptKeyResponse, error)
	mustEmbedUnimplementedAssetWalletServer()
}

// UnimplementedAssetWalletServer must be embedded to have forward compatible implementations.
type UnimplementedAssetWalletServer struct {
}

func (UnimplementedAssetWalletServer) CreateInteractiveSendTemplate(context.Context, *CreateInteractiveSendTemplateRequest) (*CreateInteractiveSendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInteractiveSendTemplate not implemented")
}
func (UnimplementedAssetWalletServer) PrepareAnchoringTemplate(context.Context, *PrepareAnchoringTemplateRequest) (*PrepareAnchoringTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareAnchoringTemplate not implemented")
}
func (UnimplementedAssetWalletServer) UpdateVirtualPsbt(context.Context, *UpdateVirtualPsbtRequest) (*UpdateVirtualPsbtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVirtualPsbt not implemented")
}
func (UnimplementedAssetWalletServer) FundVirtualPsbt(context.Context, *FundVirtualPsbtRequest) (*FundVirtualPsbtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundVirtualPsbt not implemented")
}
func (UnimplementedAssetWalletServer) SignVirtualPsbt(context.Context, *SignVirtualPsbtRequest) (*SignVirtualPsbtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignVirtualPsbt not implemented")
}
func (UnimplementedAssetWalletServer) AnchorVirtualPsbts(context.Context, *AnchorVirtualPsbtsRequest) (*taprpc.SendAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnchorVirtualPsbts not implemented")
}
func (UnimplementedAssetWalletServer) CommitVirtualPsbts(context.Context, *CommitVirtualPsbtsRequest) (*CommitVirtualPsbtsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitVirtualPsbts not implemented")
}
func (UnimplementedAssetWalletServer) PublishAndLogTransfer(context.Context, *PublishAndLogRequest) (*taprpc.SendAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAndLogTransfer not implemented")
}
func (UnimplementedAssetWalletServer) NextInternalKey(context.Context, *NextInternalKeyRequest) (*NextInternalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextInternalKey not implemented")
}
func (UnimplementedAssetWalletServer) NextScriptKey(context.Context, *NextScriptKeyRequest) (*NextScriptKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextScriptKey not implemented")
}
func (UnimplementedAssetWalletServer) QueryInternalKey(context.Context, *QueryInternalKeyRequest) (*QueryInternalKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInternalKey not implemented")
}
func (UnimplementedAssetWalletServer) QueryScriptKey(context.Context, *QueryScriptKeyRequest) (*QueryScriptKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryScriptKey not implemented")
}
func (UnimplementedAssetWalletServer) ProveAssetOwnership(context.Context, *ProveAssetOwnershipRequest) (*ProveAssetOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProveAssetOwnership not implemented")
}
func (UnimplementedAssetWalletServer) VerifyAssetOwnership(context.Context, *VerifyAssetOwnershipRequest) (*VerifyAssetOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAssetOwnership not implemented")
}
func (UnimplementedAssetWalletServer) RemoveUTXOLease(context.Context, *RemoveUTXOLeaseRequest) (*RemoveUTXOLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUTXOLease not implemented")
}
func (UnimplementedAssetWalletServer) DeclareScriptKey(context.Context, *DeclareScriptKeyRequest) (*DeclareScriptKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclareScriptKey not implemented")
}
func (UnimplementedAssetWalletServer) mustEmbedUnimplementedAssetWalletServer() {}

// UnsafeAssetWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetWalletServer will
// result in compilation errors.
type UnsafeAssetWalletServer interface {
	mustEmbedUnimplementedAssetWalletServer()
}

func RegisterAssetWalletServer(s grpc.ServiceRegistrar, srv AssetWalletServer) {
	s.RegisterService(&AssetWallet_ServiceDesc, srv)
}

func _AssetWallet_CreateInteractiveSendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInteractiveSendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).CreateInteractiveSendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/CreateInteractiveSendTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).CreateInteractiveSendTemplate(ctx, req.(*CreateInteractiveSendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_PrepareAnchoringTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareAnchoringTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).PrepareAnchoringTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/PrepareAnchoringTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).PrepareAnchoringTemplate(ctx, req.(*PrepareAnchoringTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_UpdateVirtualPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVirtualPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).UpdateVirtualPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/UpdateVirtualPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).UpdateVirtualPsbt(ctx, req.(*UpdateVirtualPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_FundVirtualPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundVirtualPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).FundVirtualPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/FundVirtualPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).FundVirtualPsbt(ctx, req.(*FundVirtualPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_SignVirtualPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignVirtualPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).SignVirtualPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/SignVirtualPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).SignVirtualPsbt(ctx, req.(*SignVirtualPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_AnchorVirtualPsbts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnchorVirtualPsbtsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).AnchorVirtualPsbts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/AnchorVirtualPsbts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).AnchorVirtualPsbts(ctx, req.(*AnchorVirtualPsbtsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_CommitVirtualPsbts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitVirtualPsbtsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).CommitVirtualPsbts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/CommitVirtualPsbts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).CommitVirtualPsbts(ctx, req.(*CommitVirtualPsbtsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_PublishAndLogTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishAndLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).PublishAndLogTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/PublishAndLogTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).PublishAndLogTransfer(ctx, req.(*PublishAndLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_NextInternalKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextInternalKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).NextInternalKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/NextInternalKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).NextInternalKey(ctx, req.(*NextInternalKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_NextScriptKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextScriptKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).NextScriptKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/NextScriptKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).NextScriptKey(ctx, req.(*NextScriptKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_QueryInternalKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInternalKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).QueryInternalKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/QueryInternalKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).QueryInternalKey(ctx, req.(*QueryInternalKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_QueryScriptKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryScriptKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).QueryScriptKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/QueryScriptKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).QueryScriptKey(ctx, req.(*QueryScriptKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_ProveAssetOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProveAssetOwnershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).ProveAssetOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/ProveAssetOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).ProveAssetOwnership(ctx, req.(*ProveAssetOwnershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_VerifyAssetOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAssetOwnershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).VerifyAssetOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/VerifyAssetOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).VerifyAssetOwnership(ctx, req.(*VerifyAssetOwnershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_RemoveUTXOLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUTXOLeaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).RemoveUTXOLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/RemoveUTXOLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).RemoveUTXOLease(ctx, req.(*RemoveUTXOLeaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetWallet_DeclareScriptKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclareScriptKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetWalletServer).DeclareScriptKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetwalletrpc.AssetWallet/DeclareScriptKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetWalletServer).DeclareScriptKey(ctx, req.(*DeclareScriptKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetWallet_ServiceDesc is the grpc.ServiceDesc for AssetWallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetWallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assetwalletrpc.AssetWallet",
	HandlerType: (*AssetWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInteractiveSendTemplate",
			Handler:    _AssetWallet_CreateInteractiveSendTemplate_Handler,
		},
		{
			MethodName: "PrepareAnchoringTemplate",
			Handler:    _AssetWallet_PrepareAnchoringTemplate_Handler,
		},
		{
			MethodName: "UpdateVirtualPsbt",
			Handler:    _AssetWallet_UpdateVirtualPsbt_Handler,
		},
		{
			MethodName: "FundVirtualPsbt",
			Handler:    _AssetWallet_FundVirtualPsbt_Handler,
		},
		{
			MethodName: "SignVirtualPsbt",
			Handler:    _AssetWallet_SignVirtualPsbt_Handler,
		},
		{
			MethodName: "AnchorVirtualPsbts",
			Handler:    _AssetWallet_AnchorVirtualPsbts_Handler,
		},
		{
			MethodName: "CommitVirtualPsbts",
			Handler:    _AssetWallet_CommitVirtualPsbts_Handler,
		},
		{
			MethodName: "PublishAndLogTransfer",
			Handler:    _AssetWallet_PublishAndLogTransfer_Handler,
		},
		{
			MethodName: "NextInternalKey",
			Handler:    _AssetWallet_NextInternalKey_Handler,
		},
		{
			MethodName: "NextScriptKey",
			Handler:    _AssetWallet_NextScriptKey_Handler,
		},
		{
			MethodName: "QueryInternalKey",
			Handler:    _AssetWallet_QueryInternalKey_Handler,
		},
		{
			MethodName: "QueryScriptKey",
			Handler:    _AssetWallet_QueryScriptKey_Handler,
		},
		{
			MethodName: "ProveAssetOwnership",
			Handler:    _AssetWallet_ProveAssetOwnership_Handler,
		},
		{
			MethodName: "VerifyAssetOwnership",
			Handler:    _AssetWallet_VerifyAssetOwnership_Handler,
		},
		{
			MethodName: "RemoveUTXOLease",
			Handler:    _AssetWallet_RemoveUTXOLease_Handler,
		},
		{
			MethodName: "DeclareScriptKey",
			Handler:    _AssetWallet_DeclareScriptKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assetwalletrpc/assetwallet.proto",
}
