// Code generated by protoc-gen-go.
// source: segmentation.proto
// DO NOT EDIT!

/*
Package segmentation is a generated protocol buffer package.

It is generated from these files:
	segmentation.proto

It has these top-level messages:
	AdRequest
	SupplySegment
	BlacklistResponse
*/
package segmentation

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdRequest struct {
	DeviceId string  `protobuf:"bytes,1,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Lat      float64 `protobuf:"fixed64,2,opt,name=lat" json:"lat,omitempty"`
	Lng      float64 `protobuf:"fixed64,3,opt,name=lng" json:"lng,omitempty"`
}

func (m *AdRequest) Reset()                    { *m = AdRequest{} }
func (m *AdRequest) String() string            { return proto.CompactTextString(m) }
func (*AdRequest) ProtoMessage()               {}
func (*AdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SupplySegment struct {
	Audience  string `protobuf:"bytes,1,opt,name=audience" json:"audience,omitempty"`
	PlaceName string `protobuf:"bytes,2,opt,name=place_name,json=placeName" json:"place_name,omitempty"`
	PlaceType string `protobuf:"bytes,3,opt,name=place_type,json=placeType" json:"place_type,omitempty"`
}

func (m *SupplySegment) Reset()                    { *m = SupplySegment{} }
func (m *SupplySegment) String() string            { return proto.CompactTextString(m) }
func (*SupplySegment) ProtoMessage()               {}
func (*SupplySegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BlacklistResponse struct {
	Blacklist bool `protobuf:"varint,1,opt,name=blacklist" json:"blacklist,omitempty"`
}

func (m *BlacklistResponse) Reset()                    { *m = BlacklistResponse{} }
func (m *BlacklistResponse) String() string            { return proto.CompactTextString(m) }
func (*BlacklistResponse) ProtoMessage()               {}
func (*BlacklistResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*AdRequest)(nil), "segmentation.AdRequest")
	proto.RegisterType((*SupplySegment)(nil), "segmentation.SupplySegment")
	proto.RegisterType((*BlacklistResponse)(nil), "segmentation.BlacklistResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Segmentation service

type SegmentationClient interface {
	GetSegmentInfo(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*SupplySegment, error)
}

type segmentationClient struct {
	cc *grpc.ClientConn
}

func NewSegmentationClient(cc *grpc.ClientConn) SegmentationClient {
	return &segmentationClient{cc}
}

func (c *segmentationClient) GetSegmentInfo(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*SupplySegment, error) {
	out := new(SupplySegment)
	err := grpc.Invoke(ctx, "/segmentation.Segmentation/GetSegmentInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Segmentation service

type SegmentationServer interface {
	GetSegmentInfo(context.Context, *AdRequest) (*SupplySegment, error)
}

func RegisterSegmentationServer(s *grpc.Server, srv SegmentationServer) {
	s.RegisterService(&_Segmentation_serviceDesc, srv)
}

func _Segmentation_GetSegmentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentationServer).GetSegmentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/segmentation.Segmentation/GetSegmentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentationServer).GetSegmentInfo(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Segmentation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "segmentation.Segmentation",
	HandlerType: (*SegmentationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSegmentInfo",
			Handler:    _Segmentation_GetSegmentInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Blacklist service

type BlacklistClient interface {
	GetBlacklistInfo(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*BlacklistResponse, error)
}

type blacklistClient struct {
	cc *grpc.ClientConn
}

func NewBlacklistClient(cc *grpc.ClientConn) BlacklistClient {
	return &blacklistClient{cc}
}

func (c *blacklistClient) GetBlacklistInfo(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*BlacklistResponse, error) {
	out := new(BlacklistResponse)
	err := grpc.Invoke(ctx, "/segmentation.Blacklist/GetBlacklistInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Blacklist service

type BlacklistServer interface {
	GetBlacklistInfo(context.Context, *AdRequest) (*BlacklistResponse, error)
}

func RegisterBlacklistServer(s *grpc.Server, srv BlacklistServer) {
	s.RegisterService(&_Blacklist_serviceDesc, srv)
}

func _Blacklist_GetBlacklistInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlacklistServer).GetBlacklistInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/segmentation.Blacklist/GetBlacklistInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlacklistServer).GetBlacklistInfo(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blacklist_serviceDesc = grpc.ServiceDesc{
	ServiceName: "segmentation.Blacklist",
	HandlerType: (*BlacklistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlacklistInfo",
			Handler:    _Blacklist_GetBlacklistInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("segmentation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xbb, 0x16, 0xa4, 0x19, 0xaa, 0xd6, 0x5c, 0x2c, 0xad, 0x62, 0xd9, 0x53, 0x4f, 0x8b,
	0xd6, 0x5f, 0x60, 0x2f, 0xb5, 0x20, 0xa5, 0xa4, 0x1e, 0x04, 0x0f, 0x25, 0xdd, 0x8c, 0x4b, 0x30,
	0x9b, 0xc4, 0x26, 0x2b, 0xf6, 0xdf, 0xcb, 0xa6, 0xdd, 0xba, 0x8b, 0xe0, 0x6d, 0xf2, 0x3d, 0x78,
	0x2f, 0x6f, 0x06, 0xa8, 0xc3, 0x2c, 0x47, 0xed, 0xb9, 0x97, 0x46, 0x27, 0x76, 0x6b, 0xbc, 0xa1,
	0xdd, 0x3a, 0x8b, 0x9f, 0x81, 0x3c, 0x0a, 0x86, 0x9f, 0x05, 0x3a, 0x4f, 0x87, 0x40, 0x04, 0x7e,
	0xc9, 0x14, 0xd7, 0x52, 0xf4, 0xa3, 0x51, 0x34, 0x26, 0xac, 0xb3, 0x07, 0x73, 0x41, 0x7b, 0xd0,
	0x56, 0xdc, 0xf7, 0x4f, 0x46, 0xd1, 0x38, 0x62, 0xe5, 0x18, 0x88, 0xce, 0xfa, 0xed, 0x03, 0xd1,
	0x59, 0x2c, 0xe1, 0x6c, 0x55, 0x58, 0xab, 0x76, 0xab, 0x7d, 0x06, 0x1d, 0x40, 0x87, 0x17, 0x42,
	0xa2, 0x4e, 0xb1, 0x32, 0xac, 0xde, 0xf4, 0x06, 0xc0, 0x2a, 0x9e, 0xe2, 0x5a, 0xf3, 0x1c, 0x83,
	0x2f, 0x61, 0x24, 0x90, 0x05, 0xcf, 0x6b, 0xb2, 0xdf, 0x59, 0x0c, 0x21, 0x95, 0xfc, 0xb2, 0xb3,
	0x18, 0xdf, 0xc3, 0xe5, 0x54, 0xf1, 0xf4, 0x43, 0x49, 0xe7, 0x19, 0x3a, 0x6b, 0xb4, 0x43, 0x7a,
	0x0d, 0x64, 0x53, 0xc1, 0x90, 0xd7, 0x61, 0xbf, 0x60, 0xf2, 0x0a, 0xdd, 0x55, 0xad, 0x3b, 0x7d,
	0x82, 0xf3, 0x19, 0xfa, 0x03, 0x9a, 0xeb, 0x77, 0x43, 0xaf, 0x92, 0xc6, 0xc2, 0x8e, 0x9b, 0x19,
	0x0c, 0x9b, 0x42, 0xa3, 0x64, 0xdc, 0x9a, 0xbc, 0x01, 0x39, 0x7e, 0x86, 0x2e, 0xa0, 0x37, 0x43,
	0x7f, 0x7c, 0xff, 0x6f, 0x7c, 0xdb, 0x14, 0xfe, 0x54, 0x8a, 0x5b, 0xd3, 0x3b, 0x18, 0x4a, 0x93,
	0x64, 0x5b, 0x9b, 0x26, 0xf8, 0xcd, 0x73, 0xab, 0xd0, 0x25, 0x5b, 0x53, 0x78, 0xcc, 0x0a, 0x29,
	0x70, 0x7a, 0xc1, 0xca, 0x79, 0x56, 0xce, 0xcb, 0xf2, 0xc0, 0xcb, 0x68, 0x73, 0x1a, 0x2e, 0xfd,
	0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x94, 0x8b, 0x26, 0xd4, 0xff, 0x01, 0x00, 0x00,
}
