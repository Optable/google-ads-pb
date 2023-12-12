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
// source: google/ads/googleads/v15/resources/android_privacy_shared_key_google_network_type.proto

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

// An Android privacy shared key view for Google network type key.
type AndroidPrivacySharedKeyGoogleNetworkType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the Android privacy shared key.
	// Android privacy shared key resource names have the form:
	//
	// `customers/{customer_id}/androidPrivacySharedKeyGoogleNetworkTypes/{campaign_id}~{android_privacy_interaction_type}~{android_privacy_network_type}~{android_privacy_interaction_date(yyyy-mm-dd)}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign ID used in the share key encoding.
	CampaignId int64 `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Output only. The interaction type enum used in the share key encoding.
	AndroidPrivacyInteractionType enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType `protobuf:"varint,3,opt,name=android_privacy_interaction_type,json=androidPrivacyInteractionType,proto3,enum=google.ads.googleads.v15.enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType" json:"android_privacy_interaction_type,omitempty"`
	// Output only. The interaction date used in the shared key encoding in the
	// format of "YYYY-MM-DD" in UTC timezone.
	AndroidPrivacyInteractionDate string `protobuf:"bytes,4,opt,name=android_privacy_interaction_date,json=androidPrivacyInteractionDate,proto3" json:"android_privacy_interaction_date,omitempty"`
	// Output only. The network type enum used in the share key encoding.
	AndroidPrivacyNetworkType enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType `protobuf:"varint,5,opt,name=android_privacy_network_type,json=androidPrivacyNetworkType,proto3,enum=google.ads.googleads.v15.enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType" json:"android_privacy_network_type,omitempty"`
	// Output only. 128 bit hex string of the encoded shared network type key,
	// including a '0x' prefix. This key can be used to do a bitwise OR operator
	// with the aggregate conversion key to create a full aggregation key to
	// retrieve the Aggregate API Report in Android Privacy Sandbox.
	SharedNetworkTypeKey string `protobuf:"bytes,6,opt,name=shared_network_type_key,json=sharedNetworkTypeKey,proto3" json:"shared_network_type_key,omitempty"`
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) Reset() {
	*x = AndroidPrivacySharedKeyGoogleNetworkType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidPrivacySharedKeyGoogleNetworkType) ProtoMessage() {}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidPrivacySharedKeyGoogleNetworkType.ProtoReflect.Descriptor instead.
func (*AndroidPrivacySharedKeyGoogleNetworkType) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescGZIP(), []int{0}
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetAndroidPrivacyInteractionType() enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType {
	if x != nil {
		return x.AndroidPrivacyInteractionType
	}
	return enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType(0)
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetAndroidPrivacyInteractionDate() string {
	if x != nil {
		return x.AndroidPrivacyInteractionDate
	}
	return ""
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetAndroidPrivacyNetworkType() enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType {
	if x != nil {
		return x.AndroidPrivacyNetworkType
	}
	return enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType(0)
}

func (x *AndroidPrivacySharedKeyGoogleNetworkType) GetSharedNetworkTypeKey() string {
	if x != nil {
		return x.SharedNetworkTypeKey
	}
	return ""
}

var File_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDesc = []byte{
	0x0a, 0x57, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x45, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x07, 0x0a, 0x28, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x6e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x49, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x43, 0x0a,
	0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65,
	0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0xad, 0x01, 0x0a, 0x20, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x5f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x20, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1c, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x19, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x17, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4b, 0x65, 0x79,
	0x3a, 0xfe, 0x01, 0xea, 0x41, 0xfa, 0x01, 0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0xb4, 0x01, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e,
	0x7b, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b,
	0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x7d, 0x42, 0x9f, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x2d, 0x41, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescData = file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDesc
)

func file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDescData
}

var file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_goTypes = []interface{}{
	(*AndroidPrivacySharedKeyGoogleNetworkType)(nil),                           // 0: google.ads.googleads.v15.resources.AndroidPrivacySharedKeyGoogleNetworkType
	(enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType)(0), // 1: google.ads.googleads.v15.enums.AndroidPrivacyInteractionTypeEnum.AndroidPrivacyInteractionType
	(enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType)(0),         // 2: google.ads.googleads.v15.enums.AndroidPrivacyNetworkTypeEnum.AndroidPrivacyNetworkType
}
var file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v15.resources.AndroidPrivacySharedKeyGoogleNetworkType.android_privacy_interaction_type:type_name -> google.ads.googleads.v15.enums.AndroidPrivacyInteractionTypeEnum.AndroidPrivacyInteractionType
	2, // 1: google.ads.googleads.v15.resources.AndroidPrivacySharedKeyGoogleNetworkType.android_privacy_network_type:type_name -> google.ads.googleads.v15.enums.AndroidPrivacyNetworkTypeEnum.AndroidPrivacyNetworkType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_init()
}
func file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_init() {
	if File_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidPrivacySharedKeyGoogleNetworkType); i {
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
			RawDescriptor: file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto = out.File
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_rawDesc = nil
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_goTypes = nil
	file_google_ads_googleads_v15_resources_android_privacy_shared_key_google_network_type_proto_depIdxs = nil
}
