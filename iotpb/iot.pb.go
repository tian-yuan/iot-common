// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iot.proto

package iotpb

import (
	fmt "fmt"
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

///////////////////// message dispatcher service //////////////////////////////
type MessageHeader struct {
	Retain               bool     `protobuf:"varint,1,opt,name=retain,proto3" json:"retain,omitempty"`
	Qos                  int32    `protobuf:"varint,2,opt,name=qos,proto3" json:"qos,omitempty"`
	Dup                  bool     `protobuf:"varint,3,opt,name=dup,proto3" json:"dup,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageHeader) Reset()         { *m = MessageHeader{} }
func (m *MessageHeader) String() string { return proto.CompactTextString(m) }
func (*MessageHeader) ProtoMessage()    {}
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{0}
}

func (m *MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHeader.Unmarshal(m, b)
}
func (m *MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHeader.Marshal(b, m, deterministic)
}
func (m *MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHeader.Merge(m, src)
}
func (m *MessageHeader) XXX_Size() int {
	return xxx_messageInfo_MessageHeader.Size(m)
}
func (m *MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHeader proto.InternalMessageInfo

func (m *MessageHeader) GetRetain() bool {
	if m != nil {
		return m.Retain
	}
	return false
}

func (m *MessageHeader) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *MessageHeader) GetDup() bool {
	if m != nil {
		return m.Dup
	}
	return false
}

func (m *MessageHeader) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type SubscribeOptions struct {
	NoLocal              int32    `protobuf:"varint,1,opt,name=noLocal,proto3" json:"noLocal,omitempty"`
	RetainAsPublished    int32    `protobuf:"varint,2,opt,name=retainAsPublished,proto3" json:"retainAsPublished,omitempty"`
	RetainHandling       int32    `protobuf:"varint,3,opt,name=retainHandling,proto3" json:"retainHandling,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeOptions) Reset()         { *m = SubscribeOptions{} }
func (m *SubscribeOptions) String() string { return proto.CompactTextString(m) }
func (*SubscribeOptions) ProtoMessage()    {}
func (*SubscribeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{1}
}

func (m *SubscribeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeOptions.Unmarshal(m, b)
}
func (m *SubscribeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeOptions.Marshal(b, m, deterministic)
}
func (m *SubscribeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeOptions.Merge(m, src)
}
func (m *SubscribeOptions) XXX_Size() int {
	return xxx_messageInfo_SubscribeOptions.Size(m)
}
func (m *SubscribeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeOptions proto.InternalMessageInfo

func (m *SubscribeOptions) GetNoLocal() int32 {
	if m != nil {
		return m.NoLocal
	}
	return 0
}

func (m *SubscribeOptions) GetRetainAsPublished() int32 {
	if m != nil {
		return m.RetainAsPublished
	}
	return 0
}

func (m *SubscribeOptions) GetRetainHandling() int32 {
	if m != nil {
		return m.RetainHandling
	}
	return 0
}

type SubscribeMessageRequest struct {
	Header               *SubscribeOptions `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TopicFilter          string            `protobuf:"bytes,2,opt,name=topicFilter,proto3" json:"topicFilter,omitempty"`
	Qos                  int32             `protobuf:"varint,3,opt,name=qos,proto3" json:"qos,omitempty"`
	MsgId                int32             `protobuf:"varint,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Guid                 uint32            `protobuf:"varint,5,opt,name=guid,proto3" json:"guid,omitempty"`
	Properties           map[string]string `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SubscribeMessageRequest) Reset()         { *m = SubscribeMessageRequest{} }
func (m *SubscribeMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeMessageRequest) ProtoMessage()    {}
func (*SubscribeMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{2}
}

func (m *SubscribeMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeMessageRequest.Unmarshal(m, b)
}
func (m *SubscribeMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeMessageRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeMessageRequest.Merge(m, src)
}
func (m *SubscribeMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeMessageRequest.Size(m)
}
func (m *SubscribeMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeMessageRequest proto.InternalMessageInfo

func (m *SubscribeMessageRequest) GetHeader() *SubscribeOptions {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubscribeMessageRequest) GetTopicFilter() string {
	if m != nil {
		return m.TopicFilter
	}
	return ""
}

func (m *SubscribeMessageRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *SubscribeMessageRequest) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *SubscribeMessageRequest) GetGuid() uint32 {
	if m != nil {
		return m.Guid
	}
	return 0
}

func (m *SubscribeMessageRequest) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type SubscribeMessageResponse struct {
	MsgId                int32    `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	TopicId              uint32   `protobuf:"varint,2,opt,name=topicId,proto3" json:"topicId,omitempty"`
	Code                 int32    `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeMessageResponse) Reset()         { *m = SubscribeMessageResponse{} }
func (m *SubscribeMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeMessageResponse) ProtoMessage()    {}
func (*SubscribeMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{3}
}

func (m *SubscribeMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeMessageResponse.Unmarshal(m, b)
}
func (m *SubscribeMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeMessageResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeMessageResponse.Merge(m, src)
}
func (m *SubscribeMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeMessageResponse.Size(m)
}
func (m *SubscribeMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeMessageResponse proto.InternalMessageInfo

func (m *SubscribeMessageResponse) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *SubscribeMessageResponse) GetTopicId() uint32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *SubscribeMessageResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SubscribeMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UnSubscribeMessageRequest struct {
	TopicFilter          string   `protobuf:"bytes,1,opt,name=topicFilter,proto3" json:"topicFilter,omitempty"`
	Guid                 uint32   `protobuf:"varint,5,opt,name=guid,proto3" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnSubscribeMessageRequest) Reset()         { *m = UnSubscribeMessageRequest{} }
func (m *UnSubscribeMessageRequest) String() string { return proto.CompactTextString(m) }
func (*UnSubscribeMessageRequest) ProtoMessage()    {}
func (*UnSubscribeMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{4}
}

func (m *UnSubscribeMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnSubscribeMessageRequest.Unmarshal(m, b)
}
func (m *UnSubscribeMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnSubscribeMessageRequest.Marshal(b, m, deterministic)
}
func (m *UnSubscribeMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnSubscribeMessageRequest.Merge(m, src)
}
func (m *UnSubscribeMessageRequest) XXX_Size() int {
	return xxx_messageInfo_UnSubscribeMessageRequest.Size(m)
}
func (m *UnSubscribeMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnSubscribeMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnSubscribeMessageRequest proto.InternalMessageInfo

func (m *UnSubscribeMessageRequest) GetTopicFilter() string {
	if m != nil {
		return m.TopicFilter
	}
	return ""
}

func (m *UnSubscribeMessageRequest) GetGuid() uint32 {
	if m != nil {
		return m.Guid
	}
	return 0
}

type UnSubscribeMessageResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnSubscribeMessageResponse) Reset()         { *m = UnSubscribeMessageResponse{} }
func (m *UnSubscribeMessageResponse) String() string { return proto.CompactTextString(m) }
func (*UnSubscribeMessageResponse) ProtoMessage()    {}
func (*UnSubscribeMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{5}
}

func (m *UnSubscribeMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnSubscribeMessageResponse.Unmarshal(m, b)
}
func (m *UnSubscribeMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnSubscribeMessageResponse.Marshal(b, m, deterministic)
}
func (m *UnSubscribeMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnSubscribeMessageResponse.Merge(m, src)
}
func (m *UnSubscribeMessageResponse) XXX_Size() int {
	return xxx_messageInfo_UnSubscribeMessageResponse.Size(m)
}
func (m *UnSubscribeMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnSubscribeMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnSubscribeMessageResponse proto.InternalMessageInfo

func (m *UnSubscribeMessageResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UnSubscribeMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PublishMessageRequest struct {
	Header               *MessageHeader    `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	Topic                string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	MsgId                int32             `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MqttVersion          int32             `protobuf:"varint,3,opt,name=mqttVersion,proto3" json:"mqttVersion,omitempty"`
	Properties           map[string]string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishMessageRequest) Reset()         { *m = PublishMessageRequest{} }
func (m *PublishMessageRequest) String() string { return proto.CompactTextString(m) }
func (*PublishMessageRequest) ProtoMessage()    {}
func (*PublishMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{6}
}

func (m *PublishMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishMessageRequest.Unmarshal(m, b)
}
func (m *PublishMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishMessageRequest.Marshal(b, m, deterministic)
}
func (m *PublishMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishMessageRequest.Merge(m, src)
}
func (m *PublishMessageRequest) XXX_Size() int {
	return xxx_messageInfo_PublishMessageRequest.Size(m)
}
func (m *PublishMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishMessageRequest proto.InternalMessageInfo

func (m *PublishMessageRequest) GetHeader() *MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PublishMessageRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishMessageRequest) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *PublishMessageRequest) GetMqttVersion() int32 {
	if m != nil {
		return m.MqttVersion
	}
	return 0
}

func (m *PublishMessageRequest) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *PublishMessageRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PublishMessageResponse struct {
	MsgId                int32    `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishMessageResponse) Reset()         { *m = PublishMessageResponse{} }
func (m *PublishMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PublishMessageResponse) ProtoMessage()    {}
func (*PublishMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{7}
}

func (m *PublishMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishMessageResponse.Unmarshal(m, b)
}
func (m *PublishMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishMessageResponse.Marshal(b, m, deterministic)
}
func (m *PublishMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishMessageResponse.Merge(m, src)
}
func (m *PublishMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PublishMessageResponse.Size(m)
}
func (m *PublishMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishMessageResponse proto.InternalMessageInfo

func (m *PublishMessageResponse) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *PublishMessageResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PublishMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

///////////////////// topic manager service //////////////////////////////
type SubTopicIdRange struct {
	Begin                uint32   `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  uint32   `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTopicIdRange) Reset()         { *m = SubTopicIdRange{} }
func (m *SubTopicIdRange) String() string { return proto.CompactTextString(m) }
func (*SubTopicIdRange) ProtoMessage()    {}
func (*SubTopicIdRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{8}
}

func (m *SubTopicIdRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTopicIdRange.Unmarshal(m, b)
}
func (m *SubTopicIdRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTopicIdRange.Marshal(b, m, deterministic)
}
func (m *SubTopicIdRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTopicIdRange.Merge(m, src)
}
func (m *SubTopicIdRange) XXX_Size() int {
	return xxx_messageInfo_SubTopicIdRange.Size(m)
}
func (m *SubTopicIdRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTopicIdRange.DiscardUnknown(m)
}

var xxx_messageInfo_SubTopicIdRange proto.InternalMessageInfo

func (m *SubTopicIdRange) GetBegin() uint32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *SubTopicIdRange) GetEnd() uint32 {
	if m != nil {
		return m.End
	}
	return 0
}

type SubTopicLoadRequest struct {
	ProductKey           string             `protobuf:"bytes,1,opt,name=productKey,proto3" json:"productKey,omitempty"`
	SubTopicIdRange      []*SubTopicIdRange `protobuf:"bytes,2,rep,name=subTopicIdRange,proto3" json:"subTopicIdRange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SubTopicLoadRequest) Reset()         { *m = SubTopicLoadRequest{} }
func (m *SubTopicLoadRequest) String() string { return proto.CompactTextString(m) }
func (*SubTopicLoadRequest) ProtoMessage()    {}
func (*SubTopicLoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{9}
}

func (m *SubTopicLoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTopicLoadRequest.Unmarshal(m, b)
}
func (m *SubTopicLoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTopicLoadRequest.Marshal(b, m, deterministic)
}
func (m *SubTopicLoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTopicLoadRequest.Merge(m, src)
}
func (m *SubTopicLoadRequest) XXX_Size() int {
	return xxx_messageInfo_SubTopicLoadRequest.Size(m)
}
func (m *SubTopicLoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTopicLoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubTopicLoadRequest proto.InternalMessageInfo

func (m *SubTopicLoadRequest) GetProductKey() string {
	if m != nil {
		return m.ProductKey
	}
	return ""
}

func (m *SubTopicLoadRequest) GetSubTopicIdRange() []*SubTopicIdRange {
	if m != nil {
		return m.SubTopicIdRange
	}
	return nil
}

type SubTopicLoadResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTopicLoadResponse) Reset()         { *m = SubTopicLoadResponse{} }
func (m *SubTopicLoadResponse) String() string { return proto.CompactTextString(m) }
func (*SubTopicLoadResponse) ProtoMessage()    {}
func (*SubTopicLoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{10}
}

func (m *SubTopicLoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTopicLoadResponse.Unmarshal(m, b)
}
func (m *SubTopicLoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTopicLoadResponse.Marshal(b, m, deterministic)
}
func (m *SubTopicLoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTopicLoadResponse.Merge(m, src)
}
func (m *SubTopicLoadResponse) XXX_Size() int {
	return xxx_messageInfo_SubTopicLoadResponse.Size(m)
}
func (m *SubTopicLoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTopicLoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubTopicLoadResponse proto.InternalMessageInfo

func (m *SubTopicLoadResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SubTopicLoadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

///////////////////// registry service //////////////////////////////
type ConnectMessageRequest struct {
	Header               *MessageHeader `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	Username             string         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string         `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	WillRetain           bool           `protobuf:"varint,3,opt,name=willRetain,proto3" json:"willRetain,omitempty"`
	Will                 bool           `protobuf:"varint,4,opt,name=will,proto3" json:"will,omitempty"`
	CleanStart           bool           `protobuf:"varint,5,opt,name=cleanStart,proto3" json:"cleanStart,omitempty"`
	WillQoS              uint32         `protobuf:"varint,6,opt,name=willQoS,proto3" json:"willQoS,omitempty"`
	WillTopic            string         `protobuf:"bytes,7,opt,name=willTopic,proto3" json:"willTopic,omitempty"`
	ClientId             string         `protobuf:"bytes,8,opt,name=clientId,proto3" json:"clientId,omitempty"`
	WillMsg              string         `protobuf:"bytes,9,opt,name=willMsg,proto3" json:"willMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConnectMessageRequest) Reset()         { *m = ConnectMessageRequest{} }
func (m *ConnectMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectMessageRequest) ProtoMessage()    {}
func (*ConnectMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{11}
}

func (m *ConnectMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectMessageRequest.Unmarshal(m, b)
}
func (m *ConnectMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectMessageRequest.Marshal(b, m, deterministic)
}
func (m *ConnectMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectMessageRequest.Merge(m, src)
}
func (m *ConnectMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectMessageRequest.Size(m)
}
func (m *ConnectMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectMessageRequest proto.InternalMessageInfo

func (m *ConnectMessageRequest) GetHeader() *MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ConnectMessageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ConnectMessageRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ConnectMessageRequest) GetWillRetain() bool {
	if m != nil {
		return m.WillRetain
	}
	return false
}

func (m *ConnectMessageRequest) GetWill() bool {
	if m != nil {
		return m.Will
	}
	return false
}

func (m *ConnectMessageRequest) GetCleanStart() bool {
	if m != nil {
		return m.CleanStart
	}
	return false
}

func (m *ConnectMessageRequest) GetWillQoS() uint32 {
	if m != nil {
		return m.WillQoS
	}
	return 0
}

func (m *ConnectMessageRequest) GetWillTopic() string {
	if m != nil {
		return m.WillTopic
	}
	return ""
}

func (m *ConnectMessageRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ConnectMessageRequest) GetWillMsg() string {
	if m != nil {
		return m.WillMsg
	}
	return ""
}

type ConnectMessageResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectMessageResponse) Reset()         { *m = ConnectMessageResponse{} }
func (m *ConnectMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectMessageResponse) ProtoMessage()    {}
func (*ConnectMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{12}
}

func (m *ConnectMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectMessageResponse.Unmarshal(m, b)
}
func (m *ConnectMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectMessageResponse.Marshal(b, m, deterministic)
}
func (m *ConnectMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectMessageResponse.Merge(m, src)
}
func (m *ConnectMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectMessageResponse.Size(m)
}
func (m *ConnectMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectMessageResponse proto.InternalMessageInfo

func (m *ConnectMessageResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConnectMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageHeader)(nil), "iotpb.MessageHeader")
	proto.RegisterType((*SubscribeOptions)(nil), "iotpb.SubscribeOptions")
	proto.RegisterType((*SubscribeMessageRequest)(nil), "iotpb.SubscribeMessageRequest")
	proto.RegisterMapType((map[string]string)(nil), "iotpb.SubscribeMessageRequest.PropertiesEntry")
	proto.RegisterType((*SubscribeMessageResponse)(nil), "iotpb.SubscribeMessageResponse")
	proto.RegisterType((*UnSubscribeMessageRequest)(nil), "iotpb.UnSubscribeMessageRequest")
	proto.RegisterType((*UnSubscribeMessageResponse)(nil), "iotpb.UnSubscribeMessageResponse")
	proto.RegisterType((*PublishMessageRequest)(nil), "iotpb.PublishMessageRequest")
	proto.RegisterMapType((map[string]string)(nil), "iotpb.PublishMessageRequest.PropertiesEntry")
	proto.RegisterType((*PublishMessageResponse)(nil), "iotpb.PublishMessageResponse")
	proto.RegisterType((*SubTopicIdRange)(nil), "iotpb.SubTopicIdRange")
	proto.RegisterType((*SubTopicLoadRequest)(nil), "iotpb.SubTopicLoadRequest")
	proto.RegisterType((*SubTopicLoadResponse)(nil), "iotpb.SubTopicLoadResponse")
	proto.RegisterType((*ConnectMessageRequest)(nil), "iotpb.ConnectMessageRequest")
	proto.RegisterType((*ConnectMessageResponse)(nil), "iotpb.ConnectMessageResponse")
}

func init() { proto.RegisterFile("iot.proto", fileDescriptor_1804728d01c3c43d) }

var fileDescriptor_1804728d01c3c43d = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xef, 0xae, 0x63, 0xc7, 0x7e, 0x8e, 0x9b, 0x76, 0x48, 0xd3, 0xc5, 0x94, 0x62, 0xf6, 0x80,
	0x72, 0x88, 0x8c, 0x64, 0x2e, 0x80, 0x84, 0x04, 0xa2, 0xad, 0x6a, 0x48, 0xfa, 0x67, 0x52, 0x38,
	0x55, 0x82, 0xf1, 0xee, 0x68, 0x33, 0x62, 0x33, 0xb3, 0xd9, 0x99, 0x25, 0xf8, 0xca, 0x91, 0x2f,
	0x84, 0xc4, 0x85, 0xaf, 0xc0, 0x47, 0x42, 0xf3, 0x6f, 0xbd, 0xde, 0x38, 0x3e, 0xb4, 0x70, 0xe8,
	0xed, 0xbd, 0x37, 0xb3, 0xef, 0xfd, 0x7e, 0xbf, 0xf7, 0xde, 0xd8, 0x30, 0x60, 0x42, 0x4d, 0x8b,
	0x52, 0x28, 0x81, 0xba, 0x4c, 0xa8, 0x62, 0x11, 0xff, 0x04, 0xa3, 0x53, 0x2a, 0x25, 0xc9, 0xe8,
	0x53, 0x4a, 0x52, 0x5a, 0xa2, 0x43, 0xe8, 0x95, 0x54, 0x11, 0xc6, 0xa3, 0x60, 0x12, 0x1c, 0xf5,
	0xb1, 0xf3, 0xd0, 0x1d, 0xe8, 0x5c, 0x0a, 0x19, 0x85, 0x93, 0xe0, 0xa8, 0x8b, 0xb5, 0xa9, 0x23,
	0x69, 0x55, 0x44, 0x1d, 0x73, 0x4d, 0x9b, 0x08, 0xc1, 0x8e, 0x5a, 0x16, 0x34, 0xda, 0x99, 0x04,
	0x47, 0x23, 0x6c, 0xec, 0xf8, 0xf7, 0x00, 0xee, 0x9c, 0x55, 0x0b, 0x99, 0x94, 0x6c, 0x41, 0x9f,
	0x17, 0x8a, 0x09, 0x2e, 0x51, 0x04, 0xbb, 0x5c, 0x9c, 0x88, 0x84, 0xe4, 0xa6, 0x4a, 0x17, 0x7b,
	0x17, 0x1d, 0xc3, 0x5d, 0x5b, 0xf0, 0x1b, 0xf9, 0xa2, 0x5a, 0xe4, 0x4c, 0x9e, 0xd3, 0xd4, 0x15,
	0xbd, 0x7e, 0x80, 0x3e, 0x81, 0xdb, 0x36, 0xf8, 0x94, 0xf0, 0x34, 0x67, 0x3c, 0x33, 0x68, 0xba,
	0xb8, 0x15, 0x8d, 0xff, 0x0c, 0xe1, 0x7e, 0x0d, 0xc2, 0xf1, 0xc5, 0xf4, 0xb2, 0xa2, 0x52, 0xa1,
	0x4f, 0xa1, 0x77, 0x6e, 0xa8, 0x1b, 0x28, 0xc3, 0xd9, 0xfd, 0xa9, 0x51, 0x66, 0xda, 0x06, 0x8d,
	0xdd, 0x35, 0x34, 0x81, 0xa1, 0x12, 0x05, 0x4b, 0x9e, 0xb0, 0x5c, 0xd1, 0xd2, 0x80, 0x1b, 0xe0,
	0x66, 0xc8, 0x6b, 0xd5, 0x59, 0x69, 0x75, 0x00, 0xdd, 0x0b, 0x99, 0xcd, 0x53, 0x23, 0x4d, 0x17,
	0x5b, 0x47, 0xeb, 0x95, 0x55, 0x2c, 0x8d, 0xba, 0x56, 0x2f, 0x6d, 0xa3, 0x67, 0x00, 0x45, 0x29,
	0x0a, 0x5a, 0x2a, 0x46, 0x65, 0xd4, 0x9b, 0x74, 0x8e, 0x86, 0xb3, 0x69, 0x1b, 0xd2, 0x3a, 0x85,
	0xe9, 0x8b, 0xfa, 0x83, 0xc7, 0x5c, 0x95, 0x4b, 0xdc, 0xc8, 0x30, 0xfe, 0x0a, 0xf6, 0x5b, 0xc7,
	0x1a, 0xde, 0x2f, 0x74, 0x69, 0xe8, 0x0e, 0xb0, 0x36, 0x35, 0xbc, 0x5f, 0x49, 0x5e, 0x51, 0x47,
	0xc6, 0x3a, 0x5f, 0x86, 0x9f, 0x07, 0xf1, 0x6f, 0x10, 0x5d, 0xaf, 0x2a, 0x0b, 0xc1, 0x25, 0x5d,
	0x91, 0x0a, 0x9a, 0xa4, 0x22, 0xd8, 0x35, 0x5a, 0xcc, 0x6d, 0xdf, 0x46, 0xd8, 0xbb, 0x9a, 0x6e,
	0x22, 0x52, 0xea, 0x74, 0x31, 0xb6, 0xbe, 0x7d, 0x61, 0xd3, 0x1a, 0x69, 0x06, 0xd8, 0xbb, 0xf1,
	0x4b, 0x78, 0xff, 0x07, 0x7e, 0x53, 0xd3, 0x5a, 0x3d, 0x08, 0xae, 0xf7, 0x60, 0x83, 0xb6, 0xf1,
	0x77, 0x30, 0xde, 0x94, 0xd2, 0xd1, 0xf1, 0xf0, 0x82, 0xcd, 0xf0, 0xc2, 0x75, 0x78, 0x7f, 0x85,
	0x70, 0xcf, 0x0d, 0x62, 0x0b, 0xdb, 0x71, 0x3d, 0x50, 0xa9, 0x19, 0xa8, 0x03, 0xd7, 0xbd, 0xb5,
	0x3d, 0xab, 0xa7, 0xe9, 0x00, 0xba, 0x06, 0xb6, 0xe3, 0x60, 0x9d, 0x95, 0xb4, 0x61, 0x53, 0xda,
	0x09, 0x0c, 0x2f, 0x2e, 0x95, 0xfa, 0x91, 0x96, 0x92, 0x09, 0xee, 0x74, 0x6c, 0x86, 0xd0, 0xc9,
	0xda, 0xf4, 0xec, 0x98, 0xe9, 0x39, 0x76, 0xf5, 0x37, 0xa2, 0xdd, 0x36, 0x3b, 0x9a, 0x7d, 0x41,
	0x96, 0xb9, 0x20, 0x56, 0xc6, 0x3d, 0xec, 0xdd, 0xb7, 0x9d, 0xaa, 0xd7, 0x70, 0xd8, 0x46, 0xb3,
	0x75, 0xa6, 0x7c, 0x6b, 0xc2, 0xcd, 0xad, 0xe9, 0xac, 0xb7, 0xe6, 0x0b, 0xd8, 0x3f, 0xab, 0x16,
	0xaf, 0xec, 0xd4, 0x61, 0xc2, 0x33, 0x93, 0x76, 0x41, 0x33, 0xf7, 0xa8, 0x8d, 0xb0, 0x75, 0x34,
	0x64, 0xca, 0xfd, 0x98, 0x6a, 0x33, 0xbe, 0x82, 0xf7, 0xfc, 0xa7, 0x27, 0x82, 0xa4, 0xbe, 0xa5,
	0x0f, 0x8d, 0xac, 0x69, 0x95, 0xa8, 0xef, 0x6b, 0x8a, 0x8d, 0x08, 0xfa, 0x1a, 0xf6, 0xe5, 0x7a,
	0xc5, 0x28, 0x34, 0xda, 0x1f, 0xae, 0x36, 0xb7, 0x79, 0x8a, 0xdb, 0xd7, 0xe3, 0x47, 0x70, 0xb0,
	0x5e, 0xf8, 0x8d, 0x86, 0xf2, 0xef, 0x10, 0xee, 0x7d, 0x2b, 0x38, 0xa7, 0x89, 0x7a, 0xab, 0xa1,
	0x1c, 0x43, 0xbf, 0x92, 0xb4, 0xe4, 0xe4, 0x82, 0x3a, 0xb6, 0xb5, 0xaf, 0xcf, 0x0a, 0x22, 0xe5,
	0x95, 0x28, 0x53, 0x57, 0xbe, 0xf6, 0xb5, 0x4e, 0x57, 0x2c, 0xcf, 0xb1, 0xfd, 0x01, 0xb1, 0xbf,
	0x0c, 0x8d, 0x88, 0x66, 0xa3, 0x3d, 0xb3, 0xea, 0x7d, 0x6c, 0x6c, 0xfd, 0x4d, 0x92, 0x53, 0xc2,
	0xcf, 0x14, 0x29, 0x95, 0x99, 0xb3, 0x3e, 0x6e, 0x44, 0x34, 0x5b, 0x7d, 0xef, 0xa5, 0x38, 0x8b,
	0x7a, 0xf6, 0x3d, 0x71, 0x2e, 0x7a, 0x00, 0x03, 0x6d, 0x1a, 0xd1, 0xa2, 0x5d, 0x03, 0x65, 0x15,
	0xd0, 0x38, 0x93, 0x9c, 0x51, 0xae, 0xe6, 0x69, 0xd4, 0xb7, 0x38, 0xbd, 0xef, 0x73, 0x9e, 0xca,
	0x2c, 0x1a, 0x58, 0x05, 0x9d, 0x1b, 0x3f, 0x81, 0xc3, 0xb6, 0x80, 0x6f, 0xd2, 0x89, 0xd9, 0x1f,
	0x21, 0xdc, 0x75, 0x19, 0x1e, 0x31, 0x59, 0x10, 0x95, 0x9c, 0xd3, 0x12, 0x3d, 0x87, 0xdb, 0xeb,
	0x73, 0x8f, 0x1e, 0x6c, 0x5b, 0xce, 0xf1, 0x87, 0x37, 0x9c, 0x5a, 0x48, 0xf1, 0x2d, 0xf4, 0x0c,
	0x06, 0xf5, 0x7b, 0x86, 0x1e, 0x6e, 0xff, 0x99, 0x18, 0x7f, 0x74, 0xe3, 0x79, 0x9d, 0xef, 0x15,
	0x0c, 0x1b, 0x2f, 0x24, 0x9a, 0xb8, 0x2f, 0x6e, 0x7c, 0x88, 0xc7, 0x1f, 0x6f, 0xb9, 0xe1, 0xb3,
	0xce, 0x7e, 0x86, 0x91, 0x63, 0xf0, 0x98, 0x67, 0x8c, 0xd3, 0xff, 0x5c, 0x87, 0xd9, 0x3f, 0x21,
	0xec, 0x99, 0xb6, 0x9f, 0x12, 0x4e, 0xb2, 0xff, 0x43, 0xe9, 0x39, 0xec, 0xe9, 0xc5, 0xf4, 0x4b,
	0x8a, 0xc6, 0xad, 0xcd, 0x6e, 0x3c, 0x17, 0xe3, 0x0f, 0x36, 0x9e, 0xbd, 0x63, 0x4d, 0x7b, 0x0d,
	0xfb, 0x98, 0x66, 0x4c, 0xaa, 0x72, 0xe9, 0x45, 0x9d, 0x43, 0xdf, 0x87, 0x6a, 0x39, 0x37, 0x3e,
	0x37, 0xb5, 0x9c, 0x9b, 0x77, 0x29, 0xbe, 0xb5, 0xe8, 0x99, 0x7f, 0xa1, 0x9f, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0xe7, 0x56, 0x4c, 0x92, 0x0a, 0x00, 0x00,
}
