// Copyright 2022 Google LLC
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
// 	protoc        v3.21.7
// source: google/ads/googleads/v12/enums/device.proto

package enums

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

// Enumerates Google Ads devices available for targeting.
type DeviceEnum_Device int32

const (
	// Not specified.
	DeviceEnum_UNSPECIFIED DeviceEnum_Device = 0
	// The value is unknown in this version.
	DeviceEnum_UNKNOWN DeviceEnum_Device = 1
	// Mobile devices with full browsers.
	DeviceEnum_MOBILE DeviceEnum_Device = 2
	// Tablets with full browsers.
	DeviceEnum_TABLET DeviceEnum_Device = 3
	// Computers.
	DeviceEnum_DESKTOP DeviceEnum_Device = 4
	// Smart TVs and game consoles.
	DeviceEnum_CONNECTED_TV DeviceEnum_Device = 6
	// Other device types.
	DeviceEnum_OTHER DeviceEnum_Device = 5
)

// Enum value maps for DeviceEnum_Device.
var (
	DeviceEnum_Device_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "MOBILE",
		3: "TABLET",
		4: "DESKTOP",
		6: "CONNECTED_TV",
		5: "OTHER",
	}
	DeviceEnum_Device_value = map[string]int32{
		"UNSPECIFIED":  0,
		"UNKNOWN":      1,
		"MOBILE":       2,
		"TABLET":       3,
		"DESKTOP":      4,
		"CONNECTED_TV": 6,
		"OTHER":        5,
	}
)

func (x DeviceEnum_Device) Enum() *DeviceEnum_Device {
	p := new(DeviceEnum_Device)
	*p = x
	return p
}

func (x DeviceEnum_Device) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceEnum_Device) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_device_proto_enumTypes[0].Descriptor()
}

func (DeviceEnum_Device) Type() protoreflect.EnumType {
	return &file_enums_device_proto_enumTypes[0]
}

func (x DeviceEnum_Device) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceEnum_Device.Descriptor instead.
func (DeviceEnum_Device) EnumDescriptor() ([]byte, []int) {
	return file_enums_device_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enumeration of Google Ads devices available for targeting.
type DeviceEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeviceEnum) Reset() {
	*x = DeviceEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceEnum) ProtoMessage() {}

func (x *DeviceEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceEnum.ProtoReflect.Descriptor instead.
func (*DeviceEnum) Descriptor() ([]byte, []int) {
	return file_enums_device_proto_rawDescGZIP(), []int{0}
}

var File_enums_device_proto protoreflect.FileDescriptor

var file_enums_device_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x76, 0x0a,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x68, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x54, 0x56, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x05, 0x42, 0xe5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0b, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_device_proto_rawDescOnce sync.Once
	file_enums_device_proto_rawDescData = file_enums_device_proto_rawDesc
)

func file_enums_device_proto_rawDescGZIP() []byte {
	file_enums_device_proto_rawDescOnce.Do(func() {
		file_enums_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_device_proto_rawDescData)
	})
	return file_enums_device_proto_rawDescData
}

var file_enums_device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_device_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_device_proto_goTypes = []interface{}{
	(DeviceEnum_Device)(0), // 0: google.ads.googleads.v12.enums.DeviceEnum.Device
	(*DeviceEnum)(nil),     // 1: google.ads.googleads.v12.enums.DeviceEnum
}
var file_enums_device_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_device_proto_init() }
func file_enums_device_proto_init() {
	if File_enums_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceEnum); i {
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
			RawDescriptor: file_enums_device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_device_proto_goTypes,
		DependencyIndexes: file_enums_device_proto_depIdxs,
		EnumInfos:         file_enums_device_proto_enumTypes,
		MessageInfos:      file_enums_device_proto_msgTypes,
	}.Build()
	File_enums_device_proto = out.File
	file_enums_device_proto_rawDesc = nil
	file_enums_device_proto_goTypes = nil
	file_enums_device_proto_depIdxs = nil
}
