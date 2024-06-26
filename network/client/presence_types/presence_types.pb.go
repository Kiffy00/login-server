// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: presence_types.proto

package presence_types

import (
	attribute_types "bnet-mock/network/client/attribute_types"
	channel_types "bnet-mock/network/client/channel_types"
	entity_types "bnet-mock/network/client/entity_types"
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

type FieldOperation_OperationType int32

const (
	FieldOperation_SET   FieldOperation_OperationType = 0
	FieldOperation_CLEAR FieldOperation_OperationType = 1
)

// Enum value maps for FieldOperation_OperationType.
var (
	FieldOperation_OperationType_name = map[int32]string{
		0: "SET",
		1: "CLEAR",
	}
	FieldOperation_OperationType_value = map[string]int32{
		"SET":   0,
		"CLEAR": 1,
	}
)

func (x FieldOperation_OperationType) Enum() *FieldOperation_OperationType {
	p := new(FieldOperation_OperationType)
	*p = x
	return p
}

func (x FieldOperation_OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldOperation_OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_presence_types_proto_enumTypes[0].Descriptor()
}

func (FieldOperation_OperationType) Type() protoreflect.EnumType {
	return &file_presence_types_proto_enumTypes[0]
}

func (x FieldOperation_OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FieldOperation_OperationType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FieldOperation_OperationType(num)
	return nil
}

// Deprecated: Use FieldOperation_OperationType.Descriptor instead.
func (FieldOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{3, 0}
}

type RichPresenceLocalizationKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program        *uint32 `protobuf:"fixed32,1,req,name=program" json:"program,omitempty"`
	Stream         *uint32 `protobuf:"fixed32,2,req,name=stream" json:"stream,omitempty"`
	LocalizationId *uint32 `protobuf:"varint,3,req,name=localization_id,json=localizationId" json:"localization_id,omitempty"`
}

func (x *RichPresenceLocalizationKey) Reset() {
	*x = RichPresenceLocalizationKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RichPresenceLocalizationKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RichPresenceLocalizationKey) ProtoMessage() {}

func (x *RichPresenceLocalizationKey) ProtoReflect() protoreflect.Message {
	mi := &file_presence_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RichPresenceLocalizationKey.ProtoReflect.Descriptor instead.
func (*RichPresenceLocalizationKey) Descriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{0}
}

func (x *RichPresenceLocalizationKey) GetProgram() uint32 {
	if x != nil && x.Program != nil {
		return *x.Program
	}
	return 0
}

func (x *RichPresenceLocalizationKey) GetStream() uint32 {
	if x != nil && x.Stream != nil {
		return *x.Stream
	}
	return 0
}

func (x *RichPresenceLocalizationKey) GetLocalizationId() uint32 {
	if x != nil && x.LocalizationId != nil {
		return *x.LocalizationId
	}
	return 0
}

type FieldKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program  *uint32 `protobuf:"varint,1,req,name=program" json:"program,omitempty"`
	Group    *uint32 `protobuf:"varint,2,req,name=group" json:"group,omitempty"`
	Field    *uint32 `protobuf:"varint,3,req,name=field" json:"field,omitempty"`
	UniqueId *uint64 `protobuf:"varint,4,opt,name=unique_id,json=uniqueId,def=0" json:"unique_id,omitempty"`
}

// Default values for FieldKey fields.
const (
	Default_FieldKey_UniqueId = uint64(0)
)

func (x *FieldKey) Reset() {
	*x = FieldKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldKey) ProtoMessage() {}

func (x *FieldKey) ProtoReflect() protoreflect.Message {
	mi := &file_presence_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldKey.ProtoReflect.Descriptor instead.
func (*FieldKey) Descriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{1}
}

func (x *FieldKey) GetProgram() uint32 {
	if x != nil && x.Program != nil {
		return *x.Program
	}
	return 0
}

func (x *FieldKey) GetGroup() uint32 {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return 0
}

func (x *FieldKey) GetField() uint32 {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return 0
}

func (x *FieldKey) GetUniqueId() uint64 {
	if x != nil && x.UniqueId != nil {
		return *x.UniqueId
	}
	return Default_FieldKey_UniqueId
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *FieldKey                `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *attribute_types.Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_presence_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetKey() *FieldKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Field) GetValue() *attribute_types.Variant {
	if x != nil {
		return x.Value
	}
	return nil
}

type FieldOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field     *Field                        `protobuf:"bytes,1,req,name=field" json:"field,omitempty"`
	Operation *FieldOperation_OperationType `protobuf:"varint,2,opt,name=operation,enum=bgs.protocol.presence.v1.FieldOperation_OperationType,def=0" json:"operation,omitempty"`
}

// Default values for FieldOperation fields.
const (
	Default_FieldOperation_Operation = FieldOperation_SET
)

func (x *FieldOperation) Reset() {
	*x = FieldOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOperation) ProtoMessage() {}

func (x *FieldOperation) ProtoReflect() protoreflect.Message {
	mi := &file_presence_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOperation.ProtoReflect.Descriptor instead.
func (*FieldOperation) Descriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{3}
}

func (x *FieldOperation) GetField() *Field {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *FieldOperation) GetOperation() FieldOperation_OperationType {
	if x != nil && x.Operation != nil {
		return *x.Operation
	}
	return Default_FieldOperation_Operation
}

type ChannelState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId       *entity_types.EntityId `protobuf:"bytes,1,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation []*FieldOperation      `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	Healing        *bool                  `protobuf:"varint,3,opt,name=healing,def=0" json:"healing,omitempty"`
}

// Default values for ChannelState fields.
const (
	Default_ChannelState_Healing = bool(false)
)

func (x *ChannelState) Reset() {
	*x = ChannelState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelState) ProtoMessage() {}

func (x *ChannelState) ProtoReflect() protoreflect.Message {
	mi := &file_presence_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelState.ProtoReflect.Descriptor instead.
func (*ChannelState) Descriptor() ([]byte, []int) {
	return file_presence_types_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelState) GetEntityId() *entity_types.EntityId {
	if x != nil {
		return x.EntityId
	}
	return nil
}

func (x *ChannelState) GetFieldOperation() []*FieldOperation {
	if x != nil {
		return x.FieldOperation
	}
	return nil
}

func (x *ChannelState) GetHealing() bool {
	if x != nil && x.Healing != nil {
		return *x.Healing
	}
	return Default_ChannelState_Healing
}

var file_presence_types_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*channel_types.ChannelState)(nil),
		ExtensionType: (*ChannelState)(nil),
		Field:         101,
		Name:          "bgs.protocol.presence.v1.ChannelState.presence",
		Tag:           "bytes,101,opt,name=presence",
		Filename:      "presence_types.proto",
	},
}

// Extension fields to channel_types.ChannelState.
var (
	// optional bgs.protocol.presence.v1.ChannelState presence = 101;
	E_ChannelState_Presence = &file_presence_types_proto_extTypes[0]
)

var File_presence_types_proto protoreflect.FileDescriptor

var file_presence_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x15, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x78, 0x0a, 0x1b, 0x52, 0x69, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x07,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x02, 0x28, 0x07, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x08, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x09,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x34, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x59, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x03, 0x53,
	0x45, 0x54, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x45, 0x41, 0x52,
	0x10, 0x01, 0x22, 0xa2, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x32, 0x69, 0x0a, 0x08,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x2b, 0x48, 0x01, 0x5a, 0x27, 0x62, 0x6e, 0x65,
	0x74, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73,
}

var (
	file_presence_types_proto_rawDescOnce sync.Once
	file_presence_types_proto_rawDescData = file_presence_types_proto_rawDesc
)

func file_presence_types_proto_rawDescGZIP() []byte {
	file_presence_types_proto_rawDescOnce.Do(func() {
		file_presence_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_presence_types_proto_rawDescData)
	})
	return file_presence_types_proto_rawDescData
}

var file_presence_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_presence_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_presence_types_proto_goTypes = []interface{}{
	(FieldOperation_OperationType)(0),   // 0: bgs.protocol.presence.v1.FieldOperation.OperationType
	(*RichPresenceLocalizationKey)(nil), // 1: bgs.protocol.presence.v1.RichPresenceLocalizationKey
	(*FieldKey)(nil),                    // 2: bgs.protocol.presence.v1.FieldKey
	(*Field)(nil),                       // 3: bgs.protocol.presence.v1.Field
	(*FieldOperation)(nil),              // 4: bgs.protocol.presence.v1.FieldOperation
	(*ChannelState)(nil),                // 5: bgs.protocol.presence.v1.ChannelState
	(*attribute_types.Variant)(nil),     // 6: bgs.protocol.Variant
	(*entity_types.EntityId)(nil),       // 7: bgs.protocol.EntityId
	(*channel_types.ChannelState)(nil),  // 8: bgs.protocol.channel.v1.ChannelState
}
var file_presence_types_proto_depIdxs = []int32{
	2, // 0: bgs.protocol.presence.v1.Field.key:type_name -> bgs.protocol.presence.v1.FieldKey
	6, // 1: bgs.protocol.presence.v1.Field.value:type_name -> bgs.protocol.Variant
	3, // 2: bgs.protocol.presence.v1.FieldOperation.field:type_name -> bgs.protocol.presence.v1.Field
	0, // 3: bgs.protocol.presence.v1.FieldOperation.operation:type_name -> bgs.protocol.presence.v1.FieldOperation.OperationType
	7, // 4: bgs.protocol.presence.v1.ChannelState.entity_id:type_name -> bgs.protocol.EntityId
	4, // 5: bgs.protocol.presence.v1.ChannelState.field_operation:type_name -> bgs.protocol.presence.v1.FieldOperation
	8, // 6: bgs.protocol.presence.v1.ChannelState.presence:extendee -> bgs.protocol.channel.v1.ChannelState
	5, // 7: bgs.protocol.presence.v1.ChannelState.presence:type_name -> bgs.protocol.presence.v1.ChannelState
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	7, // [7:8] is the sub-list for extension type_name
	6, // [6:7] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_presence_types_proto_init() }
func file_presence_types_proto_init() {
	if File_presence_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_presence_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RichPresenceLocalizationKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_presence_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_presence_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_presence_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_presence_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_presence_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_presence_types_proto_goTypes,
		DependencyIndexes: file_presence_types_proto_depIdxs,
		EnumInfos:         file_presence_types_proto_enumTypes,
		MessageInfos:      file_presence_types_proto_msgTypes,
		ExtensionInfos:    file_presence_types_proto_extTypes,
	}.Build()
	File_presence_types_proto = out.File
	file_presence_types_proto_rawDesc = nil
	file_presence_types_proto_goTypes = nil
	file_presence_types_proto_depIdxs = nil
}
