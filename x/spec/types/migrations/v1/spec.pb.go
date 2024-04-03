// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/spec/spec.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Spec_ProvidersTypes int32

const (
	Spec_dynamic Spec_ProvidersTypes = 0
	Spec_static  Spec_ProvidersTypes = 1
)

var Spec_ProvidersTypes_name = map[int32]string{
	0: "dynamic",
	1: "static",
}

var Spec_ProvidersTypes_value = map[string]int32{
	"dynamic": 0,
	"static":  1,
}

func (x Spec_ProvidersTypes) String() string {
	return proto.EnumName(Spec_ProvidersTypes_name, int32(x))
}

func (Spec_ProvidersTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_789140b95c48dfce, []int{0, 0}
}

type Spec struct {
	Index                         string                                  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name                          string                                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Enabled                       bool                                    `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ReliabilityThreshold          uint32                                  `protobuf:"varint,5,opt,name=reliability_threshold,json=reliabilityThreshold,proto3" json:"reliability_threshold,omitempty"`
	DataReliabilityEnabled        bool                                    `protobuf:"varint,6,opt,name=data_reliability_enabled,json=dataReliabilityEnabled,proto3" json:"data_reliability_enabled,omitempty"`
	BlockDistanceForFinalizedData uint32                                  `protobuf:"varint,7,opt,name=block_distance_for_finalized_data,json=blockDistanceForFinalizedData,proto3" json:"block_distance_for_finalized_data,omitempty"`
	BlocksInFinalizationProof     uint32                                  `protobuf:"varint,8,opt,name=blocks_in_finalization_proof,json=blocksInFinalizationProof,proto3" json:"blocks_in_finalization_proof,omitempty"`
	AverageBlockTime              int64                                   `protobuf:"varint,9,opt,name=average_block_time,json=averageBlockTime,proto3" json:"average_block_time,omitempty"`
	AllowedBlockLagForQosSync     int64                                   `protobuf:"varint,10,opt,name=allowed_block_lag_for_qos_sync,json=allowedBlockLagForQosSync,proto3" json:"allowed_block_lag_for_qos_sync,omitempty"`
	BlockLastUpdated              uint64                                  `protobuf:"varint,11,opt,name=block_last_updated,json=blockLastUpdated,proto3" json:"block_last_updated,omitempty"`
	MinStakeProvider              types.Coin                              `protobuf:"bytes,12,opt,name=min_stake_provider,json=minStakeProvider,proto3" json:"min_stake_provider"`
	ProvidersTypes                Spec_ProvidersTypes                     `protobuf:"varint,14,opt,name=providers_types,json=providersTypes,proto3,enum=lavanet.lava.spec.Spec_ProvidersTypes" json:"providers_types,omitempty"`
	Imports                       []string                                `protobuf:"bytes,15,rep,name=imports,proto3" json:"imports,omitempty"`
	ApiCollections                []*ApiCollection                        `protobuf:"bytes,16,rep,name=api_collections,json=apiCollections,proto3" json:"api_collections,omitempty"`
	Contributor                   []string                                `protobuf:"bytes,17,rep,name=contributor,proto3" json:"contributor,omitempty"`
	ContributorPercentage         *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,18,opt,name=contributor_percentage,json=contributorPercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"contributor_percentage,omitempty"`
	Shares                        uint64                                  `protobuf:"varint,19,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_789140b95c48dfce, []int{0}
}
func (m *Spec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return m.Size()
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Spec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Spec) GetReliabilityThreshold() uint32 {
	if m != nil {
		return m.ReliabilityThreshold
	}
	return 0
}

func (m *Spec) GetDataReliabilityEnabled() bool {
	if m != nil {
		return m.DataReliabilityEnabled
	}
	return false
}

func (m *Spec) GetBlockDistanceForFinalizedData() uint32 {
	if m != nil {
		return m.BlockDistanceForFinalizedData
	}
	return 0
}

func (m *Spec) GetBlocksInFinalizationProof() uint32 {
	if m != nil {
		return m.BlocksInFinalizationProof
	}
	return 0
}

func (m *Spec) GetAverageBlockTime() int64 {
	if m != nil {
		return m.AverageBlockTime
	}
	return 0
}

func (m *Spec) GetAllowedBlockLagForQosSync() int64 {
	if m != nil {
		return m.AllowedBlockLagForQosSync
	}
	return 0
}

func (m *Spec) GetBlockLastUpdated() uint64 {
	if m != nil {
		return m.BlockLastUpdated
	}
	return 0
}

func (m *Spec) GetMinStakeProvider() types.Coin {
	if m != nil {
		return m.MinStakeProvider
	}
	return types.Coin{}
}

func (m *Spec) GetProvidersTypes() Spec_ProvidersTypes {
	if m != nil {
		return m.ProvidersTypes
	}
	return Spec_dynamic
}

func (m *Spec) GetImports() []string {
	if m != nil {
		return m.Imports
	}
	return nil
}

func (m *Spec) GetApiCollections() []*ApiCollection {
	if m != nil {
		return m.ApiCollections
	}
	return nil
}

func (m *Spec) GetContributor() []string {
	if m != nil {
		return m.Contributor
	}
	return nil
}

func (m *Spec) GetShares() uint64 {
	if m != nil {
		return m.Shares
	}
	return 0
}

func init() {
	proto.RegisterEnum("lavanet.lava.spec.Spec_ProvidersTypes", Spec_ProvidersTypes_name, Spec_ProvidersTypes_value)
	proto.RegisterType((*Spec)(nil), "lavanet.lava.spec.Spec")
}

func init() { proto.RegisterFile("lavanet/lava/spec/spec.proto", fileDescriptor_789140b95c48dfce) }

var fileDescriptor_789140b95c48dfce = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcf, 0x4e, 0xfb, 0x36,
	0x1c, 0x6f, 0xd6, 0xfc, 0xda, 0xe2, 0xee, 0x57, 0x82, 0x07, 0xc8, 0x20, 0x16, 0x32, 0x34, 0xa1,
	0x6c, 0xda, 0x12, 0x01, 0x97, 0xdd, 0x26, 0x0a, 0xab, 0x06, 0xda, 0x34, 0x16, 0xd8, 0x65, 0x97,
	0xc8, 0x71, 0x4c, 0x6b, 0x91, 0xd8, 0x59, 0x6c, 0x3a, 0xba, 0xa7, 0xd8, 0x63, 0xec, 0x21, 0xf6,
	0x00, 0x1c, 0x39, 0x4e, 0x3b, 0xa0, 0xa9, 0xbc, 0xc8, 0x64, 0x27, 0x61, 0xa9, 0xd8, 0xa5, 0xf6,
	0xd7, 0x9f, 0x3f, 0xdf, 0x6f, 0xdd, 0x4f, 0x0d, 0xf6, 0x32, 0x3c, 0xc7, 0x9c, 0xaa, 0x50, 0xaf,
	0xa1, 0x2c, 0x28, 0x31, 0x1f, 0x41, 0x51, 0x0a, 0x25, 0xe0, 0x46, 0x8d, 0x06, 0x7a, 0x0d, 0x34,
	0xb0, 0xbb, 0x39, 0x15, 0x53, 0x61, 0xd0, 0x50, 0xef, 0x2a, 0xe2, 0xee, 0xe1, 0x5b, 0x1b, 0x5c,
	0xb0, 0x98, 0x88, 0x2c, 0xa3, 0x44, 0x31, 0xc1, 0x6b, 0x9e, 0x4b, 0x84, 0xcc, 0x85, 0x0c, 0x13,
	0x2c, 0x69, 0x38, 0x3f, 0x4a, 0xa8, 0xc2, 0x47, 0x21, 0x11, 0xac, 0xc6, 0x0f, 0xfe, 0xec, 0x03,
	0xfb, 0xba, 0xa0, 0x04, 0x6e, 0x82, 0x77, 0x8c, 0xa7, 0xf4, 0x01, 0x59, 0x9e, 0xe5, 0xaf, 0x45,
	0x55, 0x01, 0x21, 0xb0, 0x39, 0xce, 0x29, 0xfa, 0xc0, 0x1c, 0x9a, 0x3d, 0x44, 0xa0, 0x4f, 0x39,
	0x4e, 0x32, 0x9a, 0x22, 0xdb, 0xb3, 0xfc, 0x41, 0xd4, 0x94, 0xf0, 0x04, 0x6c, 0x95, 0x34, 0x63,
	0x38, 0x61, 0x19, 0x53, 0x8b, 0x58, 0xcd, 0x4a, 0x2a, 0x67, 0x22, 0x4b, 0xd1, 0x3b, 0xcf, 0xf2,
	0xdf, 0x47, 0x9b, 0x2d, 0xf0, 0xa6, 0xc1, 0xe0, 0x57, 0x00, 0xa5, 0x58, 0xe1, 0xb8, 0xad, 0x6c,
	0xfc, 0x7b, 0xc6, 0x7f, 0x5b, 0xe3, 0xd1, 0x7f, 0xf0, 0x37, 0x75, 0xbb, 0x6f, 0xc1, 0x27, 0x49,
	0x26, 0xc8, 0x5d, 0x9c, 0x32, 0xa9, 0x30, 0x27, 0x34, 0xbe, 0x15, 0x65, 0x7c, 0xcb, 0x38, 0xce,
	0xd8, 0x6f, 0x34, 0x8d, 0xb5, 0x0c, 0xf5, 0x4d, 0xeb, 0x8f, 0x0d, 0xf1, 0xbc, 0xe6, 0x4d, 0x44,
	0x39, 0x69, 0x58, 0xe7, 0x58, 0x61, 0xf8, 0x35, 0xd8, 0x33, 0x04, 0x19, 0x33, 0xde, 0x18, 0x60,
	0x7d, 0x8b, 0x71, 0x51, 0x0a, 0x71, 0x8b, 0x06, 0xc6, 0x64, 0xa7, 0xe2, 0x5c, 0xf0, 0x49, 0x8b,
	0x71, 0xa5, 0x09, 0xf0, 0x0b, 0x00, 0xf1, 0x9c, 0x96, 0x78, 0x4a, 0xe3, 0x6a, 0x24, 0xc5, 0x72,
	0x8a, 0xd6, 0x3c, 0xcb, 0xef, 0x46, 0x4e, 0x8d, 0x8c, 0x35, 0x70, 0xc3, 0x72, 0x0a, 0x4f, 0x81,
	0x8b, 0xb3, 0x4c, 0xfc, 0x4a, 0xd3, 0x9a, 0x9d, 0xe1, 0xa9, 0x99, 0xfd, 0x17, 0x21, 0x63, 0xb9,
	0xe0, 0x04, 0x01, 0xa3, 0xdc, 0xa9, 0x59, 0x46, 0xf9, 0x1d, 0x9e, 0x4e, 0x44, 0xf9, 0xa3, 0x90,
	0xd7, 0x0b, 0x4e, 0x74, 0xc3, 0x46, 0x2a, 0x55, 0x7c, 0x5f, 0xa4, 0x58, 0xd1, 0x14, 0x0d, 0x3d,
	0xcb, 0xb7, 0x23, 0x27, 0xa9, 0xf8, 0x52, 0xfd, 0x54, 0x9d, 0xc3, 0xef, 0x01, 0xcc, 0x19, 0x8f,
	0xa5, 0xc2, 0x77, 0x54, 0x7f, 0xa5, 0x39, 0x4b, 0x69, 0x89, 0x3e, 0xf4, 0x2c, 0x7f, 0x78, 0xbc,
	0x13, 0x54, 0x11, 0x09, 0x74, 0x44, 0x82, 0x3a, 0x22, 0xc1, 0x99, 0x60, 0x7c, 0x6c, 0x3f, 0x3e,
	0xef, 0x77, 0x22, 0x27, 0x67, 0xfc, 0x5a, 0x2b, 0xaf, 0x6a, 0x21, 0xfc, 0x01, 0xac, 0x37, 0x26,
	0x32, 0x56, 0x8b, 0x82, 0x4a, 0x34, 0xf2, 0x2c, 0x7f, 0x74, 0x7c, 0x18, 0xbc, 0xc9, 0x6f, 0xa0,
	0xd3, 0x15, 0x34, 0x52, 0x79, 0xa3, 0xd9, 0xd1, 0xa8, 0x58, 0xa9, 0x75, 0xa4, 0x58, 0x5e, 0x88,
	0x52, 0x49, 0xb4, 0xee, 0x75, 0xfd, 0xb5, 0xa8, 0x29, 0xe1, 0x05, 0x58, 0x5f, 0xcd, 0xb5, 0x44,
	0x8e, 0xd7, 0xf5, 0x87, 0xc7, 0xde, 0xff, 0xb4, 0x3a, 0x2d, 0xd8, 0xd9, 0x2b, 0x31, 0x1a, 0xe1,
	0x76, 0x29, 0xa1, 0x07, 0x86, 0x44, 0x70, 0x55, 0xb2, 0xe4, 0x5e, 0x89, 0x12, 0x6d, 0x98, 0x46,
	0xed, 0x23, 0x88, 0xc1, 0x76, 0xab, 0x8c, 0x0b, 0x5a, 0x12, 0xca, 0x15, 0x9e, 0x52, 0x04, 0x75,
	0xfe, 0xc7, 0x9f, 0xff, 0xfd, 0xbc, 0x7f, 0x38, 0x65, 0x6a, 0x76, 0x9f, 0x04, 0x44, 0xe4, 0x61,
	0xfd, 0xdf, 0xaa, 0x96, 0x2f, 0x65, 0x7a, 0x17, 0x9a, 0xcb, 0x08, 0xce, 0x29, 0x89, 0xb6, 0x5a,
	0x4e, 0x57, 0xaf, 0x46, 0x70, 0x1b, 0xf4, 0xe4, 0x0c, 0x97, 0x54, 0xa2, 0x8f, 0xcc, 0x6f, 0x55,
	0x57, 0x07, 0x9f, 0x81, 0xd1, 0xea, 0x1d, 0xc1, 0x21, 0xe8, 0xa7, 0x0b, 0x8e, 0x73, 0x46, 0x9c,
	0x0e, 0x04, 0xa0, 0x27, 0x15, 0x56, 0x8c, 0x38, 0xd6, 0xa5, 0x3d, 0xe8, 0x3a, 0xf6, 0xa5, 0x3d,
	0x78, 0xef, 0x8c, 0xc6, 0xe3, 0x3f, 0x96, 0xae, 0xf5, 0xb8, 0x74, 0xad, 0xa7, 0xa5, 0x6b, 0xfd,
	0xb3, 0x74, 0xad, 0xdf, 0x5f, 0xdc, 0xce, 0xd3, 0x8b, 0xdb, 0xf9, 0xeb, 0xc5, 0xed, 0xfc, 0xfc,
	0x69, 0x6b, 0xd6, 0x95, 0xf7, 0xe2, 0xa1, 0x7a, 0x31, 0xcc, 0xb4, 0x49, 0xcf, 0xbc, 0x04, 0x27,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x67, 0xe7, 0xdf, 0x36, 0x9a, 0x04, 0x00, 0x00,
}

func (this *Spec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Spec)
	if !ok {
		that2, ok := that.(Spec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ReliabilityThreshold != that1.ReliabilityThreshold {
		return false
	}
	if this.DataReliabilityEnabled != that1.DataReliabilityEnabled {
		return false
	}
	if this.BlockDistanceForFinalizedData != that1.BlockDistanceForFinalizedData {
		return false
	}
	if this.BlocksInFinalizationProof != that1.BlocksInFinalizationProof {
		return false
	}
	if this.AverageBlockTime != that1.AverageBlockTime {
		return false
	}
	if this.AllowedBlockLagForQosSync != that1.AllowedBlockLagForQosSync {
		return false
	}
	if this.BlockLastUpdated != that1.BlockLastUpdated {
		return false
	}
	if !this.MinStakeProvider.Equal(&that1.MinStakeProvider) {
		return false
	}
	if this.ProvidersTypes != that1.ProvidersTypes {
		return false
	}
	if len(this.Imports) != len(that1.Imports) {
		return false
	}
	for i := range this.Imports {
		if this.Imports[i] != that1.Imports[i] {
			return false
		}
	}
	if len(this.ApiCollections) != len(that1.ApiCollections) {
		return false
	}
	for i := range this.ApiCollections {
		if !this.ApiCollections[i].Equal(that1.ApiCollections[i]) {
			return false
		}
	}
	if len(this.Contributor) != len(that1.Contributor) {
		return false
	}
	for i := range this.Contributor {
		if this.Contributor[i] != that1.Contributor[i] {
			return false
		}
	}
	if that1.ContributorPercentage == nil {
		if this.ContributorPercentage != nil {
			return false
		}
	} else if !this.ContributorPercentage.Equal(*that1.ContributorPercentage) {
		return false
	}
	if this.Shares != that1.Shares {
		return false
	}
	return true
}
func (m *Spec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Spec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Spec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Shares != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.Shares))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.ContributorPercentage != nil {
		{
			size := m.ContributorPercentage.Size()
			i -= size
			if _, err := m.ContributorPercentage.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSpec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.Contributor) > 0 {
		for iNdEx := len(m.Contributor) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributor[iNdEx])
			copy(dAtA[i:], m.Contributor[iNdEx])
			i = encodeVarintSpec(dAtA, i, uint64(len(m.Contributor[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.ApiCollections) > 0 {
		for iNdEx := len(m.ApiCollections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiCollections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.Imports) > 0 {
		for iNdEx := len(m.Imports) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Imports[iNdEx])
			copy(dAtA[i:], m.Imports[iNdEx])
			i = encodeVarintSpec(dAtA, i, uint64(len(m.Imports[iNdEx])))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.ProvidersTypes != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.ProvidersTypes))
		i--
		dAtA[i] = 0x70
	}
	{
		size, err := m.MinStakeProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpec(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.BlockLastUpdated != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlockLastUpdated))
		i--
		dAtA[i] = 0x58
	}
	if m.AllowedBlockLagForQosSync != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.AllowedBlockLagForQosSync))
		i--
		dAtA[i] = 0x50
	}
	if m.AverageBlockTime != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.AverageBlockTime))
		i--
		dAtA[i] = 0x48
	}
	if m.BlocksInFinalizationProof != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlocksInFinalizationProof))
		i--
		dAtA[i] = 0x40
	}
	if m.BlockDistanceForFinalizedData != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.BlockDistanceForFinalizedData))
		i--
		dAtA[i] = 0x38
	}
	if m.DataReliabilityEnabled {
		i--
		if m.DataReliabilityEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ReliabilityThreshold != 0 {
		i = encodeVarintSpec(dAtA, i, uint64(m.ReliabilityThreshold))
		i--
		dAtA[i] = 0x28
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpec(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Spec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.ReliabilityThreshold != 0 {
		n += 1 + sovSpec(uint64(m.ReliabilityThreshold))
	}
	if m.DataReliabilityEnabled {
		n += 2
	}
	if m.BlockDistanceForFinalizedData != 0 {
		n += 1 + sovSpec(uint64(m.BlockDistanceForFinalizedData))
	}
	if m.BlocksInFinalizationProof != 0 {
		n += 1 + sovSpec(uint64(m.BlocksInFinalizationProof))
	}
	if m.AverageBlockTime != 0 {
		n += 1 + sovSpec(uint64(m.AverageBlockTime))
	}
	if m.AllowedBlockLagForQosSync != 0 {
		n += 1 + sovSpec(uint64(m.AllowedBlockLagForQosSync))
	}
	if m.BlockLastUpdated != 0 {
		n += 1 + sovSpec(uint64(m.BlockLastUpdated))
	}
	l = m.MinStakeProvider.Size()
	n += 1 + l + sovSpec(uint64(l))
	if m.ProvidersTypes != 0 {
		n += 1 + sovSpec(uint64(m.ProvidersTypes))
	}
	if len(m.Imports) > 0 {
		for _, s := range m.Imports {
			l = len(s)
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if len(m.ApiCollections) > 0 {
		for _, e := range m.ApiCollections {
			l = e.Size()
			n += 2 + l + sovSpec(uint64(l))
		}
	}
	if len(m.Contributor) > 0 {
		for _, s := range m.Contributor {
			l = len(s)
			n += 2 + l + sovSpec(uint64(l))
		}
	}
	if m.ContributorPercentage != nil {
		l = m.ContributorPercentage.Size()
		n += 2 + l + sovSpec(uint64(l))
	}
	if m.Shares != 0 {
		n += 2 + sovSpec(uint64(m.Shares))
	}
	return n
}

func sovSpec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpec(x uint64) (n int) {
	return sovSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Spec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpec
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
			return fmt.Errorf("proto: Spec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Spec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReliabilityThreshold", wireType)
			}
			m.ReliabilityThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReliabilityThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataReliabilityEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
			m.DataReliabilityEnabled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockDistanceForFinalizedData", wireType)
			}
			m.BlockDistanceForFinalizedData = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockDistanceForFinalizedData |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksInFinalizationProof", wireType)
			}
			m.BlocksInFinalizationProof = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksInFinalizationProof |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTime", wireType)
			}
			m.AverageBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedBlockLagForQosSync", wireType)
			}
			m.AllowedBlockLagForQosSync = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowedBlockLagForQosSync |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockLastUpdated", wireType)
			}
			m.BlockLastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockLastUpdated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakeProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinStakeProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidersTypes", wireType)
			}
			m.ProvidersTypes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProvidersTypes |= Spec_ProvidersTypes(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imports", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imports = append(m.Imports, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiCollections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiCollections = append(m.ApiCollections, &ApiCollection{})
			if err := m.ApiCollections[len(m.ApiCollections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contributor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contributor = append(m.Contributor, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContributorPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.ContributorPercentage = &v
			if err := m.ContributorPercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			m.Shares = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shares |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpec
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
func skipSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
				return 0, ErrInvalidLengthSpec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpec = fmt.Errorf("proto: unexpected end of group")
)
