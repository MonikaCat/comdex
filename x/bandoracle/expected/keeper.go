package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	// ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"

	assettypes "github.com/MonikaCat/comdex/v5/x/asset/types"
	marketttypes "github.com/MonikaCat/comdex/v5/x/market/types"
)

type MarketKeeper interface {
	DeleteTwaData(ctx sdk.Context, assetID uint64)
	GetAllTwa(ctx sdk.Context) (twa []marketttypes.TimeWeightedAverage)
}

type AssetKeeper interface {
	GetAssets(ctx sdk.Context, id uint64) (assettypes.Asset, bool)
}

type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, sourcePort string, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error)
}

// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

type ScopedKeeper interface {
	AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool
	ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error
	GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool)
}
