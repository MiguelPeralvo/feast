// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/serving/Serving.proto

package serving

import (
	context "context"
	fmt "fmt"
	types "github.com/gojek/feast/protos/generated/go/feast/types"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type QueryFeaturesRequest struct {
	// e.g. "driver", "customer", "city".
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// List of entity ID.
	EntityId []string `protobuf:"bytes,2,rep,name=entityId,proto3" json:"entityId,omitempty"`
	// List of feature ID.
	// feature ID is in the form of [entity_name].[granularity].[feature_name]
	// e.g: "driver.day.total_accepted_booking"
	// all requested feature ID shall have same entity name.
	FeatureId            []string `protobuf:"bytes,3,rep,name=featureId,proto3" json:"featureId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryFeaturesRequest) Reset()         { *m = QueryFeaturesRequest{} }
func (m *QueryFeaturesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeaturesRequest) ProtoMessage()    {}
func (*QueryFeaturesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7609de6de542e6f0, []int{0}
}

func (m *QueryFeaturesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFeaturesRequest.Unmarshal(m, b)
}
func (m *QueryFeaturesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFeaturesRequest.Marshal(b, m, deterministic)
}
func (m *QueryFeaturesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeaturesRequest.Merge(m, src)
}
func (m *QueryFeaturesRequest) XXX_Size() int {
	return xxx_messageInfo_QueryFeaturesRequest.Size(m)
}
func (m *QueryFeaturesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeaturesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeaturesRequest proto.InternalMessageInfo

func (m *QueryFeaturesRequest) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *QueryFeaturesRequest) GetEntityId() []string {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *QueryFeaturesRequest) GetFeatureId() []string {
	if m != nil {
		return m.FeatureId
	}
	return nil
}

type QueryFeaturesResponse struct {
	// Entity name of the response
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// map of entity ID and its entity's properties.
	Entities             map[string]*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryFeaturesResponse) Reset()         { *m = QueryFeaturesResponse{} }
func (m *QueryFeaturesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeaturesResponse) ProtoMessage()    {}
func (*QueryFeaturesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7609de6de542e6f0, []int{1}
}

func (m *QueryFeaturesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryFeaturesResponse.Unmarshal(m, b)
}
func (m *QueryFeaturesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryFeaturesResponse.Marshal(b, m, deterministic)
}
func (m *QueryFeaturesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeaturesResponse.Merge(m, src)
}
func (m *QueryFeaturesResponse) XXX_Size() int {
	return xxx_messageInfo_QueryFeaturesResponse.Size(m)
}
func (m *QueryFeaturesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeaturesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeaturesResponse proto.InternalMessageInfo

func (m *QueryFeaturesResponse) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *QueryFeaturesResponse) GetEntities() map[string]*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type Entity struct {
	// map of feature ID and its feature value.
	Features             map[string]*FeatureValue `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7609de6de542e6f0, []int{2}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetFeatures() map[string]*FeatureValue {
	if m != nil {
		return m.Features
	}
	return nil
}

type FeatureValue struct {
	// value of feature
	Value *types.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp of the feature
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureValue) Reset()         { *m = FeatureValue{} }
func (m *FeatureValue) String() string { return proto.CompactTextString(m) }
func (*FeatureValue) ProtoMessage()    {}
func (*FeatureValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_7609de6de542e6f0, []int{3}
}

func (m *FeatureValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureValue.Unmarshal(m, b)
}
func (m *FeatureValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureValue.Marshal(b, m, deterministic)
}
func (m *FeatureValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureValue.Merge(m, src)
}
func (m *FeatureValue) XXX_Size() int {
	return xxx_messageInfo_FeatureValue.Size(m)
}
func (m *FeatureValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureValue.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureValue proto.InternalMessageInfo

func (m *FeatureValue) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FeatureValue) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryFeaturesRequest)(nil), "feast.serving.QueryFeaturesRequest")
	proto.RegisterType((*QueryFeaturesResponse)(nil), "feast.serving.QueryFeaturesResponse")
	proto.RegisterMapType((map[string]*Entity)(nil), "feast.serving.QueryFeaturesResponse.EntitiesEntry")
	proto.RegisterType((*Entity)(nil), "feast.serving.Entity")
	proto.RegisterMapType((map[string]*FeatureValue)(nil), "feast.serving.Entity.FeaturesEntry")
	proto.RegisterType((*FeatureValue)(nil), "feast.serving.FeatureValue")
}

func init() { proto.RegisterFile("feast/serving/Serving.proto", fileDescriptor_7609de6de542e6f0) }

var fileDescriptor_7609de6de542e6f0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x75, 0x5a, 0x5c, 0x36, 0x77, 0x0d, 0xca, 0xe0, 0x62, 0xc8, 0x8a, 0x96, 0xac, 0x0f, 0x01,
	0x61, 0x06, 0xe3, 0x4b, 0xf1, 0x45, 0x5c, 0x58, 0xa1, 0x2f, 0x4b, 0x8d, 0x22, 0x52, 0x7c, 0x49,
	0xed, 0x6d, 0x4c, 0xdb, 0x64, 0x62, 0x66, 0x52, 0xc8, 0xf7, 0xf8, 0x59, 0xfe, 0x8c, 0x74, 0x66,
	0xd2, 0x26, 0xa5, 0xe8, 0x3e, 0x25, 0x33, 0xe7, 0xde, 0x73, 0xce, 0xcc, 0x3d, 0x03, 0x57, 0x4b,
	0x4c, 0xa4, 0xe2, 0x12, 0xab, 0x6d, 0x56, 0xa4, 0xfc, 0xb3, 0xf9, 0xb2, 0xb2, 0x12, 0x4a, 0x50,
	0x57, 0x83, 0xcc, 0x82, 0xfe, 0xcb, 0x54, 0x88, 0x74, 0x83, 0x5c, 0x83, 0xf3, 0x7a, 0xc9, 0x55,
	0x96, 0xa3, 0x54, 0x49, 0x5e, 0x9a, 0x7a, 0xff, 0x99, 0x21, 0x53, 0x4d, 0x89, 0x92, 0x7f, 0x4d,
	0x36, 0x35, 0x1a, 0x20, 0x28, 0xe1, 0xe9, 0xa7, 0x1a, 0xab, 0xe6, 0x23, 0x26, 0xaa, 0xae, 0x50,
	0xc6, 0xf8, 0xab, 0x46, 0xa9, 0xe8, 0x0b, 0x00, 0x2c, 0x54, 0xa6, 0x9a, 0xbb, 0x24, 0x47, 0x8f,
	0x8c, 0x48, 0xe8, 0xc4, 0x9d, 0x1d, 0xea, 0xc3, 0xb9, 0x59, 0x4d, 0x16, 0xde, 0x60, 0x34, 0x0c,
	0x9d, 0x78, 0xbf, 0xa6, 0xcf, 0xc1, 0x59, 0x1a, 0xba, 0xc9, 0xc2, 0x1b, 0x6a, 0xf0, 0xb0, 0x11,
	0xfc, 0x21, 0x70, 0x79, 0x24, 0x29, 0x4b, 0x51, 0x48, 0xfc, 0xaf, 0xe6, 0x9d, 0xd5, 0xcc, 0x50,
	0x6a, 0xcd, 0x8b, 0x28, 0x62, 0xbd, 0x7b, 0x60, 0x27, 0x79, 0xd9, 0xad, 0x6d, 0xba, 0x2d, 0x54,
	0xd5, 0xc4, 0x7b, 0x0e, 0x3f, 0x06, 0xb7, 0x07, 0xd1, 0x27, 0x30, 0x5c, 0x63, 0x63, 0x95, 0x77,
	0xbf, 0xf4, 0x35, 0x3c, 0xdc, 0xee, 0x6e, 0xcb, 0x1b, 0x8c, 0x48, 0x78, 0x11, 0x5d, 0x1e, 0xe9,
	0xe9, 0xf6, 0x26, 0x36, 0x35, 0xef, 0x06, 0x63, 0x12, 0xfc, 0x26, 0x70, 0x66, 0x76, 0xe9, 0x7b,
	0x38, 0xb7, 0xa7, 0x96, 0x1e, 0xd1, 0x76, 0xaf, 0x4f, 0xb6, 0xb3, 0xd6, 0xb0, 0xf5, 0xd7, 0x36,
	0xf9, 0xdf, 0xc0, 0xed, 0x41, 0x27, 0xfc, 0xbd, 0xe9, 0xfb, 0xbb, 0x3a, 0x12, 0xb0, 0xed, 0x7a,
	0xe0, 0x5d, 0x97, 0x15, 0x3c, 0xea, 0x42, 0x34, 0x6c, 0x69, 0x88, 0xa6, 0xa1, 0x96, 0x46, 0xc7,
	0x85, 0x75, 0xbb, 0xe9, 0x18, 0x9c, 0x7d, 0xb6, 0xac, 0xa8, 0xcf, 0x4c, 0xfa, 0x58, 0x9b, 0x3e,
	0xf6, 0xa5, 0xad, 0x88, 0x0f, 0xc5, 0xd1, 0x0a, 0xc0, 0x66, 0xf8, 0xc3, 0x74, 0x42, 0xbf, 0x83,
	0xdb, 0x1b, 0x16, 0xbd, 0xfe, 0xf7, 0x28, 0x75, 0x2a, 0xfd, 0x57, 0xf7, 0x99, 0x77, 0xf0, 0xe0,
	0x66, 0x06, 0xfd, 0x07, 0x72, 0xf3, 0xf8, 0x20, 0x3d, 0xdd, 0xb9, 0x9c, 0x8d, 0xd3, 0x4c, 0xfd,
	0xac, 0xe7, 0xec, 0x87, 0xc8, 0x79, 0x2a, 0x56, 0xb8, 0xe6, 0xe6, 0x85, 0xe8, 0x33, 0x48, 0x9e,
	0x62, 0x81, 0x55, 0xa2, 0x70, 0xc1, 0x53, 0xc1, 0x7b, 0x0f, 0x71, 0x7e, 0xa6, 0x4b, 0xde, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x77, 0xb5, 0x0e, 0xd1, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServingAPIClient is the client API for ServingAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServingAPIClient interface {
	// Query features from Feast serving storage
	QueryFeatures(ctx context.Context, in *QueryFeaturesRequest, opts ...grpc.CallOption) (*QueryFeaturesResponse, error)
}

type servingAPIClient struct {
	cc *grpc.ClientConn
}

func NewServingAPIClient(cc *grpc.ClientConn) ServingAPIClient {
	return &servingAPIClient{cc}
}

func (c *servingAPIClient) QueryFeatures(ctx context.Context, in *QueryFeaturesRequest, opts ...grpc.CallOption) (*QueryFeaturesResponse, error) {
	out := new(QueryFeaturesResponse)
	err := c.cc.Invoke(ctx, "/feast.serving.ServingAPI/QueryFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServingAPIServer is the server API for ServingAPI service.
type ServingAPIServer interface {
	// Query features from Feast serving storage
	QueryFeatures(context.Context, *QueryFeaturesRequest) (*QueryFeaturesResponse, error)
}

// UnimplementedServingAPIServer can be embedded to have forward compatible implementations.
type UnimplementedServingAPIServer struct {
}

func (*UnimplementedServingAPIServer) QueryFeatures(ctx context.Context, req *QueryFeaturesRequest) (*QueryFeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFeatures not implemented")
}

func RegisterServingAPIServer(s *grpc.Server, srv ServingAPIServer) {
	s.RegisterService(&_ServingAPI_serviceDesc, srv)
}

func _ServingAPI_QueryFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingAPIServer).QueryFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.serving.ServingAPI/QueryFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingAPIServer).QueryFeatures(ctx, req.(*QueryFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServingAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.serving.ServingAPI",
	HandlerType: (*ServingAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFeatures",
			Handler:    _ServingAPI_QueryFeatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/serving/Serving.proto",
}
