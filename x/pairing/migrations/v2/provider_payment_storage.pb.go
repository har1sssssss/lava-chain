// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/pairing/provider_payment_storage.proto

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

type ProviderPaymentStorage struct {
	Index                                  string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Epoch                                  uint64   `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	UniquePaymentStorageClientProviderKeys []string `protobuf:"bytes,5,rep,name=uniquePaymentStorageClientProviderKeys,proto3" json:"uniquePaymentStorageClientProviderKeys,omitempty"`
	ComplainersTotalCu                     uint64   `protobuf:"varint,6,opt,name=complainersTotalCu,proto3" json:"complainersTotalCu,omitempty"`
}

func (m *ProviderPaymentStorage) Reset()         { *m = ProviderPaymentStorage{} }
func (m *ProviderPaymentStorage) String() string { return proto.CompactTextString(m) }
func (*ProviderPaymentStorage) ProtoMessage()    {}
func (*ProviderPaymentStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_39d1051245241a4e, []int{0}
}
func (m *ProviderPaymentStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderPaymentStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderPaymentStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderPaymentStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderPaymentStorage.Merge(m, src)
}
func (m *ProviderPaymentStorage) XXX_Size() int {
	return m.Size()
}
func (m *ProviderPaymentStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderPaymentStorage.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderPaymentStorage proto.InternalMessageInfo

func (m *ProviderPaymentStorage) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ProviderPaymentStorage) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ProviderPaymentStorage) GetUniquePaymentStorageClientProviderKeys() []string {
	if m != nil {
		return m.UniquePaymentStorageClientProviderKeys
	}
	return nil
}

func (m *ProviderPaymentStorage) GetComplainersTotalCu() uint64 {
	if m != nil {
		return m.ComplainersTotalCu
	}
	return 0
}

func init() {
	proto.RegisterType((*ProviderPaymentStorage)(nil), "lavanet.lava.pairing.ProviderPaymentStorage")
}

func init() {
	proto.RegisterFile("lavanet/lava/pairing/provider_payment_storage.proto", fileDescriptor_39d1051245241a4e)
}

var fileDescriptor_39d1051245241a4e = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xce, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0x05, 0x89, 0x99, 0x45, 0x99, 0x79, 0xe9, 0xfa, 0x05,
	0x45, 0xf9, 0x65, 0x99, 0x29, 0xa9, 0x45, 0xf1, 0x05, 0x89, 0x95, 0xb9, 0xa9, 0x79, 0x25, 0xf1,
	0xc5, 0x25, 0xf9, 0x45, 0x89, 0xe9, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x22, 0x50,
	0x4d, 0x7a, 0x20, 0x5a, 0x0f, 0xaa, 0x49, 0xca, 0x11, 0xab, 0x51, 0xa5, 0x79, 0x99, 0x85, 0xa5,
	0xa9, 0xe8, 0x06, 0xc5, 0x27, 0xe7, 0x64, 0x82, 0xb8, 0x30, 0x8b, 0x20, 0x06, 0x2b, 0xdd, 0x60,
	0xe4, 0x12, 0x0b, 0x80, 0x0a, 0x05, 0x40, 0x74, 0x04, 0x43, 0x34, 0x08, 0x89, 0x70, 0xb1, 0x66,
	0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x20, 0xd1, 0xd4,
	0x82, 0xfc, 0xe4, 0x0c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x28, 0x8c, 0x4b,
	0x0d, 0x62, 0x2d, 0xaa, 0x19, 0xce, 0x60, 0x3b, 0x61, 0xe6, 0x7b, 0xa7, 0x56, 0x16, 0x4b, 0xb0,
	0x2a, 0x30, 0x6b, 0x70, 0x06, 0x11, 0xa9, 0x5a, 0x48, 0x8f, 0x4b, 0x28, 0x39, 0x3f, 0xb7, 0x20,
	0x27, 0x31, 0x33, 0x2f, 0xb5, 0xa8, 0x38, 0x24, 0xbf, 0x24, 0x31, 0xc7, 0xb9, 0x54, 0x82, 0x0d,
	0x6c, 0x35, 0x16, 0x19, 0x2f, 0x16, 0x0e, 0x26, 0x01, 0x66, 0x2f, 0x16, 0x0e, 0x16, 0x01, 0x56,
	0x27, 0xc7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x4f, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x09, 0xc2, 0x0a, 0x78, 0x20, 0x96, 0x54,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x03, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xda, 0x1f,
	0x87, 0x29, 0xb4, 0x01, 0x00, 0x00,
}

func (m *ProviderPaymentStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderPaymentStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderPaymentStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ComplainersTotalCu != 0 {
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(m.ComplainersTotalCu))
		i--
		dAtA[i] = 0x30
	}
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for iNdEx := len(m.UniquePaymentStorageClientProviderKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UniquePaymentStorageClientProviderKeys[iNdEx])
			copy(dAtA[i:], m.UniquePaymentStorageClientProviderKeys[iNdEx])
			i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.UniquePaymentStorageClientProviderKeys[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProviderPaymentStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovProviderPaymentStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderPaymentStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProviderPaymentStorage(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovProviderPaymentStorage(uint64(m.Epoch))
	}
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for _, s := range m.UniquePaymentStorageClientProviderKeys {
			l = len(s)
			n += 1 + l + sovProviderPaymentStorage(uint64(l))
		}
	}
	if m.ComplainersTotalCu != 0 {
		n += 1 + sovProviderPaymentStorage(uint64(m.ComplainersTotalCu))
	}
	return n
}

func sovProviderPaymentStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProviderPaymentStorage(x uint64) (n int) {
	return sovProviderPaymentStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderPaymentStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProviderPaymentStorage
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
			return fmt.Errorf("proto: ProviderPaymentStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderPaymentStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniquePaymentStorageClientProviderKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniquePaymentStorageClientProviderKeys = append(m.UniquePaymentStorageClientProviderKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComplainersTotalCu", wireType)
			}
			m.ComplainersTotalCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComplainersTotalCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProviderPaymentStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProviderPaymentStorage
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
func skipProviderPaymentStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
				return 0, ErrInvalidLengthProviderPaymentStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProviderPaymentStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProviderPaymentStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProviderPaymentStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProviderPaymentStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProviderPaymentStorage = fmt.Errorf("proto: unexpected end of group")
)
