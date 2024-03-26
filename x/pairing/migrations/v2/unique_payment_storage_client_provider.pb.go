// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/pairing/unique_payment_storage_client_provider.proto

package v2

import (
	fmt "fmt"
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

type UniquePaymentStorageClientProvider struct {
	Index  string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block  uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	UsedCU uint64 `protobuf:"varint,3,opt,name=usedCU,proto3" json:"usedCU,omitempty"`
}

func (m *UniquePaymentStorageClientProvider) Reset()         { *m = UniquePaymentStorageClientProvider{} }
func (m *UniquePaymentStorageClientProvider) String() string { return proto.CompactTextString(m) }
func (*UniquePaymentStorageClientProvider) ProtoMessage()    {}
func (*UniquePaymentStorageClientProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_de8e31c04724b0b7, []int{0}
}
func (m *UniquePaymentStorageClientProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UniquePaymentStorageClientProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UniquePaymentStorageClientProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UniquePaymentStorageClientProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniquePaymentStorageClientProvider.Merge(m, src)
}
func (m *UniquePaymentStorageClientProvider) XXX_Size() int {
	return m.Size()
}
func (m *UniquePaymentStorageClientProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_UniquePaymentStorageClientProvider.DiscardUnknown(m)
}

var xxx_messageInfo_UniquePaymentStorageClientProvider proto.InternalMessageInfo

func (m *UniquePaymentStorageClientProvider) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *UniquePaymentStorageClientProvider) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *UniquePaymentStorageClientProvider) GetUsedCU() uint64 {
	if m != nil {
		return m.UsedCU
	}
	return 0
}

func init() {
	proto.RegisterType((*UniquePaymentStorageClientProvider)(nil), "lavanet.lava.pairing.UniquePaymentStorageClientProvider")
}

func init() {
	proto.RegisterFile("lavanet/lava/pairing/unique_payment_storage_client_provider.proto", fileDescriptor_de8e31c04724b0b7)
}

var fileDescriptor_de8e31c04724b0b7 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcc, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0x05, 0x89, 0x99, 0x45, 0x99, 0x79, 0xe9, 0xfa, 0xa5,
	0x79, 0x99, 0x85, 0xa5, 0xa9, 0xf1, 0x05, 0x89, 0x95, 0xb9, 0xa9, 0x79, 0x25, 0xf1, 0xc5, 0x25,
	0xf9, 0x45, 0x89, 0xe9, 0xa9, 0xf1, 0xc9, 0x39, 0x99, 0x20, 0x6e, 0x41, 0x51, 0x7e, 0x59, 0x66,
	0x4a, 0x6a, 0x91, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x08, 0xd4, 0x08, 0x3d, 0x10, 0xad,
	0x07, 0x35, 0x42, 0x29, 0x83, 0x4b, 0x29, 0x14, 0x6c, 0x4a, 0x00, 0xc4, 0x90, 0x60, 0x88, 0x19,
	0xce, 0x60, 0x23, 0x02, 0xa0, 0x26, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x20, 0xd1, 0xa4, 0x9c, 0xfc, 0xe4, 0x6c, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0xad, 0xb4, 0x38, 0x35, 0xc5,
	0x39, 0x54, 0x82, 0x19, 0x2c, 0x0c, 0xe5, 0x39, 0x39, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e,
	0x8a, 0x3f, 0x2b, 0xe0, 0x3e, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xc4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x9a, 0x51, 0x54, 0x0e, 0x01, 0x00, 0x00,
}

func (m *UniquePaymentStorageClientProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UniquePaymentStorageClientProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UniquePaymentStorageClientProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsedCU != 0 {
		i = encodeVarintUniquePaymentStorageClientProvider(dAtA, i, uint64(m.UsedCU))
		i--
		dAtA[i] = 0x18
	}
	if m.Block != 0 {
		i = encodeVarintUniquePaymentStorageClientProvider(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintUniquePaymentStorageClientProvider(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUniquePaymentStorageClientProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovUniquePaymentStorageClientProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UniquePaymentStorageClientProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovUniquePaymentStorageClientProvider(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovUniquePaymentStorageClientProvider(uint64(m.Block))
	}
	if m.UsedCU != 0 {
		n += 1 + sovUniquePaymentStorageClientProvider(uint64(m.UsedCU))
	}
	return n
}

func sovUniquePaymentStorageClientProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUniquePaymentStorageClientProvider(x uint64) (n int) {
	return sovUniquePaymentStorageClientProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UniquePaymentStorageClientProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUniquePaymentStorageClientProvider
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
			return fmt.Errorf("proto: UniquePaymentStorageClientProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UniquePaymentStorageClientProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniquePaymentStorageClientProvider
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
				return ErrInvalidLengthUniquePaymentStorageClientProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniquePaymentStorageClientProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniquePaymentStorageClientProvider
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedCU", wireType)
			}
			m.UsedCU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniquePaymentStorageClientProvider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedCU |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUniquePaymentStorageClientProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUniquePaymentStorageClientProvider
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
func skipUniquePaymentStorageClientProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUniquePaymentStorageClientProvider
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
					return 0, ErrIntOverflowUniquePaymentStorageClientProvider
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
					return 0, ErrIntOverflowUniquePaymentStorageClientProvider
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
				return 0, ErrInvalidLengthUniquePaymentStorageClientProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUniquePaymentStorageClientProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUniquePaymentStorageClientProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUniquePaymentStorageClientProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUniquePaymentStorageClientProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUniquePaymentStorageClientProvider = fmt.Errorf("proto: unexpected end of group")
)
