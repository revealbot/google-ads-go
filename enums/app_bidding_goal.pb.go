// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: google/ads/googleads/v19/enums/app_bidding_goal.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the goal towards which the bidding strategy, of an app
// campaign, should optimize for.
type AppBiddingGoalEnum_AppBiddingGoal int32

const (
	// Not specified.
	AppBiddingGoalEnum_UNSPECIFIED AppBiddingGoalEnum_AppBiddingGoal = 0
	// Represents value unknown in this version of the API.
	AppBiddingGoalEnum_UNKNOWN AppBiddingGoalEnum_AppBiddingGoal = 1
	// The bidding strategy of the app campaign should aim to maximize
	// installation of the app.
	AppBiddingGoalEnum_OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME AppBiddingGoalEnum_AppBiddingGoal = 2
	// The bidding strategy of the app campaign should aim to maximize
	// the selected in-app conversions' volume.
	AppBiddingGoalEnum_OPTIMIZE_FOR_IN_APP_CONVERSION_VOLUME AppBiddingGoalEnum_AppBiddingGoal = 3
	// The bidding strategy of the app campaign should aim to maximize
	// all conversions' value, that is, install and selected in-app conversions.
	AppBiddingGoalEnum_OPTIMIZE_FOR_TOTAL_CONVERSION_VALUE AppBiddingGoalEnum_AppBiddingGoal = 4
	// The bidding strategy of the app campaign should aim to maximize
	// just the selected in-app conversion's volume, while achieving or
	// exceeding target cost per in-app conversion.
	AppBiddingGoalEnum_OPTIMIZE_FOR_TARGET_IN_APP_CONVERSION AppBiddingGoalEnum_AppBiddingGoal = 5
	// The bidding strategy of the app campaign should aim to maximize
	// all conversions' value, that is, install and selected in-app conversions
	// while achieving or exceeding target return on advertising spend.
	AppBiddingGoalEnum_OPTIMIZE_FOR_RETURN_ON_ADVERTISING_SPEND AppBiddingGoalEnum_AppBiddingGoal = 6
	// This bidding strategy of the app campaign should aim to
	// maximize installation of the app without advertiser-provided target
	// cost-per-install.
	AppBiddingGoalEnum_OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME_WITHOUT_TARGET_CPI AppBiddingGoalEnum_AppBiddingGoal = 7
	// This bidding strategy of the app campaign should aim to
	// maximize pre-registration of the app.
	AppBiddingGoalEnum_OPTIMIZE_FOR_PRE_REGISTRATION_CONVERSION_VOLUME AppBiddingGoalEnum_AppBiddingGoal = 8
)

// Enum value maps for AppBiddingGoalEnum_AppBiddingGoal.
var (
	AppBiddingGoalEnum_AppBiddingGoal_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME",
		3: "OPTIMIZE_FOR_IN_APP_CONVERSION_VOLUME",
		4: "OPTIMIZE_FOR_TOTAL_CONVERSION_VALUE",
		5: "OPTIMIZE_FOR_TARGET_IN_APP_CONVERSION",
		6: "OPTIMIZE_FOR_RETURN_ON_ADVERTISING_SPEND",
		7: "OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME_WITHOUT_TARGET_CPI",
		8: "OPTIMIZE_FOR_PRE_REGISTRATION_CONVERSION_VOLUME",
	}
	AppBiddingGoalEnum_AppBiddingGoal_value = map[string]int32{
		"UNSPECIFIED":                                               0,
		"UNKNOWN":                                                   1,
		"OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME":                    2,
		"OPTIMIZE_FOR_IN_APP_CONVERSION_VOLUME":                     3,
		"OPTIMIZE_FOR_TOTAL_CONVERSION_VALUE":                       4,
		"OPTIMIZE_FOR_TARGET_IN_APP_CONVERSION":                     5,
		"OPTIMIZE_FOR_RETURN_ON_ADVERTISING_SPEND":                  6,
		"OPTIMIZE_FOR_INSTALL_CONVERSION_VOLUME_WITHOUT_TARGET_CPI": 7,
		"OPTIMIZE_FOR_PRE_REGISTRATION_CONVERSION_VOLUME":           8,
	}
)

func (x AppBiddingGoalEnum_AppBiddingGoal) Enum() *AppBiddingGoalEnum_AppBiddingGoal {
	p := new(AppBiddingGoalEnum_AppBiddingGoal)
	*p = x
	return p
}

func (x AppBiddingGoalEnum_AppBiddingGoal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppBiddingGoalEnum_AppBiddingGoal) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_app_bidding_goal_proto_enumTypes[0].Descriptor()
}

func (AppBiddingGoalEnum_AppBiddingGoal) Type() protoreflect.EnumType {
	return &file_enums_app_bidding_goal_proto_enumTypes[0]
}

func (x AppBiddingGoalEnum_AppBiddingGoal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppBiddingGoalEnum_AppBiddingGoal.Descriptor instead.
func (AppBiddingGoalEnum_AppBiddingGoal) EnumDescriptor() ([]byte, []int) {
	return file_enums_app_bidding_goal_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing an app bidding goal for raise Target CPA
// recommendation.
type AppBiddingGoalEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppBiddingGoalEnum) Reset() {
	*x = AppBiddingGoalEnum{}
	mi := &file_enums_app_bidding_goal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppBiddingGoalEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppBiddingGoalEnum) ProtoMessage() {}

func (x *AppBiddingGoalEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_app_bidding_goal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppBiddingGoalEnum.ProtoReflect.Descriptor instead.
func (*AppBiddingGoalEnum) Descriptor() ([]byte, []int) {
	return file_enums_app_bidding_goal_proto_rawDescGZIP(), []int{0}
}

var File_enums_app_bidding_goal_proto protoreflect.FileDescriptor

var file_enums_app_bidding_goal_proto_rawDesc = string([]byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x92, 0x03, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfb,
	0x02, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x61,
	0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x2a, 0x0a, 0x26, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x4f,
	0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x41,
	0x50, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x4f,
	0x4c, 0x55, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49,
	0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x04, 0x12,
	0x29, 0x0a, 0x25, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x2c, 0x0a, 0x28, 0x4f, 0x50,
	0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52,
	0x4e, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x3d, 0x0a, 0x39, 0x4f, 0x50, 0x54, 0x49,
	0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x4f, 0x4c, 0x55,
	0x4d, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x43, 0x50, 0x49, 0x10, 0x07, 0x12, 0x33, 0x0a, 0x2f, 0x4f, 0x50, 0x54, 0x49, 0x4d,
	0x49, 0x5a, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49,
	0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x08, 0x42, 0xed, 0x01, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x42, 0x13, 0x41, 0x70, 0x70, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x47,
	0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_app_bidding_goal_proto_rawDescOnce sync.Once
	file_enums_app_bidding_goal_proto_rawDescData []byte
)

func file_enums_app_bidding_goal_proto_rawDescGZIP() []byte {
	file_enums_app_bidding_goal_proto_rawDescOnce.Do(func() {
		file_enums_app_bidding_goal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_app_bidding_goal_proto_rawDesc), len(file_enums_app_bidding_goal_proto_rawDesc)))
	})
	return file_enums_app_bidding_goal_proto_rawDescData
}

var file_enums_app_bidding_goal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_app_bidding_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_app_bidding_goal_proto_goTypes = []any{
	(AppBiddingGoalEnum_AppBiddingGoal)(0), // 0: google.ads.googleads.v19.enums.AppBiddingGoalEnum.AppBiddingGoal
	(*AppBiddingGoalEnum)(nil),             // 1: google.ads.googleads.v19.enums.AppBiddingGoalEnum
}
var file_enums_app_bidding_goal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_app_bidding_goal_proto_init() }
func file_enums_app_bidding_goal_proto_init() {
	if File_enums_app_bidding_goal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_app_bidding_goal_proto_rawDesc), len(file_enums_app_bidding_goal_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_app_bidding_goal_proto_goTypes,
		DependencyIndexes: file_enums_app_bidding_goal_proto_depIdxs,
		EnumInfos:         file_enums_app_bidding_goal_proto_enumTypes,
		MessageInfos:      file_enums_app_bidding_goal_proto_msgTypes,
	}.Build()
	File_enums_app_bidding_goal_proto = out.File
	file_enums_app_bidding_goal_proto_goTypes = nil
	file_enums_app_bidding_goal_proto_depIdxs = nil
}
