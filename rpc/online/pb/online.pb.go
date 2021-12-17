// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/online/online.proto

package online_service // import "rpc/online/pb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_online_41e77e251fc3842e, []int{0}
}
func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (dst *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(dst, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type OnlineResponse struct {
	Uids                 []int64  `protobuf:"varint,1,rep,packed,name=uids" json:"uids,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineResponse) Reset()         { *m = OnlineResponse{} }
func (m *OnlineResponse) String() string { return proto.CompactTextString(m) }
func (*OnlineResponse) ProtoMessage()    {}
func (*OnlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_online_41e77e251fc3842e, []int{1}
}
func (m *OnlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineResponse.Unmarshal(m, b)
}
func (m *OnlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineResponse.Marshal(b, m, deterministic)
}
func (dst *OnlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineResponse.Merge(dst, src)
}
func (m *OnlineResponse) XXX_Size() int {
	return xxx_messageInfo_OnlineResponse.Size(m)
}
func (m *OnlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineResponse proto.InternalMessageInfo

func (m *OnlineResponse) GetUids() []int64 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *OnlineResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type OnlineRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineRequest) Reset()         { *m = OnlineRequest{} }
func (m *OnlineRequest) String() string { return proto.CompactTextString(m) }
func (*OnlineRequest) ProtoMessage()    {}
func (*OnlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_online_41e77e251fc3842e, []int{2}
}
func (m *OnlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineRequest.Unmarshal(m, b)
}
func (m *OnlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineRequest.Marshal(b, m, deterministic)
}
func (dst *OnlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineRequest.Merge(dst, src)
}
func (m *OnlineRequest) XXX_Size() int {
	return xxx_messageInfo_OnlineRequest.Size(m)
}
func (m *OnlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineRequest proto.InternalMessageInfo

func (m *OnlineRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_online_41e77e251fc3842e, []int{3}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (dst *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(dst, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "online.EmptyRequest")
	proto.RegisterType((*OnlineResponse)(nil), "online.OnlineResponse")
	proto.RegisterType((*OnlineRequest)(nil), "online.OnlineRequest")
	proto.RegisterType((*EmptyResponse)(nil), "online.EmptyResponse")
}

func init() { proto.RegisterFile("proto/online/online.proto", fileDescriptor_online_41e77e251fc3842e) }

var fileDescriptor_online_41e77e251fc3842e = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcf, 0xcb, 0xc9, 0xcc, 0x4b, 0x85, 0x52, 0x7a, 0x60, 0x31, 0x21, 0x36, 0x08,
	0x4f, 0x89, 0x8f, 0x8b, 0xc7, 0x35, 0xb7, 0xa0, 0xa4, 0x32, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8,
	0x44, 0xc9, 0x8a, 0x8b, 0xcf, 0x1f, 0x2c, 0x13, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a,
	0x24, 0xc4, 0xc5, 0x52, 0x9a, 0x99, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x1c, 0x04, 0x66,
	0x0b, 0x89, 0x70, 0xb1, 0x26, 0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07,
	0x41, 0x38, 0x4a, 0x8a, 0x5c, 0xbc, 0x30, 0xbd, 0x60, 0xc3, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x33,
	0x53, 0x24, 0x18, 0xc1, 0x8a, 0x40, 0x4c, 0x25, 0x7e, 0x2e, 0x5e, 0xa8, 0x75, 0x10, 0xd3, 0x8d,
	0xf6, 0x33, 0x72, 0xf1, 0x40, 0x34, 0x05, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x09, 0x99, 0x73, 0xb1,
	0x07, 0x97, 0x24, 0x16, 0x95, 0x84, 0x16, 0x08, 0x89, 0xea, 0x41, 0x9d, 0x8c, 0x62, 0xaa, 0x14,
	0x5c, 0x18, 0xc5, 0x24, 0x21, 0x0b, 0x2e, 0x8e, 0xe0, 0x8c, 0xd2, 0x92, 0x94, 0xfc, 0xf2, 0x3c,
	0x12, 0x75, 0xda, 0x70, 0xf1, 0xb8, 0xa7, 0x96, 0x40, 0x94, 0x3a, 0xe6, 0xe4, 0x08, 0x89, 0xa0,
	0x29, 0x83, 0x68, 0x16, 0x43, 0x37, 0x13, 0xa2, 0xdb, 0x49, 0x2e, 0x4a, 0xa6, 0xa8, 0x20, 0x19,
	0x16, 0xc8, 0x05, 0x49, 0xd6, 0x10, 0x56, 0x7c, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0x6a, 0x12,
	0x1b, 0x38, 0xc0, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x9b, 0xec, 0x1e, 0x8d, 0x01,
	0x00, 0x00,
}
