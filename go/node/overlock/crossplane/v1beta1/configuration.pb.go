// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: overlock/crossplane/v1beta1/configuration.proto

package v1beta1

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

// Configuration holds the configuration details.
type Configuration struct {
	Id       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Creator  string    `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bb9d8cd8d5cbea, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return m.Size()
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Configuration) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Configuration) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "overlock.crossplane.v1beta1.Configuration")
}

func init() {
	proto.RegisterFile("overlock/crossplane/v1beta1/configuration.proto", fileDescriptor_b7bb9d8cd8d5cbea)
}

var fileDescriptor_b7bb9d8cd8d5cbea = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcf, 0x2f, 0x4b, 0x2d,
	0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x2e, 0xc8, 0x49, 0xcc, 0x4b, 0xd5,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2d,
	0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0x69, 0xd0,
	0x43, 0x68, 0xd0, 0x83, 0x6a, 0x90, 0xd2, 0xc2, 0x67, 0x5a, 0x6e, 0x6a, 0x49, 0x62, 0x4a, 0x62,
	0x49, 0x22, 0xc4, 0x20, 0xa5, 0x1a, 0x2e, 0x5e, 0x67, 0x64, 0xf3, 0x85, 0xf8, 0xb8, 0x98, 0x32,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x1c, 0xb9, 0x38, 0x60,
	0x5a, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x54, 0xf5, 0xf0, 0x58, 0xae, 0xe7, 0x0b, 0x55,
	0x1c, 0x04, 0xd7, 0x26, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1,
	0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x3a, 0xc5, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x6b, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x79, 0x6a, 0x92, 0x6e, 0x71, 0x6a, 0x59, 0x6a, 0x1e, 0x3c, 0x98, 0x74, 0x13, 0x0b, 0x32, 0xf5,
	0xd3, 0xf3, 0xf5, 0xf3, 0xf2, 0x53, 0x52, 0xf1, 0x85, 0x5d, 0x12, 0x1b, 0xd8, 0x97, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x7e, 0x4d, 0x81, 0x61, 0x01, 0x00, 0x00,
}

func (m *Configuration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Configuration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintConfiguration(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfiguration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintConfiguration(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Configuration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovConfiguration(uint64(m.Id))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovConfiguration(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovConfiguration(uint64(l))
	}
	return n
}

func sovConfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfiguration(x uint64) (n int) {
	return sovConfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Configuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfiguration
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
func skipConfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfiguration
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
					return 0, ErrIntOverflowConfiguration
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
					return 0, ErrIntOverflowConfiguration
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
				return 0, ErrInvalidLengthConfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfiguration = fmt.Errorf("proto: unexpected end of group")
)
