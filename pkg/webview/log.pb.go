// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/webview/log.proto

package webview

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LogLevel int32

const (
	// For backwards-compatibility, the int value doesn't say
	// anything about relative severity.
	LogLevel_NONE    LogLevel = 0
	LogLevel_INFO    LogLevel = 1
	LogLevel_VERBOSE LogLevel = 2
	LogLevel_DEBUG   LogLevel = 3
	LogLevel_WARN    LogLevel = 4
	LogLevel_ERROR   LogLevel = 5
)

var LogLevel_name = map[int32]string{
	0: "NONE",
	1: "INFO",
	2: "VERBOSE",
	3: "DEBUG",
	4: "WARN",
	5: "ERROR",
}

var LogLevel_value = map[string]int32{
	"NONE":    0,
	"INFO":    1,
	"VERBOSE": 2,
	"DEBUG":   3,
	"WARN":    4,
	"ERROR":   5,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}

func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3da405ef8c289d22, []int{0}
}

type LogSegment struct {
	SpanId string               `protobuf:"bytes,1,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Time   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Text   string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Level  LogLevel             `protobuf:"varint,4,opt,name=level,proto3,enum=webview.LogLevel" json:"level,omitempty"`
	// When we store warnings in the LogStore, we break them up into lines and
	// store them as a series of line segments. 'anchor' marks the beginning of a
	// series of logs that should be kept together.
	//
	// Anchor warning1, line1
	//        warning1, line2
	// Anchor warning2, line1
	Anchor               bool     `protobuf:"varint,5,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSegment) Reset()         { *m = LogSegment{} }
func (m *LogSegment) String() string { return proto.CompactTextString(m) }
func (*LogSegment) ProtoMessage()    {}
func (*LogSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3da405ef8c289d22, []int{0}
}

func (m *LogSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSegment.Unmarshal(m, b)
}
func (m *LogSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSegment.Marshal(b, m, deterministic)
}
func (m *LogSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSegment.Merge(m, src)
}
func (m *LogSegment) XXX_Size() int {
	return xxx_messageInfo_LogSegment.Size(m)
}
func (m *LogSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSegment.DiscardUnknown(m)
}

var xxx_messageInfo_LogSegment proto.InternalMessageInfo

func (m *LogSegment) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *LogSegment) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *LogSegment) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *LogSegment) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_NONE
}

func (m *LogSegment) GetAnchor() bool {
	if m != nil {
		return m.Anchor
	}
	return false
}

type LogSpan struct {
	ManifestName         string   `protobuf:"bytes,1,opt,name=manifest_name,json=manifestName,proto3" json:"manifest_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogSpan) Reset()         { *m = LogSpan{} }
func (m *LogSpan) String() string { return proto.CompactTextString(m) }
func (*LogSpan) ProtoMessage()    {}
func (*LogSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_3da405ef8c289d22, []int{1}
}

func (m *LogSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSpan.Unmarshal(m, b)
}
func (m *LogSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSpan.Marshal(b, m, deterministic)
}
func (m *LogSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSpan.Merge(m, src)
}
func (m *LogSpan) XXX_Size() int {
	return xxx_messageInfo_LogSpan.Size(m)
}
func (m *LogSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSpan.DiscardUnknown(m)
}

var xxx_messageInfo_LogSpan proto.InternalMessageInfo

func (m *LogSpan) GetManifestName() string {
	if m != nil {
		return m.ManifestName
	}
	return ""
}

type LogList struct {
	Spans    map[string]*LogSpan `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Segments []*LogSegment       `protobuf:"bytes,2,rep,name=segments,proto3" json:"segments,omitempty"`
	// from_checkpoint and to_checkpoint express an interval on the
	// central log-store, with an inclusive start and an exclusive end
	//
	// [from_checkpoint, to_checkpoint)
	//
	// An interval of [0, 0) means that the server isn't using
	// the incremental load protocol.
	//
	// An interval of [-1, -1) means that the server doesn't have new logs
	// to send down.
	FromCheckpoint       int32    `protobuf:"varint,3,opt,name=from_checkpoint,json=fromCheckpoint,proto3" json:"from_checkpoint,omitempty"`
	ToCheckpoint         int32    `protobuf:"varint,4,opt,name=to_checkpoint,json=toCheckpoint,proto3" json:"to_checkpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogList) Reset()         { *m = LogList{} }
func (m *LogList) String() string { return proto.CompactTextString(m) }
func (*LogList) ProtoMessage()    {}
func (*LogList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3da405ef8c289d22, []int{2}
}

func (m *LogList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogList.Unmarshal(m, b)
}
func (m *LogList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogList.Marshal(b, m, deterministic)
}
func (m *LogList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogList.Merge(m, src)
}
func (m *LogList) XXX_Size() int {
	return xxx_messageInfo_LogList.Size(m)
}
func (m *LogList) XXX_DiscardUnknown() {
	xxx_messageInfo_LogList.DiscardUnknown(m)
}

var xxx_messageInfo_LogList proto.InternalMessageInfo

func (m *LogList) GetSpans() map[string]*LogSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *LogList) GetSegments() []*LogSegment {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *LogList) GetFromCheckpoint() int32 {
	if m != nil {
		return m.FromCheckpoint
	}
	return 0
}

func (m *LogList) GetToCheckpoint() int32 {
	if m != nil {
		return m.ToCheckpoint
	}
	return 0
}

func init() {
	proto.RegisterEnum("webview.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterType((*LogSegment)(nil), "webview.LogSegment")
	proto.RegisterType((*LogSpan)(nil), "webview.LogSpan")
	proto.RegisterType((*LogList)(nil), "webview.LogList")
	proto.RegisterMapType((map[string]*LogSpan)(nil), "webview.LogList.SpansEntry")
}

func init() { proto.RegisterFile("pkg/webview/log.proto", fileDescriptor_3da405ef8c289d22) }

var fileDescriptor_3da405ef8c289d22 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xb5, 0xd0, 0x02, 0x7b, 0x59, 0xd7, 0x3a, 0x46, 0x6d, 0xf0, 0xc1, 0x06, 0x13, 0xa9, 0x3e,
	0x4c, 0x23, 0xbe, 0x18, 0xdf, 0x44, 0xab, 0x59, 0x25, 0x25, 0x99, 0xf5, 0x23, 0xf1, 0x85, 0x14,
	0x18, 0x86, 0x09, 0x9d, 0x99, 0x86, 0x19, 0xc0, 0xfd, 0x05, 0xfe, 0x17, 0x7f, 0xa5, 0x99, 0x7e,
	0xec, 0xb2, 0x6f, 0x77, 0xce, 0x3d, 0xe7, 0xf4, 0xdc, 0xdb, 0x0b, 0x8f, 0x8b, 0x2d, 0x8b, 0x8f,
	0x74, 0x71, 0xe0, 0xf4, 0x18, 0xe7, 0x8a, 0xe1, 0x62, 0xa7, 0x8c, 0x42, 0xdd, 0x1a, 0x1a, 0x3c,
	0x67, 0x4a, 0xb1, 0x9c, 0xc6, 0x25, 0xbc, 0xd8, 0xaf, 0x63, 0xc3, 0x05, 0xd5, 0x26, 0x13, 0x45,
	0xc5, 0x1c, 0xfe, 0x73, 0x00, 0xa6, 0x8a, 0x5d, 0x51, 0x26, 0xa8, 0x34, 0xe8, 0x29, 0x74, 0x75,
	0x91, 0xc9, 0x39, 0x5f, 0x05, 0x4e, 0xe8, 0x44, 0x67, 0xa4, 0x63, 0x9f, 0x97, 0x2b, 0x84, 0xc1,
	0xb5, 0xd2, 0xa0, 0x15, 0x3a, 0x51, 0x7f, 0x3c, 0xc0, 0x95, 0x2f, 0x6e, 0x7c, 0xf1, 0xf7, 0xc6,
	0x97, 0x94, 0x3c, 0x84, 0xc0, 0x35, 0xf4, 0x8f, 0x09, 0xda, 0xa5, 0x4b, 0x59, 0xa3, 0x11, 0x78,
	0x39, 0x3d, 0xd0, 0x3c, 0x70, 0x43, 0x27, 0xba, 0x18, 0x3f, 0xc4, 0x75, 0x4a, 0x3c, 0x55, 0x6c,
	0x6a, 0x1b, 0xa4, 0xea, 0xa3, 0x27, 0xd0, 0xc9, 0xe4, 0x72, 0xa3, 0x76, 0x81, 0x17, 0x3a, 0x51,
	0x8f, 0xd4, 0xaf, 0x21, 0x86, 0xae, 0xcd, 0x5a, 0x64, 0x12, 0xbd, 0x80, 0xfb, 0x22, 0x93, 0x7c,
	0x4d, 0xb5, 0x99, 0xcb, 0x4c, 0xd0, 0x3a, 0xee, 0x79, 0x03, 0xa6, 0x99, 0xa0, 0xc3, 0xbf, 0xad,
	0x52, 0x30, 0xe5, 0xda, 0xa0, 0x37, 0xe0, 0xd9, 0x51, 0x74, 0xe0, 0x84, 0xed, 0xa8, 0x3f, 0x7e,
	0x76, 0xe7, 0xe3, 0x5c, 0x1b, 0x6c, 0x6d, 0x75, 0x22, 0xcd, 0xee, 0x9a, 0x54, 0x4c, 0x14, 0x43,
	0x4f, 0x57, 0x7b, 0xd1, 0x41, 0xab, 0x54, 0x3d, 0x3a, 0x55, 0xd5, 0x3b, 0x23, 0x37, 0x24, 0x34,
	0x82, 0x07, 0xeb, 0x9d, 0x12, 0xf3, 0xe5, 0x86, 0x2e, 0xb7, 0x85, 0xe2, 0xb2, 0x9a, 0xdf, 0x23,
	0x17, 0x16, 0xfe, 0x78, 0x83, 0xda, 0xf4, 0x46, 0x9d, 0xd2, 0xdc, 0x92, 0x76, 0x6e, 0xd4, 0x2d,
	0x69, 0xf0, 0x15, 0xe0, 0x36, 0x13, 0xf2, 0xa1, 0xbd, 0xa5, 0xd7, 0xf5, 0x98, 0xb6, 0x44, 0x2f,
	0xc1, 0x3b, 0x64, 0xf9, 0xbe, 0xf9, 0x27, 0xfe, 0x9d, 0x6c, 0x45, 0x26, 0x49, 0xd5, 0x7e, 0xdf,
	0x7a, 0xe7, 0xbc, 0xfe, 0x06, 0xbd, 0x66, 0xc9, 0xa8, 0x07, 0x6e, 0x3a, 0x4b, 0x13, 0xff, 0x9e,
	0xad, 0x2e, 0xd3, 0xcf, 0x33, 0xdf, 0x41, 0x7d, 0xe8, 0xfe, 0x4c, 0xc8, 0x64, 0x76, 0x95, 0xf8,
	0x2d, 0x74, 0x06, 0xde, 0xa7, 0x64, 0xf2, 0xe3, 0x8b, 0xdf, 0xb6, 0x8c, 0x5f, 0x1f, 0x48, 0xea,
	0xbb, 0x16, 0x4c, 0x08, 0x99, 0x11, 0xdf, 0x9b, 0xbc, 0xfa, 0x3d, 0x62, 0xdc, 0x6c, 0xf6, 0x0b,
	0xbc, 0x54, 0x22, 0x3e, 0x72, 0xb9, 0x12, 0x3c, 0xcf, 0xa9, 0x64, 0xb1, 0xe1, 0xb9, 0x89, 0x4f,
	0x4e, 0x72, 0xd1, 0x29, 0x0f, 0xe4, 0xed, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe8, 0x23,
	0xf4, 0xa8, 0x02, 0x00, 0x00,
}
