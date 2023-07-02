package types

import (
	"github.com/comdex-official/comdex/x/nft/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "nft/MsgCreateDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "nft/MsgUpdateDenom", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "nft/MsgTransferDenom", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "nft/MsgBurnNFT", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgUpdateDenom{},
		&MsgTransferDenom{},
		&MsgMintNFT{},
		&MsgTransferNFT{},
		&MsgBurnNFT{},
	)

	registry.RegisterImplementations((*exported.NFT)(nil),
		&NFT{},
	)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func MustMarshalSupply(cdc codec.BinaryCodec, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyWrap)
}

func MustUnMarshalSupply(cdc codec.BinaryCodec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

func MustMarshalNFTID(cdc codec.BinaryCodec, nftID string) []byte {
	nftIDWrap := gogotypes.StringValue{Value: nftID}
	return cdc.MustMarshal(&nftIDWrap)
}

func MustMarshalDenomID(cdc codec.BinaryCodec, denomID string) []byte {
	denomIDWrap := gogotypes.StringValue{Value: denomID}
	return cdc.MustMarshal(&denomIDWrap)
}

func MustUnMarshalDenomID(cdc codec.BinaryCodec, value []byte) string {
	var denomIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &denomIDWrap)
	return denomIDWrap.Value
}
