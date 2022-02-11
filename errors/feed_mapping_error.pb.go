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
// source: google/ads/googleads/v10/errors/feed_mapping_error.proto

package errors

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

// Enum describing possible feed item errors.
type FeedMappingErrorEnum_FeedMappingError int32

const (
	// Enum unspecified.
	FeedMappingErrorEnum_UNSPECIFIED FeedMappingErrorEnum_FeedMappingError = 0
	// The received error code is not known in this version.
	FeedMappingErrorEnum_UNKNOWN FeedMappingErrorEnum_FeedMappingError = 1
	// The given placeholder field does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 2
	// The given criterion field does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 3
	// The given placeholder type does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 4
	// The given criterion type does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 5
	// A feed mapping must contain at least one attribute field mapping.
	FeedMappingErrorEnum_NO_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 7
	// The type of the feed attribute referenced in the attribute field mapping
	// must match the type of the placeholder field.
	FeedMappingErrorEnum_FEED_ATTRIBUTE_TYPE_MISMATCH FeedMappingErrorEnum_FeedMappingError = 8
	// A feed mapping for a system generated feed cannot be operated on.
	FeedMappingErrorEnum_CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 9
	// Only one feed mapping for a placeholder type is allowed per feed or
	// customer (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 10
	// Only one feed mapping for a criterion type is allowed per customer.
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 11
	// Only one feed attribute mapping for a placeholder field is allowed
	// (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 12
	// Only one feed attribute mapping for a criterion field is allowed
	// (depending on the criterion type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 13
	// This feed mapping may not contain any explicit attribute field mappings.
	FeedMappingErrorEnum_UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 14
	// Location placeholder feed mappings can only be created for Places feeds.
	FeedMappingErrorEnum_LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS FeedMappingErrorEnum_FeedMappingError = 15
	// Mappings for typed feeds cannot be modified.
	FeedMappingErrorEnum_CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED FeedMappingErrorEnum_FeedMappingError = 16
	// The given placeholder type can only be mapped to system generated feeds.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 17
	// The given placeholder type cannot be mapped to a system generated feed
	// with the given type.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE FeedMappingErrorEnum_FeedMappingError = 18
	// The "field" oneof was not set in an AttributeFieldMapping.
	FeedMappingErrorEnum_ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD FeedMappingErrorEnum_FeedMappingError = 19
)

// Enum value maps for FeedMappingErrorEnum_FeedMappingError.
var (
	FeedMappingErrorEnum_FeedMappingError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_PLACEHOLDER_FIELD",
		3:  "INVALID_CRITERION_FIELD",
		4:  "INVALID_PLACEHOLDER_TYPE",
		5:  "INVALID_CRITERION_TYPE",
		7:  "NO_ATTRIBUTE_FIELD_MAPPINGS",
		8:  "FEED_ATTRIBUTE_TYPE_MISMATCH",
		9:  "CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED",
		10: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE",
		11: "MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE",
		12: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD",
		13: "MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD",
		14: "UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS",
		15: "LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS",
		16: "CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED",
		17: "INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED",
		18: "INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE",
		19: "ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD",
	}
	FeedMappingErrorEnum_FeedMappingError_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"INVALID_PLACEHOLDER_FIELD":    2,
		"INVALID_CRITERION_FIELD":      3,
		"INVALID_PLACEHOLDER_TYPE":     4,
		"INVALID_CRITERION_TYPE":       5,
		"NO_ATTRIBUTE_FIELD_MAPPINGS":  7,
		"FEED_ATTRIBUTE_TYPE_MISMATCH": 8,
		"CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED":    9,
		"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE":                  10,
		"MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE":                    11,
		"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD":                 12,
		"MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD":                   13,
		"UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS":                     14,
		"LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS":              15,
		"CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED":                   16,
		"INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED":  17,
		"INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE": 18,
		"ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD":                   19,
	}
)

func (x FeedMappingErrorEnum_FeedMappingError) Enum() *FeedMappingErrorEnum_FeedMappingError {
	p := new(FeedMappingErrorEnum_FeedMappingError)
	*p = x
	return p
}

func (x FeedMappingErrorEnum_FeedMappingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedMappingErrorEnum_FeedMappingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_feed_mapping_error_proto_enumTypes[0].Descriptor()
}

func (FeedMappingErrorEnum_FeedMappingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_feed_mapping_error_proto_enumTypes[0]
}

func (x FeedMappingErrorEnum_FeedMappingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedMappingErrorEnum_FeedMappingError.Descriptor instead.
func (FeedMappingErrorEnum_FeedMappingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item errors.
type FeedMappingErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedMappingErrorEnum) Reset() {
	*x = FeedMappingErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_feed_mapping_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedMappingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedMappingErrorEnum) ProtoMessage() {}

func (x *FeedMappingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_feed_mapping_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedMappingErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedMappingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_feed_mapping_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x06, 0x0a, 0x14, 0x46, 0x65,
	0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xf9, 0x05, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10,
	0x03, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x04, 0x12,
	0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45,
	0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x4e,
	0x4f, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c,
	0x46, 0x45, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x08, 0x12, 0x38,
	0x0a, 0x34, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x09, 0x12, 0x2a, 0x0a, 0x26, 0x4d, 0x55, 0x4c, 0x54,
	0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x0a, 0x12, 0x28, 0x0a, 0x24, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45,
	0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x52,
	0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0b, 0x12, 0x2b,
	0x0a, 0x27, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49,
	0x4e, 0x47, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c,
	0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x0c, 0x12, 0x29, 0x0a, 0x25, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x10, 0x0d, 0x12, 0x27, 0x0a, 0x23, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x0e, 0x12,
	0x2e, 0x0a, 0x2a, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x53, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x53, 0x10, 0x0f, 0x12,
	0x29, 0x0a, 0x25, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59,
	0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x44, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x10, 0x12, 0x3a, 0x0a, 0x36, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x46, 0x45, 0x45, 0x44, 0x10, 0x11, 0x12, 0x3b, 0x0a, 0x37, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x12, 0x12, 0x29, 0x0a, 0x25, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x13, 0x42, 0xf5,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescData = file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_feed_mapping_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_feed_mapping_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_feed_mapping_error_proto_goTypes = []interface{}{
	(FeedMappingErrorEnum_FeedMappingError)(0), // 0: google.ads.googleads.v10.errors.FeedMappingErrorEnum.FeedMappingError
	(*FeedMappingErrorEnum)(nil),               // 1: google.ads.googleads.v10.errors.FeedMappingErrorEnum
}
var file_google_ads_googleads_v10_errors_feed_mapping_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_feed_mapping_error_proto_init() }
func file_google_ads_googleads_v10_errors_feed_mapping_error_proto_init() {
	if File_google_ads_googleads_v10_errors_feed_mapping_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_feed_mapping_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedMappingErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_feed_mapping_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_feed_mapping_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_feed_mapping_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_feed_mapping_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_feed_mapping_error_proto = out.File
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_feed_mapping_error_proto_depIdxs = nil
}
