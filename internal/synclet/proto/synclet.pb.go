// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/synclet/synclet.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	LogLevel_INFO    LogLevel = 0
	LogLevel_VERBOSE LogLevel = 1
	LogLevel_DEBUG   LogLevel = 2
)

var LogLevel_name = map[int32]string{
	0: "INFO",
	1: "VERBOSE",
	2: "DEBUG",
}

var LogLevel_value = map[string]int32{
	"INFO":    0,
	"VERBOSE": 1,
	"DEBUG":   2,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}

func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{0}
}

type Cmd struct {
	Argv                 []string `protobuf:"bytes,1,rep,name=argv,proto3" json:"argv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{0}
}

func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type UpdateContainerRequest struct {
	ContainerId          string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	TarArchive           []byte    `protobuf:"bytes,2,opt,name=tar_archive,json=tarArchive,proto3" json:"tar_archive,omitempty"`
	FilesToDelete        []string  `protobuf:"bytes,3,rep,name=files_to_delete,json=filesToDelete,proto3" json:"files_to_delete,omitempty"`
	Commands             []*Cmd    `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	LogStyle             *LogStyle `protobuf:"bytes,5,opt,name=log_style,json=logStyle,proto3" json:"log_style,omitempty"`
	HotReload            bool      `protobuf:"varint,6,opt,name=hot_reload,json=hotReload,proto3" json:"hot_reload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateContainerRequest) Reset()         { *m = UpdateContainerRequest{} }
func (m *UpdateContainerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerRequest) ProtoMessage()    {}
func (*UpdateContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{1}
}

func (m *UpdateContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerRequest.Unmarshal(m, b)
}
func (m *UpdateContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerRequest.Merge(m, src)
}
func (m *UpdateContainerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerRequest.Size(m)
}
func (m *UpdateContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerRequest proto.InternalMessageInfo

func (m *UpdateContainerRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *UpdateContainerRequest) GetTarArchive() []byte {
	if m != nil {
		return m.TarArchive
	}
	return nil
}

func (m *UpdateContainerRequest) GetFilesToDelete() []string {
	if m != nil {
		return m.FilesToDelete
	}
	return nil
}

func (m *UpdateContainerRequest) GetCommands() []*Cmd {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *UpdateContainerRequest) GetLogStyle() *LogStyle {
	if m != nil {
		return m.LogStyle
	}
	return nil
}

func (m *UpdateContainerRequest) GetHotReload() bool {
	if m != nil {
		return m.HotReload
	}
	return false
}

type LogMessage struct {
	Level                LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=synclet.LogLevel" json:"level,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{2}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_INFO
}

func (m *LogMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type FailedRunStep struct {
	Cmd                  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	ExitCode             int32    `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailedRunStep) Reset()         { *m = FailedRunStep{} }
func (m *FailedRunStep) String() string { return proto.CompactTextString(m) }
func (*FailedRunStep) ProtoMessage()    {}
func (*FailedRunStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{3}
}

func (m *FailedRunStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailedRunStep.Unmarshal(m, b)
}
func (m *FailedRunStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailedRunStep.Marshal(b, m, deterministic)
}
func (m *FailedRunStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailedRunStep.Merge(m, src)
}
func (m *FailedRunStep) XXX_Size() int {
	return xxx_messageInfo_FailedRunStep.Size(m)
}
func (m *FailedRunStep) XXX_DiscardUnknown() {
	xxx_messageInfo_FailedRunStep.DiscardUnknown(m)
}

var xxx_messageInfo_FailedRunStep proto.InternalMessageInfo

func (m *FailedRunStep) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *FailedRunStep) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type UpdateContainerReply struct {
	LogMessage *LogMessage `protobuf:"bytes,1,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	// Contains info about run step failure (if any)
	FailedRunStep        *FailedRunStep `protobuf:"bytes,2,opt,name=failed_run_step,json=failedRunStep,proto3" json:"failed_run_step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateContainerReply) Reset()         { *m = UpdateContainerReply{} }
func (m *UpdateContainerReply) String() string { return proto.CompactTextString(m) }
func (*UpdateContainerReply) ProtoMessage()    {}
func (*UpdateContainerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{4}
}

func (m *UpdateContainerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContainerReply.Unmarshal(m, b)
}
func (m *UpdateContainerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContainerReply.Marshal(b, m, deterministic)
}
func (m *UpdateContainerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContainerReply.Merge(m, src)
}
func (m *UpdateContainerReply) XXX_Size() int {
	return xxx_messageInfo_UpdateContainerReply.Size(m)
}
func (m *UpdateContainerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContainerReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContainerReply proto.InternalMessageInfo

func (m *UpdateContainerReply) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func (m *UpdateContainerReply) GetFailedRunStep() *FailedRunStep {
	if m != nil {
		return m.FailedRunStep
	}
	return nil
}

type LogStyle struct {
	ColorsEnabled        bool     `protobuf:"varint,1,opt,name=colors_enabled,json=colorsEnabled,proto3" json:"colors_enabled,omitempty"`
	Level                LogLevel `protobuf:"varint,2,opt,name=level,proto3,enum=synclet.LogLevel" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStyle) Reset()         { *m = LogStyle{} }
func (m *LogStyle) String() string { return proto.CompactTextString(m) }
func (*LogStyle) ProtoMessage()    {}
func (*LogStyle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f1080a1433c7af1, []int{5}
}

func (m *LogStyle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStyle.Unmarshal(m, b)
}
func (m *LogStyle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStyle.Marshal(b, m, deterministic)
}
func (m *LogStyle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStyle.Merge(m, src)
}
func (m *LogStyle) XXX_Size() int {
	return xxx_messageInfo_LogStyle.Size(m)
}
func (m *LogStyle) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStyle.DiscardUnknown(m)
}

var xxx_messageInfo_LogStyle proto.InternalMessageInfo

func (m *LogStyle) GetColorsEnabled() bool {
	if m != nil {
		return m.ColorsEnabled
	}
	return false
}

func (m *LogStyle) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_INFO
}

func init() {
	proto.RegisterEnum("synclet.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterType((*Cmd)(nil), "synclet.Cmd")
	proto.RegisterType((*UpdateContainerRequest)(nil), "synclet.UpdateContainerRequest")
	proto.RegisterType((*LogMessage)(nil), "synclet.LogMessage")
	proto.RegisterType((*FailedRunStep)(nil), "synclet.FailedRunStep")
	proto.RegisterType((*UpdateContainerReply)(nil), "synclet.UpdateContainerReply")
	proto.RegisterType((*LogStyle)(nil), "synclet.LogStyle")
}

func init() { proto.RegisterFile("internal/synclet/synclet.proto", fileDescriptor_6f1080a1433c7af1) }

var fileDescriptor_6f1080a1433c7af1 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0x75, 0x5d, 0x93, 0xcb, 0xba, 0x15, 0x83, 0xa6, 0x00, 0x1a, 0x0b, 0x91, 0x80,
	0x08, 0xa1, 0x16, 0x15, 0x1e, 0x10, 0x48, 0x93, 0x58, 0xd7, 0xa1, 0x49, 0x85, 0x4a, 0x2e, 0xe5,
	0x61, 0x2f, 0xc1, 0x8d, 0xaf, 0x69, 0x24, 0x27, 0x0e, 0x89, 0x5b, 0xe8, 0xff, 0xc0, 0xbf, 0x8c,
	0x84, 0xea, 0x34, 0xdd, 0xc6, 0x80, 0x97, 0xfa, 0xee, 0x73, 0xbe, 0x1f, 0xdf, 0x6b, 0x0c, 0x8f,
	0xe2, 0x54, 0x61, 0x9e, 0x32, 0xd1, 0x29, 0x96, 0x69, 0x28, 0x50, 0x55, 0x67, 0x3b, 0xcb, 0xa5,
	0x92, 0xa4, 0xb1, 0x76, 0xbd, 0xfb, 0x50, 0xeb, 0x25, 0x9c, 0x10, 0xd8, 0x61, 0x79, 0xb4, 0x70,
	0x0c, 0xb7, 0xe6, 0x5b, 0x54, 0xdb, 0xde, 0x2f, 0x03, 0x0e, 0xc7, 0x19, 0x67, 0x0a, 0x7b, 0x32,
	0x55, 0x2c, 0x4e, 0x31, 0xa7, 0xf8, 0x6d, 0x8e, 0x85, 0x22, 0x8f, 0x61, 0x2f, 0xac, 0x58, 0x10,
	0x73, 0xc7, 0x70, 0x0d, 0xdf, 0xa2, 0xf6, 0x86, 0x5d, 0x70, 0x72, 0x0c, 0xb6, 0x62, 0x79, 0xc0,
	0xf2, 0x70, 0x16, 0x2f, 0xd0, 0xd9, 0x76, 0x0d, 0x7f, 0x8f, 0x82, 0x62, 0xf9, 0xfb, 0x92, 0x90,
	0xa7, 0x70, 0x30, 0x8d, 0x05, 0x16, 0x81, 0x92, 0x01, 0x47, 0x81, 0x0a, 0x9d, 0x9a, 0xee, 0xde,
	0xd4, 0xf8, 0xb3, 0x3c, 0xd3, 0x90, 0xf8, 0x60, 0x86, 0x32, 0x49, 0x58, 0xca, 0x0b, 0x67, 0xc7,
	0xad, 0xf9, 0x76, 0x77, 0xaf, 0x5d, 0x89, 0xe9, 0x25, 0x9c, 0x6e, 0xa2, 0xa4, 0x0d, 0x96, 0x90,
	0x51, 0x50, 0xa8, 0xa5, 0x40, 0xa7, 0xee, 0x1a, 0xbe, 0xdd, 0xbd, 0xb3, 0xb9, 0x3a, 0x90, 0xd1,
	0x68, 0x15, 0xa0, 0xa6, 0x58, 0x5b, 0xe4, 0x08, 0x60, 0x26, 0x55, 0x90, 0xa3, 0x90, 0x8c, 0x3b,
	0xbb, 0xae, 0xe1, 0x9b, 0xd4, 0x9a, 0x49, 0x45, 0x35, 0xf0, 0x86, 0x00, 0x03, 0x19, 0x7d, 0xc4,
	0xa2, 0x60, 0x11, 0x92, 0x67, 0x50, 0x17, 0xb8, 0x40, 0xa1, 0xb5, 0xee, 0xdf, 0x2c, 0x3c, 0x58,
	0x05, 0x68, 0x19, 0x27, 0x0e, 0x34, 0x92, 0x32, 0x67, 0x2d, 0xba, 0x72, 0xbd, 0x13, 0x68, 0x9e,
	0xb3, 0x58, 0x20, 0xa7, 0xf3, 0x74, 0xa4, 0x30, 0x23, 0x2d, 0xa8, 0x85, 0x49, 0xb5, 0xbd, 0x95,
	0x49, 0x1e, 0x82, 0x85, 0x3f, 0x62, 0x15, 0x84, 0x92, 0x97, 0xe9, 0x75, 0x6a, 0xae, 0x40, 0x4f,
	0x72, 0xf4, 0x7e, 0x1a, 0x70, 0xef, 0xd6, 0x1f, 0x92, 0x89, 0x25, 0x79, 0x0d, 0xf6, 0x4a, 0x78,
	0xd5, 0xd6, 0xd0, 0xd2, 0xef, 0x5e, 0x9f, 0x70, 0xad, 0x82, 0x82, 0xb8, 0x52, 0x74, 0x02, 0x07,
	0x53, 0x3d, 0x4e, 0x90, 0xcf, 0xd3, 0xa0, 0x50, 0x98, 0xe9, 0x8e, 0x76, 0xf7, 0x70, 0x93, 0x79,
	0x63, 0x5c, 0xda, 0x9c, 0x5e, 0x77, 0xbd, 0x4b, 0x30, 0xab, 0xa5, 0x92, 0x27, 0xb0, 0x1f, 0x4a,
	0x21, 0xf3, 0x22, 0xc0, 0x94, 0x4d, 0x04, 0x96, 0xa2, 0x4c, 0xda, 0x2c, 0x69, 0xbf, 0x84, 0x57,
	0x4b, 0xdc, 0xfe, 0xff, 0x12, 0x9f, 0xbf, 0xd0, 0xb5, 0x35, 0x22, 0x26, 0xec, 0x5c, 0x7c, 0x3a,
	0x1f, 0xb6, 0xb6, 0x88, 0x0d, 0x8d, 0x2f, 0x7d, 0x7a, 0x3a, 0x1c, 0xf5, 0x5b, 0x06, 0xb1, 0xa0,
	0x7e, 0xd6, 0x3f, 0x1d, 0x7f, 0x68, 0x6d, 0x77, 0xbf, 0x42, 0x63, 0x54, 0x16, 0x22, 0x63, 0x38,
	0xf8, 0x63, 0x45, 0xe4, 0x78, 0xd3, 0xe5, 0xef, 0x5f, 0xf3, 0x83, 0xa3, 0x7f, 0x5f, 0xc8, 0xc4,
	0xd2, 0xdb, 0x7a, 0x69, 0x9c, 0xbe, 0xbd, 0x7c, 0x13, 0xc5, 0x6a, 0x36, 0x9f, 0xb4, 0x43, 0x99,
	0x74, 0xbe, 0xc7, 0x29, 0x4f, 0x62, 0x21, 0x30, 0x8d, 0x3a, 0x2a, 0x16, 0xaa, 0x73, 0xeb, 0xb5,
	0xe9, 0x57, 0xf6, 0x4e, 0xff, 0x4e, 0x76, 0xf5, 0xf1, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x77, 0xf9, 0xa2, 0x42, 0x94, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncletClient is the client API for Synclet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncletClient interface {
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error)
}

type syncletClient struct {
	cc *grpc.ClientConn
}

func NewSyncletClient(cc *grpc.ClientConn) SyncletClient {
	return &syncletClient{cc}
}

func (c *syncletClient) UpdateContainer(ctx context.Context, in *UpdateContainerRequest, opts ...grpc.CallOption) (Synclet_UpdateContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synclet_serviceDesc.Streams[0], "/synclet.Synclet/UpdateContainer", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncletUpdateContainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synclet_UpdateContainerClient interface {
	Recv() (*UpdateContainerReply, error)
	grpc.ClientStream
}

type syncletUpdateContainerClient struct {
	grpc.ClientStream
}

func (x *syncletUpdateContainerClient) Recv() (*UpdateContainerReply, error) {
	m := new(UpdateContainerReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncletServer is the server API for Synclet service.
type SyncletServer interface {
	// updates the specified container and then restarts it
	// (much functionality packed into one rpc to minimize latency)
	UpdateContainer(*UpdateContainerRequest, Synclet_UpdateContainerServer) error
}

// UnimplementedSyncletServer can be embedded to have forward compatible implementations.
type UnimplementedSyncletServer struct {
}

func (*UnimplementedSyncletServer) UpdateContainer(req *UpdateContainerRequest, srv Synclet_UpdateContainerServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateContainer not implemented")
}

func RegisterSyncletServer(s *grpc.Server, srv SyncletServer) {
	s.RegisterService(&_Synclet_serviceDesc, srv)
}

func _Synclet_UpdateContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateContainerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SyncletServer).UpdateContainer(m, &syncletUpdateContainerServer{stream})
}

type Synclet_UpdateContainerServer interface {
	Send(*UpdateContainerReply) error
	grpc.ServerStream
}

type syncletUpdateContainerServer struct {
	grpc.ServerStream
}

func (x *syncletUpdateContainerServer) Send(m *UpdateContainerReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Synclet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synclet.Synclet",
	HandlerType: (*SyncletServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateContainer",
			Handler:       _Synclet_UpdateContainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/synclet/synclet.proto",
}
