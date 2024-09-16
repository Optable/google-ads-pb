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
// source: google/ads/googleads/v17/enums/android_privacy_interaction_type.proto

package enums

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

// Enumerates interaction types
type AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType int32

const (
	// Not specified.
	AndroidPrivacyInteractionTypeEnum_UNSPECIFIED AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType = 0
	// The value is unknown in this version.
	AndroidPrivacyInteractionTypeEnum_UNKNOWN AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType = 1
	// The physical click interaction type.
	AndroidPrivacyInteractionTypeEnum_CLICK AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType = 2
	// The 10 seconds engaged view interaction type.
	AndroidPrivacyInteractionTypeEnum_ENGAGED_VIEW AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType = 3
	// The view (ad impression) interaction type.
	AndroidPrivacyInteractionTypeEnum_VIEW AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType = 4
)

// Enum value maps for AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType.
var (
	AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CLICK",
		3: "ENGAGED_VIEW",
		4: "VIEW",
	}
	AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType_value = map[string]int32{
		"UNSPECIFIED":  0,
		"UNKNOWN":      1,
		"CLICK":        2,
		"ENGAGED_VIEW": 3,
		"VIEW":         4,
	}
)

func (x AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) Enum() *AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType {
	p := new(AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType)
	*p = x
	return p
}

func (x AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_enumTypes[0].Descriptor()
}

func (AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_enumTypes[0]
}

func (x AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType.Descriptor instead.
func (AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescGZIP(), []int{0, 0}
}

// The interaction type enum for Android privacy shared key.
type AndroidPrivacyInteractionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AndroidPrivacyInteractionTypeEnum) Reset() {
	*x = AndroidPrivacyInteractionTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidPrivacyInteractionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidPrivacyInteractionTypeEnum) ProtoMessage() {}

func (x *AndroidPrivacyInteractionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidPrivacyInteractionTypeEnum.ProtoReflect.Descriptor instead.
func (*AndroidPrivacyInteractionTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x21, 0x41, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x64, 0x0a,
	0x1d, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x47, 0x41, 0x47,
	0x45, 0x44, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x04, 0x42, 0xfc, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x22, 0x41, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescData = file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDesc
)

func file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDescData
}

var file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_goTypes = []any{
	(AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType)(0), // 0: google.ads.googleads.v17.enums.AndroidPrivacyInteractionTypeEnum.AndroidPrivacyInteractionType
	(*AndroidPrivacyInteractionTypeEnum)(nil),                            // 1: google.ads.googleads.v17.enums.AndroidPrivacyInteractionTypeEnum
}
var file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_init() }
func file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_init() {
	if File_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AndroidPrivacyInteractionTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto = out.File
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_rawDesc = nil
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_goTypes = nil
	file_google_ads_googleads_v17_enums_android_privacy_interaction_type_proto_depIdxs = nil
}
