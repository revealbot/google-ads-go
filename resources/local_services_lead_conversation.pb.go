// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/ads/googleads/v17/resources/local_services_lead_conversation.proto

package resources

import (
	enums "github.com/revealbot/google-ads-go/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Data from Local Services Lead Conversation.
// Contains details of Lead Conversation which is generated when user calls,
// messages or books service from advertiser. These are appended to a Lead.
// More info: https://ads.google.com/local-services-ads
type LocalServicesLeadConversation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the local services lead conversation
	// data. Local Services Lead Conversation resource name have the form
	//
	// `customers/{customer_id}/localServicesLeadConversation/{local_services_lead_conversation_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this Lead Conversation.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Type of GLS lead conversation, EMAIL, MESSAGE, PHONE_CALL,
	// SMS, etc.
	ConversationChannel enums.LocalServicesLeadConversationTypeEnum_ConversationType `protobuf:"varint,3,opt,name=conversation_channel,json=conversationChannel,proto3,enum=google.ads.googleads.v17.enums.LocalServicesLeadConversationTypeEnum_ConversationType" json:"conversation_channel,omitempty"`
	// Output only. Type of participant in the lead conversation, ADVERTISER or
	// CONSUMER.
	ParticipantType enums.LocalServicesParticipantTypeEnum_ParticipantType `protobuf:"varint,4,opt,name=participant_type,json=participantType,proto3,enum=google.ads.googleads.v17.enums.LocalServicesParticipantTypeEnum_ParticipantType" json:"participant_type,omitempty"`
	// Output only. Resource name of Lead associated to the Lead Conversation.
	Lead string `protobuf:"bytes,5,opt,name=lead,proto3" json:"lead,omitempty"`
	// Output only. The date time at which lead conversation was created by Local
	// Services Ads. The format is "YYYY-MM-DD HH:MM:SS" in the Google Ads
	// account's timezone. Examples: "2018-03-05 09:15:00" or "2018-02-01
	// 14:34:30"
	EventDateTime string `protobuf:"bytes,6,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
	// Output only. Details of phone call conversation in case of PHONE_CALL.
	PhoneCallDetails *PhoneCallDetails `protobuf:"bytes,7,opt,name=phone_call_details,json=phoneCallDetails,proto3,oneof" json:"phone_call_details,omitempty"`
	// Output only. Details of message conversation in case of EMAIL, MESSAGE or
	// SMS.
	MessageDetails *MessageDetails `protobuf:"bytes,8,opt,name=message_details,json=messageDetails,proto3,oneof" json:"message_details,omitempty"`
}

func (x *LocalServicesLeadConversation) Reset() {
	*x = LocalServicesLeadConversation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_local_services_lead_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalServicesLeadConversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesLeadConversation) ProtoMessage() {}

func (x *LocalServicesLeadConversation) ProtoReflect() protoreflect.Message {
	mi := &file_resources_local_services_lead_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesLeadConversation.ProtoReflect.Descriptor instead.
func (*LocalServicesLeadConversation) Descriptor() ([]byte, []int) {
	return file_resources_local_services_lead_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *LocalServicesLeadConversation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *LocalServicesLeadConversation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LocalServicesLeadConversation) GetConversationChannel() enums.LocalServicesLeadConversationTypeEnum_ConversationType {
	if x != nil {
		return x.ConversationChannel
	}
	return enums.LocalServicesLeadConversationTypeEnum_ConversationType(0)
}

func (x *LocalServicesLeadConversation) GetParticipantType() enums.LocalServicesParticipantTypeEnum_ParticipantType {
	if x != nil {
		return x.ParticipantType
	}
	return enums.LocalServicesParticipantTypeEnum_ParticipantType(0)
}

func (x *LocalServicesLeadConversation) GetLead() string {
	if x != nil {
		return x.Lead
	}
	return ""
}

func (x *LocalServicesLeadConversation) GetEventDateTime() string {
	if x != nil {
		return x.EventDateTime
	}
	return ""
}

func (x *LocalServicesLeadConversation) GetPhoneCallDetails() *PhoneCallDetails {
	if x != nil {
		return x.PhoneCallDetails
	}
	return nil
}

func (x *LocalServicesLeadConversation) GetMessageDetails() *MessageDetails {
	if x != nil {
		return x.MessageDetails
	}
	return nil
}

// Represents details of a phone call conversation.
type PhoneCallDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The duration (in milliseconds) of the phone call (end to end).
	CallDurationMillis int64 `protobuf:"varint,1,opt,name=call_duration_millis,json=callDurationMillis,proto3" json:"call_duration_millis,omitempty"`
	// Output only. URL to the call recording audio file.
	CallRecordingUrl string `protobuf:"bytes,2,opt,name=call_recording_url,json=callRecordingUrl,proto3" json:"call_recording_url,omitempty"`
}

func (x *PhoneCallDetails) Reset() {
	*x = PhoneCallDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_local_services_lead_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneCallDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneCallDetails) ProtoMessage() {}

func (x *PhoneCallDetails) ProtoReflect() protoreflect.Message {
	mi := &file_resources_local_services_lead_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneCallDetails.ProtoReflect.Descriptor instead.
func (*PhoneCallDetails) Descriptor() ([]byte, []int) {
	return file_resources_local_services_lead_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *PhoneCallDetails) GetCallDurationMillis() int64 {
	if x != nil {
		return x.CallDurationMillis
	}
	return 0
}

func (x *PhoneCallDetails) GetCallRecordingUrl() string {
	if x != nil {
		return x.CallRecordingUrl
	}
	return ""
}

// Represents details of text message in case of email, message or SMS.
type MessageDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Textual content of the message.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Output only. URL to the SMS or email attachments. These URLs can be used to
	// download the contents of the attachment by using the developer token.
	AttachmentUrls []string `protobuf:"bytes,2,rep,name=attachment_urls,json=attachmentUrls,proto3" json:"attachment_urls,omitempty"`
}

func (x *MessageDetails) Reset() {
	*x = MessageDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_local_services_lead_conversation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDetails) ProtoMessage() {}

func (x *MessageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_resources_local_services_lead_conversation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDetails.ProtoReflect.Descriptor instead.
func (*MessageDetails) Descriptor() ([]byte, []int) {
	return file_resources_local_services_lead_conversation_proto_rawDescGZIP(), []int{2}
}

func (x *MessageDetails) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MessageDetails) GetAttachmentUrls() []string {
	if x != nil {
		return x.AttachmentUrls
	}
	return nil
}

var File_resources_local_services_lead_conversation_proto protoreflect.FileDescriptor

var file_resources_local_services_lead_conversation_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x07, 0x0a, 0x1d, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3e, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x38, 0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x8e, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x56, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c, 0x0a, 0x2a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x64,
	0x12, 0x2b, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6c, 0x0a,
	0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x10, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x12, 0x65, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x88,
	0x01, 0x01, 0x3a, 0x9a, 0x01, 0xea, 0x41, 0x96, 0x01, 0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x7c, 0x0a, 0x10, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x35,
	0x0a, 0x14, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x57, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c,
	0x73, 0x42, 0x94, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x22, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea,
	0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_local_services_lead_conversation_proto_rawDescOnce sync.Once
	file_resources_local_services_lead_conversation_proto_rawDescData = file_resources_local_services_lead_conversation_proto_rawDesc
)

func file_resources_local_services_lead_conversation_proto_rawDescGZIP() []byte {
	file_resources_local_services_lead_conversation_proto_rawDescOnce.Do(func() {
		file_resources_local_services_lead_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_local_services_lead_conversation_proto_rawDescData)
	})
	return file_resources_local_services_lead_conversation_proto_rawDescData
}

var file_resources_local_services_lead_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_local_services_lead_conversation_proto_goTypes = []interface{}{
	(*LocalServicesLeadConversation)(nil),                             // 0: google.ads.googleads.v17.resources.LocalServicesLeadConversation
	(*PhoneCallDetails)(nil),                                          // 1: google.ads.googleads.v17.resources.PhoneCallDetails
	(*MessageDetails)(nil),                                            // 2: google.ads.googleads.v17.resources.MessageDetails
	(enums.LocalServicesLeadConversationTypeEnum_ConversationType)(0), // 3: google.ads.googleads.v17.enums.LocalServicesLeadConversationTypeEnum.ConversationType
	(enums.LocalServicesParticipantTypeEnum_ParticipantType)(0),       // 4: google.ads.googleads.v17.enums.LocalServicesParticipantTypeEnum.ParticipantType
}
var file_resources_local_services_lead_conversation_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v17.resources.LocalServicesLeadConversation.conversation_channel:type_name -> google.ads.googleads.v17.enums.LocalServicesLeadConversationTypeEnum.ConversationType
	4, // 1: google.ads.googleads.v17.resources.LocalServicesLeadConversation.participant_type:type_name -> google.ads.googleads.v17.enums.LocalServicesParticipantTypeEnum.ParticipantType
	1, // 2: google.ads.googleads.v17.resources.LocalServicesLeadConversation.phone_call_details:type_name -> google.ads.googleads.v17.resources.PhoneCallDetails
	2, // 3: google.ads.googleads.v17.resources.LocalServicesLeadConversation.message_details:type_name -> google.ads.googleads.v17.resources.MessageDetails
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_resources_local_services_lead_conversation_proto_init() }
func file_resources_local_services_lead_conversation_proto_init() {
	if File_resources_local_services_lead_conversation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_local_services_lead_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalServicesLeadConversation); i {
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
		file_resources_local_services_lead_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneCallDetails); i {
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
		file_resources_local_services_lead_conversation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDetails); i {
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
	file_resources_local_services_lead_conversation_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_local_services_lead_conversation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_local_services_lead_conversation_proto_goTypes,
		DependencyIndexes: file_resources_local_services_lead_conversation_proto_depIdxs,
		MessageInfos:      file_resources_local_services_lead_conversation_proto_msgTypes,
	}.Build()
	File_resources_local_services_lead_conversation_proto = out.File
	file_resources_local_services_lead_conversation_proto_rawDesc = nil
	file_resources_local_services_lead_conversation_proto_goTypes = nil
	file_resources_local_services_lead_conversation_proto_depIdxs = nil
}
