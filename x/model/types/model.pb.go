// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model/model.proto

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

type Model struct {
	Vid                                        int32  `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Pid                                        int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	DeviceTypeId                               int32  `protobuf:"varint,3,opt,name=device_type_id,json=deviceTypeId,proto3" json:"device_type_id,omitempty"`
	ProductName                                string `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductLabel                               string `protobuf:"bytes,5,opt,name=product_label,json=productLabel,proto3" json:"product_label,omitempty"`
	PartNumber                                 string `protobuf:"bytes,6,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	CommissioningCustomFlow                    int32  `protobuf:"varint,7,opt,name=commissioning_custom_flow,json=commissioningCustomFlow,proto3" json:"commissioning_custom_flow,omitempty"`
	CommissioningCustomFlowUrl                 string `protobuf:"bytes,8,opt,name=commissioning_custom_flow_url,json=commissioningCustomFlowUrl,proto3" json:"commissioning_custom_flow_url,omitempty"`
	CommissioningModeInitialStepsHint          uint32 `protobuf:"varint,9,opt,name=commissioning_mode_initial_steps_hint,json=commissioningModeInitialStepsHint,proto3" json:"commissioning_mode_initial_steps_hint,omitempty"`
	CommissioningModeInitialStepsInstruction   string `protobuf:"bytes,10,opt,name=commissioning_mode_initial_steps_instruction,json=commissioningModeInitialStepsInstruction,proto3" json:"commissioning_mode_initial_steps_instruction,omitempty"`
	CommissioningModeSecondaryStepsHint        uint32 `protobuf:"varint,11,opt,name=commissioning_mode_secondary_steps_hint,json=commissioningModeSecondaryStepsHint,proto3" json:"commissioning_mode_secondary_steps_hint,omitempty"`
	CommissioningModeSecondaryStepsInstruction string `protobuf:"bytes,12,opt,name=commissioning_mode_secondary_steps_instruction,json=commissioningModeSecondaryStepsInstruction,proto3" json:"commissioning_mode_secondary_steps_instruction,omitempty"`
	UserManualUrl                              string `protobuf:"bytes,13,opt,name=user_manual_url,json=userManualUrl,proto3" json:"user_manual_url,omitempty"`
	SupportUrl                                 string `protobuf:"bytes,14,opt,name=support_url,json=supportUrl,proto3" json:"support_url,omitempty"`
	ProductUrl                                 string `protobuf:"bytes,15,opt,name=product_url,json=productUrl,proto3" json:"product_url,omitempty"`
	Creator                                    string `protobuf:"bytes,16,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{0}
}
func (m *Model) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Model.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return m.Size()
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *Model) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Model) GetDeviceTypeId() int32 {
	if m != nil {
		return m.DeviceTypeId
	}
	return 0
}

func (m *Model) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *Model) GetProductLabel() string {
	if m != nil {
		return m.ProductLabel
	}
	return ""
}

func (m *Model) GetPartNumber() string {
	if m != nil {
		return m.PartNumber
	}
	return ""
}

func (m *Model) GetCommissioningCustomFlow() int32 {
	if m != nil {
		return m.CommissioningCustomFlow
	}
	return 0
}

func (m *Model) GetCommissioningCustomFlowUrl() string {
	if m != nil {
		return m.CommissioningCustomFlowUrl
	}
	return ""
}

func (m *Model) GetCommissioningModeInitialStepsHint() uint32 {
	if m != nil {
		return m.CommissioningModeInitialStepsHint
	}
	return 0
}

func (m *Model) GetCommissioningModeInitialStepsInstruction() string {
	if m != nil {
		return m.CommissioningModeInitialStepsInstruction
	}
	return ""
}

func (m *Model) GetCommissioningModeSecondaryStepsHint() uint32 {
	if m != nil {
		return m.CommissioningModeSecondaryStepsHint
	}
	return 0
}

func (m *Model) GetCommissioningModeSecondaryStepsInstruction() string {
	if m != nil {
		return m.CommissioningModeSecondaryStepsInstruction
	}
	return ""
}

func (m *Model) GetUserManualUrl() string {
	if m != nil {
		return m.UserManualUrl
	}
	return ""
}

func (m *Model) GetSupportUrl() string {
	if m != nil {
		return m.SupportUrl
	}
	return ""
}

func (m *Model) GetProductUrl() string {
	if m != nil {
		return m.ProductUrl
	}
	return ""
}

func (m *Model) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Model)(nil), "zigbeealliance.distributedcomplianceledger.model.Model")
}

func init() { proto.RegisterFile("model/model.proto", fileDescriptor_312ac5bcab6cbb43) }

var fileDescriptor_312ac5bcab6cbb43 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x3e,
	0x1c, 0xc6, 0x97, 0xdf, 0x7e, 0xed, 0x98, 0xdb, 0x6e, 0x23, 0x42, 0xc2, 0x9b, 0x44, 0xd6, 0x31,
	0xfe, 0x54, 0x88, 0x36, 0x08, 0x6e, 0xdc, 0x36, 0xa4, 0x89, 0x4a, 0x6c, 0x42, 0xed, 0xb8, 0x70,
	0xc0, 0x38, 0xb1, 0xe9, 0x2c, 0x39, 0xb6, 0x65, 0x3b, 0x1b, 0xe5, 0x55, 0xf0, 0x62, 0x78, 0x07,
	0x5c, 0x38, 0x4e, 0x9c, 0x38, 0xa2, 0xf6, 0x8d, 0x20, 0xdb, 0x8d, 0xda, 0x0a, 0xc6, 0xb8, 0x44,
	0xc9, 0xe3, 0xcf, 0xf7, 0xc9, 0x23, 0xdb, 0x0f, 0xb8, 0x59, 0x48, 0x42, 0x79, 0xea, 0x9f, 0x3d,
	0xa5, 0xa5, 0x95, 0xf1, 0x93, 0x4f, 0x6c, 0x94, 0x51, 0x8a, 0x39, 0x67, 0x58, 0xe4, 0xb4, 0x47,
	0x98, 0xb1, 0x9a, 0x65, 0xa5, 0xa5, 0x24, 0x97, 0x85, 0x0a, 0x2a, 0xa7, 0x64, 0x44, 0x75, 0xcf,
	0xcf, 0xed, 0x6c, 0xe7, 0xd2, 0x14, 0xd2, 0x20, 0x3f, 0x9f, 0x86, 0x8f, 0x60, 0x76, 0xf7, 0x6b,
	0x1d, 0xd4, 0x8e, 0x1d, 0x14, 0x6f, 0x81, 0xd5, 0x73, 0x46, 0x60, 0xd4, 0x8e, 0x3a, 0xb5, 0x81,
	0x7b, 0x75, 0x8a, 0x62, 0x04, 0xfe, 0x17, 0x14, 0xc5, 0x48, 0x7c, 0x0f, 0x6c, 0x10, 0x7a, 0xce,
	0x72, 0x8a, 0xec, 0x58, 0x51, 0xc4, 0x08, 0x5c, 0xf5, 0x8b, 0xcd, 0xa0, 0x9e, 0x8e, 0x15, 0xed,
	0x93, 0x78, 0x0f, 0x34, 0x95, 0x96, 0xa4, 0xcc, 0x2d, 0x12, 0xb8, 0xa0, 0xf0, 0xff, 0x76, 0xd4,
	0x59, 0x1f, 0x34, 0x66, 0xda, 0x09, 0x2e, 0x68, 0xbc, 0x0f, 0x5a, 0x15, 0xc2, 0x71, 0x46, 0x39,
	0xac, 0x79, 0xa6, 0x9a, 0x7b, 0xe5, 0xb4, 0x78, 0x17, 0x34, 0x14, 0xd6, 0x16, 0x89, 0xb2, 0xc8,
	0xa8, 0x86, 0x75, 0x8f, 0x00, 0x27, 0x9d, 0x78, 0x25, 0x7e, 0x0e, 0xb6, 0x73, 0x59, 0x14, 0xcc,
	0x18, 0x26, 0x05, 0x13, 0x23, 0x94, 0x97, 0xc6, 0xca, 0x02, 0x7d, 0xe0, 0xf2, 0x02, 0xae, 0xf9,
	0x64, 0xb7, 0x97, 0x80, 0x17, 0x7e, 0xfd, 0x88, 0xcb, 0x8b, 0xf8, 0x00, 0xdc, 0xb9, 0x72, 0x16,
	0x95, 0x9a, 0xc3, 0x1b, 0xfe, 0x77, 0x3b, 0x57, 0xcc, 0xbf, 0xd1, 0x3c, 0x7e, 0x0d, 0xee, 0x2f,
	0x5b, 0xb8, 0xdd, 0x46, 0x4c, 0x30, 0xcb, 0x30, 0x47, 0xc6, 0x52, 0x65, 0xd0, 0x19, 0x13, 0x16,
	0xae, 0xb7, 0xa3, 0x4e, 0x6b, 0xb0, 0xb7, 0x04, 0xbb, 0x4d, 0xef, 0x07, 0x74, 0xe8, 0xc8, 0x97,
	0x4c, 0xd8, 0xf8, 0x1d, 0x78, 0x7c, 0xad, 0x23, 0x13, 0xc6, 0xea, 0x32, 0xb7, 0x4c, 0x0a, 0x08,
	0x7c, 0xc6, 0xce, 0x5f, 0x8d, 0xfb, 0x73, 0x3e, 0x3e, 0x05, 0x0f, 0xff, 0xe0, 0x6f, 0x68, 0x2e,
	0x05, 0xc1, 0x7a, 0xbc, 0x98, 0xb9, 0xe1, 0x33, 0xef, 0xff, 0x66, 0x3d, 0xac, 0xe0, 0x79, 0xea,
	0x0c, 0xf4, 0xfe, 0xc1, 0x75, 0x31, 0x77, 0xd3, 0xe7, 0x7e, 0x74, 0x8d, 0xf9, 0x62, 0xf2, 0x07,
	0x60, 0xb3, 0x34, 0x54, 0xa3, 0x02, 0x8b, 0x12, 0x73, 0x7f, 0x40, 0x2d, 0x6f, 0xd2, 0x72, 0xf2,
	0xb1, 0x57, 0xdd, 0x99, 0xec, 0x82, 0x86, 0x29, 0x95, 0x92, 0xda, 0x7a, 0x66, 0x23, 0xdc, 0x99,
	0x99, 0x34, 0x03, 0xaa, 0x9b, 0xe7, 0x80, 0xcd, 0xd9, 0xa5, 0x0a, 0x92, 0x03, 0x9e, 0x82, 0xb5,
	0x5c, 0x53, 0x6c, 0xa5, 0x86, 0x5b, 0x6e, 0xf1, 0x10, 0x7e, 0xff, 0xd2, 0xbd, 0x35, 0x2b, 0xcd,
	0x01, 0x21, 0x9a, 0x1a, 0x33, 0xb4, 0x9a, 0x89, 0xd1, 0xa0, 0x02, 0x0f, 0xdf, 0x7f, 0x9b, 0x24,
	0xd1, 0xe5, 0x24, 0x89, 0x7e, 0x4e, 0x92, 0xe8, 0xf3, 0x34, 0x59, 0xb9, 0x9c, 0x26, 0x2b, 0x3f,
	0xa6, 0xc9, 0xca, 0xdb, 0xa3, 0x11, 0xb3, 0x67, 0x65, 0xe6, 0x36, 0x28, 0x0d, 0xbd, 0xed, 0x56,
	0xc5, 0x4d, 0x17, 0x8a, 0xdb, 0x9d, 0x37, 0xb7, 0x1b, 0xaa, 0x9b, 0x7e, 0x0c, 0xa5, 0x4f, 0x5d,
	0xd3, 0x4c, 0x56, 0xf7, 0x75, 0x7d, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x31, 0xd6, 0x4b, 0x54,
	0x10, 0x04, 0x00, 0x00,
}

func (m *Model) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Model) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Model) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModel(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.ProductUrl) > 0 {
		i -= len(m.ProductUrl)
		copy(dAtA[i:], m.ProductUrl)
		i = encodeVarintModel(dAtA, i, uint64(len(m.ProductUrl)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.SupportUrl) > 0 {
		i -= len(m.SupportUrl)
		copy(dAtA[i:], m.SupportUrl)
		i = encodeVarintModel(dAtA, i, uint64(len(m.SupportUrl)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.UserManualUrl) > 0 {
		i -= len(m.UserManualUrl)
		copy(dAtA[i:], m.UserManualUrl)
		i = encodeVarintModel(dAtA, i, uint64(len(m.UserManualUrl)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.CommissioningModeSecondaryStepsInstruction) > 0 {
		i -= len(m.CommissioningModeSecondaryStepsInstruction)
		copy(dAtA[i:], m.CommissioningModeSecondaryStepsInstruction)
		i = encodeVarintModel(dAtA, i, uint64(len(m.CommissioningModeSecondaryStepsInstruction)))
		i--
		dAtA[i] = 0x62
	}
	if m.CommissioningModeSecondaryStepsHint != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.CommissioningModeSecondaryStepsHint))
		i--
		dAtA[i] = 0x58
	}
	if len(m.CommissioningModeInitialStepsInstruction) > 0 {
		i -= len(m.CommissioningModeInitialStepsInstruction)
		copy(dAtA[i:], m.CommissioningModeInitialStepsInstruction)
		i = encodeVarintModel(dAtA, i, uint64(len(m.CommissioningModeInitialStepsInstruction)))
		i--
		dAtA[i] = 0x52
	}
	if m.CommissioningModeInitialStepsHint != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.CommissioningModeInitialStepsHint))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CommissioningCustomFlowUrl) > 0 {
		i -= len(m.CommissioningCustomFlowUrl)
		copy(dAtA[i:], m.CommissioningCustomFlowUrl)
		i = encodeVarintModel(dAtA, i, uint64(len(m.CommissioningCustomFlowUrl)))
		i--
		dAtA[i] = 0x42
	}
	if m.CommissioningCustomFlow != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.CommissioningCustomFlow))
		i--
		dAtA[i] = 0x38
	}
	if len(m.PartNumber) > 0 {
		i -= len(m.PartNumber)
		copy(dAtA[i:], m.PartNumber)
		i = encodeVarintModel(dAtA, i, uint64(len(m.PartNumber)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ProductLabel) > 0 {
		i -= len(m.ProductLabel)
		copy(dAtA[i:], m.ProductLabel)
		i = encodeVarintModel(dAtA, i, uint64(len(m.ProductLabel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProductName) > 0 {
		i -= len(m.ProductName)
		copy(dAtA[i:], m.ProductName)
		i = encodeVarintModel(dAtA, i, uint64(len(m.ProductName)))
		i--
		dAtA[i] = 0x22
	}
	if m.DeviceTypeId != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.DeviceTypeId))
		i--
		dAtA[i] = 0x18
	}
	if m.Pid != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x10
	}
	if m.Vid != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Model) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovModel(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovModel(uint64(m.Pid))
	}
	if m.DeviceTypeId != 0 {
		n += 1 + sovModel(uint64(m.DeviceTypeId))
	}
	l = len(m.ProductName)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.ProductLabel)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.PartNumber)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.CommissioningCustomFlow != 0 {
		n += 1 + sovModel(uint64(m.CommissioningCustomFlow))
	}
	l = len(m.CommissioningCustomFlowUrl)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.CommissioningModeInitialStepsHint != 0 {
		n += 1 + sovModel(uint64(m.CommissioningModeInitialStepsHint))
	}
	l = len(m.CommissioningModeInitialStepsInstruction)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.CommissioningModeSecondaryStepsHint != 0 {
		n += 1 + sovModel(uint64(m.CommissioningModeSecondaryStepsHint))
	}
	l = len(m.CommissioningModeSecondaryStepsInstruction)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.UserManualUrl)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.SupportUrl)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.ProductUrl)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 2 + l + sovModel(uint64(l))
	}
	return n
}

func sovModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Model) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
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
			return fmt.Errorf("proto: Model: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Model: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceTypeId", wireType)
			}
			m.DeviceTypeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceTypeId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningCustomFlow", wireType)
			}
			m.CommissioningCustomFlow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissioningCustomFlow |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningCustomFlowUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommissioningCustomFlowUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningModeInitialStepsHint", wireType)
			}
			m.CommissioningModeInitialStepsHint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissioningModeInitialStepsHint |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningModeInitialStepsInstruction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommissioningModeInitialStepsInstruction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningModeSecondaryStepsHint", wireType)
			}
			m.CommissioningModeSecondaryStepsHint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissioningModeSecondaryStepsHint |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissioningModeSecondaryStepsInstruction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommissioningModeSecondaryStepsInstruction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserManualUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserManualUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
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
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModel
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
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
				return 0, ErrInvalidLengthModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModel = fmt.Errorf("proto: unexpected end of group")
)