// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: google/ads/googleads/v17/resources/change_status.proto

package resources

import (
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

// Describes the status of returned resource. ChangeStatus could have up to 3
// minutes delay to reflect a new change.
type ChangeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Time at which the most recent change has occurred on this
	// resource.
	LastChangeDateTime *string `protobuf:"bytes,24,opt,name=last_change_date_time,json=lastChangeDateTime,proto3,oneof" json:"last_change_date_time,omitempty"`
	// Output only. Represents the type of the changed resource. This dictates
	// what fields will be set. For example, for AD_GROUP, campaign and ad_group
	// fields will be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v17.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// Output only. The Campaign affected by this change.
	Campaign *string `protobuf:"bytes,17,opt,name=campaign,proto3,oneof" json:"campaign,omitempty"`
	// Output only. The AdGroup affected by this change.
	AdGroup *string `protobuf:"bytes,18,opt,name=ad_group,json=adGroup,proto3,oneof" json:"ad_group,omitempty"`
	// Output only. Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v17.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// Output only. The AdGroupAd affected by this change.
	AdGroupAd *string `protobuf:"bytes,25,opt,name=ad_group_ad,json=adGroupAd,proto3,oneof" json:"ad_group_ad,omitempty"`
	// Output only. The AdGroupCriterion affected by this change.
	AdGroupCriterion *string `protobuf:"bytes,26,opt,name=ad_group_criterion,json=adGroupCriterion,proto3,oneof" json:"ad_group_criterion,omitempty"`
	// Output only. The CampaignCriterion affected by this change.
	CampaignCriterion *string `protobuf:"bytes,27,opt,name=campaign_criterion,json=campaignCriterion,proto3,oneof" json:"campaign_criterion,omitempty"`
	// Output only. The Feed affected by this change.
	Feed *string `protobuf:"bytes,28,opt,name=feed,proto3,oneof" json:"feed,omitempty"`
	// Output only. The FeedItem affected by this change.
	FeedItem *string `protobuf:"bytes,29,opt,name=feed_item,json=feedItem,proto3,oneof" json:"feed_item,omitempty"`
	// Output only. The AdGroupFeed affected by this change.
	AdGroupFeed *string `protobuf:"bytes,30,opt,name=ad_group_feed,json=adGroupFeed,proto3,oneof" json:"ad_group_feed,omitempty"`
	// Output only. The CampaignFeed affected by this change.
	CampaignFeed *string `protobuf:"bytes,31,opt,name=campaign_feed,json=campaignFeed,proto3,oneof" json:"campaign_feed,omitempty"`
	// Output only. The AdGroupBidModifier affected by this change.
	AdGroupBidModifier *string `protobuf:"bytes,32,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3,oneof" json:"ad_group_bid_modifier,omitempty"`
	// Output only. The SharedSet affected by this change.
	SharedSet string `protobuf:"bytes,33,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// Output only. The CampaignSharedSet affected by this change.
	CampaignSharedSet string `protobuf:"bytes,34,opt,name=campaign_shared_set,json=campaignSharedSet,proto3" json:"campaign_shared_set,omitempty"`
	// Output only. The Asset affected by this change.
	Asset string `protobuf:"bytes,35,opt,name=asset,proto3" json:"asset,omitempty"`
	// Output only. The CustomerAsset affected by this change.
	CustomerAsset string `protobuf:"bytes,36,opt,name=customer_asset,json=customerAsset,proto3" json:"customer_asset,omitempty"`
	// Output only. The CampaignAsset affected by this change.
	CampaignAsset string `protobuf:"bytes,37,opt,name=campaign_asset,json=campaignAsset,proto3" json:"campaign_asset,omitempty"`
	// Output only. The AdGroupAsset affected by this change.
	AdGroupAsset string `protobuf:"bytes,38,opt,name=ad_group_asset,json=adGroupAsset,proto3" json:"ad_group_asset,omitempty"`
	// Output only. The CombinedAudience affected by this change.
	CombinedAudience string `protobuf:"bytes,40,opt,name=combined_audience,json=combinedAudience,proto3" json:"combined_audience,omitempty"`
	// Output only. The AssetGroup affected by this change.
	AssetGroup string `protobuf:"bytes,41,opt,name=asset_group,json=assetGroup,proto3" json:"asset_group,omitempty"`
}

func (x *ChangeStatus) Reset() {
	*x = ChangeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_change_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatus) ProtoMessage() {}

func (x *ChangeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_change_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatus.ProtoReflect.Descriptor instead.
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_change_status_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeStatus) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ChangeStatus) GetLastChangeDateTime() string {
	if x != nil && x.LastChangeDateTime != nil {
		return *x.LastChangeDateTime
	}
	return ""
}

func (x *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType(0)
}

func (x *ChangeStatus) GetCampaign() string {
	if x != nil && x.Campaign != nil {
		return *x.Campaign
	}
	return ""
}

func (x *ChangeStatus) GetAdGroup() string {
	if x != nil && x.AdGroup != nil {
		return *x.AdGroup
	}
	return ""
}

func (x *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if x != nil {
		return x.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_ChangeStatusOperation(0)
}

func (x *ChangeStatus) GetAdGroupAd() string {
	if x != nil && x.AdGroupAd != nil {
		return *x.AdGroupAd
	}
	return ""
}

func (x *ChangeStatus) GetAdGroupCriterion() string {
	if x != nil && x.AdGroupCriterion != nil {
		return *x.AdGroupCriterion
	}
	return ""
}

func (x *ChangeStatus) GetCampaignCriterion() string {
	if x != nil && x.CampaignCriterion != nil {
		return *x.CampaignCriterion
	}
	return ""
}

func (x *ChangeStatus) GetFeed() string {
	if x != nil && x.Feed != nil {
		return *x.Feed
	}
	return ""
}

func (x *ChangeStatus) GetFeedItem() string {
	if x != nil && x.FeedItem != nil {
		return *x.FeedItem
	}
	return ""
}

func (x *ChangeStatus) GetAdGroupFeed() string {
	if x != nil && x.AdGroupFeed != nil {
		return *x.AdGroupFeed
	}
	return ""
}

func (x *ChangeStatus) GetCampaignFeed() string {
	if x != nil && x.CampaignFeed != nil {
		return *x.CampaignFeed
	}
	return ""
}

func (x *ChangeStatus) GetAdGroupBidModifier() string {
	if x != nil && x.AdGroupBidModifier != nil {
		return *x.AdGroupBidModifier
	}
	return ""
}

func (x *ChangeStatus) GetSharedSet() string {
	if x != nil {
		return x.SharedSet
	}
	return ""
}

func (x *ChangeStatus) GetCampaignSharedSet() string {
	if x != nil {
		return x.CampaignSharedSet
	}
	return ""
}

func (x *ChangeStatus) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *ChangeStatus) GetCustomerAsset() string {
	if x != nil {
		return x.CustomerAsset
	}
	return ""
}

func (x *ChangeStatus) GetCampaignAsset() string {
	if x != nil {
		return x.CampaignAsset
	}
	return ""
}

func (x *ChangeStatus) GetAdGroupAsset() string {
	if x != nil {
		return x.AdGroupAsset
	}
	return ""
}

func (x *ChangeStatus) GetCombinedAudience() string {
	if x != nil {
		return x.CombinedAudience
	}
	return ""
}

func (x *ChangeStatus) GetAssetGroup() string {
	if x != nil {
		return x.AssetGroup
	}
	return ""
}

var File_google_ads_googleads_v17_resources_change_status_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_change_status_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x11, 0x0a, 0x0c, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2d, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x7f, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48, 0x01, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x08, 0x61, 0x64, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa, 0x41,
	0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x48, 0x02, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01,
	0x01, 0x12, 0x7d, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x4f, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x48, 0x03, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x64, 0x0a, 0x12, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x48, 0x04, 0x52, 0x10, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x66, 0x0a, 0x12, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x48, 0x05, 0x52, 0x11, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x3e, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x46, 0x65, 0x65, 0x64, 0x48, 0x06, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x4b, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x1d, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x07, 0x52,
	0x08, 0x66, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0d,
	0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65,
	0x64, 0x48, 0x08, 0x52, 0x0b, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x57, 0x0a, 0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x03, 0xfa,
	0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x48, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x6b, 0x0a, 0x15,
	0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x03,
	0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x48, 0x0a, 0x52, 0x12, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x49, 0x0a, 0x0a, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x12, 0x62, 0x0a, 0x13, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x22, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x32, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x11, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x20, 0x0a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x55, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e,
	0xe0, 0x41, 0x03, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x55, 0x0a,
	0x0e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x25, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x53, 0x0a, 0x0e, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41,
	0x03, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0c, 0x61, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x5e, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x64, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b,
	0xe0, 0x41, 0x03, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x25, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x37, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x18, 0x0a, 0x16,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x61, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x83, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v17_resources_change_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_change_status_proto_rawDescData = file_google_ads_googleads_v17_resources_change_status_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_change_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_change_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_change_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_change_status_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_change_status_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_change_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_resources_change_status_proto_goTypes = []any{
	(*ChangeStatus)(nil), // 0: google.ads.googleads.v17.resources.ChangeStatus
	(enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType)(0), // 1: google.ads.googleads.v17.enums.ChangeStatusResourceTypeEnum.ChangeStatusResourceType
	(enums.ChangeStatusOperationEnum_ChangeStatusOperation)(0),       // 2: google.ads.googleads.v17.enums.ChangeStatusOperationEnum.ChangeStatusOperation
}
var file_google_ads_googleads_v17_resources_change_status_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v17.resources.ChangeStatus.resource_type:type_name -> google.ads.googleads.v17.enums.ChangeStatusResourceTypeEnum.ChangeStatusResourceType
	2, // 1: google.ads.googleads.v17.resources.ChangeStatus.resource_status:type_name -> google.ads.googleads.v17.enums.ChangeStatusOperationEnum.ChangeStatusOperation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_change_status_proto_init() }
func file_google_ads_googleads_v17_resources_change_status_proto_init() {
	if File_google_ads_googleads_v17_resources_change_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_change_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeStatus); i {
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
	file_google_ads_googleads_v17_resources_change_status_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_resources_change_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_change_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_change_status_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_change_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_change_status_proto = out.File
	file_google_ads_googleads_v17_resources_change_status_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_change_status_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_change_status_proto_depIdxs = nil
}
