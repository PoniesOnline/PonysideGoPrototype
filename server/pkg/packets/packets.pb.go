// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: packets.proto

package packets

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_packets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IdMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdMessage) Reset() {
	*x = IdMessage{}
	mi := &file_packets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMessage) ProtoMessage() {}

func (x *IdMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMessage.ProtoReflect.Descriptor instead.
func (*IdMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{1}
}

func (x *IdMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LoginRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequestMessage) Reset() {
	*x = LoginRequestMessage{}
	mi := &file_packets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequestMessage) ProtoMessage() {}

func (x *LoginRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequestMessage.ProtoReflect.Descriptor instead.
func (*LoginRequestMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequestMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequestMessage) Reset() {
	*x = RegisterRequestMessage{}
	mi := &file_packets_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequestMessage) ProtoMessage() {}

func (x *RegisterRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequestMessage.ProtoReflect.Descriptor instead.
func (*RegisterRequestMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequestMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type OkResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OkResponseMessage) Reset() {
	*x = OkResponseMessage{}
	mi := &file_packets_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OkResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkResponseMessage) ProtoMessage() {}

func (x *OkResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkResponseMessage.ProtoReflect.Descriptor instead.
func (*OkResponseMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{4}
}

type DenyResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *DenyResponseMessage) Reset() {
	*x = DenyResponseMessage{}
	mi := &file_packets_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenyResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenyResponseMessage) ProtoMessage() {}

func (x *DenyResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenyResponseMessage.ProtoReflect.Descriptor instead.
func (*DenyResponseMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{5}
}

func (x *DenyResponseMessage) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type PlayerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	X    float64 `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y    float64 `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *PlayerMessage) Reset() {
	*x = PlayerMessage{}
	mi := &file_packets_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMessage) ProtoMessage() {}

func (x *PlayerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMessage.ProtoReflect.Descriptor instead.
func (*PlayerMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{6}
}

func (x *PlayerMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerMessage) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerMessage) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PlayerInputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dx int32 `protobuf:"varint,1,opt,name=dx,proto3" json:"dx,omitempty"`
	Dy int32 `protobuf:"varint,2,opt,name=dy,proto3" json:"dy,omitempty"`
}

func (x *PlayerInputMessage) Reset() {
	*x = PlayerInputMessage{}
	mi := &file_packets_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerInputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInputMessage) ProtoMessage() {}

func (x *PlayerInputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInputMessage.ProtoReflect.Descriptor instead.
func (*PlayerInputMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{7}
}

func (x *PlayerInputMessage) GetDx() int32 {
	if x != nil {
		return x.Dx
	}
	return 0
}

func (x *PlayerInputMessage) GetDy() int32 {
	if x != nil {
		return x.Dy
	}
	return 0
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId uint64 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*Packet_Chat
	//	*Packet_Id
	//	*Packet_LoginRequest
	//	*Packet_RegisterRequest
	//	*Packet_OkResponse
	//	*Packet_DenyResponse
	//	*Packet_Player
	//	*Packet_PlayerInput
	Msg isPacket_Msg `protobuf_oneof:"msg"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	mi := &file_packets_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{8}
}

func (x *Packet) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (m *Packet) GetMsg() isPacket_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *Packet) GetChat() *ChatMessage {
	if x, ok := x.GetMsg().(*Packet_Chat); ok {
		return x.Chat
	}
	return nil
}

func (x *Packet) GetId() *IdMessage {
	if x, ok := x.GetMsg().(*Packet_Id); ok {
		return x.Id
	}
	return nil
}

func (x *Packet) GetLoginRequest() *LoginRequestMessage {
	if x, ok := x.GetMsg().(*Packet_LoginRequest); ok {
		return x.LoginRequest
	}
	return nil
}

func (x *Packet) GetRegisterRequest() *RegisterRequestMessage {
	if x, ok := x.GetMsg().(*Packet_RegisterRequest); ok {
		return x.RegisterRequest
	}
	return nil
}

func (x *Packet) GetOkResponse() *OkResponseMessage {
	if x, ok := x.GetMsg().(*Packet_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

func (x *Packet) GetDenyResponse() *DenyResponseMessage {
	if x, ok := x.GetMsg().(*Packet_DenyResponse); ok {
		return x.DenyResponse
	}
	return nil
}

func (x *Packet) GetPlayer() *PlayerMessage {
	if x, ok := x.GetMsg().(*Packet_Player); ok {
		return x.Player
	}
	return nil
}

func (x *Packet) GetPlayerInput() *PlayerInputMessage {
	if x, ok := x.GetMsg().(*Packet_PlayerInput); ok {
		return x.PlayerInput
	}
	return nil
}

type isPacket_Msg interface {
	isPacket_Msg()
}

type Packet_Chat struct {
	Chat *ChatMessage `protobuf:"bytes,2,opt,name=chat,proto3,oneof"`
}

type Packet_Id struct {
	Id *IdMessage `protobuf:"bytes,3,opt,name=id,proto3,oneof"`
}

type Packet_LoginRequest struct {
	LoginRequest *LoginRequestMessage `protobuf:"bytes,4,opt,name=login_request,json=loginRequest,proto3,oneof"`
}

type Packet_RegisterRequest struct {
	RegisterRequest *RegisterRequestMessage `protobuf:"bytes,5,opt,name=register_request,json=registerRequest,proto3,oneof"`
}

type Packet_OkResponse struct {
	OkResponse *OkResponseMessage `protobuf:"bytes,6,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

type Packet_DenyResponse struct {
	DenyResponse *DenyResponseMessage `protobuf:"bytes,7,opt,name=deny_response,json=denyResponse,proto3,oneof"`
}

type Packet_Player struct {
	Player *PlayerMessage `protobuf:"bytes,8,opt,name=player,proto3,oneof"`
}

type Packet_PlayerInput struct {
	PlayerInput *PlayerInputMessage `protobuf:"bytes,9,opt,name=player_input,json=playerInput,proto3,oneof"`
}

func (*Packet_Chat) isPacket_Msg() {}

func (*Packet_Id) isPacket_Msg() {}

func (*Packet_LoginRequest) isPacket_Msg() {}

func (*Packet_RegisterRequest) isPacket_Msg() {}

func (*Packet_OkResponse) isPacket_Msg() {}

func (*Packet_DenyResponse) isPacket_Msg() {}

func (*Packet_Player) isPacket_Msg() {}

func (*Packet_PlayerInput) isPacket_Msg() {}

var File_packets_proto protoreflect.FileDescriptor

var file_packets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x13,
	0x44, 0x65, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x34, 0x0a, 0x12,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x64, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x64, 0x79, 0x22, 0x89, 0x04, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4c, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x0b, 0x6f, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x4f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x44, 0x65, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0d,
	0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packets_proto_rawDescOnce sync.Once
	file_packets_proto_rawDescData = file_packets_proto_rawDesc
)

func file_packets_proto_rawDescGZIP() []byte {
	file_packets_proto_rawDescOnce.Do(func() {
		file_packets_proto_rawDescData = protoimpl.X.CompressGZIP(file_packets_proto_rawDescData)
	})
	return file_packets_proto_rawDescData
}

var file_packets_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_packets_proto_goTypes = []any{
	(*ChatMessage)(nil),            // 0: packets.ChatMessage
	(*IdMessage)(nil),              // 1: packets.IdMessage
	(*LoginRequestMessage)(nil),    // 2: packets.LoginRequestMessage
	(*RegisterRequestMessage)(nil), // 3: packets.RegisterRequestMessage
	(*OkResponseMessage)(nil),      // 4: packets.OkResponseMessage
	(*DenyResponseMessage)(nil),    // 5: packets.DenyResponseMessage
	(*PlayerMessage)(nil),          // 6: packets.PlayerMessage
	(*PlayerInputMessage)(nil),     // 7: packets.PlayerInputMessage
	(*Packet)(nil),                 // 8: packets.Packet
}
var file_packets_proto_depIdxs = []int32{
	0, // 0: packets.Packet.chat:type_name -> packets.ChatMessage
	1, // 1: packets.Packet.id:type_name -> packets.IdMessage
	2, // 2: packets.Packet.login_request:type_name -> packets.LoginRequestMessage
	3, // 3: packets.Packet.register_request:type_name -> packets.RegisterRequestMessage
	4, // 4: packets.Packet.ok_response:type_name -> packets.OkResponseMessage
	5, // 5: packets.Packet.deny_response:type_name -> packets.DenyResponseMessage
	6, // 6: packets.Packet.player:type_name -> packets.PlayerMessage
	7, // 7: packets.Packet.player_input:type_name -> packets.PlayerInputMessage
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_packets_proto_init() }
func file_packets_proto_init() {
	if File_packets_proto != nil {
		return
	}
	file_packets_proto_msgTypes[8].OneofWrappers = []any{
		(*Packet_Chat)(nil),
		(*Packet_Id)(nil),
		(*Packet_LoginRequest)(nil),
		(*Packet_RegisterRequest)(nil),
		(*Packet_OkResponse)(nil),
		(*Packet_DenyResponse)(nil),
		(*Packet_Player)(nil),
		(*Packet_PlayerInput)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packets_proto_goTypes,
		DependencyIndexes: file_packets_proto_depIdxs,
		MessageInfos:      file_packets_proto_msgTypes,
	}.Build()
	File_packets_proto = out.File
	file_packets_proto_rawDesc = nil
	file_packets_proto_goTypes = nil
	file_packets_proto_depIdxs = nil
}
