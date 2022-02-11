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
// source: google/ads/googleads/v10/errors/mutate_error.proto

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

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
	// The resource isn't in Google Ads. It belongs to another ads system.
	MutateErrorEnum_RESOURCE_NOT_IN_GOOGLE_ADS MutateErrorEnum_MutateError = 10
	// The resource being created already exists.
	MutateErrorEnum_RESOURCE_ALREADY_EXISTS MutateErrorEnum_MutateError = 11
	// This resource cannot be used with "validate_only".
	MutateErrorEnum_RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY MutateErrorEnum_MutateError = 12
	// This operation cannot be used with "partial_failure".
	MutateErrorEnum_OPERATION_DOES_NOT_SUPPORT_PARTIAL_FAILURE MutateErrorEnum_MutateError = 16
	// Attempt to write to read-only fields.
	MutateErrorEnum_RESOURCE_READ_ONLY MutateErrorEnum_MutateError = 13
)

// Enum value maps for MutateErrorEnum_MutateError.
var (
	MutateErrorEnum_MutateError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "RESOURCE_NOT_FOUND",
		7:  "ID_EXISTS_IN_MULTIPLE_MUTATES",
		8:  "INCONSISTENT_FIELD_VALUES",
		9:  "MUTATE_NOT_ALLOWED",
		10: "RESOURCE_NOT_IN_GOOGLE_ADS",
		11: "RESOURCE_ALREADY_EXISTS",
		12: "RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY",
		16: "OPERATION_DOES_NOT_SUPPORT_PARTIAL_FAILURE",
		13: "RESOURCE_READ_ONLY",
	}
	MutateErrorEnum_MutateError_value = map[string]int32{
		"UNSPECIFIED":                                0,
		"UNKNOWN":                                    1,
		"RESOURCE_NOT_FOUND":                         3,
		"ID_EXISTS_IN_MULTIPLE_MUTATES":              7,
		"INCONSISTENT_FIELD_VALUES":                  8,
		"MUTATE_NOT_ALLOWED":                         9,
		"RESOURCE_NOT_IN_GOOGLE_ADS":                 10,
		"RESOURCE_ALREADY_EXISTS":                    11,
		"RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY":    12,
		"OPERATION_DOES_NOT_SUPPORT_PARTIAL_FAILURE": 16,
		"RESOURCE_READ_ONLY":                         13,
	}
)

func (x MutateErrorEnum_MutateError) Enum() *MutateErrorEnum_MutateError {
	p := new(MutateErrorEnum_MutateError)
	*p = x
	return p
}

func (x MutateErrorEnum_MutateError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MutateErrorEnum_MutateError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_mutate_error_proto_enumTypes[0].Descriptor()
}

func (MutateErrorEnum_MutateError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_mutate_error_proto_enumTypes[0]
}

func (x MutateErrorEnum_MutateError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MutateErrorEnum_MutateError.Descriptor instead.
func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MutateErrorEnum) Reset() {
	*x = MutateErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_mutate_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateErrorEnum) ProtoMessage() {}

func (x *MutateErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_mutate_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateErrorEnum.ProtoReflect.Descriptor instead.
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_mutate_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_mutate_error_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x0f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xcf, 0x02, 0x0a, 0x0b, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x21, 0x0a,
	0x1d, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x4d, 0x55,
	0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x53, 0x10, 0x07,
	0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x08, 0x12,
	0x16, 0x0a, 0x12, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x53, 0x10, 0x0b, 0x12, 0x2b, 0x0a, 0x27, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10,
	0x0c, 0x12, 0x2e, 0x0a, 0x2a, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x10, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x0d, 0x42, 0xf0, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x42, 0x10, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescData = file_google_ads_googleads_v10_errors_mutate_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_mutate_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_mutate_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_mutate_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_mutate_error_proto_goTypes = []interface{}{
	(MutateErrorEnum_MutateError)(0), // 0: google.ads.googleads.v10.errors.MutateErrorEnum.MutateError
	(*MutateErrorEnum)(nil),          // 1: google.ads.googleads.v10.errors.MutateErrorEnum
}
var file_google_ads_googleads_v10_errors_mutate_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_mutate_error_proto_init() }
func file_google_ads_googleads_v10_errors_mutate_error_proto_init() {
	if File_google_ads_googleads_v10_errors_mutate_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_mutate_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_mutate_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_mutate_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_mutate_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_mutate_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_mutate_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_mutate_error_proto = out.File
	file_google_ads_googleads_v10_errors_mutate_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_mutate_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_mutate_error_proto_depIdxs = nil
}
