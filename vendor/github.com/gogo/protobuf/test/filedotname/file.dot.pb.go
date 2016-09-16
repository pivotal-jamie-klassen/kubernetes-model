// Code generated by protoc-gen-gogo.
// source: file.dot.proto
// DO NOT EDIT!

/*
Package filedotname is a generated protocol buffer package.

It is generated from these files:
	file.dot.proto

It has these top-level messages:
	M
*/
package filedotname

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
import sort "sort"
import strconv "strconv"
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

type M struct {
	A                *string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *M) Reset()                    { *m = M{} }
func (*M) ProtoMessage()               {}
func (*M) Descriptor() ([]byte, []int) { return fileDescriptorFileDot, []int{0} }

func init() {
	proto.RegisterType((*M)(nil), "filedotname.M")
}
func (this *M) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return FileDotDescription()
}
func FileDotDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3426 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x5d, 0x6c, 0x1c, 0xd5,
		0x15, 0xf6, 0xec, 0x8f, 0xbd, 0x7b, 0x76, 0xbd, 0x1e, 0x8f, 0x8d, 0xb3, 0x31, 0xb0, 0x71, 0x0c,
		0x14, 0x03, 0xc5, 0x41, 0x21, 0x09, 0xc9, 0xa6, 0x10, 0xad, 0xed, 0x8d, 0x71, 0xe4, 0xbf, 0xce,
		0xda, 0x10, 0xe0, 0x61, 0x74, 0x3d, 0x7b, 0x77, 0x3d, 0xc9, 0xec, 0xcc, 0x76, 0x66, 0x36, 0x89,
		0xf3, 0x14, 0x44, 0x7f, 0x84, 0x50, 0xff, 0x2b, 0x95, 0xff, 0x02, 0x52, 0x0b, 0xa5, 0xa5, 0x85,
		0xfe, 0xa9, 0xea, 0x13, 0x2f, 0xb4, 0x3c, 0x55, 0xf0, 0xd6, 0x87, 0x3e, 0x10, 0x17, 0xa9, 0xb4,
		0xa5, 0x2d, 0x95, 0x22, 0x15, 0x89, 0x97, 0xea, 0xfe, 0xcd, 0xce, 0xfe, 0x38, 0xb3, 0x46, 0x02,
		0xfa, 0xe4, 0xbd, 0xe7, 0x9e, 0xef, 0x9b, 0x73, 0xcf, 0x3d, 0xf7, 0x9c, 0x33, 0x77, 0x0c, 0x0f,
		0x1f, 0x82, 0x89, 0xaa, 0x6d, 0x57, 0x4d, 0x7c, 0xa0, 0xee, 0xd8, 0x9e, 0xbd, 0xd1, 0xa8, 0x1c,
		0x28, 0x63, 0x57, 0x77, 0x8c, 0xba, 0x67, 0x3b, 0xd3, 0x54, 0xa6, 0x0c, 0x31, 0x8d, 0x69, 0xa1,
		0x31, 0xb9, 0x04, 0xc3, 0x27, 0x0d, 0x13, 0xcf, 0xf9, 0x8a, 0x25, 0xec, 0x29, 0x47, 0x21, 0x56,
		0x31, 0x4c, 0x9c, 0x95, 0x26, 0xa2, 0x53, 0xa9, 0x83, 0x37, 0x4e, 0xb7, 0x81, 0xa6, 0x5b, 0x11,
		0xab, 0x44, 0xac, 0x52, 0xc4, 0xe4, 0xbb, 0x31, 0x18, 0xe9, 0x32, 0xab, 0x28, 0x10, 0xb3, 0x50,
		0x8d, 0x30, 0x4a, 0x53, 0x49, 0x95, 0xfe, 0x56, 0xb2, 0x30, 0x50, 0x47, 0xfa, 0x59, 0x54, 0xc5,
		0xd9, 0x08, 0x15, 0x8b, 0xa1, 0x92, 0x03, 0x28, 0xe3, 0x3a, 0xb6, 0xca, 0xd8, 0xd2, 0xb7, 0xb2,
		0xd1, 0x89, 0xe8, 0x54, 0x52, 0x0d, 0x48, 0x94, 0xdb, 0x60, 0xb8, 0xde, 0xd8, 0x30, 0x0d, 0x5d,
		0x0b, 0xa8, 0xc1, 0x44, 0x74, 0x2a, 0xae, 0xca, 0x6c, 0x62, 0xae, 0xa9, 0x7c, 0x33, 0x0c, 0x9d,
		0xc7, 0xe8, 0x6c, 0x50, 0x35, 0x45, 0x55, 0x33, 0x44, 0x1c, 0x50, 0x9c, 0x85, 0x74, 0x0d, 0xbb,
		0x2e, 0xaa, 0x62, 0xcd, 0xdb, 0xaa, 0xe3, 0x6c, 0x8c, 0xae, 0x7e, 0xa2, 0x63, 0xf5, 0xed, 0x2b,
		0x4f, 0x71, 0xd4, 0xda, 0x56, 0x1d, 0x2b, 0x05, 0x48, 0x62, 0xab, 0x51, 0x63, 0x0c, 0xf1, 0x1d,
		0xfc, 0x57, 0xb4, 0x1a, 0xb5, 0x76, 0x96, 0x04, 0x81, 0x71, 0x8a, 0x01, 0x17, 0x3b, 0xe7, 0x0c,
		0x1d, 0x67, 0xfb, 0x29, 0xc1, 0xcd, 0x1d, 0x04, 0x25, 0x36, 0xdf, 0xce, 0x21, 0x70, 0xca, 0x2c,
		0x24, 0xf1, 0x05, 0x0f, 0x5b, 0xae, 0x61, 0x5b, 0xd9, 0x01, 0x4a, 0x72, 0x53, 0x97, 0x5d, 0xc4,
		0x66, 0xb9, 0x9d, 0xa2, 0x89, 0x53, 0x8e, 0xc0, 0x80, 0x5d, 0xf7, 0x0c, 0xdb, 0x72, 0xb3, 0x89,
		0x09, 0x69, 0x2a, 0x75, 0xf0, 0xba, 0xae, 0x81, 0xb0, 0xc2, 0x74, 0x54, 0xa1, 0xac, 0x2c, 0x80,
		0xec, 0xda, 0x0d, 0x47, 0xc7, 0x9a, 0x6e, 0x97, 0xb1, 0x66, 0x58, 0x15, 0x3b, 0x9b, 0xa4, 0x04,
		0xfb, 0x3a, 0x17, 0x42, 0x15, 0x67, 0xed, 0x32, 0x5e, 0xb0, 0x2a, 0xb6, 0x9a, 0x71, 0x5b, 0xc6,
		0xca, 0x18, 0xf4, 0xbb, 0x5b, 0x96, 0x87, 0x2e, 0x64, 0xd3, 0x34, 0x42, 0xf8, 0x68, 0xf2, 0xbf,
		0x71, 0x18, 0xea, 0x25, 0xc4, 0x8e, 0x43, 0xbc, 0x42, 0x56, 0x99, 0x8d, 0xec, 0xc6, 0x07, 0x0c,
		0xd3, 0xea, 0xc4, 0xfe, 0x8f, 0xe9, 0xc4, 0x02, 0xa4, 0x2c, 0xec, 0x7a, 0xb8, 0xcc, 0x22, 0x22,
		0xda, 0x63, 0x4c, 0x01, 0x03, 0x75, 0x86, 0x54, 0xec, 0x63, 0x85, 0xd4, 0x69, 0x18, 0xf2, 0x4d,
		0xd2, 0x1c, 0x64, 0x55, 0x45, 0x6c, 0x1e, 0x08, 0xb3, 0x64, 0xba, 0x28, 0x70, 0x2a, 0x81, 0xa9,
		0x19, 0xdc, 0x32, 0x56, 0xe6, 0x00, 0x6c, 0x0b, 0xdb, 0x15, 0xad, 0x8c, 0x75, 0x33, 0x9b, 0xd8,
		0xc1, 0x4b, 0x2b, 0x44, 0xa5, 0xc3, 0x4b, 0x36, 0x93, 0xea, 0xa6, 0x72, 0xac, 0x19, 0x6a, 0x03,
		0x3b, 0x44, 0xca, 0x12, 0x3b, 0x64, 0x1d, 0xd1, 0xb6, 0x0e, 0x19, 0x07, 0x93, 0xb8, 0xc7, 0x65,
		0xbe, 0xb2, 0x24, 0x35, 0x62, 0x3a, 0x74, 0x65, 0x2a, 0x87, 0xb1, 0x85, 0x0d, 0x3a, 0xc1, 0xa1,
		0x72, 0x03, 0xf8, 0x02, 0x8d, 0x86, 0x15, 0xd0, 0x2c, 0x94, 0x16, 0xc2, 0x65, 0x54, 0xc3, 0xe3,
		0x47, 0x21, 0xd3, 0xea, 0x1e, 0x65, 0x14, 0xe2, 0xae, 0x87, 0x1c, 0x8f, 0x46, 0x61, 0x5c, 0x65,
		0x03, 0x45, 0x86, 0x28, 0xb6, 0xca, 0x34, 0xcb, 0xc5, 0x55, 0xf2, 0x73, 0xfc, 0x2e, 0x18, 0x6c,
		0x79, 0x7c, 0xaf, 0xc0, 0xc9, 0xc7, 0xfb, 0x61, 0xb4, 0x5b, 0xcc, 0x75, 0x0d, 0xff, 0x31, 0xe8,
		0xb7, 0x1a, 0xb5, 0x0d, 0xec, 0x64, 0xa3, 0x94, 0x81, 0x8f, 0x94, 0x02, 0xc4, 0x4d, 0xb4, 0x81,
		0xcd, 0x6c, 0x6c, 0x42, 0x9a, 0xca, 0x1c, 0xbc, 0xad, 0xa7, 0xa8, 0x9e, 0x5e, 0x24, 0x10, 0x95,
		0x21, 0x95, 0x7b, 0x20, 0xc6, 0x53, 0x1c, 0x61, 0xb8, 0xb5, 0x37, 0x06, 0x12, 0x8b, 0x2a, 0xc5,
		0x29, 0xd7, 0x42, 0x92, 0xfc, 0x65, 0xbe, 0xed, 0xa7, 0x36, 0x27, 0x88, 0x80, 0xf8, 0x55, 0x19,
		0x87, 0x04, 0x0d, 0xb3, 0x32, 0x16, 0xa5, 0xc1, 0x1f, 0x93, 0x8d, 0x29, 0xe3, 0x0a, 0x6a, 0x98,
		0x9e, 0x76, 0x0e, 0x99, 0x0d, 0x4c, 0x03, 0x26, 0xa9, 0xa6, 0xb9, 0xf0, 0x3e, 0x22, 0x53, 0xf6,
		0x41, 0x8a, 0x45, 0xa5, 0x61, 0x95, 0xf1, 0x05, 0x9a, 0x7d, 0xe2, 0x2a, 0x0b, 0xd4, 0x05, 0x22,
		0x21, 0x8f, 0x3f, 0xe3, 0xda, 0x96, 0xd8, 0x5a, 0xfa, 0x08, 0x22, 0xa0, 0x8f, 0xbf, 0xab, 0x3d,
		0xf1, 0x5d, 0xdf, 0x7d, 0x79, 0xed, 0xb1, 0x38, 0xf9, 0x9b, 0x08, 0xc4, 0xe8, 0x79, 0x1b, 0x82,
		0xd4, 0xda, 0x03, 0xab, 0x45, 0x6d, 0x6e, 0x65, 0x7d, 0x66, 0xb1, 0x28, 0x4b, 0x4a, 0x06, 0x80,
		0x0a, 0x4e, 0x2e, 0xae, 0x14, 0xd6, 0xe4, 0x88, 0x3f, 0x5e, 0x58, 0x5e, 0x3b, 0x72, 0x48, 0x8e,
		0xfa, 0x80, 0x75, 0x26, 0x88, 0x05, 0x15, 0xee, 0x3c, 0x28, 0xc7, 0x15, 0x19, 0xd2, 0x8c, 0x60,
		0xe1, 0x74, 0x71, 0xee, 0xc8, 0x21, 0xb9, 0xbf, 0x55, 0x72, 0xe7, 0x41, 0x79, 0x40, 0x19, 0x84,
		0x24, 0x95, 0xcc, 0xac, 0xac, 0x2c, 0xca, 0x09, 0x9f, 0xb3, 0xb4, 0xa6, 0x2e, 0x2c, 0xcf, 0xcb,
		0x49, 0x9f, 0x73, 0x5e, 0x5d, 0x59, 0x5f, 0x95, 0xc1, 0x67, 0x58, 0x2a, 0x96, 0x4a, 0x85, 0xf9,
		0xa2, 0x9c, 0xf2, 0x35, 0x66, 0x1e, 0x58, 0x2b, 0x96, 0xe4, 0x74, 0x8b, 0x59, 0x77, 0x1e, 0x94,
		0x07, 0xfd, 0x47, 0x14, 0x97, 0xd7, 0x97, 0xe4, 0x8c, 0x32, 0x0c, 0x83, 0xec, 0x11, 0xc2, 0x88,
		0xa1, 0x36, 0xd1, 0x91, 0x43, 0xb2, 0xdc, 0x34, 0x84, 0xb1, 0x0c, 0xb7, 0x08, 0x8e, 0x1c, 0x92,
		0x95, 0xc9, 0x59, 0x88, 0xd3, 0xe8, 0x52, 0x14, 0xc8, 0x2c, 0x16, 0x66, 0x8a, 0x8b, 0xda, 0xca,
		0xea, 0xda, 0xc2, 0xca, 0x72, 0x61, 0x51, 0x96, 0x9a, 0x32, 0xb5, 0xf8, 0xc5, 0xf5, 0x05, 0xb5,
		0x38, 0x27, 0x47, 0x82, 0xb2, 0xd5, 0x62, 0x61, 0xad, 0x38, 0x27, 0x47, 0x27, 0x75, 0x18, 0xed,
		0x96, 0x67, 0xba, 0x9e, 0x8c, 0xc0, 0x16, 0x47, 0x76, 0xd8, 0x62, 0xca, 0xd5, 0xb1, 0xc5, 0x2f,
		0x48, 0x30, 0xd2, 0x25, 0xd7, 0x76, 0x7d, 0xc8, 0x09, 0x88, 0xb3, 0x10, 0x65, 0xd5, 0xe7, 0x96,
		0xae, 0x49, 0x9b, 0x06, 0x6c, 0x47, 0x05, 0xa2, 0xb8, 0x60, 0x05, 0x8e, 0xee, 0x50, 0x81, 0x09,
		0x45, 0x87, 0x91, 0x8f, 0x48, 0x90, 0xdd, 0x89, 0x3b, 0x24, 0x51, 0x44, 0x5a, 0x12, 0xc5, 0xf1,
		0x76, 0x03, 0xf6, 0xef, 0xbc, 0x86, 0x0e, 0x2b, 0x5e, 0x94, 0x60, 0xac, 0x7b, 0xa3, 0xd2, 0xd5,
		0x86, 0x7b, 0xa0, 0xbf, 0x86, 0xbd, 0x4d, 0x5b, 0x14, 0xeb, 0xcf, 0x75, 0x29, 0x01, 0x64, 0xba,
		0xdd, 0x57, 0x1c, 0x15, 0xac, 0x21, 0xd1, 0x9d, 0xba, 0x0d, 0x66, 0x4d, 0x87, 0xa5, 0x8f, 0x46,
		0xe0, 0x9a, 0xae, 0xe4, 0x5d, 0x0d, 0xbd, 0x1e, 0xc0, 0xb0, 0xea, 0x0d, 0x8f, 0x15, 0x64, 0x96,
		0x9f, 0x92, 0x54, 0x42, 0xcf, 0x3e, 0xc9, 0x3d, 0x0d, 0xcf, 0x9f, 0x8f, 0xd2, 0x79, 0x60, 0x22,
		0xaa, 0x70, 0xb4, 0x69, 0x68, 0x8c, 0x1a, 0x9a, 0xdb, 0x61, 0xa5, 0x1d, 0xb5, 0xee, 0x0e, 0x90,
		0x75, 0xd3, 0xc0, 0x96, 0xa7, 0xb9, 0x9e, 0x83, 0x51, 0xcd, 0xb0, 0xaa, 0x34, 0x01, 0x27, 0xf2,
		0xf1, 0x0a, 0x32, 0x5d, 0xac, 0x0e, 0xb1, 0xe9, 0x92, 0x98, 0x25, 0x08, 0x5a, 0x65, 0x9c, 0x00,
		0xa2, 0xbf, 0x05, 0xc1, 0xa6, 0x7d, 0xc4, 0xe4, 0x63, 0x03, 0x90, 0x0a, 0xb4, 0x75, 0xca, 0x7e,
		0x48, 0x9f, 0x41, 0xe7, 0x90, 0x26, 0x5a, 0x75, 0xe6, 0x89, 0x14, 0x91, 0xad, 0xf2, 0x76, 0xfd,
		0x0e, 0x18, 0xa5, 0x2a, 0x76, 0xc3, 0xc3, 0x8e, 0xa6, 0x9b, 0xc8, 0x75, 0xa9, 0xd3, 0x12, 0x54,
		0x55, 0x21, 0x73, 0x2b, 0x64, 0x6a, 0x56, 0xcc, 0x28, 0x87, 0x61, 0x84, 0x22, 0x6a, 0x0d, 0xd3,
		0x33, 0xea, 0x26, 0xd6, 0xc8, 0xcb, 0x83, 0x4b, 0x13, 0xb1, 0x6f, 0xd9, 0x30, 0xd1, 0x58, 0xe2,
		0x0a, 0xc4, 0x22, 0x57, 0x99, 0x87, 0xeb, 0x29, 0xac, 0x8a, 0x2d, 0xec, 0x20, 0x0f, 0x6b, 0xf8,
		0x4b, 0x0d, 0x64, 0xba, 0x1a, 0xb2, 0xca, 0xda, 0x26, 0x72, 0x37, 0xb3, 0xa3, 0x41, 0x82, 0xbd,
		0x44, 0x77, 0x9e, 0xab, 0x16, 0xa9, 0x66, 0xc1, 0x2a, 0xdf, 0x8b, 0xdc, 0x4d, 0x25, 0x0f, 0x63,
		0x94, 0xc8, 0xf5, 0x1c, 0xc3, 0xaa, 0x6a, 0xfa, 0x26, 0xd6, 0xcf, 0x6a, 0x0d, 0xaf, 0x72, 0x34,
		0x7b, 0x6d, 0x90, 0x81, 0x1a, 0x59, 0xa2, 0x3a, 0xb3, 0x44, 0x65, 0xdd, 0xab, 0x1c, 0x55, 0x4a,
		0x90, 0x26, 0xfb, 0x51, 0x33, 0x2e, 0x62, 0xad, 0x62, 0x3b, 0xb4, 0xb8, 0x64, 0xba, 0x1c, 0xee,
		0x80, 0x13, 0xa7, 0x57, 0x38, 0x60, 0xc9, 0x2e, 0xe3, 0x7c, 0xbc, 0xb4, 0x5a, 0x2c, 0xce, 0xa9,
		0x29, 0xc1, 0x72, 0xd2, 0x76, 0x48, 0x4c, 0x55, 0x6d, 0xdf, 0xc7, 0x29, 0x16, 0x53, 0x55, 0x5b,
		0x78, 0xf8, 0x30, 0x8c, 0xe8, 0x3a, 0x5b, 0xb6, 0xa1, 0x6b, 0xbc, 0xcb, 0x77, 0xb3, 0x72, 0x8b,
		0xbf, 0x74, 0x7d, 0x9e, 0x29, 0xf0, 0x30, 0x77, 0x95, 0x63, 0x70, 0x4d, 0xd3, 0x5f, 0x41, 0xe0,
		0x70, 0xc7, 0x2a, 0xdb, 0xa1, 0x87, 0x61, 0xa4, 0xbe, 0xd5, 0x09, 0x54, 0x5a, 0x9e, 0x58, 0xdf,
		0x6a, 0x87, 0xdd, 0x44, 0xdf, 0xdc, 0x1c, 0xac, 0x23, 0x0f, 0x97, 0xb3, 0x7b, 0x82, 0xda, 0x81,
		0x09, 0xe5, 0x00, 0xc8, 0xba, 0xae, 0x61, 0x0b, 0x6d, 0x98, 0x58, 0x43, 0x0e, 0xb6, 0x90, 0x9b,
		0xdd, 0x17, 0x54, 0xce, 0xe8, 0x7a, 0x91, 0xce, 0x16, 0xe8, 0xa4, 0x72, 0x2b, 0x0c, 0xdb, 0x1b,
		0x67, 0x74, 0x16, 0x5c, 0x5a, 0xdd, 0xc1, 0x15, 0xe3, 0x42, 0xf6, 0x46, 0xea, 0xa6, 0x21, 0x32,
		0x41, 0x43, 0x6b, 0x95, 0x8a, 0x95, 0x5b, 0x40, 0xd6, 0xdd, 0x4d, 0xe4, 0xd4, 0x69, 0x75, 0x77,
		0xeb, 0x48, 0xc7, 0xd9, 0x9b, 0x98, 0x2a, 0x93, 0x2f, 0x0b, 0xb1, 0x72, 0x1a, 0x46, 0x1b, 0x96,
		0x61, 0x79, 0xd8, 0xa9, 0x3b, 0x98, 0x34, 0xe9, 0xec, 0xa4, 0x65, 0xff, 0x3a, 0xb0, 0x43, 0x9b,
		0xbd, 0x1e, 0xd4, 0x66, 0xbb, 0xab, 0x8e, 0x34, 0x3a, 0x85, 0x93, 0x79, 0x48, 0x07, 0x37, 0x5d,
		0x49, 0x02, 0xdb, 0x76, 0x59, 0x22, 0x35, 0x74, 0x76, 0x65, 0x8e, 0x54, 0xbf, 0x07, 0x8b, 0x72,
		0x84, 0x54, 0xe1, 0xc5, 0x85, 0xb5, 0xa2, 0xa6, 0xae, 0x2f, 0xaf, 0x2d, 0x2c, 0x15, 0xe5, 0xe8,
		0xad, 0xc9, 0xc4, 0x7b, 0x03, 0xf2, 0xa5, 0x4b, 0x97, 0x2e, 0x45, 0x26, 0xdf, 0x88, 0x40, 0xa6,
		0xb5, 0xf3, 0x55, 0xbe, 0x00, 0x7b, 0xc4, 0x6b, 0xaa, 0x8b, 0x3d, 0xed, 0xbc, 0xe1, 0xd0, 0x38,
		0xac, 0x21, 0xd6, 0x3b, 0xfa, 0x2e, 0x1c, 0xe5, 0x5a, 0x25, 0xec, 0xdd, 0x6f, 0x38, 0x24, 0xca,
		0x6a, 0xc8, 0x53, 0x16, 0x61, 0x9f, 0x65, 0x6b, 0xae, 0x87, 0xac, 0x32, 0x72, 0xca, 0x5a, 0xf3,
		0x82, 0x40, 0x43, 0xba, 0x8e, 0x5d, 0xd7, 0x66, 0x25, 0xc0, 0x67, 0xb9, 0xce, 0xb2, 0x4b, 0x5c,
		0xb9, 0x99, 0x1b, 0x0b, 0x5c, 0xb5, 0x6d, 0xbb, 0xa3, 0x3b, 0x6d, 0xf7, 0xb5, 0x90, 0xac, 0xa1,
		0xba, 0x86, 0x2d, 0xcf, 0xd9, 0xa2, 0xfd, 0x5a, 0x42, 0x4d, 0xd4, 0x50, 0xbd, 0x48, 0xc6, 0x9f,
		0xdc, 0x1e, 0x04, 0xfd, 0xf8, 0xe7, 0x28, 0xa4, 0x83, 0x3d, 0x1b, 0x69, 0x81, 0x75, 0x9a, 0x9f,
		0x25, 0x7a, 0x7c, 0x6f, 0xb8, 0x6a, 0x87, 0x37, 0x3d, 0x4b, 0x12, 0x77, 0xbe, 0x9f, 0x75, 0x52,
		0x2a, 0x43, 0x92, 0xa2, 0x49, 0x0e, 0x2c, 0x66, 0xfd, 0x79, 0x42, 0xe5, 0x23, 0x65, 0x1e, 0xfa,
		0xcf, 0xb8, 0x94, 0xbb, 0x9f, 0x72, 0xdf, 0x78, 0x75, 0xee, 0x53, 0x25, 0x4a, 0x9e, 0x3c, 0x55,
		0xd2, 0x96, 0x57, 0xd4, 0xa5, 0xc2, 0xa2, 0xca, 0xe1, 0xca, 0x5e, 0x88, 0x99, 0xe8, 0xe2, 0x56,
		0x6b, 0x8a, 0xa7, 0xa2, 0x5e, 0x1d, 0xbf, 0x17, 0x62, 0xe7, 0x31, 0x3a, 0xdb, 0x9a, 0x58, 0xa9,
		0xe8, 0x13, 0x0c, 0xfd, 0x03, 0x10, 0xa7, 0xfe, 0x52, 0x00, 0xb8, 0xc7, 0xe4, 0x3e, 0x25, 0x01,
		0xb1, 0xd9, 0x15, 0x95, 0x84, 0xbf, 0x0c, 0x69, 0x26, 0xd5, 0x56, 0x17, 0x8a, 0xb3, 0x45, 0x39,
		0x32, 0x79, 0x18, 0xfa, 0x99, 0x13, 0xc8, 0xd1, 0xf0, 0xdd, 0x20, 0xf7, 0xf1, 0x21, 0xe7, 0x90,
		0xc4, 0xec, 0xfa, 0xd2, 0x4c, 0x51, 0x95, 0x23, 0xc1, 0xed, 0x75, 0x21, 0x1d, 0x6c, 0xd7, 0x3e,
		0x9d, 0x98, 0xfa, 0x9d, 0x04, 0xa9, 0x40, 0xfb, 0x45, 0x0a, 0x3f, 0x32, 0x4d, 0xfb, 0xbc, 0x86,
		0x4c, 0x03, 0xb9, 0x3c, 0x28, 0x80, 0x8a, 0x0a, 0x44, 0xd2, 0xeb, 0xa6, 0x7d, 0x2a, 0xc6, 0x3f,
		0x2b, 0x81, 0xdc, 0xde, 0xba, 0xb5, 0x19, 0x28, 0x7d, 0xa6, 0x06, 0x3e, 0x2d, 0x41, 0xa6, 0xb5,
		0x5f, 0x6b, 0x33, 0x6f, 0xff, 0x67, 0x6a, 0xde, 0x53, 0x12, 0x0c, 0xb6, 0x74, 0x69, 0xff, 0x57,
		0xd6, 0x3d, 0x19, 0x85, 0x91, 0x2e, 0x38, 0xa5, 0xc0, 0xdb, 0x59, 0xd6, 0x61, 0xdf, 0xde, 0xcb,
		0xb3, 0xa6, 0x49, 0xb5, 0x5c, 0x45, 0x8e, 0xc7, 0xbb, 0xdf, 0x5b, 0x40, 0x36, 0xca, 0xd8, 0xf2,
		0x8c, 0x8a, 0x81, 0x1d, 0xfe, 0x0a, 0xce, 0x7a, 0xdc, 0xa1, 0xa6, 0x9c, 0xbd, 0x85, 0x7f, 0x1e,
		0x94, 0xba, 0xed, 0x1a, 0x9e, 0x71, 0x0e, 0x6b, 0x86, 0x25, 0xde, 0xd7, 0x49, 0xcf, 0x1b, 0x53,
		0x65, 0x31, 0xb3, 0x60, 0x79, 0xbe, 0xb6, 0x85, 0xab, 0xa8, 0x4d, 0x9b, 0xe4, 0xbe, 0xa8, 0x2a,
		0x8b, 0x19, 0x5f, 0x7b, 0x3f, 0xa4, 0xcb, 0x76, 0x83, 0xb4, 0x0f, 0x4c, 0x8f, 0xa4, 0x5a, 0x49,
		0x4d, 0x31, 0x99, 0xaf, 0xc2, 0xfb, 0xbb, 0xe6, 0x45, 0x41, 0x5a, 0x4d, 0x31, 0x19, 0x53, 0xb9,
		0x19, 0x86, 0x50, 0xb5, 0xea, 0x10, 0x72, 0x41, 0xc4, 0x9a, 0xd6, 0x8c, 0x2f, 0xa6, 0x8a, 0xe3,
		0xa7, 0x20, 0x21, 0xfc, 0x40, 0xaa, 0x19, 0xf1, 0x84, 0x56, 0x67, 0xd7, 0x35, 0x91, 0xa9, 0xa4,
		0x9a, 0xb0, 0xc4, 0xe4, 0x7e, 0x48, 0x1b, 0xae, 0xd6, 0xbc, 0x37, 0x8c, 0x4c, 0x44, 0xa6, 0x12,
		0x6a, 0xca, 0x70, 0xfd, 0x8b, 0xa2, 0xc9, 0x17, 0x23, 0x90, 0x69, 0xbd, 0xf7, 0x54, 0xe6, 0x20,
		0x61, 0xda, 0x3a, 0xa2, 0x81, 0xc0, 0x2e, 0xdd, 0xa7, 0x42, 0xae, 0x4a, 0xa7, 0x17, 0xb9, 0xbe,
		0xea, 0x23, 0xc7, 0xff, 0x28, 0x41, 0x42, 0x88, 0x95, 0x31, 0x88, 0xd5, 0x91, 0xb7, 0x49, 0xe9,
		0xe2, 0x33, 0x11, 0x59, 0x52, 0xe9, 0x98, 0xc8, 0xdd, 0x3a, 0xb2, 0x68, 0x08, 0x70, 0x39, 0x19,
		0x93, 0x7d, 0x35, 0x31, 0x2a, 0xd3, 0x76, 0xd8, 0xae, 0xd5, 0xb0, 0xe5, 0xb9, 0x62, 0x5f, 0xb9,
		0x7c, 0x96, 0x8b, 0x95, 0xdb, 0x60, 0xd8, 0x73, 0x90, 0x61, 0xb6, 0xe8, 0xc6, 0xa8, 0xae, 0x2c,
		0x26, 0x7c, 0xe5, 0x3c, 0xec, 0x15, 0xbc, 0x65, 0xec, 0x21, 0x7d, 0x13, 0x97, 0x9b, 0xa0, 0x7e,
		0x7a, 0xa9, 0xb6, 0x87, 0x2b, 0xcc, 0xf1, 0x79, 0x81, 0x9d, 0x7c, 0x5b, 0x82, 0x61, 0xd1, 0xc0,
		0x97, 0x7d, 0x67, 0x2d, 0x01, 0x20, 0xcb, 0xb2, 0xbd, 0xa0, 0xbb, 0x3a, 0x43, 0xb9, 0x03, 0x37,
		0x5d, 0xf0, 0x41, 0x6a, 0x80, 0x60, 0xbc, 0x06, 0xd0, 0x9c, 0xd9, 0xd1, 0x6d, 0xfb, 0x20, 0xc5,
		0x2f, 0xb5, 0xe9, 0x97, 0x11, 0xf6, 0xd6, 0x07, 0x4c, 0x44, 0x3a, 0x7d, 0x65, 0x14, 0xe2, 0x1b,
		0xb8, 0x6a, 0x58, 0xfc, 0xaa, 0x8d, 0x0d, 0xc4, 0x05, 0x5e, 0xcc, 0xbf, 0xc0, 0x9b, 0x79, 0x08,
		0x46, 0x74, 0xbb, 0xd6, 0x6e, 0xee, 0x8c, 0xdc, 0xf6, 0xe6, 0xe9, 0xde, 0x2b, 0x3d, 0x08, 0xcd,
		0xee, 0xec, 0x39, 0x49, 0x7a, 0x21, 0x12, 0x9d, 0x5f, 0x9d, 0x79, 0x39, 0x32, 0x3e, 0xcf, 0xa0,
		0xab, 0x62, 0xa5, 0x2a, 0xae, 0x98, 0x58, 0x27, 0xd6, 0xc3, 0xf3, 0x37, 0xc0, 0xed, 0x55, 0xc3,
		0xdb, 0x6c, 0x6c, 0x4c, 0xeb, 0x76, 0xed, 0x40, 0xd5, 0xae, 0xda, 0xcd, 0x8f, 0x41, 0x64, 0x44,
		0x07, 0xf4, 0x17, 0xff, 0x20, 0x94, 0xf4, 0xa5, 0xe3, 0xa1, 0x5f, 0x8f, 0xf2, 0xcb, 0x30, 0xc2,
		0x95, 0x35, 0x7a, 0x23, 0xcd, 0xfa, 0x70, 0xe5, 0xaa, 0xb7, 0x12, 0xd9, 0xd7, 0xde, 0xa5, 0x95,
		0x4e, 0x1d, 0xe6, 0x50, 0x32, 0xc7, 0x3a, 0xf5, 0xbc, 0x0a, 0xd7, 0xb4, 0xf0, 0xb1, 0xa3, 0x89,
		0x9d, 0x10, 0xc6, 0x37, 0x38, 0xe3, 0x48, 0x80, 0xb1, 0xc4, 0xa1, 0xf9, 0x59, 0x18, 0xdc, 0x0d,
		0xd7, 0xef, 0x39, 0x57, 0x1a, 0x07, 0x49, 0xe6, 0x61, 0x88, 0x92, 0xe8, 0x0d, 0xd7, 0xb3, 0x6b,
		0x34, 0xef, 0x5d, 0x9d, 0xe6, 0x0f, 0xef, 0xb2, 0xb3, 0x92, 0x21, 0xb0, 0x59, 0x1f, 0x95, 0xbf,
		0x0f, 0x46, 0x89, 0x84, 0xa6, 0x96, 0x20, 0x5b, 0xf8, 0x3d, 0x4a, 0xf6, 0xed, 0x47, 0xd8, 0x91,
		0x1a, 0xf1, 0x09, 0x02, 0xbc, 0x81, 0x9d, 0xa8, 0x62, 0xcf, 0xc3, 0x8e, 0xab, 0x21, 0xd3, 0x54,
		0xae, 0xfa, 0x85, 0x26, 0xfb, 0xc4, 0xfb, 0xad, 0x3b, 0x31, 0xcf, 0x90, 0x05, 0xd3, 0xcc, 0xaf,
		0xc3, 0x9e, 0x2e, 0x3b, 0xdb, 0x03, 0xe7, 0x93, 0x9c, 0x73, 0xb4, 0x63, 0x77, 0x09, 0xed, 0x2a,
		0x08, 0xb9, 0xbf, 0x1f, 0x3d, 0x70, 0x3e, 0xc5, 0x39, 0x15, 0x8e, 0x15, 0xdb, 0x42, 0x18, 0x4f,
		0xc1, 0xf0, 0x39, 0xec, 0x6c, 0xd8, 0x2e, 0x7f, 0xf9, 0xef, 0x81, 0xee, 0x69, 0x4e, 0x37, 0xc4,
		0x81, 0xf4, 0x2a, 0x80, 0x70, 0x1d, 0x83, 0x44, 0x05, 0xe9, 0xb8, 0x07, 0x8a, 0x67, 0x38, 0xc5,
		0x00, 0xd1, 0x27, 0xd0, 0x02, 0xa4, 0xab, 0x36, 0xaf, 0x2e, 0xe1, 0xf0, 0x67, 0x39, 0x3c, 0x25,
		0x30, 0x9c, 0xa2, 0x6e, 0xd7, 0x1b, 0x26, 0x29, 0x3d, 0xe1, 0x14, 0x3f, 0x10, 0x14, 0x02, 0xc3,
		0x29, 0x76, 0xe1, 0xd6, 0xe7, 0x04, 0x85, 0x1b, 0xf0, 0xe7, 0x09, 0x48, 0xd9, 0x96, 0xb9, 0x65,
		0x5b, 0xbd, 0x18, 0xf1, 0x3c, 0x67, 0x00, 0x0e, 0x21, 0x04, 0xc7, 0x21, 0xd9, 0xeb, 0x46, 0xfc,
		0x90, 0xc3, 0x13, 0x58, 0xec, 0xc0, 0x3c, 0x0c, 0x89, 0x24, 0x63, 0xd8, 0x56, 0x0f, 0x14, 0x3f,
		0xe2, 0x14, 0x99, 0x00, 0x8c, 0x2f, 0xc3, 0xc3, 0xae, 0x57, 0xc5, 0xbd, 0x90, 0xbc, 0x28, 0x96,
		0xc1, 0x21, 0xdc, 0x95, 0x1b, 0xd8, 0xd2, 0x37, 0x7b, 0x63, 0x78, 0x49, 0xb8, 0x52, 0x60, 0x08,
		0xc5, 0x2c, 0x0c, 0xd6, 0x90, 0xe3, 0x6e, 0x22, 0xb3, 0xa7, 0xed, 0xf8, 0x31, 0xe7, 0x48, 0xfb,
		0x20, 0xee, 0x91, 0x86, 0xb5, 0x1b, 0x9a, 0x97, 0x85, 0x47, 0x02, 0x30, 0x7e, 0xf4, 0x5c, 0x8f,
		0xde, 0xaf, 0xec, 0x86, 0xed, 0x27, 0xe2, 0xe8, 0x31, 0xec, 0x52, 0x90, 0xf1, 0x38, 0x24, 0x5d,
		0xe3, 0x62, 0x4f, 0x34, 0x3f, 0x15, 0x3b, 0x4d, 0x01, 0x04, 0xfc, 0x00, 0xec, 0xed, 0x9a, 0xea,
		0x7b, 0x20, 0x7b, 0x85, 0x93, 0x8d, 0x75, 0x49, 0xf7, 0x3c, 0x25, 0xec, 0x96, 0xf2, 0x67, 0x22,
		0x25, 0xe0, 0x36, 0xae, 0x55, 0xd2, 0x9d, 0xbb, 0xa8, 0xb2, 0x3b, 0xaf, 0xfd, 0x5c, 0x78, 0x8d,
		0x61, 0x5b, 0xbc, 0xb6, 0x06, 0x63, 0x9c, 0x71, 0x77, 0xfb, 0xfa, 0xaa, 0x48, 0xac, 0x0c, 0xbd,
		0xde, 0xba, 0xbb, 0x0f, 0xc1, 0xb8, 0xef, 0x4e, 0xd1, 0x58, 0xba, 0x5a, 0x0d, 0xd5, 0x7b, 0x60,
		0x7e, 0x8d, 0x33, 0x8b, 0x8c, 0xef, 0x77, 0xa6, 0xee, 0x12, 0xaa, 0x13, 0xf2, 0xd3, 0x90, 0x15,
		0xe4, 0x0d, 0xcb, 0xc1, 0xba, 0x5d, 0xb5, 0x8c, 0x8b, 0xb8, 0xdc, 0x03, 0xf5, 0x2f, 0xda, 0xb6,
		0x6a, 0x3d, 0x00, 0x27, 0xcc, 0x0b, 0x20, 0xfb, 0xfd, 0x86, 0x66, 0xd4, 0xea, 0xb6, 0xe3, 0x85,
		0x30, 0xfe, 0x52, 0xec, 0x94, 0x8f, 0x5b, 0xa0, 0xb0, 0x7c, 0x11, 0x32, 0x74, 0xd8, 0x6b, 0x48,
		0xfe, 0x8a, 0x13, 0x0d, 0x36, 0x51, 0x3c, 0x71, 0xe8, 0x76, 0xad, 0x8e, 0x9c, 0x5e, 0xf2, 0xdf,
		0xaf, 0x45, 0xe2, 0xe0, 0x10, 0x16, 0x7d, 0x43, 0x6d, 0x95, 0x58, 0x09, 0xfb, 0x78, 0x9d, 0x7d,
		0xf8, 0x0a, 0x3f, 0xb3, 0xad, 0x85, 0x38, 0xbf, 0x48, 0xdc, 0xd3, 0x5a, 0x2e, 0xc3, 0xc9, 0x1e,
		0xb9, 0xe2, 0x7b, 0xa8, 0xa5, 0x5a, 0xe6, 0x4f, 0xc2, 0x60, 0x4b, 0xa9, 0x0c, 0xa7, 0xfa, 0x32,
		0xa7, 0x4a, 0x07, 0x2b, 0x65, 0xfe, 0x30, 0xc4, 0x48, 0xd9, 0x0b, 0x87, 0x7f, 0x85, 0xc3, 0xa9,
		0x7a, 0xfe, 0x6e, 0x48, 0x88, 0x72, 0x17, 0x0e, 0xfd, 0x2a, 0x87, 0xfa, 0x10, 0x02, 0x17, 0xa5,
		0x2e, 0x1c, 0xfe, 0x35, 0x01, 0x17, 0x10, 0x02, 0xef, 0xdd, 0x85, 0xaf, 0x3f, 0x16, 0xe3, 0xe9,
		0x4a, 0xf8, 0xee, 0x38, 0x0c, 0xf0, 0x1a, 0x17, 0x8e, 0x7e, 0x94, 0x3f, 0x5c, 0x20, 0xf2, 0x77,
		0x41, 0xbc, 0x47, 0x87, 0x7f, 0x9d, 0x43, 0x99, 0x7e, 0x7e, 0x16, 0x52, 0x81, 0xba, 0x16, 0x0e,
		0xff, 0x06, 0x87, 0x07, 0x51, 0xc4, 0x74, 0x5e, 0xd7, 0xc2, 0x09, 0xbe, 0x29, 0x4c, 0xe7, 0x08,
		0xe2, 0x36, 0x51, 0xd2, 0xc2, 0xd1, 0xdf, 0x12, 0x5e, 0x17, 0x90, 0xfc, 0x09, 0x48, 0xfa, 0x69,
		0x2a, 0x1c, 0xff, 0x6d, 0x8e, 0x6f, 0x62, 0x88, 0x07, 0x02, 0x69, 0x32, 0x9c, 0xe2, 0x3b, 0xc2,
		0x03, 0x01, 0x14, 0x39, 0x46, 0xed, 0xa5, 0x2f, 0x9c, 0xe9, 0xbb, 0xe2, 0x18, 0xb5, 0x55, 0x3e,
		0xb2, 0x9b, 0x34, 0x5b, 0x84, 0x53, 0x7c, 0x4f, 0xec, 0x26, 0xd5, 0x27, 0x66, 0xb4, 0xd7, 0x92,
		0x70, 0x8e, 0xef, 0x0b, 0x33, 0xda, 0x4a, 0x49, 0x7e, 0x15, 0x94, 0xce, 0x3a, 0x12, 0xce, 0xf7,
		0x38, 0xe7, 0x1b, 0xee, 0x28, 0x23, 0xf9, 0xfb, 0x61, 0xac, 0x7b, 0x0d, 0x09, 0x67, 0x7d, 0xe2,
		0x4a, 0x5b, 0xd7, 0x1f, 0x2c, 0x21, 0xf9, 0xb5, 0x66, 0xd7, 0x1f, 0xac, 0x1f, 0xe1, 0xb4, 0x4f,
		0x5e, 0x69, 0x7d, 0xb1, 0x0b, 0x96, 0x8f, 0x7c, 0x01, 0xa0, 0x99, 0xba, 0xc3, 0xb9, 0x9e, 0xe6,
		0x5c, 0x01, 0x10, 0x39, 0x1a, 0x3c, 0x73, 0x87, 0xe3, 0x9f, 0x11, 0x47, 0x83, 0x23, 0xf2, 0xc7,
		0x21, 0x61, 0x35, 0x4c, 0x93, 0x04, 0x87, 0x72, 0xf5, 0x7f, 0x08, 0xc9, 0xfe, 0xed, 0x23, 0x7e,
		0x30, 0x04, 0x20, 0x7f, 0x18, 0xe2, 0xb8, 0xb6, 0x81, 0xcb, 0x61, 0xc8, 0xbf, 0x7f, 0x24, 0x12,
		0x02, 0xd1, 0xce, 0x9f, 0x00, 0x60, 0x2f, 0x8d, 0xf4, 0x7b, 0x40, 0x08, 0xf6, 0x1f, 0x1f, 0xf1,
		0x6f, 0xcd, 0x4d, 0x48, 0x93, 0x80, 0x7d, 0xb9, 0xbe, 0x3a, 0xc1, 0xfb, 0xad, 0x04, 0xf4, 0x45,
		0xf3, 0x18, 0x0c, 0x9c, 0x71, 0x6d, 0xcb, 0x43, 0xd5, 0x30, 0xf4, 0x3f, 0x39, 0x5a, 0xe8, 0x13,
		0x87, 0xd5, 0x6c, 0x07, 0x7b, 0xa8, 0xea, 0x86, 0x61, 0xff, 0xc5, 0xb1, 0x3e, 0x80, 0x80, 0x75,
		0xe4, 0x7a, 0xbd, 0xac, 0xfb, 0xdf, 0x02, 0x2c, 0x00, 0xc4, 0x68, 0xf2, 0xfb, 0x2c, 0xde, 0x0a,
		0xc3, 0x7e, 0x20, 0x8c, 0xe6, 0xfa, 0xf9, 0xbb, 0x21, 0x49, 0x7e, 0xb2, 0xff, 0xbf, 0x08, 0x01,
		0xff, 0x87, 0x83, 0x9b, 0x88, 0x99, 0xfd, 0xdd, 0x6f, 0x77, 0x60, 0xde, 0x9e, 0xb7, 0xd9, 0xbd,
		0x0e, 0xbc, 0x22, 0x41, 0xa6, 0x62, 0x98, 0x78, 0xba, 0x6c, 0x7b, 0xfc, 0x12, 0x26, 0x45, 0xc6,
		0x65, 0xdb, 0x23, 0x1e, 0x1f, 0xdf, 0xdd, 0x05, 0xce, 0xe4, 0x30, 0x48, 0x4b, 0x4a, 0x1a, 0x24,
		0xc4, 0xbf, 0xcc, 0x4b, 0x68, 0x66, 0xf1, 0xcd, 0xcb, 0xb9, 0xbe, 0xb7, 0x2e, 0xe7, 0xfa, 0xfe,
		0x74, 0x39, 0xd7, 0xf7, 0xce, 0xe5, 0x9c, 0xf4, 0xde, 0xe5, 0x9c, 0xf4, 0xc1, 0xe5, 0x9c, 0xf4,
		0xe1, 0xe5, 0x9c, 0x74, 0x69, 0x3b, 0x27, 0xbd, 0xb4, 0x9d, 0x93, 0x5e, 0xdd, 0xce, 0x49, 0xbf,
		0xdd, 0xce, 0x49, 0xaf, 0x6f, 0xe7, 0xa4, 0x37, 0xb7, 0x73, 0x7d, 0x6f, 0x6d, 0xe7, 0xfa, 0xde,
		0xd9, 0xce, 0x49, 0xef, 0x6d, 0xe7, 0xfa, 0x3e, 0xd8, 0xce, 0x49, 0x1f, 0x6e, 0xe7, 0xfa, 0x2e,
		0xfd, 0x25, 0xd7, 0xf7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xdb, 0x4d, 0xde, 0x58, 0x2c,
		0x00, 0x00,
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
func (this *M) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *M")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *M but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *M but is not nil && this == nil")
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return fmt.Errorf("A this(%v) Not Equal that(%v)", *this.A, *that1.A)
		}
	} else if this.A != nil {
		return fmt.Errorf("this.A == nil && that.A != nil")
	} else if that1.A != nil {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *M) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return false
		}
	} else if this.A != nil {
		return false
	} else if that1.A != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type MFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetA() *string
}

func (this *M) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *M) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewMFromFace(this)
}

func (this *M) GetA() *string {
	return this.A
}

func NewMFromFace(that MFace) *M {
	this := &M{}
	this.A = that.GetA()
	return this
}

func (this *M) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filedotname.M{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringFileDot(this.A, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFileDot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFileDot(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func NewPopulatedM(r randyFileDot, easy bool) *M {
	this := &M{}
	if r.Intn(10) != 0 {
		v1 := randStringFileDot(r)
		this.A = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFileDot(r, 2)
	}
	return this
}

type randyFileDot interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFileDot(r randyFileDot) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFileDot(r randyFileDot) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFileDot(r)
	}
	return string(tmps)
}
func randUnrecognizedFileDot(r randyFileDot, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFileDot(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFileDot(data []byte, r randyFileDot, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateFileDot(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFileDot(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFileDot(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *M) Size() (n int) {
	var l int
	_ = l
	if m.A != nil {
		l = len(*m.A)
		n += 1 + l + sovFileDot(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileDot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileDot(x uint64) (n int) {
	return sovFileDot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *M) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&M{`,
		`A:` + valueToStringFileDot(this.A) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFileDot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() { proto.RegisterFile("file.dot.proto", fileDescriptorFileDot) }

var fileDescriptorFileDot = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x24, 0xcb, 0xaf, 0x6e, 0xc2, 0x50,
	0x1c, 0xc5, 0xf1, 0xdf, 0x91, 0xeb, 0x96, 0x25, 0xab, 0x5a, 0x26, 0x4e, 0x96, 0xa9, 0x99, 0xb5,
	0xef, 0x30, 0x0d, 0x86, 0x37, 0x68, 0xe9, 0x1f, 0x9a, 0x50, 0x2e, 0x21, 0xb7, 0xbe, 0x8f, 0x83,
	0x44, 0x22, 0x91, 0x95, 0x95, 0xc8, 0xde, 0x1f, 0xa6, 0xb2, 0xb2, 0x92, 0x70, 0x71, 0xe7, 0x93,
	0x9c, 0x6f, 0xf0, 0x5e, 0x54, 0xdb, 0x3c, 0xca, 0x8c, 0x8d, 0xf6, 0x07, 0x63, 0x4d, 0xf8, 0xfa,
	0x70, 0x66, 0xec, 0x2e, 0xa9, 0xf3, 0xaf, 0xbf, 0xb2, 0xb2, 0x9b, 0x26, 0x8d, 0xd6, 0xa6, 0x8e,
	0x4b, 0x53, 0x9a, 0xd8, 0x7f, 0xd2, 0xa6, 0xf0, 0xf2, 0xf0, 0xeb, 0xd9, 0xfe, 0x7c, 0x04, 0x58,
	0x86, 0x6f, 0x01, 0x92, 0x4f, 0x7c, 0xe3, 0xf7, 0x65, 0x85, 0xe4, 0x7f, 0xd1, 0x39, 0x4a, 0xef,
	0x28, 0x57, 0x47, 0x19, 0x1c, 0x31, 0x3a, 0x62, 0x72, 0xc4, 0xec, 0x88, 0x56, 0x89, 0xa3, 0x12,
	0x27, 0x25, 0xce, 0x4a, 0x5c, 0x94, 0xe8, 0x94, 0xd2, 0x2b, 0x65, 0x50, 0x62, 0x54, 0xca, 0xa4,
	0xc4, 0xac, 0x94, 0xf6, 0x46, 0xb9, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x59, 0x32, 0x8a, 0xad,
	0x00, 0x00, 0x00,
}
