// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v9/resources/ad.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/revealbot/google-ads-go/common"
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

// An ad.
type Ad struct {
	// The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v9.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v9.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_DynamicSearchAd
	//	*Ad_ResponsiveDisplayAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	//	*Ad_VideoAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae4d7986ec9bb9b, []int{0}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_DynamicSearchAd struct {
	DynamicSearchAd *common.DynamicSearchAdInfo `protobuf:"bytes,8,opt,name=dynamic_search_ad,json=dynamicSearchAd,proto3,oneof"`
}

type Ad_ResponsiveDisplayAd struct {
	ResponsiveDisplayAd *common.ResponsiveDisplayAdInfo `protobuf:"bytes,9,opt,name=responsive_display_ad,json=responsiveDisplayAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

type Ad_VideoAd struct {
	VideoAd *common.VideoAdInfo `protobuf:"bytes,24,opt,name=video_ad,json=videoAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_DynamicSearchAd) isAd_AdData() {}

func (*Ad_ResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (*Ad_VideoAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetDynamicSearchAd() *common.DynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_DynamicSearchAd); ok {
		return x.DynamicSearchAd
	}
	return nil
}

func (m *Ad) GetResponsiveDisplayAd() *common.ResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveDisplayAd); ok {
		return x.ResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

func (m *Ad) GetVideoAd() *common.VideoAdInfo {
	if x, ok := m.GetAdData().(*Ad_VideoAd); ok {
		return x.VideoAd
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ad) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_DynamicSearchAd)(nil),
		(*Ad_ResponsiveDisplayAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
		(*Ad_VideoAd)(nil),
	}
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v9.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v9/resources/ad.proto", fileDescriptor_6ae4d7986ec9bb9b)
}

var fileDescriptor_6ae4d7986ec9bb9b = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0x23, 0x35,
	0x18, 0x86, 0x49, 0xb6, 0x34, 0x5b, 0x77, 0xe9, 0x8f, 0x43, 0xd9, 0x51, 0x41, 0xa8, 0x8b, 0xb4,
	0x52, 0xd5, 0x4a, 0x93, 0x28, 0xcb, 0x82, 0x94, 0x8a, 0x83, 0x09, 0xad, 0xda, 0x22, 0x21, 0xa2,
	0xb4, 0xcd, 0x01, 0x2a, 0x8c, 0xdc, 0xb1, 0x33, 0x1d, 0xe1, 0xb1, 0x47, 0xf6, 0x4c, 0x68, 0xb8,
	0x12, 0x8e, 0x39, 0xe4, 0x52, 0xb8, 0x14, 0x8e, 0xb9, 0x00, 0xe4, 0xcf, 0x33, 0x8e, 0xba, 0xfd,
	0x99, 0x39, 0xf3, 0xcf, 0xfb, 0xbc, 0xef, 0xd7, 0xcf, 0x8e, 0x3b, 0xe8, 0x20, 0x96, 0x32, 0xe6,
	0xac, 0x47, 0xa8, 0xee, 0xd9, 0xa1, 0x19, 0xcd, 0xfb, 0x3d, 0xc5, 0xb4, 0x2c, 0x54, 0xc4, 0x74,
	0x8f, 0x50, 0x3f, 0x53, 0x32, 0x97, 0xf8, 0x8d, 0x15, 0xf8, 0x84, 0x6a, 0xdf, 0x69, 0xfd, 0x79,
	0xdf, 0x77, 0xda, 0xdd, 0xc1, 0x53, 0x76, 0x91, 0x4c, 0x53, 0x29, 0x7a, 0x84, 0x86, 0xf9, 0x22,
	0x63, 0x61, 0x22, 0x66, 0x52, 0x5b, 0xdb, 0xdd, 0xf7, 0x35, 0x4c, 0x54, 0xe8, 0x5c, 0xa6, 0x61,
	0x46, 0x14, 0x49, 0x59, 0xce, 0x54, 0x89, 0x1d, 0x3e, 0x85, 0x31, 0x51, 0xa4, 0xba, 0x4a, 0x2a,
	0xc5, 0x07, 0xcf, 0x8b, 0x29, 0x9b, 0x27, 0x51, 0xa5, 0xfd, 0xb2, 0xd4, 0xc2, 0xec, 0xa6, 0x98,
	0xf5, 0x7e, 0x57, 0x24, 0xcb, 0x98, 0x2a, 0xeb, 0xfd, 0xea, 0xcf, 0x0d, 0xd4, 0x0e, 0x28, 0x3e,
	0x44, 0xed, 0x84, 0x7a, 0xad, 0xbd, 0xd6, 0xfe, 0xfa, 0xe0, 0xf3, 0xb2, 0x1f, 0x7e, 0xc5, 0xf8,
	0xe7, 0x22, 0xff, 0xe6, 0xeb, 0x29, 0xe1, 0x05, 0x9b, 0xb4, 0x13, 0x8a, 0x8f, 0x10, 0x9a, 0x25,
	0x82, 0xf0, 0xb0, 0x50, 0x5c, 0x7b, 0xed, 0xbd, 0x17, 0xfb, 0xeb, 0x83, 0x2f, 0x1e, 0x40, 0x17,
	0xb9, 0x4a, 0x44, 0x6c, 0xa9, 0x35, 0xd0, 0x5f, 0x29, 0xae, 0xf1, 0x19, 0xda, 0xb6, 0x70, 0x2a,
	0x6f, 0x12, 0xce, 0xac, 0xc7, 0x56, 0x03, 0x8f, 0x4d, 0xc0, 0x7e, 0x04, 0x0a, 0x9c, 0xc6, 0x68,
	0x27, 0x57, 0x24, 0xfa, 0x2d, 0x11, 0xb1, 0x71, 0x09, 0x73, 0x96, 0x66, 0x9c, 0xe4, 0xcc, 0x7b,
	0x05, 0x7f, 0xc6, 0xf3, 0x6e, 0xdd, 0x0a, 0xbd, 0x52, 0xfc, 0xb2, 0x04, 0x71, 0x84, 0x76, 0x8c,
	0xd1, 0x87, 0x67, 0xa4, 0x3d, 0x04, 0xf5, 0xf5, 0xfc, 0xa7, 0xee, 0x8c, 0x3d, 0x5c, 0xff, 0x7b,
	0x00, 0xc7, 0x15, 0x37, 0xe9, 0x16, 0x8a, 0x7f, 0xb0, 0xa6, 0xf1, 0x77, 0x68, 0x9d, 0x26, 0x3a,
	0xe3, 0x64, 0x61, 0xaa, 0xf6, 0x56, 0x1a, 0x14, 0x8b, 0x4a, 0xe0, 0x4a, 0x71, 0x7c, 0x8c, 0x56,
	0xcc, 0x55, 0xf0, 0x3e, 0xde, 0x6b, 0xed, 0x6f, 0x0c, 0xfa, 0x4f, 0x96, 0x04, 0x77, 0xc1, 0x0f,
	0xe8, 0xe5, 0x22, 0x63, 0x27, 0xa2, 0x48, 0xcb, 0xe1, 0x04, 0x68, 0x7c, 0x8e, 0xba, 0x84, 0x52,
	0x46, 0xc3, 0x9b, 0x45, 0x68, 0xb1, 0x90, 0x50, 0xed, 0x75, 0xa1, 0x98, 0xdd, 0x07, 0xc5, 0x8c,
	0xa4, 0xe4, 0xb6, 0x94, 0x2d, 0xc0, 0x46, 0x8b, 0x53, 0x50, 0x04, 0x54, 0xe3, 0x5f, 0xd0, 0xb6,
	0xbd, 0x71, 0x61, 0xa6, 0xd8, 0x8c, 0x29, 0x26, 0x22, 0xe6, 0x7d, 0xda, 0xa8, 0xba, 0x63, 0xe0,
	0xa0, 0x3a, 0x3b, 0x9c, 0x6c, 0x59, 0xab, 0xb1, 0x73, 0xc2, 0x7d, 0xb4, 0x22, 0x48, 0xca, 0xbc,
	0xd7, 0x0d, 0xfa, 0x04, 0x4a, 0x7c, 0x82, 0x3a, 0x39, 0xbb, 0xcb, 0x43, 0x42, 0xbd, 0x55, 0x80,
	0x0e, 0xea, 0xce, 0xed, 0x92, 0xdd, 0xe5, 0x01, 0x3d, 0x17, 0x33, 0x79, 0xf6, 0xd1, 0x64, 0x35,
	0x87, 0x19, 0xfe, 0x15, 0x6d, 0xb1, 0xbb, 0x8c, 0x08, 0xd3, 0xa5, 0xca, 0xaf, 0x03, 0x7e, 0x83,
	0x3a, 0xbf, 0x93, 0x92, 0xbb, 0xe7, 0xbb, 0xc1, 0xee, 0xad, 0x62, 0x82, 0xb6, 0xe9, 0x42, 0x90,
	0x34, 0x89, 0x42, 0xcd, 0x88, 0x8a, 0x6e, 0x4d, 0xc0, 0x4b, 0x08, 0x78, 0x57, 0x17, 0x70, 0x6c,
	0xc1, 0x0b, 0xe0, 0x5c, 0xc2, 0x26, 0xbd, 0xbf, 0x8c, 0x53, 0xb4, 0xa3, 0x98, 0xce, 0xa4, 0xd0,
	0xc9, 0x9c, 0x85, 0xd5, 0xad, 0x23, 0xd4, 0x5b, 0x83, 0x98, 0x6f, 0xeb, 0x62, 0x26, 0x0e, 0x3e,
	0xb6, 0xac, 0x8b, 0xea, 0xaa, 0x87, 0x5b, 0x78, 0x82, 0x5e, 0x45, 0x84, 0xf3, 0x50, 0x0a, 0x0e,
	0x29, 0x9f, 0x40, 0x8a, 0x5f, 0xfb, 0xab, 0x21, 0x9c, 0xff, 0x24, 0xf8, 0xd2, 0x1c, 0x45, 0x6e,
	0x05, 0xff, 0x81, 0x76, 0xdd, 0x29, 0x3c, 0x6c, 0xd7, 0x06, 0x24, 0x1c, 0x35, 0x3d, 0x8f, 0xc7,
	0xdb, 0xf6, 0x9a, 0x3d, 0xbe, 0x8d, 0xcf, 0xd0, 0xcb, 0x5b, 0x99, 0x33, 0x6e, 0x92, 0x36, 0x21,
	0xe9, 0xb0, 0x2e, 0xe9, 0xcc, 0xe8, 0x9d, 0x73, 0xe7, 0xd6, 0x4e, 0xcd, 0x59, 0xeb, 0x5b, 0x99,
	0x65, 0xe6, 0xa9, 0xd2, 0x29, 0x51, 0x70, 0x99, 0xb6, 0x9b, 0x9d, 0xf5, 0x45, 0x09, 0x5e, 0x18,
	0x6e, 0x79, 0xd6, 0xfa, 0xfe, 0x32, 0x8e, 0x51, 0xd7, 0x45, 0x64, 0x4a, 0xd2, 0x22, 0x82, 0x10,
	0x0c, 0x21, 0xef, 0x9b, 0x86, 0x8c, 0x2d, 0xe9, 0x62, 0x5c, 0xd9, 0x6e, 0xc3, 0x74, 0x25, 0x4e,
	0x49, 0x02, 0x5d, 0xd9, 0x69, 0xd6, 0x95, 0x53, 0xa3, 0x5f, 0x76, 0x25, 0xb6, 0x53, 0xe3, 0x94,
	0xa4, 0x24, 0x36, 0x4f, 0x8f, 0xf7, 0x59, 0x33, 0xa7, 0x73, 0xa3, 0x5f, 0x3a, 0x25, 0x76, 0x6a,
	0x9c, 0xe6, 0x09, 0x65, 0xd2, 0x38, 0x79, 0xcd, 0x9c, 0xa6, 0x46, 0xbf, 0x74, 0x9a, 0xdb, 0xe9,
	0x68, 0x0d, 0x75, 0x08, 0x0d, 0x29, 0xc9, 0xc9, 0xe8, 0xbf, 0x16, 0x7a, 0x1b, 0xc9, 0xd4, 0xaf,
	0xfd, 0x50, 0x18, 0x75, 0x02, 0x3a, 0x36, 0xef, 0xd1, 0xb8, 0xf5, 0xf3, 0x0f, 0xa5, 0x3a, 0x96,
	0x9c, 0x88, 0xd8, 0x97, 0x2a, 0xee, 0xc5, 0x4c, 0xc0, 0x6b, 0x55, 0xfd, 0xaf, 0xce, 0x12, 0xfd,
	0xcc, 0x17, 0xca, 0x91, 0x1b, 0xfd, 0xd5, 0x7e, 0x71, 0x1a, 0x04, 0x7f, 0xb7, 0xdf, 0xd8, 0xb7,
	0xd6, 0x0f, 0xa8, 0xf6, 0xdd, 0xb3, 0xeb, 0x4f, 0xfb, 0xe6, 0xe7, 0x69, 0x95, 0xff, 0x54, 0x9a,
	0xeb, 0x80, 0xea, 0x6b, 0xa7, 0xb9, 0x9e, 0xf6, 0xaf, 0x9d, 0xe6, 0xdf, 0xf6, 0x5b, 0xbb, 0x31,
	0x1c, 0x06, 0x54, 0x0f, 0x87, 0x4e, 0x35, 0x1c, 0x4e, 0xfb, 0xc3, 0xa1, 0xd3, 0xdd, 0xac, 0x42,
	0xb1, 0xef, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xc6, 0x60, 0x4c, 0x4d, 0x09, 0x00, 0x00,
}
