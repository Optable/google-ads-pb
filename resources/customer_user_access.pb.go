// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: google/ads/googleads/v11/resources/customer_user_access.proto

package resources

import (
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

// Represents the permission of a single user onto a single customer.
type CustomerUserAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Name of the resource.
	// Resource names have the form:
	// `customers/{customer_id}/customerUserAccesses/{user_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. User id of the user with the customer access.
	// Read only field
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Output only. Email address of the user.
	// Read only field
	EmailAddress *string `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3,oneof" json:"email_address,omitempty"`
	// Access role of the user.
	AccessRole enums.AccessRoleEnum_AccessRole `protobuf:"varint,4,opt,name=access_role,json=accessRole,proto3,enum=google.ads.googleads.v11.enums.AccessRoleEnum_AccessRole" json:"access_role,omitempty"`
	// Output only. The customer user access creation time.
	// Read only field
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	AccessCreationDateTime *string `protobuf:"bytes,6,opt,name=access_creation_date_time,json=accessCreationDateTime,proto3,oneof" json:"access_creation_date_time,omitempty"`
	// Output only. The email address of the inviter user.
	// Read only field
	InviterUserEmailAddress *string `protobuf:"bytes,7,opt,name=inviter_user_email_address,json=inviterUserEmailAddress,proto3,oneof" json:"inviter_user_email_address,omitempty"`
}

func (x *CustomerUserAccess) Reset() {
	*x = CustomerUserAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUserAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUserAccess) ProtoMessage() {}

func (x *CustomerUserAccess) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUserAccess.ProtoReflect.Descriptor instead.
func (*CustomerUserAccess) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerUserAccess) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerUserAccess) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CustomerUserAccess) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

func (x *CustomerUserAccess) GetAccessRole() enums.AccessRoleEnum_AccessRole {
	if x != nil {
		return x.AccessRole
	}
	return enums.AccessRoleEnum_AccessRole(0)
}

func (x *CustomerUserAccess) GetAccessCreationDateTime() string {
	if x != nil && x.AccessCreationDateTime != nil {
		return *x.AccessCreationDateTime
	}
	return ""
}

func (x *CustomerUserAccess) GetInviterUserEmailAddress() string {
	if x != nil && x.InviterUserEmailAddress != nil {
		return *x.InviterUserEmailAddress
	}
	return ""
}

var File_google_ads_googleads_v11_resources_customer_user_access_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdc, 0x04, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x33, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0c,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x5a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x19, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x16, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x45, 0x0a, 0x1a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x17, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x68, 0xea, 0x41, 0x65, 0x0a, 0x2b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x1c, 0x0a, 0x1a, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x89, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescData = file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDesc
)

func file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDescData
}

var file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_resources_customer_user_access_proto_goTypes = []interface{}{
	(*CustomerUserAccess)(nil),           // 0: google.ads.googleads.v11.resources.CustomerUserAccess
	(enums.AccessRoleEnum_AccessRole)(0), // 1: google.ads.googleads.v11.enums.AccessRoleEnum.AccessRole
}
var file_google_ads_googleads_v11_resources_customer_user_access_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.CustomerUserAccess.access_role:type_name -> google.ads.googleads.v11.enums.AccessRoleEnum.AccessRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_resources_customer_user_access_proto_init() }
func file_google_ads_googleads_v11_resources_customer_user_access_proto_init() {
	if File_google_ads_googleads_v11_resources_customer_user_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUserAccess); i {
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
	file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_resources_customer_user_access_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_resources_customer_user_access_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v11_resources_customer_user_access_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_resources_customer_user_access_proto = out.File
	file_google_ads_googleads_v11_resources_customer_user_access_proto_rawDesc = nil
	file_google_ads_googleads_v11_resources_customer_user_access_proto_goTypes = nil
	file_google_ads_googleads_v11_resources_customer_user_access_proto_depIdxs = nil
}
