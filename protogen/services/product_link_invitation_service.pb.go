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
// source: google/ads/googleads/v17/services/product_link_invitation_service.proto

package services

import (
	enums "github.com/Optable/google-ads-pb/protogen/enums"
	resources "github.com/Optable/google-ads-pb/protogen/resources"
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

// Request message for
// [ProductLinkInvitationService.CreateProductLinkInvitation][google.ads.googleads.v17.services.ProductLinkInvitationService.CreateProductLinkInvitation].
type CreateProductLinkInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The product link invitation to be created.
	ProductLinkInvitation *resources.ProductLinkInvitation `protobuf:"bytes,2,opt,name=product_link_invitation,json=productLinkInvitation,proto3" json:"product_link_invitation,omitempty"`
}

func (x *CreateProductLinkInvitationRequest) Reset() {
	*x = CreateProductLinkInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductLinkInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductLinkInvitationRequest) ProtoMessage() {}

func (x *CreateProductLinkInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductLinkInvitationRequest.ProtoReflect.Descriptor instead.
func (*CreateProductLinkInvitationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductLinkInvitationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateProductLinkInvitationRequest) GetProductLinkInvitation() *resources.ProductLinkInvitation {
	if x != nil {
		return x.ProductLinkInvitation
	}
	return nil
}

// Response message for product link invitation create.
type CreateProductLinkInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the product link invitation.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *CreateProductLinkInvitationResponse) Reset() {
	*x = CreateProductLinkInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductLinkInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductLinkInvitationResponse) ProtoMessage() {}

func (x *CreateProductLinkInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductLinkInvitationResponse.ProtoReflect.Descriptor instead.
func (*CreateProductLinkInvitationResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductLinkInvitationResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [ProductLinkInvitationService.UpdateProductLinkInvitation][google.ads.googleads.v17.services.ProductLinkInvitationService.UpdateProductLinkInvitation].
type UpdateProductLinkInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The product link invitation to be created.
	ProductLinkInvitationStatus enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus `protobuf:"varint,2,opt,name=product_link_invitation_status,json=productLinkInvitationStatus,proto3,enum=google.ads.googleads.v17.enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus" json:"product_link_invitation_status,omitempty"`
	// Required. Resource name of the product link invitation.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *UpdateProductLinkInvitationRequest) Reset() {
	*x = UpdateProductLinkInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductLinkInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductLinkInvitationRequest) ProtoMessage() {}

func (x *UpdateProductLinkInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductLinkInvitationRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductLinkInvitationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductLinkInvitationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UpdateProductLinkInvitationRequest) GetProductLinkInvitationStatus() enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus {
	if x != nil {
		return x.ProductLinkInvitationStatus
	}
	return enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus(0)
}

func (x *UpdateProductLinkInvitationRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Response message for product link invitation update.
type UpdateProductLinkInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the update.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *UpdateProductLinkInvitationResponse) Reset() {
	*x = UpdateProductLinkInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductLinkInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductLinkInvitationResponse) ProtoMessage() {}

func (x *UpdateProductLinkInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductLinkInvitationResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductLinkInvitationResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductLinkInvitationResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [ProductLinkinvitationService.RemoveProductLinkInvitation][].
type RemoveProductLinkInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the product link invitation being removed.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The resource name of the product link invitation being removed.
	// expected, in this format:
	//
	// ` `
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *RemoveProductLinkInvitationRequest) Reset() {
	*x = RemoveProductLinkInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductLinkInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductLinkInvitationRequest) ProtoMessage() {}

func (x *RemoveProductLinkInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductLinkInvitationRequest.ProtoReflect.Descriptor instead.
func (*RemoveProductLinkInvitationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveProductLinkInvitationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *RemoveProductLinkInvitationRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Response message for product link invitation removeal.
type RemoveProductLinkInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result for the remove request.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *RemoveProductLinkInvitationResponse) Reset() {
	*x = RemoveProductLinkInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductLinkInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductLinkInvitationResponse) ProtoMessage() {}

func (x *RemoveProductLinkInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductLinkInvitationResponse.ProtoReflect.Descriptor instead.
func (*RemoveProductLinkInvitationResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveProductLinkInvitationResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v17_services_product_link_invitation_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x76, 0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x23, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x02, 0x0a,
	0x22, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0xa5, 0x01, 0x0a, 0x1e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x5b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30,
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f,
	0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xfa, 0x41,
	0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5b, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x23, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xcb, 0x07, 0x0a, 0x1c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x02, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0xda, 0x41, 0x23, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x01, 0x2a, 0x22, 0x3c, 0x2f, 0x76, 0x31,
	0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xb1, 0x02, 0x0a, 0x1b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x82, 0x01, 0xda, 0x41, 0x38, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x01, 0x2a, 0x22,
	0x3c, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x91, 0x02,
	0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0xda, 0x41,
	0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41,
	0x3a, 0x01, 0x2a, 0x22, 0x3c, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41,
	0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8d, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x42, 0x21, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x37, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescData = file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDesc
)

func file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDescData
}

var file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_ads_googleads_v17_services_product_link_invitation_service_proto_goTypes = []any{
	(*CreateProductLinkInvitationRequest)(nil),                             // 0: google.ads.googleads.v17.services.CreateProductLinkInvitationRequest
	(*CreateProductLinkInvitationResponse)(nil),                            // 1: google.ads.googleads.v17.services.CreateProductLinkInvitationResponse
	(*UpdateProductLinkInvitationRequest)(nil),                             // 2: google.ads.googleads.v17.services.UpdateProductLinkInvitationRequest
	(*UpdateProductLinkInvitationResponse)(nil),                            // 3: google.ads.googleads.v17.services.UpdateProductLinkInvitationResponse
	(*RemoveProductLinkInvitationRequest)(nil),                             // 4: google.ads.googleads.v17.services.RemoveProductLinkInvitationRequest
	(*RemoveProductLinkInvitationResponse)(nil),                            // 5: google.ads.googleads.v17.services.RemoveProductLinkInvitationResponse
	(*resources.ProductLinkInvitation)(nil),                                // 6: google.ads.googleads.v17.resources.ProductLinkInvitation
	(enums.ProductLinkInvitationStatusEnum_ProductLinkInvitationStatus)(0), // 7: google.ads.googleads.v17.enums.ProductLinkInvitationStatusEnum.ProductLinkInvitationStatus
}
var file_google_ads_googleads_v17_services_product_link_invitation_service_proto_depIdxs = []int32{
	6, // 0: google.ads.googleads.v17.services.CreateProductLinkInvitationRequest.product_link_invitation:type_name -> google.ads.googleads.v17.resources.ProductLinkInvitation
	7, // 1: google.ads.googleads.v17.services.UpdateProductLinkInvitationRequest.product_link_invitation_status:type_name -> google.ads.googleads.v17.enums.ProductLinkInvitationStatusEnum.ProductLinkInvitationStatus
	0, // 2: google.ads.googleads.v17.services.ProductLinkInvitationService.CreateProductLinkInvitation:input_type -> google.ads.googleads.v17.services.CreateProductLinkInvitationRequest
	2, // 3: google.ads.googleads.v17.services.ProductLinkInvitationService.UpdateProductLinkInvitation:input_type -> google.ads.googleads.v17.services.UpdateProductLinkInvitationRequest
	4, // 4: google.ads.googleads.v17.services.ProductLinkInvitationService.RemoveProductLinkInvitation:input_type -> google.ads.googleads.v17.services.RemoveProductLinkInvitationRequest
	1, // 5: google.ads.googleads.v17.services.ProductLinkInvitationService.CreateProductLinkInvitation:output_type -> google.ads.googleads.v17.services.CreateProductLinkInvitationResponse
	3, // 6: google.ads.googleads.v17.services.ProductLinkInvitationService.UpdateProductLinkInvitation:output_type -> google.ads.googleads.v17.services.UpdateProductLinkInvitationResponse
	5, // 7: google.ads.googleads.v17.services.ProductLinkInvitationService.RemoveProductLinkInvitation:output_type -> google.ads.googleads.v17.services.RemoveProductLinkInvitationResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_services_product_link_invitation_service_proto_init() }
func file_google_ads_googleads_v17_services_product_link_invitation_service_proto_init() {
	if File_google_ads_googleads_v17_services_product_link_invitation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProductLinkInvitationRequest); i {
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
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProductLinkInvitationResponse); i {
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
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProductLinkInvitationRequest); i {
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
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProductLinkInvitationResponse); i {
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
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveProductLinkInvitationRequest); i {
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
		file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveProductLinkInvitationResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v17_services_product_link_invitation_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_services_product_link_invitation_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_services_product_link_invitation_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_services_product_link_invitation_service_proto = out.File
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_rawDesc = nil
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_goTypes = nil
	file_google_ads_googleads_v17_services_product_link_invitation_service_proto_depIdxs = nil
}
