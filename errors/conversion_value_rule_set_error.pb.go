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
// source: google/ads/googleads/v10/errors/conversion_value_rule_set_error.proto

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

// Enum describing possible conversion value rule set errors.
type ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError int32

const (
	// Enum unspecified.
	ConversionValueRuleSetErrorEnum_UNSPECIFIED ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 0
	// The received error code is not known in this version.
	ConversionValueRuleSetErrorEnum_UNKNOWN ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 1
	// Two value rules in this value rule set contain conflicting conditions.
	ConversionValueRuleSetErrorEnum_CONFLICTING_VALUE_RULE_CONDITIONS ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 2
	// This value rule set includes a value rule that cannot be found, has been
	// permanently removed or belongs to a different customer.
	ConversionValueRuleSetErrorEnum_INVALID_VALUE_RULE ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 3
	// An error that's thrown when a mutate operation is trying to
	// replace/remove some existing elements in the dimensions field. In other
	// words, ADD op is always fine and UPDATE op is fine if it's only appending
	// new elements into dimensions list.
	ConversionValueRuleSetErrorEnum_DIMENSIONS_UPDATE_ONLY_ALLOW_APPEND ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 4
	// An error that's thrown when a mutate is adding new value rule(s) into a
	// value rule set and the added value rule(s) include conditions that are
	// not specified in the dimensions of the value rule set.
	ConversionValueRuleSetErrorEnum_CONDITION_TYPE_NOT_ALLOWED ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 5
	// The dimensions field contains duplicate elements.
	ConversionValueRuleSetErrorEnum_DUPLICATE_DIMENSIONS ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 6
	// This value rule set is attached to an invalid campaign id. Either a
	// campaign with this campaign id doesn't exist or it belongs to a different
	// customer.
	ConversionValueRuleSetErrorEnum_INVALID_CAMPAIGN_ID ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 7
	// When a mutate request tries to pause a value rule set, the enabled
	// value rules in this set must be paused in the same command, or this error
	// will be thrown.
	ConversionValueRuleSetErrorEnum_CANNOT_PAUSE_UNLESS_ALL_VALUE_RULES_ARE_PAUSED ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 8
	// When a mutate request tries to pause all the value rules in a value rule
	// set, the value rule set must be paused, or this error will be thrown.
	ConversionValueRuleSetErrorEnum_SHOULD_PAUSE_WHEN_ALL_VALUE_RULES_ARE_PAUSED ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 9
	// This value rule set is attached to a campaign that does not support value
	// rules. Currently, campaign level value rule sets can only be created on
	// Search, or Display campaigns.
	ConversionValueRuleSetErrorEnum_VALUE_RULES_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 10
	// To add a value rule set that applies on Store Visits/Store Sales
	// conversion action categories, the customer must have valid Store Visits/
	// Store Sales conversion actions.
	ConversionValueRuleSetErrorEnum_INELIGIBLE_CONVERSION_ACTION_CATEGORIES ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 11
	// If NO_CONDITION is used as a dimension of a value rule set, it must be
	// the only dimension.
	ConversionValueRuleSetErrorEnum_DIMENSION_NO_CONDITION_USED_WITH_OTHER_DIMENSIONS ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 12
	// Dimension NO_CONDITION can only be used by Store Visits/Store Sales value
	// rule set.
	ConversionValueRuleSetErrorEnum_DIMENSION_NO_CONDITION_NOT_ALLOWED ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError = 13
)

// Enum value maps for ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError.
var (
	ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CONFLICTING_VALUE_RULE_CONDITIONS",
		3:  "INVALID_VALUE_RULE",
		4:  "DIMENSIONS_UPDATE_ONLY_ALLOW_APPEND",
		5:  "CONDITION_TYPE_NOT_ALLOWED",
		6:  "DUPLICATE_DIMENSIONS",
		7:  "INVALID_CAMPAIGN_ID",
		8:  "CANNOT_PAUSE_UNLESS_ALL_VALUE_RULES_ARE_PAUSED",
		9:  "SHOULD_PAUSE_WHEN_ALL_VALUE_RULES_ARE_PAUSED",
		10: "VALUE_RULES_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE",
		11: "INELIGIBLE_CONVERSION_ACTION_CATEGORIES",
		12: "DIMENSION_NO_CONDITION_USED_WITH_OTHER_DIMENSIONS",
		13: "DIMENSION_NO_CONDITION_NOT_ALLOWED",
	}
	ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError_value = map[string]int32{
		"UNSPECIFIED":                                       0,
		"UNKNOWN":                                           1,
		"CONFLICTING_VALUE_RULE_CONDITIONS":                 2,
		"INVALID_VALUE_RULE":                                3,
		"DIMENSIONS_UPDATE_ONLY_ALLOW_APPEND":               4,
		"CONDITION_TYPE_NOT_ALLOWED":                        5,
		"DUPLICATE_DIMENSIONS":                              6,
		"INVALID_CAMPAIGN_ID":                               7,
		"CANNOT_PAUSE_UNLESS_ALL_VALUE_RULES_ARE_PAUSED":    8,
		"SHOULD_PAUSE_WHEN_ALL_VALUE_RULES_ARE_PAUSED":      9,
		"VALUE_RULES_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE":       10,
		"INELIGIBLE_CONVERSION_ACTION_CATEGORIES":           11,
		"DIMENSION_NO_CONDITION_USED_WITH_OTHER_DIMENSIONS": 12,
		"DIMENSION_NO_CONDITION_NOT_ALLOWED":                13,
	}
)

func (x ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) Enum() *ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError {
	p := new(ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError)
	*p = x
	return p
}

func (x ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_enumTypes[0].Descriptor()
}

func (ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_enumTypes[0]
}

func (x ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError.Descriptor instead.
func (ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion value rule set errors.
type ConversionValueRuleSetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionValueRuleSetErrorEnum) Reset() {
	*x = ConversionValueRuleSetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRuleSetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRuleSetErrorEnum) ProtoMessage() {}

func (x *ConversionValueRuleSetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRuleSetErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionValueRuleSetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x04, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x99, 0x04, 0x0a, 0x1b, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x4e, 0x46,
	0x4c, 0x49, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x52, 0x55,
	0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12,
	0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x5f, 0x52, 0x55, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x49, 0x4d, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x04,
	0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x49,
	0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x49,
	0x44, 0x10, 0x07, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x41,
	0x55, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x53, 0x5f, 0x41, 0x52, 0x45, 0x5f, 0x50,
	0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x08, 0x12, 0x30, 0x0a, 0x2c, 0x53, 0x48, 0x4f, 0x55, 0x4c,
	0x44, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x5f, 0x57, 0x48, 0x45, 0x4e, 0x5f, 0x41, 0x4c, 0x4c,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x53, 0x5f, 0x41, 0x52, 0x45,
	0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x2f, 0x0a, 0x2b, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x2b, 0x0a, 0x27, 0x49, 0x4e,
	0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x49, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x35, 0x0a, 0x31, 0x44, 0x49, 0x4d, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x4f, 0x54, 0x48, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x0c, 0x12, 0x26,
	0x0a, 0x22, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x10, 0x0d, 0x42, 0x80, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x20,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescData = file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_goTypes = []interface{}{
	(ConversionValueRuleSetErrorEnum_ConversionValueRuleSetError)(0), // 0: google.ads.googleads.v10.errors.ConversionValueRuleSetErrorEnum.ConversionValueRuleSetError
	(*ConversionValueRuleSetErrorEnum)(nil),                          // 1: google.ads.googleads.v10.errors.ConversionValueRuleSetErrorEnum
}
var file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_init() }
func file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_init() {
	if File_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRuleSetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto = out.File
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_conversion_value_rule_set_error_proto_depIdxs = nil
}
