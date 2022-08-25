// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/device-provisioner/admin/admin.proto

// Package admin defines the administrative gRPC interfaces.

package admin

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_onosproject_onos_api_go_onos_p4rt_v1 "github.com/onosproject/onos-api/go/onos/p4rt/v1"
	v1 "github.com/onosproject/onos-api/go/onos/p4rt/v1"
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

// GetPipelineRequest get pipeline request
type GetPipelineRequest struct {
	PipelineConfigID github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID `protobuf:"bytes,1,opt,name=pipelineconfig_id,json=pipelineconfigId,proto3,casttype=github.com/onosproject/onos-api/go/onos/p4rt/v1.PipelineConfigID" json:"pipelineconfig_id,omitempty"`
}

func (m *GetPipelineRequest) Reset()         { *m = GetPipelineRequest{} }
func (m *GetPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRequest) ProtoMessage()    {}
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{0}
}
func (m *GetPipelineRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPipelineRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineRequest.Merge(m, src)
}
func (m *GetPipelineRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineRequest proto.InternalMessageInfo

func (m *GetPipelineRequest) GetPipelineConfigID() github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID {
	if m != nil {
		return m.PipelineConfigID
	}
	return ""
}

// GetPipelineResponse get pipeline response
type GetPipelineResponse struct {
	Pipelineconfig *v1.PipelineConfig `protobuf:"bytes,1,opt,name=pipelineconfig,proto3" json:"pipelineconfig,omitempty"`
}

func (m *GetPipelineResponse) Reset()         { *m = GetPipelineResponse{} }
func (m *GetPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineResponse) ProtoMessage()    {}
func (*GetPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{1}
}
func (m *GetPipelineResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPipelineResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineResponse.Merge(m, src)
}
func (m *GetPipelineResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineResponse proto.InternalMessageInfo

func (m *GetPipelineResponse) GetPipelineconfig() *v1.PipelineConfig {
	if m != nil {
		return m.Pipelineconfig
	}
	return nil
}

// ListPipelineRequest
type ListPipelinesRequest struct {
}

func (m *ListPipelinesRequest) Reset()         { *m = ListPipelinesRequest{} }
func (m *ListPipelinesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesRequest) ProtoMessage()    {}
func (*ListPipelinesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{2}
}
func (m *ListPipelinesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPipelinesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPipelinesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPipelinesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesRequest.Merge(m, src)
}
func (m *ListPipelinesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListPipelinesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesRequest proto.InternalMessageInfo

// ListPipelineResponse
type ListPipelinesResponse struct {
	Pipelineconfig *v1.PipelineConfig `protobuf:"bytes,1,opt,name=pipelineconfig,proto3" json:"pipelineconfig,omitempty"`
}

func (m *ListPipelinesResponse) Reset()         { *m = ListPipelinesResponse{} }
func (m *ListPipelinesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesResponse) ProtoMessage()    {}
func (*ListPipelinesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{3}
}
func (m *ListPipelinesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPipelinesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPipelinesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPipelinesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesResponse.Merge(m, src)
}
func (m *ListPipelinesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListPipelinesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesResponse proto.InternalMessageInfo

func (m *ListPipelinesResponse) GetPipelineconfig() *v1.PipelineConfig {
	if m != nil {
		return m.Pipelineconfig
	}
	return nil
}

// WatchPipelineRequest
type WatchPipelinesRequest struct {
	PipelineConfigID github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID `protobuf:"bytes,1,opt,name=pipelineconfig_id,json=pipelineconfigId,proto3,casttype=github.com/onosproject/onos-api/go/onos/p4rt/v1.PipelineConfigID" json:"pipelineconfig_id,omitempty"`
	Noreplay         bool                                                             `protobuf:"varint,2,opt,name=noreplay,proto3" json:"noreplay,omitempty"`
}

func (m *WatchPipelinesRequest) Reset()         { *m = WatchPipelinesRequest{} }
func (m *WatchPipelinesRequest) String() string { return proto.CompactTextString(m) }
func (*WatchPipelinesRequest) ProtoMessage()    {}
func (*WatchPipelinesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{4}
}
func (m *WatchPipelinesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WatchPipelinesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchPipelinesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WatchPipelinesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchPipelinesRequest.Merge(m, src)
}
func (m *WatchPipelinesRequest) XXX_Size() int {
	return m.Size()
}
func (m *WatchPipelinesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchPipelinesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchPipelinesRequest proto.InternalMessageInfo

func (m *WatchPipelinesRequest) GetPipelineConfigID() github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID {
	if m != nil {
		return m.PipelineConfigID
	}
	return ""
}

func (m *WatchPipelinesRequest) GetNoreplay() bool {
	if m != nil {
		return m.Noreplay
	}
	return false
}

// WatchPipelineResponse
type WatchPipelinesResponse struct {
	v1.PipelineConfig `protobuf:"bytes,1,opt,name=pipelineconfig,proto3,embedded=pipelineconfig" json:"pipelineconfig"`
}

func (m *WatchPipelinesResponse) Reset()         { *m = WatchPipelinesResponse{} }
func (m *WatchPipelinesResponse) String() string { return proto.CompactTextString(m) }
func (*WatchPipelinesResponse) ProtoMessage()    {}
func (*WatchPipelinesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02bb4de82d3ef5e, []int{5}
}
func (m *WatchPipelinesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WatchPipelinesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchPipelinesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WatchPipelinesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchPipelinesResponse.Merge(m, src)
}
func (m *WatchPipelinesResponse) XXX_Size() int {
	return m.Size()
}
func (m *WatchPipelinesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchPipelinesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchPipelinesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetPipelineRequest)(nil), "onos.deviceprovisioner.admin.GetPipelineRequest")
	proto.RegisterType((*GetPipelineResponse)(nil), "onos.deviceprovisioner.admin.GetPipelineResponse")
	proto.RegisterType((*ListPipelinesRequest)(nil), "onos.deviceprovisioner.admin.ListPipelinesRequest")
	proto.RegisterType((*ListPipelinesResponse)(nil), "onos.deviceprovisioner.admin.ListPipelinesResponse")
	proto.RegisterType((*WatchPipelinesRequest)(nil), "onos.deviceprovisioner.admin.WatchPipelinesRequest")
	proto.RegisterType((*WatchPipelinesResponse)(nil), "onos.deviceprovisioner.admin.WatchPipelinesResponse")
}

func init() {
	proto.RegisterFile("onos/device-provisioner/admin/admin.proto", fileDescriptor_a02bb4de82d3ef5e)
}

var fileDescriptor_a02bb4de82d3ef5e = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0xcf, 0xcb, 0x2f,
	0xd6, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0xd5, 0x2d, 0x28, 0xca, 0x2f, 0xcb, 0x2c, 0xce, 0xcc,
	0xcf, 0x4b, 0x2d, 0xd2, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0x83, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x32, 0x20, 0xa5, 0x7a, 0x10, 0xa5, 0x48, 0x2a, 0xf5, 0xc0, 0x6a, 0xa4, 0x94, 0xc0,
	0x06, 0x15, 0x98, 0x14, 0x95, 0xe8, 0x97, 0x19, 0xea, 0x17, 0x64, 0x16, 0xa4, 0xe6, 0x64, 0xe6,
	0xa5, 0xc6, 0x27, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0x43, 0x4c, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0xd2, 0x1c, 0x46, 0x2e, 0x21, 0xf7, 0xd4, 0x92, 0x00,
	0xa8, 0x96, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x56, 0x46, 0x2e, 0x41, 0x98, 0x31,
	0x10, 0x53, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x22, 0x1e, 0xdd, 0x93,
	0x17, 0x80, 0x69, 0x70, 0x06, 0x4b, 0x7a, 0xba, 0xfc, 0xba, 0x27, 0xef, 0x90, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x72, 0x4f, 0x41, 0x51, 0x7e, 0x56, 0x6a, 0x72,
	0x09, 0x98, 0xad, 0x9b, 0x58, 0x90, 0xa9, 0x9f, 0x9e, 0xaf, 0x8f, 0xec, 0x4e, 0x3d, 0x74, 0x33,
	0x82, 0x04, 0x50, 0xad, 0xf4, 0x4c, 0x51, 0x8a, 0xe6, 0x12, 0x46, 0x71, 0x5d, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x0b, 0x17, 0x1f, 0xaa, 0x52, 0xb0, 0xd3, 0xb8, 0x8d, 0x64, 0xf4, 0xc0,
	0xc1, 0x04, 0xb2, 0x40, 0x0f, 0xc3, 0x82, 0x20, 0x34, 0x3d, 0x4a, 0x62, 0x5c, 0x22, 0x3e, 0x99,
	0xc5, 0x70, 0xd3, 0x8b, 0xa1, 0x9e, 0x57, 0x8a, 0xe5, 0x12, 0x45, 0x13, 0xa7, 0xaa, 0xb5, 0xbb,
	0x19, 0xb9, 0x44, 0xc3, 0x13, 0x4b, 0x92, 0x33, 0xd0, 0x2d, 0x1e, 0x2c, 0xa1, 0x2e, 0x24, 0xc5,
	0xc5, 0x91, 0x97, 0x5f, 0x94, 0x5a, 0x90, 0x93, 0x58, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11,
	0x04, 0xe7, 0x2b, 0x65, 0x70, 0x89, 0xa1, 0x3b, 0x1e, 0x1a, 0x3a, 0x7e, 0xe4, 0x84, 0x8e, 0x13,
	0xc7, 0x89, 0x7b, 0xf2, 0x0c, 0x17, 0xee, 0xc9, 0x33, 0xa2, 0x87, 0x93, 0x51, 0x0f, 0x33, 0x97,
	0x28, 0xaa, 0xe2, 0xe0, 0xd4, 0x22, 0x50, 0xfa, 0x17, 0x2a, 0xe1, 0xe2, 0x46, 0x4a, 0x15, 0x42,
	0x06, 0x7a, 0xf8, 0x32, 0x87, 0x1e, 0x66, 0xf2, 0x96, 0x32, 0x24, 0x41, 0x07, 0xc4, 0x77, 0x4a,
	0x0c, 0x42, 0x35, 0x5c, 0xbc, 0x28, 0xc9, 0x42, 0xc8, 0x08, 0xbf, 0x29, 0xd8, 0xd2, 0x96, 0x94,
	0x31, 0x49, 0x7a, 0x60, 0x76, 0x1b, 0x30, 0x0a, 0xd5, 0x73, 0xf1, 0xa1, 0x86, 0xbb, 0x10, 0x01,
	0xa3, 0xb0, 0x26, 0x31, 0x29, 0x13, 0xd2, 0x34, 0x21, 0x1c, 0xe0, 0x24, 0x71, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0xe0, 0xa2, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x59, 0x99, 0x09, 0xcf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelineConfigServiceClient is the client API for PipelineConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineConfigServiceClient interface {
	// Get pipeline config based on a given ID
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error)
	// List returns all target pipelines
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (PipelineConfigService_ListPipelinesClient, error)
	// Watch returns a stream of pipeline change notifications
	WatchPipelines(ctx context.Context, in *WatchPipelinesRequest, opts ...grpc.CallOption) (PipelineConfigService_WatchPipelinesClient, error)
}

type pipelineConfigServiceClient struct {
	cc *grpc.ClientConn
}

func NewPipelineConfigServiceClient(cc *grpc.ClientConn) PipelineConfigServiceClient {
	return &pipelineConfigServiceClient{cc}
}

func (c *pipelineConfigServiceClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error) {
	out := new(GetPipelineResponse)
	err := c.cc.Invoke(ctx, "/onos.deviceprovisioner.admin.PipelineConfigService/GetPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineConfigServiceClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (PipelineConfigService_ListPipelinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PipelineConfigService_serviceDesc.Streams[0], "/onos.deviceprovisioner.admin.PipelineConfigService/ListPipelines", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineConfigServiceListPipelinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PipelineConfigService_ListPipelinesClient interface {
	Recv() (*ListPipelinesResponse, error)
	grpc.ClientStream
}

type pipelineConfigServiceListPipelinesClient struct {
	grpc.ClientStream
}

func (x *pipelineConfigServiceListPipelinesClient) Recv() (*ListPipelinesResponse, error) {
	m := new(ListPipelinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipelineConfigServiceClient) WatchPipelines(ctx context.Context, in *WatchPipelinesRequest, opts ...grpc.CallOption) (PipelineConfigService_WatchPipelinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PipelineConfigService_serviceDesc.Streams[1], "/onos.deviceprovisioner.admin.PipelineConfigService/WatchPipelines", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineConfigServiceWatchPipelinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PipelineConfigService_WatchPipelinesClient interface {
	Recv() (*WatchPipelinesResponse, error)
	grpc.ClientStream
}

type pipelineConfigServiceWatchPipelinesClient struct {
	grpc.ClientStream
}

func (x *pipelineConfigServiceWatchPipelinesClient) Recv() (*WatchPipelinesResponse, error) {
	m := new(WatchPipelinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineConfigServiceServer is the server API for PipelineConfigService service.
type PipelineConfigServiceServer interface {
	// Get pipeline config based on a given ID
	GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error)
	// List returns all target pipelines
	ListPipelines(*ListPipelinesRequest, PipelineConfigService_ListPipelinesServer) error
	// Watch returns a stream of pipeline change notifications
	WatchPipelines(*WatchPipelinesRequest, PipelineConfigService_WatchPipelinesServer) error
}

// UnimplementedPipelineConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPipelineConfigServiceServer struct {
}

func (*UnimplementedPipelineConfigServiceServer) GetPipeline(ctx context.Context, req *GetPipelineRequest) (*GetPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipeline not implemented")
}
func (*UnimplementedPipelineConfigServiceServer) ListPipelines(req *ListPipelinesRequest, srv PipelineConfigService_ListPipelinesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}
func (*UnimplementedPipelineConfigServiceServer) WatchPipelines(req *WatchPipelinesRequest, srv PipelineConfigService_WatchPipelinesServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPipelines not implemented")
}

func RegisterPipelineConfigServiceServer(s *grpc.Server, srv PipelineConfigServiceServer) {
	s.RegisterService(&_PipelineConfigService_serviceDesc, srv)
}

func _PipelineConfigService_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineConfigServiceServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.deviceprovisioner.admin.PipelineConfigService/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineConfigServiceServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineConfigService_ListPipelines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPipelinesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PipelineConfigServiceServer).ListPipelines(m, &pipelineConfigServiceListPipelinesServer{stream})
}

type PipelineConfigService_ListPipelinesServer interface {
	Send(*ListPipelinesResponse) error
	grpc.ServerStream
}

type pipelineConfigServiceListPipelinesServer struct {
	grpc.ServerStream
}

func (x *pipelineConfigServiceListPipelinesServer) Send(m *ListPipelinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PipelineConfigService_WatchPipelines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPipelinesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PipelineConfigServiceServer).WatchPipelines(m, &pipelineConfigServiceWatchPipelinesServer{stream})
}

type PipelineConfigService_WatchPipelinesServer interface {
	Send(*WatchPipelinesResponse) error
	grpc.ServerStream
}

type pipelineConfigServiceWatchPipelinesServer struct {
	grpc.ServerStream
}

func (x *pipelineConfigServiceWatchPipelinesServer) Send(m *WatchPipelinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PipelineConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.deviceprovisioner.admin.PipelineConfigService",
	HandlerType: (*PipelineConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPipeline",
			Handler:    _PipelineConfigService_GetPipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPipelines",
			Handler:       _PipelineConfigService_ListPipelines_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchPipelines",
			Handler:       _PipelineConfigService_WatchPipelines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "onos/device-provisioner/admin/admin.proto",
}

func (m *GetPipelineRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPipelineRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPipelineRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PipelineConfigID) > 0 {
		i -= len(m.PipelineConfigID)
		copy(dAtA[i:], m.PipelineConfigID)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.PipelineConfigID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetPipelineResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPipelineResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPipelineResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pipelineconfig != nil {
		{
			size, err := m.Pipelineconfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListPipelinesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPipelinesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPipelinesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListPipelinesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPipelinesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPipelinesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pipelineconfig != nil {
		{
			size, err := m.Pipelineconfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdmin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchPipelinesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchPipelinesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchPipelinesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Noreplay {
		i--
		if m.Noreplay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.PipelineConfigID) > 0 {
		i -= len(m.PipelineConfigID)
		copy(dAtA[i:], m.PipelineConfigID)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.PipelineConfigID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchPipelinesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchPipelinesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchPipelinesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PipelineConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAdmin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetPipelineRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PipelineConfigID)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *GetPipelineResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pipelineconfig != nil {
		l = m.Pipelineconfig.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *ListPipelinesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListPipelinesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pipelineconfig != nil {
		l = m.Pipelineconfig.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *WatchPipelinesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PipelineConfigID)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.Noreplay {
		n += 2
	}
	return n
}

func (m *WatchPipelinesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PipelineConfig.Size()
	n += 1 + l + sovAdmin(uint64(l))
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetPipelineRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: GetPipelineRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPipelineRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipelineConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PipelineConfigID = github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *GetPipelineResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: GetPipelineResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPipelineResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelineconfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pipelineconfig == nil {
				m.Pipelineconfig = &v1.PipelineConfig{}
			}
			if err := m.Pipelineconfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *ListPipelinesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: ListPipelinesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPipelinesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *ListPipelinesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: ListPipelinesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPipelinesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelineconfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pipelineconfig == nil {
				m.Pipelineconfig = &v1.PipelineConfig{}
			}
			if err := m.Pipelineconfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *WatchPipelinesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: WatchPipelinesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchPipelinesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipelineConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PipelineConfigID = github_com_onosproject_onos_api_go_onos_p4rt_v1.PipelineConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Noreplay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Noreplay = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *WatchPipelinesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: WatchPipelinesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchPipelinesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipelineConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PipelineConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)
