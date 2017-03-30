// Code generated by protoc-gen-gogo.
// source: proto3.proto
// DO NOT EDIT!

/*
Package prototests is a generated protocol buffer package.

It is generated from these files:
	proto3.proto

It has these top-level messages:
	Proto3
	SmallMsg3
*/
package prototests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Proto3 is a proto3 message.
type Proto3 struct {
	Field int64      `protobuf:"varint,1,opt,name=Field,proto3" json:"Field,omitempty"`
	Msg   *SmallMsg3 `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	Ints  []int64    `protobuf:"varint,4,rep,packed,name=Ints" json:"Ints,omitempty"`
}

func (m *Proto3) Reset()                    { *m = Proto3{} }
func (m *Proto3) String() string            { return proto.CompactTextString(m) }
func (*Proto3) ProtoMessage()               {}
func (*Proto3) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func (m *Proto3) GetField() int64 {
	if m != nil {
		return m.Field
	}
	return 0
}

func (m *Proto3) GetMsg() *SmallMsg3 {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Proto3) GetInts() []int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

// SmallMsg3 only contains some native fields.
type SmallMsg3 struct {
	ScarBusStop     string   `protobuf:"bytes,1,opt,name=ScarBusStop,proto3" json:"ScarBusStop,omitempty"`
	FlightParachute []uint32 `protobuf:"fixed32,12,rep,packed,name=FlightParachute" json:"FlightParachute,omitempty"`
}

func (m *SmallMsg3) Reset()                    { *m = SmallMsg3{} }
func (m *SmallMsg3) String() string            { return proto.CompactTextString(m) }
func (*SmallMsg3) ProtoMessage()               {}
func (*SmallMsg3) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{1} }

func (m *SmallMsg3) GetScarBusStop() string {
	if m != nil {
		return m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg3) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func init() {
	proto.RegisterType((*Proto3)(nil), "prototests.Proto3")
	proto.RegisterType((*SmallMsg3)(nil), "prototests.SmallMsg3")
}
func (this *Proto3) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return Proto3Description()
}
func (this *SmallMsg3) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return Proto3Description()
}
func Proto3Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3669 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x70, 0x1b, 0xd7,
		0x75, 0x36, 0xfe, 0x48, 0xe0, 0x00, 0x04, 0x97, 0x97, 0xb4, 0xb4, 0xa2, 0x63, 0x89, 0x62, 0xec,
		0x88, 0xb6, 0x1b, 0x28, 0x23, 0x59, 0xb2, 0xbc, 0x6a, 0xe2, 0x01, 0x41, 0x88, 0x81, 0x4a, 0x12,
		0xc8, 0x82, 0x8c, 0xe5, 0xe4, 0x61, 0xe7, 0x72, 0x71, 0x01, 0xae, 0xb4, 0xd8, 0x45, 0x76, 0x17,
		0x92, 0xe9, 0x27, 0x77, 0xdc, 0x9f, 0xc9, 0x74, 0xfa, 0xdf, 0x4e, 0x13, 0xd7, 0x71, 0xdb, 0xcc,
		0xb4, 0x6e, 0x93, 0xfe, 0x24, 0xfd, 0x49, 0x3b, 0x7d, 0xea, 0x4b, 0xda, 0x3e, 0x75, 0x26, 0xef,
		0x7d, 0x68, 0x5a, 0xcf, 0xf4, 0xcf, 0x6d, 0xd3, 0xc6, 0x33, 0xed, 0x8c, 0x5e, 0x3a, 0xf7, 0x6f,
		0xb1, 0x0b, 0x80, 0x5a, 0x30, 0x33, 0x8e, 0x9f, 0xc8, 0x7b, 0xee, 0xf9, 0xbe, 0x3d, 0x7b, 0xee,
		0xb9, 0xe7, 0x9c, 0x7b, 0x17, 0xf0, 0x4f, 0xd7, 0x60, 0xad, 0xe7, 0xba, 0x3d, 0x9b, 0x5c, 0x1e,
		0x78, 0x6e, 0xe0, 0x1e, 0x0e, 0xbb, 0x97, 0x3b, 0xc4, 0x37, 0x3d, 0x6b, 0x10, 0xb8, 0x5e, 0x85,
		0xc9, 0xd0, 0x22, 0xd7, 0xa8, 0x48, 0x8d, 0xf5, 0x5d, 0x58, 0xba, 0x65, 0xd9, 0x64, 0x2b, 0x54,
		0x6c, 0x93, 0x00, 0xdd, 0x80, 0x6c, 0xd7, 0xb2, 0x89, 0x9a, 0x5a, 0xcb, 0x6c, 0x14, 0xaf, 0x3c,
		0x55, 0x19, 0x03, 0x55, 0xe2, 0x88, 0x16, 0x15, 0xeb, 0x0c, 0xb1, 0xfe, 0x6e, 0x16, 0x96, 0xa7,
		0xcc, 0x22, 0x04, 0x59, 0x07, 0xf7, 0x29, 0x63, 0x6a, 0xa3, 0xa0, 0xb3, 0xff, 0x91, 0x0a, 0xf3,
		0x03, 0x6c, 0xde, 0xc3, 0x3d, 0xa2, 0xa6, 0x99, 0x58, 0x0e, 0xd1, 0x79, 0x80, 0x0e, 0x19, 0x10,
		0xa7, 0x43, 0x1c, 0xf3, 0x58, 0xcd, 0xac, 0x65, 0x36, 0x0a, 0x7a, 0x44, 0x82, 0x9e, 0x83, 0xa5,
		0xc1, 0xf0, 0xd0, 0xb6, 0x4c, 0x23, 0xa2, 0x06, 0x6b, 0x99, 0x8d, 0x9c, 0xae, 0xf0, 0x89, 0xad,
		0x91, 0xf2, 0x25, 0x58, 0x7c, 0x40, 0xf0, 0xbd, 0xa8, 0x6a, 0x91, 0xa9, 0x96, 0xa9, 0x38, 0xa2,
		0x58, 0x83, 0x52, 0x9f, 0xf8, 0x3e, 0xee, 0x11, 0x23, 0x38, 0x1e, 0x10, 0x35, 0xcb, 0xde, 0x7e,
		0x6d, 0xe2, 0xed, 0xc7, 0xdf, 0xbc, 0x28, 0x50, 0xfb, 0xc7, 0x03, 0x82, 0xaa, 0x50, 0x20, 0xce,
		0xb0, 0xcf, 0x19, 0x72, 0x27, 0xf8, 0xaf, 0xee, 0x0c, 0xfb, 0xe3, 0x2c, 0x79, 0x0a, 0x13, 0x14,
		0xf3, 0x3e, 0xf1, 0xee, 0x5b, 0x26, 0x51, 0xe7, 0x18, 0xc1, 0xa5, 0x09, 0x82, 0x36, 0x9f, 0x1f,
		0xe7, 0x90, 0x38, 0x54, 0x83, 0x02, 0x79, 0x35, 0x20, 0x8e, 0x6f, 0xb9, 0x8e, 0x3a, 0xcf, 0x48,
		0x9e, 0x9e, 0xb2, 0x8a, 0xc4, 0xee, 0x8c, 0x53, 0x8c, 0x70, 0xe8, 0x3a, 0xcc, 0xbb, 0x83, 0xc0,
		0x72, 0x1d, 0x5f, 0xcd, 0xaf, 0xa5, 0x36, 0x8a, 0x57, 0x3e, 0x32, 0x35, 0x10, 0x9a, 0x5c, 0x47,
		0x97, 0xca, 0xa8, 0x01, 0x8a, 0xef, 0x0e, 0x3d, 0x93, 0x18, 0xa6, 0xdb, 0x21, 0x86, 0xe5, 0x74,
		0x5d, 0xb5, 0xc0, 0x08, 0x2e, 0x4c, 0xbe, 0x08, 0x53, 0xac, 0xb9, 0x1d, 0xd2, 0x70, 0xba, 0xae,
		0x5e, 0xf6, 0x63, 0x63, 0x74, 0x06, 0xe6, 0xfc, 0x63, 0x27, 0xc0, 0xaf, 0xaa, 0x25, 0x16, 0x21,
		0x62, 0xb4, 0xfe, 0xbf, 0x39, 0x58, 0x9c, 0x25, 0xc4, 0x6e, 0x42, 0xae, 0x4b, 0xdf, 0x52, 0x4d,
		0x9f, 0xc6, 0x07, 0x1c, 0x13, 0x77, 0xe2, 0xdc, 0x0f, 0xe8, 0xc4, 0x2a, 0x14, 0x1d, 0xe2, 0x07,
		0xa4, 0xc3, 0x23, 0x22, 0x33, 0x63, 0x4c, 0x01, 0x07, 0x4d, 0x86, 0x54, 0xf6, 0x07, 0x0a, 0xa9,
		0x3b, 0xb0, 0x18, 0x9a, 0x64, 0x78, 0xd8, 0xe9, 0xc9, 0xd8, 0xbc, 0x9c, 0x64, 0x49, 0xa5, 0x2e,
		0x71, 0x3a, 0x85, 0xe9, 0x65, 0x12, 0x1b, 0xa3, 0x2d, 0x00, 0xd7, 0x21, 0x6e, 0xd7, 0xe8, 0x10,
		0xd3, 0x56, 0xf3, 0x27, 0x78, 0xa9, 0x49, 0x55, 0x26, 0xbc, 0xe4, 0x72, 0xa9, 0x69, 0xa3, 0x17,
		0x47, 0xa1, 0x36, 0x7f, 0x42, 0xa4, 0xec, 0xf2, 0x4d, 0x36, 0x11, 0x6d, 0x07, 0x50, 0xf6, 0x08,
		0x8d, 0x7b, 0xd2, 0x11, 0x6f, 0x56, 0x60, 0x46, 0x54, 0x12, 0xdf, 0x4c, 0x17, 0x30, 0xfe, 0x62,
		0x0b, 0x5e, 0x74, 0x88, 0x3e, 0x0a, 0xa1, 0xc0, 0x60, 0x61, 0x05, 0x2c, 0x0b, 0x95, 0xa4, 0x70,
		0x0f, 0xf7, 0xc9, 0xea, 0x0d, 0x28, 0xc7, 0xdd, 0x83, 0x56, 0x20, 0xe7, 0x07, 0xd8, 0x0b, 0x58,
		0x14, 0xe6, 0x74, 0x3e, 0x40, 0x0a, 0x64, 0x88, 0xd3, 0x61, 0x59, 0x2e, 0xa7, 0xd3, 0x7f, 0x57,
		0x5f, 0x80, 0x85, 0xd8, 0xe3, 0x67, 0x05, 0xae, 0x7f, 0x69, 0x0e, 0x56, 0xa6, 0xc5, 0xdc, 0xd4,
		0xf0, 0x3f, 0x03, 0x73, 0xce, 0xb0, 0x7f, 0x48, 0x3c, 0x35, 0xc3, 0x18, 0xc4, 0x08, 0x55, 0x21,
		0x67, 0xe3, 0x43, 0x62, 0xab, 0xd9, 0xb5, 0xd4, 0x46, 0xf9, 0xca, 0x73, 0x33, 0x45, 0x75, 0x65,
		0x87, 0x42, 0x74, 0x8e, 0x44, 0x9f, 0x82, 0xac, 0x48, 0x71, 0x94, 0xe1, 0xd9, 0xd9, 0x18, 0x68,
		0x2c, 0xea, 0x0c, 0x87, 0x9e, 0x80, 0x02, 0xfd, 0xcb, 0x7d, 0x3b, 0xc7, 0x6c, 0xce, 0x53, 0x01,
		0xf5, 0x2b, 0x5a, 0x85, 0x3c, 0x0b, 0xb3, 0x0e, 0x91, 0xa5, 0x21, 0x1c, 0xd3, 0x85, 0xe9, 0x90,
		0x2e, 0x1e, 0xda, 0x81, 0x71, 0x1f, 0xdb, 0x43, 0xc2, 0x02, 0xa6, 0xa0, 0x97, 0x84, 0xf0, 0xb3,
		0x54, 0x86, 0x2e, 0x40, 0x91, 0x47, 0xa5, 0xe5, 0x74, 0xc8, 0xab, 0x2c, 0xfb, 0xe4, 0x74, 0x1e,
		0xa8, 0x0d, 0x2a, 0xa1, 0x8f, 0xbf, 0xeb, 0xbb, 0x8e, 0x5c, 0x5a, 0xf6, 0x08, 0x2a, 0x60, 0x8f,
		0x7f, 0x61, 0x3c, 0xf1, 0x3d, 0x39, 0xfd, 0xf5, 0xc6, 0x63, 0x71, 0xfd, 0x5b, 0x69, 0xc8, 0xb2,
		0xfd, 0xb6, 0x08, 0xc5, 0xfd, 0x57, 0x5a, 0x75, 0x63, 0xab, 0x79, 0xb0, 0xb9, 0x53, 0x57, 0x52,
		0xa8, 0x0c, 0xc0, 0x04, 0xb7, 0x76, 0x9a, 0xd5, 0x7d, 0x25, 0x1d, 0x8e, 0x1b, 0x7b, 0xfb, 0xd7,
		0x9f, 0x57, 0x32, 0x21, 0xe0, 0x80, 0x0b, 0xb2, 0x51, 0x85, 0xab, 0x57, 0x94, 0x1c, 0x52, 0xa0,
		0xc4, 0x09, 0x1a, 0x77, 0xea, 0x5b, 0xd7, 0x9f, 0x57, 0xe6, 0xe2, 0x92, 0xab, 0x57, 0x94, 0x79,
		0xb4, 0x00, 0x05, 0x26, 0xd9, 0x6c, 0x36, 0x77, 0x94, 0x7c, 0xc8, 0xd9, 0xde, 0xd7, 0x1b, 0x7b,
		0xdb, 0x4a, 0x21, 0xe4, 0xdc, 0xd6, 0x9b, 0x07, 0x2d, 0x05, 0x42, 0x86, 0xdd, 0x7a, 0xbb, 0x5d,
		0xdd, 0xae, 0x2b, 0xc5, 0x50, 0x63, 0xf3, 0x95, 0xfd, 0x7a, 0x5b, 0x29, 0xc5, 0xcc, 0xba, 0x7a,
		0x45, 0x59, 0x08, 0x1f, 0x51, 0xdf, 0x3b, 0xd8, 0x55, 0xca, 0x68, 0x09, 0x16, 0xf8, 0x23, 0xa4,
		0x11, 0x8b, 0x63, 0xa2, 0xeb, 0xcf, 0x2b, 0xca, 0xc8, 0x10, 0xce, 0xb2, 0x14, 0x13, 0x5c, 0x7f,
		0x5e, 0x41, 0xeb, 0x35, 0xc8, 0xb1, 0xe8, 0x42, 0x08, 0xca, 0x3b, 0xd5, 0xcd, 0xfa, 0x8e, 0xd1,
		0x6c, 0xed, 0x37, 0x9a, 0x7b, 0xd5, 0x1d, 0x25, 0x35, 0x92, 0xe9, 0xf5, 0xcf, 0x1c, 0x34, 0xf4,
		0xfa, 0x96, 0x92, 0x8e, 0xca, 0x5a, 0xf5, 0xea, 0x7e, 0x7d, 0x4b, 0xc9, 0xac, 0x9b, 0xb0, 0x32,
		0x2d, 0xcf, 0x4c, 0xdd, 0x19, 0x91, 0x25, 0x4e, 0x9f, 0xb0, 0xc4, 0x8c, 0x6b, 0x62, 0x89, 0xbf,
		0x9a, 0x82, 0xe5, 0x29, 0xb9, 0x76, 0xea, 0x43, 0x5e, 0x82, 0x1c, 0x0f, 0x51, 0x5e, 0x7d, 0x9e,
		0x99, 0x9a, 0xb4, 0x59, 0xc0, 0x4e, 0x54, 0x20, 0x86, 0x8b, 0x56, 0xe0, 0xcc, 0x09, 0x15, 0x98,
		0x52, 0x4c, 0x18, 0xf9, 0x46, 0x0a, 0xd4, 0x93, 0xb8, 0x13, 0x12, 0x45, 0x3a, 0x96, 0x28, 0x6e,
		0x8e, 0x1b, 0x70, 0xf1, 0xe4, 0x77, 0x98, 0xb0, 0xe2, 0x9d, 0x14, 0x9c, 0x99, 0xde, 0xa8, 0x4c,
		0xb5, 0xe1, 0x53, 0x30, 0xd7, 0x27, 0xc1, 0x91, 0x2b, 0x8b, 0xf5, 0xc7, 0xa6, 0x94, 0x00, 0x3a,
		0x3d, 0xee, 0x2b, 0x81, 0x8a, 0xd6, 0x90, 0xcc, 0x49, 0xdd, 0x06, 0xb7, 0x66, 0xc2, 0xd2, 0x2f,
		0xa6, 0xe1, 0xf1, 0xa9, 0xe4, 0x53, 0x0d, 0x7d, 0x12, 0xc0, 0x72, 0x06, 0xc3, 0x80, 0x17, 0x64,
		0x9e, 0x9f, 0x0a, 0x4c, 0xc2, 0xf6, 0x3e, 0xcd, 0x3d, 0xc3, 0x20, 0x9c, 0xcf, 0xb0, 0x79, 0xe0,
		0x22, 0xa6, 0x70, 0x63, 0x64, 0x68, 0x96, 0x19, 0x7a, 0xfe, 0x84, 0x37, 0x9d, 0xa8, 0x75, 0x9f,
		0x00, 0xc5, 0xb4, 0x2d, 0xe2, 0x04, 0x86, 0x1f, 0x78, 0x04, 0xf7, 0x2d, 0xa7, 0xc7, 0x12, 0x70,
		0x5e, 0xcb, 0x75, 0xb1, 0xed, 0x13, 0x7d, 0x91, 0x4f, 0xb7, 0xe5, 0x2c, 0x45, 0xb0, 0x2a, 0xe3,
		0x45, 0x10, 0x73, 0x31, 0x04, 0x9f, 0x0e, 0x11, 0xeb, 0x5f, 0x9f, 0x87, 0x62, 0xa4, 0xad, 0x43,
		0x17, 0xa1, 0x74, 0x17, 0xdf, 0xc7, 0x86, 0x6c, 0xd5, 0xb9, 0x27, 0x8a, 0x54, 0xd6, 0x12, 0xed,
		0xfa, 0x27, 0x60, 0x85, 0xa9, 0xb8, 0xc3, 0x80, 0x78, 0x86, 0x69, 0x63, 0xdf, 0x67, 0x4e, 0xcb,
		0x33, 0x55, 0x44, 0xe7, 0x9a, 0x74, 0xaa, 0x26, 0x67, 0xd0, 0x35, 0x58, 0x66, 0x88, 0xfe, 0xd0,
		0x0e, 0xac, 0x81, 0x4d, 0x0c, 0x7a, 0x78, 0xf0, 0x59, 0x22, 0x0e, 0x2d, 0x5b, 0xa2, 0x1a, 0xbb,
		0x42, 0x81, 0x5a, 0xe4, 0xa3, 0x2d, 0x78, 0x92, 0xc1, 0x7a, 0xc4, 0x21, 0x1e, 0x0e, 0x88, 0x41,
		0xbe, 0x30, 0xc4, 0xb6, 0x6f, 0x60, 0xa7, 0x63, 0x1c, 0x61, 0xff, 0x48, 0x5d, 0xa1, 0x04, 0x9b,
		0x69, 0x35, 0xa5, 0x9f, 0xa3, 0x8a, 0xdb, 0x42, 0xaf, 0xce, 0xd4, 0xaa, 0x4e, 0xe7, 0xd3, 0xd8,
		0x3f, 0x42, 0x1a, 0x9c, 0x61, 0x2c, 0x7e, 0xe0, 0x59, 0x4e, 0xcf, 0x30, 0x8f, 0x88, 0x79, 0xcf,
		0x18, 0x06, 0xdd, 0x1b, 0xea, 0x13, 0xd1, 0xe7, 0x33, 0x0b, 0xdb, 0x4c, 0xa7, 0x46, 0x55, 0x0e,
		0x82, 0xee, 0x0d, 0xd4, 0x86, 0x12, 0x5d, 0x8c, 0xbe, 0xf5, 0x1a, 0x31, 0xba, 0xae, 0xc7, 0x2a,
		0x4b, 0x79, 0xca, 0xce, 0x8e, 0x78, 0xb0, 0xd2, 0x14, 0x80, 0x5d, 0xb7, 0x43, 0xb4, 0x5c, 0xbb,
		0x55, 0xaf, 0x6f, 0xe9, 0x45, 0xc9, 0x72, 0xcb, 0xf5, 0x68, 0x40, 0xf5, 0xdc, 0xd0, 0xc1, 0x45,
		0x1e, 0x50, 0x3d, 0x57, 0xba, 0xf7, 0x1a, 0x2c, 0x9b, 0x26, 0x7f, 0x67, 0xcb, 0x34, 0x44, 0x8b,
		0xef, 0xab, 0x4a, 0xcc, 0x59, 0xa6, 0xb9, 0xcd, 0x15, 0x44, 0x8c, 0xfb, 0xe8, 0x45, 0x78, 0x7c,
		0xe4, 0xac, 0x28, 0x70, 0x69, 0xe2, 0x2d, 0xc7, 0xa1, 0xd7, 0x60, 0x79, 0x70, 0x3c, 0x09, 0x44,
		0xb1, 0x27, 0x0e, 0x8e, 0xc7, 0x61, 0x4f, 0xb3, 0x63, 0x9b, 0x47, 0x4c, 0x1c, 0x90, 0x8e, 0x7a,
		0x36, 0xaa, 0x1d, 0x99, 0x40, 0x97, 0x41, 0x31, 0x4d, 0x83, 0x38, 0xf8, 0xd0, 0x26, 0x06, 0xf6,
		0x88, 0x83, 0x7d, 0xf5, 0x42, 0x54, 0xb9, 0x6c, 0x9a, 0x75, 0x36, 0x5b, 0x65, 0x93, 0xe8, 0x59,
		0x58, 0x72, 0x0f, 0xef, 0x9a, 0x3c, 0xb2, 0x8c, 0x81, 0x47, 0xba, 0xd6, 0xab, 0xea, 0x53, 0xcc,
		0x4d, 0x8b, 0x74, 0x82, 0xc5, 0x55, 0x8b, 0x89, 0xd1, 0x33, 0xa0, 0x98, 0xfe, 0x11, 0xf6, 0x06,
		0xac, 0xb4, 0xfb, 0x03, 0x6c, 0x12, 0xf5, 0x69, 0xae, 0xca, 0xe5, 0x7b, 0x52, 0x4c, 0x23, 0xdb,
		0x7f, 0x60, 0x75, 0x03, 0xc9, 0x78, 0x89, 0x47, 0x36, 0x93, 0x09, 0xb6, 0x3b, 0xb0, 0x32, 0x74,
		0x2c, 0x27, 0x20, 0xde, 0xc0, 0x23, 0xb4, 0x89, 0xe7, 0x3b, 0x51, 0xfd, 0xe7, 0xf9, 0x13, 0xda,
		0xf0, 0x83, 0xa8, 0x36, 0x0f, 0x00, 0x7d, 0x79, 0x38, 0x29, 0x5c, 0xd7, 0xa0, 0x14, 0x8d, 0x0b,
		0x54, 0x00, 0x1e, 0x19, 0x4a, 0x8a, 0xd6, 0xd8, 0x5a, 0x73, 0x8b, 0x56, 0xc7, 0xcf, 0xd5, 0x95,
		0x34, 0xad, 0xd2, 0x3b, 0x8d, 0xfd, 0xba, 0xa1, 0x1f, 0xec, 0xed, 0x37, 0x76, 0xeb, 0x4a, 0xe6,
		0xd9, 0x42, 0xfe, 0x5f, 0xe6, 0x95, 0xd7, 0x5f, 0x7f, 0xfd, 0xf5, 0xf4, 0xfa, 0xb7, 0xd3, 0x50,
		0x8e, 0x77, 0xc6, 0xe8, 0x47, 0xe1, 0xac, 0x3c, 0xc6, 0xfa, 0x24, 0x30, 0x1e, 0x58, 0x1e, 0x0b,
		0xd5, 0x3e, 0xe6, 0xbd, 0x65, 0xe8, 0xe5, 0x15, 0xa1, 0xd5, 0x26, 0xc1, 0xcb, 0x96, 0x47, 0x03,
		0xb1, 0x8f, 0x03, 0xb4, 0x03, 0x17, 0x1c, 0xd7, 0xf0, 0x03, 0xec, 0x74, 0xb0, 0xd7, 0x31, 0x46,
		0x17, 0x08, 0x06, 0x36, 0x4d, 0xe2, 0xfb, 0x2e, 0x2f, 0x11, 0x21, 0xcb, 0x47, 0x1c, 0xb7, 0x2d,
		0x94, 0x47, 0xb9, 0xb3, 0x2a, 0x54, 0xc7, 0x22, 0x22, 0x73, 0x52, 0x44, 0x3c, 0x01, 0x85, 0x3e,
		0x1e, 0x18, 0xc4, 0x09, 0xbc, 0x63, 0xd6, 0xcf, 0xe5, 0xf5, 0x7c, 0x1f, 0x0f, 0xea, 0x74, 0xfc,
		0xc1, 0xad, 0x41, 0xd4, 0x8f, 0x7f, 0x9f, 0x81, 0x52, 0xb4, 0xa7, 0xa3, 0x2d, 0xb2, 0xc9, 0xf2,
		0x77, 0x8a, 0xed, 0xf0, 0x8f, 0x3e, 0xb2, 0x03, 0xac, 0xd4, 0x68, 0x62, 0xd7, 0xe6, 0x78, 0xa7,
		0xa5, 0x73, 0x24, 0x2d, 0xaa, 0x74, 0x4f, 0x13, 0xde, 0xbf, 0xe7, 0x75, 0x31, 0x42, 0xdb, 0x30,
		0x77, 0xd7, 0x67, 0xdc, 0x73, 0x8c, 0xfb, 0xa9, 0x47, 0x73, 0xdf, 0x6e, 0x33, 0xf2, 0xc2, 0xed,
		0xb6, 0xb1, 0xd7, 0xd4, 0x77, 0xab, 0x3b, 0xba, 0x80, 0xa3, 0x73, 0x90, 0xb5, 0xf1, 0x6b, 0xc7,
		0xf1, 0x12, 0xc0, 0x44, 0xb3, 0x3a, 0xfe, 0x1c, 0x64, 0x1f, 0x10, 0x7c, 0x2f, 0x9e, 0x78, 0x99,
		0xe8, 0x03, 0x0c, 0xfd, 0xcb, 0x90, 0x63, 0xfe, 0x42, 0x00, 0xc2, 0x63, 0xca, 0x63, 0x28, 0x0f,
		0xd9, 0x5a, 0x53, 0xa7, 0xe1, 0xaf, 0x40, 0x89, 0x4b, 0x8d, 0x56, 0xa3, 0x5e, 0xab, 0x2b, 0xe9,
		0xf5, 0x6b, 0x30, 0xc7, 0x9d, 0x40, 0xb7, 0x46, 0xe8, 0x06, 0xe5, 0x31, 0x31, 0x14, 0x1c, 0x29,
		0x39, 0x7b, 0xb0, 0xbb, 0x59, 0xd7, 0x95, 0x74, 0x74, 0x79, 0x7d, 0x28, 0x45, 0xdb, 0xb9, 0x1f,
		0x4e, 0x4c, 0xfd, 0x65, 0x0a, 0x8a, 0x91, 0xf6, 0x8c, 0x36, 0x06, 0xd8, 0xb6, 0xdd, 0x07, 0x06,
		0xb6, 0x2d, 0xec, 0x8b, 0xa0, 0x00, 0x26, 0xaa, 0x52, 0xc9, 0xac, 0x8b, 0xf6, 0x43, 0x31, 0xfe,
		0xed, 0x14, 0x28, 0xe3, 0xad, 0xdd, 0x98, 0x81, 0xa9, 0x0f, 0xd5, 0xc0, 0xb7, 0x52, 0x50, 0x8e,
		0xf7, 0x73, 0x63, 0xe6, 0x5d, 0xfc, 0x50, 0xcd, 0xfb, 0x87, 0x34, 0x2c, 0xc4, 0xba, 0xb8, 0x59,
		0xad, 0xfb, 0x02, 0x2c, 0x59, 0x1d, 0xd2, 0x1f, 0xb8, 0x01, 0x71, 0xcc, 0x63, 0xc3, 0x26, 0xf7,
		0x89, 0xad, 0xae, 0xb3, 0x44, 0x71, 0xf9, 0xd1, 0x7d, 0x62, 0xa5, 0x31, 0xc2, 0xed, 0x50, 0x98,
		0xb6, 0xdc, 0xd8, 0xaa, 0xef, 0xb6, 0x9a, 0xfb, 0xf5, 0xbd, 0xda, 0x2b, 0xc6, 0xc1, 0xde, 0x8f,
		0xed, 0x35, 0x5f, 0xde, 0xd3, 0x15, 0x6b, 0x4c, 0xed, 0x03, 0xdc, 0xea, 0x2d, 0x50, 0xc6, 0x8d,
		0x42, 0x67, 0x61, 0x9a, 0x59, 0xca, 0x63, 0x68, 0x19, 0x16, 0xf7, 0x9a, 0x46, 0xbb, 0xb1, 0x55,
		0x37, 0xea, 0xb7, 0x6e, 0xd5, 0x6b, 0xfb, 0x6d, 0x7e, 0x70, 0x0e, 0xb5, 0xf7, 0xe3, 0x9b, 0xfa,
		0xcd, 0x0c, 0x2c, 0x4f, 0xb1, 0x04, 0x55, 0x45, 0xcf, 0xce, 0x8f, 0x11, 0x1f, 0x9f, 0xc5, 0xfa,
		0x0a, 0xed, 0x0a, 0x5a, 0xd8, 0x0b, 0x44, 0x8b, 0xff, 0x0c, 0x50, 0x2f, 0x39, 0x81, 0xd5, 0xb5,
		0x88, 0x27, 0xee, 0x19, 0x78, 0x23, 0xbf, 0x38, 0x92, 0xf3, 0xab, 0x86, 0x1f, 0x01, 0x34, 0x70,
		0x7d, 0x2b, 0xb0, 0xee, 0x13, 0xc3, 0x72, 0xe4, 0xa5, 0x04, 0x6d, 0xec, 0xb3, 0xba, 0x22, 0x67,
		0x1a, 0x4e, 0x10, 0x6a, 0x3b, 0xa4, 0x87, 0xc7, 0xb4, 0x69, 0x02, 0xcf, 0xe8, 0x8a, 0x9c, 0x09,
		0xb5, 0x2f, 0x42, 0xa9, 0xe3, 0x0e, 0x69, 0x9b, 0xc4, 0xf5, 0x68, 0xbd, 0x48, 0xe9, 0x45, 0x2e,
		0x0b, 0x55, 0x44, 0x1f, 0x3b, 0xba, 0x0d, 0x29, 0xe9, 0x45, 0x2e, 0xe3, 0x2a, 0x97, 0x60, 0x11,
		0xf7, 0x7a, 0x1e, 0x25, 0x97, 0x44, 0xbc, 0x33, 0x2f, 0x87, 0x62, 0xa6, 0xb8, 0x7a, 0x1b, 0xf2,
		0xd2, 0x0f, 0xb4, 0x24, 0x53, 0x4f, 0x18, 0x03, 0x7e, 0x27, 0x95, 0xde, 0x28, 0xe8, 0x79, 0x47,
		0x4e, 0x5e, 0x84, 0x92, 0xe5, 0x1b, 0xa3, 0xcb, 0xd1, 0xf4, 0x5a, 0x7a, 0x23, 0xaf, 0x17, 0x2d,
		0x3f, 0xbc, 0x0d, 0x5b, 0x7f, 0x27, 0x0d, 0xe5, 0xf8, 0xe5, 0x2e, 0xda, 0x82, 0xbc, 0xed, 0x9a,
		0x98, 0x85, 0x16, 0xff, 0xb2, 0xb0, 0x91, 0x70, 0x1f, 0x5c, 0xd9, 0x11, 0xfa, 0x7a, 0x88, 0x5c,
		0xfd, 0xbb, 0x14, 0xe4, 0xa5, 0x18, 0x9d, 0x81, 0xec, 0x00, 0x07, 0x47, 0x8c, 0x2e, 0xb7, 0x99,
		0x56, 0x52, 0x3a, 0x1b, 0x53, 0xb9, 0x3f, 0xc0, 0x0e, 0x0b, 0x01, 0x21, 0xa7, 0x63, 0xba, 0xae,
		0x36, 0xc1, 0x1d, 0xd6, 0xf6, 0xbb, 0xfd, 0x3e, 0x71, 0x02, 0x5f, 0xae, 0xab, 0x90, 0xd7, 0x84,
		0x18, 0x3d, 0x07, 0x4b, 0x81, 0x87, 0x2d, 0x3b, 0xa6, 0x9b, 0x65, 0xba, 0x8a, 0x9c, 0x08, 0x95,
		0x35, 0x38, 0x27, 0x79, 0x3b, 0x24, 0xc0, 0xe6, 0x11, 0xe9, 0x8c, 0x40, 0x73, 0xec, 0xe6, 0xf0,
		0xac, 0x50, 0xd8, 0x12, 0xf3, 0x12, 0xbb, 0xfe, 0x9d, 0x14, 0x2c, 0xc9, 0x83, 0x4a, 0x27, 0x74,
		0xd6, 0x2e, 0x00, 0x76, 0x1c, 0x37, 0x88, 0xba, 0x6b, 0x32, 0x94, 0x27, 0x70, 0x95, 0x6a, 0x08,
		0xd2, 0x23, 0x04, 0xab, 0x7d, 0x80, 0xd1, 0xcc, 0x89, 0x6e, 0xbb, 0x00, 0x45, 0x71, 0x73, 0xcf,
		0x3e, 0xff, 0xf0, 0xa3, 0x2d, 0x70, 0x11, 0x3d, 0xd1, 0xa0, 0x15, 0xc8, 0x1d, 0x92, 0x9e, 0xe5,
		0x88, 0xfb, 0x44, 0x3e, 0x90, 0xb7, 0x94, 0xd9, 0xf0, 0x96, 0x72, 0xf3, 0x0e, 0x2c, 0x9b, 0x6e,
		0x7f, 0xdc, 0xdc, 0x4d, 0x65, 0xec, 0x78, 0xed, 0x7f, 0x3a, 0xf5, 0x39, 0x18, 0xb5, 0x98, 0x5f,
		0x4d, 0x67, 0xb6, 0x5b, 0x9b, 0x5f, 0x4b, 0xaf, 0x6e, 0x73, 0x5c, 0x4b, 0xbe, 0xa6, 0x4e, 0xba,
		0x36, 0x31, 0xa9, 0xe9, 0xf0, 0xfd, 0x8f, 0xc1, 0xc7, 0x7b, 0x56, 0x70, 0x34, 0x3c, 0xac, 0x98,
		0x6e, 0xff, 0x72, 0xcf, 0xed, 0xb9, 0xa3, 0xcf, 0x5d, 0x74, 0xc4, 0x06, 0xec, 0x3f, 0xf1, 0xc9,
		0xab, 0x10, 0x4a, 0x57, 0x13, 0xbf, 0x8f, 0x69, 0x7b, 0xb0, 0x2c, 0x94, 0x0d, 0x76, 0xe7, 0xce,
		0x8f, 0x06, 0xe8, 0x91, 0xf7, 0x2e, 0xea, 0x37, 0xdf, 0x65, 0xb5, 0x5a, 0x5f, 0x12, 0x50, 0x3a,
		0xc7, 0x0f, 0x10, 0x9a, 0x0e, 0x8f, 0xc7, 0xf8, 0xf8, 0xbe, 0x24, 0x5e, 0x02, 0xe3, 0xb7, 0x05,
		0xe3, 0x72, 0x84, 0xb1, 0x2d, 0xa0, 0x5a, 0x0d, 0x16, 0x4e, 0xc3, 0xf5, 0xd7, 0x82, 0xab, 0x44,
		0xa2, 0x24, 0xdb, 0xb0, 0xc8, 0x48, 0xcc, 0xa1, 0x1f, 0xb8, 0x7d, 0x96, 0xf4, 0x1e, 0x4d, 0xf3,
		0x37, 0xef, 0xf2, 0x8d, 0x52, 0xa6, 0xb0, 0x5a, 0x88, 0xd2, 0x34, 0x60, 0x9f, 0x19, 0x3a, 0xc4,
		0xb4, 0x13, 0x18, 0xfe, 0x56, 0x18, 0x12, 0xea, 0x6b, 0x9f, 0x85, 0x15, 0xfa, 0x3f, 0xcb, 0x49,
		0x51, 0x4b, 0x92, 0x6f, 0x99, 0xd4, 0xef, 0xbc, 0xc1, 0xf7, 0xe2, 0x72, 0x48, 0x10, 0xb1, 0x29,
		0xb2, 0x8a, 0x3d, 0x12, 0x04, 0xc4, 0xf3, 0x0d, 0x6c, 0x4f, 0x33, 0x2f, 0x72, 0x4c, 0x57, 0xbf,
		0xfc, 0x5e, 0x7c, 0x15, 0xb7, 0x39, 0xb2, 0x6a, 0xdb, 0xda, 0x01, 0x9c, 0x9d, 0x12, 0x15, 0x33,
		0x70, 0xbe, 0x29, 0x38, 0x57, 0x26, 0x22, 0x83, 0xd2, 0xb6, 0x40, 0xca, 0xc3, 0xb5, 0x9c, 0x81,
		0xf3, 0xd7, 0x05, 0x27, 0x12, 0x58, 0xb9, 0xa4, 0x94, 0xf1, 0x36, 0x2c, 0xdd, 0x27, 0xde, 0xa1,
		0xeb, 0x8b, 0xab, 0x91, 0x19, 0xe8, 0xde, 0x12, 0x74, 0x8b, 0x02, 0xc8, 0xee, 0x4a, 0x28, 0xd7,
		0x8b, 0x90, 0xef, 0x62, 0x93, 0xcc, 0x40, 0xf1, 0x15, 0x41, 0x31, 0x4f, 0xf5, 0x29, 0xb4, 0x0a,
		0xa5, 0x9e, 0x2b, 0xca, 0x52, 0x32, 0xfc, 0x6d, 0x01, 0x2f, 0x4a, 0x8c, 0xa0, 0x18, 0xb8, 0x83,
		0xa1, 0x4d, 0x6b, 0x56, 0x32, 0xc5, 0x6f, 0x48, 0x0a, 0x89, 0x11, 0x14, 0xa7, 0x70, 0xeb, 0x6f,
		0x4a, 0x0a, 0x3f, 0xe2, 0xcf, 0x97, 0xa0, 0xe8, 0x3a, 0xf6, 0xb1, 0xeb, 0xcc, 0x62, 0xc4, 0x6f,
		0x09, 0x06, 0x10, 0x10, 0x4a, 0x70, 0x13, 0x0a, 0xb3, 0x2e, 0xc4, 0x6f, 0xbf, 0x27, 0xb7, 0x87,
		0x5c, 0x81, 0x6d, 0x58, 0x94, 0x09, 0xca, 0x72, 0x9d, 0x19, 0x28, 0x7e, 0x47, 0x50, 0x94, 0x23,
		0x30, 0xf1, 0x1a, 0x01, 0xf1, 0x83, 0x1e, 0x99, 0x85, 0xe4, 0x1d, 0xf9, 0x1a, 0x02, 0x22, 0x5c,
		0x79, 0x48, 0x1c, 0xf3, 0x68, 0x36, 0x86, 0xdf, 0x95, 0xae, 0x94, 0x18, 0x4a, 0x51, 0x83, 0x85,
		0x3e, 0xf6, 0xfc, 0x23, 0x6c, 0xcf, 0xb4, 0x1c, 0xbf, 0x27, 0x38, 0x4a, 0x21, 0x48, 0x78, 0x64,
		0xe8, 0x9c, 0x86, 0xe6, 0x6b, 0xd2, 0x23, 0x11, 0x98, 0xd8, 0x7a, 0x7e, 0xc0, 0x2e, 0xa0, 0x4e,
		0xc3, 0xf6, 0x75, 0xb9, 0xf5, 0x38, 0x76, 0x37, 0xca, 0x78, 0x13, 0x0a, 0xbe, 0xf5, 0xda, 0x4c,
		0x34, 0xbf, 0x2f, 0x57, 0x9a, 0x01, 0x28, 0xf8, 0x15, 0x38, 0x37, 0xb5, 0x4c, 0xcc, 0x40, 0xf6,
		0x07, 0x82, 0xec, 0xcc, 0x94, 0x52, 0x21, 0x52, 0xc2, 0x69, 0x29, 0xff, 0x50, 0xa6, 0x04, 0x32,
		0xc6, 0xd5, 0xa2, 0x07, 0x05, 0x1f, 0x77, 0x4f, 0xe7, 0xb5, 0x3f, 0x92, 0x5e, 0xe3, 0xd8, 0x98,
		0xd7, 0xf6, 0xe1, 0x8c, 0x60, 0x3c, 0xdd, 0xba, 0x7e, 0x43, 0x26, 0x56, 0x8e, 0x3e, 0x88, 0xaf,
		0xee, 0xe7, 0x61, 0x35, 0x74, 0xa7, 0xec, 0x48, 0x7d, 0xa3, 0x8f, 0x07, 0x33, 0x30, 0x7f, 0x53,
		0x30, 0xcb, 0x8c, 0x1f, 0xb6, 0xb4, 0xfe, 0x2e, 0x1e, 0x50, 0xf2, 0x3b, 0xa0, 0x4a, 0xf2, 0xa1,
		0xe3, 0x11, 0xd3, 0xed, 0x39, 0xd6, 0x6b, 0xa4, 0x33, 0x03, 0xf5, 0x1f, 0x8f, 0x2d, 0xd5, 0x41,
		0x04, 0x4e, 0x99, 0x1b, 0xa0, 0x84, 0xbd, 0x8a, 0x61, 0xf5, 0x07, 0xae, 0x17, 0x24, 0x30, 0xfe,
		0x89, 0x5c, 0xa9, 0x10, 0xd7, 0x60, 0x30, 0xad, 0x0e, 0x65, 0x36, 0x9c, 0x35, 0x24, 0xff, 0x54,
		0x10, 0x2d, 0x8c, 0x50, 0x22, 0x71, 0x98, 0x6e, 0x7f, 0x80, 0xbd, 0x59, 0xf2, 0xdf, 0x9f, 0xc9,
		0xc4, 0x21, 0x20, 0x22, 0x71, 0x04, 0xc7, 0x03, 0x42, 0xab, 0xfd, 0x0c, 0x0c, 0xdf, 0x92, 0x89,
		0x43, 0x62, 0x04, 0x85, 0x6c, 0x18, 0x66, 0xa0, 0xf8, 0x73, 0x49, 0x21, 0x31, 0x94, 0xe2, 0x33,
		0xa3, 0x42, 0xeb, 0x91, 0x9e, 0xe5, 0x07, 0x1e, 0xef, 0x83, 0x1f, 0x4d, 0xf5, 0x17, 0xef, 0xc5,
		0x9b, 0x30, 0x3d, 0x02, 0xd5, 0x6e, 0xc3, 0xe2, 0x58, 0x8b, 0x81, 0x92, 0x7e, 0xb3, 0xa0, 0xfe,
		0xf8, 0xfb, 0x22, 0x19, 0xc5, 0x3b, 0x0c, 0x6d, 0x87, 0xae, 0x7b, 0xbc, 0x0f, 0x48, 0x26, 0x7b,
		0xe3, 0xfd, 0x70, 0xe9, 0x63, 0x6d, 0x80, 0x76, 0x0b, 0x16, 0x62, 0x3d, 0x40, 0x32, 0xd5, 0x4f,
		0x08, 0xaa, 0x52, 0xb4, 0x05, 0xd0, 0xae, 0x41, 0x96, 0xd6, 0xf3, 0x64, 0xf8, 0x4f, 0x0a, 0x38,
		0x53, 0xd7, 0x3e, 0x09, 0x79, 0x59, 0xc7, 0x93, 0xa1, 0x3f, 0x25, 0xa0, 0x21, 0x84, 0xc2, 0x65,
		0x0d, 0x4f, 0x86, 0xff, 0xb4, 0x84, 0x4b, 0x08, 0x85, 0xcf, 0xee, 0xc2, 0xbf, 0xfa, 0x99, 0xac,
		0xc8, 0xc3, 0xd2, 0x77, 0x37, 0x61, 0x5e, 0x14, 0xef, 0x64, 0xf4, 0x17, 0xc5, 0xc3, 0x25, 0x42,
		0x7b, 0x01, 0x72, 0x33, 0x3a, 0xfc, 0x67, 0x05, 0x94, 0xeb, 0x6b, 0x35, 0x28, 0x46, 0x0a, 0x76,
		0x32, 0xfc, 0xe7, 0x04, 0x3c, 0x8a, 0xa2, 0xa6, 0x8b, 0x82, 0x9d, 0x4c, 0xf0, 0xf3, 0xd2, 0x74,
		0x81, 0xa0, 0x6e, 0x93, 0xb5, 0x3a, 0x19, 0xfd, 0x0b, 0xd2, 0xeb, 0x12, 0xa2, 0xbd, 0x04, 0x85,
		0x30, 0xff, 0x26, 0xe3, 0x7f, 0x51, 0xe0, 0x47, 0x18, 0xea, 0x81, 0x48, 0xfe, 0x4f, 0xa6, 0xf8,
		0x25, 0xe9, 0x81, 0x08, 0x8a, 0x6e, 0xa3, 0xf1, 0x9a, 0x9e, 0xcc, 0xf4, 0xcb, 0x72, 0x1b, 0x8d,
		0x95, 0x74, 0xba, 0x9a, 0x2c, 0x0d, 0x26, 0x53, 0xfc, 0x8a, 0x5c, 0x4d, 0xa6, 0x4f, 0xcd, 0x18,
		0x2f, 0x92, 0xc9, 0x1c, 0xbf, 0x26, 0xcd, 0x18, 0xab, 0x91, 0x5a, 0x0b, 0xd0, 0x64, 0x81, 0x4c,
		0xe6, 0xfb, 0x92, 0xe0, 0x5b, 0x9a, 0xa8, 0x8f, 0xda, 0xcb, 0x70, 0x66, 0x7a, 0x71, 0x4c, 0x66,
		0xfd, 0xf2, 0xfb, 0x63, 0xc7, 0x99, 0x68, 0x6d, 0xd4, 0xf6, 0x47, 0x59, 0x36, 0x5a, 0x18, 0x93,
		0x69, 0xdf, 0x7c, 0x3f, 0x9e, 0x68, 0xa3, 0x75, 0x51, 0xab, 0x02, 0x8c, 0x6a, 0x52, 0x32, 0xd7,
		0x5b, 0x82, 0x2b, 0x02, 0xa2, 0x5b, 0x43, 0x94, 0xa4, 0x64, 0xfc, 0x57, 0xe4, 0xd6, 0x10, 0x08,
		0xba, 0x35, 0x64, 0x35, 0x4a, 0x46, 0xbf, 0x2d, 0xb7, 0x86, 0x84, 0x68, 0x37, 0x21, 0xef, 0x0c,
		0x6d, 0x9b, 0xc6, 0x16, 0x7a, 0xf4, 0xcf, 0x88, 0xd4, 0x7f, 0x7d, 0x28, 0xc0, 0x12, 0xa0, 0x5d,
		0x83, 0x1c, 0xe9, 0x1f, 0x92, 0x4e, 0x12, 0xf2, 0xdf, 0x1e, 0xca, 0x7c, 0x42, 0xb5, 0xb5, 0x97,
		0x00, 0xf8, 0x61, 0x9a, 0x7d, 0x25, 0x4a, 0xc0, 0xfe, 0xfb, 0x43, 0xf1, 0x0b, 0x85, 0x11, 0x64,
		0x44, 0xc0, 0x7f, 0xef, 0xf0, 0x68, 0x82, 0xf7, 0xe2, 0x04, 0xec, 0x00, 0xfe, 0x22, 0xcc, 0xdf,
		0xf5, 0x5d, 0x27, 0xc0, 0xbd, 0x24, 0xf4, 0x7f, 0x08, 0xb4, 0xd4, 0xa7, 0x0e, 0xeb, 0xbb, 0x1e,
		0x09, 0x70, 0xcf, 0x4f, 0xc2, 0xfe, 0xa7, 0xc0, 0x86, 0x00, 0x0a, 0x36, 0xb1, 0x1f, 0xcc, 0xf2,
		0xde, 0xff, 0x25, 0xc1, 0x12, 0x40, 0x8d, 0xa6, 0xff, 0xdf, 0x23, 0xc7, 0x49, 0xd8, 0xef, 0x49,
		0xa3, 0x85, 0xbe, 0xf6, 0x49, 0x28, 0xd0, 0x7f, 0xf9, 0xaf, 0x76, 0x12, 0xc0, 0xff, 0x2d, 0xc0,
		0x23, 0x04, 0x7d, 0xb2, 0x1f, 0x74, 0x02, 0x2b, 0xd9, 0xd9, 0xff, 0x23, 0x56, 0x5a, 0xea, 0x6b,
		0x55, 0x28, 0xfa, 0x41, 0xa7, 0x33, 0x14, 0x1d, 0x4d, 0x02, 0xfc, 0xfb, 0x0f, 0xc3, 0x43, 0x6e,
		0x88, 0xd9, 0xbc, 0x38, 0xfd, 0xb2, 0x0e, 0xb6, 0xdd, 0x6d, 0x97, 0x5f, 0xd3, 0xc1, 0xaf, 0xa6,
		0xa1, 0xc4, 0x26, 0xae, 0x8a, 0x4b, 0x35, 0xbe, 0xbd, 0x68, 0xf9, 0xf0, 0x57, 0x4f, 0x77, 0x1f,
		0xb7, 0xfe, 0x79, 0x98, 0x63, 0xac, 0x57, 0xd1, 0x0a, 0xe4, 0x98, 0x75, 0xec, 0x4b, 0x52, 0x46,
		0xe7, 0x03, 0x74, 0x09, 0x32, 0xbb, 0x7e, 0x4f, 0xfc, 0x86, 0xe7, 0xf1, 0xca, 0xe8, 0x41, 0x95,
		0x76, 0x1f, 0xdb, 0xf6, 0xae, 0xdf, 0xbb, 0xaa, 0x53, 0x0d, 0x84, 0x20, 0xdb, 0xe0, 0x97, 0xb2,
		0x99, 0x8d, 0x8c, 0xce, 0xfe, 0x5f, 0x7f, 0x19, 0x0a, 0xa1, 0x16, 0x5a, 0x83, 0x62, 0xdb, 0xc4,
		0xde, 0xe6, 0xd0, 0x6f, 0x07, 0xee, 0x40, 0xfe, 0x72, 0x25, 0x22, 0x42, 0x1b, 0xb0, 0x78, 0xcb,
		0xb6, 0x7a, 0x47, 0x41, 0x0b, 0x7b, 0xd8, 0x3c, 0x1a, 0x06, 0x44, 0x2d, 0xad, 0x65, 0x36, 0xe6,
		0xf5, 0x71, 0xf1, 0x66, 0xe9, 0x7b, 0xdf, 0x3d, 0x9f, 0xfa, 0xbf, 0xef, 0x9e, 0x4f, 0x7d, 0xe3,
		0x1f, 0xcf, 0xa7, 0x0e, 0xe7, 0xb8, 0x33, 0xfe, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xd4, 0xf5,
		0x46, 0x74, 0x2f, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *Proto3) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.Proto3{")
	s = append(s, "Field: "+fmt.Sprintf("%#v", this.Field)+",\n")
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	s = append(s, "Ints: "+fmt.Sprintf("%#v", this.Ints)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg3) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.SmallMsg3{")
	s = append(s, "ScarBusStop: "+fmt.Sprintf("%#v", this.ScarBusStop)+",\n")
	s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProto3(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedProto3(r randyProto3, easy bool) *Proto3 {
	this := &Proto3{}
	this.Field = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field *= -1
	}
	if r.Intn(10) != 0 {
		this.Msg = NewPopulatedSmallMsg3(r, easy)
	}
	v1 := r.Intn(10)
	this.Ints = make([]int64, v1)
	for i := 0; i < v1; i++ {
		this.Ints[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Ints[i] *= -1
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSmallMsg3(r randyProto3, easy bool) *SmallMsg3 {
	this := &SmallMsg3{}
	this.ScarBusStop = string(randStringProto3(r))
	v2 := r.Intn(10)
	this.FlightParachute = make([]uint32, v2)
	for i := 0; i < v2; i++ {
		this.FlightParachute[i] = uint32(r.Uint32())
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyProto3 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProto3(r randyProto3) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProto3(r randyProto3) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneProto3(r)
	}
	return string(tmps)
}
func randUnrecognizedProto3(r randyProto3, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProto3(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProto3(dAtA []byte, r randyProto3, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProto3(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProto3(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptorProto3) }

var fileDescriptorProto3 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd6, 0x03, 0x53, 0x42, 0x5c, 0x60, 0xaa, 0x24, 0xb5, 0xb8, 0xa4, 0x58, 0x4a, 0x37,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3d, 0x3f, 0x3d, 0x5f, 0x1f,
	0x2c, 0x97, 0x54, 0x9a, 0x06, 0xe6, 0x81, 0x39, 0x60, 0x16, 0x44, 0xab, 0x52, 0x34, 0x17, 0x5b,
	0x00, 0xd8, 0x28, 0x21, 0x11, 0x2e, 0x56, 0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xe6, 0x20, 0x08, 0x47, 0x48, 0x9d, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x82, 0x59, 0x81, 0x51,
	0x83, 0xdb, 0x48, 0x54, 0x0f, 0x61, 0x91, 0x5e, 0x70, 0x6e, 0x62, 0x4e, 0x8e, 0x6f, 0x71, 0xba,
	0x71, 0x10, 0x48, 0x85, 0x90, 0x10, 0x17, 0x8b, 0x67, 0x5e, 0x49, 0xb1, 0x04, 0x8b, 0x02, 0xb3,
	0x06, 0x73, 0x10, 0x98, 0xad, 0x14, 0xce, 0xc5, 0x09, 0x57, 0x25, 0xa4, 0xc0, 0xc5, 0x1d, 0x9c,
	0x9c, 0x58, 0xe4, 0x54, 0x5a, 0x1c, 0x5c, 0x92, 0x5f, 0x00, 0xb6, 0x85, 0x33, 0x08, 0x59, 0x48,
	0x48, 0x83, 0x8b, 0xdf, 0x2d, 0x27, 0x33, 0x3d, 0xa3, 0x24, 0x20, 0xb1, 0x28, 0x31, 0x39, 0xa3,
	0xb4, 0x24, 0x55, 0x82, 0x47, 0x81, 0x59, 0x83, 0x3d, 0x08, 0x5d, 0xd8, 0x89, 0xe7, 0xc3, 0x43,
	0x39, 0xc6, 0x1f, 0x0f, 0xe5, 0x18, 0x37, 0x3c, 0x92, 0x63, 0x4c, 0x62, 0x83, 0x04, 0x06, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xfb, 0x93, 0xeb, 0x15, 0x01, 0x00, 0x00,
}
