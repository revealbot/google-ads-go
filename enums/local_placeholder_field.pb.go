// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/enums/local_placeholder_field.proto

package enums

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

// Possible values for Local placeholder fields.
type LocalPlaceholderFieldEnum_LocalPlaceholderField int32

const (
	// Not specified.
	LocalPlaceholderFieldEnum_UNSPECIFIED LocalPlaceholderFieldEnum_LocalPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	LocalPlaceholderFieldEnum_UNKNOWN LocalPlaceholderFieldEnum_LocalPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	LocalPlaceholderFieldEnum_DEAL_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with local deal title to be
	// shown in dynamic ad.
	LocalPlaceholderFieldEnum_DEAL_NAME LocalPlaceholderFieldEnum_LocalPlaceholderField = 3
	// Data Type: STRING. Local deal subtitle to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_SUBTITLE LocalPlaceholderFieldEnum_LocalPlaceholderField = 4
	// Data Type: STRING. Description of local deal to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_DESCRIPTION LocalPlaceholderFieldEnum_LocalPlaceholderField = 5
	// Data Type: STRING. Price to be shown in the ad. Highly recommended for
	// dynamic ads. Example: "100.00 USD"
	LocalPlaceholderFieldEnum_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 6
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	LocalPlaceholderFieldEnum_FORMATTED_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 7
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	LocalPlaceholderFieldEnum_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 8
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	LocalPlaceholderFieldEnum_FORMATTED_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 9
	// Data Type: URL. Image to be displayed in the ad.
	LocalPlaceholderFieldEnum_IMAGE_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 10
	// Data Type: STRING. Complete property address, including postal code.
	LocalPlaceholderFieldEnum_ADDRESS LocalPlaceholderFieldEnum_LocalPlaceholderField = 11
	// Data Type: STRING. Category of local deal used to group like items
	// together for recommendation engine.
	LocalPlaceholderFieldEnum_CATEGORY LocalPlaceholderFieldEnum_LocalPlaceholderField = 12
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	LocalPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 13
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific local deal and its location).
	LocalPlaceholderFieldEnum_FINAL_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 14
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	LocalPlaceholderFieldEnum_FINAL_MOBILE_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 15
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	LocalPlaceholderFieldEnum_TRACKING_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 16
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	LocalPlaceholderFieldEnum_ANDROID_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 17
	// Data Type: STRING_LIST. List of recommended local deal IDs to show
	// together with this item.
	LocalPlaceholderFieldEnum_SIMILAR_DEAL_IDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 18
	// Data Type: STRING. iOS app link.
	LocalPlaceholderFieldEnum_IOS_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 19
	// Data Type: INT64. iOS app store ID.
	LocalPlaceholderFieldEnum_IOS_APP_STORE_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 20
)

var LocalPlaceholderFieldEnum_LocalPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DEAL_ID",
	3:  "DEAL_NAME",
	4:  "SUBTITLE",
	5:  "DESCRIPTION",
	6:  "PRICE",
	7:  "FORMATTED_PRICE",
	8:  "SALE_PRICE",
	9:  "FORMATTED_SALE_PRICE",
	10: "IMAGE_URL",
	11: "ADDRESS",
	12: "CATEGORY",
	13: "CONTEXTUAL_KEYWORDS",
	14: "FINAL_URLS",
	15: "FINAL_MOBILE_URLS",
	16: "TRACKING_URL",
	17: "ANDROID_APP_LINK",
	18: "SIMILAR_DEAL_IDS",
	19: "IOS_APP_LINK",
	20: "IOS_APP_STORE_ID",
}

var LocalPlaceholderFieldEnum_LocalPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"DEAL_ID":              2,
	"DEAL_NAME":            3,
	"SUBTITLE":             4,
	"DESCRIPTION":          5,
	"PRICE":                6,
	"FORMATTED_PRICE":      7,
	"SALE_PRICE":           8,
	"FORMATTED_SALE_PRICE": 9,
	"IMAGE_URL":            10,
	"ADDRESS":              11,
	"CATEGORY":             12,
	"CONTEXTUAL_KEYWORDS":  13,
	"FINAL_URLS":           14,
	"FINAL_MOBILE_URLS":    15,
	"TRACKING_URL":         16,
	"ANDROID_APP_LINK":     17,
	"SIMILAR_DEAL_IDS":     18,
	"IOS_APP_LINK":         19,
	"IOS_APP_STORE_ID":     20,
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) String() string {
	return proto.EnumName(LocalPlaceholderFieldEnum_LocalPlaceholderField_name, int32(x))
}

func (LocalPlaceholderFieldEnum_LocalPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_73ca15daa532ccbb, []int{0, 0}
}

// Values for Local placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type LocalPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalPlaceholderFieldEnum) Reset()         { *m = LocalPlaceholderFieldEnum{} }
func (m *LocalPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*LocalPlaceholderFieldEnum) ProtoMessage()    {}
func (*LocalPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ca15daa532ccbb, []int{0}
}

func (m *LocalPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *LocalPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *LocalPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalPlaceholderFieldEnum.Merge(m, src)
}
func (m *LocalPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Size(m)
}
func (m *LocalPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LocalPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.enums.LocalPlaceholderFieldEnum_LocalPlaceholderField", LocalPlaceholderFieldEnum_LocalPlaceholderField_name, LocalPlaceholderFieldEnum_LocalPlaceholderField_value)
	proto.RegisterType((*LocalPlaceholderFieldEnum)(nil), "google.ads.googleads.v9.enums.LocalPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/enums/local_placeholder_field.proto", fileDescriptor_73ca15daa532ccbb)
}

var fileDescriptor_73ca15daa532ccbb = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0x1e, 0xb0, 0xfe, 0x60, 0x68, 0x71, 0x0d, 0xd5, 0x7e, 0xa4, 0x5e, 0xb4, 0x0f, 0x10, 0x90,
	0x76, 0x97, 0x5e, 0x99, 0xd8, 0x20, 0x0b, 0x63, 0x47, 0xb6, 0x43, 0xd7, 0x09, 0x29, 0xca, 0x48,
	0x96, 0x55, 0x0a, 0x04, 0x91, 0xb5, 0x0f, 0xb1, 0xc7, 0xd8, 0xd5, 0xb4, 0x47, 0xd9, 0xa3, 0xec,
	0x66, 0xaf, 0x30, 0x39, 0xe1, 0x67, 0x17, 0xdd, 0x6e, 0xa2, 0x73, 0xbe, 0x9f, 0x73, 0x8e, 0xe2,
	0x0f, 0xdc, 0xa6, 0x79, 0x9e, 0x66, 0x49, 0x3f, 0x8a, 0x8b, 0x7e, 0x55, 0xda, 0xea, 0x69, 0xd0,
	0x4f, 0x56, 0x8f, 0xcb, 0xa2, 0x9f, 0xe5, 0x8b, 0x28, 0x0b, 0xd7, 0x59, 0xb4, 0x48, 0x3e, 0xe7,
	0x59, 0x9c, 0x6c, 0xc2, 0x4f, 0x0f, 0x49, 0x16, 0x3b, 0xeb, 0x4d, 0xfe, 0x25, 0x47, 0x57, 0x95,
	0xc3, 0x89, 0xe2, 0xc2, 0xd9, 0x9b, 0x9d, 0xa7, 0x81, 0x53, 0x9a, 0x6f, 0xbe, 0x37, 0xc0, 0x1b,
	0x6e, 0x07, 0xf8, 0x07, 0xff, 0xc8, 0xda, 0xe9, 0xea, 0x71, 0x79, 0xf3, 0xb5, 0x01, 0x2e, 0x9f,
	0x65, 0x51, 0x07, 0xb4, 0x02, 0xa1, 0x7d, 0xea, 0xb1, 0x11, 0xa3, 0x04, 0xbe, 0x40, 0x2d, 0x70,
	0x12, 0x88, 0x89, 0x90, 0x77, 0x02, 0xd6, 0x6c, 0x43, 0x28, 0xe6, 0x21, 0x23, 0xb0, 0x8e, 0xce,
	0x40, 0xb3, 0x6c, 0x04, 0x9e, 0x52, 0xd8, 0x40, 0x6d, 0x70, 0xaa, 0x83, 0xa1, 0x61, 0x86, 0x53,
	0xf8, 0xd2, 0xce, 0x21, 0x54, 0x7b, 0x8a, 0xf9, 0x86, 0x49, 0x01, 0x8f, 0x50, 0x13, 0x1c, 0xf9,
	0x8a, 0x79, 0x14, 0x1e, 0xa3, 0x2e, 0xe8, 0x8c, 0xa4, 0x9a, 0x62, 0x63, 0x28, 0x09, 0x2b, 0xf0,
	0x04, 0x9d, 0x03, 0xa0, 0x31, 0xa7, 0xdb, 0xfe, 0x14, 0xbd, 0x06, 0xbd, 0x83, 0xe8, 0x2f, 0xa6,
	0x69, 0xf7, 0xb2, 0x29, 0x1e, 0xd3, 0x30, 0x50, 0x1c, 0x02, 0x7b, 0x13, 0x26, 0x44, 0x51, 0xad,
	0x61, 0xcb, 0x1e, 0xe1, 0x61, 0x43, 0xc7, 0x52, 0xdd, 0xc3, 0x36, 0x7a, 0x05, 0xba, 0x9e, 0x14,
	0x86, 0xbe, 0x37, 0x01, 0xe6, 0xe1, 0x84, 0xde, 0xdf, 0x49, 0x45, 0x34, 0x3c, 0xb3, 0xcb, 0x46,
	0x4c, 0x60, 0x6e, 0x47, 0x68, 0x78, 0x8e, 0x2e, 0xc1, 0x45, 0xd5, 0x4f, 0xe5, 0x90, 0x71, 0x5a,
	0xc1, 0x1d, 0x04, 0x41, 0xdb, 0x28, 0xec, 0x4d, 0x98, 0x18, 0x97, 0xcb, 0x20, 0xea, 0x01, 0x88,
	0x05, 0x51, 0x92, 0x91, 0x10, 0xfb, 0x7e, 0xc8, 0x99, 0x98, 0xc0, 0x0b, 0x8b, 0x6a, 0x36, 0x65,
	0x1c, 0xab, 0x70, 0xfb, 0x7b, 0x34, 0x44, 0xd6, 0xcd, 0xa4, 0x3e, 0xe8, 0xba, 0x56, 0xb7, 0x43,
	0xb4, 0x91, 0x8a, 0xda, 0xff, 0xd8, 0x1b, 0xfe, 0xae, 0x81, 0xeb, 0x45, 0xbe, 0x74, 0xfe, 0xfb,
	0xa0, 0xc3, 0xb7, 0xcf, 0xbe, 0x97, 0x6f, 0xb3, 0xe0, 0xd7, 0x3e, 0x0c, 0xb7, 0xe6, 0x34, 0xcf,
	0xa2, 0x55, 0xea, 0xe4, 0x9b, 0xb4, 0x9f, 0x26, 0xab, 0x32, 0x29, 0xbb, 0x68, 0xad, 0x1f, 0x8a,
	0x7f, 0x24, 0xed, 0xb6, 0xfc, 0x7e, 0xab, 0x37, 0xc6, 0x18, 0xff, 0xa8, 0x5f, 0x8d, 0xab, 0x51,
	0x38, 0x2e, 0x9c, 0xaa, 0xb4, 0xd5, 0x6c, 0xe0, 0xd8, 0xe4, 0x14, 0x3f, 0x77, 0xfc, 0x1c, 0xc7,
	0xc5, 0x7c, 0xcf, 0xcf, 0x67, 0x83, 0x79, 0xc9, 0xff, 0xaa, 0x5f, 0x57, 0xa0, 0xeb, 0xe2, 0xb8,
	0x70, 0xdd, 0xbd, 0xc2, 0x75, 0x67, 0x03, 0xd7, 0x2d, 0x35, 0x1f, 0x8f, 0xcb, 0xc3, 0xde, 0xfd,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xae, 0x98, 0x53, 0x81, 0x01, 0x03, 0x00, 0x00,
}
