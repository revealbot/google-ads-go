// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/keyword_plan_campaign.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/revealbot/google-ads-go/enums"
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

// A Keyword Plan campaign.
// Max number of keyword plan campaigns per plan allowed: 1.
type KeywordPlanCampaign struct {
	// The resource name of the Keyword Plan campaign.
	// KeywordPlanCampaign resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaigns/{kp_campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan this campaign belongs to.
	KeywordPlan *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan,json=keywordPlan,proto3" json:"keyword_plan,omitempty"`
	// The ID of the Keyword Plan campaign.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Keyword Plan campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The languages targeted for the Keyword Plan campaign.
	// Max allowed: 1.
	LanguageConstants []*wrappers.StringValue `protobuf:"bytes,5,rep,name=language_constants,json=languageConstants,proto3" json:"language_constants,omitempty"`
	// Targeting network.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,6,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v0.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// A default max cpc bid in micros, and in the account currency, for all ad
	// groups under the campaign.
	//
	// This field is required and should not be empty when creating Keyword Plan
	// campaigns.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The geo targets.
	// Max number allowed: 20.
	GeoTargets           []*KeywordPlanGeoTarget `protobuf:"bytes,8,rep,name=geo_targets,json=geoTargets,proto3" json:"geo_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordPlanCampaign) Reset()         { *m = KeywordPlanCampaign{} }
func (m *KeywordPlanCampaign) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaign) ProtoMessage()    {}
func (*KeywordPlanCampaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ee1a602123b5150, []int{0}
}

func (m *KeywordPlanCampaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaign.Unmarshal(m, b)
}
func (m *KeywordPlanCampaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaign.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaign.Merge(m, src)
}
func (m *KeywordPlanCampaign) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaign.Size(m)
}
func (m *KeywordPlanCampaign) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaign.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaign proto.InternalMessageInfo

func (m *KeywordPlanCampaign) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanCampaign) GetKeywordPlan() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlan
	}
	return nil
}

func (m *KeywordPlanCampaign) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanCampaign) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanCampaign) GetLanguageConstants() []*wrappers.StringValue {
	if m != nil {
		return m.LanguageConstants
	}
	return nil
}

func (m *KeywordPlanCampaign) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

func (m *KeywordPlanCampaign) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *KeywordPlanCampaign) GetGeoTargets() []*KeywordPlanGeoTarget {
	if m != nil {
		return m.GeoTargets
	}
	return nil
}

// A geo target.
// Next ID: 3
type KeywordPlanGeoTarget struct {
	// Required. The resource name of the geo target.
	GeoTargetConstant    *wrappers.StringValue `protobuf:"bytes,1,opt,name=geo_target_constant,json=geoTargetConstant,proto3" json:"geo_target_constant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeywordPlanGeoTarget) Reset()         { *m = KeywordPlanGeoTarget{} }
func (m *KeywordPlanGeoTarget) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanGeoTarget) ProtoMessage()    {}
func (*KeywordPlanGeoTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ee1a602123b5150, []int{1}
}

func (m *KeywordPlanGeoTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanGeoTarget.Unmarshal(m, b)
}
func (m *KeywordPlanGeoTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanGeoTarget.Marshal(b, m, deterministic)
}
func (m *KeywordPlanGeoTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanGeoTarget.Merge(m, src)
}
func (m *KeywordPlanGeoTarget) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanGeoTarget.Size(m)
}
func (m *KeywordPlanGeoTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanGeoTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanGeoTarget proto.InternalMessageInfo

func (m *KeywordPlanGeoTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstant
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanCampaign)(nil), "google.ads.googleads.v0.resources.KeywordPlanCampaign")
	proto.RegisterType((*KeywordPlanGeoTarget)(nil), "google.ads.googleads.v0.resources.KeywordPlanGeoTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/keyword_plan_campaign.proto", fileDescriptor_8ee1a602123b5150)
}

var fileDescriptor_8ee1a602123b5150 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x61, 0x6b, 0xd4, 0x30,
	0x18, 0xc7, 0x69, 0x77, 0x4e, 0xcd, 0x9d, 0x03, 0xb3, 0xbd, 0x28, 0x53, 0xe4, 0x36, 0x19, 0x1c,
	0x08, 0x69, 0x99, 0xa2, 0x52, 0x11, 0xe9, 0x0d, 0x39, 0x75, 0x3a, 0x8e, 0x2a, 0x87, 0xc8, 0x41,
	0xc9, 0x35, 0x31, 0x94, 0x6b, 0x93, 0x92, 0xb4, 0x3b, 0xf4, 0xbd, 0x5f, 0xc4, 0x97, 0x7e, 0x0a,
	0x5f, 0xfb, 0x51, 0xfc, 0x14, 0x72, 0xc9, 0xa5, 0x3b, 0xdc, 0xcd, 0xfa, 0xee, 0x69, 0xf2, 0xff,
	0x3d, 0xff, 0xa7, 0xff, 0x24, 0xe0, 0x39, 0x13, 0x82, 0xe5, 0xd4, 0xc7, 0x44, 0xf9, 0xa6, 0x5c,
	0x56, 0xe7, 0x81, 0x2f, 0xa9, 0x12, 0xb5, 0x4c, 0xa9, 0xf2, 0xe7, 0xf4, 0xcb, 0x42, 0x48, 0x92,
	0x94, 0x39, 0xe6, 0x49, 0x8a, 0x8b, 0x12, 0x67, 0x8c, 0xa3, 0x52, 0x8a, 0x4a, 0xc0, 0x03, 0xc3,
	0x20, 0x4c, 0x14, 0x6a, 0x70, 0x74, 0x1e, 0xa0, 0x06, 0xdf, 0x7f, 0x7a, 0x95, 0x03, 0xe5, 0x75,
	0xf1, 0x57, 0x77, 0x4e, 0xab, 0x85, 0x90, 0x73, 0xd3, 0x7c, 0xff, 0xde, 0x8a, 0xd4, 0x5f, 0xb3,
	0xfa, 0xb3, 0xbf, 0x90, 0xb8, 0x2c, 0xa9, 0x54, 0x66, 0xff, 0xf0, 0x67, 0x07, 0xec, 0x9e, 0x1a,
	0x7c, 0x9c, 0x63, 0x7e, 0xb2, 0x1a, 0x0d, 0xde, 0x07, 0xb7, 0xac, 0x7d, 0xc2, 0x71, 0x41, 0x3d,
	0xa7, 0xef, 0x0c, 0x6e, 0xc6, 0x3d, 0xbb, 0x78, 0x86, 0x0b, 0x0a, 0x5f, 0x80, 0xde, 0xba, 0xb5,
	0xe7, 0xf6, 0x9d, 0x41, 0xf7, 0xf8, 0xee, 0xea, 0x2f, 0x90, 0xf5, 0x44, 0xef, 0x2b, 0x99, 0x71,
	0x36, 0xc1, 0x79, 0x4d, 0xe3, 0xee, 0xfc, 0xc2, 0x0d, 0x3e, 0x00, 0x6e, 0x46, 0xbc, 0x2d, 0x8d,
	0xdd, 0xb9, 0x84, 0xbd, 0xe6, 0xd5, 0xe3, 0x47, 0x86, 0x72, 0x33, 0x02, 0x03, 0xd0, 0xd1, 0x93,
	0x74, 0xfe, 0xc3, 0x45, 0x2b, 0xe1, 0x29, 0x80, 0x39, 0xe6, 0xac, 0xc6, 0x8c, 0x26, 0xa9, 0xe0,
	0xaa, 0xc2, 0xbc, 0x52, 0xde, 0xb5, 0xfe, 0x56, 0x2b, 0x7f, 0xdb, 0x72, 0x27, 0x16, 0x83, 0x5f,
	0xc1, 0xde, 0xa6, 0x9c, 0xbd, 0xed, 0xbe, 0x33, 0xd8, 0x39, 0x7e, 0x85, 0xae, 0x3a, 0x45, 0x7d,
	0x44, 0x68, 0x2d, 0xe3, 0x33, 0x03, 0xbe, 0xe4, 0x75, 0xb1, 0x61, 0x39, 0x86, 0xf3, 0x4b, 0x6b,
	0x30, 0x02, 0x3b, 0x69, 0x99, 0x26, 0xb3, 0x8c, 0x24, 0x45, 0x96, 0x4a, 0xa1, 0xbc, 0xeb, 0xed,
	0x99, 0xf5, 0xd2, 0x32, 0x1d, 0x66, 0xe4, 0x9d, 0x06, 0xe0, 0x47, 0xd0, 0x65, 0x54, 0x24, 0x15,
	0x96, 0x8c, 0x56, 0xca, 0xbb, 0xa1, 0x43, 0x78, 0x82, 0x5a, 0xef, 0xde, 0xfa, 0x88, 0x23, 0x2a,
	0x3e, 0x68, 0x3e, 0x06, 0xcc, 0x96, 0xea, 0x90, 0x80, 0xbd, 0x4d, 0x1a, 0xf8, 0x16, 0xec, 0x5e,
	0x38, 0x36, 0xf9, 0xeb, 0x8b, 0xd4, 0x1a, 0x7f, 0xd3, 0xde, 0xe6, 0x3f, 0xfc, 0xe6, 0x82, 0xa3,
	0x54, 0x14, 0xed, 0x03, 0x0f, 0xbd, 0x0d, 0xf7, 0x79, 0xbc, 0x74, 0x19, 0x3b, 0x9f, 0xde, 0xac,
	0x70, 0x26, 0x96, 0x07, 0x8c, 0x84, 0x64, 0x3e, 0xa3, 0x5c, 0xcf, 0x60, 0x1f, 0x56, 0x99, 0xa9,
	0x7f, 0xbc, 0xe4, 0x67, 0x4d, 0xf5, 0xdd, 0xdd, 0x1a, 0x45, 0xd1, 0x0f, 0xf7, 0x60, 0x64, 0x5a,
	0x46, 0x44, 0x21, 0x53, 0x2e, 0xab, 0x49, 0x80, 0x62, 0xab, 0xfc, 0x65, 0x35, 0xd3, 0x88, 0xa8,
	0x69, 0xa3, 0x99, 0x4e, 0x82, 0x69, 0xa3, 0xf9, 0xed, 0x1e, 0x99, 0x8d, 0x30, 0x8c, 0x88, 0x0a,
	0xc3, 0x46, 0x15, 0x86, 0x93, 0x20, 0x0c, 0x1b, 0xdd, 0x6c, 0x5b, 0x0f, 0xfb, 0xf0, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x74, 0xe7, 0x19, 0xc7, 0x75, 0x04, 0x00, 0x00,
}
