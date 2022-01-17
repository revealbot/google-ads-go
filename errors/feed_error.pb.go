// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/errors/feed_error.proto

package errors

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

// Enum describing possible feed errors.
type FeedErrorEnum_FeedError int32

const (
	// Enum unspecified.
	FeedErrorEnum_UNSPECIFIED FeedErrorEnum_FeedError = 0
	// The received error code is not known in this version.
	FeedErrorEnum_UNKNOWN FeedErrorEnum_FeedError = 1
	// The names of the FeedAttributes must be unique.
	FeedErrorEnum_ATTRIBUTE_NAMES_NOT_UNIQUE FeedErrorEnum_FeedError = 2
	// The attribute list must be an exact copy of the existing list if the
	// attribute ID's are present.
	FeedErrorEnum_ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES FeedErrorEnum_FeedError = 3
	// Cannot specify USER origin for a system generated feed.
	FeedErrorEnum_CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED FeedErrorEnum_FeedError = 4
	// Cannot specify GOOGLE origin for a non-system generated feed.
	FeedErrorEnum_CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED FeedErrorEnum_FeedError = 5
	// Cannot specify feed attributes for system feed.
	FeedErrorEnum_CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED FeedErrorEnum_FeedError = 6
	// Cannot update FeedAttributes on feed with origin GOOGLE.
	FeedErrorEnum_CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE FeedErrorEnum_FeedError = 7
	// The given ID refers to a removed Feed. Removed Feeds are immutable.
	FeedErrorEnum_FEED_REMOVED FeedErrorEnum_FeedError = 8
	// The origin of the feed is not valid for the client.
	FeedErrorEnum_INVALID_ORIGIN_VALUE FeedErrorEnum_FeedError = 9
	// A user can only create and modify feeds with USER origin.
	FeedErrorEnum_FEED_ORIGIN_IS_NOT_USER FeedErrorEnum_FeedError = 10
	// Invalid auth token for the given email.
	FeedErrorEnum_INVALID_AUTH_TOKEN_FOR_EMAIL FeedErrorEnum_FeedError = 11
	// Invalid email specified.
	FeedErrorEnum_INVALID_EMAIL FeedErrorEnum_FeedError = 12
	// Feed name matches that of another active Feed.
	FeedErrorEnum_DUPLICATE_FEED_NAME FeedErrorEnum_FeedError = 13
	// Name of feed is not allowed.
	FeedErrorEnum_INVALID_FEED_NAME FeedErrorEnum_FeedError = 14
	// Missing OAuthInfo.
	FeedErrorEnum_MISSING_OAUTH_INFO FeedErrorEnum_FeedError = 15
	// New FeedAttributes must not affect the unique key.
	FeedErrorEnum_NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY FeedErrorEnum_FeedError = 16
	// Too many FeedAttributes for a Feed.
	FeedErrorEnum_TOO_MANY_ATTRIBUTES FeedErrorEnum_FeedError = 17
	// The business account is not valid.
	FeedErrorEnum_INVALID_BUSINESS_ACCOUNT FeedErrorEnum_FeedError = 18
	// Business account cannot access Google My Business account.
	FeedErrorEnum_BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT FeedErrorEnum_FeedError = 19
	// Invalid chain ID provided for affiliate location feed.
	FeedErrorEnum_INVALID_AFFILIATE_CHAIN_ID FeedErrorEnum_FeedError = 20
)

var FeedErrorEnum_FeedError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ATTRIBUTE_NAMES_NOT_UNIQUE",
	3:  "ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES",
	4:  "CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED",
	5:  "CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED",
	6:  "CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED",
	7:  "CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE",
	8:  "FEED_REMOVED",
	9:  "INVALID_ORIGIN_VALUE",
	10: "FEED_ORIGIN_IS_NOT_USER",
	11: "INVALID_AUTH_TOKEN_FOR_EMAIL",
	12: "INVALID_EMAIL",
	13: "DUPLICATE_FEED_NAME",
	14: "INVALID_FEED_NAME",
	15: "MISSING_OAUTH_INFO",
	16: "NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY",
	17: "TOO_MANY_ATTRIBUTES",
	18: "INVALID_BUSINESS_ACCOUNT",
	19: "BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT",
	20: "INVALID_AFFILIATE_CHAIN_ID",
}

var FeedErrorEnum_FeedError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"ATTRIBUTE_NAMES_NOT_UNIQUE": 2,
	"ATTRIBUTES_DO_NOT_MATCH_EXISTING_ATTRIBUTES":      3,
	"CANNOT_SPECIFY_USER_ORIGIN_FOR_SYSTEM_FEED":       4,
	"CANNOT_SPECIFY_GOOGLE_ORIGIN_FOR_NON_SYSTEM_FEED": 5,
	"CANNOT_SPECIFY_FEED_ATTRIBUTES_FOR_SYSTEM_FEED":   6,
	"CANNOT_UPDATE_FEED_ATTRIBUTES_WITH_ORIGIN_GOOGLE": 7,
	"FEED_REMOVED":                                    8,
	"INVALID_ORIGIN_VALUE":                            9,
	"FEED_ORIGIN_IS_NOT_USER":                         10,
	"INVALID_AUTH_TOKEN_FOR_EMAIL":                    11,
	"INVALID_EMAIL":                                   12,
	"DUPLICATE_FEED_NAME":                             13,
	"INVALID_FEED_NAME":                               14,
	"MISSING_OAUTH_INFO":                              15,
	"NEW_ATTRIBUTE_CANNOT_BE_PART_OF_UNIQUE_KEY":      16,
	"TOO_MANY_ATTRIBUTES":                             17,
	"INVALID_BUSINESS_ACCOUNT":                        18,
	"BUSINESS_ACCOUNT_CANNOT_ACCESS_LOCATION_ACCOUNT": 19,
	"INVALID_AFFILIATE_CHAIN_ID":                      20,
}

func (x FeedErrorEnum_FeedError) String() string {
	return proto.EnumName(FeedErrorEnum_FeedError_name, int32(x))
}

func (FeedErrorEnum_FeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_431d33a773e35cd7, []int{0, 0}
}

// Container for enum describing possible feed errors.
type FeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedErrorEnum) Reset()         { *m = FeedErrorEnum{} }
func (m *FeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedErrorEnum) ProtoMessage()    {}
func (*FeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_431d33a773e35cd7, []int{0}
}

func (m *FeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedErrorEnum.Unmarshal(m, b)
}
func (m *FeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedErrorEnum.Merge(m, src)
}
func (m *FeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedErrorEnum.Size(m)
}
func (m *FeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v9.errors.FeedErrorEnum_FeedError", FeedErrorEnum_FeedError_name, FeedErrorEnum_FeedError_value)
	proto.RegisterType((*FeedErrorEnum)(nil), "google.ads.googleads.v9.errors.FeedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/errors/feed_error.proto", fileDescriptor_431d33a773e35cd7)
}

var fileDescriptor_431d33a773e35cd7 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0xeb, 0x7e, 0xdb, 0x98, 0xf7, 0xcf, 0xf3, 0x06, 0x9b, 0x60, 0x9a, 0x50, 0x2f,
	0x41, 0x4a, 0x2a, 0xc6, 0x55, 0xb8, 0x72, 0x93, 0x93, 0xd6, 0x5a, 0x62, 0x97, 0xd8, 0xe9, 0x28,
	0xaa, 0x64, 0x0d, 0x12, 0x2a, 0xa4, 0x6d, 0x99, 0x1a, 0xd8, 0x03, 0x71, 0xc9, 0xa3, 0x20, 0x9e,
	0x81, 0x7b, 0x78, 0x0a, 0xe4, 0x38, 0xc9, 0xca, 0x26, 0xb8, 0xea, 0xe9, 0xf1, 0xe7, 0x9b, 0xef,
	0xf1, 0x39, 0x3e, 0xc8, 0x9d, 0x15, 0xc5, 0xec, 0x22, 0x77, 0xcf, 0xb3, 0xb2, 0x0e, 0x4d, 0x74,
	0xd3, 0x73, 0xf3, 0xf9, 0xbc, 0x98, 0x97, 0xee, 0x87, 0x3c, 0xcf, 0x74, 0x15, 0x3b, 0xd7, 0xf3,
	0xe2, 0x53, 0x41, 0x8e, 0x2d, 0xe5, 0x9c, 0x67, 0xa5, 0xd3, 0x0a, 0x9c, 0x9b, 0x9e, 0x63, 0x05,
	0xdd, 0x1f, 0x2b, 0x68, 0x2b, 0xcc, 0xf3, 0x0c, 0xcc, 0x5f, 0xb8, 0xfa, 0x7c, 0xd9, 0xfd, 0xbe,
	0x82, 0xd6, 0xdb, 0x0c, 0xd9, 0x41, 0x1b, 0x29, 0x97, 0x23, 0xf0, 0x59, 0xc8, 0x20, 0xc0, 0xff,
	0x91, 0x0d, 0xb4, 0x96, 0xf2, 0x53, 0x2e, 0xce, 0x38, 0x5e, 0x22, 0xc7, 0xe8, 0x31, 0x55, 0x2a,
	0x61, 0xfd, 0x54, 0x81, 0xe6, 0x34, 0x06, 0xa9, 0xb9, 0x50, 0x3a, 0xe5, 0xec, 0x75, 0x0a, 0xb8,
	0x43, 0x5c, 0xf4, 0xbc, 0x3d, 0x97, 0x3a, 0x10, 0xd5, 0x69, 0x4c, 0x95, 0x3f, 0xd4, 0xf0, 0x86,
	0x49, 0xc5, 0xf8, 0x40, 0xdf, 0x02, 0x78, 0x99, 0x38, 0xe8, 0x99, 0x4f, 0xb9, 0xa1, 0xac, 0xe7,
	0x44, 0xa7, 0x12, 0x12, 0x2d, 0x12, 0x36, 0x60, 0x5c, 0x87, 0x22, 0xd1, 0x72, 0x22, 0x15, 0xc4,
	0x3a, 0x04, 0x08, 0xf0, 0xff, 0xe4, 0x25, 0xea, 0xdd, 0xe1, 0x07, 0x42, 0x0c, 0x22, 0x58, 0x54,
	0x70, 0xc1, 0xff, 0x50, 0xad, 0x90, 0x17, 0xc8, 0xb9, 0xa3, 0x32, 0x07, 0x0b, 0x95, 0xdc, 0x73,
	0x5a, 0x5d, 0x70, 0x4a, 0x47, 0x01, 0x55, 0x70, 0x4f, 0x72, 0xc6, 0xd4, 0xb0, 0xb1, 0xb5, 0x45,
	0xe0, 0x35, 0x82, 0xd1, 0x66, 0xc5, 0x25, 0x10, 0x8b, 0x31, 0x04, 0xf8, 0x01, 0x39, 0x44, 0xfb,
	0x8c, 0x8f, 0x69, 0xc4, 0x82, 0x06, 0x1e, 0xd3, 0x28, 0x05, 0xbc, 0x4e, 0x9e, 0xa0, 0x83, 0x8a,
	0xad, 0xd3, 0xac, 0xee, 0xa5, 0x84, 0x04, 0x23, 0xf2, 0x14, 0x1d, 0x35, 0x32, 0x9a, 0xaa, 0xa1,
	0x56, 0xe2, 0x14, 0xec, 0xf5, 0x20, 0xa6, 0x2c, 0xc2, 0x1b, 0x64, 0x17, 0x6d, 0x35, 0x84, 0x4d,
	0x6d, 0x92, 0x03, 0xb4, 0x17, 0xa4, 0xa3, 0x88, 0xf9, 0x6d, 0xbd, 0x66, 0x46, 0x78, 0x8b, 0x3c,
	0x44, 0xbb, 0x0d, 0x7b, 0x9b, 0xde, 0x26, 0x8f, 0x10, 0x89, 0x99, 0x94, 0x66, 0x2a, 0xa2, 0x72,
	0x61, 0x3c, 0x14, 0x78, 0xc7, 0x4c, 0x85, 0xc3, 0xd9, 0xed, 0x65, 0x75, 0xdd, 0x89, 0x3e, 0xe8,
	0x11, 0x4d, 0x94, 0x16, 0x61, 0x3d, 0x72, 0x7d, 0x0a, 0x13, 0x8c, 0x8d, 0xaf, 0x12, 0x42, 0xc7,
	0x94, 0x4f, 0x16, 0xc7, 0xbb, 0x4b, 0x8e, 0xd0, 0x61, 0xe3, 0xdb, 0x4f, 0x25, 0xe3, 0x20, 0xa5,
	0xa6, 0xbe, 0x2f, 0x52, 0xae, 0x30, 0x21, 0x27, 0xc8, 0xbd, 0x9b, 0x6d, 0x9c, 0xa8, 0xef, 0x9b,
	0x6c, 0x24, 0x7c, 0xaa, 0x98, 0xe0, 0xad, 0x68, 0xcf, 0x3c, 0xc1, 0xb6, 0x31, 0x61, 0xc8, 0x22,
	0x66, 0xee, 0xea, 0x0f, 0xa9, 0xe9, 0x60, 0x80, 0xf7, 0xfb, 0x3f, 0x97, 0x50, 0xf7, 0x7d, 0x71,
	0xe9, 0xfc, 0x7b, 0x0f, 0xfa, 0xdb, 0xed, 0x93, 0x1f, 0x99, 0xbd, 0x19, 0x2d, 0xbd, 0x0d, 0x6a,
	0xc5, 0xac, 0xb8, 0x38, 0xbf, 0x9a, 0x39, 0xc5, 0x7c, 0xe6, 0xce, 0xf2, 0xab, 0x6a, 0xab, 0x9a,
	0xd5, 0xbb, 0xfe, 0x58, 0xfe, 0x6d, 0x13, 0x5f, 0xd9, 0x9f, 0x2f, 0x9d, 0xe5, 0x01, 0xa5, 0x5f,
	0x3b, 0xc7, 0x03, 0xfb, 0x31, 0x9a, 0x95, 0x8e, 0x0d, 0x4d, 0x34, 0xee, 0x39, 0x95, 0x65, 0xf9,
	0xad, 0x01, 0xa6, 0x34, 0x2b, 0xa7, 0x2d, 0x30, 0x1d, 0xf7, 0xa6, 0x16, 0xf8, 0xd5, 0xe9, 0xda,
	0xac, 0xe7, 0xd1, 0xac, 0xf4, 0xbc, 0x16, 0xf1, 0xbc, 0x71, 0xcf, 0xf3, 0x2c, 0xf4, 0x6e, 0xb5,
	0xaa, 0xee, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x09, 0x5c, 0x50, 0x26, 0x04, 0x00,
	0x00,
}
