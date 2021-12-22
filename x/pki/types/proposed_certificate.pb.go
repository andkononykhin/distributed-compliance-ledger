// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/proposed_certificate.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ProposedCertificate struct {
	Subject      string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId string   `protobuf:"bytes,2,opt,name=subject_key_id,json=subjectKeyId,proto3" json:"subject_key_id,omitempty"`
	PemCert      string   `protobuf:"bytes,3,opt,name=pem_cert,json=pemCert,proto3" json:"pem_cert,omitempty"`
	SerialNumber string   `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Owner        string   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Approvals    []string `protobuf:"bytes,6,rep,name=approvals,proto3" json:"approvals,omitempty"`
}

func (m *ProposedCertificate) Reset()         { *m = ProposedCertificate{} }
func (m *ProposedCertificate) String() string { return proto.CompactTextString(m) }
func (*ProposedCertificate) ProtoMessage()    {}
func (*ProposedCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_589d443e04a789ec, []int{0}
}
func (m *ProposedCertificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedCertificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedCertificate.Merge(m, src)
}
func (m *ProposedCertificate) XXX_Size() int {
	return m.Size()
}
func (m *ProposedCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedCertificate proto.InternalMessageInfo

func (m *ProposedCertificate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ProposedCertificate) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *ProposedCertificate) GetPemCert() string {
	if m != nil {
		return m.PemCert
	}
	return ""
}

func (m *ProposedCertificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *ProposedCertificate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProposedCertificate) GetApprovals() []string {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposedCertificate)(nil), "zigbeealliance.distributedcomplianceledger.pki.ProposedCertificate")
}

func init() { proto.RegisterFile("pki/proposed_certificate.proto", fileDescriptor_589d443e04a789ec) }

var fileDescriptor_589d443e04a789ec = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x9b, 0xbf, 0x7f, 0x0b, 0xb5, 0x80, 0xc1, 0x30, 0xb8, 0x0c, 0x56, 0x05, 0x0c, 0x2c,
	0x49, 0x06, 0x24, 0x76, 0x0a, 0x0b, 0x42, 0x42, 0xa8, 0x6c, 0x0c, 0x44, 0x89, 0x7d, 0x09, 0xa6,
	0x49, 0x6c, 0xd9, 0x0e, 0x50, 0x9e, 0x82, 0x87, 0xe1, 0x21, 0x18, 0x2b, 0x26, 0x46, 0xd4, 0x3e,
	0x07, 0x12, 0xaa, 0x9d, 0xaa, 0x4c, 0x6c, 0xb9, 0x27, 0xe7, 0xdc, 0x73, 0xf5, 0x19, 0x51, 0x35,
	0x16, 0xb1, 0xd2, 0x52, 0x49, 0x03, 0x3c, 0x61, 0xa0, 0xad, 0xb8, 0x13, 0x2c, 0xb5, 0x10, 0x29,
	0x2d, 0xad, 0xc4, 0xd1, 0x8b, 0xc8, 0x33, 0x80, 0xb4, 0x28, 0x44, 0x5a, 0x31, 0x88, 0xb8, 0x30,
	0x56, 0x8b, 0xac, 0xb6, 0xc0, 0x99, 0x2c, 0x95, 0x57, 0x0b, 0xe0, 0x39, 0xe8, 0x48, 0x8d, 0xc5,
	0x6e, 0x9f, 0x49, 0x53, 0x4a, 0x93, 0xb8, 0x74, 0xec, 0x07, 0xbf, 0x6a, 0xef, 0x3b, 0x40, 0xdb,
	0x57, 0x4d, 0xd3, 0xe9, 0xaa, 0x08, 0x13, 0xb4, 0x66, 0xea, 0xec, 0x01, 0x98, 0x25, 0xc1, 0x20,
	0x38, 0xec, 0x8d, 0x96, 0x23, 0x3e, 0x40, 0x5b, 0xcd, 0x67, 0x32, 0x86, 0x49, 0x22, 0x38, 0xf9,
	0xe7, 0x0c, 0x1b, 0x8d, 0x7a, 0x01, 0x93, 0x73, 0x8e, 0xfb, 0x68, 0x5d, 0x41, 0xe9, 0x6e, 0x27,
	0x6d, 0xbf, 0x40, 0x41, 0xb9, 0x68, 0xc0, 0xfb, 0x68, 0xd3, 0x80, 0x16, 0x69, 0x91, 0x54, 0x75,
	0x99, 0x81, 0x26, 0xff, 0x9b, 0xbc, 0x13, 0x2f, 0x9d, 0x86, 0x23, 0xd4, 0x91, 0x4f, 0x15, 0x68,
	0xd2, 0x59, 0xfc, 0x1c, 0x92, 0x8f, 0xb7, 0x70, 0xa7, 0x39, 0xfc, 0x84, 0x73, 0x0d, 0xc6, 0x5c,
	0x5b, 0x2d, 0xaa, 0x7c, 0xe4, 0x6d, 0xf8, 0x18, 0xf5, 0x52, 0xa5, 0xb4, 0x7c, 0x4c, 0x0b, 0x43,
	0xba, 0x83, 0xf6, 0x9f, 0x99, 0x95, 0x75, 0x78, 0xfb, 0x3e, 0xa3, 0xc1, 0x74, 0x46, 0x83, 0xaf,
	0x19, 0x0d, 0x5e, 0xe7, 0xb4, 0x35, 0x9d, 0xd3, 0xd6, 0xe7, 0x9c, 0xb6, 0x6e, 0xce, 0x72, 0x61,
	0xef, 0xeb, 0x2c, 0x62, 0xb2, 0x8c, 0x3d, 0xef, 0x70, 0x09, 0x3c, 0xfe, 0x05, 0x3c, 0x5c, 0x11,
	0x0f, 0x3d, 0xf2, 0xf8, 0x39, 0x5e, 0xbc, 0x9f, 0x9d, 0x28, 0x30, 0x59, 0xd7, 0x61, 0x3e, 0xfa,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x86, 0x65, 0x24, 0xd3, 0x01, 0x00, 0x00,
}

func (m *ProposedCertificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedCertificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedCertificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Approvals[iNdEx])
			copy(dAtA[i:], m.Approvals[iNdEx])
			i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.Approvals[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintProposedCertificate(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposedCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedCertificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProposedCertificate(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, s := range m.Approvals {
			l = len(s)
			n += 1 + l + sovProposedCertificate(uint64(l))
		}
	}
	return n
}

func sovProposedCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedCertificate(x uint64) (n int) {
	return sovProposedCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedCertificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedCertificate
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
			return fmt.Errorf("proto: ProposedCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificate
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
				return ErrInvalidLengthProposedCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposedCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedCertificate
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
func skipProposedCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedCertificate
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
					return 0, ErrIntOverflowProposedCertificate
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
					return 0, ErrIntOverflowProposedCertificate
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
				return 0, ErrInvalidLengthProposedCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedCertificate = fmt.Errorf("proto: unexpected end of group")
)
