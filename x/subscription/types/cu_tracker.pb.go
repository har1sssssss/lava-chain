// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/cu_tracker.proto

package types

import (
	fmt "fmt"
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

type TrackedCu struct {
	Cu uint64 `protobuf:"varint,1,opt,name=cu,proto3" json:"cu,omitempty"`
}

func (m *TrackedCu) Reset()         { *m = TrackedCu{} }
func (m *TrackedCu) String() string { return proto.CompactTextString(m) }
func (*TrackedCu) ProtoMessage()    {}
func (*TrackedCu) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{0}
}
func (m *TrackedCu) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrackedCu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrackedCu.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrackedCu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedCu.Merge(m, src)
}
func (m *TrackedCu) XXX_Size() int {
	return m.Size()
}
func (m *TrackedCu) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedCu.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedCu proto.InternalMessageInfo

func (m *TrackedCu) GetCu() uint64 {
	if m != nil {
		return m.Cu
	}
	return 0
}

type CuTrackerTimerData struct {
	Block  uint64     `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Credit types.Coin `protobuf:"bytes,2,opt,name=credit,proto3" json:"credit"`
}

func (m *CuTrackerTimerData) Reset()         { *m = CuTrackerTimerData{} }
func (m *CuTrackerTimerData) String() string { return proto.CompactTextString(m) }
func (*CuTrackerTimerData) ProtoMessage()    {}
func (*CuTrackerTimerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{1}
}
func (m *CuTrackerTimerData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CuTrackerTimerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CuTrackerTimerData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CuTrackerTimerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CuTrackerTimerData.Merge(m, src)
}
func (m *CuTrackerTimerData) XXX_Size() int {
	return m.Size()
}
func (m *CuTrackerTimerData) XXX_DiscardUnknown() {
	xxx_messageInfo_CuTrackerTimerData.DiscardUnknown(m)
}

var xxx_messageInfo_CuTrackerTimerData proto.InternalMessageInfo

func (m *CuTrackerTimerData) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *CuTrackerTimerData) GetCredit() types.Coin {
	if m != nil {
		return m.Credit
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*TrackedCu)(nil), "lavanet.lava.subscription.TrackedCu")
	proto.RegisterType((*CuTrackerTimerData)(nil), "lavanet.lava.subscription.CuTrackerTimerData")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/cu_tracker.proto", fileDescriptor_5974e118ddf7c543)
}

var fileDescriptor_5974e118ddf7c543 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xaa, 0x54, 0xc2, 0x48, 0x0c, 0x51, 0x87, 0xb6, 0x48, 0xa6, 0xea, 0x54, 0x31,
	0xd8, 0x6a, 0x19, 0xd8, 0x1b, 0x16, 0xd6, 0xaa, 0x13, 0x0b, 0xb2, 0x6f, 0xad, 0x60, 0xb5, 0xc9,
	0x8d, 0xfc, 0x13, 0xc1, 0x5b, 0xf0, 0x58, 0x1d, 0x3b, 0x32, 0x21, 0x94, 0xbc, 0x08, 0xca, 0xcf,
	0x40, 0xa7, 0x73, 0x6c, 0x7f, 0xd6, 0x3d, 0xf7, 0xd0, 0x87, 0xa3, 0x2c, 0x65, 0xae, 0xbd, 0x68,
	0x54, 0xb8, 0xa0, 0x1c, 0x58, 0x53, 0x78, 0x83, 0xb9, 0x80, 0xf0, 0xe6, 0xad, 0x84, 0x83, 0xb6,
	0xbc, 0xb0, 0xe8, 0x31, 0x9e, 0xf6, 0x2c, 0x6f, 0x94, 0xff, 0x67, 0x67, 0x0c, 0xd0, 0x65, 0xe8,
	0x84, 0x92, 0x4e, 0x8b, 0x72, 0xa5, 0xb4, 0x97, 0x2b, 0x01, 0x68, 0xf2, 0xee, 0xeb, 0x6c, 0x9c,
	0x62, 0x8a, 0xad, 0x15, 0x8d, 0xeb, 0x6e, 0x17, 0x77, 0xf4, 0x7a, 0xd7, 0x4e, 0xd8, 0x27, 0x21,
	0xbe, 0xa5, 0x03, 0x08, 0x13, 0x32, 0x27, 0xcb, 0xe1, 0x76, 0x00, 0x61, 0x01, 0x34, 0x4e, 0x42,
	0xf7, 0x6c, 0x77, 0x26, 0xd3, 0xf6, 0x59, 0x7a, 0x19, 0x8f, 0xe9, 0x95, 0x3a, 0x22, 0x1c, 0x7a,
	0xb0, 0x3b, 0xc4, 0x4f, 0x74, 0x04, 0x56, 0xef, 0x8d, 0x9f, 0x0c, 0xe6, 0x64, 0x79, 0xb3, 0x9e,
	0xf2, 0x2e, 0x0f, 0x6f, 0xf2, 0xf0, 0x3e, 0x0f, 0x4f, 0xd0, 0xe4, 0x9b, 0xe1, 0xe9, 0xe7, 0x3e,
	0xda, 0xf6, 0xf8, 0xe6, 0xe5, 0x54, 0x31, 0x72, 0xae, 0x18, 0xf9, 0xad, 0x18, 0xf9, 0xaa, 0x59,
	0x74, 0xae, 0x59, 0xf4, 0x5d, 0xb3, 0xe8, 0x55, 0xa4, 0xc6, 0xbf, 0x07, 0xc5, 0x01, 0x33, 0x71,
	0xd1, 0x51, 0xb9, 0x16, 0x1f, 0x97, 0x45, 0xf9, 0xcf, 0x42, 0x3b, 0x35, 0x6a, 0x77, 0x7a, 0xfc,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xaf, 0x9e, 0xac, 0x52, 0x01, 0x00, 0x00,
}

func (m *TrackedCu) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrackedCu) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrackedCu) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cu != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.Cu))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CuTrackerTimerData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CuTrackerTimerData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CuTrackerTimerData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Credit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCuTracker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Block != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCuTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovCuTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TrackedCu) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cu != 0 {
		n += 1 + sovCuTracker(uint64(m.Cu))
	}
	return n
}

func (m *CuTrackerTimerData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovCuTracker(uint64(m.Block))
	}
	l = m.Credit.Size()
	n += 1 + l + sovCuTracker(uint64(l))
	return n
}

func sovCuTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCuTracker(x uint64) (n int) {
	return sovCuTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrackedCu) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: TrackedCu: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrackedCu: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cu", wireType)
			}
			m.Cu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func (m *CuTrackerTimerData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: CuTrackerTimerData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CuTrackerTimerData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
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
				return ErrInvalidLengthCuTracker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCuTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Credit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func skipCuTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
				return 0, ErrInvalidLengthCuTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCuTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCuTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCuTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCuTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCuTracker = fmt.Errorf("proto: unexpected end of group")
)
