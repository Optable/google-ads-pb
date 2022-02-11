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
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: google/ads/googleads/v10/resources/payments_account.proto

package resources

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

// A payments account, which can be used to set up billing for an Ads customer.
type PaymentsAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the payments account.
	// PaymentsAccount resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. A 16 digit ID used to identify a payments account.
	PaymentsAccountId *string `protobuf:"bytes,8,opt,name=payments_account_id,json=paymentsAccountId,proto3,oneof" json:"payments_account_id,omitempty"`
	// Output only. The name of the payments account.
	Name *string `protobuf:"bytes,9,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The currency code of the payments account.
	// A subset of the currency codes derived from the ISO 4217 standard is
	// supported.
	CurrencyCode *string `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3,oneof" json:"currency_code,omitempty"`
	// Output only. A 12 digit ID used to identify the payments profile associated with the
	// payments account.
	PaymentsProfileId *string `protobuf:"bytes,11,opt,name=payments_profile_id,json=paymentsProfileId,proto3,oneof" json:"payments_profile_id,omitempty"`
	// Output only. A secondary payments profile ID present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *string `protobuf:"bytes,12,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3,oneof" json:"secondary_payments_profile_id,omitempty"`
	// Output only. Paying manager of this payment account.
	PayingManagerCustomer *string `protobuf:"bytes,13,opt,name=paying_manager_customer,json=payingManagerCustomer,proto3,oneof" json:"paying_manager_customer,omitempty"`
}

func (x *PaymentsAccount) Reset() {
	*x = PaymentsAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentsAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentsAccount) ProtoMessage() {}

func (x *PaymentsAccount) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentsAccount.ProtoReflect.Descriptor instead.
func (*PaymentsAccount) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_payments_account_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentsAccount) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *PaymentsAccount) GetPaymentsAccountId() string {
	if x != nil && x.PaymentsAccountId != nil {
		return *x.PaymentsAccountId
	}
	return ""
}

func (x *PaymentsAccount) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PaymentsAccount) GetCurrencyCode() string {
	if x != nil && x.CurrencyCode != nil {
		return *x.CurrencyCode
	}
	return ""
}

func (x *PaymentsAccount) GetPaymentsProfileId() string {
	if x != nil && x.PaymentsProfileId != nil {
		return *x.PaymentsProfileId
	}
	return ""
}

func (x *PaymentsAccount) GetSecondaryPaymentsProfileId() string {
	if x != nil && x.SecondaryPaymentsProfileId != nil {
		return *x.SecondaryPaymentsProfileId
	}
	return ""
}

func (x *PaymentsAccount) GetPayingManagerCustomer() string {
	if x != nil && x.PayingManagerCustomer != nil {
		return *x.PayingManagerCustomer
	}
	return ""
}

var File_google_ads_googleads_v10_resources_payments_account_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_payments_account_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x05, 0x0a, 0x0f, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x13, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52,
	0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x1d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x04, 0x52, 0x1a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x66, 0x0a, 0x17, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x05,
	0x52, 0x15, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x3a, 0x6d, 0xea, 0x41, 0x6a, 0x0a,
	0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x70, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x86, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x14, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30,
	0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_payments_account_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_payments_account_proto_rawDescData = file_google_ads_googleads_v10_resources_payments_account_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_payments_account_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_payments_account_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_payments_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_payments_account_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_payments_account_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_payments_account_proto_goTypes = []interface{}{
	(*PaymentsAccount)(nil), // 0: google.ads.googleads.v10.resources.PaymentsAccount
}
var file_google_ads_googleads_v10_resources_payments_account_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_payments_account_proto_init() }
func file_google_ads_googleads_v10_resources_payments_account_proto_init() {
	if File_google_ads_googleads_v10_resources_payments_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentsAccount); i {
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
	file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_payments_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_payments_account_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_payments_account_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_payments_account_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_payments_account_proto = out.File
	file_google_ads_googleads_v10_resources_payments_account_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_payments_account_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_payments_account_proto_depIdxs = nil
}
