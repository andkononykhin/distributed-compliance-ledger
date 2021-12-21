// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/vendor_info_type.proto

package types

import (
	fmt "fmt"
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

type VendorInfoType struct {
	VendorID             uint64 `protobuf:"varint,1,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	VendorName           string `protobuf:"bytes,2,opt,name=vendorName,proto3" json:"vendorName,omitempty"`
	CompanyLegalName     string `protobuf:"bytes,3,opt,name=companyLegalName,proto3" json:"companyLegalName,omitempty"`
	CompanyPrefferedName string `protobuf:"bytes,4,opt,name=companyPrefferedName,proto3" json:"companyPrefferedName,omitempty"`
	VendorLandingPageURL string `protobuf:"bytes,5,opt,name=vendorLandingPageURL,proto3" json:"vendorLandingPageURL,omitempty"`
	Creator              string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *VendorInfoType) Reset()         { *m = VendorInfoType{} }
func (m *VendorInfoType) String() string { return proto.CompactTextString(m) }
func (*VendorInfoType) ProtoMessage()    {}
func (*VendorInfoType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9214c3d87077372e, []int{0}
}
func (m *VendorInfoType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VendorInfoType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VendorInfoType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VendorInfoType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VendorInfoType.Merge(m, src)
}
func (m *VendorInfoType) XXX_Size() int {
	return m.Size()
}
func (m *VendorInfoType) XXX_DiscardUnknown() {
	xxx_messageInfo_VendorInfoType.DiscardUnknown(m)
}

var xxx_messageInfo_VendorInfoType proto.InternalMessageInfo

func (m *VendorInfoType) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *VendorInfoType) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *VendorInfoType) GetCompanyLegalName() string {
	if m != nil {
		return m.CompanyLegalName
	}
	return ""
}

func (m *VendorInfoType) GetCompanyPrefferedName() string {
	if m != nil {
		return m.CompanyPrefferedName
	}
	return ""
}

func (m *VendorInfoType) GetVendorLandingPageURL() string {
	if m != nil {
		return m.VendorLandingPageURL
	}
	return ""
}

func (m *VendorInfoType) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*VendorInfoType)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.VendorInfoType")
}

func init() { proto.RegisterFile("vendorinfo/vendor_info_type.proto", fileDescriptor_9214c3d87077372e) }

var fileDescriptor_9214c3d87077372e = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x97, 0xff, 0x7f, 0x4e, 0xcd, 0x41, 0x24, 0x78, 0x28, 0x1e, 0xc2, 0xf4, 0x34, 0x84,
	0xb5, 0xa0, 0xf8, 0x05, 0xc4, 0xcb, 0x70, 0xc8, 0x28, 0xea, 0xc1, 0xcb, 0x48, 0x9b, 0xb7, 0x31,
	0xd0, 0x26, 0x25, 0xcb, 0xc4, 0xfa, 0x29, 0xfc, 0x58, 0x1e, 0x77, 0xf4, 0x28, 0xed, 0xb7, 0xf0,
	0x24, 0x4d, 0x9c, 0x2b, 0xb8, 0xdb, 0xfb, 0x3c, 0xf9, 0xbd, 0x79, 0xe1, 0x79, 0xf0, 0xc9, 0x33,
	0x28, 0xae, 0x8d, 0x54, 0x99, 0x8e, 0xfc, 0x38, 0x6f, 0xe7, 0xb9, 0xad, 0x4a, 0x08, 0x4b, 0xa3,
	0xad, 0x26, 0x97, 0xaf, 0x52, 0x24, 0x00, 0x2c, 0xcf, 0x25, 0x53, 0x29, 0x84, 0x5c, 0x2e, 0xac,
	0x91, 0xc9, 0xd2, 0x02, 0x4f, 0x75, 0x51, 0x7a, 0x37, 0x07, 0x2e, 0xc0, 0x84, 0x9b, 0xdf, 0x4e,
	0xbf, 0x10, 0x3e, 0x78, 0x70, 0x72, 0xa2, 0x32, 0x7d, 0x57, 0x95, 0x40, 0x8e, 0xf1, 0x9e, 0x07,
	0x26, 0xd7, 0x01, 0x1a, 0xa2, 0x51, 0x3f, 0xfe, 0xd5, 0x84, 0x62, 0xec, 0xe7, 0x5b, 0x56, 0x40,
	0xf0, 0x6f, 0x88, 0x46, 0xfb, 0x71, 0xc7, 0x21, 0x67, 0xf8, 0xb0, 0xbd, 0xc6, 0x54, 0x35, 0x05,
	0xc1, 0x72, 0x47, 0xfd, 0x77, 0xd4, 0x1f, 0x9f, 0x9c, 0xe3, 0xa3, 0x1f, 0x6f, 0x66, 0x20, 0xcb,
	0xc0, 0x00, 0x77, 0x7c, 0xdf, 0xf1, 0x5b, 0xdf, 0xda, 0x1d, 0x7f, 0x6d, 0xca, 0x14, 0x97, 0x4a,
	0xcc, 0x98, 0x80, 0xfb, 0x78, 0x1a, 0xec, 0xf8, 0x9d, 0x6d, 0x6f, 0x24, 0xc0, 0xbb, 0xa9, 0x01,
	0x66, 0xb5, 0x09, 0x06, 0x0e, 0x5b, 0xcb, 0x2b, 0x78, 0xaf, 0x29, 0x5a, 0xd5, 0x14, 0x7d, 0xd6,
	0x14, 0xbd, 0x35, 0xb4, 0xb7, 0x6a, 0x68, 0xef, 0xa3, 0xa1, 0xbd, 0xc7, 0x1b, 0x21, 0xed, 0xd3,
	0x32, 0x09, 0x53, 0x5d, 0x44, 0x3e, 0xd8, 0xf1, 0x3a, 0xd9, 0xa8, 0x93, 0xec, 0x78, 0x13, 0xed,
	0xd8, 0x67, 0x1b, 0xbd, 0x44, 0x9d, 0xae, 0xda, 0x7e, 0x16, 0xc9, 0xc0, 0x35, 0x74, 0xf1, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x97, 0xf7, 0xcd, 0x2d, 0xc6, 0x01, 0x00, 0x00,
}

func (m *VendorInfoType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VendorInfoType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VendorInfoType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVendorInfoType(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.VendorLandingPageURL) > 0 {
		i -= len(m.VendorLandingPageURL)
		copy(dAtA[i:], m.VendorLandingPageURL)
		i = encodeVarintVendorInfoType(dAtA, i, uint64(len(m.VendorLandingPageURL)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CompanyPrefferedName) > 0 {
		i -= len(m.CompanyPrefferedName)
		copy(dAtA[i:], m.CompanyPrefferedName)
		i = encodeVarintVendorInfoType(dAtA, i, uint64(len(m.CompanyPrefferedName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompanyLegalName) > 0 {
		i -= len(m.CompanyLegalName)
		copy(dAtA[i:], m.CompanyLegalName)
		i = encodeVarintVendorInfoType(dAtA, i, uint64(len(m.CompanyLegalName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VendorName) > 0 {
		i -= len(m.VendorName)
		copy(dAtA[i:], m.VendorName)
		i = encodeVarintVendorInfoType(dAtA, i, uint64(len(m.VendorName)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintVendorInfoType(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVendorInfoType(dAtA []byte, offset int, v uint64) int {
	offset -= sovVendorInfoType(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VendorInfoType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovVendorInfoType(uint64(m.VendorID))
	}
	l = len(m.VendorName)
	if l > 0 {
		n += 1 + l + sovVendorInfoType(uint64(l))
	}
	l = len(m.CompanyLegalName)
	if l > 0 {
		n += 1 + l + sovVendorInfoType(uint64(l))
	}
	l = len(m.CompanyPrefferedName)
	if l > 0 {
		n += 1 + l + sovVendorInfoType(uint64(l))
	}
	l = len(m.VendorLandingPageURL)
	if l > 0 {
		n += 1 + l + sovVendorInfoType(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVendorInfoType(uint64(l))
	}
	return n
}

func sovVendorInfoType(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVendorInfoType(x uint64) (n int) {
	return sovVendorInfoType(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VendorInfoType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVendorInfoType
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
			return fmt.Errorf("proto: VendorInfoType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VendorInfoType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
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
				return ErrInvalidLengthVendorInfoType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfoType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyLegalName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
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
				return ErrInvalidLengthVendorInfoType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfoType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyLegalName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyPrefferedName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
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
				return ErrInvalidLengthVendorInfoType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfoType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyPrefferedName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorLandingPageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
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
				return ErrInvalidLengthVendorInfoType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfoType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorLandingPageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfoType
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
				return ErrInvalidLengthVendorInfoType
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfoType
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVendorInfoType(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVendorInfoType
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
func skipVendorInfoType(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVendorInfoType
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
					return 0, ErrIntOverflowVendorInfoType
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
					return 0, ErrIntOverflowVendorInfoType
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
				return 0, ErrInvalidLengthVendorInfoType
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVendorInfoType
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVendorInfoType
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVendorInfoType        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVendorInfoType          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVendorInfoType = fmt.Errorf("proto: unexpected end of group")
)
