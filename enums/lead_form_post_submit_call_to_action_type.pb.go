// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: google/ads/googleads/v9/enums/lead_form_post_submit_call_to_action_type.proto

package enums

import (
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

// Enum describing the type of post-submit call-to-action phrases for a lead
// form.
type LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType int32

const (
	// Not specified.
	LeadFormPostSubmitCallToActionTypeEnum_UNSPECIFIED LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 0
	// Used for return value only. Represents value unknown in this version.
	LeadFormPostSubmitCallToActionTypeEnum_UNKNOWN LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 1
	// Visit site.
	LeadFormPostSubmitCallToActionTypeEnum_VISIT_SITE LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 2
	// Download.
	LeadFormPostSubmitCallToActionTypeEnum_DOWNLOAD LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 3
	// Learn more.
	LeadFormPostSubmitCallToActionTypeEnum_LEARN_MORE LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 4
	// Shop now.
	LeadFormPostSubmitCallToActionTypeEnum_SHOP_NOW LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType = 5
)

// Enum value maps for LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType.
var (
	LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "VISIT_SITE",
		3: "DOWNLOAD",
		4: "LEARN_MORE",
		5: "SHOP_NOW",
	}
	LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"VISIT_SITE":  2,
		"DOWNLOAD":    3,
		"LEARN_MORE":  4,
		"SHOP_NOW":    5,
	}
)

func (x LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) Enum() *LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType {
	p := new(LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType)
	*p = x
	return p
}

func (x LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_lead_form_post_submit_call_to_action_type_proto_enumTypes[0].Descriptor()
}

func (LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) Type() protoreflect.EnumType {
	return &file_enums_lead_form_post_submit_call_to_action_type_proto_enumTypes[0]
}

func (x LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType.Descriptor instead.
func (LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType) EnumDescriptor() ([]byte, []int) {
	return file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescGZIP(), []int{0, 0}
}

// Describes the type of post-submit call-to-action phrases for a lead form.
type LeadFormPostSubmitCallToActionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeadFormPostSubmitCallToActionTypeEnum) Reset() {
	*x = LeadFormPostSubmitCallToActionTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_lead_form_post_submit_call_to_action_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadFormPostSubmitCallToActionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadFormPostSubmitCallToActionTypeEnum) ProtoMessage() {}

func (x *LeadFormPostSubmitCallToActionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_lead_form_post_submit_call_to_action_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadFormPostSubmitCallToActionTypeEnum.ProtoReflect.Descriptor instead.
func (*LeadFormPostSubmitCallToActionTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_lead_form_post_submit_call_to_action_type_proto protoreflect.FileDescriptor

var file_enums_lead_form_post_submit_call_to_action_type_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a,
	0x26, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x7e, 0x0a, 0x22, 0x4c, 0x65, 0x61, 0x64, 0x46,
	0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56,
	0x49, 0x53, 0x49, 0x54, 0x5f, 0x53, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x41,
	0x52, 0x4e, 0x5f, 0x4d, 0x4f, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x4f,
	0x50, 0x5f, 0x4e, 0x4f, 0x57, 0x10, 0x05, 0x42, 0xfc, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x27, 0x4c,
	0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescOnce sync.Once
	file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescData = file_enums_lead_form_post_submit_call_to_action_type_proto_rawDesc
)

func file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescGZIP() []byte {
	file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescOnce.Do(func() {
		file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescData)
	})
	return file_enums_lead_form_post_submit_call_to_action_type_proto_rawDescData
}

var file_enums_lead_form_post_submit_call_to_action_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_lead_form_post_submit_call_to_action_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_lead_form_post_submit_call_to_action_type_proto_goTypes = []interface{}{
	(LeadFormPostSubmitCallToActionTypeEnum_LeadFormPostSubmitCallToActionType)(0), // 0: google.ads.googleads.v9.enums.LeadFormPostSubmitCallToActionTypeEnum.LeadFormPostSubmitCallToActionType
	(*LeadFormPostSubmitCallToActionTypeEnum)(nil),                                 // 1: google.ads.googleads.v9.enums.LeadFormPostSubmitCallToActionTypeEnum
}
var file_enums_lead_form_post_submit_call_to_action_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_enums_lead_form_post_submit_call_to_action_type_proto_init()
}
func file_enums_lead_form_post_submit_call_to_action_type_proto_init() {
	if File_enums_lead_form_post_submit_call_to_action_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_lead_form_post_submit_call_to_action_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadFormPostSubmitCallToActionTypeEnum); i {
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
			RawDescriptor: file_enums_lead_form_post_submit_call_to_action_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_lead_form_post_submit_call_to_action_type_proto_goTypes,
		DependencyIndexes: file_enums_lead_form_post_submit_call_to_action_type_proto_depIdxs,
		EnumInfos:         file_enums_lead_form_post_submit_call_to_action_type_proto_enumTypes,
		MessageInfos:      file_enums_lead_form_post_submit_call_to_action_type_proto_msgTypes,
	}.Build()
	File_enums_lead_form_post_submit_call_to_action_type_proto = out.File
	file_enums_lead_form_post_submit_call_to_action_type_proto_rawDesc = nil
	file_enums_lead_form_post_submit_call_to_action_type_proto_goTypes = nil
	file_enums_lead_form_post_submit_call_to_action_type_proto_depIdxs = nil
}
