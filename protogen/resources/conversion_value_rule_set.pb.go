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
// source: google/ads/googleads/v17/resources/conversion_value_rule_set.proto

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

// A conversion value rule set
type ConversionValueRuleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the conversion value rule set.
	// Conversion value rule set resource names have the form:
	//
	// `customers/{customer_id}/conversionValueRuleSets/{conversion_value_rule_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the conversion value rule set.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Resource names of rules within the rule set.
	ConversionValueRules []string `protobuf:"bytes,3,rep,name=conversion_value_rules,json=conversionValueRules,proto3" json:"conversion_value_rules,omitempty"`
	// Defines dimensions for Value Rule conditions. The condition types of value
	// rules within this value rule set must be of these dimensions. The first
	// entry in this list is the primary dimension of the included value rules.
	// When using value rule primary dimension segmentation, conversion values
	// will be segmented into the values adjusted by value rules and the original
	// values, if some value rules apply.
	Dimensions []enums.ValueRuleSetDimensionEnum_ValueRuleSetDimension `protobuf:"varint,4,rep,packed,name=dimensions,proto3,enum=google.ads.googleads.v17.enums.ValueRuleSetDimensionEnum_ValueRuleSetDimension" json:"dimensions,omitempty"`
	// Output only. The resource name of the conversion value rule set's owner
	// customer. When the value rule set is inherited from a manager customer,
	// owner_customer will be the resource name of the manager whereas the
	// customer in the resource_name will be of the requesting serving customer.
	// ** Read-only **
	OwnerCustomer string `protobuf:"bytes,5,opt,name=owner_customer,json=ownerCustomer,proto3" json:"owner_customer,omitempty"`
	// Immutable. Defines the scope where the conversion value rule set is
	// attached.
	AttachmentType enums.ValueRuleSetAttachmentTypeEnum_ValueRuleSetAttachmentType `protobuf:"varint,6,opt,name=attachment_type,json=attachmentType,proto3,enum=google.ads.googleads.v17.enums.ValueRuleSetAttachmentTypeEnum_ValueRuleSetAttachmentType" json:"attachment_type,omitempty"`
	// The resource name of the campaign when the conversion value rule
	// set is attached to a campaign.
	Campaign string `protobuf:"bytes,7,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. The status of the conversion value rule set.
	// ** Read-only **
	Status enums.ConversionValueRuleSetStatusEnum_ConversionValueRuleSetStatus `protobuf:"varint,8,opt,name=status,proto3,enum=google.ads.googleads.v17.enums.ConversionValueRuleSetStatusEnum_ConversionValueRuleSetStatus" json:"status,omitempty"`
	// Immutable. The conversion action categories of the conversion value rule
	// set.
	ConversionActionCategories []enums.ConversionActionCategoryEnum_ConversionActionCategory `protobuf:"varint,9,rep,packed,name=conversion_action_categories,json=conversionActionCategories,proto3,enum=google.ads.googleads.v17.enums.ConversionActionCategoryEnum_ConversionActionCategory" json:"conversion_action_categories,omitempty"`
}

func (x *ConversionValueRuleSet) Reset() {
	*x = ConversionValueRuleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRuleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRuleSet) ProtoMessage() {}

func (x *ConversionValueRuleSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRuleSet.ProtoReflect.Descriptor instead.
func (*ConversionValueRuleSet) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescGZIP(), []int{0}
}

func (x *ConversionValueRuleSet) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ConversionValueRuleSet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConversionValueRuleSet) GetConversionValueRules() []string {
	if x != nil {
		return x.ConversionValueRules
	}
	return nil
}

func (x *ConversionValueRuleSet) GetDimensions() []enums.ValueRuleSetDimensionEnum_ValueRuleSetDimension {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *ConversionValueRuleSet) GetOwnerCustomer() string {
	if x != nil {
		return x.OwnerCustomer
	}
	return ""
}

func (x *ConversionValueRuleSet) GetAttachmentType() enums.ValueRuleSetAttachmentTypeEnum_ValueRuleSetAttachmentType {
	if x != nil {
		return x.AttachmentType
	}
	return enums.ValueRuleSetAttachmentTypeEnum_ValueRuleSetAttachmentType(0)
}

func (x *ConversionValueRuleSet) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *ConversionValueRuleSet) GetStatus() enums.ConversionValueRuleSetStatusEnum_ConversionValueRuleSetStatus {
	if x != nil {
		return x.Status
	}
	return enums.ConversionValueRuleSetStatusEnum_ConversionValueRuleSetStatus(0)
}

func (x *ConversionValueRuleSet) GetConversionActionCategories() []enums.ConversionActionCategoryEnum_ConversionActionCategory {
	if x != nil {
		return x.ConversionActionCategories
	}
	return nil
}

var File_google_ads_googleads_v17_resources_conversion_value_rule_set_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa8, 0x08, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x67,
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x31,
	0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x6f, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x0e, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0d, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x87, 0x01, 0x0a, 0x0f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x7a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x55, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x1a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x3a, 0x85, 0x01, 0xea, 0x41, 0x81, 0x01, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x4e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x8d, 0x02, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37,
	0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescData = file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_goTypes = []any{
	(*ConversionValueRuleSet)(nil),                                           // 0: google.ads.googleads.v17.resources.ConversionValueRuleSet
	(enums.ValueRuleSetDimensionEnum_ValueRuleSetDimension)(0),               // 1: google.ads.googleads.v17.enums.ValueRuleSetDimensionEnum.ValueRuleSetDimension
	(enums.ValueRuleSetAttachmentTypeEnum_ValueRuleSetAttachmentType)(0),     // 2: google.ads.googleads.v17.enums.ValueRuleSetAttachmentTypeEnum.ValueRuleSetAttachmentType
	(enums.ConversionValueRuleSetStatusEnum_ConversionValueRuleSetStatus)(0), // 3: google.ads.googleads.v17.enums.ConversionValueRuleSetStatusEnum.ConversionValueRuleSetStatus
	(enums.ConversionActionCategoryEnum_ConversionActionCategory)(0),         // 4: google.ads.googleads.v17.enums.ConversionActionCategoryEnum.ConversionActionCategory
}
var file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v17.resources.ConversionValueRuleSet.dimensions:type_name -> google.ads.googleads.v17.enums.ValueRuleSetDimensionEnum.ValueRuleSetDimension
	2, // 1: google.ads.googleads.v17.resources.ConversionValueRuleSet.attachment_type:type_name -> google.ads.googleads.v17.enums.ValueRuleSetAttachmentTypeEnum.ValueRuleSetAttachmentType
	3, // 2: google.ads.googleads.v17.resources.ConversionValueRuleSet.status:type_name -> google.ads.googleads.v17.enums.ConversionValueRuleSetStatusEnum.ConversionValueRuleSetStatus
	4, // 3: google.ads.googleads.v17.resources.ConversionValueRuleSet.conversion_action_categories:type_name -> google.ads.googleads.v17.enums.ConversionActionCategoryEnum.ConversionActionCategory
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_init() }
func file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_init() {
	if File_google_ads_googleads_v17_resources_conversion_value_rule_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConversionValueRuleSet); i {
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
			RawDescriptor: file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_conversion_value_rule_set_proto = out.File
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_conversion_value_rule_set_proto_depIdxs = nil
}
