// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pricefeed/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AddOracleProposal struct {
	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Oracles     []string `protobuf:"bytes,3,rep,name=oracles,proto3" json:"oracles,omitempty" yaml:"oracles"`
	Pairs       []string `protobuf:"bytes,4,rep,name=pairs,proto3" json:"pairs,omitempty" yaml:"pairs"`
}

func (m *AddOracleProposal) Reset()         { *m = AddOracleProposal{} }
func (m *AddOracleProposal) String() string { return proto.CompactTextString(m) }
func (*AddOracleProposal) ProtoMessage()    {}
func (*AddOracleProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_93b5d0eeea70f3f6, []int{0}
}
func (m *AddOracleProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddOracleProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddOracleProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddOracleProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddOracleProposal.Merge(m, src)
}
func (m *AddOracleProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddOracleProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddOracleProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddOracleProposal proto.InternalMessageInfo

func (m *AddOracleProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddOracleProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddOracleProposal) GetOracles() []string {
	if m != nil {
		return m.Oracles
	}
	return nil
}

func (m *AddOracleProposal) GetPairs() []string {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func init() {
	proto.RegisterType((*AddOracleProposal)(nil), "nibiru.pricefeed.v1.AddOracleProposal")
}

func init() { proto.RegisterFile("pricefeed/gov.proto", fileDescriptor_93b5d0eeea70f3f6) }

var fileDescriptor_93b5d0eeea70f3f6 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x4c,
	0x4e, 0x4d, 0x4b, 0x4d, 0x4d, 0xd1, 0x4f, 0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xce, 0xcb, 0x4c, 0xca, 0x2c, 0x2a, 0xd5, 0x83, 0xcb, 0xe9, 0x95, 0x19, 0x4a, 0x89, 0xa4,
	0xe7, 0xa7, 0xe7, 0x83, 0xe5, 0xf5, 0x41, 0x2c, 0x88, 0x52, 0xa5, 0x85, 0x8c, 0x5c, 0x82, 0x8e,
	0x29, 0x29, 0xfe, 0x45, 0x89, 0xc9, 0x39, 0xa9, 0x01, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89, 0x39,
	0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x10, 0x8e, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e,
	0x9e, 0x04, 0x13, 0x58, 0x0e, 0x59, 0x48, 0x48, 0x87, 0x8b, 0x3d, 0x1f, 0x6c, 0x52, 0xb1, 0x04,
	0xb3, 0x02, 0xb3, 0x06, 0xa7, 0x93, 0xd0, 0xa7, 0x7b, 0xf2, 0x7c, 0x95, 0x89, 0xb9, 0x39, 0x56,
	0x4a, 0x50, 0x09, 0xa5, 0x20, 0x98, 0x12, 0x21, 0x35, 0x2e, 0xd6, 0x82, 0xc4, 0xcc, 0xa2, 0x62,
	0x09, 0x16, 0xb0, 0x5a, 0x81, 0x4f, 0xf7, 0xe4, 0x79, 0x20, 0x6a, 0xc1, 0xc2, 0x4a, 0x41, 0x10,
	0x69, 0x27, 0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4f, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xf7, 0x03, 0xfb, 0xd9, 0x39, 0x23, 0x31,
	0x33, 0x4f, 0x1f, 0xe2, 0x7f, 0xfd, 0x0a, 0x7d, 0x44, 0xe8, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x7d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xab, 0x16, 0x50, 0x37, 0x01,
	0x00, 0x00,
}

func (m *AddOracleProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddOracleProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddOracleProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pairs) > 0 {
		for iNdEx := len(m.Pairs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Pairs[iNdEx])
			copy(dAtA[i:], m.Pairs[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.Pairs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Oracles) > 0 {
		for iNdEx := len(m.Oracles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Oracles[iNdEx])
			copy(dAtA[i:], m.Oracles[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.Oracles[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddOracleProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Oracles) > 0 {
		for _, s := range m.Oracles {
			l = len(s)
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.Pairs) > 0 {
		for _, s := range m.Pairs {
			l = len(s)
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddOracleProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: AddOracleProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddOracleProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracles = append(m.Oracles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pairs = append(m.Pairs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
