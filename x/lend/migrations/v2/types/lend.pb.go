package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LendAsset struct {
	ID                  uint64                                  `protobuf:"varint,1,opt,name=lending_id,json=lendingId,proto3" json:"lending_id,omitempty" yaml:"lending_id"`
	AssetID             uint64                                  `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	PoolID              uint64                                  `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	Owner               string                                  `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	AmountIn            github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	LendingTime         time.Time                               `protobuf:"bytes,6,opt,name=lending_time,json=lendingTime,proto3,stdtime" json:"lending_time" yaml:"lending_time"`
	AvailableToBorrow   github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,7,opt,name=available_to_borrow,json=availableToBorrow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"available_to_borrow" yaml:"available_to_borrow"`
	AppID               uint64                                  `protobuf:"varint,8,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	GlobalIndex         github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,9,opt,name=global_index,json=globalIndex,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"global_index" yaml:"global_index"`
	LastInteractionTime time.Time                               `protobuf:"bytes,10,opt,name=last_interaction_time,json=lastInteractionTime,proto3,stdtime" json:"last_interaction_time" yaml:"last_interaction_time"`
	CPoolName           string                                  `protobuf:"bytes,11,opt,name=cpool_name,json=cpoolName,proto3" json:"cpool_name,omitempty" yaml:"cpool_name"`
	TotalRewards        github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,12,opt,name=total_rewards,json=totalRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_rewards" yaml:"total_rewards"`
}

func (m *LendAsset) Reset()         { *m = LendAsset{} }
func (m *LendAsset) String() string { return proto.CompactTextString(m) }
func (*LendAsset) ProtoMessage()    {}
func (*LendAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{0}
}
func (m *LendAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendAsset.Merge(m, src)
}
func (m *LendAsset) XXX_Size() int {
	return m.Size()
}
func (m *LendAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_LendAsset.DiscardUnknown(m)
}

var xxx_messageInfo_LendAsset proto.InternalMessageInfo

func (m *LendAsset) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LendAsset) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *LendAsset) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *LendAsset) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *LendAsset) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *LendAsset) GetLendingTime() time.Time {
	if m != nil {
		return m.LendingTime
	}
	return time.Time{}
}

func (m *LendAsset) GetAppID() uint64 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *LendAsset) GetLastInteractionTime() time.Time {
	if m != nil {
		return m.LastInteractionTime
	}
	return time.Time{}
}

func (m *LendAsset) GetCPoolName() string {
	if m != nil {
		return m.CPoolName
	}
	return ""
}

type BorrowAsset struct {
	ID                  uint64                                  `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty" yaml:"borrowing_id"`
	LendingID           uint64                                  `protobuf:"varint,2,opt,name=lending_id,json=lendingId,proto3" json:"lending_id,omitempty" yaml:"lending_id"`
	IsStableBorrow      bool                                    `protobuf:"varint,3,opt,name=is_stable_borrow,json=isStableBorrow,proto3" json:"is_stable_borrow,omitempty" yaml:"is_stable_borrow"`
	PairID              uint64                                  `protobuf:"varint,4,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	AmountIn            github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	AmountOut           github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=amount_out,json=amountOut,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_out" yaml:"amount_out"`
	BridgedAssetAmount  github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,7,opt,name=bridged_asset_amount,json=bridgedAssetAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"bridged_asset_amount" yaml:"bridged_asset_amount"`
	BorrowingTime       time.Time                               `protobuf:"bytes,8,opt,name=borrowing_time,json=borrowingTime,proto3,stdtime" json:"borrowing_time" yaml:"borrowing_time"`
	StableBorrowRate    github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,9,opt,name=stable_borrow_rate,json=stableBorrowRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_borrow_rate" yaml:"stable_borrow_rate"`
	InterestAccumulated github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,10,opt,name=interest_accumulated,json=interestAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_accumulated" yaml:"interest_accumulated"`
	GlobalIndex         github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,11,opt,name=global_index,json=globalIndex,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"global_index" yaml:"global_index"`
	ReserveGlobalIndex  github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,12,opt,name=reserve_global_index,json=reserveGlobalIndex,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve_global_index" yaml:"reserve_global_index"`
	LastInteractionTime time.Time                               `protobuf:"bytes,13,opt,name=last_interaction_time,json=lastInteractionTime,proto3,stdtime" json:"last_interaction_time" yaml:"last_interaction_time"`
	CPoolName           string                                  `protobuf:"bytes,14,opt,name=cpool_name,json=cpoolName,proto3" json:"cpool_name,omitempty" yaml:"cpool_name"`
	IsLiquidated        bool                                    `protobuf:"varint,15,opt,name=is_liquidated,json=isLiquidated,proto3" json:"is_liquidated,omitempty" yaml:"is_liquidated"`
}

func (m *BorrowAsset) Reset()         { *m = BorrowAsset{} }
func (m *BorrowAsset) String() string { return proto.CompactTextString(m) }
func (*BorrowAsset) ProtoMessage()    {}
func (*BorrowAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{1}
}
func (m *BorrowAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BorrowAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BorrowAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BorrowAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowAsset.Merge(m, src)
}
func (m *BorrowAsset) XXX_Size() int {
	return m.Size()
}
func (m *BorrowAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowAsset.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowAsset proto.InternalMessageInfo

func (m *BorrowAsset) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *BorrowAsset) GetLendingID() uint64 {
	if m != nil {
		return m.LendingID
	}
	return 0
}

func (m *BorrowAsset) GetIsStableBorrow() bool {
	if m != nil {
		return m.IsStableBorrow
	}
	return false
}

func (m *BorrowAsset) GetPairID() uint64 {
	if m != nil {
		return m.PairID
	}
	return 0
}

func (m *BorrowAsset) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *BorrowAsset) GetAmountOut() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountOut
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *BorrowAsset) GetBridgedAssetAmount() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.BridgedAssetAmount
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *BorrowAsset) GetBorrowingTime() time.Time {
	if m != nil {
		return m.BorrowingTime
	}
	return time.Time{}
}

func (m *BorrowAsset) GetLastInteractionTime() time.Time {
	if m != nil {
		return m.LastInteractionTime
	}
	return time.Time{}
}

func (m *BorrowAsset) GetCPoolName() string {
	if m != nil {
		return m.CPoolName
	}
	return ""
}

func (m *BorrowAsset) GetIsLiquidated() bool {
	if m != nil {
		return m.IsLiquidated
	}
	return false
}

type Pool struct {
	PoolID     uint64                  `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	ModuleName string                  `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty" yaml:"module_name"`
	CPoolName  string                  `protobuf:"bytes,3,opt,name=cpool_name,json=cpoolName,proto3" json:"cpool_name,omitempty" yaml:"cpool_name"`
	AssetData  []*AssetDataPoolMapping `protobuf:"bytes,4,rep,name=asset_data,json=assetData,proto3" json:"asset_data,omitempty" yaml:"asset_data"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{2}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *Pool) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *Pool) GetCPoolName() string {
	if m != nil {
		return m.CPoolName
	}
	return ""
}

func (m *Pool) GetAssetData() []*AssetDataPoolMapping {
	if m != nil {
		return m.AssetData
	}
	return nil
}

type UserAssetLendBorrowMapping struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// to check if poool id is needed
	LendId   uint64   `protobuf:"varint,2,opt,name=lend_id,json=lendId,proto3" json:"lend_id,omitempty" yaml:"lend_id"`
	PoolId   uint64   `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	BorrowId []uint64 `protobuf:"varint,4,rep,packed,name=borrow_id,json=borrowId,proto3" json:"borrow_id,omitempty" yaml:"borrow_id"`
}

func (m *UserAssetLendBorrowMapping) Reset()         { *m = UserAssetLendBorrowMapping{} }
func (m *UserAssetLendBorrowMapping) String() string { return proto.CompactTextString(m) }
func (*UserAssetLendBorrowMapping) ProtoMessage()    {}
func (*UserAssetLendBorrowMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{3}
}
func (m *UserAssetLendBorrowMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAssetLendBorrowMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserAssetLendBorrowMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserAssetLendBorrowMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAssetLendBorrowMapping.Merge(m, src)
}
func (m *UserAssetLendBorrowMapping) XXX_Size() int {
	return m.Size()
}
func (m *UserAssetLendBorrowMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAssetLendBorrowMapping.DiscardUnknown(m)
}

var xxx_messageInfo_UserAssetLendBorrowMapping proto.InternalMessageInfo

func (m *UserAssetLendBorrowMapping) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UserAssetLendBorrowMapping) GetLendId() uint64 {
	if m != nil {
		return m.LendId
	}
	return 0
}

func (m *UserAssetLendBorrowMapping) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *UserAssetLendBorrowMapping) GetBorrowId() []uint64 {
	if m != nil {
		return m.BorrowId
	}
	return nil
}

type AssetDataPoolMapping struct {
	AssetID uint64 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	// 1 for main_asset, 2 for 1st transit_asset, 3 for 2nd transit_asset
	AssetTransitType uint64                                 `protobuf:"varint,2,opt,name=asset_transit_type,json=assetTransitType,proto3" json:"asset_transit_type,omitempty" yaml:"asset_transit_type"`
	SupplyCap        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=supply_cap,json=supplyCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"supply_cap" yaml:"supply_cap"`
}

func (m *AssetDataPoolMapping) Reset()         { *m = AssetDataPoolMapping{} }
func (m *AssetDataPoolMapping) String() string { return proto.CompactTextString(m) }
func (*AssetDataPoolMapping) ProtoMessage()    {}
func (*AssetDataPoolMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{4}
}
func (m *AssetDataPoolMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetDataPoolMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetDataPoolMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetDataPoolMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetDataPoolMapping.Merge(m, src)
}
func (m *AssetDataPoolMapping) XXX_Size() int {
	return m.Size()
}
func (m *AssetDataPoolMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetDataPoolMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AssetDataPoolMapping proto.InternalMessageInfo

func (m *AssetDataPoolMapping) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *AssetDataPoolMapping) GetAssetTransitType() uint64 {
	if m != nil {
		return m.AssetTransitType
	}
	return 0
}

type Extended_Pair struct {
	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetIn         uint64 `protobuf:"varint,2,opt,name=asset_in,json=assetIn,proto3" json:"asset_in,omitempty" yaml:"asset_in"`
	AssetOut        uint64 `protobuf:"varint,3,opt,name=asset_out,json=assetOut,proto3" json:"asset_out,omitempty" yaml:"asset_out"`
	IsInterPool     bool   `protobuf:"varint,4,opt,name=is_inter_pool,json=isInterPool,proto3" json:"is_inter_pool,omitempty" yaml:"is_inter_pool"`
	AssetOutPoolID  uint64 `protobuf:"varint,5,opt,name=asset_out_pool_id,json=assetOutPoolId,proto3" json:"asset_out_pool_id,omitempty" yaml:"asset_out_pool_id"`
	MinUsdValueLeft uint64 `protobuf:"varint,6,opt,name=min_usd_value_left,json=minUsdValueLeft,proto3" json:"min_usd_value_left,omitempty" yaml:"min_usd_value_left"`
}

func (m *Extended_Pair) Reset()         { *m = Extended_Pair{} }
func (m *Extended_Pair) String() string { return proto.CompactTextString(m) }
func (*Extended_Pair) ProtoMessage()    {}
func (*Extended_Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{5}
}
func (m *Extended_Pair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Extended_Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Extended_Pair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Extended_Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extended_Pair.Merge(m, src)
}
func (m *Extended_Pair) XXX_Size() int {
	return m.Size()
}
func (m *Extended_Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Extended_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Extended_Pair proto.InternalMessageInfo

func (m *Extended_Pair) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Extended_Pair) GetAssetIn() uint64 {
	if m != nil {
		return m.AssetIn
	}
	return 0
}

func (m *Extended_Pair) GetAssetOut() uint64 {
	if m != nil {
		return m.AssetOut
	}
	return 0
}

func (m *Extended_Pair) GetIsInterPool() bool {
	if m != nil {
		return m.IsInterPool
	}
	return false
}

func (m *Extended_Pair) GetAssetOutPoolID() uint64 {
	if m != nil {
		return m.AssetOutPoolID
	}
	return 0
}

func (m *Extended_Pair) GetMinUsdValueLeft() uint64 {
	if m != nil {
		return m.MinUsdValueLeft
	}
	return 0
}

type AssetToPairMapping struct {
	PoolID  uint64   `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	AssetID uint64   `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	PairID  []uint64 `protobuf:"varint,3,rep,packed,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
}

func (m *AssetToPairMapping) Reset()         { *m = AssetToPairMapping{} }
func (m *AssetToPairMapping) String() string { return proto.CompactTextString(m) }
func (*AssetToPairMapping) ProtoMessage()    {}
func (*AssetToPairMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{6}
}
func (m *AssetToPairMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetToPairMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetToPairMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetToPairMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetToPairMapping.Merge(m, src)
}
func (m *AssetToPairMapping) XXX_Size() int {
	return m.Size()
}
func (m *AssetToPairMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetToPairMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AssetToPairMapping proto.InternalMessageInfo

func (m *AssetToPairMapping) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *AssetToPairMapping) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *AssetToPairMapping) GetPairID() []uint64 {
	if m != nil {
		return m.PairID
	}
	return nil
}

type PoolAssetLBMapping struct {
	PoolID                   uint64                                 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	AssetID                  uint64                                 `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	LendIds                  []uint64                               `protobuf:"varint,3,rep,packed,name=lend_ids,json=lendIds,proto3" json:"lend_ids,omitempty" yaml:"lend_ids"`
	BorrowIds                []uint64                               `protobuf:"varint,4,rep,packed,name=borrow_ids,json=borrowIds,proto3" json:"borrow_ids,omitempty" yaml:"borrow_ids"`
	TotalBorrowed            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_borrowed,json=totalBorrowed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_borrowed" yaml:"total_borrowed"`
	TotalStableBorrowed      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=total_stable_borrowed,json=totalStableBorrowed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_stable_borrowed" yaml:"total_stable_borrowed"`
	TotalLend                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=total_lend,json=totalLend,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_lend" yaml:"total_lend"`
	TotalInterestAccumulated github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=total_interest_accumulated,json=totalInterestAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_interest_accumulated" yaml:"total_interest_accumulated"`
	LendApr                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=lend_apr,json=lendApr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"lend_apr" yaml:"lend_apr"`
	BorrowApr                github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=borrow_apr,json=borrowApr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrow_apr" yaml:"borrow_apr"`
	StableBorrowApr          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=stable_borrow_apr,json=stableBorrowApr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_borrow_apr" yaml:"stable_borrow_apr"`
	UtilisationRatio         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=utilisation_ratio,json=utilisationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"utilisation_ratio" yaml:"utilisation_ratio"`
}

func (m *PoolAssetLBMapping) Reset()         { *m = PoolAssetLBMapping{} }
func (m *PoolAssetLBMapping) String() string { return proto.CompactTextString(m) }
func (*PoolAssetLBMapping) ProtoMessage()    {}
func (*PoolAssetLBMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{7}
}
func (m *PoolAssetLBMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAssetLBMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAssetLBMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAssetLBMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAssetLBMapping.Merge(m, src)
}
func (m *PoolAssetLBMapping) XXX_Size() int {
	return m.Size()
}
func (m *PoolAssetLBMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAssetLBMapping.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAssetLBMapping proto.InternalMessageInfo

func (m *PoolAssetLBMapping) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *PoolAssetLBMapping) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *PoolAssetLBMapping) GetLendIds() []uint64 {
	if m != nil {
		return m.LendIds
	}
	return nil
}

func (m *PoolAssetLBMapping) GetBorrowIds() []uint64 {
	if m != nil {
		return m.BorrowIds
	}
	return nil
}

type AssetRatesParams struct {
	AssetID              uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	UOptimal             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=u_optimal,json=uOptimal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"u_optimal" yaml:"u_optimal"`
	Base                 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base" yaml:"base"`
	Slope1               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slope1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slope1" yaml:"slope1"`
	Slope2               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slope2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slope2" yaml:"slope2"`
	EnableStableBorrow   bool                                   `protobuf:"varint,6,opt,name=enable_stable_borrow,json=enableStableBorrow,proto3" json:"enable_stable_borrow,omitempty" yaml:"enable_stable_borrow"`
	StableBase           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=stable_base,json=stableBase,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_base" yaml:"stable_base"`
	StableSlope1         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=stable_slope1,json=stableSlope1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_slope1" yaml:"stable_slope1"`
	StableSlope2         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=stable_slope2,json=stableSlope2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_slope2" yaml:"stable_slope2"`
	Ltv                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=ltv,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ltv" yaml:"ltv"`
	LiquidationThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=liquidation_threshold,json=liquidationThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_threshold" yaml:"liquidation_threshold"`
	LiquidationPenalty   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_penalty" yaml:"liquidation_penalty"`
	LiquidationBonus     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=liquidation_bonus,json=liquidationBonus,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_bonus" yaml:"liquidation_bonus"`
	ReserveFactor        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,14,opt,name=reserve_factor,json=reserveFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve_factor" yaml:"reserve_factor"`
	CAssetID             uint64                                 `protobuf:"varint,15,opt,name=c_asset_id,json=cAssetId,proto3" json:"c_asset_id,omitempty" yaml:"c_asset_id"`
}

func (m *AssetRatesParams) Reset()         { *m = AssetRatesParams{} }
func (m *AssetRatesParams) String() string { return proto.CompactTextString(m) }
func (*AssetRatesParams) ProtoMessage()    {}
func (*AssetRatesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{8}
}
func (m *AssetRatesParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetRatesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetRatesParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetRatesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetRatesParams.Merge(m, src)
}
func (m *AssetRatesParams) XXX_Size() int {
	return m.Size()
}
func (m *AssetRatesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetRatesParams.DiscardUnknown(m)
}

var xxx_messageInfo_AssetRatesParams proto.InternalMessageInfo

func (m *AssetRatesParams) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *AssetRatesParams) GetEnableStableBorrow() bool {
	if m != nil {
		return m.EnableStableBorrow
	}
	return false
}

func (m *AssetRatesParams) GetCAssetID() uint64 {
	if m != nil {
		return m.CAssetID
	}
	return 0
}

type ReserveBuybackAssetData struct {
	AssetID       uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	ReserveAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=reserve_amount,json=reserveAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserve_amount" yaml:"reserve_amount"`
	BuybackAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=buyback_amount,json=buybackAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"buyback_amount" yaml:"buyback_amount"`
}

func (m *ReserveBuybackAssetData) Reset()         { *m = ReserveBuybackAssetData{} }
func (m *ReserveBuybackAssetData) String() string { return proto.CompactTextString(m) }
func (*ReserveBuybackAssetData) ProtoMessage()    {}
func (*ReserveBuybackAssetData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{9}
}
func (m *ReserveBuybackAssetData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReserveBuybackAssetData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReserveBuybackAssetData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReserveBuybackAssetData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveBuybackAssetData.Merge(m, src)
}
func (m *ReserveBuybackAssetData) XXX_Size() int {
	return m.Size()
}
func (m *ReserveBuybackAssetData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveBuybackAssetData.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveBuybackAssetData proto.InternalMessageInfo

func (m *ReserveBuybackAssetData) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

type AuctionParams struct {
	AppId                  uint64                                 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	AuctionDurationSeconds uint64                                 `protobuf:"varint,2,opt,name=auction_duration_seconds,json=auctionDurationSeconds,proto3" json:"auction_duration_seconds,omitempty" yaml:"auction_duration_seconds"`
	Buffer                 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=buffer,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"buffer" yaml:"buffer"`
	Cusp                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=cusp,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"cusp" yaml:"cusp"`
	Step                   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=step,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"step" yaml:"step"`
	PriceFunctionType      uint64                                 `protobuf:"varint,6,opt,name=price_function_type,json=priceFunctionType,proto3" json:"price_function_type,omitempty" yaml:"price_function_type"`
	DutchId                uint64                                 `protobuf:"varint,7,opt,name=dutch_id,json=dutchId,proto3" json:"dutch_id,omitempty" yaml:"dutch_id"`
	BidDurationSeconds     uint64                                 `protobuf:"varint,8,opt,name=bid_duration_seconds,json=bidDurationSeconds,proto3" json:"bid_duration_seconds,omitempty" yaml:"bid_duration_seconds"`
}

func (m *AuctionParams) Reset()         { *m = AuctionParams{} }
func (m *AuctionParams) String() string { return proto.CompactTextString(m) }
func (*AuctionParams) ProtoMessage()    {}
func (*AuctionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{10}
}
func (m *AuctionParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionParams.Merge(m, src)
}
func (m *AuctionParams) XXX_Size() int {
	return m.Size()
}
func (m *AuctionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionParams.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionParams proto.InternalMessageInfo

func (m *AuctionParams) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *AuctionParams) GetAuctionDurationSeconds() uint64 {
	if m != nil {
		return m.AuctionDurationSeconds
	}
	return 0
}

func (m *AuctionParams) GetPriceFunctionType() uint64 {
	if m != nil {
		return m.PriceFunctionType
	}
	return 0
}

func (m *AuctionParams) GetDutchId() uint64 {
	if m != nil {
		return m.DutchId
	}
	return 0
}

func (m *AuctionParams) GetBidDurationSeconds() uint64 {
	if m != nil {
		return m.BidDurationSeconds
	}
	return 0
}

type BorrowInterestTracker struct {
	BorrowingId         uint64                                 `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty" yaml:"borrowing_id"`
	ReservePoolInterest github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=reserve_pool_interest,json=reservePoolInterest,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve_pool_interest" yaml:"reserve_pool_interest"`
}

func (m *BorrowInterestTracker) Reset()         { *m = BorrowInterestTracker{} }
func (m *BorrowInterestTracker) String() string { return proto.CompactTextString(m) }
func (*BorrowInterestTracker) ProtoMessage()    {}
func (*BorrowInterestTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{11}
}
func (m *BorrowInterestTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BorrowInterestTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BorrowInterestTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BorrowInterestTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowInterestTracker.Merge(m, src)
}
func (m *BorrowInterestTracker) XXX_Size() int {
	return m.Size()
}
func (m *BorrowInterestTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowInterestTracker.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowInterestTracker proto.InternalMessageInfo

func (m *BorrowInterestTracker) GetBorrowingId() uint64 {
	if m != nil {
		return m.BorrowingId
	}
	return 0
}

type LendRewardsTracker struct {
	LendingId          uint64                                 `protobuf:"varint,1,opt,name=lending_id,json=lendingId,proto3" json:"lending_id,omitempty" yaml:"lending_id"`
	RewardsAccumulated github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rewards_accumulated,json=rewardsAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rewards_accumulated" yaml:"interest_accumulated"`
}

func (m *LendRewardsTracker) Reset()         { *m = LendRewardsTracker{} }
func (m *LendRewardsTracker) String() string { return proto.CompactTextString(m) }
func (*LendRewardsTracker) ProtoMessage()    {}
func (*LendRewardsTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{12}
}
func (m *LendRewardsTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendRewardsTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendRewardsTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendRewardsTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendRewardsTracker.Merge(m, src)
}
func (m *LendRewardsTracker) XXX_Size() int {
	return m.Size()
}
func (m *LendRewardsTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_LendRewardsTracker.DiscardUnknown(m)
}

var xxx_messageInfo_LendRewardsTracker proto.InternalMessageInfo

func (m *LendRewardsTracker) GetLendingId() uint64 {
	if m != nil {
		return m.LendingId
	}
	return 0
}

type ModuleBalance struct {
	PoolID             uint64               `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	ModuleBalanceStats []ModuleBalanceStats `protobuf:"bytes,2,rep,name=module_balance_stats,json=moduleBalanceStats,proto3" json:"module_balance_stats" yaml:"module_balance_stats"`
}

func (m *ModuleBalance) Reset()         { *m = ModuleBalance{} }
func (m *ModuleBalance) String() string { return proto.CompactTextString(m) }
func (*ModuleBalance) ProtoMessage()    {}
func (*ModuleBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{13}
}
func (m *ModuleBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleBalance.Merge(m, src)
}
func (m *ModuleBalance) XXX_Size() int {
	return m.Size()
}
func (m *ModuleBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleBalance.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleBalance proto.InternalMessageInfo

func (m *ModuleBalance) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *ModuleBalance) GetModuleBalanceStats() []ModuleBalanceStats {
	if m != nil {
		return m.ModuleBalanceStats
	}
	return nil
}

type ModuleBalanceStats struct {
	AssetID uint64                                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	Balance github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=balance,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"balance" yaml:"balance"`
}

func (m *ModuleBalanceStats) Reset()         { *m = ModuleBalanceStats{} }
func (m *ModuleBalanceStats) String() string { return proto.CompactTextString(m) }
func (*ModuleBalanceStats) ProtoMessage()    {}
func (*ModuleBalanceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{14}
}
func (m *ModuleBalanceStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleBalanceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleBalanceStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleBalanceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleBalanceStats.Merge(m, src)
}
func (m *ModuleBalanceStats) XXX_Size() int {
	return m.Size()
}
func (m *ModuleBalanceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleBalanceStats.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleBalanceStats proto.InternalMessageInfo

func (m *ModuleBalanceStats) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *ModuleBalanceStats) GetBalance() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Balance
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

type ModBal struct {
	FundModuleBalance []FundModBal `protobuf:"bytes,1,rep,name=fund_module_balance,json=fundModuleBalance,proto3" json:"fund_module_balance" yaml:"fund_module_balance"`
}

func (m *ModBal) Reset()         { *m = ModBal{} }
func (m *ModBal) String() string { return proto.CompactTextString(m) }
func (*ModBal) ProtoMessage()    {}
func (*ModBal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{15}
}
func (m *ModBal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModBal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModBal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModBal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModBal.Merge(m, src)
}
func (m *ModBal) XXX_Size() int {
	return m.Size()
}
func (m *ModBal) XXX_DiscardUnknown() {
	xxx_messageInfo_ModBal.DiscardUnknown(m)
}

var xxx_messageInfo_ModBal proto.InternalMessageInfo

func (m *ModBal) GetFundModuleBalance() []FundModBal {
	if m != nil {
		return m.FundModuleBalance
	}
	return nil
}

type ReserveBal struct {
	FundReserveBalance []FundReserveBal `protobuf:"bytes,1,rep,name=fund_reserve_balance,json=fundReserveBalance,proto3" json:"fund_reserve_balance" yaml:"fund_reserve_balance"`
}

func (m *ReserveBal) Reset()         { *m = ReserveBal{} }
func (m *ReserveBal) String() string { return proto.CompactTextString(m) }
func (*ReserveBal) ProtoMessage()    {}
func (*ReserveBal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{16}
}
func (m *ReserveBal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReserveBal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReserveBal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReserveBal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveBal.Merge(m, src)
}
func (m *ReserveBal) XXX_Size() int {
	return m.Size()
}
func (m *ReserveBal) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveBal.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveBal proto.InternalMessageInfo

func (m *ReserveBal) GetFundReserveBalance() []FundReserveBal {
	if m != nil {
		return m.FundReserveBalance
	}
	return nil
}

type FundModBal struct {
	AssetID     uint64                                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	PoolID      uint64                                  `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	AmountIn    github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	DepositTime time.Time                               `protobuf:"bytes,4,opt,name=deposit_time,json=depositTime,proto3,stdtime" json:"deposit_time" yaml:"deposit_time"`
	Funder      string                                  `protobuf:"bytes,5,opt,name=funder,proto3" json:"funder,omitempty" yaml:"funder"`
}

func (m *FundModBal) Reset()         { *m = FundModBal{} }
func (m *FundModBal) String() string { return proto.CompactTextString(m) }
func (*FundModBal) ProtoMessage()    {}
func (*FundModBal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{17}
}
func (m *FundModBal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundModBal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundModBal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundModBal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundModBal.Merge(m, src)
}
func (m *FundModBal) XXX_Size() int {
	return m.Size()
}
func (m *FundModBal) XXX_DiscardUnknown() {
	xxx_messageInfo_FundModBal.DiscardUnknown(m)
}

var xxx_messageInfo_FundModBal proto.InternalMessageInfo

func (m *FundModBal) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *FundModBal) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *FundModBal) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *FundModBal) GetDepositTime() time.Time {
	if m != nil {
		return m.DepositTime
	}
	return time.Time{}
}

func (m *FundModBal) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

type FundReserveBal struct {
	AssetID     uint64                                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	AmountIn    github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	DepositTime time.Time                               `protobuf:"bytes,3,opt,name=deposit_time,json=depositTime,proto3,stdtime" json:"deposit_time" yaml:"deposit_time"`
	Funder      string                                  `protobuf:"bytes,4,opt,name=funder,proto3" json:"funder,omitempty" yaml:"funder"`
}

func (m *FundReserveBal) Reset()         { *m = FundReserveBal{} }
func (m *FundReserveBal) String() string { return proto.CompactTextString(m) }
func (*FundReserveBal) ProtoMessage()    {}
func (*FundReserveBal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{18}
}
func (m *FundReserveBal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundReserveBal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundReserveBal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundReserveBal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundReserveBal.Merge(m, src)
}
func (m *FundReserveBal) XXX_Size() int {
	return m.Size()
}
func (m *FundReserveBal) XXX_DiscardUnknown() {
	xxx_messageInfo_FundReserveBal.DiscardUnknown(m)
}

var xxx_messageInfo_FundReserveBal proto.InternalMessageInfo

func (m *FundReserveBal) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *FundReserveBal) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *FundReserveBal) GetDepositTime() time.Time {
	if m != nil {
		return m.DepositTime
	}
	return time.Time{}
}

func (m *FundReserveBal) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

type AllReserveStats struct {
	AssetID                        uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	AmountOutFromReserveToLenders  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount_out_from_reserve_to_lenders,json=amountOutFromReserveToLenders,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out_from_reserve_to_lenders" yaml:"amount_out_from_reserve_to_lenders"`
	AmountOutFromReserveForAuction github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount_out_from_reserve_for_auction,json=amountOutFromReserveForAuction,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out_from_reserve_for_auction" yaml:"amount_out_from_reserve_for_auction"`
	AmountInFromLiqPenalty         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount_in_from_liq_penalty,json=amountInFromLiqPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in_from_liq_penalty" yaml:"amount_in_from_liq_penalty"`
	AmountInFromRepayments         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount_in_from_repayments,json=amountInFromRepayments,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in_from_repayments" yaml:"amount_in_from_repayments"`
	TotalAmountOutToLenders        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=total_amount_out_to_lenders,json=totalAmountOutToLenders,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_amount_out_to_lenders" yaml:"total_amount_out_to_lenders"`
}

func (m *AllReserveStats) Reset()         { *m = AllReserveStats{} }
func (m *AllReserveStats) String() string { return proto.CompactTextString(m) }
func (*AllReserveStats) ProtoMessage()    {}
func (*AllReserveStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{19}
}
func (m *AllReserveStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllReserveStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllReserveStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllReserveStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllReserveStats.Merge(m, src)
}
func (m *AllReserveStats) XXX_Size() int {
	return m.Size()
}
func (m *AllReserveStats) XXX_DiscardUnknown() {
	xxx_messageInfo_AllReserveStats.DiscardUnknown(m)
}

var xxx_messageInfo_AllReserveStats proto.InternalMessageInfo

func (m *AllReserveStats) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

type AssetToPairSingleMapping struct {
	PoolID  uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	AssetID uint64 `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	PairID  uint64 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
}

func (m *AssetToPairSingleMapping) Reset()         { *m = AssetToPairSingleMapping{} }
func (m *AssetToPairSingleMapping) String() string { return proto.CompactTextString(m) }
func (*AssetToPairSingleMapping) ProtoMessage()    {}
func (*AssetToPairSingleMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{20}
}
func (m *AssetToPairSingleMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetToPairSingleMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetToPairSingleMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetToPairSingleMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetToPairSingleMapping.Merge(m, src)
}
func (m *AssetToPairSingleMapping) XXX_Size() int {
	return m.Size()
}
func (m *AssetToPairSingleMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetToPairSingleMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AssetToPairSingleMapping proto.InternalMessageInfo

func (m *AssetToPairSingleMapping) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *AssetToPairSingleMapping) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *AssetToPairSingleMapping) GetPairID() uint64 {
	if m != nil {
		return m.PairID
	}
	return 0
}

type PoolPairs struct {
	PoolID          uint64                  `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	ModuleName      string                  `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty" yaml:"module_name"`
	CPoolName       string                  `protobuf:"bytes,3,opt,name=cpool_name,json=cpoolName,proto3" json:"cpool_name,omitempty" yaml:"cpool_name"`
	AssetData       []*AssetDataPoolMapping `protobuf:"bytes,4,rep,name=asset_data,json=assetData,proto3" json:"asset_data,omitempty" yaml:"asset_data"`
	MinUsdValueLeft uint64                  `protobuf:"varint,5,opt,name=min_usd_value_left,json=minUsdValueLeft,proto3" json:"min_usd_value_left,omitempty" yaml:"min_usd_value_left"`
}

func (m *PoolPairs) Reset()         { *m = PoolPairs{} }
func (m *PoolPairs) String() string { return proto.CompactTextString(m) }
func (*PoolPairs) ProtoMessage()    {}
func (*PoolPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{21}
}
func (m *PoolPairs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolPairs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolPairs.Merge(m, src)
}
func (m *PoolPairs) XXX_Size() int {
	return m.Size()
}
func (m *PoolPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolPairs.DiscardUnknown(m)
}

var xxx_messageInfo_PoolPairs proto.InternalMessageInfo

func (m *PoolPairs) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *PoolPairs) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *PoolPairs) GetCPoolName() string {
	if m != nil {
		return m.CPoolName
	}
	return ""
}

func (m *PoolPairs) GetAssetData() []*AssetDataPoolMapping {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *PoolPairs) GetMinUsdValueLeft() uint64 {
	if m != nil {
		return m.MinUsdValueLeft
	}
	return 0
}

type PoolInterestData struct {
	AssetID      uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	LendInterest github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=lend_interest,json=lendInterest,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"lend_interest" yaml:"lend_interest"`
}

func (m *PoolInterestData) Reset()         { *m = PoolInterestData{} }
func (m *PoolInterestData) String() string { return proto.CompactTextString(m) }
func (*PoolInterestData) ProtoMessage()    {}
func (*PoolInterestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{22}
}
func (m *PoolInterestData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInterestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInterestData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInterestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInterestData.Merge(m, src)
}
func (m *PoolInterestData) XXX_Size() int {
	return m.Size()
}
func (m *PoolInterestData) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInterestData.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInterestData proto.InternalMessageInfo

func (m *PoolInterestData) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

type PoolInterest struct {
	PoolID           uint64             `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	PoolInterestData []PoolInterestData `protobuf:"bytes,2,rep,name=pool_interest_data,json=poolInterestData,proto3" json:"pool_interest_data" yaml:"pool_interest_data"`
}

func (m *PoolInterest) Reset()         { *m = PoolInterest{} }
func (m *PoolInterest) String() string { return proto.CompactTextString(m) }
func (*PoolInterest) ProtoMessage()    {}
func (*PoolInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{23}
}
func (m *PoolInterest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInterest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInterest.Merge(m, src)
}
func (m *PoolInterest) XXX_Size() int {
	return m.Size()
}
func (m *PoolInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInterest.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInterest proto.InternalMessageInfo

func (m *PoolInterest) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *PoolInterest) GetPoolInterestData() []PoolInterestData {
	if m != nil {
		return m.PoolInterestData
	}
	return nil
}

type PoolInterestDataB struct {
	AssetID        uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	BorrowInterest github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=borrow_interest,json=borrowInterest,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"borrow_interest" yaml:"borrow_interest"`
}

func (m *PoolInterestDataB) Reset()         { *m = PoolInterestDataB{} }
func (m *PoolInterestDataB) String() string { return proto.CompactTextString(m) }
func (*PoolInterestDataB) ProtoMessage()    {}
func (*PoolInterestDataB) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{24}
}
func (m *PoolInterestDataB) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInterestDataB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInterestDataB.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInterestDataB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInterestDataB.Merge(m, src)
}
func (m *PoolInterestDataB) XXX_Size() int {
	return m.Size()
}
func (m *PoolInterestDataB) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInterestDataB.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInterestDataB proto.InternalMessageInfo

func (m *PoolInterestDataB) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

type PoolInterestB struct {
	PoolID           uint64              `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	PoolInterestData []PoolInterestDataB `protobuf:"bytes,2,rep,name=pool_interest_data,json=poolInterestData,proto3" json:"pool_interest_data" yaml:"pool_interest_data"`
}

func (m *PoolInterestB) Reset()         { *m = PoolInterestB{} }
func (m *PoolInterestB) String() string { return proto.CompactTextString(m) }
func (*PoolInterestB) ProtoMessage()    {}
func (*PoolInterestB) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{25}
}
func (m *PoolInterestB) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInterestB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInterestB.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInterestB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInterestB.Merge(m, src)
}
func (m *PoolInterestB) XXX_Size() int {
	return m.Size()
}
func (m *PoolInterestB) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInterestB.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInterestB proto.InternalMessageInfo

func (m *PoolInterestB) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *PoolInterestB) GetPoolInterestData() []PoolInterestDataB {
	if m != nil {
		return m.PoolInterestData
	}
	return nil
}

type AssetRatesPoolPairs struct {
	AssetID              uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	UOptimal             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=u_optimal,json=uOptimal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"u_optimal" yaml:"u_optimal"`
	Base                 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base" yaml:"base"`
	Slope1               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slope1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slope1" yaml:"slope1"`
	Slope2               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slope2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slope2" yaml:"slope2"`
	EnableStableBorrow   bool                                   `protobuf:"varint,6,opt,name=enable_stable_borrow,json=enableStableBorrow,proto3" json:"enable_stable_borrow,omitempty" yaml:"enable_stable_borrow"`
	StableBase           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=stable_base,json=stableBase,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_base" yaml:"stable_base"`
	StableSlope1         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=stable_slope1,json=stableSlope1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_slope1" yaml:"stable_slope1"`
	StableSlope2         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=stable_slope2,json=stableSlope2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stable_slope2" yaml:"stable_slope2"`
	Ltv                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=ltv,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ltv" yaml:"ltv"`
	LiquidationThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=liquidation_threshold,json=liquidationThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_threshold" yaml:"liquidation_threshold"`
	LiquidationPenalty   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_penalty" yaml:"liquidation_penalty"`
	LiquidationBonus     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=liquidation_bonus,json=liquidationBonus,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_bonus" yaml:"liquidation_bonus"`
	ReserveFactor        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,14,opt,name=reserve_factor,json=reserveFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserve_factor" yaml:"reserve_factor"`
	CAssetID             uint64                                 `protobuf:"varint,15,opt,name=c_asset_id,json=cAssetId,proto3" json:"c_asset_id,omitempty" yaml:"c_asset_id"`
	ModuleName           string                                 `protobuf:"bytes,16,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty" yaml:"module_name"`
	CPoolName            string                                 `protobuf:"bytes,17,opt,name=cpool_name,json=cpoolName,proto3" json:"cpool_name,omitempty" yaml:"cpool_name"`
	AssetData            []*AssetDataPoolMapping                `protobuf:"bytes,18,rep,name=asset_data,json=assetData,proto3" json:"asset_data,omitempty" yaml:"asset_data"`
	MinUsdValueLeft      uint64                                 `protobuf:"varint,19,opt,name=min_usd_value_left,json=minUsdValueLeft,proto3" json:"min_usd_value_left,omitempty" yaml:"min_usd_value_left"`
}

func (m *AssetRatesPoolPairs) Reset()         { *m = AssetRatesPoolPairs{} }
func (m *AssetRatesPoolPairs) String() string { return proto.CompactTextString(m) }
func (*AssetRatesPoolPairs) ProtoMessage()    {}
func (*AssetRatesPoolPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{26}
}
func (m *AssetRatesPoolPairs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetRatesPoolPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetRatesPoolPairs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetRatesPoolPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetRatesPoolPairs.Merge(m, src)
}
func (m *AssetRatesPoolPairs) XXX_Size() int {
	return m.Size()
}
func (m *AssetRatesPoolPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetRatesPoolPairs.DiscardUnknown(m)
}

var xxx_messageInfo_AssetRatesPoolPairs proto.InternalMessageInfo

func (m *AssetRatesPoolPairs) GetAssetID() uint64 {
	if m != nil {
		return m.AssetID
	}
	return 0
}

func (m *AssetRatesPoolPairs) GetEnableStableBorrow() bool {
	if m != nil {
		return m.EnableStableBorrow
	}
	return false
}

func (m *AssetRatesPoolPairs) GetCAssetID() uint64 {
	if m != nil {
		return m.CAssetID
	}
	return 0
}

func (m *AssetRatesPoolPairs) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *AssetRatesPoolPairs) GetCPoolName() string {
	if m != nil {
		return m.CPoolName
	}
	return ""
}

func (m *AssetRatesPoolPairs) GetAssetData() []*AssetDataPoolMapping {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *AssetRatesPoolPairs) GetMinUsdValueLeft() uint64 {
	if m != nil {
		return m.MinUsdValueLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*LendAsset)(nil), "comdex.lend.v1beta1.LendAsset")
	proto.RegisterType((*BorrowAsset)(nil), "comdex.lend.v1beta1.BorrowAsset")
	proto.RegisterType((*Pool)(nil), "comdex.lend.v1beta1.Pool")
	proto.RegisterType((*UserAssetLendBorrowMapping)(nil), "comdex.lend.v1beta1.UserAssetLendBorrowMapping")
	proto.RegisterType((*AssetDataPoolMapping)(nil), "comdex.lend.v1beta1.AssetDataPoolMapping")
	proto.RegisterType((*Extended_Pair)(nil), "comdex.lend.v1beta1.Extended_Pair")
	proto.RegisterType((*AssetToPairMapping)(nil), "comdex.lend.v1beta1.AssetToPairMapping")
	proto.RegisterType((*PoolAssetLBMapping)(nil), "comdex.lend.v1beta1.PoolAssetLBMapping")
	proto.RegisterType((*AssetRatesParams)(nil), "comdex.lend.v1beta1.AssetRatesParams")
	proto.RegisterType((*ReserveBuybackAssetData)(nil), "comdex.lend.v1beta1.ReserveBuybackAssetData")
	proto.RegisterType((*AuctionParams)(nil), "comdex.lend.v1beta1.AuctionParams")
	proto.RegisterType((*BorrowInterestTracker)(nil), "comdex.lend.v1beta1.Borrow_interest_tracker")
	proto.RegisterType((*LendRewardsTracker)(nil), "comdex.lend.v1beta1.Lend_rewards_tracker")
	proto.RegisterType((*ModuleBalance)(nil), "comdex.lend.v1beta1.ModuleBalance")
	proto.RegisterType((*ModuleBalanceStats)(nil), "comdex.lend.v1beta1.ModuleBalanceStats")
	proto.RegisterType((*ModBal)(nil), "comdex.lend.v1beta1.ModBal")
	proto.RegisterType((*ReserveBal)(nil), "comdex.lend.v1beta1.ReserveBal")
	proto.RegisterType((*FundModBal)(nil), "comdex.lend.v1beta1.FundModBal")
	proto.RegisterType((*FundReserveBal)(nil), "comdex.lend.v1beta1.FundReserveBal")
	proto.RegisterType((*AllReserveStats)(nil), "comdex.lend.v1beta1.AllReserveStats")
	proto.RegisterType((*AssetToPairSingleMapping)(nil), "comdex.lend.v1beta1.AssetToPairSingleMapping")
	proto.RegisterType((*PoolPairs)(nil), "comdex.lend.v1beta1.PoolPairs")
	proto.RegisterType((*PoolInterestData)(nil), "comdex.lend.v1beta1.PoolInterestData")
	proto.RegisterType((*PoolInterest)(nil), "comdex.lend.v1beta1.PoolInterest")
	proto.RegisterType((*PoolInterestDataB)(nil), "comdex.lend.v1beta1.PoolInterestDataB")
	proto.RegisterType((*PoolInterestB)(nil), "comdex.lend.v1beta1.PoolInterestB")
	proto.RegisterType((*AssetRatesPoolPairs)(nil), "comdex.lend.v1beta1.AssetRatesPoolPairs")
}

func init() { proto.RegisterFile("comdex/lend/v1beta1/lend.proto", fileDescriptor_b87bb4bef8334ddd) }

var fileDescriptor_b87bb4bef8334ddd = []byte{
	// 2982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5b, 0x4d, 0x6c, 0x24, 0x47,
	0xf5, 0xdf, 0x1e, 0x7b, 0xed, 0x99, 0x37, 0xb6, 0xd7, 0xae, 0xb1, 0xb3, 0x13, 0x6f, 0xe2, 0xd9,
	0xd4, 0xfe, 0xff, 0x89, 0x03, 0xca, 0x58, 0x6b, 0x92, 0x03, 0x51, 0xa2, 0xe0, 0xf1, 0xee, 0x26,
	0x93, 0x78, 0x37, 0x49, 0xad, 0x03, 0x02, 0x02, 0xad, 0x9e, 0xe9, 0xb2, 0xb7, 0xb5, 0x3d, 0xdd,
	0xbd, 0xdd, 0x3d, 0xde, 0x35, 0x90, 0x80, 0x12, 0x09, 0x21, 0x20, 0x22, 0x70, 0xe5, 0x8e, 0xc4,
	0x1d, 0x2e, 0x48, 0x1c, 0x10, 0x52, 0xc8, 0x81, 0x43, 0xf8, 0x38, 0x44, 0x39, 0x0c, 0xc8, 0x39,
	0x20, 0x71, 0xdc, 0x03, 0x07, 0x4e, 0xa8, 0xaa, 0x5e, 0x7f, 0xcd, 0xb4, 0xbd, 0xdb, 0x9e, 0x24,
	0x0b, 0xd2, 0x9e, 0x3c, 0xf5, 0xaa, 0xea, 0xf7, 0x5e, 0xbd, 0xf7, 0xea, 0xd5, 0x7b, 0x55, 0x6d,
	0x58, 0xe9, 0xba, 0x3d, 0x93, 0xdf, 0x5a, 0xb3, 0xb9, 0x63, 0xae, 0xed, 0x9d, 0xef, 0xf0, 0xd0,
	0x38, 0x2f, 0x1b, 0x4d, 0xcf, 0x77, 0x43, 0x97, 0xd4, 0x54, 0x7f, 0x53, 0x92, 0xb0, 0x7f, 0x79,
	0x71, 0xd7, 0xdd, 0x75, 0x65, 0xff, 0x9a, 0xf8, 0xa5, 0x86, 0x2e, 0x37, 0x76, 0x5d, 0x77, 0xd7,
	0xe6, 0x6b, 0xb2, 0xd5, 0xe9, 0xef, 0xac, 0x85, 0x56, 0x8f, 0x07, 0xa1, 0xd1, 0xf3, 0x70, 0xc0,
	0x4a, 0xd7, 0x0d, 0x7a, 0x6e, 0xb0, 0xd6, 0x31, 0x02, 0x1e, 0xf3, 0xea, 0xba, 0x96, 0xa3, 0xfa,
	0xe9, 0xdb, 0x65, 0xa8, 0x6c, 0x71, 0xc7, 0xdc, 0x08, 0x02, 0x1e, 0x92, 0xa7, 0x01, 0x04, 0x53,
	0xcb, 0xd9, 0xd5, 0x2d, 0xb3, 0xae, 0x9d, 0xd5, 0x56, 0x27, 0x5b, 0x67, 0x0e, 0x06, 0x8d, 0x52,
	0xfb, 0xc2, 0xed, 0x41, 0x63, 0x61, 0xdf, 0xe8, 0xd9, 0x4f, 0xd3, 0x64, 0x04, 0x65, 0x15, 0x6c,
	0xb4, 0x4d, 0xf2, 0x45, 0x28, 0x1b, 0x02, 0x44, 0xcc, 0x2c, 0xc9, 0x99, 0x2b, 0x07, 0x83, 0xc6,
	0xb4, 0x04, 0x96, 0xd3, 0x4f, 0xa9, 0xe9, 0xd1, 0x20, 0xca, 0xa6, 0xe5, 0xcf, 0xb6, 0x49, 0x9e,
	0x82, 0x69, 0xcf, 0x75, 0x6d, 0x31, 0x73, 0x42, 0xce, 0x7c, 0xe8, 0x60, 0xd0, 0x98, 0x7a, 0xc5,
	0x75, 0x6d, 0x39, 0x71, 0x4e, 0x4d, 0xc4, 0x21, 0x94, 0x4d, 0x89, 0x5f, 0x6d, 0x93, 0x3c, 0x0a,
	0x27, 0xdd, 0x9b, 0x0e, 0xf7, 0xeb, 0x93, 0x67, 0xb5, 0xd5, 0x4a, 0x6b, 0xfe, 0xf6, 0xa0, 0x31,
	0xa3, 0x86, 0x4a, 0x32, 0x65, 0xaa, 0x9b, 0x7c, 0x1b, 0x2a, 0x46, 0xcf, 0xed, 0x3b, 0xa1, 0x6e,
	0x39, 0xf5, 0x93, 0x67, 0xb5, 0xd5, 0xea, 0xfa, 0x83, 0x4d, 0xa5, 0x97, 0xa6, 0xd0, 0x4b, 0xa4,
	0xe3, 0xe6, 0xa6, 0x6b, 0x39, 0xad, 0xcd, 0xf7, 0x07, 0x8d, 0x13, 0xb7, 0x07, 0x8d, 0x79, 0x14,
	0x37, 0x9a, 0x49, 0xff, 0x3d, 0x68, 0x3c, 0xb6, 0x6b, 0x85, 0xd7, 0xfa, 0x9d, 0x66, 0xd7, 0xed,
	0xad, 0xa1, 0x62, 0xd5, 0x9f, 0x27, 0x02, 0xf3, 0xfa, 0x5a, 0xb8, 0xef, 0xf1, 0x40, 0x82, 0xb0,
	0xb2, 0x9a, 0xd6, 0x76, 0xc8, 0x37, 0x61, 0x26, 0x52, 0x98, 0xb0, 0x4d, 0x7d, 0x4a, 0xf2, 0x5f,
	0x6e, 0x2a, 0xc3, 0x35, 0x23, 0xc3, 0x35, 0xb7, 0x23, 0xc3, 0xb5, 0x1a, 0x28, 0x40, 0x2d, 0xab,
	0x6e, 0x31, 0x9b, 0xbe, 0xfb, 0xb7, 0x86, 0xc6, 0xaa, 0x48, 0x12, 0x53, 0xc8, 0x77, 0xa0, 0x66,
	0xec, 0x19, 0x96, 0x6d, 0x74, 0x6c, 0xae, 0x87, 0xae, 0xde, 0x71, 0x7d, 0xdf, 0xbd, 0x59, 0x9f,
	0x96, 0x2a, 0xd9, 0x12, 0x50, 0x1f, 0x0d, 0x1a, 0x8f, 0xde, 0x85, 0xdc, 0x6d, 0x27, 0xbc, 0x3d,
	0x68, 0x2c, 0xe3, 0xaa, 0x47, 0x21, 0x29, 0x5b, 0x88, 0xa9, 0xdb, 0x6e, 0x4b, 0xd2, 0xc8, 0x79,
	0x98, 0x32, 0x3c, 0x4f, 0x18, 0xae, 0x2c, 0x0d, 0xb7, 0x7c, 0x30, 0x68, 0x9c, 0xdc, 0xf0, 0x3c,
	0x69, 0xb7, 0x59, 0xc4, 0x92, 0x03, 0x28, 0x3b, 0x69, 0x78, 0x5e, 0xdb, 0x24, 0xd7, 0x60, 0x66,
	0xd7, 0x76, 0x3b, 0x86, 0xad, 0x5b, 0x8e, 0xc9, 0x6f, 0xd5, 0x2b, 0x52, 0xd2, 0x8b, 0x05, 0x24,
	0xbd, 0xc0, 0xbb, 0x89, 0x7a, 0xd2, 0x58, 0x94, 0x55, 0x55, 0xb3, 0x2d, 0x5a, 0xe4, 0x16, 0x2c,
	0xd9, 0x46, 0x20, 0x6c, 0x17, 0x72, 0xdf, 0xe8, 0x86, 0x96, 0xeb, 0x28, 0x1b, 0xc0, 0x1d, 0x6d,
	0xb0, 0x8a, 0x36, 0x78, 0x08, 0x6d, 0x90, 0x07, 0xa3, 0x8c, 0x51, 0x13, 0x7d, 0xed, 0xa4, 0x4b,
	0x1a, 0x65, 0x03, 0xa0, 0x2b, 0xdd, 0xd5, 0x31, 0x7a, 0xbc, 0x5e, 0x95, 0x2b, 0xa4, 0x07, 0x83,
	0x46, 0x65, 0x53, 0x38, 0xf5, 0x15, 0xa3, 0xc7, 0x93, 0xed, 0x94, 0x0c, 0xa4, 0xac, 0x22, 0x1b,
	0xa2, 0x9f, 0x5c, 0x87, 0xd9, 0xd0, 0x0d, 0x0d, 0x5b, 0xf7, 0xf9, 0x4d, 0xc3, 0x37, 0x83, 0xfa,
	0x8c, 0x44, 0xb9, 0x54, 0xd8, 0xa2, 0x8b, 0x8a, 0x4d, 0x06, 0x8c, 0xb2, 0x19, 0xd9, 0x66, 0xd8,
	0xfc, 0x57, 0x15, 0xaa, 0xca, 0xa2, 0x2a, 0x0e, 0x7c, 0x09, 0x66, 0x94, 0xd1, 0x33, 0x91, 0xe0,
	0xe1, 0x38, 0x12, 0xa0, 0xee, 0xd3, 0x63, 0x28, 0xab, 0xc6, 0xcd, 0xb6, 0x29, 0x34, 0x90, 0x8a,
	0x24, 0x2a, 0x1e, 0x48, 0x0d, 0x6c, 0x61, 0xc0, 0xb8, 0x73, 0x40, 0xb9, 0x08, 0xf3, 0x56, 0xa0,
	0x07, 0xa1, 0x74, 0x43, 0x74, 0x6b, 0x11, 0x1e, 0xca, 0xad, 0x33, 0xb7, 0x07, 0x8d, 0xd3, 0x6a,
	0xee, 0xf0, 0x08, 0xca, 0xe6, 0xac, 0xe0, 0xaa, 0xa4, 0xa0, 0x8b, 0x8a, 0xe0, 0x62, 0x58, 0xbe,
	0x10, 0x63, 0x32, 0x15, 0x5c, 0x0c, 0xcb, 0xcf, 0x04, 0x17, 0x35, 0x44, 0x04, 0x17, 0xd1, 0x63,
	0xde, 0xdb, 0xa0, 0xf1, 0x26, 0x00, 0x42, 0xb8, 0xfd, 0x10, 0x43, 0xc6, 0x11, 0xdc, 0x2f, 0x20,
	0xf7, 0x85, 0x0c, 0x77, 0xb7, 0x1f, 0x16, 0x62, 0x8f, 0xeb, 0x7d, 0xb9, 0x1f, 0x92, 0x9f, 0x6b,
	0xb0, 0xd8, 0xf1, 0x2d, 0x73, 0x97, 0x9b, 0xba, 0x8a, 0xd7, 0xaa, 0x4f, 0x86, 0x95, 0x23, 0x45,
	0xb9, 0x82, 0xa2, 0x9c, 0x41, 0x0f, 0xc9, 0x01, 0x29, 0x24, 0x14, 0x41, 0x04, 0xe9, 0x97, 0x1b,
	0x72, 0x3e, 0x31, 0x61, 0x2e, 0xf1, 0x3c, 0xb9, 0xa1, 0xcb, 0x77, 0xdc, 0xd0, 0x8f, 0xa0, 0x5c,
	0x4b, 0xc3, 0x9e, 0x9b, 0xec, 0xe4, 0xd9, 0x98, 0x28, 0xf7, 0xf0, 0x3e, 0x90, 0x8c, 0x67, 0xe9,
	0xbe, 0x11, 0x72, 0x8c, 0x56, 0x2f, 0x15, 0x8e, 0x56, 0x0f, 0x2a, 0xbe, 0xa3, 0x88, 0x94, 0xcd,
	0x07, 0x29, 0x77, 0x65, 0x46, 0xc8, 0xc9, 0xf7, 0x34, 0x58, 0x94, 0xd1, 0x86, 0x07, 0xa1, 0x6e,
	0x74, 0xbb, 0xfd, 0x5e, 0xdf, 0x36, 0x42, 0x6e, 0xca, 0xc0, 0x55, 0x69, 0x5d, 0x2e, 0xcc, 0x1d,
	0xad, 0x91, 0x87, 0x49, 0x59, 0x2d, 0x22, 0x6f, 0x24, 0xd4, 0x91, 0x28, 0x5d, 0xfd, 0xd4, 0xa2,
	0xf4, 0x77, 0x61, 0xd1, 0xe7, 0x01, 0xf7, 0xf7, 0xb8, 0x9e, 0xe1, 0x38, 0x33, 0xde, 0x5a, 0xf3,
	0x30, 0x29, 0x23, 0x48, 0x7e, 0xfe, 0x6e, 0x8e, 0x89, 0xd9, 0xcf, 0xf6, 0x98, 0x98, 0x3b, 0xce,
	0x31, 0xf1, 0x2c, 0xcc, 0x5a, 0x81, 0x6e, 0x5b, 0x37, 0xfa, 0x96, 0x29, 0x5d, 0xe4, 0x94, 0x8c,
	0x90, 0xf5, 0x24, 0xf0, 0x67, 0xba, 0x29, 0x9b, 0xb1, 0x82, 0xad, 0xa4, 0xf9, 0xeb, 0x12, 0x4c,
	0x0a, 0x5e, 0xe9, 0x14, 0x4c, 0x2b, 0x90, 0x82, 0x5d, 0x84, 0x6a, 0xcf, 0x35, 0xfb, 0x36, 0x57,
	0x4b, 0x28, 0xc9, 0x25, 0xfc, 0xdf, 0xc1, 0xa0, 0x01, 0x97, 0x25, 0x19, 0xd7, 0x40, 0xd4, 0xf4,
	0xd4, 0x50, 0xca, 0xa0, 0x17, 0x8f, 0x18, 0x52, 0xc4, 0xc4, 0x71, 0x14, 0x61, 0x03, 0xa8, 0x20,
	0x63, 0x1a, 0xa1, 0x51, 0x9f, 0x3c, 0x3b, 0xb1, 0x5a, 0x5d, 0x7f, 0xbc, 0x99, 0x93, 0x49, 0x37,
	0x65, 0x28, 0xb9, 0x60, 0x84, 0x86, 0xc0, 0xbe, 0x6c, 0x78, 0x9e, 0xe5, 0xec, 0x2a, 0x6e, 0x71,
	0x4f, 0x2a, 0x96, 0xc6, 0x98, 0x94, 0x55, 0x8c, 0xa8, 0x9f, 0xfe, 0x49, 0x83, 0xe5, 0xd7, 0x02,
	0xee, 0xcb, 0x19, 0xe2, 0x48, 0x53, 0xbb, 0x17, 0xd1, 0x92, 0xcc, 0x54, 0x3b, 0x3a, 0x33, 0xfd,
	0x3c, 0x4c, 0x0b, 0xd1, 0x92, 0x23, 0x92, 0x24, 0xba, 0xc6, 0x0e, 0xca, 0xa6, 0xc4, 0xaf, 0xb6,
	0x29, 0x06, 0x67, 0xb3, 0x64, 0x72, 0x84, 0x61, 0xce, 0x43, 0x05, 0x83, 0x8c, 0x3c, 0xf7, 0x26,
	0x56, 0x27, 0x5b, 0x8b, 0xc9, 0xf9, 0x14, 0x77, 0x51, 0x56, 0x56, 0xbf, 0xdb, 0x26, 0x7d, 0xab,
	0x04, 0x8b, 0x79, 0xba, 0xc9, 0x64, 0xf6, 0x5a, 0xb1, 0xcc, 0xfe, 0x25, 0x20, 0x8a, 0x1a, 0xfa,
	0x86, 0x13, 0x58, 0xa1, 0x2e, 0x76, 0x2a, 0xae, 0xf5, 0xe1, 0x24, 0x2c, 0x8e, 0x8e, 0xa1, 0x6c,
	0x5e, 0x12, 0xb7, 0x15, 0x6d, 0x7b, 0xdf, 0xe3, 0xa4, 0x03, 0x10, 0xf4, 0x3d, 0xcf, 0xde, 0xd7,
	0xbb, 0x86, 0x87, 0x5e, 0xb2, 0x59, 0x38, 0x3e, 0xa0, 0x61, 0x13, 0x24, 0xca, 0x2a, 0xaa, 0xb1,
	0x69, 0x78, 0xf4, 0x1f, 0x25, 0x98, 0xbd, 0x78, 0x2b, 0xe4, 0x8e, 0xc9, 0x4d, 0x5d, 0x24, 0x09,
	0x64, 0x0e, 0x4a, 0xd1, 0xba, 0x59, 0xc9, 0x32, 0x49, 0x33, 0xd6, 0x86, 0x83, 0x0b, 0xa9, 0x8d,
	0xa8, 0xc0, 0x89, 0x55, 0xe0, 0x08, 0x4b, 0x28, 0xaa, 0x38, 0xca, 0x95, 0xe1, 0x52, 0x96, 0x88,
	0xbb, 0x28, 0x53, 0xb0, 0xe2, 0xf8, 0x7d, 0x46, 0x6e, 0x6a, 0x19, 0x48, 0x74, 0x61, 0x4f, 0x99,
	0xb8, 0x0c, 0x6f, 0xea, 0xa4, 0x9b, 0xb2, 0xaa, 0x15, 0xc8, 0xd8, 0x22, 0xb7, 0xf2, 0x57, 0x61,
	0x21, 0x46, 0xd5, 0x23, 0x8f, 0x39, 0x29, 0x19, 0x37, 0x0f, 0x06, 0x8d, 0xb9, 0x0d, 0x64, 0x13,
	0x6f, 0xee, 0xfa, 0x90, 0x28, 0x7a, 0xec, 0x4d, 0x73, 0x46, 0x7a, 0xac, 0x49, 0x5e, 0x04, 0xd2,
	0xb3, 0x1c, 0xbd, 0x1f, 0x98, 0xfa, 0x9e, 0x61, 0xf7, 0xb9, 0x6e, 0xf3, 0x1d, 0x95, 0x9f, 0x64,
	0xcc, 0x39, 0x3a, 0x86, 0xb2, 0x53, 0x3d, 0xcb, 0x79, 0x2d, 0x30, 0xbf, 0x2c, 0x48, 0x5b, 0x82,
	0xf2, 0x5b, 0x0d, 0x88, 0x14, 0x65, 0xdb, 0x15, 0x7a, 0x8e, 0x9c, 0xed, 0x98, 0x81, 0x68, 0xcc,
	0xea, 0x13, 0x13, 0xc4, 0x09, 0xb9, 0x51, 0xee, 0x2a, 0x41, 0xa4, 0x3f, 0xae, 0x00, 0x11, 0x62,
	0xa9, 0x10, 0xd0, 0xba, 0x77, 0xf2, 0x37, 0xa1, 0x8c, 0xb1, 0x22, 0xc0, 0x05, 0xa4, 0x1c, 0x32,
	0xea, 0xa1, 0x6c, 0x5a, 0x85, 0x91, 0x80, 0x3c, 0x09, 0x10, 0xef, 0xff, 0x00, 0x63, 0xc3, 0x52,
	0xb2, 0x31, 0x92, 0x3e, 0xca, 0x2a, 0x51, 0x70, 0x08, 0x88, 0x03, 0x73, 0xaa, 0x84, 0x50, 0x24,
	0xae, 0x5c, 0xaa, 0xd2, 0x7a, 0xbe, 0x70, 0x41, 0xb2, 0x94, 0x2e, 0x48, 0x22, 0x34, 0xca, 0x54,
	0xb9, 0xd3, 0xc2, 0x36, 0x79, 0x4b, 0x83, 0x25, 0x35, 0x24, 0x93, 0x33, 0x71, 0x53, 0xba, 0x5b,
	0x45, 0x25, 0x9a, 0x85, 0xf8, 0x3e, 0x94, 0xe6, 0x3b, 0x04, 0x4a, 0x59, 0x4d, 0xd2, 0xd3, 0x95,
	0x03, 0x37, 0x45, 0xc4, 0x51, 0xc3, 0x85, 0xee, 0xb0, 0xa6, 0xde, 0x2c, 0xcc, 0x78, 0x21, 0xcd,
	0x58, 0x20, 0x51, 0x56, 0x91, 0x0d, 0x71, 0x70, 0x90, 0x9f, 0x6a, 0xb0, 0xac, 0xba, 0x72, 0x53,
	0xbe, 0xb2, 0x64, 0x7a, 0xb5, 0x30, 0xd3, 0x47, 0xd2, 0x4c, 0xf3, 0x13, 0xbf, 0xba, 0xec, 0x6c,
	0xe7, 0x64, 0x7f, 0xaf, 0xa3, 0x4b, 0x19, 0x9e, 0x8f, 0x19, 0xef, 0x46, 0xe1, 0x38, 0x9b, 0x76,
	0x40, 0xc3, 0xf3, 0xd1, 0x01, 0x37, 0x3c, 0x5f, 0x68, 0x15, 0x9d, 0x4c, 0xe0, 0xc3, 0x78, 0x71,
	0x3c, 0x41, 0x8a, 0xdd, 0x55, 0xf0, 0xd8, 0x83, 0x85, 0x6c, 0xae, 0x2d, 0x58, 0xa9, 0x24, 0xf6,
	0xc5, 0xc2, 0xac, 0xea, 0x79, 0xc9, 0xbb, 0xe4, 0x78, 0x2a, 0x9d, 0xbb, 0x0b, 0xbe, 0x37, 0x61,
	0xa1, 0x1f, 0x5a, 0xb6, 0x15, 0x18, 0x32, 0x01, 0xf4, 0xc5, 0x1f, 0x4c, 0x65, 0x8f, 0xcd, 0x77,
	0x04, 0x90, 0xb2, 0xf9, 0x14, 0x8d, 0x49, 0xd2, 0x7b, 0x55, 0x98, 0x97, 0xd1, 0x42, 0x54, 0x10,
	0xc1, 0x2b, 0x86, 0x6f, 0xf4, 0x82, 0x71, 0x4e, 0x6e, 0x1d, 0x2a, 0x7d, 0xdd, 0xf5, 0x42, 0xab,
	0x67, 0xd8, 0x98, 0xd7, 0xb5, 0x0a, 0x2f, 0x00, 0x0f, 0xb9, 0x18, 0x88, 0xb2, 0x72, 0xff, 0x65,
	0xf5, 0x93, 0xbc, 0x0a, 0x93, 0xa2, 0x7c, 0xc4, 0x73, 0xfc, 0xd9, 0xc2, 0xd8, 0x55, 0xb4, 0xbf,
	0x11, 0x70, 0xca, 0x24, 0x14, 0xf9, 0x0a, 0x4c, 0x05, 0xb6, 0xeb, 0xf1, 0xf3, 0x78, 0x23, 0xf8,
	0x5c, 0x61, 0x50, 0xbc, 0xb2, 0x52, 0x28, 0x94, 0x21, 0x5c, 0x0c, 0xbc, 0x8e, 0x41, 0x6f, 0x3c,
	0xe0, 0xf5, 0x08, 0x78, 0x9d, 0xbc, 0x0a, 0x8b, 0xdc, 0x91, 0x5e, 0x95, 0xbd, 0xe7, 0x98, 0x92,
	0x07, 0x7e, 0x23, 0x29, 0x67, 0xf2, 0x46, 0x51, 0x46, 0x14, 0x39, 0x73, 0xdf, 0xc1, 0xa1, 0x1a,
	0x8d, 0x12, 0xea, 0x55, 0x41, 0xeb, 0x42, 0x61, 0x81, 0x49, 0xd6, 0xe7, 0xa5, 0x96, 0x01, 0xbd,
	0x5d, 0xe8, 0xfa, 0x3a, 0xcc, 0x62, 0x1f, 0xaa, 0xbc, 0x5c, 0xf8, 0x7e, 0x4a, 0x31, 0x5a, 0xcc,
	0x30, 0x8a, 0x34, 0x3f, 0xa3, 0xda, 0x57, 0x95, 0xfe, 0x87, 0x98, 0xad, 0x63, 0x50, 0xfa, 0x44,
	0x98, 0xad, 0x67, 0x99, 0xad, 0x93, 0x2b, 0x30, 0x61, 0x87, 0x7b, 0x18, 0x97, 0x9e, 0x29, 0xcc,
	0x02, 0x30, 0xee, 0x85, 0x7b, 0x94, 0x09, 0x20, 0xf2, 0xb6, 0x06, 0x4b, 0x51, 0x05, 0x26, 0x8b,
	0xc2, 0x6b, 0x3e, 0x0f, 0xae, 0xb9, 0xb6, 0x89, 0xf1, 0xe8, 0x4a, 0x61, 0x16, 0x51, 0xb9, 0x99,
	0x07, 0x4a, 0xd9, 0x62, 0x8a, 0xbe, 0x1d, 0x91, 0xc9, 0x1b, 0x50, 0x4b, 0x8f, 0xf7, 0xb8, 0x63,
	0xd8, 0xe1, 0x3e, 0x86, 0xa6, 0xad, 0xc2, 0x22, 0x2c, 0x8f, 0x8a, 0x80, 0x90, 0x94, 0x91, 0x14,
	0xf5, 0x15, 0x45, 0x14, 0x71, 0x31, 0x3d, 0xb6, 0xe3, 0x3a, 0xfd, 0x40, 0x16, 0xd8, 0x63, 0xc4,
	0xc5, 0x11, 0x40, 0xca, 0xe6, 0x53, 0xb4, 0x96, 0x20, 0x89, 0xbc, 0x25, 0xba, 0x0a, 0xd8, 0x31,
	0xba, 0xa1, 0xeb, 0x63, 0x9d, 0xfd, 0x7c, 0x61, 0xae, 0x4b, 0xd9, 0x8b, 0x05, 0x85, 0x46, 0xd9,
	0x2c, 0x12, 0x2e, 0xc9, 0x36, 0x79, 0x0e, 0xa0, 0xab, 0xc7, 0x41, 0xf7, 0x94, 0x0c, 0xba, 0x8f,
	0x1c, 0x0c, 0x1a, 0xe5, 0xcd, 0x24, 0xea, 0x46, 0x95, 0xac, 0x9e, 0xc4, 0xdd, 0x72, 0x57, 0x75,
	0x9b, 0xf4, 0x57, 0x25, 0x38, 0xcd, 0x14, 0x64, 0xab, 0xbf, 0xdf, 0x31, 0xba, 0xd7, 0xe3, 0xa2,
	0x6c, 0x9c, 0x78, 0x9e, 0xd2, 0x03, 0xde, 0xe5, 0x95, 0xc6, 0xcb, 0xdf, 0xb2, 0x68, 0x89, 0x1e,
	0xf0, 0x92, 0xce, 0x81, 0xb9, 0x8e, 0x12, 0x3f, 0xe2, 0x37, 0x31, 0x1e, 0xbf, 0x2c, 0x1a, 0x65,
	0xb3, 0x48, 0x50, 0xfc, 0xe8, 0x3f, 0x27, 0x61, 0x76, 0xa3, 0x2f, 0xef, 0x56, 0xf0, 0xf0, 0x5b,
	0x8d, 0xdf, 0x26, 0x94, 0xaa, 0x16, 0x0e, 0x7d, 0x92, 0xf8, 0x06, 0xd4, 0x0d, 0x35, 0x55, 0x37,
	0xfb, 0xbe, 0x72, 0xa8, 0x80, 0x77, 0x5d, 0xc7, 0x0c, 0x30, 0x19, 0x3f, 0x77, 0x7b, 0xd0, 0x68,
	0xe0, 0xdc, 0x43, 0x46, 0x52, 0xf6, 0x00, 0x76, 0x5d, 0xc0, 0x9e, 0xab, 0xaa, 0x43, 0x9c, 0x1e,
	0x9d, 0xfe, 0xce, 0x0e, 0xf7, 0x51, 0x05, 0xc7, 0x3e, 0x3d, 0x14, 0x0a, 0x65, 0x08, 0x27, 0x8e,
	0xd0, 0x6e, 0x3f, 0xf0, 0xf0, 0xb4, 0x3b, 0xf6, 0x11, 0x2a, 0x30, 0x28, 0x93, 0x50, 0x02, 0x32,
	0x08, 0xb9, 0x87, 0xe7, 0xdc, 0xb3, 0x85, 0x8d, 0x55, 0x8d, 0x02, 0x2c, 0x17, 0x90, 0xe2, 0x0f,
	0xb9, 0x02, 0x35, 0xcf, 0xb7, 0xba, 0x5c, 0xdf, 0xe9, 0x3b, 0x78, 0x2d, 0xb6, 0xef, 0x71, 0xac,
	0x1a, 0x57, 0x92, 0x58, 0x92, 0x33, 0x88, 0xb2, 0x05, 0x49, 0xbd, 0x84, 0x44, 0x79, 0x0d, 0xd0,
	0x84, 0xb2, 0xd9, 0x0f, 0xbb, 0xd7, 0x84, 0x65, 0xa7, 0x87, 0x0b, 0xf0, 0xa8, 0x87, 0xb2, 0x69,
	0xf9, 0xb3, 0x6d, 0x8a, 0x33, 0xb6, 0x63, 0x99, 0xa3, 0x96, 0x55, 0x2f, 0x56, 0xa9, 0x33, 0x36,
	0x6f, 0x14, 0x65, 0xa4, 0x63, 0x99, 0x43, 0x16, 0xa5, 0x1f, 0x69, 0x70, 0xba, 0x85, 0x75, 0x52,
	0x94, 0x5a, 0x87, 0xbe, 0xd1, 0xbd, 0xce, 0x7d, 0xf2, 0x74, 0xee, 0xdb, 0xc9, 0xe9, 0xbb, 0x7a,
	0x35, 0x11, 0x45, 0x4f, 0xb4, 0xaf, 0x54, 0x89, 0x88, 0xe8, 0xe8, 0x39, 0xc7, 0x3e, 0x2a, 0x72,
	0x41, 0x29, 0xab, 0x21, 0x5d, 0x96, 0xa7, 0x11, 0xf5, 0x8f, 0x1a, 0x2c, 0x8a, 0xca, 0x24, 0x7a,
	0x2c, 0x8a, 0x57, 0xf6, 0x64, 0xce, 0xeb, 0xf0, 0xd2, 0x1d, 0x9f, 0x71, 0xde, 0x84, 0x5a, 0x04,
	0x94, 0xae, 0x6b, 0x4a, 0x9f, 0xc6, 0x55, 0x36, 0x41, 0x4e, 0xa9, 0x5a, 0x86, 0xbe, 0xa7, 0xc1,
	0xac, 0xba, 0x8c, 0x6c, 0x19, 0xb6, 0xe1, 0x74, 0xf9, 0x71, 0x4b, 0xf4, 0x37, 0x61, 0x11, 0x2f,
	0x30, 0x3b, 0x0a, 0x48, 0x64, 0x63, 0xa1, 0x88, 0x10, 0x13, 0xab, 0xd5, 0xf5, 0xc7, 0x72, 0xef,
	0x1a, 0x33, 0x8c, 0xaf, 0x8a, 0xe1, 0xad, 0x73, 0xd9, 0x17, 0x92, 0x3c, 0x48, 0xca, 0x48, 0x6f,
	0x64, 0x22, 0xfd, 0x83, 0x06, 0x64, 0x14, 0x6f, 0x9c, 0x33, 0x61, 0x0f, 0xa6, 0x91, 0xaf, 0x34,
	0xc7, 0x91, 0x0f, 0x3b, 0x1b, 0x28, 0xf6, 0x5c, 0x94, 0x76, 0xcb, 0x79, 0x85, 0xde, 0x72, 0x22,
	0x66, 0xf4, 0x0d, 0x98, 0xba, 0xec, 0x9a, 0x2d, 0xc3, 0x26, 0x01, 0xd4, 0x76, 0xfa, 0x8e, 0xa9,
	0x67, 0xb5, 0x50, 0xd7, 0xa4, 0x4a, 0x1b, 0xb9, 0x2a, 0xbd, 0xd4, 0x77, 0x4c, 0x35, 0xbb, 0x45,
	0x51, 0x26, 0x0c, 0x20, 0x39, 0x48, 0x94, 0x2d, 0xec, 0xa8, 0xf1, 0x89, 0xda, 0xe8, 0x0f, 0x34,
	0x80, 0xe8, 0x84, 0x35, 0x6c, 0xf2, 0x2d, 0x58, 0x94, 0x33, 0xa3, 0x3d, 0x92, 0x15, 0xe2, 0xdc,
	0xa1, 0x42, 0x24, 0x10, 0xc3, 0x36, 0xcd, 0x83, 0xa3, 0x8c, 0xec, 0x64, 0x26, 0x49, 0xe2, 0xf7,
	0x27, 0x00, 0x92, 0x05, 0x8d, 0x63, 0xcb, 0x94, 0x53, 0x97, 0x0a, 0x38, 0x75, 0xe6, 0x99, 0x73,
	0xe2, 0xb3, 0xff, 0x36, 0xc2, 0xe4, 0x9e, 0x2b, 0xef, 0x7c, 0xad, 0x1e, 0x97, 0xe7, 0x58, 0xa1,
	0x6f, 0x23, 0xd2, 0xb3, 0xf1, 0xdb, 0x08, 0x24, 0xc9, 0xf7, 0x95, 0xc7, 0x61, 0x4a, 0xe8, 0x9c,
	0xfb, 0x78, 0x9c, 0xa5, 0x32, 0x00, 0x45, 0xa7, 0x0c, 0x07, 0xd0, 0xbf, 0x94, 0x60, 0x2e, 0x6b,
	0xd4, 0x71, 0x8c, 0x91, 0xd1, 0x6a, 0xe9, 0x1e, 0x6b, 0x75, 0xe2, 0x53, 0xd3, 0xea, 0xe4, 0x9d,
	0xb4, 0xfa, 0xe1, 0x14, 0x9c, 0xda, 0xb0, 0x6d, 0x54, 0xea, 0xd8, 0xf1, 0xea, 0x17, 0x1a, 0xa4,
	0x1e, 0xb7, 0xf5, 0x1d, 0xdf, 0xed, 0xc5, 0xdb, 0x2c, 0x74, 0xe5, 0xd5, 0x1a, 0xf7, 0x03, 0x3c,
	0x5a, 0xbe, 0x5e, 0x38, 0x77, 0x79, 0x7c, 0xf8, 0xf9, 0xfc, 0x30, 0x0e, 0x94, 0x3d, 0x1c, 0xbf,
	0x95, 0x5f, 0xf2, 0xdd, 0x1e, 0xae, 0x6f, 0xdb, 0xdd, 0x52, 0xfd, 0xe4, 0x97, 0x1a, 0x9c, 0x3b,
	0x0c, 0x66, 0xc7, 0xf5, 0x75, 0x4c, 0x14, 0xf1, 0x54, 0x7f, 0xbd, 0xb0, 0xa4, 0x9f, 0x3b, 0x5a,
	0xd2, 0x14, 0x0b, 0xca, 0x56, 0xf2, 0x44, 0xbd, 0xe4, 0xfa, 0x98, 0x2c, 0x93, 0x9f, 0x68, 0xb0,
	0x1c, 0xbb, 0x9c, 0xc2, 0xb1, 0xad, 0x1b, 0x71, 0x81, 0x38, 0x39, 0xde, 0xfd, 0xe3, 0xe1, 0xc8,
	0x22, 0x5f, 0x46, 0x97, 0x15, 0x82, 0x6d, 0x59, 0x37, 0xa2, 0x5a, 0xf1, 0x1d, 0x0d, 0x1e, 0x1c,
	0x9a, 0xe7, 0x73, 0xcf, 0xd8, 0xef, 0x71, 0x27, 0x0c, 0x70, 0x2b, 0xb3, 0xc2, 0x02, 0x9d, 0xcd,
	0x15, 0x28, 0x01, 0x1e, 0x92, 0x87, 0xc5, 0x1d, 0xe4, 0x67, 0x1a, 0x9c, 0x51, 0xf7, 0xa8, 0x29,
	0x85, 0xa7, 0xfc, 0x4d, 0x5d, 0x48, 0x6f, 0x17, 0x96, 0x88, 0xa6, 0xaf, 0x68, 0x73, 0xa1, 0x29,
	0x3b, 0x2d, 0x7b, 0x37, 0x22, 0x13, 0xc6, 0x2e, 0x46, 0x7f, 0xaf, 0x41, 0x3d, 0xf5, 0x7c, 0x72,
	0xd5, 0x72, 0x76, 0x6d, 0xfe, 0x5f, 0xf2, 0x88, 0x72, 0xd7, 0x5f, 0xd9, 0x88, 0xf3, 0xaf, 0x22,
	0xc4, 0x12, 0x03, 0x83, 0xfb, 0x8f, 0xd0, 0x85, 0x1e, 0xa1, 0x0f, 0x79, 0x8d, 0x3b, 0x79, 0xac,
	0xd7, 0xb8, 0xdf, 0x68, 0x30, 0x9f, 0xae, 0x02, 0xc6, 0xbd, 0x6e, 0xb8, 0x0e, 0xb3, 0xea, 0xe9,
	0x29, 0x2a, 0x60, 0x4a, 0xe3, 0x7d, 0xbe, 0x96, 0x01, 0xa3, 0x4c, 0x7e, 0x53, 0x19, 0x57, 0x2c,
	0x7f, 0xd6, 0x60, 0x26, 0x2d, 0xfc, 0x71, 0x1d, 0xe9, 0x87, 0x1a, 0x90, 0x4c, 0x85, 0xa4, 0xec,
	0xa8, 0x12, 0xfc, 0xff, 0xcf, 0xb5, 0xe3, 0xb0, 0xce, 0x5a, 0x4f, 0x89, 0x15, 0x1e, 0x0c, 0x1a,
	0x23, 0xda, 0x4c, 0x0c, 0x32, 0xca, 0x82, 0xb2, 0x79, 0x6f, 0x68, 0x38, 0xfd, 0x9d, 0x06, 0x0b,
	0x23, 0xe8, 0xe3, 0x98, 0xe4, 0x06, 0x9c, 0xea, 0x64, 0x6b, 0x56, 0x34, 0xca, 0x0b, 0x85, 0x8d,
	0xf2, 0x40, 0xf6, 0xa9, 0x30, 0x36, 0x0b, 0x7e, 0x97, 0x15, 0x1b, 0xe6, 0xaf, 0x1a, 0xcc, 0xa6,
	0xd7, 0xd0, 0x3a, 0xae, 0x65, 0x7e, 0x74, 0x94, 0x65, 0x1e, 0xbd, 0x3b, 0xcb, 0x7c, 0x72, 0xa6,
	0x79, 0x67, 0x0e, 0x6a, 0xa9, 0xb7, 0x96, 0x38, 0x7e, 0xdd, 0x7f, 0x6e, 0xb9, 0xff, 0xdc, 0x72,
	0xff, 0xb9, 0xe5, 0xfe, 0x73, 0xcb, 0xfd, 0xe7, 0x96, 0xff, 0x99, 0xe7, 0x96, 0xe1, 0xe4, 0x71,
	0xfe, 0x13, 0x49, 0x1e, 0x17, 0xc6, 0x4f, 0x1e, 0xc9, 0x3d, 0x49, 0x1e, 0x6b, 0xc7, 0x49, 0x1e,
	0x5b, 0x2f, 0xbc, 0x7f, 0xb0, 0xa2, 0x7d, 0x70, 0xb0, 0xa2, 0xfd, 0xfd, 0x60, 0x45, 0x7b, 0xf7,
	0xe3, 0x95, 0x13, 0x1f, 0x7c, 0xbc, 0x72, 0xe2, 0xc3, 0x8f, 0x57, 0x4e, 0x7c, 0xad, 0x99, 0x31,
	0xb7, 0x58, 0xc9, 0x13, 0xee, 0xce, 0x8e, 0xd5, 0xb5, 0x0c, 0x1b, 0xdb, 0x6b, 0xf8, 0x7f, 0x50,
	0xd2, 0xf4, 0x9d, 0x29, 0x79, 0x3b, 0xf1, 0x85, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xb0,
	0x53, 0x62, 0x23, 0x35, 0x00, 0x00,
}

func (m *LendAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalRewards.Size()
		i -= size
		if _, err := m.TotalRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.CPoolName) > 0 {
		i -= len(m.CPoolName)
		copy(dAtA[i:], m.CPoolName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.CPoolName)))
		i--
		dAtA[i] = 0x5a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastInteractionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastInteractionTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLend(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size := m.GlobalIndex.Size()
		i -= size
		if _, err := m.GlobalIndex.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.AppID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AppID))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.AvailableToBorrow.Size()
		i -= size
		if _, err := m.AvailableToBorrow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LendingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LendingTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLend(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLend(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x18
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BorrowAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BorrowAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BorrowAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsLiquidated {
		i--
		if m.IsLiquidated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.CPoolName) > 0 {
		i -= len(m.CPoolName)
		copy(dAtA[i:], m.CPoolName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.CPoolName)))
		i--
		dAtA[i] = 0x72
	}
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastInteractionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastInteractionTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintLend(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x6a
	{
		size := m.ReserveGlobalIndex.Size()
		i -= size
		if _, err := m.ReserveGlobalIndex.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.GlobalIndex.Size()
		i -= size
		if _, err := m.GlobalIndex.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.InterestAccumulated.Size()
		i -= size
		if _, err := m.InterestAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.StableBorrowRate.Size()
		i -= size
		if _, err := m.StableBorrowRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BorrowingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BorrowingTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintLend(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x42
	{
		size, err := m.BridgedAssetAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.AmountOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.PairID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PairID))
		i--
		dAtA[i] = 0x20
	}
	if m.IsStableBorrow {
		i--
		if m.IsStableBorrow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.LendingID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.LendingID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetData) > 0 {
		for iNdEx := len(m.AssetData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CPoolName) > 0 {
		i -= len(m.CPoolName)
		copy(dAtA[i:], m.CPoolName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.CPoolName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserAssetLendBorrowMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAssetLendBorrowMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAssetLendBorrowMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BorrowId) > 0 {
		dAtA10 := make([]byte, len(m.BorrowId)*10)
		var j9 int
		for _, num := range m.BorrowId {
			for num >= 1<<7 {
				dAtA10[j9] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j9++
			}
			dAtA10[j9] = uint8(num)
			j9++
		}
		i -= j9
		copy(dAtA[i:], dAtA10[:j9])
		i = encodeVarintLend(dAtA, i, uint64(j9))
		i--
		dAtA[i] = 0x22
	}
	if m.PoolId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if m.LendId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.LendId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLend(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AssetDataPoolMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetDataPoolMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetDataPoolMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SupplyCap.Size()
		i -= size
		if _, err := m.SupplyCap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AssetTransitType != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetTransitType))
		i--
		dAtA[i] = 0x10
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Extended_Pair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Extended_Pair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Extended_Pair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinUsdValueLeft != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.MinUsdValueLeft))
		i--
		dAtA[i] = 0x30
	}
	if m.AssetOutPoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetOutPoolID))
		i--
		dAtA[i] = 0x28
	}
	if m.IsInterPool {
		i--
		if m.IsInterPool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AssetOut != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetOut))
		i--
		dAtA[i] = 0x18
	}
	if m.AssetIn != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetIn))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetToPairMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetToPairMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetToPairMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PairID) > 0 {
		dAtA12 := make([]byte, len(m.PairID)*10)
		var j11 int
		for _, num := range m.PairID {
			for num >= 1<<7 {
				dAtA12[j11] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j11++
			}
			dAtA12[j11] = uint8(num)
			j11++
		}
		i -= j11
		copy(dAtA[i:], dAtA12[:j11])
		i = encodeVarintLend(dAtA, i, uint64(j11))
		i--
		dAtA[i] = 0x1a
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolAssetLBMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAssetLBMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAssetLBMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.UtilisationRatio.Size()
		i -= size
		if _, err := m.UtilisationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.StableBorrowApr.Size()
		i -= size
		if _, err := m.StableBorrowApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.BorrowApr.Size()
		i -= size
		if _, err := m.BorrowApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.LendApr.Size()
		i -= size
		if _, err := m.LendApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.TotalInterestAccumulated.Size()
		i -= size
		if _, err := m.TotalInterestAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.TotalLend.Size()
		i -= size
		if _, err := m.TotalLend.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.TotalStableBorrowed.Size()
		i -= size
		if _, err := m.TotalStableBorrowed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.TotalBorrowed.Size()
		i -= size
		if _, err := m.TotalBorrowed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.BorrowIds) > 0 {
		dAtA14 := make([]byte, len(m.BorrowIds)*10)
		var j13 int
		for _, num := range m.BorrowIds {
			for num >= 1<<7 {
				dAtA14[j13] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j13++
			}
			dAtA14[j13] = uint8(num)
			j13++
		}
		i -= j13
		copy(dAtA[i:], dAtA14[:j13])
		i = encodeVarintLend(dAtA, i, uint64(j13))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LendIds) > 0 {
		dAtA16 := make([]byte, len(m.LendIds)*10)
		var j15 int
		for _, num := range m.LendIds {
			for num >= 1<<7 {
				dAtA16[j15] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j15++
			}
			dAtA16[j15] = uint8(num)
			j15++
		}
		i -= j15
		copy(dAtA[i:], dAtA16[:j15])
		i = encodeVarintLend(dAtA, i, uint64(j15))
		i--
		dAtA[i] = 0x1a
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetRatesParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetRatesParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetRatesParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CAssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.CAssetID))
		i--
		dAtA[i] = 0x78
	}
	{
		size := m.ReserveFactor.Size()
		i -= size
		if _, err := m.ReserveFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.LiquidationBonus.Size()
		i -= size
		if _, err := m.LiquidationBonus.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.LiquidationThreshold.Size()
		i -= size
		if _, err := m.LiquidationThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.Ltv.Size()
		i -= size
		if _, err := m.Ltv.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.StableSlope2.Size()
		i -= size
		if _, err := m.StableSlope2.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.StableSlope1.Size()
		i -= size
		if _, err := m.StableSlope1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.StableBase.Size()
		i -= size
		if _, err := m.StableBase.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.EnableStableBorrow {
		i--
		if m.EnableStableBorrow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Slope2.Size()
		i -= size
		if _, err := m.Slope2.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Slope1.Size()
		i -= size
		if _, err := m.Slope1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Base.Size()
		i -= size
		if _, err := m.Base.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UOptimal.Size()
		i -= size
		if _, err := m.UOptimal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReserveBuybackAssetData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReserveBuybackAssetData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReserveBuybackAssetData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BuybackAmount.Size()
		i -= size
		if _, err := m.BuybackAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReserveAmount.Size()
		i -= size
		if _, err := m.ReserveAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BidDurationSeconds != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.BidDurationSeconds))
		i--
		dAtA[i] = 0x40
	}
	if m.DutchId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.DutchId))
		i--
		dAtA[i] = 0x38
	}
	if m.PriceFunctionType != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PriceFunctionType))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Step.Size()
		i -= size
		if _, err := m.Step.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Cusp.Size()
		i -= size
		if _, err := m.Cusp.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Buffer.Size()
		i -= size
		if _, err := m.Buffer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AuctionDurationSeconds != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AuctionDurationSeconds))
		i--
		dAtA[i] = 0x10
	}
	if m.AppId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BorrowInterestTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BorrowInterestTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BorrowInterestTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ReservePoolInterest.Size()
		i -= size
		if _, err := m.ReservePoolInterest.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.BorrowingId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.BorrowingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LendRewardsTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendRewardsTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendRewardsTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RewardsAccumulated.Size()
		i -= size
		if _, err := m.RewardsAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LendingId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.LendingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ModuleBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleBalanceStats) > 0 {
		for iNdEx := len(m.ModuleBalanceStats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModuleBalanceStats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ModuleBalanceStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleBalanceStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleBalanceStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ModBal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModBal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModBal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FundModuleBalance) > 0 {
		for iNdEx := len(m.FundModuleBalance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FundModuleBalance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReserveBal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReserveBal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReserveBal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FundReserveBalance) > 0 {
		for iNdEx := len(m.FundReserveBalance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FundReserveBalance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FundModBal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundModBal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundModBal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintLend(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0x2a
	}
	n18, err18 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DepositTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositTime):])
	if err18 != nil {
		return 0, err18
	}
	i -= n18
	i = encodeVarintLend(dAtA, i, uint64(n18))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x10
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FundReserveBal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundReserveBal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundReserveBal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintLend(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0x22
	}
	n20, err20 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DepositTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositTime):])
	if err20 != nil {
		return 0, err20
	}
	i -= n20
	i = encodeVarintLend(dAtA, i, uint64(n20))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AllReserveStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllReserveStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllReserveStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalAmountOutToLenders.Size()
		i -= size
		if _, err := m.TotalAmountOutToLenders.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.AmountInFromRepayments.Size()
		i -= size
		if _, err := m.AmountInFromRepayments.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.AmountInFromLiqPenalty.Size()
		i -= size
		if _, err := m.AmountInFromLiqPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.AmountOutFromReserveForAuction.Size()
		i -= size
		if _, err := m.AmountOutFromReserveForAuction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AmountOutFromReserveToLenders.Size()
		i -= size
		if _, err := m.AmountOutFromReserveToLenders.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetToPairSingleMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetToPairSingleMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetToPairSingleMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PairID))
		i--
		dAtA[i] = 0x18
	}
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolPairs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolPairs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolPairs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinUsdValueLeft != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.MinUsdValueLeft))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AssetData) > 0 {
		for iNdEx := len(m.AssetData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CPoolName) > 0 {
		i -= len(m.CPoolName)
		copy(dAtA[i:], m.CPoolName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.CPoolName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolInterestData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInterestData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInterestData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LendInterest.Size()
		i -= size
		if _, err := m.LendInterest.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolInterest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInterest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInterest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolInterestData) > 0 {
		for iNdEx := len(m.PoolInterestData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolInterestData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolInterestDataB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInterestDataB) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInterestDataB) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BorrowInterest.Size()
		i -= size
		if _, err := m.BorrowInterest.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolInterestB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInterestB) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInterestB) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolInterestData) > 0 {
		for iNdEx := len(m.PoolInterestData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolInterestData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PoolID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetRatesPoolPairs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetRatesPoolPairs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetRatesPoolPairs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinUsdValueLeft != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.MinUsdValueLeft))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if len(m.AssetData) > 0 {
		for iNdEx := len(m.AssetData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x92
		}
	}
	if len(m.CPoolName) > 0 {
		i -= len(m.CPoolName)
		copy(dAtA[i:], m.CPoolName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.CPoolName)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.CAssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.CAssetID))
		i--
		dAtA[i] = 0x78
	}
	{
		size := m.ReserveFactor.Size()
		i -= size
		if _, err := m.ReserveFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.LiquidationBonus.Size()
		i -= size
		if _, err := m.LiquidationBonus.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.LiquidationThreshold.Size()
		i -= size
		if _, err := m.LiquidationThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.Ltv.Size()
		i -= size
		if _, err := m.Ltv.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.StableSlope2.Size()
		i -= size
		if _, err := m.StableSlope2.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.StableSlope1.Size()
		i -= size
		if _, err := m.StableSlope1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.StableBase.Size()
		i -= size
		if _, err := m.StableBase.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.EnableStableBorrow {
		i--
		if m.EnableStableBorrow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Slope2.Size()
		i -= size
		if _, err := m.Slope2.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Slope1.Size()
		i -= size
		if _, err := m.Slope1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Base.Size()
		i -= size
		if _, err := m.Base.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UOptimal.Size()
		i -= size
		if _, err := m.UOptimal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLend(dAtA []byte, offset int, v uint64) int {
	offset -= sovLend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LendAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLend(uint64(m.ID))
	}
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LendingTime)
	n += 1 + l + sovLend(uint64(l))
	l = m.AvailableToBorrow.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.AppID != 0 {
		n += 1 + sovLend(uint64(m.AppID))
	}
	l = m.GlobalIndex.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastInteractionTime)
	n += 1 + l + sovLend(uint64(l))
	l = len(m.CPoolName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	l = m.TotalRewards.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *BorrowAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLend(uint64(m.ID))
	}
	if m.LendingID != 0 {
		n += 1 + sovLend(uint64(m.LendingID))
	}
	if m.IsStableBorrow {
		n += 2
	}
	if m.PairID != 0 {
		n += 1 + sovLend(uint64(m.PairID))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.BridgedAssetAmount.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BorrowingTime)
	n += 1 + l + sovLend(uint64(l))
	l = m.StableBorrowRate.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.InterestAccumulated.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.GlobalIndex.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.ReserveGlobalIndex.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastInteractionTime)
	n += 1 + l + sovLend(uint64(l))
	l = len(m.CPoolName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	if m.IsLiquidated {
		n += 2
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	l = len(m.CPoolName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	if len(m.AssetData) > 0 {
		for _, e := range m.AssetData {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *UserAssetLendBorrowMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	if m.LendId != 0 {
		n += 1 + sovLend(uint64(m.LendId))
	}
	if m.PoolId != 0 {
		n += 1 + sovLend(uint64(m.PoolId))
	}
	if len(m.BorrowId) > 0 {
		l = 0
		for _, e := range m.BorrowId {
			l += sovLend(uint64(e))
		}
		n += 1 + sovLend(uint64(l)) + l
	}
	return n
}

func (m *AssetDataPoolMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if m.AssetTransitType != 0 {
		n += 1 + sovLend(uint64(m.AssetTransitType))
	}
	l = m.SupplyCap.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *Extended_Pair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLend(uint64(m.Id))
	}
	if m.AssetIn != 0 {
		n += 1 + sovLend(uint64(m.AssetIn))
	}
	if m.AssetOut != 0 {
		n += 1 + sovLend(uint64(m.AssetOut))
	}
	if m.IsInterPool {
		n += 2
	}
	if m.AssetOutPoolID != 0 {
		n += 1 + sovLend(uint64(m.AssetOutPoolID))
	}
	if m.MinUsdValueLeft != 0 {
		n += 1 + sovLend(uint64(m.MinUsdValueLeft))
	}
	return n
}

func (m *AssetToPairMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if len(m.PairID) > 0 {
		l = 0
		for _, e := range m.PairID {
			l += sovLend(uint64(e))
		}
		n += 1 + sovLend(uint64(l)) + l
	}
	return n
}

func (m *PoolAssetLBMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if len(m.LendIds) > 0 {
		l = 0
		for _, e := range m.LendIds {
			l += sovLend(uint64(e))
		}
		n += 1 + sovLend(uint64(l)) + l
	}
	if len(m.BorrowIds) > 0 {
		l = 0
		for _, e := range m.BorrowIds {
			l += sovLend(uint64(e))
		}
		n += 1 + sovLend(uint64(l)) + l
	}
	l = m.TotalBorrowed.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.TotalStableBorrowed.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.TotalLend.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.TotalInterestAccumulated.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LendApr.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.BorrowApr.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.StableBorrowApr.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.UtilisationRatio.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *AssetRatesParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.UOptimal.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Base.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Slope1.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Slope2.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.EnableStableBorrow {
		n += 2
	}
	l = m.StableBase.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.StableSlope1.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.StableSlope2.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Ltv.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationThreshold.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationBonus.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.ReserveFactor.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.CAssetID != 0 {
		n += 1 + sovLend(uint64(m.CAssetID))
	}
	return n
}

func (m *ReserveBuybackAssetData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.ReserveAmount.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.BuybackAmount.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *AuctionParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovLend(uint64(m.AppId))
	}
	if m.AuctionDurationSeconds != 0 {
		n += 1 + sovLend(uint64(m.AuctionDurationSeconds))
	}
	l = m.Buffer.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Cusp.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Step.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.PriceFunctionType != 0 {
		n += 1 + sovLend(uint64(m.PriceFunctionType))
	}
	if m.DutchId != 0 {
		n += 1 + sovLend(uint64(m.DutchId))
	}
	if m.BidDurationSeconds != 0 {
		n += 1 + sovLend(uint64(m.BidDurationSeconds))
	}
	return n
}

func (m *BorrowInterestTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BorrowingId != 0 {
		n += 1 + sovLend(uint64(m.BorrowingId))
	}
	l = m.ReservePoolInterest.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *LendRewardsTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LendingId != 0 {
		n += 1 + sovLend(uint64(m.LendingId))
	}
	l = m.RewardsAccumulated.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *ModuleBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if len(m.ModuleBalanceStats) > 0 {
		for _, e := range m.ModuleBalanceStats {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *ModuleBalanceStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.Balance.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *ModBal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FundModuleBalance) > 0 {
		for _, e := range m.FundModuleBalance {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *ReserveBal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FundReserveBalance) > 0 {
		for _, e := range m.FundReserveBalance {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *FundModBal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositTime)
	n += 1 + l + sovLend(uint64(l))
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	return n
}

func (m *FundReserveBal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositTime)
	n += 1 + l + sovLend(uint64(l))
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	return n
}

func (m *AllReserveStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.AmountOutFromReserveToLenders.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.AmountOutFromReserveForAuction.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.AmountInFromLiqPenalty.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.AmountInFromRepayments.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.TotalAmountOutToLenders.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *AssetToPairSingleMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	if m.PairID != 0 {
		n += 1 + sovLend(uint64(m.PairID))
	}
	return n
}

func (m *PoolPairs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	l = len(m.CPoolName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	if len(m.AssetData) > 0 {
		for _, e := range m.AssetData {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	if m.MinUsdValueLeft != 0 {
		n += 1 + sovLend(uint64(m.MinUsdValueLeft))
	}
	return n
}

func (m *PoolInterestData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.LendInterest.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *PoolInterest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if len(m.PoolInterestData) > 0 {
		for _, e := range m.PoolInterestData {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *PoolInterestDataB) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.BorrowInterest.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *PoolInterestB) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLend(uint64(m.PoolID))
	}
	if len(m.PoolInterestData) > 0 {
		for _, e := range m.PoolInterestData {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *AssetRatesPoolPairs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovLend(uint64(m.AssetID))
	}
	l = m.UOptimal.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Base.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Slope1.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Slope2.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.EnableStableBorrow {
		n += 2
	}
	l = m.StableBase.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.StableSlope1.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.StableSlope2.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.Ltv.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationThreshold.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.LiquidationBonus.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.ReserveFactor.Size()
	n += 1 + l + sovLend(uint64(l))
	if m.CAssetID != 0 {
		n += 1 + sovLend(uint64(m.CAssetID))
	}
	l = len(m.ModuleName)
	if l > 0 {
		n += 2 + l + sovLend(uint64(l))
	}
	l = len(m.CPoolName)
	if l > 0 {
		n += 2 + l + sovLend(uint64(l))
	}
	if len(m.AssetData) > 0 {
		for _, e := range m.AssetData {
			l = e.Size()
			n += 2 + l + sovLend(uint64(l))
		}
	}
	if m.MinUsdValueLeft != 0 {
		n += 2 + sovLend(uint64(m.MinUsdValueLeft))
	}
	return n
}

func sovLend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLend(x uint64) (n int) {
	return sovLend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LendAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LendAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LendAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LendingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableToBorrow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AvailableToBorrow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			m.AppID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GlobalIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastInteractionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastInteractionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BorrowAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BorrowAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BorrowAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendingID", wireType)
			}
			m.LendingID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LendingID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsStableBorrow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsStableBorrow = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			m.PairID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgedAssetAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BridgedAssetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BorrowingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableBorrowRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableBorrowRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GlobalIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveGlobalIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveGlobalIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastInteractionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastInteractionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLiquidated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLiquidated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetData = append(m.AssetData, &AssetDataPoolMapping{})
			if err := m.AssetData[len(m.AssetData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserAssetLendBorrowMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserAssetLendBorrowMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAssetLendBorrowMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendId", wireType)
			}
			m.LendId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LendId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BorrowId = append(m.BorrowId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLend
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLend
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BorrowId) == 0 {
					m.BorrowId = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLend
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BorrowId = append(m.BorrowId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowId", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetDataPoolMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetDataPoolMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetDataPoolMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetTransitType", wireType)
			}
			m.AssetTransitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetTransitType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplyCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SupplyCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Extended_Pair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Extended_Pair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Extended_Pair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetIn", wireType)
			}
			m.AssetIn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetIn |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOut", wireType)
			}
			m.AssetOut = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOut |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInterPool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsInterPool = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutPoolID", wireType)
			}
			m.AssetOutPoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOutPoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUsdValueLeft", wireType)
			}
			m.MinUsdValueLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUsdValueLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetToPairMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetToPairMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetToPairMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PairID = append(m.PairID, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLend
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLend
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PairID) == 0 {
					m.PairID = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLend
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PairID = append(m.PairID, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolAssetLBMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolAssetLBMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAssetLBMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.LendIds = append(m.LendIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLend
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLend
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.LendIds) == 0 {
					m.LendIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLend
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.LendIds = append(m.LendIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field LendIds", wireType)
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BorrowIds = append(m.BorrowIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLend
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLend
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLend
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BorrowIds) == 0 {
					m.BorrowIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLend
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BorrowIds = append(m.BorrowIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowIds", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBorrowed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBorrowed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStableBorrowed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalStableBorrowed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalLend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalLend.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalInterestAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalInterestAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LendApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableBorrowApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableBorrowApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UtilisationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UtilisationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetRatesParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetRatesParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetRatesParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UOptimal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UOptimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Slope1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Slope2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableStableBorrow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableStableBorrow = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableBase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableSlope1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableSlope1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableSlope2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableSlope2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ltv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ltv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationBonus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationBonus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CAssetID", wireType)
			}
			m.CAssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CAssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReserveBuybackAssetData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReserveBuybackAssetData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReserveBuybackAssetData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuybackAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BuybackAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuctionParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuctionParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionDurationSeconds", wireType)
			}
			m.AuctionDurationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionDurationSeconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Buffer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cusp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cusp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Step.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFunctionType", wireType)
			}
			m.PriceFunctionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceFunctionType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutchId", wireType)
			}
			m.DutchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DutchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidDurationSeconds", wireType)
			}
			m.BidDurationSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BidDurationSeconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BorrowInterestTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Borrow_interest_tracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Borrow_interest_tracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowingId", wireType)
			}
			m.BorrowingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BorrowingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservePoolInterest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservePoolInterest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LendRewardsTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lend_rewards_tracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lend_rewards_tracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendingId", wireType)
			}
			m.LendingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LendingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardsAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ModuleBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModuleBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleBalanceStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleBalanceStats = append(m.ModuleBalanceStats, ModuleBalanceStats{})
			if err := m.ModuleBalanceStats[len(m.ModuleBalanceStats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ModuleBalanceStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModuleBalanceStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleBalanceStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ModBal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModBal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModBal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundModuleBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FundModuleBalance = append(m.FundModuleBalance, FundModBal{})
			if err := m.FundModuleBalance[len(m.FundModuleBalance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReserveBal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReserveBal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReserveBal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundReserveBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FundReserveBalance = append(m.FundReserveBalance, FundReserveBal{})
			if err := m.FundReserveBalance[len(m.FundReserveBalance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FundModBal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundModBal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundModBal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DepositTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FundReserveBal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundReserveBal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundReserveBal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DepositTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllReserveStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllReserveStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllReserveStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOutFromReserveToLenders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOutFromReserveToLenders.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOutFromReserveForAuction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOutFromReserveForAuction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInFromLiqPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountInFromLiqPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInFromRepayments", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountInFromRepayments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmountOutToLenders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAmountOutToLenders.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetToPairSingleMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetToPairSingleMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetToPairSingleMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			m.PairID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolPairs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolPairs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolPairs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetData = append(m.AssetData, &AssetDataPoolMapping{})
			if err := m.AssetData[len(m.AssetData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUsdValueLeft", wireType)
			}
			m.MinUsdValueLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUsdValueLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolInterestData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolInterestData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInterestData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendInterest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LendInterest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolInterest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolInterest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInterest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolInterestData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolInterestData = append(m.PoolInterestData, PoolInterestData{})
			if err := m.PoolInterestData[len(m.PoolInterestData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolInterestDataB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolInterestDataB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInterestDataB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowInterest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowInterest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolInterestB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolInterestB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInterestB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolInterestData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolInterestData = append(m.PoolInterestData, PoolInterestDataB{})
			if err := m.PoolInterestData[len(m.PoolInterestData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetRatesPoolPairs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetRatesPoolPairs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetRatesPoolPairs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UOptimal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UOptimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Slope1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Slope2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableStableBorrow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableStableBorrow = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableBase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableSlope1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableSlope1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StableSlope2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StableSlope2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ltv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ltv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationBonus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationBonus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CAssetID", wireType)
			}
			m.CAssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CAssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetData = append(m.AssetData, &AssetDataPoolMapping{})
			if err := m.AssetData[len(m.AssetData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUsdValueLeft", wireType)
			}
			m.MinUsdValueLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUsdValueLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLend
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLend = fmt.Errorf("proto: unexpected end of group")
)
