// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryGetVendorInfoTypeRequest struct {
	VendorID uint64 `protobuf:"varint,1,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
}

func (m *QueryGetVendorInfoTypeRequest) Reset()         { *m = QueryGetVendorInfoTypeRequest{} }
func (m *QueryGetVendorInfoTypeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetVendorInfoTypeRequest) ProtoMessage()    {}
func (*QueryGetVendorInfoTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{0}
}
func (m *QueryGetVendorInfoTypeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVendorInfoTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVendorInfoTypeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVendorInfoTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVendorInfoTypeRequest.Merge(m, src)
}
func (m *QueryGetVendorInfoTypeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVendorInfoTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVendorInfoTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVendorInfoTypeRequest proto.InternalMessageInfo

func (m *QueryGetVendorInfoTypeRequest) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

type QueryGetVendorInfoTypeResponse struct {
	VendorInfoType VendorInfoType `protobuf:"bytes,1,opt,name=vendorInfoType,proto3" json:"vendorInfoType"`
}

func (m *QueryGetVendorInfoTypeResponse) Reset()         { *m = QueryGetVendorInfoTypeResponse{} }
func (m *QueryGetVendorInfoTypeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetVendorInfoTypeResponse) ProtoMessage()    {}
func (*QueryGetVendorInfoTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{1}
}
func (m *QueryGetVendorInfoTypeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVendorInfoTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVendorInfoTypeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVendorInfoTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVendorInfoTypeResponse.Merge(m, src)
}
func (m *QueryGetVendorInfoTypeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVendorInfoTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVendorInfoTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVendorInfoTypeResponse proto.InternalMessageInfo

func (m *QueryGetVendorInfoTypeResponse) GetVendorInfoType() VendorInfoType {
	if m != nil {
		return m.VendorInfoType
	}
	return VendorInfoType{}
}

type QueryAllVendorInfoTypeRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVendorInfoTypeRequest) Reset()         { *m = QueryAllVendorInfoTypeRequest{} }
func (m *QueryAllVendorInfoTypeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllVendorInfoTypeRequest) ProtoMessage()    {}
func (*QueryAllVendorInfoTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{2}
}
func (m *QueryAllVendorInfoTypeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVendorInfoTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVendorInfoTypeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVendorInfoTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVendorInfoTypeRequest.Merge(m, src)
}
func (m *QueryAllVendorInfoTypeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVendorInfoTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVendorInfoTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVendorInfoTypeRequest proto.InternalMessageInfo

func (m *QueryAllVendorInfoTypeRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllVendorInfoTypeResponse struct {
	VendorInfoType []VendorInfoType    `protobuf:"bytes,1,rep,name=vendorInfoType,proto3" json:"vendorInfoType"`
	Pagination     *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVendorInfoTypeResponse) Reset()         { *m = QueryAllVendorInfoTypeResponse{} }
func (m *QueryAllVendorInfoTypeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllVendorInfoTypeResponse) ProtoMessage()    {}
func (*QueryAllVendorInfoTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4af47a04389bb67, []int{3}
}
func (m *QueryAllVendorInfoTypeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVendorInfoTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVendorInfoTypeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVendorInfoTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVendorInfoTypeResponse.Merge(m, src)
}
func (m *QueryAllVendorInfoTypeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVendorInfoTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVendorInfoTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVendorInfoTypeResponse proto.InternalMessageInfo

func (m *QueryAllVendorInfoTypeResponse) GetVendorInfoType() []VendorInfoType {
	if m != nil {
		return m.VendorInfoType
	}
	return nil
}

func (m *QueryAllVendorInfoTypeResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetVendorInfoTypeRequest)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryGetVendorInfoTypeRequest")
	proto.RegisterType((*QueryGetVendorInfoTypeResponse)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryGetVendorInfoTypeResponse")
	proto.RegisterType((*QueryAllVendorInfoTypeRequest)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryAllVendorInfoTypeRequest")
	proto.RegisterType((*QueryAllVendorInfoTypeResponse)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.QueryAllVendorInfoTypeResponse")
}

func init() { proto.RegisterFile("vendorinfo/query.proto", fileDescriptor_e4af47a04389bb67) }

var fileDescriptor_e4af47a04389bb67 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0x31, 0x8a, 0x8c, 0x50, 0x70, 0x10, 0x91, 0xa0, 0xab, 0xee, 0x41, 0x45, 0xc8,
	0x0c, 0xad, 0x78, 0xf2, 0xd4, 0xa2, 0xad, 0x45, 0x10, 0x1b, 0xab, 0xa0, 0x97, 0x32, 0x9b, 0xbc,
	0x8e, 0x03, 0x9b, 0x79, 0xdb, 0x9d, 0x49, 0x30, 0x8a, 0x17, 0x3f, 0x41, 0x41, 0x04, 0xbf, 0x88,
	0xdf, 0xa1, 0xc7, 0x82, 0x17, 0x4f, 0x45, 0x12, 0x3f, 0x88, 0xec, 0xcc, 0xae, 0xc9, 0x6a, 0x13,
	0xa1, 0x8d, 0xb7, 0x97, 0xc9, 0xbc, 0xff, 0x7f, 0xfe, 0xbf, 0x7d, 0x33, 0xe4, 0xf2, 0x00, 0x74,
	0x17, 0x33, 0xa5, 0x77, 0x91, 0xef, 0xf5, 0x21, 0x1b, 0xb2, 0x34, 0x43, 0x8b, 0xf4, 0xfe, 0x3b,
	0x25, 0x63, 0x00, 0x91, 0x24, 0x4a, 0xe8, 0x0e, 0xb0, 0xae, 0x32, 0x36, 0x53, 0x71, 0xdf, 0x42,
	0xb7, 0x83, 0xbd, 0xd4, 0xaf, 0x26, 0xd0, 0x95, 0x90, 0xb1, 0x89, 0x44, 0xf3, 0xaa, 0x44, 0x94,
	0x09, 0x70, 0x91, 0x2a, 0x2e, 0xb4, 0x46, 0x2b, 0xac, 0x42, 0x6d, 0xbc, 0x68, 0xf3, 0x6e, 0x07,
	0x4d, 0x0f, 0x0d, 0x8f, 0x85, 0x01, 0xef, 0xc6, 0x07, 0xcb, 0x31, 0x58, 0xb1, 0xcc, 0x53, 0x21,
	0x95, 0x76, 0x9b, 0x8b, 0xbd, 0x37, 0xa7, 0x0e, 0xe6, 0xcb, 0x9d, 0xbc, 0xde, 0xb1, 0xc3, 0x14,
	0x8a, 0x2d, 0x97, 0x24, 0x4a, 0x74, 0x25, 0xcf, 0x2b, 0xbf, 0x1a, 0x3d, 0x20, 0xd7, 0xb6, 0x72,
	0xe9, 0x0d, 0xb0, 0x2f, 0x5d, 0xdf, 0xa6, 0xde, 0xc5, 0xed, 0x61, 0x0a, 0x6d, 0xd8, 0xeb, 0x83,
	0xb1, 0xb4, 0x49, 0xce, 0x7b, 0xc1, 0xcd, 0x87, 0x57, 0x82, 0x1b, 0xc1, 0x9d, 0x46, 0xfb, 0xf7,
	0xef, 0xe8, 0x73, 0x40, 0xc2, 0x59, 0xdd, 0x26, 0x45, 0x6d, 0x80, 0x1a, 0xb2, 0x34, 0xa8, 0xfc,
	0xe3, 0x44, 0x2e, 0xac, 0x3c, 0x62, 0x27, 0x42, 0xc6, 0xaa, 0x36, 0x6b, 0x8d, 0x83, 0xa3, 0xeb,
	0xb5, 0xf6, 0x1f, 0x16, 0x91, 0x2c, 0x42, 0xad, 0x26, 0xc9, 0xf1, 0xa1, 0xd6, 0x09, 0x99, 0x20,
	0x2c, 0x4e, 0x74, 0x8b, 0x79, 0xde, 0x2c, 0xe7, 0xcd, 0xfc, 0xd7, 0x2d, 0x78, 0xb3, 0x67, 0x42,
	0x96, 0xbd, 0xed, 0xa9, 0xce, 0xe8, 0xa8, 0x04, 0x70, 0x8c, 0xd3, 0x1c, 0x00, 0x67, 0xfe, 0x33,
	0x00, 0xba, 0x51, 0xc9, 0x57, 0x77, 0xf9, 0x6e, 0xff, 0x33, 0x9f, 0x3f, 0xf1, 0x74, 0xc0, 0x95,
	0xaf, 0x0d, 0x72, 0xd6, 0x05, 0xa4, 0x5f, 0xea, 0x64, 0xa9, 0xea, 0x4d, 0xb7, 0x4f, 0x18, 0x61,
	0xee, 0xc0, 0x35, 0x5f, 0x2c, 0x58, 0xd5, 0xa7, 0x8a, 0x5e, 0x7d, 0xfc, 0xf6, 0xf3, 0x53, 0xfd,
	0x39, 0xdd, 0xe2, 0x5e, 0xbe, 0x55, 0xea, 0xf3, 0x39, 0xfa, 0xfc, 0xaf, 0x6b, 0x55, 0xea, 0xf2,
	0xf7, 0xe5, 0x2d, 0xf8, 0x40, 0xf7, 0xeb, 0xe4, 0x62, 0xd5, 0x75, 0x35, 0x49, 0x4e, 0x47, 0x67,
	0xd6, 0xe4, 0x9e, 0x8e, 0xce, 0xcc, 0x29, 0x8d, 0x9e, 0x3a, 0x3a, 0x8f, 0xe9, 0xfa, 0x62, 0xe8,
	0xac, 0xc1, 0xc1, 0x28, 0x0c, 0x0e, 0x47, 0x61, 0xf0, 0x63, 0x14, 0x06, 0xfb, 0xe3, 0xb0, 0x76,
	0x38, 0x0e, 0x6b, 0xdf, 0xc7, 0x61, 0xed, 0xf5, 0x13, 0xa9, 0xec, 0x9b, 0x7e, 0xcc, 0x3a, 0xd8,
	0x9b, 0xe7, 0xd5, 0x9a, 0x98, 0xb5, 0x0a, 0xb7, 0xb7, 0xd3, 0x7e, 0xf9, 0xc3, 0x66, 0xe2, 0x73,
	0xee, 0x11, 0xbb, 0xf7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x16, 0x71, 0xf6, 0x26, 0x98, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a vendorInfoType by index.
	VendorInfoType(ctx context.Context, in *QueryGetVendorInfoTypeRequest, opts ...grpc.CallOption) (*QueryGetVendorInfoTypeResponse, error)
	// Queries a list of vendorInfoType items.
	VendorInfoTypeAll(ctx context.Context, in *QueryAllVendorInfoTypeRequest, opts ...grpc.CallOption) (*QueryAllVendorInfoTypeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) VendorInfoType(ctx context.Context, in *QueryGetVendorInfoTypeRequest, opts ...grpc.CallOption) (*QueryGetVendorInfoTypeResponse, error) {
	out := new(QueryGetVendorInfoTypeResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VendorInfoTypeAll(ctx context.Context, in *QueryAllVendorInfoTypeRequest, opts ...grpc.CallOption) (*QueryAllVendorInfoTypeResponse, error) {
	out := new(QueryAllVendorInfoTypeResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoTypeAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a vendorInfoType by index.
	VendorInfoType(context.Context, *QueryGetVendorInfoTypeRequest) (*QueryGetVendorInfoTypeResponse, error)
	// Queries a list of vendorInfoType items.
	VendorInfoTypeAll(context.Context, *QueryAllVendorInfoTypeRequest) (*QueryAllVendorInfoTypeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) VendorInfoType(ctx context.Context, req *QueryGetVendorInfoTypeRequest) (*QueryGetVendorInfoTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorInfoType not implemented")
}
func (*UnimplementedQueryServer) VendorInfoTypeAll(ctx context.Context, req *QueryAllVendorInfoTypeRequest) (*QueryAllVendorInfoTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorInfoTypeAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_VendorInfoType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetVendorInfoTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VendorInfoType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VendorInfoType(ctx, req.(*QueryGetVendorInfoTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VendorInfoTypeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllVendorInfoTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VendorInfoTypeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.vendorinfo.Query/VendorInfoTypeAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VendorInfoTypeAll(ctx, req.(*QueryAllVendorInfoTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zigbeealliance.distributedcomplianceledger.vendorinfo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VendorInfoType",
			Handler:    _Query_VendorInfoType_Handler,
		},
		{
			MethodName: "VendorInfoTypeAll",
			Handler:    _Query_VendorInfoTypeAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendorinfo/query.proto",
}

func (m *QueryGetVendorInfoTypeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVendorInfoTypeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVendorInfoTypeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VendorID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetVendorInfoTypeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVendorInfoTypeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVendorInfoTypeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VendorInfoType.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllVendorInfoTypeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVendorInfoTypeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVendorInfoTypeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllVendorInfoTypeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVendorInfoTypeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVendorInfoTypeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.VendorInfoType) > 0 {
		for iNdEx := len(m.VendorInfoType) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VendorInfoType[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetVendorInfoTypeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovQuery(uint64(m.VendorID))
	}
	return n
}

func (m *QueryGetVendorInfoTypeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VendorInfoType.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllVendorInfoTypeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllVendorInfoTypeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VendorInfoType) > 0 {
		for _, e := range m.VendorInfoType {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetVendorInfoTypeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetVendorInfoTypeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVendorInfoTypeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetVendorInfoTypeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetVendorInfoTypeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVendorInfoTypeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorInfoType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VendorInfoType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllVendorInfoTypeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllVendorInfoTypeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVendorInfoTypeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllVendorInfoTypeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllVendorInfoTypeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVendorInfoTypeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorInfoType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorInfoType = append(m.VendorInfoType, VendorInfoType{})
			if err := m.VendorInfoType[len(m.VendorInfoType)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
