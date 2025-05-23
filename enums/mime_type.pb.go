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
// source: google/ads/googleads/v19/enums/mime_type.proto

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

// The mime type
type MimeTypeEnum_MimeType int32

const (
	// The mime type has not been specified.
	MimeTypeEnum_UNSPECIFIED MimeTypeEnum_MimeType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	MimeTypeEnum_UNKNOWN MimeTypeEnum_MimeType = 1
	// MIME type of image/jpeg.
	MimeTypeEnum_IMAGE_JPEG MimeTypeEnum_MimeType = 2
	// MIME type of image/gif.
	MimeTypeEnum_IMAGE_GIF MimeTypeEnum_MimeType = 3
	// MIME type of image/png.
	MimeTypeEnum_IMAGE_PNG MimeTypeEnum_MimeType = 4
	// MIME type of application/x-shockwave-flash.
	MimeTypeEnum_FLASH MimeTypeEnum_MimeType = 5
	// MIME type of text/html.
	MimeTypeEnum_TEXT_HTML MimeTypeEnum_MimeType = 6
	// MIME type of application/pdf.
	MimeTypeEnum_PDF MimeTypeEnum_MimeType = 7
	// MIME type of application/msword.
	MimeTypeEnum_MSWORD MimeTypeEnum_MimeType = 8
	// MIME type of application/vnd.ms-excel.
	MimeTypeEnum_MSEXCEL MimeTypeEnum_MimeType = 9
	// MIME type of application/rtf.
	MimeTypeEnum_RTF MimeTypeEnum_MimeType = 10
	// MIME type of audio/wav.
	MimeTypeEnum_AUDIO_WAV MimeTypeEnum_MimeType = 11
	// MIME type of audio/mp3.
	MimeTypeEnum_AUDIO_MP3 MimeTypeEnum_MimeType = 12
	// MIME type of application/x-html5-ad-zip.
	MimeTypeEnum_HTML5_AD_ZIP MimeTypeEnum_MimeType = 13
)

// Enum value maps for MimeTypeEnum_MimeType.
var (
	MimeTypeEnum_MimeType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "IMAGE_JPEG",
		3:  "IMAGE_GIF",
		4:  "IMAGE_PNG",
		5:  "FLASH",
		6:  "TEXT_HTML",
		7:  "PDF",
		8:  "MSWORD",
		9:  "MSEXCEL",
		10: "RTF",
		11: "AUDIO_WAV",
		12: "AUDIO_MP3",
		13: "HTML5_AD_ZIP",
	}
	MimeTypeEnum_MimeType_value = map[string]int32{
		"UNSPECIFIED":  0,
		"UNKNOWN":      1,
		"IMAGE_JPEG":   2,
		"IMAGE_GIF":    3,
		"IMAGE_PNG":    4,
		"FLASH":        5,
		"TEXT_HTML":    6,
		"PDF":          7,
		"MSWORD":       8,
		"MSEXCEL":      9,
		"RTF":          10,
		"AUDIO_WAV":    11,
		"AUDIO_MP3":    12,
		"HTML5_AD_ZIP": 13,
	}
)

func (x MimeTypeEnum_MimeType) Enum() *MimeTypeEnum_MimeType {
	p := new(MimeTypeEnum_MimeType)
	*p = x
	return p
}

func (x MimeTypeEnum_MimeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MimeTypeEnum_MimeType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_mime_type_proto_enumTypes[0].Descriptor()
}

func (MimeTypeEnum_MimeType) Type() protoreflect.EnumType {
	return &file_enums_mime_type_proto_enumTypes[0]
}

func (x MimeTypeEnum_MimeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MimeTypeEnum_MimeType.Descriptor instead.
func (MimeTypeEnum_MimeType) EnumDescriptor() ([]byte, []int) {
	return file_enums_mime_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the mime types.
type MimeTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MimeTypeEnum) Reset() {
	*x = MimeTypeEnum{}
	mi := &file_enums_mime_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MimeTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MimeTypeEnum) ProtoMessage() {}

func (x *MimeTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_mime_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MimeTypeEnum.ProtoReflect.Descriptor instead.
func (*MimeTypeEnum) Descriptor() ([]byte, []int) {
	return file_enums_mime_type_proto_rawDescGZIP(), []int{0}
}

var File_enums_mime_type_proto protoreflect.FileDescriptor

var file_enums_mime_type_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x4a, 0x50, 0x45, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x47, 0x49, 0x46, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c,
	0x41, 0x53, 0x48, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x48, 0x54,
	0x4d, 0x4c, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x44, 0x46, 0x10, 0x07, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x53, 0x45,
	0x58, 0x43, 0x45, 0x4c, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x54, 0x46, 0x10, 0x0a, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x57, 0x41, 0x56, 0x10, 0x0b, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x50, 0x33, 0x10, 0x0c, 0x12, 0x10, 0x0a,
	0x0c, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x41, 0x44, 0x5f, 0x5a, 0x49, 0x50, 0x10, 0x0d, 0x42,
	0xe7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0d, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_enums_mime_type_proto_rawDescOnce sync.Once
	file_enums_mime_type_proto_rawDescData []byte
)

func file_enums_mime_type_proto_rawDescGZIP() []byte {
	file_enums_mime_type_proto_rawDescOnce.Do(func() {
		file_enums_mime_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_mime_type_proto_rawDesc), len(file_enums_mime_type_proto_rawDesc)))
	})
	return file_enums_mime_type_proto_rawDescData
}

var file_enums_mime_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_mime_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_mime_type_proto_goTypes = []any{
	(MimeTypeEnum_MimeType)(0), // 0: google.ads.googleads.v19.enums.MimeTypeEnum.MimeType
	(*MimeTypeEnum)(nil),       // 1: google.ads.googleads.v19.enums.MimeTypeEnum
}
var file_enums_mime_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_mime_type_proto_init() }
func file_enums_mime_type_proto_init() {
	if File_enums_mime_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_mime_type_proto_rawDesc), len(file_enums_mime_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_mime_type_proto_goTypes,
		DependencyIndexes: file_enums_mime_type_proto_depIdxs,
		EnumInfos:         file_enums_mime_type_proto_enumTypes,
		MessageInfos:      file_enums_mime_type_proto_msgTypes,
	}.Build()
	File_enums_mime_type_proto = out.File
	file_enums_mime_type_proto_goTypes = nil
	file_enums_mime_type_proto_depIdxs = nil
}
