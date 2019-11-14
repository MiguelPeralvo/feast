// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/storage/Redis.proto

package storage

import (
	fmt "fmt"
	types "github.com/gojek/feast/sdk/go/protos/feast/types"
	proto "github.com/golang/protobuf/proto"
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

type RedisKey struct {
	// FeatureSet this row belongs to, this is defined as featureSetName:version.
	FeatureSet string `protobuf:"bytes,2,opt,name=feature_set,json=featureSet,proto3" json:"feature_set,omitempty"`
	// List of fields containing entity names and their respective values
	// contained within this feature row.
	Entities             []*types.Field `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RedisKey) Reset()         { *m = RedisKey{} }
func (m *RedisKey) String() string { return proto.CompactTextString(m) }
func (*RedisKey) ProtoMessage()    {}
func (*RedisKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e898a359fc9e5d, []int{0}
}

func (m *RedisKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisKey.Unmarshal(m, b)
}
func (m *RedisKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisKey.Marshal(b, m, deterministic)
}
func (m *RedisKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisKey.Merge(m, src)
}
func (m *RedisKey) XXX_Size() int {
	return xxx_messageInfo_RedisKey.Size(m)
}
func (m *RedisKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisKey.DiscardUnknown(m)
}

var xxx_messageInfo_RedisKey proto.InternalMessageInfo

func (m *RedisKey) GetFeatureSet() string {
	if m != nil {
		return m.FeatureSet
	}
	return ""
}

func (m *RedisKey) GetEntities() []*types.Field {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisKey)(nil), "feast.storage.RedisKey")
}

func init() { proto.RegisterFile("feast/storage/Redis.proto", fileDescriptor_64e898a359fc9e5d) }

var fileDescriptor_64e898a359fc9e5d = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0xaf, 0x82, 0x40,
	0x10, 0x84, 0xf3, 0x1e, 0xc9, 0x0b, 0xef, 0x88, 0xcd, 0x35, 0xa2, 0x8d, 0xc4, 0x8a, 0xea, 0x36,
	0xc1, 0x7f, 0x40, 0x61, 0x63, 0xa1, 0xc1, 0x4e, 0x0b, 0x03, 0xb2, 0x9c, 0x27, 0xea, 0x11, 0x6e,
	0x29, 0xf8, 0xf7, 0xc6, 0x85, 0x98, 0xd0, 0xee, 0xcc, 0xce, 0x37, 0x23, 0x16, 0x15, 0xe6, 0x8e,
	0xc0, 0x91, 0x6d, 0x73, 0x8d, 0x90, 0x61, 0x69, 0x9c, 0x6a, 0x5a, 0x4b, 0x56, 0xce, 0x58, 0x52,
	0xa3, 0xb4, 0x9c, 0x0f, 0x4e, 0xea, 0x1b, 0x74, 0xb0, 0x35, 0xf8, 0x28, 0x07, 0xdf, 0xfa, 0x2c,
	0x7c, 0x7e, 0xdb, 0x61, 0x2f, 0x57, 0x22, 0xa8, 0x30, 0xa7, 0xae, 0xc5, 0x8b, 0x43, 0x0a, 0x7f,
	0xa3, 0x9f, 0xf8, 0x3f, 0x13, 0xe3, 0xe9, 0x88, 0x24, 0x95, 0xf0, 0xf1, 0x45, 0x86, 0x0c, 0xba,
	0xd0, 0x8b, 0xbc, 0x38, 0x48, 0xa4, 0x1a, 0x38, 0x1c, 0xac, 0x38, 0x38, 0xfb, 0x7a, 0xd2, 0xbd,
	0x98, 0xd6, 0x48, 0x05, 0xb3, 0x0e, 0x1f, 0xf2, 0x29, 0xd1, 0x86, 0x6e, 0x5d, 0xa1, 0xae, 0xf6,
	0x09, 0xda, 0xde, 0xb1, 0x86, 0x71, 0x4d, 0x59, 0x83, 0xb6, 0xc0, 0xf5, 0x1c, 0x4c, 0x16, 0x16,
	0x7f, 0x7c, 0xdd, 0xbc, 0x03, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2b, 0x58, 0x50, 0xf9, 0x00, 0x00,
	0x00,
}
