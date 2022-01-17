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
// source: google/ads/googleads/v9/enums/promotion_extension_occasion.proto

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

// A promotion extension occasion.
type PromotionExtensionOccasionEnum_PromotionExtensionOccasion int32

const (
	// Not specified.
	PromotionExtensionOccasionEnum_UNSPECIFIED PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionExtensionOccasionEnum_UNKNOWN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 1
	// New Year's.
	PromotionExtensionOccasionEnum_NEW_YEARS PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 2
	// Chinese New Year.
	PromotionExtensionOccasionEnum_CHINESE_NEW_YEAR PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 3
	// Valentine's Day.
	PromotionExtensionOccasionEnum_VALENTINES_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 4
	// Easter.
	PromotionExtensionOccasionEnum_EASTER PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 5
	// Mother's Day.
	PromotionExtensionOccasionEnum_MOTHERS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 6
	// Father's Day.
	PromotionExtensionOccasionEnum_FATHERS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 7
	// Labor Day.
	PromotionExtensionOccasionEnum_LABOR_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 8
	// Back To School.
	PromotionExtensionOccasionEnum_BACK_TO_SCHOOL PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 9
	// Halloween.
	PromotionExtensionOccasionEnum_HALLOWEEN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 10
	// Black Friday.
	PromotionExtensionOccasionEnum_BLACK_FRIDAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 11
	// Cyber Monday.
	PromotionExtensionOccasionEnum_CYBER_MONDAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 12
	// Christmas.
	PromotionExtensionOccasionEnum_CHRISTMAS PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 13
	// Boxing Day.
	PromotionExtensionOccasionEnum_BOXING_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 14
	// Independence Day in any country.
	PromotionExtensionOccasionEnum_INDEPENDENCE_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 15
	// National Day in any country.
	PromotionExtensionOccasionEnum_NATIONAL_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 16
	// End of any season.
	PromotionExtensionOccasionEnum_END_OF_SEASON PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 17
	// Winter Sale.
	PromotionExtensionOccasionEnum_WINTER_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 18
	// Summer sale.
	PromotionExtensionOccasionEnum_SUMMER_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 19
	// Fall Sale.
	PromotionExtensionOccasionEnum_FALL_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 20
	// Spring Sale.
	PromotionExtensionOccasionEnum_SPRING_SALE PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 21
	// Ramadan.
	PromotionExtensionOccasionEnum_RAMADAN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 22
	// Eid al-Fitr.
	PromotionExtensionOccasionEnum_EID_AL_FITR PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 23
	// Eid al-Adha.
	PromotionExtensionOccasionEnum_EID_AL_ADHA PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 24
	// Singles Day.
	PromotionExtensionOccasionEnum_SINGLES_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 25
	// Women's Day.
	PromotionExtensionOccasionEnum_WOMENS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 26
	// Holi.
	PromotionExtensionOccasionEnum_HOLI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 27
	// Parent's Day.
	PromotionExtensionOccasionEnum_PARENTS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 28
	// St. Nicholas Day.
	PromotionExtensionOccasionEnum_ST_NICHOLAS_DAY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 29
	// Carnival.
	PromotionExtensionOccasionEnum_CARNIVAL PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 30
	// Epiphany, also known as Three Kings' Day.
	PromotionExtensionOccasionEnum_EPIPHANY PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 31
	// Rosh Hashanah.
	PromotionExtensionOccasionEnum_ROSH_HASHANAH PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 32
	// Passover.
	PromotionExtensionOccasionEnum_PASSOVER PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 33
	// Hanukkah.
	PromotionExtensionOccasionEnum_HANUKKAH PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 34
	// Diwali.
	PromotionExtensionOccasionEnum_DIWALI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 35
	// Navratri.
	PromotionExtensionOccasionEnum_NAVRATRI PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 36
	// Available in Thai: Songkran.
	PromotionExtensionOccasionEnum_SONGKRAN PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 37
	// Available in Japanese: Year-end Gift.
	PromotionExtensionOccasionEnum_YEAR_END_GIFT PromotionExtensionOccasionEnum_PromotionExtensionOccasion = 38
)

// Enum value maps for PromotionExtensionOccasionEnum_PromotionExtensionOccasion.
var (
	PromotionExtensionOccasionEnum_PromotionExtensionOccasion_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "NEW_YEARS",
		3:  "CHINESE_NEW_YEAR",
		4:  "VALENTINES_DAY",
		5:  "EASTER",
		6:  "MOTHERS_DAY",
		7:  "FATHERS_DAY",
		8:  "LABOR_DAY",
		9:  "BACK_TO_SCHOOL",
		10: "HALLOWEEN",
		11: "BLACK_FRIDAY",
		12: "CYBER_MONDAY",
		13: "CHRISTMAS",
		14: "BOXING_DAY",
		15: "INDEPENDENCE_DAY",
		16: "NATIONAL_DAY",
		17: "END_OF_SEASON",
		18: "WINTER_SALE",
		19: "SUMMER_SALE",
		20: "FALL_SALE",
		21: "SPRING_SALE",
		22: "RAMADAN",
		23: "EID_AL_FITR",
		24: "EID_AL_ADHA",
		25: "SINGLES_DAY",
		26: "WOMENS_DAY",
		27: "HOLI",
		28: "PARENTS_DAY",
		29: "ST_NICHOLAS_DAY",
		30: "CARNIVAL",
		31: "EPIPHANY",
		32: "ROSH_HASHANAH",
		33: "PASSOVER",
		34: "HANUKKAH",
		35: "DIWALI",
		36: "NAVRATRI",
		37: "SONGKRAN",
		38: "YEAR_END_GIFT",
	}
	PromotionExtensionOccasionEnum_PromotionExtensionOccasion_value = map[string]int32{
		"UNSPECIFIED":      0,
		"UNKNOWN":          1,
		"NEW_YEARS":        2,
		"CHINESE_NEW_YEAR": 3,
		"VALENTINES_DAY":   4,
		"EASTER":           5,
		"MOTHERS_DAY":      6,
		"FATHERS_DAY":      7,
		"LABOR_DAY":        8,
		"BACK_TO_SCHOOL":   9,
		"HALLOWEEN":        10,
		"BLACK_FRIDAY":     11,
		"CYBER_MONDAY":     12,
		"CHRISTMAS":        13,
		"BOXING_DAY":       14,
		"INDEPENDENCE_DAY": 15,
		"NATIONAL_DAY":     16,
		"END_OF_SEASON":    17,
		"WINTER_SALE":      18,
		"SUMMER_SALE":      19,
		"FALL_SALE":        20,
		"SPRING_SALE":      21,
		"RAMADAN":          22,
		"EID_AL_FITR":      23,
		"EID_AL_ADHA":      24,
		"SINGLES_DAY":      25,
		"WOMENS_DAY":       26,
		"HOLI":             27,
		"PARENTS_DAY":      28,
		"ST_NICHOLAS_DAY":  29,
		"CARNIVAL":         30,
		"EPIPHANY":         31,
		"ROSH_HASHANAH":    32,
		"PASSOVER":         33,
		"HANUKKAH":         34,
		"DIWALI":           35,
		"NAVRATRI":         36,
		"SONGKRAN":         37,
		"YEAR_END_GIFT":    38,
	}
)

func (x PromotionExtensionOccasionEnum_PromotionExtensionOccasion) Enum() *PromotionExtensionOccasionEnum_PromotionExtensionOccasion {
	p := new(PromotionExtensionOccasionEnum_PromotionExtensionOccasion)
	*p = x
	return p
}

func (x PromotionExtensionOccasionEnum_PromotionExtensionOccasion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PromotionExtensionOccasionEnum_PromotionExtensionOccasion) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_promotion_extension_occasion_proto_enumTypes[0].Descriptor()
}

func (PromotionExtensionOccasionEnum_PromotionExtensionOccasion) Type() protoreflect.EnumType {
	return &file_enums_promotion_extension_occasion_proto_enumTypes[0]
}

func (x PromotionExtensionOccasionEnum_PromotionExtensionOccasion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PromotionExtensionOccasionEnum_PromotionExtensionOccasion.Descriptor instead.
func (PromotionExtensionOccasionEnum_PromotionExtensionOccasion) EnumDescriptor() ([]byte, []int) {
	return file_enums_promotion_extension_occasion_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing a promotion extension occasion.
// For more information about the occasions please check:
// https://support.google.com/google-ads/answer/7367521
type PromotionExtensionOccasionEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromotionExtensionOccasionEnum) Reset() {
	*x = PromotionExtensionOccasionEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_promotion_extension_occasion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionExtensionOccasionEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionExtensionOccasionEnum) ProtoMessage() {}

func (x *PromotionExtensionOccasionEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_promotion_extension_occasion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionExtensionOccasionEnum.ProtoReflect.Descriptor instead.
func (*PromotionExtensionOccasionEnum) Descriptor() ([]byte, []int) {
	return file_enums_promotion_extension_occasion_proto_rawDescGZIP(), []int{0}
}

var File_enums_promotion_extension_occasion_proto protoreflect.FileDescriptor

var file_enums_promotion_extension_occasion_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x63, 0x63, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbc, 0x05, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x63, 0x63, 0x61, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x99, 0x05, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x63, 0x63, 0x61, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x45, 0x57, 0x5f, 0x59, 0x45, 0x41, 0x52, 0x53, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x53, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x59, 0x45,
	0x41, 0x52, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x41, 0x4c, 0x45, 0x4e, 0x54, 0x49, 0x4e,
	0x45, 0x53, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x41, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x53, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x54, 0x48, 0x45, 0x52, 0x53,
	0x5f, 0x44, 0x41, 0x59, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x41, 0x42, 0x4f, 0x52, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x54, 0x4f,
	0x5f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x45, 0x4e, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4c, 0x41, 0x43,
	0x4b, 0x5f, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x59,
	0x42, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x48, 0x52, 0x49, 0x53, 0x54, 0x4d, 0x41, 0x53, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x42,
	0x4f, 0x58, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x49,
	0x4e, 0x44, 0x45, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x10,
	0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x44, 0x41,
	0x59, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x10, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x55, 0x4d, 0x4d, 0x45,
	0x52, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x13, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x41, 0x4c, 0x4c,
	0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x14, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x52, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x15, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x41, 0x4d, 0x41,
	0x44, 0x41, 0x4e, 0x10, 0x16, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x5f,
	0x46, 0x49, 0x54, 0x52, 0x10, 0x17, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x49, 0x44, 0x5f, 0x41, 0x4c,
	0x5f, 0x41, 0x44, 0x48, 0x41, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4e, 0x47, 0x4c,
	0x45, 0x53, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x19, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x4f, 0x4d, 0x45,
	0x4e, 0x53, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x1a, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4c, 0x49,
	0x10, 0x1b, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x44, 0x41,
	0x59, 0x10, 0x1c, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x5f, 0x4e, 0x49, 0x43, 0x48, 0x4f, 0x4c,
	0x41, 0x53, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x1d, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x52, 0x4e,
	0x49, 0x56, 0x41, 0x4c, 0x10, 0x1e, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x50, 0x49, 0x50, 0x48, 0x41,
	0x4e, 0x59, 0x10, 0x1f, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53,
	0x48, 0x41, 0x4e, 0x41, 0x48, 0x10, 0x20, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53, 0x4f,
	0x56, 0x45, 0x52, 0x10, 0x21, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x41, 0x4e, 0x55, 0x4b, 0x4b, 0x41,
	0x48, 0x10, 0x22, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x57, 0x41, 0x4c, 0x49, 0x10, 0x23, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x41, 0x56, 0x52, 0x41, 0x54, 0x52, 0x49, 0x10, 0x24, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x4f, 0x4e, 0x47, 0x4b, 0x52, 0x41, 0x4e, 0x10, 0x25, 0x12, 0x11, 0x0a, 0x0d, 0x59,
	0x45, 0x41, 0x52, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x47, 0x49, 0x46, 0x54, 0x10, 0x26, 0x42, 0xf4,
	0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x63, 0x63, 0x61, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_promotion_extension_occasion_proto_rawDescOnce sync.Once
	file_enums_promotion_extension_occasion_proto_rawDescData = file_enums_promotion_extension_occasion_proto_rawDesc
)

func file_enums_promotion_extension_occasion_proto_rawDescGZIP() []byte {
	file_enums_promotion_extension_occasion_proto_rawDescOnce.Do(func() {
		file_enums_promotion_extension_occasion_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_promotion_extension_occasion_proto_rawDescData)
	})
	return file_enums_promotion_extension_occasion_proto_rawDescData
}

var file_enums_promotion_extension_occasion_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_promotion_extension_occasion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_promotion_extension_occasion_proto_goTypes = []interface{}{
	(PromotionExtensionOccasionEnum_PromotionExtensionOccasion)(0), // 0: google.ads.googleads.v9.enums.PromotionExtensionOccasionEnum.PromotionExtensionOccasion
	(*PromotionExtensionOccasionEnum)(nil),                         // 1: google.ads.googleads.v9.enums.PromotionExtensionOccasionEnum
}
var file_enums_promotion_extension_occasion_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_promotion_extension_occasion_proto_init() }
func file_enums_promotion_extension_occasion_proto_init() {
	if File_enums_promotion_extension_occasion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_promotion_extension_occasion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionExtensionOccasionEnum); i {
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
			RawDescriptor: file_enums_promotion_extension_occasion_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_promotion_extension_occasion_proto_goTypes,
		DependencyIndexes: file_enums_promotion_extension_occasion_proto_depIdxs,
		EnumInfos:         file_enums_promotion_extension_occasion_proto_enumTypes,
		MessageInfos:      file_enums_promotion_extension_occasion_proto_msgTypes,
	}.Build()
	File_enums_promotion_extension_occasion_proto = out.File
	file_enums_promotion_extension_occasion_proto_rawDesc = nil
	file_enums_promotion_extension_occasion_proto_goTypes = nil
	file_enums_promotion_extension_occasion_proto_depIdxs = nil
}
