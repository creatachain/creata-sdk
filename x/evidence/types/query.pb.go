// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: creata/evidence/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_creatachain_augusteum_libs_bytes "github.com/creatachain/augusteum/libs/bytes"
	types "github.com/creatachain/creata-sdk/codec/types"
	query "github.com/creatachain/creata-sdk/types/query"
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

// QueryEvidenceRequest is the request type for the Query/Evidence RPC method.
type QueryEvidenceRequest struct {
	// evidence_hash defines the hash of the requested evidence.
	EvidenceHash github_com_creatachain_augusteum_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=evidence_hash,json=evidenceHash,proto3,casttype=github.com/creatachain/augusteum/libs/bytes.HexBytes" json:"evidence_hash,omitempty"`
}

func (m *QueryEvidenceRequest) Reset()         { *m = QueryEvidenceRequest{} }
func (m *QueryEvidenceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEvidenceRequest) ProtoMessage()    {}
func (*QueryEvidenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf9cbcf68147add, []int{0}
}
func (m *QueryEvidenceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEvidenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEvidenceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEvidenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEvidenceRequest.Merge(m, src)
}
func (m *QueryEvidenceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEvidenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEvidenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEvidenceRequest proto.InternalMessageInfo

func (m *QueryEvidenceRequest) GetEvidenceHash() github_com_creatachain_augusteum_libs_bytes.HexBytes {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

// QueryEvidenceResponse is the response type for the Query/Evidence RPC method.
type QueryEvidenceResponse struct {
	// evidence returns the requested evidence.
	Evidence *types.Any `protobuf:"bytes,1,opt,name=evidence,proto3" json:"evidence,omitempty"`
}

func (m *QueryEvidenceResponse) Reset()         { *m = QueryEvidenceResponse{} }
func (m *QueryEvidenceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEvidenceResponse) ProtoMessage()    {}
func (*QueryEvidenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf9cbcf68147add, []int{1}
}
func (m *QueryEvidenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEvidenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEvidenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEvidenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEvidenceResponse.Merge(m, src)
}
func (m *QueryEvidenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEvidenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEvidenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEvidenceResponse proto.InternalMessageInfo

func (m *QueryEvidenceResponse) GetEvidence() *types.Any {
	if m != nil {
		return m.Evidence
	}
	return nil
}

// QueryEvidenceRequest is the request type for the Query/AllEvidence RPC
// method.
type QueryAllEvidenceRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllEvidenceRequest) Reset()         { *m = QueryAllEvidenceRequest{} }
func (m *QueryAllEvidenceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllEvidenceRequest) ProtoMessage()    {}
func (*QueryAllEvidenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf9cbcf68147add, []int{2}
}
func (m *QueryAllEvidenceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllEvidenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllEvidenceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllEvidenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllEvidenceRequest.Merge(m, src)
}
func (m *QueryAllEvidenceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllEvidenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllEvidenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllEvidenceRequest proto.InternalMessageInfo

func (m *QueryAllEvidenceRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllEvidenceResponse is the response type for the Query/AllEvidence RPC
// method.
type QueryAllEvidenceResponse struct {
	// evidence returns all evidences.
	Evidence []*types.Any `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllEvidenceResponse) Reset()         { *m = QueryAllEvidenceResponse{} }
func (m *QueryAllEvidenceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllEvidenceResponse) ProtoMessage()    {}
func (*QueryAllEvidenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bf9cbcf68147add, []int{3}
}
func (m *QueryAllEvidenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllEvidenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllEvidenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllEvidenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllEvidenceResponse.Merge(m, src)
}
func (m *QueryAllEvidenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllEvidenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllEvidenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllEvidenceResponse proto.InternalMessageInfo

func (m *QueryAllEvidenceResponse) GetEvidence() []*types.Any {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *QueryAllEvidenceResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryEvidenceRequest)(nil), "creata.evidence.v1beta1.QueryEvidenceRequest")
	proto.RegisterType((*QueryEvidenceResponse)(nil), "creata.evidence.v1beta1.QueryEvidenceResponse")
	proto.RegisterType((*QueryAllEvidenceRequest)(nil), "creata.evidence.v1beta1.QueryAllEvidenceRequest")
	proto.RegisterType((*QueryAllEvidenceResponse)(nil), "creata.evidence.v1beta1.QueryAllEvidenceResponse")
}

func init() {
	proto.RegisterFile("creata/evidence/v1beta1/query.proto", fileDescriptor_7bf9cbcf68147add)
}

var fileDescriptor_7bf9cbcf68147add = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x37, 0x2b, 0x4a, 0x49, 0xeb, 0x25, 0xac, 0xb4, 0x0e, 0x32, 0xca, 0x08, 0xb5, 0x08,
	0x4d, 0xba, 0xab, 0x48, 0x4f, 0x42, 0x17, 0xd4, 0x0a, 0x1e, 0x74, 0x8e, 0x82, 0x48, 0x66, 0x1a,
	0x33, 0x83, 0xd3, 0x64, 0xba, 0x49, 0x4a, 0x07, 0xf1, 0xe2, 0x27, 0x10, 0xc4, 0x2f, 0xe0, 0xcd,
	0x6f, 0xe2, 0xb1, 0xe0, 0xa5, 0x27, 0x91, 0x5d, 0x3f, 0x85, 0x27, 0x99, 0x24, 0xb3, 0xdd, 0xed,
	0xb6, 0xee, 0xf6, 0xf6, 0xce, 0xe4, 0x7d, 0x9f, 0xe7, 0xf7, 0xfe, 0x81, 0x77, 0xd3, 0x01, 0xa3,
	0x9a, 0x12, 0x76, 0x98, 0xef, 0x31, 0x91, 0x32, 0x72, 0xd8, 0x4d, 0x98, 0xa6, 0x5d, 0x72, 0x60,
	0xd8, 0xa0, 0xc2, 0xe5, 0x40, 0x6a, 0x89, 0x56, 0x5d, 0x12, 0x6e, 0x92, 0xb0, 0x4f, 0x0a, 0xee,
	0xfb, 0xea, 0x84, 0x2a, 0xe6, 0x2a, 0xc6, 0xf5, 0x25, 0xe5, 0xb9, 0xa0, 0x3a, 0x97, 0xc2, 0x89,
	0x04, 0x1d, 0x2e, 0xb9, 0xb4, 0x21, 0xa9, 0x23, 0xff, 0xf7, 0x26, 0x97, 0x92, 0x17, 0x8c, 0xd8,
	0xaf, 0xc4, 0xbc, 0x23, 0x54, 0x78, 0xd7, 0xe0, 0x96, 0x7f, 0xa2, 0x65, 0x4e, 0xa8, 0x10, 0x52,
	0x5b, 0x35, 0xe5, 0x5e, 0x23, 0x03, 0x3b, 0xaf, 0x6a, 0xc3, 0x27, 0x9e, 0x29, 0x66, 0x07, 0x86,
	0x29, 0x8d, 0xde, 0xc0, 0xeb, 0x0d, 0xe6, 0xdb, 0x8c, 0xaa, 0x6c, 0x0d, 0xdc, 0x01, 0x1b, 0x2b,
	0xfd, 0xed, 0xbf, 0xbf, 0x6e, 0x3f, 0xe4, 0xb9, 0xce, 0x4c, 0x82, 0x53, 0xb9, 0x4f, 0x1c, 0x78,
	0x9a, 0xd1, 0x5c, 0x10, 0x6a, 0xb8, 0x51, 0x9a, 0x99, 0x7d, 0x52, 0xe4, 0x89, 0x22, 0x49, 0xa5,
	0x99, 0xc2, 0xbb, 0xec, 0xa8, 0x5f, 0x07, 0xf1, 0x4a, 0x23, 0xb7, 0x4b, 0x55, 0x16, 0x3d, 0x87,
	0x37, 0xce, 0xd8, 0xaa, 0x52, 0x0a, 0xc5, 0xd0, 0x16, 0x5c, 0x6a, 0x12, 0xad, 0xe5, 0x72, 0xaf,
	0x83, 0x5d, 0x03, 0xb8, 0xe9, 0x0d, 0xef, 0x88, 0x2a, 0x1e, 0x67, 0x45, 0x14, 0xae, 0x5a, 0xa9,
	0x9d, 0xa2, 0x38, 0xdb, 0xc4, 0x53, 0x08, 0x4f, 0xe7, 0xe7, 0xe5, 0xd6, 0xb1, 0xdf, 0x42, 0x3d,
	0x6c, 0xec, 0xd6, 0xe3, 0x87, 0x8d, 0x5f, 0x52, 0xde, 0xd4, 0xc6, 0x13, 0x95, 0xd1, 0x57, 0x00,
	0xd7, 0x66, 0x3d, 0xce, 0x25, 0xbe, 0x32, 0x9f, 0x18, 0x3d, 0x9b, 0xc2, 0x6a, 0x5b, 0xac, 0x7b,
	0x73, 0xb1, 0x9c, 0xdd, 0x24, 0x57, 0xef, 0xa4, 0x0d, 0xaf, 0x5a, 0x2e, 0xf4, 0x1d, 0xc0, 0xa5,
	0x86, 0x0c, 0x6d, 0xe2, 0x0b, 0x0e, 0x0d, 0x9f, 0xb7, 0xea, 0x00, 0x2f, 0x9a, 0xee, 0x08, 0xa2,
	0xc7, 0x9f, 0x7e, 0xfe, 0xf9, 0xd2, 0xde, 0x46, 0x8f, 0xa6, 0xb6, 0x3f, 0x73, 0xf9, 0xe3, 0x1f,
	0x1f, 0xa6, 0x0e, 0xe9, 0x23, 0xfa, 0x06, 0xe0, 0xf2, 0xc4, 0x20, 0xd1, 0xd6, 0xff, 0xfd, 0x67,
	0xf7, 0x1a, 0x74, 0x2f, 0x51, 0xe1, 0xa1, 0xb1, 0x85, 0xde, 0x40, 0xeb, 0x8b, 0x41, 0xf7, 0x5f,
	0xfc, 0x18, 0x86, 0xe0, 0x78, 0x18, 0x82, 0xdf, 0xc3, 0x10, 0x7c, 0x1e, 0x85, 0xad, 0xe3, 0x51,
	0xd8, 0x3a, 0x19, 0x85, 0xad, 0xd7, 0xbd, 0x0b, 0xce, 0xdf, 0xc5, 0x9b, 0x6a, 0xef, 0x3d, 0x39,
	0x3a, 0x15, 0xd6, 0x55, 0xc9, 0x54, 0x72, 0xcd, 0x5e, 0xc2, 0x83, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x83, 0x41, 0x9f, 0x27, 0x04, 0x00, 0x00,
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
	// Evidence queries evidence based on evidence hash.
	Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error) {
	out := new(QueryEvidenceResponse)
	err := c.cc.Invoke(ctx, "/creata.evidence.v1beta1.Query/Evidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error) {
	out := new(QueryAllEvidenceResponse)
	err := c.cc.Invoke(ctx, "/creata.evidence.v1beta1.Query/AllEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Evidence(ctx context.Context, req *QueryEvidenceRequest) (*QueryEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evidence not implemented")
}
func (*UnimplementedQueryServer) AllEvidence(ctx context.Context, req *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllEvidence not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Evidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Evidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creata.evidence.v1beta1.Query/Evidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Evidence(ctx, req.(*QueryEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllEvidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creata.evidence.v1beta1.Query/AllEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllEvidence(ctx, req.(*QueryAllEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "creata.evidence.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evidence",
			Handler:    _Query_Evidence_Handler,
		},
		{
			MethodName: "AllEvidence",
			Handler:    _Query_AllEvidence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creata/evidence/v1beta1/query.proto",
}

func (m *QueryEvidenceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEvidenceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEvidenceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EvidenceHash) > 0 {
		i -= len(m.EvidenceHash)
		copy(dAtA[i:], m.EvidenceHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EvidenceHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEvidenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEvidenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEvidenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllEvidenceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllEvidenceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllEvidenceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllEvidenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllEvidenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllEvidenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryEvidenceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EvidenceHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEvidenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllEvidenceRequest) Size() (n int) {
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

func (m *QueryAllEvidenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
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
func (m *QueryEvidenceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEvidenceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEvidenceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvidenceHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvidenceHash = append(m.EvidenceHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EvidenceHash == nil {
				m.EvidenceHash = []byte{}
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
func (m *QueryEvidenceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEvidenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEvidenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
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
			if m.Evidence == nil {
				m.Evidence = &types.Any{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllEvidenceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllEvidenceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllEvidenceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllEvidenceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllEvidenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllEvidenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
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
			m.Evidence = append(m.Evidence, &types.Any{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
