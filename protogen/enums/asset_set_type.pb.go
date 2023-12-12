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
// source: google/ads/googleads/v14/enums/asset_set_type.proto

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

// Possible types of an asset set.
type AssetSetTypeEnum_AssetSetType int32

const (
	// Not specified.
	AssetSetTypeEnum_UNSPECIFIED AssetSetTypeEnum_AssetSetType = 0
	// Used for return value only. Represents value unknown in this version.
	AssetSetTypeEnum_UNKNOWN AssetSetTypeEnum_AssetSetType = 1
	// Page asset set.
	AssetSetTypeEnum_PAGE_FEED AssetSetTypeEnum_AssetSetType = 2
	// Dynamic education asset set.
	AssetSetTypeEnum_DYNAMIC_EDUCATION AssetSetTypeEnum_AssetSetType = 3
	// Google Merchant Center asset set.
	AssetSetTypeEnum_MERCHANT_CENTER_FEED AssetSetTypeEnum_AssetSetType = 4
	// Dynamic real estate asset set.
	AssetSetTypeEnum_DYNAMIC_REAL_ESTATE AssetSetTypeEnum_AssetSetType = 5
	// Dynamic custom asset set.
	AssetSetTypeEnum_DYNAMIC_CUSTOM AssetSetTypeEnum_AssetSetType = 6
	// Dynamic hotels and rentals asset set.
	AssetSetTypeEnum_DYNAMIC_HOTELS_AND_RENTALS AssetSetTypeEnum_AssetSetType = 7
	// Dynamic flights asset set.
	AssetSetTypeEnum_DYNAMIC_FLIGHTS AssetSetTypeEnum_AssetSetType = 8
	// Dynamic travel asset set.
	AssetSetTypeEnum_DYNAMIC_TRAVEL AssetSetTypeEnum_AssetSetType = 9
	// Dynamic local asset set.
	AssetSetTypeEnum_DYNAMIC_LOCAL AssetSetTypeEnum_AssetSetType = 10
	// Dynamic jobs asset set.
	AssetSetTypeEnum_DYNAMIC_JOBS AssetSetTypeEnum_AssetSetType = 11
	// Location sync level asset set.
	AssetSetTypeEnum_LOCATION_SYNC AssetSetTypeEnum_AssetSetType = 12
	// Business Profile location group asset set.
	AssetSetTypeEnum_BUSINESS_PROFILE_DYNAMIC_LOCATION_GROUP AssetSetTypeEnum_AssetSetType = 13
	// Chain location group asset set which can be used for both owned
	// locations and affiliate locations.
	AssetSetTypeEnum_CHAIN_DYNAMIC_LOCATION_GROUP AssetSetTypeEnum_AssetSetType = 14
	// Static location group asset set which can be used for both owned
	// locations and affiliate locations.
	AssetSetTypeEnum_STATIC_LOCATION_GROUP AssetSetTypeEnum_AssetSetType = 15
	// Hotel Property asset set which is used to link a hotel property feed to
	// Performance Max for travel goals campaigns.
	AssetSetTypeEnum_HOTEL_PROPERTY AssetSetTypeEnum_AssetSetType = 16
)

// Enum value maps for AssetSetTypeEnum_AssetSetType.
var (
	AssetSetTypeEnum_AssetSetType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PAGE_FEED",
		3:  "DYNAMIC_EDUCATION",
		4:  "MERCHANT_CENTER_FEED",
		5:  "DYNAMIC_REAL_ESTATE",
		6:  "DYNAMIC_CUSTOM",
		7:  "DYNAMIC_HOTELS_AND_RENTALS",
		8:  "DYNAMIC_FLIGHTS",
		9:  "DYNAMIC_TRAVEL",
		10: "DYNAMIC_LOCAL",
		11: "DYNAMIC_JOBS",
		12: "LOCATION_SYNC",
		13: "BUSINESS_PROFILE_DYNAMIC_LOCATION_GROUP",
		14: "CHAIN_DYNAMIC_LOCATION_GROUP",
		15: "STATIC_LOCATION_GROUP",
		16: "HOTEL_PROPERTY",
	}
	AssetSetTypeEnum_AssetSetType_value = map[string]int32{
		"UNSPECIFIED":                0,
		"UNKNOWN":                    1,
		"PAGE_FEED":                  2,
		"DYNAMIC_EDUCATION":          3,
		"MERCHANT_CENTER_FEED":       4,
		"DYNAMIC_REAL_ESTATE":        5,
		"DYNAMIC_CUSTOM":             6,
		"DYNAMIC_HOTELS_AND_RENTALS": 7,
		"DYNAMIC_FLIGHTS":            8,
		"DYNAMIC_TRAVEL":             9,
		"DYNAMIC_LOCAL":              10,
		"DYNAMIC_JOBS":               11,
		"LOCATION_SYNC":              12,
		"BUSINESS_PROFILE_DYNAMIC_LOCATION_GROUP": 13,
		"CHAIN_DYNAMIC_LOCATION_GROUP":            14,
		"STATIC_LOCATION_GROUP":                   15,
		"HOTEL_PROPERTY":                          16,
	}
)

func (x AssetSetTypeEnum_AssetSetType) Enum() *AssetSetTypeEnum_AssetSetType {
	p := new(AssetSetTypeEnum_AssetSetType)
	*p = x
	return p
}

func (x AssetSetTypeEnum_AssetSetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetSetTypeEnum_AssetSetType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v14_enums_asset_set_type_proto_enumTypes[0].Descriptor()
}

func (AssetSetTypeEnum_AssetSetType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v14_enums_asset_set_type_proto_enumTypes[0]
}

func (x AssetSetTypeEnum_AssetSetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetSetTypeEnum_AssetSetType.Descriptor instead.
func (AssetSetTypeEnum_AssetSetType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible types of an asset set.
type AssetSetTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetSetTypeEnum) Reset() {
	*x = AssetSetTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_enums_asset_set_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSetTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetTypeEnum) ProtoMessage() {}

func (x *AssetSetTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_enums_asset_set_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetTypeEnum.ProtoReflect.Descriptor instead.
func (*AssetSetTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v14_enums_asset_set_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xad, 0x03, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x98, 0x03, 0x0a, 0x0c, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x47,
	0x45, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x59, 0x4e, 0x41,
	0x4d, 0x49, 0x43, 0x5f, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x5f, 0x43, 0x45, 0x4e, 0x54,
	0x45, 0x52, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x45, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49,
	0x43, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x4e,
	0x54, 0x41, 0x4c, 0x53, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49,
	0x43, 0x5f, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x53, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x10, 0x09, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x4a, 0x4f,
	0x42, 0x53, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x0c, 0x12, 0x2b, 0x0a, 0x27, 0x42, 0x55, 0x53, 0x49, 0x4e,
	0x45, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41,
	0x4d, 0x49, 0x43, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x0d, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x44, 0x59,
	0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43,
	0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10,
	0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45,
	0x52, 0x54, 0x59, 0x10, 0x10, 0x42, 0xeb, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x11, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescData = file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDesc
)

func file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDescData
}

var file_google_ads_googleads_v14_enums_asset_set_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v14_enums_asset_set_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v14_enums_asset_set_type_proto_goTypes = []interface{}{
	(AssetSetTypeEnum_AssetSetType)(0), // 0: google.ads.googleads.v14.enums.AssetSetTypeEnum.AssetSetType
	(*AssetSetTypeEnum)(nil),           // 1: google.ads.googleads.v14.enums.AssetSetTypeEnum
}
var file_google_ads_googleads_v14_enums_asset_set_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_enums_asset_set_type_proto_init() }
func file_google_ads_googleads_v14_enums_asset_set_type_proto_init() {
	if File_google_ads_googleads_v14_enums_asset_set_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_enums_asset_set_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSetTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_enums_asset_set_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_enums_asset_set_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v14_enums_asset_set_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v14_enums_asset_set_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_enums_asset_set_type_proto = out.File
	file_google_ads_googleads_v14_enums_asset_set_type_proto_rawDesc = nil
	file_google_ads_googleads_v14_enums_asset_set_type_proto_goTypes = nil
	file_google_ads_googleads_v14_enums_asset_set_type_proto_depIdxs = nil
}
