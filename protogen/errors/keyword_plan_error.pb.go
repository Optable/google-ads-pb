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
// source: google/ads/googleads/v17/errors/keyword_plan_error.proto

package errors

import (
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

// Enum describing possible errors from applying a keyword plan.
type KeywordPlanErrorEnum_KeywordPlanError int32

const (
	// Enum unspecified.
	KeywordPlanErrorEnum_UNSPECIFIED KeywordPlanErrorEnum_KeywordPlanError = 0
	// The received error code is not known in this version.
	KeywordPlanErrorEnum_UNKNOWN KeywordPlanErrorEnum_KeywordPlanError = 1
	// The plan's bid multiplier value is outside the valid range.
	KeywordPlanErrorEnum_BID_MULTIPLIER_OUT_OF_RANGE KeywordPlanErrorEnum_KeywordPlanError = 2
	// The plan's bid value is too high.
	KeywordPlanErrorEnum_BID_TOO_HIGH KeywordPlanErrorEnum_KeywordPlanError = 3
	// The plan's bid value is too low.
	KeywordPlanErrorEnum_BID_TOO_LOW KeywordPlanErrorEnum_KeywordPlanError = 4
	// The plan's cpc bid is not a multiple of the minimum billable unit.
	KeywordPlanErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS KeywordPlanErrorEnum_KeywordPlanError = 5
	// The plan's daily budget value is too low.
	KeywordPlanErrorEnum_DAILY_BUDGET_TOO_LOW KeywordPlanErrorEnum_KeywordPlanError = 6
	// The plan's daily budget is not a multiple of the minimum billable unit.
	KeywordPlanErrorEnum_DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS KeywordPlanErrorEnum_KeywordPlanError = 7
	// The input has an invalid value.
	KeywordPlanErrorEnum_INVALID_VALUE KeywordPlanErrorEnum_KeywordPlanError = 8
	// The plan has no keyword.
	KeywordPlanErrorEnum_KEYWORD_PLAN_HAS_NO_KEYWORDS KeywordPlanErrorEnum_KeywordPlanError = 9
	// The plan is not enabled and API cannot provide mutation, forecast or
	// stats.
	KeywordPlanErrorEnum_KEYWORD_PLAN_NOT_ENABLED KeywordPlanErrorEnum_KeywordPlanError = 10
	// The requested plan cannot be found for providing forecast or stats.
	KeywordPlanErrorEnum_KEYWORD_PLAN_NOT_FOUND KeywordPlanErrorEnum_KeywordPlanError = 11
	// The plan is missing a cpc bid.
	KeywordPlanErrorEnum_MISSING_BID KeywordPlanErrorEnum_KeywordPlanError = 13
	// The plan is missing required forecast_period field.
	KeywordPlanErrorEnum_MISSING_FORECAST_PERIOD KeywordPlanErrorEnum_KeywordPlanError = 14
	// The plan's forecast_period has invalid forecast date range.
	KeywordPlanErrorEnum_INVALID_FORECAST_DATE_RANGE KeywordPlanErrorEnum_KeywordPlanError = 15
	// The plan's name is invalid.
	KeywordPlanErrorEnum_INVALID_NAME KeywordPlanErrorEnum_KeywordPlanError = 16
)

// Enum value maps for KeywordPlanErrorEnum_KeywordPlanError.
var (
	KeywordPlanErrorEnum_KeywordPlanError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "BID_MULTIPLIER_OUT_OF_RANGE",
		3:  "BID_TOO_HIGH",
		4:  "BID_TOO_LOW",
		5:  "BID_TOO_MANY_FRACTIONAL_DIGITS",
		6:  "DAILY_BUDGET_TOO_LOW",
		7:  "DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS",
		8:  "INVALID_VALUE",
		9:  "KEYWORD_PLAN_HAS_NO_KEYWORDS",
		10: "KEYWORD_PLAN_NOT_ENABLED",
		11: "KEYWORD_PLAN_NOT_FOUND",
		13: "MISSING_BID",
		14: "MISSING_FORECAST_PERIOD",
		15: "INVALID_FORECAST_DATE_RANGE",
		16: "INVALID_NAME",
	}
	KeywordPlanErrorEnum_KeywordPlanError_value = map[string]int32{
		"UNSPECIFIED":                             0,
		"UNKNOWN":                                 1,
		"BID_MULTIPLIER_OUT_OF_RANGE":             2,
		"BID_TOO_HIGH":                            3,
		"BID_TOO_LOW":                             4,
		"BID_TOO_MANY_FRACTIONAL_DIGITS":          5,
		"DAILY_BUDGET_TOO_LOW":                    6,
		"DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS": 7,
		"INVALID_VALUE":                           8,
		"KEYWORD_PLAN_HAS_NO_KEYWORDS":            9,
		"KEYWORD_PLAN_NOT_ENABLED":                10,
		"KEYWORD_PLAN_NOT_FOUND":                  11,
		"MISSING_BID":                             13,
		"MISSING_FORECAST_PERIOD":                 14,
		"INVALID_FORECAST_DATE_RANGE":             15,
		"INVALID_NAME":                            16,
	}
)

func (x KeywordPlanErrorEnum_KeywordPlanError) Enum() *KeywordPlanErrorEnum_KeywordPlanError {
	p := new(KeywordPlanErrorEnum_KeywordPlanError)
	*p = x
	return p
}

func (x KeywordPlanErrorEnum_KeywordPlanError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeywordPlanErrorEnum_KeywordPlanError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_keyword_plan_error_proto_enumTypes[0].Descriptor()
}

func (KeywordPlanErrorEnum_KeywordPlanError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_keyword_plan_error_proto_enumTypes[0]
}

func (x KeywordPlanErrorEnum_KeywordPlanError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeywordPlanErrorEnum_KeywordPlanError.Descriptor instead.
func (KeywordPlanErrorEnum_KeywordPlanError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// resource (keyword plan, keyword plan campaign, keyword plan ad group or
// keyword plan keyword) or KeywordPlanService RPC.
type KeywordPlanErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeywordPlanErrorEnum) Reset() {
	*x = KeywordPlanErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_keyword_plan_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanErrorEnum) ProtoMessage() {}

func (x *KeywordPlanErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_keyword_plan_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanErrorEnum.ProtoReflect.Descriptor instead.
func (*KeywordPlanErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_keyword_plan_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xc8, 0x03, 0x0a, 0x14,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0xaf, 0x03, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x49, 0x44, 0x5f, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x49, 0x45, 0x52, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x49, 0x44, 0x5f,
	0x54, 0x4f, 0x4f, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49,
	0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x42,
	0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x46, 0x52, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x47, 0x49, 0x54, 0x53, 0x10, 0x05, 0x12,
	0x18, 0x0a, 0x14, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f,
	0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x06, 0x12, 0x2b, 0x0a, 0x27, 0x44, 0x41, 0x49,
	0x4c, 0x59, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41,
	0x4e, 0x59, 0x5f, 0x46, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x44, 0x49,
	0x47, 0x49, 0x54, 0x53, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x08, 0x12, 0x20, 0x0a, 0x1c, 0x4b, 0x45, 0x59,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4e, 0x4f,
	0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x4b, 0x45, 0x59,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x42, 0x49, 0x44, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x46, 0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f,
	0x44, 0x10, 0x0e, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x10, 0x42, 0xf5, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x37, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescData = file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_keyword_plan_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_keyword_plan_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_keyword_plan_error_proto_goTypes = []any{
	(KeywordPlanErrorEnum_KeywordPlanError)(0), // 0: google.ads.googleads.v17.errors.KeywordPlanErrorEnum.KeywordPlanError
	(*KeywordPlanErrorEnum)(nil),               // 1: google.ads.googleads.v17.errors.KeywordPlanErrorEnum
}
var file_google_ads_googleads_v17_errors_keyword_plan_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_keyword_plan_error_proto_init() }
func file_google_ads_googleads_v17_errors_keyword_plan_error_proto_init() {
	if File_google_ads_googleads_v17_errors_keyword_plan_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_keyword_plan_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KeywordPlanErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_keyword_plan_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_keyword_plan_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_keyword_plan_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_keyword_plan_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_keyword_plan_error_proto = out.File
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_keyword_plan_error_proto_depIdxs = nil
}
