// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: invitation_types.proto

package invitation_types

import (
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

type Invitation struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	Id                *uint64                `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	InviterIdentity   *entity_types.Identity `protobuf:"bytes,2,req,name=inviter_identity,json=inviterIdentity" json:"inviter_identity,omitempty"`
	InviteeIdentity   *entity_types.Identity `protobuf:"bytes,3,req,name=invitee_identity,json=inviteeIdentity" json:"invitee_identity,omitempty"`
	InviterName       *string                `protobuf:"bytes,4,opt,name=inviter_name,json=inviterName" json:"inviter_name,omitempty"`
	InviteeName       *string                `protobuf:"bytes,5,opt,name=invitee_name,json=inviteeName" json:"invitee_name,omitempty"`
	InvitationMessage *string                `protobuf:"bytes,6,opt,name=invitation_message,json=invitationMessage" json:"invitation_message,omitempty"`
	CreationTime      *uint64                `protobuf:"varint,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	ExpirationTime    *uint64                `protobuf:"varint,8,opt,name=expiration_time,json=expirationTime" json:"expiration_time,omitempty"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{0}
}

func (x *Invitation) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Invitation) GetInviterIdentity() *entity_types.Identity {
	if x != nil {
		return x.InviterIdentity
	}
	return nil
}

func (x *Invitation) GetInviteeIdentity() *entity_types.Identity {
	if x != nil {
		return x.InviteeIdentity
	}
	return nil
}

func (x *Invitation) GetInviterName() string {
	if x != nil && x.InviterName != nil {
		return *x.InviterName
	}
	return ""
}

func (x *Invitation) GetInviteeName() string {
	if x != nil && x.InviteeName != nil {
		return *x.InviteeName
	}
	return ""
}

func (x *Invitation) GetInvitationMessage() string {
	if x != nil && x.InvitationMessage != nil {
		return *x.InvitationMessage
	}
	return ""
}

func (x *Invitation) GetCreationTime() uint64 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *Invitation) GetExpirationTime() uint64 {
	if x != nil && x.ExpirationTime != nil {
		return *x.ExpirationTime
	}
	return 0
}

type InvitationSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId          *entity_types.EntityId `protobuf:"bytes,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	SuggesterId        *entity_types.EntityId `protobuf:"bytes,2,req,name=suggester_id,json=suggesterId" json:"suggester_id,omitempty"`
	SuggesteeId        *entity_types.EntityId `protobuf:"bytes,3,req,name=suggestee_id,json=suggesteeId" json:"suggestee_id,omitempty"`
	SuggesterName      *string                `protobuf:"bytes,4,opt,name=suggester_name,json=suggesterName" json:"suggester_name,omitempty"`
	SuggesteeName      *string                `protobuf:"bytes,5,opt,name=suggestee_name,json=suggesteeName" json:"suggestee_name,omitempty"`
	SuggesterAccountId *entity_types.EntityId `protobuf:"bytes,6,opt,name=suggester_account_id,json=suggesterAccountId" json:"suggester_account_id,omitempty"`
}

func (x *InvitationSuggestion) Reset() {
	*x = InvitationSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationSuggestion) ProtoMessage() {}

func (x *InvitationSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationSuggestion.ProtoReflect.Descriptor instead.
func (*InvitationSuggestion) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{1}
}

func (x *InvitationSuggestion) GetChannelId() *entity_types.EntityId {
	if x != nil {
		return x.ChannelId
	}
	return nil
}

func (x *InvitationSuggestion) GetSuggesterId() *entity_types.EntityId {
	if x != nil {
		return x.SuggesterId
	}
	return nil
}

func (x *InvitationSuggestion) GetSuggesteeId() *entity_types.EntityId {
	if x != nil {
		return x.SuggesteeId
	}
	return nil
}

func (x *InvitationSuggestion) GetSuggesterName() string {
	if x != nil && x.SuggesterName != nil {
		return *x.SuggesterName
	}
	return ""
}

func (x *InvitationSuggestion) GetSuggesteeName() string {
	if x != nil && x.SuggesteeName != nil {
		return *x.SuggesteeName
	}
	return ""
}

func (x *InvitationSuggestion) GetSuggesterAccountId() *entity_types.EntityId {
	if x != nil {
		return x.SuggesterAccountId
	}
	return nil
}

type InvitationTarget struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	Identity  *entity_types.Identity `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Email     *string                `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	BattleTag *string                `protobuf:"bytes,3,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
}

func (x *InvitationTarget) Reset() {
	*x = InvitationTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationTarget) ProtoMessage() {}

func (x *InvitationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationTarget.ProtoReflect.Descriptor instead.
func (*InvitationTarget) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{2}
}

func (x *InvitationTarget) GetIdentity() *entity_types.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *InvitationTarget) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *InvitationTarget) GetBattleTag() string {
	if x != nil && x.BattleTag != nil {
		return *x.BattleTag
	}
	return ""
}

type InvitationParams struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	InvitationMessage *string `protobuf:"bytes,1,opt,name=invitation_message,json=invitationMessage" json:"invitation_message,omitempty"`
	ExpirationTime    *uint64 `protobuf:"varint,2,opt,name=expiration_time,json=expirationTime,def=0" json:"expiration_time,omitempty"`
}

// Default values for InvitationParams fields.
const (
	Default_InvitationParams_ExpirationTime = uint64(0)
)

func (x *InvitationParams) Reset() {
	*x = InvitationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationParams) ProtoMessage() {}

func (x *InvitationParams) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationParams.ProtoReflect.Descriptor instead.
func (*InvitationParams) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{3}
}

func (x *InvitationParams) GetInvitationMessage() string {
	if x != nil && x.InvitationMessage != nil {
		return *x.InvitationMessage
	}
	return ""
}

func (x *InvitationParams) GetExpirationTime() uint64 {
	if x != nil && x.ExpirationTime != nil {
		return *x.ExpirationTime
	}
	return Default_InvitationParams_ExpirationTime
}

type SendInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentIdentity *entity_types.Identity `protobuf:"bytes,1,opt,name=agent_identity,json=agentIdentity" json:"agent_identity,omitempty"`
	// Deprecated: Marked as deprecated in invitation_types.proto.
	TargetId  *entity_types.EntityId    `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	Params    *InvitationParams         `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
	AgentInfo *entity_types.AccountInfo `protobuf:"bytes,4,opt,name=agent_info,json=agentInfo" json:"agent_info,omitempty"`
	Target    *InvitationTarget         `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
}

func (x *SendInvitationRequest) Reset() {
	*x = SendInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendInvitationRequest) ProtoMessage() {}

func (x *SendInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendInvitationRequest.ProtoReflect.Descriptor instead.
func (*SendInvitationRequest) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{4}
}

func (x *SendInvitationRequest) GetAgentIdentity() *entity_types.Identity {
	if x != nil {
		return x.AgentIdentity
	}
	return nil
}

// Deprecated: Marked as deprecated in invitation_types.proto.
func (x *SendInvitationRequest) GetTargetId() *entity_types.EntityId {
	if x != nil {
		return x.TargetId
	}
	return nil
}

func (x *SendInvitationRequest) GetParams() *InvitationParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SendInvitationRequest) GetAgentInfo() *entity_types.AccountInfo {
	if x != nil {
		return x.AgentInfo
	}
	return nil
}

func (x *SendInvitationRequest) GetTarget() *InvitationTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

type SendInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitation *Invitation `protobuf:"bytes,2,opt,name=invitation" json:"invitation,omitempty"`
}

func (x *SendInvitationResponse) Reset() {
	*x = SendInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendInvitationResponse) ProtoMessage() {}

func (x *SendInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendInvitationResponse.ProtoReflect.Descriptor instead.
func (*SendInvitationResponse) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{5}
}

func (x *SendInvitationResponse) GetInvitation() *Invitation {
	if x != nil {
		return x.Invitation
	}
	return nil
}

type UpdateInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentIdentity *entity_types.Identity `protobuf:"bytes,1,opt,name=agent_identity,json=agentIdentity" json:"agent_identity,omitempty"`
	InvitationId  *uint64                `protobuf:"fixed64,2,req,name=invitation_id,json=invitationId" json:"invitation_id,omitempty"`
	Params        *InvitationParams      `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
}

func (x *UpdateInvitationRequest) Reset() {
	*x = UpdateInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvitationRequest) ProtoMessage() {}

func (x *UpdateInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvitationRequest.ProtoReflect.Descriptor instead.
func (*UpdateInvitationRequest) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateInvitationRequest) GetAgentIdentity() *entity_types.Identity {
	if x != nil {
		return x.AgentIdentity
	}
	return nil
}

func (x *UpdateInvitationRequest) GetInvitationId() uint64 {
	if x != nil && x.InvitationId != nil {
		return *x.InvitationId
	}
	return 0
}

func (x *UpdateInvitationRequest) GetParams() *InvitationParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type GenericInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId      *entity_types.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	TargetId     *entity_types.EntityId `protobuf:"bytes,2,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	InvitationId *uint64                `protobuf:"fixed64,3,req,name=invitation_id,json=invitationId" json:"invitation_id,omitempty"`
	InviteeName  *string                `protobuf:"bytes,4,opt,name=invitee_name,json=inviteeName" json:"invitee_name,omitempty"`
	InviterName  *string                `protobuf:"bytes,5,opt,name=inviter_name,json=inviterName" json:"inviter_name,omitempty"`
	PreviousRole []uint32               `protobuf:"varint,6,rep,packed,name=previous_role,json=previousRole" json:"previous_role,omitempty"`
	DesiredRole  []uint32               `protobuf:"varint,7,rep,packed,name=desired_role,json=desiredRole" json:"desired_role,omitempty"`
	Reason       *uint32                `protobuf:"varint,8,opt,name=reason" json:"reason,omitempty"`
}

func (x *GenericInvitationRequest) Reset() {
	*x = GenericInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericInvitationRequest) ProtoMessage() {}

func (x *GenericInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericInvitationRequest.ProtoReflect.Descriptor instead.
func (*GenericInvitationRequest) Descriptor() ([]byte, []int) {
	return file_invitation_types_proto_rawDescGZIP(), []int{7}
}

func (x *GenericInvitationRequest) GetAgentId() *entity_types.EntityId {
	if x != nil {
		return x.AgentId
	}
	return nil
}

func (x *GenericInvitationRequest) GetTargetId() *entity_types.EntityId {
	if x != nil {
		return x.TargetId
	}
	return nil
}

func (x *GenericInvitationRequest) GetInvitationId() uint64 {
	if x != nil && x.InvitationId != nil {
		return *x.InvitationId
	}
	return 0
}

func (x *GenericInvitationRequest) GetInviteeName() string {
	if x != nil && x.InviteeName != nil {
		return *x.InviteeName
	}
	return ""
}

func (x *GenericInvitationRequest) GetInviterName() string {
	if x != nil && x.InviterName != nil {
		return *x.InviterName
	}
	return ""
}

func (x *GenericInvitationRequest) GetPreviousRole() []uint32 {
	if x != nil {
		return x.PreviousRole
	}
	return nil
}

func (x *GenericInvitationRequest) GetDesiredRole() []uint32 {
	if x != nil {
		return x.DesiredRole
	}
	return nil
}

func (x *GenericInvitationRequest) GetReason() uint32 {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return 0
}

var File_invitation_types_proto protoreflect.FileDescriptor

var file_invitation_types_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x12, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x0a, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x06, 0x52, 0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x10, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x10,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x2a, 0x05, 0x08, 0x64, 0x10, 0x90, 0x4e, 0x22, 0xdb, 0x02, 0x0a, 0x14, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x65, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x14, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x52, 0x12, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x61, 0x67, 0x2a, 0x05, 0x08, 0x64, 0x10, 0x90, 0x4e, 0x22, 0x74, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2a, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x05, 0x08, 0x64, 0x10,
	0x90, 0x4e, 0x22, 0xb9, 0x02, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x38, 0x0a, 0x0a,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x52,
	0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d,
	0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x06, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x18, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x06, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0d, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x42, 0x52, 0x0a, 0x0d, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x42, 0x14, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x01, 0x5a, 0x29, 0x62, 0x6e, 0x65,
	0x74, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_invitation_types_proto_rawDescOnce sync.Once
	file_invitation_types_proto_rawDescData = file_invitation_types_proto_rawDesc
)

func file_invitation_types_proto_rawDescGZIP() []byte {
	file_invitation_types_proto_rawDescOnce.Do(func() {
		file_invitation_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_invitation_types_proto_rawDescData)
	})
	return file_invitation_types_proto_rawDescData
}

var file_invitation_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_invitation_types_proto_goTypes = []interface{}{
	(*Invitation)(nil),               // 0: bgs.protocol.Invitation
	(*InvitationSuggestion)(nil),     // 1: bgs.protocol.InvitationSuggestion
	(*InvitationTarget)(nil),         // 2: bgs.protocol.InvitationTarget
	(*InvitationParams)(nil),         // 3: bgs.protocol.InvitationParams
	(*SendInvitationRequest)(nil),    // 4: bgs.protocol.SendInvitationRequest
	(*SendInvitationResponse)(nil),   // 5: bgs.protocol.SendInvitationResponse
	(*UpdateInvitationRequest)(nil),  // 6: bgs.protocol.UpdateInvitationRequest
	(*GenericInvitationRequest)(nil), // 7: bgs.protocol.GenericInvitationRequest
	(*entity_types.Identity)(nil),    // 8: bgs.protocol.Identity
	(*entity_types.EntityId)(nil),    // 9: bgs.protocol.EntityId
	(*entity_types.AccountInfo)(nil), // 10: bgs.protocol.AccountInfo
}
var file_invitation_types_proto_depIdxs = []int32{
	8,  // 0: bgs.protocol.Invitation.inviter_identity:type_name -> bgs.protocol.Identity
	8,  // 1: bgs.protocol.Invitation.invitee_identity:type_name -> bgs.protocol.Identity
	9,  // 2: bgs.protocol.InvitationSuggestion.channel_id:type_name -> bgs.protocol.EntityId
	9,  // 3: bgs.protocol.InvitationSuggestion.suggester_id:type_name -> bgs.protocol.EntityId
	9,  // 4: bgs.protocol.InvitationSuggestion.suggestee_id:type_name -> bgs.protocol.EntityId
	9,  // 5: bgs.protocol.InvitationSuggestion.suggester_account_id:type_name -> bgs.protocol.EntityId
	8,  // 6: bgs.protocol.InvitationTarget.identity:type_name -> bgs.protocol.Identity
	8,  // 7: bgs.protocol.SendInvitationRequest.agent_identity:type_name -> bgs.protocol.Identity
	9,  // 8: bgs.protocol.SendInvitationRequest.target_id:type_name -> bgs.protocol.EntityId
	3,  // 9: bgs.protocol.SendInvitationRequest.params:type_name -> bgs.protocol.InvitationParams
	10, // 10: bgs.protocol.SendInvitationRequest.agent_info:type_name -> bgs.protocol.AccountInfo
	2,  // 11: bgs.protocol.SendInvitationRequest.target:type_name -> bgs.protocol.InvitationTarget
	0,  // 12: bgs.protocol.SendInvitationResponse.invitation:type_name -> bgs.protocol.Invitation
	8,  // 13: bgs.protocol.UpdateInvitationRequest.agent_identity:type_name -> bgs.protocol.Identity
	3,  // 14: bgs.protocol.UpdateInvitationRequest.params:type_name -> bgs.protocol.InvitationParams
	9,  // 15: bgs.protocol.GenericInvitationRequest.agent_id:type_name -> bgs.protocol.EntityId
	9,  // 16: bgs.protocol.GenericInvitationRequest.target_id:type_name -> bgs.protocol.EntityId
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_invitation_types_proto_init() }
func file_invitation_types_proto_init() {
	if File_invitation_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invitation_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_invitation_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationSuggestion); i {
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
		file_invitation_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationTarget); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_invitation_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_invitation_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendInvitationRequest); i {
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
		file_invitation_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendInvitationResponse); i {
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
		file_invitation_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInvitationRequest); i {
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
		file_invitation_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericInvitationRequest); i {
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
			RawDescriptor: file_invitation_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_invitation_types_proto_goTypes,
		DependencyIndexes: file_invitation_types_proto_depIdxs,
		MessageInfos:      file_invitation_types_proto_msgTypes,
	}.Build()
	File_invitation_types_proto = out.File
	file_invitation_types_proto_rawDesc = nil
	file_invitation_types_proto_goTypes = nil
	file_invitation_types_proto_depIdxs = nil
}
