// Copyright 2023 Google LLC
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
// 	protoc        v3.21.12
// source: google/ads/googleads/v14/resources/media_file.proto

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

// A media file.
type MediaFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the media file.
	// Media file resource names have the form:
	//
	// `customers/{customer_id}/mediaFiles/{media_file_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the media file.
	Id *int64 `protobuf:"varint,12,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Immutable. Type of the media file.
	Type enums.MediaTypeEnum_MediaType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v14.enums.MediaTypeEnum_MediaType" json:"type,omitempty"`
	// Output only. The mime type of the media file.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,6,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v14.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Immutable. The URL of where the original media file was downloaded from (or
	// a file name). Only used for media of type AUDIO and IMAGE.
	SourceUrl *string `protobuf:"bytes,13,opt,name=source_url,json=sourceUrl,proto3,oneof" json:"source_url,omitempty"`
	// Immutable. The name of the media file. The name can be used by clients to
	// help identify previously uploaded media.
	Name *string `protobuf:"bytes,14,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The size of the media file in bytes.
	FileSize *int64 `protobuf:"varint,15,opt,name=file_size,json=fileSize,proto3,oneof" json:"file_size,omitempty"`
	// The specific type of the media file.
	//
	// Types that are assignable to Mediatype:
	//
	//	*MediaFile_Image
	//	*MediaFile_MediaBundle
	//	*MediaFile_Audio
	//	*MediaFile_Video
	Mediatype isMediaFile_Mediatype `protobuf_oneof:"mediatype"`
}

func (x *MediaFile) Reset() {
	*x = MediaFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_media_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaFile) ProtoMessage() {}

func (x *MediaFile) ProtoReflect() protoreflect.Message {
	mi := &file_resources_media_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaFile.ProtoReflect.Descriptor instead.
func (*MediaFile) Descriptor() ([]byte, []int) {
	return file_resources_media_file_proto_rawDescGZIP(), []int{0}
}

func (x *MediaFile) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MediaFile) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MediaFile) GetType() enums.MediaTypeEnum_MediaType {
	if x != nil {
		return x.Type
	}
	return enums.MediaTypeEnum_UNSPECIFIED
}

func (x *MediaFile) GetMimeType() enums.MimeTypeEnum_MimeType {
	if x != nil {
		return x.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (x *MediaFile) GetSourceUrl() string {
	if x != nil && x.SourceUrl != nil {
		return *x.SourceUrl
	}
	return ""
}

func (x *MediaFile) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MediaFile) GetFileSize() int64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (m *MediaFile) GetMediatype() isMediaFile_Mediatype {
	if m != nil {
		return m.Mediatype
	}
	return nil
}

func (x *MediaFile) GetImage() *MediaImage {
	if x, ok := x.GetMediatype().(*MediaFile_Image); ok {
		return x.Image
	}
	return nil
}

func (x *MediaFile) GetMediaBundle() *MediaBundle {
	if x, ok := x.GetMediatype().(*MediaFile_MediaBundle); ok {
		return x.MediaBundle
	}
	return nil
}

func (x *MediaFile) GetAudio() *MediaAudio {
	if x, ok := x.GetMediatype().(*MediaFile_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *MediaFile) GetVideo() *MediaVideo {
	if x, ok := x.GetMediatype().(*MediaFile_Video); ok {
		return x.Video
	}
	return nil
}

type isMediaFile_Mediatype interface {
	isMediaFile_Mediatype()
}

type MediaFile_Image struct {
	// Immutable. Encapsulates an Image.
	Image *MediaImage `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

type MediaFile_MediaBundle struct {
	// Immutable. A ZIP archive media the content of which contains HTML5
	// assets.
	MediaBundle *MediaBundle `protobuf:"bytes,4,opt,name=media_bundle,json=mediaBundle,proto3,oneof"`
}

type MediaFile_Audio struct {
	// Output only. Encapsulates an Audio.
	Audio *MediaAudio `protobuf:"bytes,10,opt,name=audio,proto3,oneof"`
}

type MediaFile_Video struct {
	// Immutable. Encapsulates a Video.
	Video *MediaVideo `protobuf:"bytes,11,opt,name=video,proto3,oneof"`
}

func (*MediaFile_Image) isMediaFile_Mediatype() {}

func (*MediaFile_MediaBundle) isMediaFile_Mediatype() {}

func (*MediaFile_Audio) isMediaFile_Mediatype() {}

func (*MediaFile_Video) isMediaFile_Mediatype() {}

// Encapsulates an Image.
type MediaImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Raw image data.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
	// Output only. The url to the full size version of the image.
	FullSizeImageUrl *string `protobuf:"bytes,2,opt,name=full_size_image_url,json=fullSizeImageUrl,proto3,oneof" json:"full_size_image_url,omitempty"`
	// Output only. The url to the preview size version of the image.
	PreviewSizeImageUrl *string `protobuf:"bytes,3,opt,name=preview_size_image_url,json=previewSizeImageUrl,proto3,oneof" json:"preview_size_image_url,omitempty"`
}

func (x *MediaImage) Reset() {
	*x = MediaImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_media_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaImage) ProtoMessage() {}

func (x *MediaImage) ProtoReflect() protoreflect.Message {
	mi := &file_resources_media_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaImage.ProtoReflect.Descriptor instead.
func (*MediaImage) Descriptor() ([]byte, []int) {
	return file_resources_media_file_proto_rawDescGZIP(), []int{1}
}

func (x *MediaImage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MediaImage) GetFullSizeImageUrl() string {
	if x != nil && x.FullSizeImageUrl != nil {
		return *x.FullSizeImageUrl
	}
	return ""
}

func (x *MediaImage) GetPreviewSizeImageUrl() string {
	if x != nil && x.PreviewSizeImageUrl != nil {
		return *x.PreviewSizeImageUrl
	}
	return ""
}

// Represents a ZIP archive media the content of which contains HTML5 assets.
type MediaBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Raw zipped data.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	// Output only. The url to access the uploaded zipped data.
	// For example, https://tpc.googlesyndication.com/simgad/123
	// This field is read-only.
	Url *string `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *MediaBundle) Reset() {
	*x = MediaBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_media_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBundle) ProtoMessage() {}

func (x *MediaBundle) ProtoReflect() protoreflect.Message {
	mi := &file_resources_media_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBundle.ProtoReflect.Descriptor instead.
func (*MediaBundle) Descriptor() ([]byte, []int) {
	return file_resources_media_file_proto_rawDescGZIP(), []int{2}
}

func (x *MediaBundle) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MediaBundle) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

// Encapsulates an Audio.
type MediaAudio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The duration of the Audio in milliseconds.
	AdDurationMillis *int64 `protobuf:"varint,2,opt,name=ad_duration_millis,json=adDurationMillis,proto3,oneof" json:"ad_duration_millis,omitempty"`
}

func (x *MediaAudio) Reset() {
	*x = MediaAudio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_media_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaAudio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaAudio) ProtoMessage() {}

func (x *MediaAudio) ProtoReflect() protoreflect.Message {
	mi := &file_resources_media_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaAudio.ProtoReflect.Descriptor instead.
func (*MediaAudio) Descriptor() ([]byte, []int) {
	return file_resources_media_file_proto_rawDescGZIP(), []int{3}
}

func (x *MediaAudio) GetAdDurationMillis() int64 {
	if x != nil && x.AdDurationMillis != nil {
		return *x.AdDurationMillis
	}
	return 0
}

// Encapsulates a Video.
type MediaVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The duration of the Video in milliseconds.
	AdDurationMillis *int64 `protobuf:"varint,5,opt,name=ad_duration_millis,json=adDurationMillis,proto3,oneof" json:"ad_duration_millis,omitempty"`
	// Immutable. The YouTube video ID (as seen in YouTube URLs). Adding prefix
	// "https://www.youtube.com/watch?v=" to this ID will get the YouTube
	// streaming URL for this video.
	YoutubeVideoId *string `protobuf:"bytes,6,opt,name=youtube_video_id,json=youtubeVideoId,proto3,oneof" json:"youtube_video_id,omitempty"`
	// Output only. The Advertising Digital Identification code for this video, as
	// defined by the American Association of Advertising Agencies, used mainly
	// for television commercials.
	AdvertisingIdCode *string `protobuf:"bytes,7,opt,name=advertising_id_code,json=advertisingIdCode,proto3,oneof" json:"advertising_id_code,omitempty"`
	// Output only. The Industry Standard Commercial Identifier code for this
	// video, used mainly for television commercials.
	IsciCode *string `protobuf:"bytes,8,opt,name=isci_code,json=isciCode,proto3,oneof" json:"isci_code,omitempty"`
}

func (x *MediaVideo) Reset() {
	*x = MediaVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_media_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaVideo) ProtoMessage() {}

func (x *MediaVideo) ProtoReflect() protoreflect.Message {
	mi := &file_resources_media_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaVideo.ProtoReflect.Descriptor instead.
func (*MediaVideo) Descriptor() ([]byte, []int) {
	return file_resources_media_file_proto_rawDescGZIP(), []int{4}
}

func (x *MediaVideo) GetAdDurationMillis() int64 {
	if x != nil && x.AdDurationMillis != nil {
		return *x.AdDurationMillis
	}
	return 0
}

func (x *MediaVideo) GetYoutubeVideoId() string {
	if x != nil && x.YoutubeVideoId != nil {
		return *x.YoutubeVideoId
	}
	return ""
}

func (x *MediaVideo) GetAdvertisingIdCode() string {
	if x != nil && x.AdvertisingIdCode != nil {
		return *x.AdvertisingIdCode
	}
	return ""
}

func (x *MediaVideo) GetIsciCode() string {
	if x != nil && x.IsciCode != nil {
		return *x.IsciCode
	}
	return ""
}

var File_resources_media_file_proto protoreflect.FileDescriptor

var file_resources_media_file_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x06, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x50, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x57, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x48, 0x02, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x05, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x4b, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x4b, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x03, 0xe0, 0x41,
	0x05, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x3a, 0x5b, 0xea, 0x41, 0x58, 0x0a,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x32, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0xde, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x37,
	0x0a, 0x13, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x10, 0x66, 0x75, 0x6c, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x13,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x16, 0x0a, 0x14, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x22, 0x58, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12,
	0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x0a,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x36, 0x0a, 0x12, 0x61, 0x64,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x10, 0x61,
	0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0xab, 0x02, 0x0a, 0x0a, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x36, 0x0a, 0x12, 0x61, 0x64, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x10, 0x61, 0x64, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x32, 0x0a, 0x10, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x01, 0x52, 0x0e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x09, 0x69, 0x73, 0x63, 0x69, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x08, 0x69, 0x73, 0x63, 0x69, 0x43, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61, 0x64, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73,
	0x63, 0x69, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x80, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x42, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x34, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a,
	0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_media_file_proto_rawDescOnce sync.Once
	file_resources_media_file_proto_rawDescData = file_resources_media_file_proto_rawDesc
)

func file_resources_media_file_proto_rawDescGZIP() []byte {
	file_resources_media_file_proto_rawDescOnce.Do(func() {
		file_resources_media_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_media_file_proto_rawDescData)
	})
	return file_resources_media_file_proto_rawDescData
}

var file_resources_media_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_resources_media_file_proto_goTypes = []interface{}{
	(*MediaFile)(nil),                  // 0: google.ads.googleads.v14.resources.MediaFile
	(*MediaImage)(nil),                 // 1: google.ads.googleads.v14.resources.MediaImage
	(*MediaBundle)(nil),                // 2: google.ads.googleads.v14.resources.MediaBundle
	(*MediaAudio)(nil),                 // 3: google.ads.googleads.v14.resources.MediaAudio
	(*MediaVideo)(nil),                 // 4: google.ads.googleads.v14.resources.MediaVideo
	(enums.MediaTypeEnum_MediaType)(0), // 5: google.ads.googleads.v14.enums.MediaTypeEnum.MediaType
	(enums.MimeTypeEnum_MimeType)(0),   // 6: google.ads.googleads.v14.enums.MimeTypeEnum.MimeType
}
var file_resources_media_file_proto_depIdxs = []int32{
	5, // 0: google.ads.googleads.v14.resources.MediaFile.type:type_name -> google.ads.googleads.v14.enums.MediaTypeEnum.MediaType
	6, // 1: google.ads.googleads.v14.resources.MediaFile.mime_type:type_name -> google.ads.googleads.v14.enums.MimeTypeEnum.MimeType
	1, // 2: google.ads.googleads.v14.resources.MediaFile.image:type_name -> google.ads.googleads.v14.resources.MediaImage
	2, // 3: google.ads.googleads.v14.resources.MediaFile.media_bundle:type_name -> google.ads.googleads.v14.resources.MediaBundle
	3, // 4: google.ads.googleads.v14.resources.MediaFile.audio:type_name -> google.ads.googleads.v14.resources.MediaAudio
	4, // 5: google.ads.googleads.v14.resources.MediaFile.video:type_name -> google.ads.googleads.v14.resources.MediaVideo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_resources_media_file_proto_init() }
func file_resources_media_file_proto_init() {
	if File_resources_media_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_media_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaFile); i {
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
		file_resources_media_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaImage); i {
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
		file_resources_media_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBundle); i {
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
		file_resources_media_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaAudio); i {
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
		file_resources_media_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaVideo); i {
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
	file_resources_media_file_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MediaFile_Image)(nil),
		(*MediaFile_MediaBundle)(nil),
		(*MediaFile_Audio)(nil),
		(*MediaFile_Video)(nil),
	}
	file_resources_media_file_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_media_file_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_resources_media_file_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_resources_media_file_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_media_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_media_file_proto_goTypes,
		DependencyIndexes: file_resources_media_file_proto_depIdxs,
		MessageInfos:      file_resources_media_file_proto_msgTypes,
	}.Build()
	File_resources_media_file_proto = out.File
	file_resources_media_file_proto_rawDesc = nil
	file_resources_media_file_proto_goTypes = nil
	file_resources_media_file_proto_depIdxs = nil
}
