// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/range.proto

package _type

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Range) Reset()         { *m = Int64Range{} }
func (m *Int64Range) String() string { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()    {}
func (*Int64Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1bd7ecef6de4614, []int{0}
}
func (m *Int64Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Range.Unmarshal(m, b)
}
func (m *Int64Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Range.Marshal(b, m, deterministic)
}
func (m *Int64Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Range.Merge(m, src)
}
func (m *Int64Range) XXX_Size() int {
	return xxx_messageInfo_Int64Range.Size(m)
}
func (m *Int64Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Range proto.InternalMessageInfo

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1bd7ecef6de4614, []int{1}
}
func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.Int64Range")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.DoubleRange")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/range.proto", fileDescriptor_b1bd7ecef6de4614)
}

var fileDescriptor_b1bd7ecef6de4614 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0x03, 0x31,
	0x10, 0x44, 0x65, 0x02, 0x14, 0x4b, 0x83, 0x4e, 0x29, 0x4e, 0x29, 0x10, 0x4a, 0x45, 0x83, 0xb7,
	0x20, 0xd0, 0x13, 0xd1, 0xa4, 0x8b, 0x52, 0xd2, 0xdd, 0x85, 0x95, 0x39, 0x30, 0x5e, 0xcb, 0xde,
	0x20, 0xdf, 0x1f, 0xf1, 0x09, 0x7c, 0x0f, 0xff, 0x40, 0x8f, 0x6c, 0x17, 0x34, 0x87, 0x44, 0xb7,
	0xf3, 0xe4, 0x19, 0x8d, 0x07, 0x36, 0x66, 0x90, 0xe7, 0x43, 0xaf, 0xf7, 0xfc, 0x86, 0x91, 0x2d,
	0x5f, 0x0f, 0x8c, 0xc6, 0x32, 0xa3, 0x0f, 0xfc, 0x42, 0x7b, 0x89, 0x55, 0x75, 0x7e, 0x40, 0x4a,
	0x42, 0xc1, 0x75, 0x16, 0xc9, 0xbd, 0xf3, 0x88, 0x32, 0x7a, 0xc2, 0xd0, 0x39, 0x43, 0xda, 0x07,
	0x16, 0x6e, 0xa0, 0x70, 0x9d, 0xf9, 0x62, 0x6e, 0xd8, 0x70, 0xc1, 0x98, 0xaf, 0xfa, 0x62, 0xd1,
	0x50, 0x92, 0x0a, 0x29, 0x49, 0x65, 0xcb, 0x15, 0xc0, 0xc6, 0xc9, 0xdd, 0x6a, 0x97, 0x93, 0x9a,
	0x39, 0x9c, 0x44, 0xe9, 0x82, 0xb4, 0xea, 0x52, 0x5d, 0xcd, 0x76, 0x55, 0x34, 0xe7, 0x30, 0x23,
	0xf7, 0xd4, 0x1e, 0x15, 0x96, 0xcf, 0xe5, 0x2d, 0x9c, 0x3d, 0xf0, 0xa1, 0xb7, 0x34, 0x61, 0x53,
	0x13, 0x36, 0x55, 0x6c, 0xeb, 0xf8, 0xf9, 0x7d, 0xac, 0x3e, 0xbe, 0x2e, 0x14, 0xb4, 0x03, 0xeb,
	0xd2, 0xd7, 0x07, 0x4e, 0xa3, 0xfe, 0xad, 0xbe, 0x86, 0x12, 0xb9, 0xcd, 0xe5, 0xb6, 0xea, 0xf1,
	0xfe, 0x7f, 0xfb, 0xf8, 0x57, 0xf3, 0xd7, 0x46, 0xfd, 0x69, 0xf9, 0xe8, 0xcd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x9a, 0xfe, 0xec, 0x6b, 0x01, 0x00, 0x00,
}

func (this *Int64Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int64Range)
	if !ok {
		that2, ok := that.(Int64Range)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DoubleRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoubleRange)
	if !ok {
		that2, ok := that.(DoubleRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}