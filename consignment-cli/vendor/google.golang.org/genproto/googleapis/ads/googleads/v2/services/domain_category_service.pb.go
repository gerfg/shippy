// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/domain_category_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [DomainCategoryService.GetDomainCategory][google.ads.googleads.v2.services.DomainCategoryService.GetDomainCategory].
type GetDomainCategoryRequest struct {
	// Resource name of the domain category to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDomainCategoryRequest) Reset()         { *m = GetDomainCategoryRequest{} }
func (m *GetDomainCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetDomainCategoryRequest) ProtoMessage()    {}
func (*GetDomainCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb15af8e20123b1c, []int{0}
}

func (m *GetDomainCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainCategoryRequest.Unmarshal(m, b)
}
func (m *GetDomainCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GetDomainCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainCategoryRequest.Merge(m, src)
}
func (m *GetDomainCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetDomainCategoryRequest.Size(m)
}
func (m *GetDomainCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainCategoryRequest proto.InternalMessageInfo

func (m *GetDomainCategoryRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDomainCategoryRequest)(nil), "google.ads.googleads.v2.services.GetDomainCategoryRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/domain_category_service.proto", fileDescriptor_cb15af8e20123b1c)
}

var fileDescriptor_cb15af8e20123b1c = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0x70, 0xe1, 0x86, 0xeb, 0xc2, 0x80, 0x58, 0xa2, 0x8b, 0x52, 0xbb, 0x90, 0x2e,
	0x66, 0x30, 0x0a, 0xc2, 0x88, 0x4a, 0xaa, 0x50, 0x57, 0x52, 0x2a, 0x74, 0x21, 0x81, 0x32, 0x26,
	0x43, 0x08, 0x34, 0x33, 0x35, 0x67, 0x5a, 0x10, 0x71, 0xa1, 0xaf, 0xe0, 0x1b, 0xb8, 0xf4, 0x21,
	0x7c, 0x80, 0x6e, 0x7d, 0x05, 0x57, 0xee, 0xdd, 0x4b, 0x3a, 0x33, 0xa9, 0xc5, 0x86, 0xee, 0x3e,
	0xe6, 0x7c, 0x3f, 0xe7, 0x7c, 0x89, 0x73, 0x92, 0x08, 0x91, 0x0c, 0x19, 0xa6, 0x31, 0x60, 0x05,
	0x0b, 0x34, 0xf1, 0x31, 0xb0, 0x7c, 0x92, 0x46, 0x0c, 0x70, 0x2c, 0x32, 0x9a, 0xf2, 0x41, 0x44,
	0x25, 0x4b, 0x44, 0x7e, 0x37, 0xd0, 0x03, 0x34, 0xca, 0x85, 0x14, 0x6e, 0x5d, 0x89, 0x10, 0x8d,
	0x01, 0x95, 0x7a, 0x34, 0xf1, 0x91, 0xd1, 0x7b, 0x87, 0x55, 0x09, 0x39, 0x03, 0x31, 0xce, 0x97,
	0x44, 0x28, 0x6b, 0x6f, 0xdb, 0x08, 0x47, 0x29, 0xa6, 0x9c, 0x0b, 0x49, 0x65, 0x2a, 0x38, 0xe8,
	0xe9, 0xe6, 0x8f, 0x69, 0x34, 0x4c, 0x19, 0x97, 0x6a, 0xd0, 0x38, 0x75, 0x6a, 0x1d, 0x26, 0xcf,
	0x67, 0x96, 0x67, 0xda, 0xb1, 0xc7, 0x6e, 0xc7, 0x0c, 0xa4, 0xbb, 0xe3, 0xac, 0x99, 0xd4, 0x01,
	0xa7, 0x19, 0xab, 0x59, 0x75, 0x6b, 0xf7, 0x5f, 0xef, 0xbf, 0x79, 0xbc, 0xa4, 0x19, 0xf3, 0xbf,
	0x2c, 0x67, 0x63, 0x51, 0x7e, 0xa5, 0x6e, 0x71, 0xdf, 0x2c, 0x67, 0xfd, 0x97, 0xb7, 0x4b, 0xd0,
	0xaa, 0x0e, 0x50, 0xd5, 0x42, 0xde, 0x5e, 0xa5, 0xb6, 0x6c, 0x07, 0x2d, 0x2a, 0x1b, 0xe4, 0xe9,
	0xfd, 0xe3, 0xd9, 0x3e, 0x70, 0xfd, 0xa2, 0xc3, 0xfb, 0x85, 0x73, 0x8e, 0xa3, 0x31, 0x48, 0x91,
	0xb1, 0x1c, 0x70, 0x4b, 0x97, 0xaa, 0x65, 0x29, 0x03, 0xdc, 0x7a, 0xf0, 0xb6, 0xa6, 0x41, 0x6d,
	0x9e, 0xa2, 0xd1, 0x28, 0x05, 0x14, 0x89, 0xac, 0xfd, 0x68, 0x3b, 0xcd, 0x48, 0x64, 0x2b, 0xaf,
	0x69, 0x7b, 0x4b, 0xdb, 0xe9, 0x16, 0xed, 0x77, 0xad, 0xeb, 0x0b, 0xad, 0x4f, 0xc4, 0x90, 0xf2,
	0x04, 0x89, 0x3c, 0xc1, 0x09, 0xe3, 0xb3, 0x6f, 0x83, 0xe7, 0x89, 0xd5, 0x3f, 0xdc, 0x91, 0x01,
	0x2f, 0xf6, 0x9f, 0x4e, 0x10, 0xbc, 0xda, 0xf5, 0x8e, 0x32, 0x0c, 0x62, 0x40, 0x0a, 0x16, 0xa8,
	0xef, 0x23, 0x1d, 0x0c, 0x53, 0x43, 0x09, 0x83, 0x18, 0xc2, 0x92, 0x12, 0xf6, 0xfd, 0xd0, 0x50,
	0x3e, 0xed, 0xa6, 0x7a, 0x27, 0x24, 0x88, 0x81, 0x90, 0x92, 0x44, 0x48, 0xdf, 0x27, 0xc4, 0xd0,
	0x6e, 0xfe, 0xce, 0xf6, 0xdc, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x66, 0x20, 0x47, 0x17,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DomainCategoryServiceClient is the client API for DomainCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DomainCategoryServiceClient interface {
	// Returns the requested domain category.
	GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error)
}

type domainCategoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDomainCategoryServiceClient(cc *grpc.ClientConn) DomainCategoryServiceClient {
	return &domainCategoryServiceClient{cc}
}

func (c *domainCategoryServiceClient) GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error) {
	out := new(resources.DomainCategory)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.DomainCategoryService/GetDomainCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainCategoryServiceServer is the server API for DomainCategoryService service.
type DomainCategoryServiceServer interface {
	// Returns the requested domain category.
	GetDomainCategory(context.Context, *GetDomainCategoryRequest) (*resources.DomainCategory, error)
}

func RegisterDomainCategoryServiceServer(s *grpc.Server, srv DomainCategoryServiceServer) {
	s.RegisterService(&_DomainCategoryService_serviceDesc, srv)
}

func _DomainCategoryService_GetDomainCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.DomainCategoryService/GetDomainCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, req.(*GetDomainCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DomainCategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.DomainCategoryService",
	HandlerType: (*DomainCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomainCategory",
			Handler:    _DomainCategoryService_GetDomainCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/domain_category_service.proto",
}
