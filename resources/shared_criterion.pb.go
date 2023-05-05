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
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: google/ads/googleads/v13/resources/shared_criterion.proto

package resources

import (
	common "github.com/shenzhencenter/google-ads-pb/common"
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// A criterion belonging to a shared set.
type SharedCriterion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the shared criterion.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedCriteria/{shared_set_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The shared set to which the shared criterion belongs.
	SharedSet *string `protobuf:"bytes,10,opt,name=shared_set,json=sharedSet,proto3,oneof" json:"shared_set,omitempty"`
	// Output only. The ID of the criterion.
	//
	// This field is ignored for mutates.
	CriterionId *int64 `protobuf:"varint,11,opt,name=criterion_id,json=criterionId,proto3,oneof" json:"criterion_id,omitempty"`
	// Output only. The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v13.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The criterion.
	//
	// Exactly one must be set.
	//
	// Types that are assignable to Criterion:
	//
	//	*SharedCriterion_Keyword
	//	*SharedCriterion_YoutubeVideo
	//	*SharedCriterion_YoutubeChannel
	//	*SharedCriterion_Placement
	//	*SharedCriterion_MobileAppCategory
	//	*SharedCriterion_MobileApplication
	Criterion isSharedCriterion_Criterion `protobuf_oneof:"criterion"`
}

func (x *SharedCriterion) Reset() {
	*x = SharedCriterion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedCriterion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedCriterion) ProtoMessage() {}

func (x *SharedCriterion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedCriterion.ProtoReflect.Descriptor instead.
func (*SharedCriterion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescGZIP(), []int{0}
}

func (x *SharedCriterion) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *SharedCriterion) GetSharedSet() string {
	if x != nil && x.SharedSet != nil {
		return *x.SharedSet
	}
	return ""
}

func (x *SharedCriterion) GetCriterionId() int64 {
	if x != nil && x.CriterionId != nil {
		return *x.CriterionId
	}
	return 0
}

func (x *SharedCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if x != nil {
		return x.Type
	}
	return enums.CriterionTypeEnum_CriterionType(0)
}

func (m *SharedCriterion) GetCriterion() isSharedCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (x *SharedCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (x *SharedCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (x *SharedCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (x *SharedCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (x *SharedCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (x *SharedCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := x.GetCriterion().(*SharedCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

type isSharedCriterion_Criterion interface {
	isSharedCriterion_Criterion()
}

type SharedCriterion_Keyword struct {
	// Immutable. Keyword.
	Keyword *common.KeywordInfo `protobuf:"bytes,3,opt,name=keyword,proto3,oneof"`
}

type SharedCriterion_YoutubeVideo struct {
	// Immutable. YouTube Video.
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,5,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type SharedCriterion_YoutubeChannel struct {
	// Immutable. YouTube Channel.
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,6,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type SharedCriterion_Placement struct {
	// Immutable. Placement.
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type SharedCriterion_MobileAppCategory struct {
	// Immutable. Mobile App Category.
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,8,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type SharedCriterion_MobileApplication struct {
	// Immutable. Mobile application.
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,9,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

func (*SharedCriterion_Keyword) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeVideo) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeChannel) isSharedCriterion_Criterion() {}

func (*SharedCriterion_Placement) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileAppCategory) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileApplication) isSharedCriterion_Criterion() {}

var File_google_ads_googleads_v13_resources_shared_criterion_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa7, 0x08, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2a, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48, 0x01, 0x52, 0x09, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x63,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x5d, 0x0a, 0x0d, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x6f, 0x75, 0x54, 0x75,
	0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x48, 0x00, 0x52, 0x0c, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x63, 0x0a, 0x0f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x6f, 0x75, 0x54,
	0x75, 0x62, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x13, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x41, 0x70, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70,
	0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x6c, 0x0a, 0x12, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x74, 0xea, 0x41, 0x71, 0x0a, 0x28, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2f, 0x7b,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b,
	0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x0a,
	0x09, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x86, 0x02, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x14, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescData = file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDesc
)

func file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDescData
}

var file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_resources_shared_criterion_proto_goTypes = []interface{}{
	(*SharedCriterion)(nil),                    // 0: google.ads.googleads.v13.resources.SharedCriterion
	(enums.CriterionTypeEnum_CriterionType)(0), // 1: google.ads.googleads.v13.enums.CriterionTypeEnum.CriterionType
	(*common.KeywordInfo)(nil),                 // 2: google.ads.googleads.v13.common.KeywordInfo
	(*common.YouTubeVideoInfo)(nil),            // 3: google.ads.googleads.v13.common.YouTubeVideoInfo
	(*common.YouTubeChannelInfo)(nil),          // 4: google.ads.googleads.v13.common.YouTubeChannelInfo
	(*common.PlacementInfo)(nil),               // 5: google.ads.googleads.v13.common.PlacementInfo
	(*common.MobileAppCategoryInfo)(nil),       // 6: google.ads.googleads.v13.common.MobileAppCategoryInfo
	(*common.MobileApplicationInfo)(nil),       // 7: google.ads.googleads.v13.common.MobileApplicationInfo
}
var file_google_ads_googleads_v13_resources_shared_criterion_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v13.resources.SharedCriterion.type:type_name -> google.ads.googleads.v13.enums.CriterionTypeEnum.CriterionType
	2, // 1: google.ads.googleads.v13.resources.SharedCriterion.keyword:type_name -> google.ads.googleads.v13.common.KeywordInfo
	3, // 2: google.ads.googleads.v13.resources.SharedCriterion.youtube_video:type_name -> google.ads.googleads.v13.common.YouTubeVideoInfo
	4, // 3: google.ads.googleads.v13.resources.SharedCriterion.youtube_channel:type_name -> google.ads.googleads.v13.common.YouTubeChannelInfo
	5, // 4: google.ads.googleads.v13.resources.SharedCriterion.placement:type_name -> google.ads.googleads.v13.common.PlacementInfo
	6, // 5: google.ads.googleads.v13.resources.SharedCriterion.mobile_app_category:type_name -> google.ads.googleads.v13.common.MobileAppCategoryInfo
	7, // 6: google.ads.googleads.v13.resources.SharedCriterion.mobile_application:type_name -> google.ads.googleads.v13.common.MobileApplicationInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_resources_shared_criterion_proto_init() }
func file_google_ads_googleads_v13_resources_shared_criterion_proto_init() {
	if File_google_ads_googleads_v13_resources_shared_criterion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedCriterion); i {
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
	file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SharedCriterion_Keyword)(nil),
		(*SharedCriterion_YoutubeVideo)(nil),
		(*SharedCriterion_YoutubeChannel)(nil),
		(*SharedCriterion_Placement)(nil),
		(*SharedCriterion_MobileAppCategory)(nil),
		(*SharedCriterion_MobileApplication)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_resources_shared_criterion_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_resources_shared_criterion_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v13_resources_shared_criterion_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_resources_shared_criterion_proto = out.File
	file_google_ads_googleads_v13_resources_shared_criterion_proto_rawDesc = nil
	file_google_ads_googleads_v13_resources_shared_criterion_proto_goTypes = nil
	file_google_ads_googleads_v13_resources_shared_criterion_proto_depIdxs = nil
}
