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
// source: google/ads/googleads/v19/resources/ad_group_simulation.proto

package resources

import (
	common "github.com/revealbot/google-ads-go/common"
	enums "github.com/revealbot/google-ads-go/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An ad group simulation. Supported combinations of advertising
// channel type, simulation type and simulation modification method is
// detailed below respectively.
//
// 1. SEARCH - CPC_BID - DEFAULT
// 2. SEARCH - CPC_BID - UNIFORM
// 3. SEARCH - TARGET_CPA - UNIFORM
// 4. SEARCH - TARGET_ROAS - UNIFORM
// 5. DISPLAY - CPC_BID - DEFAULT
// 6. DISPLAY - CPC_BID - UNIFORM
// 7. DISPLAY - TARGET_CPA - UNIFORM
type AdGroupSimulation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the ad group simulation.
	// Ad group simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupSimulations/{ad_group_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Ad group id of the simulation.
	AdGroupId *int64 `protobuf:"varint,12,opt,name=ad_group_id,json=adGroupId,proto3,oneof" json:"ad_group_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v19.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v19.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD
	// format.
	StartDate *string `protobuf:"bytes,13,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD
	// format
	EndDate *string `protobuf:"bytes,14,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//
	//	*AdGroupSimulation_CpcBidPointList
	//	*AdGroupSimulation_CpvBidPointList
	//	*AdGroupSimulation_TargetCpaPointList
	//	*AdGroupSimulation_TargetRoasPointList
	PointList     isAdGroupSimulation_PointList `protobuf_oneof:"point_list"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupSimulation) Reset() {
	*x = AdGroupSimulation{}
	mi := &file_resources_ad_group_simulation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupSimulation) ProtoMessage() {}

func (x *AdGroupSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_simulation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupSimulation.ProtoReflect.Descriptor instead.
func (*AdGroupSimulation) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupSimulation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupSimulation) GetAdGroupId() int64 {
	if x != nil && x.AdGroupId != nil {
		return *x.AdGroupId
	}
	return 0
}

func (x *AdGroupSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if x != nil {
		return x.Type
	}
	return enums.SimulationTypeEnum_SimulationType(0)
}

func (x *AdGroupSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if x != nil {
		return x.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_SimulationModificationMethod(0)
}

func (x *AdGroupSimulation) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *AdGroupSimulation) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (x *AdGroupSimulation) GetPointList() isAdGroupSimulation_PointList {
	if x != nil {
		return x.PointList
	}
	return nil
}

func (x *AdGroupSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x != nil {
		if x, ok := x.PointList.(*AdGroupSimulation_CpcBidPointList); ok {
			return x.CpcBidPointList
		}
	}
	return nil
}

func (x *AdGroupSimulation) GetCpvBidPointList() *common.CpvBidSimulationPointList {
	if x != nil {
		if x, ok := x.PointList.(*AdGroupSimulation_CpvBidPointList); ok {
			return x.CpvBidPointList
		}
	}
	return nil
}

func (x *AdGroupSimulation) GetTargetCpaPointList() *common.TargetCpaSimulationPointList {
	if x != nil {
		if x, ok := x.PointList.(*AdGroupSimulation_TargetCpaPointList); ok {
			return x.TargetCpaPointList
		}
	}
	return nil
}

func (x *AdGroupSimulation) GetTargetRoasPointList() *common.TargetRoasSimulationPointList {
	if x != nil {
		if x, ok := x.PointList.(*AdGroupSimulation_TargetRoasPointList); ok {
			return x.TargetRoasPointList
		}
	}
	return nil
}

type isAdGroupSimulation_PointList interface {
	isAdGroupSimulation_PointList()
}

type AdGroupSimulation_CpcBidPointList struct {
	// Output only. Simulation points if the simulation type is CPC_BID.
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

type AdGroupSimulation_CpvBidPointList struct {
	// Output only. Simulation points if the simulation type is CPV_BID.
	CpvBidPointList *common.CpvBidSimulationPointList `protobuf:"bytes,10,opt,name=cpv_bid_point_list,json=cpvBidPointList,proto3,oneof"`
}

type AdGroupSimulation_TargetCpaPointList struct {
	// Output only. Simulation points if the simulation type is TARGET_CPA.
	TargetCpaPointList *common.TargetCpaSimulationPointList `protobuf:"bytes,9,opt,name=target_cpa_point_list,json=targetCpaPointList,proto3,oneof"`
}

type AdGroupSimulation_TargetRoasPointList struct {
	// Output only. Simulation points if the simulation type is TARGET_ROAS.
	TargetRoasPointList *common.TargetRoasSimulationPointList `protobuf:"bytes,11,opt,name=target_roas_point_list,json=targetRoasPointList,proto3,oneof"`
}

func (*AdGroupSimulation_CpcBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_CpvBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_TargetCpaPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_TargetRoasPointList) isAdGroupSimulation_PointList() {}

var File_resources_ad_group_simulation_proto protoreflect.FileDescriptor

var file_resources_ad_group_simulation_proto_rawDesc = string([]byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x09, 0x0a, 0x11,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x57, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c,
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x61, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x93, 0x01, 0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x02, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x6e, 0x0a, 0x12, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x70, 0x63, 0x42, 0x69, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x70, 0x63, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x6e, 0x0a, 0x12, 0x63, 0x70, 0x76, 0x5f, 0x62, 0x69, 0x64, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x70, 0x76, 0x42, 0x69, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x70, 0x76, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x70, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x61, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x70, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x7a, 0x0a,
	0x16, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x61, 0x73, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x00, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x9f, 0x01, 0xea, 0x41, 0x9b, 0x01,
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x7d, 0x7e, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d,
	0x7e, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x88, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x16, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_resources_ad_group_simulation_proto_rawDescOnce sync.Once
	file_resources_ad_group_simulation_proto_rawDescData []byte
)

func file_resources_ad_group_simulation_proto_rawDescGZIP() []byte {
	file_resources_ad_group_simulation_proto_rawDescOnce.Do(func() {
		file_resources_ad_group_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resources_ad_group_simulation_proto_rawDesc), len(file_resources_ad_group_simulation_proto_rawDesc)))
	})
	return file_resources_ad_group_simulation_proto_rawDescData
}

var file_resources_ad_group_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_ad_group_simulation_proto_goTypes = []any{
	(*AdGroupSimulation)(nil),                                                // 0: google.ads.googleads.v19.resources.AdGroupSimulation
	(enums.SimulationTypeEnum_SimulationType)(0),                             // 1: google.ads.googleads.v19.enums.SimulationTypeEnum.SimulationType
	(enums.SimulationModificationMethodEnum_SimulationModificationMethod)(0), // 2: google.ads.googleads.v19.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	(*common.CpcBidSimulationPointList)(nil),                                 // 3: google.ads.googleads.v19.common.CpcBidSimulationPointList
	(*common.CpvBidSimulationPointList)(nil),                                 // 4: google.ads.googleads.v19.common.CpvBidSimulationPointList
	(*common.TargetCpaSimulationPointList)(nil),                              // 5: google.ads.googleads.v19.common.TargetCpaSimulationPointList
	(*common.TargetRoasSimulationPointList)(nil),                             // 6: google.ads.googleads.v19.common.TargetRoasSimulationPointList
}
var file_resources_ad_group_simulation_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.resources.AdGroupSimulation.type:type_name -> google.ads.googleads.v19.enums.SimulationTypeEnum.SimulationType
	2, // 1: google.ads.googleads.v19.resources.AdGroupSimulation.modification_method:type_name -> google.ads.googleads.v19.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	3, // 2: google.ads.googleads.v19.resources.AdGroupSimulation.cpc_bid_point_list:type_name -> google.ads.googleads.v19.common.CpcBidSimulationPointList
	4, // 3: google.ads.googleads.v19.resources.AdGroupSimulation.cpv_bid_point_list:type_name -> google.ads.googleads.v19.common.CpvBidSimulationPointList
	5, // 4: google.ads.googleads.v19.resources.AdGroupSimulation.target_cpa_point_list:type_name -> google.ads.googleads.v19.common.TargetCpaSimulationPointList
	6, // 5: google.ads.googleads.v19.resources.AdGroupSimulation.target_roas_point_list:type_name -> google.ads.googleads.v19.common.TargetRoasSimulationPointList
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_resources_ad_group_simulation_proto_init() }
func file_resources_ad_group_simulation_proto_init() {
	if File_resources_ad_group_simulation_proto != nil {
		return
	}
	file_resources_ad_group_simulation_proto_msgTypes[0].OneofWrappers = []any{
		(*AdGroupSimulation_CpcBidPointList)(nil),
		(*AdGroupSimulation_CpvBidPointList)(nil),
		(*AdGroupSimulation_TargetCpaPointList)(nil),
		(*AdGroupSimulation_TargetRoasPointList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resources_ad_group_simulation_proto_rawDesc), len(file_resources_ad_group_simulation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_ad_group_simulation_proto_goTypes,
		DependencyIndexes: file_resources_ad_group_simulation_proto_depIdxs,
		MessageInfos:      file_resources_ad_group_simulation_proto_msgTypes,
	}.Build()
	File_resources_ad_group_simulation_proto = out.File
	file_resources_ad_group_simulation_proto_goTypes = nil
	file_resources_ad_group_simulation_proto_depIdxs = nil
}
