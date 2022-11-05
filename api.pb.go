// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IndexService service

type IndexServiceClient interface {
	AddIndex(ctx context.Context, in *AddIndexRequest, opts ...grpc.CallOption) (*AddIndexResponse, error)
	DescribeIndex(ctx context.Context, in *DescribeIndexRequest, opts ...grpc.CallOption) (*DescribeIndexResponse, error)
	//  rpc AlterIndex(AlterIndexRequest) returns (AlterIndexResponse);
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error)
}

type indexServiceClient struct {
	cc *grpc.ClientConn
}

func NewIndexServiceClient(cc *grpc.ClientConn) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) AddIndex(ctx context.Context, in *AddIndexRequest, opts ...grpc.CallOption) (*AddIndexResponse, error) {
	out := new(AddIndexResponse)
	err := grpc.Invoke(ctx, "/api.IndexService/AddIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DescribeIndex(ctx context.Context, in *DescribeIndexRequest, opts ...grpc.CallOption) (*DescribeIndexResponse, error) {
	out := new(DescribeIndexResponse)
	err := grpc.Invoke(ctx, "/api.IndexService/DescribeIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := grpc.Invoke(ctx, "/api.IndexService/DeleteIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error) {
	out := new(ListIndexesResponse)
	err := grpc.Invoke(ctx, "/api.IndexService/ListIndexes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IndexService service

type IndexServiceServer interface {
	AddIndex(context.Context, *AddIndexRequest) (*AddIndexResponse, error)
	DescribeIndex(context.Context, *DescribeIndexRequest) (*DescribeIndexResponse, error)
	//  rpc AlterIndex(AlterIndexRequest) returns (AlterIndexResponse);
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error)
}

func RegisterIndexServiceServer(s *grpc.Server, srv IndexServiceServer) {
	s.RegisterService(&_IndexService_serviceDesc, srv)
}

func _IndexService_AddIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).AddIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IndexService/AddIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).AddIndex(ctx, req.(*AddIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DescribeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DescribeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IndexService/DescribeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DescribeIndex(ctx, req.(*DescribeIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IndexService/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_ListIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).ListIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IndexService/ListIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).ListIndexes(ctx, req.(*ListIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIndex",
			Handler:    _IndexService_AddIndex_Handler,
		},
		{
			MethodName: "DescribeIndex",
			Handler:    _IndexService_DescribeIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _IndexService_DeleteIndex_Handler,
		},
		{
			MethodName: "ListIndexes",
			Handler:    _IndexService_ListIndexes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for SegmentService service

type SegmentServiceClient interface {
	PutSegment(ctx context.Context, in *PutSegmentRequest, opts ...grpc.CallOption) (*PutSegmentResponse, error)
	GetSegment(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetSegmentResponse, error)
	DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error)
	MatchSegment(ctx context.Context, in *MatchSegmentRequest, opts ...grpc.CallOption) (*MatchSegmentResponse, error)
	MatchSegmentKey(ctx context.Context, in *MatchSegmentKeyRequest, opts ...grpc.CallOption) (*MatchSegmentKeyResponse, error)
}

type segmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewSegmentServiceClient(cc *grpc.ClientConn) SegmentServiceClient {
	return &segmentServiceClient{cc}
}

func (c *segmentServiceClient) PutSegment(ctx context.Context, in *PutSegmentRequest, opts ...grpc.CallOption) (*PutSegmentResponse, error) {
	out := new(PutSegmentResponse)
	err := grpc.Invoke(ctx, "/api.SegmentService/PutSegment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServiceClient) GetSegment(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetSegmentResponse, error) {
	out := new(GetSegmentResponse)
	err := grpc.Invoke(ctx, "/api.SegmentService/GetSegment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServiceClient) DeleteSegment(ctx context.Context, in *DeleteSegmentRequest, opts ...grpc.CallOption) (*DeleteSegmentResponse, error) {
	out := new(DeleteSegmentResponse)
	err := grpc.Invoke(ctx, "/api.SegmentService/DeleteSegment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServiceClient) MatchSegment(ctx context.Context, in *MatchSegmentRequest, opts ...grpc.CallOption) (*MatchSegmentResponse, error) {
	out := new(MatchSegmentResponse)
	err := grpc.Invoke(ctx, "/api.SegmentService/MatchSegment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentServiceClient) MatchSegmentKey(ctx context.Context, in *MatchSegmentKeyRequest, opts ...grpc.CallOption) (*MatchSegmentKeyResponse, error) {
	out := new(MatchSegmentKeyResponse)
	err := grpc.Invoke(ctx, "/api.SegmentService/MatchSegmentKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SegmentService service

type SegmentServiceServer interface {
	PutSegment(context.Context, *PutSegmentRequest) (*PutSegmentResponse, error)
	GetSegment(context.Context, *GetSegmentRequest) (*GetSegmentResponse, error)
	DeleteSegment(context.Context, *DeleteSegmentRequest) (*DeleteSegmentResponse, error)
	MatchSegment(context.Context, *MatchSegmentRequest) (*MatchSegmentResponse, error)
	MatchSegmentKey(context.Context, *MatchSegmentKeyRequest) (*MatchSegmentKeyResponse, error)
}

func RegisterSegmentServiceServer(s *grpc.Server, srv SegmentServiceServer) {
	s.RegisterService(&_SegmentService_serviceDesc, srv)
}

func _SegmentService_PutSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServiceServer).PutSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SegmentService/PutSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServiceServer).PutSegment(ctx, req.(*PutSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentService_GetSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServiceServer).GetSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SegmentService/GetSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServiceServer).GetSegment(ctx, req.(*GetSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentService_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServiceServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SegmentService/DeleteSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServiceServer).DeleteSegment(ctx, req.(*DeleteSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentService_MatchSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServiceServer).MatchSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SegmentService/MatchSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServiceServer).MatchSegment(ctx, req.(*MatchSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentService_MatchSegmentKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchSegmentKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentServiceServer).MatchSegmentKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SegmentService/MatchSegmentKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentServiceServer).MatchSegmentKey(ctx, req.(*MatchSegmentKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SegmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SegmentService",
	HandlerType: (*SegmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutSegment",
			Handler:    _SegmentService_PutSegment_Handler,
		},
		{
			MethodName: "GetSegment",
			Handler:    _SegmentService_GetSegment_Handler,
		},
		{
			MethodName: "DeleteSegment",
			Handler:    _SegmentService_DeleteSegment_Handler,
		},
		{
			MethodName: "MatchSegment",
			Handler:    _SegmentService_MatchSegment_Handler,
		},
		{
			MethodName: "MatchSegmentKey",
			Handler:    _SegmentService_MatchSegmentKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x61, 0x82, 0xe8, 0xd9, 0xa6, 0x10, 0xd4, 0xd9, 0xe8, 0x33, 0xec, 0x62, 0x5e, 0x78,
	0x25, 0x28, 0x0e, 0x45, 0x9c, 0x20, 0xee, 0x09, 0xb6, 0xf6, 0xa0, 0x01, 0x6d, 0x6b, 0x93, 0x89,
	0x3e, 0x87, 0x4f, 0xe9, 0x5b, 0xd8, 0x26, 0xe7, 0xac, 0x49, 0x9b, 0xbb, 0xe6, 0xfb, 0xcf, 0xff,
	0xd1, 0x1c, 0x02, 0xfb, 0xab, 0x52, 0x4d, 0xcb, 0xaa, 0x30, 0x85, 0xd8, 0xa9, 0x3f, 0xe5, 0x50,
	0xe5, 0x19, 0x7e, 0x3b, 0x22, 0xc7, 0x1a, 0x5f, 0x3f, 0x30, 0x37, 0xee, 0x38, 0xfb, 0x1d, 0xc0,
	0xe8, 0xa1, 0x89, 0x97, 0x58, 0x7d, 0xa9, 0x14, 0xc5, 0x25, 0xec, 0xdd, 0x64, 0x99, 0x45, 0xe2,
	0x68, 0xda, 0x98, 0xf8, 0xf8, 0x82, 0x9f, 0x1b, 0xd4, 0x46, 0x1e, 0x77, 0xa8, 0x2e, 0x8b, 0x5c,
	0xa3, 0xb8, 0x83, 0xf1, 0x1c, 0x75, 0x5a, 0xa9, 0x35, 0xba, 0x76, 0x62, 0xe7, 0x02, 0xc6, 0x0a,
	0x19, 0x8b, 0xc8, 0x73, 0x0d, 0xc3, 0x39, 0xbe, 0xa3, 0x21, 0xcb, 0x84, 0x46, 0xb7, 0x84, 0x1d,
	0xa7, 0xfd, 0xa0, 0x35, 0x2c, 0x94, 0x36, 0x16, 0xa2, 0x26, 0x83, 0x47, 0x42, 0x43, 0x10, 0x38,
	0xc3, 0xec, 0x6f, 0x00, 0x07, 0x4b, 0xb7, 0x27, 0xde, 0xcb, 0x15, 0xc0, 0xf3, 0xc6, 0x10, 0x14,
	0x27, 0xb6, 0xda, 0x02, 0x56, 0x4e, 0x7a, 0x9c, 0xfe, 0xa9, 0xae, 0xdf, 0x63, 0xa7, 0xde, 0x82,
	0xb0, 0xee, 0x73, 0x7f, 0xb9, 0xcd, 0x4d, 0xd9, 0x90, 0x78, 0xb7, 0xef, 0x48, 0x64, 0x2c, 0x22,
	0xcf, 0x2d, 0x8c, 0x9e, 0x56, 0x26, 0x7d, 0x63, 0x8d, 0x5b, 0x81, 0x8f, 0xd8, 0x92, 0x44, 0x12,
	0x92, 0x2c, 0xe0, 0xd0, 0xe7, 0x8f, 0xf8, 0x23, 0xce, 0x7a, 0xd3, 0x35, 0x65, 0xd5, 0x79, 0x3c,
	0x74, 0xb6, 0xf5, 0xae, 0x7d, 0x88, 0x17, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x67, 0xe1,
	0xf2, 0xb6, 0x02, 0x00, 0x00,
}
