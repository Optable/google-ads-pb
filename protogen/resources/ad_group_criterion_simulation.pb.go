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
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.1
// source: google/ads/googleads/v14/resources/ad_group_criterion_simulation.proto

package resources

import (
	common "github.com/Optable/google-ads-pb/protogen/common"
	enums "github.com/Optable/google-ads-pb/protogen/enums"
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

// An ad group criterion simulation. Supported combinations of advertising
// channel type, criterion type, simulation type, and simulation modification
// method are detailed below respectively. Hotel AdGroupCriterion simulation
// operations starting in V5.
//
// 1. DISPLAY - KEYWORD - CPC_BID - UNIFORM
// 2. SEARCH - KEYWORD - CPC_BID - UNIFORM
// 3. SHOPPING - LISTING_GROUP - CPC_BID - UNIFORM
// 4. HOTEL - LISTING_GROUP - CPC_BID - UNIFORM
// 5. HOTEL - LISTING_GROUP - PERCENT_CPC_BID - UNIFORM
type AdGroupCriterionSimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the ad group criterion simulation.
	// Ad group criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupCriterionSimulations/{ad_group_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. AdGroup ID of the simulation.
	AdGroupId *int64 `protobuf:"varint,9,opt,name=ad_group_id,json=adGroupId,proto3,oneof" json:"ad_group_id,omitempty"`
	// Output only. Criterion ID of the simulation.
	CriterionId *int64 `protobuf:"varint,10,opt,name=criterion_id,json=criterionId,proto3,oneof" json:"criterion_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v14.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v14.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD
	// format.
	StartDate *string `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD
	// format.
	EndDate *string `protobuf:"bytes,12,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are assignable to PointList:
	//
	//	*AdGroupCriterionSimulation_CpcBidPointList
	//	*AdGroupCriterionSimulation_PercentCpcBidPointList
	PointList isAdGroupCriterionSimulation_PointList `protobuf_oneof:"point_list"`
}

func (x *AdGroupCriterionSimulation) Reset() {
	*x = AdGroupCriterionSimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupCriterionSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupCriterionSimulation) ProtoMessage() {}

func (x *AdGroupCriterionSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupCriterionSimulation.ProtoReflect.Descriptor instead.
func (*AdGroupCriterionSimulation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupCriterionSimulation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupCriterionSimulation) GetAdGroupId() int64 {
	if x != nil && x.AdGroupId != nil {
		return *x.AdGroupId
	}
	return 0
}

func (x *AdGroupCriterionSimulation) GetCriterionId() int64 {
	if x != nil && x.CriterionId != nil {
		return *x.CriterionId
	}
	return 0
}

func (x *AdGroupCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if x != nil {
		return x.Type
	}
	return enums.SimulationTypeEnum_SimulationType(0)
}

func (x *AdGroupCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if x != nil {
		return x.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_SimulationModificationMethod(0)
}

func (x *AdGroupCriterionSimulation) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *AdGroupCriterionSimulation) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (m *AdGroupCriterionSimulation) GetPointList() isAdGroupCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (x *AdGroupCriterionSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := x.GetPointList().(*AdGroupCriterionSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

func (x *AdGroupCriterionSimulation) GetPercentCpcBidPointList() *common.PercentCpcBidSimulationPointList {
	if x, ok := x.GetPointList().(*AdGroupCriterionSimulation_PercentCpcBidPointList); ok {
		return x.PercentCpcBidPointList
	}
	return nil
}

type isAdGroupCriterionSimulation_PointList interface {
	isAdGroupCriterionSimulation_PointList()
}

type AdGroupCriterionSimulation_CpcBidPointList struct {
	// Output only. Simulation points if the simulation type is CPC_BID.
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

type AdGroupCriterionSimulation_PercentCpcBidPointList struct {
	// Output only. Simulation points if the simulation type is PERCENT_CPC_BID.
	PercentCpcBidPointList *common.PercentCpcBidSimulationPointList `protobuf:"bytes,13,opt,name=percent_cpc_bid_point_list,json=percentCpcBidPointList,proto3,oneof"`
}

func (*AdGroupCriterionSimulation_CpcBidPointList) isAdGroupCriterionSimulation_PointList() {}

func (*AdGroupCriterionSimulation_PercentCpcBidPointList) isAdGroupCriterionSimulation_PointList() {}

var File_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x08, 0x0a, 0x1a, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0xe0, 0x41, 0x03,
	0xfa, 0x41, 0x35, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x2b, 0x0a, 0x0c, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x63,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x13, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x27, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x04, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x6e, 0x0a,
	0x12, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x70, 0x63, 0x42,
	0x69, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x70,
	0x63, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x84, 0x01,
	0x0a, 0x1a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69,
	0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x70, 0x63, 0x42,
	0x69, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x16, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x70, 0x63, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x3a, 0xc1, 0x01, 0xea, 0x41, 0xbd, 0x01, 0x0a, 0x33, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x85, 0x01, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x7d,
	0x7e, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x7e, 0x7b, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x91, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1f,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x34, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescData = file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDesc
)

func file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDescData
}

var file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_goTypes = []interface{}{
	(*AdGroupCriterionSimulation)(nil),                                       // 0: google.ads.googleads.v14.resources.AdGroupCriterionSimulation
	(enums.SimulationTypeEnum_SimulationType)(0),                             // 1: google.ads.googleads.v14.enums.SimulationTypeEnum.SimulationType
	(enums.SimulationModificationMethodEnum_SimulationModificationMethod)(0), // 2: google.ads.googleads.v14.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	(*common.CpcBidSimulationPointList)(nil),                                 // 3: google.ads.googleads.v14.common.CpcBidSimulationPointList
	(*common.PercentCpcBidSimulationPointList)(nil),                          // 4: google.ads.googleads.v14.common.PercentCpcBidSimulationPointList
}
var file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v14.resources.AdGroupCriterionSimulation.type:type_name -> google.ads.googleads.v14.enums.SimulationTypeEnum.SimulationType
	2, // 1: google.ads.googleads.v14.resources.AdGroupCriterionSimulation.modification_method:type_name -> google.ads.googleads.v14.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	3, // 2: google.ads.googleads.v14.resources.AdGroupCriterionSimulation.cpc_bid_point_list:type_name -> google.ads.googleads.v14.common.CpcBidSimulationPointList
	4, // 3: google.ads.googleads.v14.resources.AdGroupCriterionSimulation.percent_cpc_bid_point_list:type_name -> google.ads.googleads.v14.common.PercentCpcBidSimulationPointList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_init() }
func file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_init() {
	if File_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupCriterionSimulation); i {
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
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AdGroupCriterionSimulation_CpcBidPointList)(nil),
		(*AdGroupCriterionSimulation_PercentCpcBidPointList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto = out.File
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_rawDesc = nil
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_goTypes = nil
	file_google_ads_googleads_v14_resources_ad_group_criterion_simulation_proto_depIdxs = nil
}
