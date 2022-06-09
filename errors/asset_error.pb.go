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
// source: google/ads/googleads/v11/errors/asset_error.proto

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

// Enum describing possible asset errors.
type AssetErrorEnum_AssetError int32

const (
	// Enum unspecified.
	AssetErrorEnum_UNSPECIFIED AssetErrorEnum_AssetError = 0
	// The received error code is not known in this version.
	AssetErrorEnum_UNKNOWN AssetErrorEnum_AssetError = 1
	// The customer is not is not on the allow-list for this asset type.
	AssetErrorEnum_CUSTOMER_NOT_ON_ALLOWLIST_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 13
	// Assets are duplicated across operations.
	AssetErrorEnum_DUPLICATE_ASSET AssetErrorEnum_AssetError = 3
	// The asset name is duplicated, either across operations or with an
	// existing asset.
	AssetErrorEnum_DUPLICATE_ASSET_NAME AssetErrorEnum_AssetError = 4
	// The Asset.asset_data oneof is empty.
	AssetErrorEnum_ASSET_DATA_IS_MISSING AssetErrorEnum_AssetError = 5
	// The asset has a name which is different from an existing duplicate that
	// represents the same content.
	AssetErrorEnum_CANNOT_MODIFY_ASSET_NAME AssetErrorEnum_AssetError = 6
	// The field cannot be set for this asset type.
	AssetErrorEnum_FIELD_INCOMPATIBLE_WITH_ASSET_TYPE AssetErrorEnum_AssetError = 7
	// Call to action must come from the list of supported values.
	AssetErrorEnum_INVALID_CALL_TO_ACTION_TEXT AssetErrorEnum_AssetError = 8
	// A lead form asset is created with an invalid combination of input fields.
	AssetErrorEnum_LEAD_FORM_INVALID_FIELDS_COMBINATION AssetErrorEnum_AssetError = 9
	// Lead forms require that the Terms of Service have been agreed to before
	// mutates can be executed.
	AssetErrorEnum_LEAD_FORM_MISSING_AGREEMENT AssetErrorEnum_AssetError = 10
	// Asset status is invalid in this operation.
	AssetErrorEnum_INVALID_ASSET_STATUS AssetErrorEnum_AssetError = 11
	// The field cannot be modified by this asset type.
	AssetErrorEnum_FIELD_CANNOT_BE_MODIFIED_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 12
	// Ad schedules for the same asset cannot overlap.
	AssetErrorEnum_SCHEDULES_CANNOT_OVERLAP AssetErrorEnum_AssetError = 14
	// Cannot set both percent off and money amount off fields of promotion
	// asset.
	AssetErrorEnum_PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF AssetErrorEnum_AssetError = 15
	// Cannot set both promotion code and orders over amount fields of promotion
	// asset.
	AssetErrorEnum_PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT AssetErrorEnum_AssetError = 16
	// The field has too many decimal places specified.
	AssetErrorEnum_TOO_MANY_DECIMAL_PLACES_SPECIFIED AssetErrorEnum_AssetError = 17
	// Duplicate assets across operations, which have identical Asset.asset_data
	// oneof, cannot have different asset level fields for asset types which are
	// deduped.
	AssetErrorEnum_DUPLICATE_ASSETS_WITH_DIFFERENT_FIELD_VALUE AssetErrorEnum_AssetError = 18
	// Carrier specific short number is not allowed.
	AssetErrorEnum_CALL_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED AssetErrorEnum_AssetError = 19
	// Customer consent required for call recording Terms of Service.
	AssetErrorEnum_CALL_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED AssetErrorEnum_AssetError = 20
	// The type of the specified phone number is not allowed.
	AssetErrorEnum_CALL_DISALLOWED_NUMBER_TYPE AssetErrorEnum_AssetError = 21
	// If the default call_conversion_action is not used, the customer must have
	// a ConversionAction with the same id and the ConversionAction must be call
	// conversion type.
	AssetErrorEnum_CALL_INVALID_CONVERSION_ACTION AssetErrorEnum_AssetError = 22
	// The country code of the phone number is invalid.
	AssetErrorEnum_CALL_INVALID_COUNTRY_CODE AssetErrorEnum_AssetError = 23
	// The format of the phone number is incorrect.
	AssetErrorEnum_CALL_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT AssetErrorEnum_AssetError = 24
	// The input phone number is not a valid phone number.
	AssetErrorEnum_CALL_INVALID_PHONE_NUMBER AssetErrorEnum_AssetError = 25
	// The phone number is not supported for country.
	AssetErrorEnum_CALL_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY AssetErrorEnum_AssetError = 26
	// Premium rate phone number is not allowed.
	AssetErrorEnum_CALL_PREMIUM_RATE_NUMBER_NOT_ALLOWED AssetErrorEnum_AssetError = 27
	// Vanity phone number is not allowed.
	AssetErrorEnum_CALL_VANITY_PHONE_NUMBER_NOT_ALLOWED AssetErrorEnum_AssetError = 28
	// PriceOffering cannot have the same value for header and description.
	AssetErrorEnum_PRICE_HEADER_SAME_AS_DESCRIPTION AssetErrorEnum_AssetError = 29
	// AppId is invalid.
	AssetErrorEnum_MOBILE_APP_INVALID_APP_ID AssetErrorEnum_AssetError = 30
	// Invalid App download URL in final URLs.
	AssetErrorEnum_MOBILE_APP_INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL AssetErrorEnum_AssetError = 31
	// Asset name is required for the asset type.
	AssetErrorEnum_NAME_REQUIRED_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 32
	// Legacy qualifying questions cannot be in the same Lead Form as
	// custom questions.
	AssetErrorEnum_LEAD_FORM_LEGACY_QUALIFYING_QUESTIONS_DISALLOWED AssetErrorEnum_AssetError = 33
	// Unique name is required for this asset type.
	AssetErrorEnum_NAME_CONFLICT_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 34
	// Cannot modify asset source.
	AssetErrorEnum_CANNOT_MODIFY_ASSET_SOURCE AssetErrorEnum_AssetError = 35
	// User can not modify the automatically created asset.
	AssetErrorEnum_CANNOT_MODIFY_AUTOMATICALLY_CREATED_ASSET AssetErrorEnum_AssetError = 36
)

// Enum value maps for AssetErrorEnum_AssetError.
var (
	AssetErrorEnum_AssetError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		13: "CUSTOMER_NOT_ON_ALLOWLIST_FOR_ASSET_TYPE",
		3:  "DUPLICATE_ASSET",
		4:  "DUPLICATE_ASSET_NAME",
		5:  "ASSET_DATA_IS_MISSING",
		6:  "CANNOT_MODIFY_ASSET_NAME",
		7:  "FIELD_INCOMPATIBLE_WITH_ASSET_TYPE",
		8:  "INVALID_CALL_TO_ACTION_TEXT",
		9:  "LEAD_FORM_INVALID_FIELDS_COMBINATION",
		10: "LEAD_FORM_MISSING_AGREEMENT",
		11: "INVALID_ASSET_STATUS",
		12: "FIELD_CANNOT_BE_MODIFIED_FOR_ASSET_TYPE",
		14: "SCHEDULES_CANNOT_OVERLAP",
		15: "PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF",
		16: "PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT",
		17: "TOO_MANY_DECIMAL_PLACES_SPECIFIED",
		18: "DUPLICATE_ASSETS_WITH_DIFFERENT_FIELD_VALUE",
		19: "CALL_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
		20: "CALL_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
		21: "CALL_DISALLOWED_NUMBER_TYPE",
		22: "CALL_INVALID_CONVERSION_ACTION",
		23: "CALL_INVALID_COUNTRY_CODE",
		24: "CALL_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT",
		25: "CALL_INVALID_PHONE_NUMBER",
		26: "CALL_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
		27: "CALL_PREMIUM_RATE_NUMBER_NOT_ALLOWED",
		28: "CALL_VANITY_PHONE_NUMBER_NOT_ALLOWED",
		29: "PRICE_HEADER_SAME_AS_DESCRIPTION",
		30: "MOBILE_APP_INVALID_APP_ID",
		31: "MOBILE_APP_INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL",
		32: "NAME_REQUIRED_FOR_ASSET_TYPE",
		33: "LEAD_FORM_LEGACY_QUALIFYING_QUESTIONS_DISALLOWED",
		34: "NAME_CONFLICT_FOR_ASSET_TYPE",
		35: "CANNOT_MODIFY_ASSET_SOURCE",
		36: "CANNOT_MODIFY_AUTOMATICALLY_CREATED_ASSET",
	}
	AssetErrorEnum_AssetError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"CUSTOMER_NOT_ON_ALLOWLIST_FOR_ASSET_TYPE":                   13,
		"DUPLICATE_ASSET":                                            3,
		"DUPLICATE_ASSET_NAME":                                       4,
		"ASSET_DATA_IS_MISSING":                                      5,
		"CANNOT_MODIFY_ASSET_NAME":                                   6,
		"FIELD_INCOMPATIBLE_WITH_ASSET_TYPE":                         7,
		"INVALID_CALL_TO_ACTION_TEXT":                                8,
		"LEAD_FORM_INVALID_FIELDS_COMBINATION":                       9,
		"LEAD_FORM_MISSING_AGREEMENT":                                10,
		"INVALID_ASSET_STATUS":                                       11,
		"FIELD_CANNOT_BE_MODIFIED_FOR_ASSET_TYPE":                    12,
		"SCHEDULES_CANNOT_OVERLAP":                                   14,
		"PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF":      15,
		"PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT": 16,
		"TOO_MANY_DECIMAL_PLACES_SPECIFIED":                          17,
		"DUPLICATE_ASSETS_WITH_DIFFERENT_FIELD_VALUE":                18,
		"CALL_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED":             19,
		"CALL_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED":          20,
		"CALL_DISALLOWED_NUMBER_TYPE":                                21,
		"CALL_INVALID_CONVERSION_ACTION":                             22,
		"CALL_INVALID_COUNTRY_CODE":                                  23,
		"CALL_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT":                  24,
		"CALL_INVALID_PHONE_NUMBER":                                  25,
		"CALL_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY":                26,
		"CALL_PREMIUM_RATE_NUMBER_NOT_ALLOWED":                       27,
		"CALL_VANITY_PHONE_NUMBER_NOT_ALLOWED":                       28,
		"PRICE_HEADER_SAME_AS_DESCRIPTION":                           29,
		"MOBILE_APP_INVALID_APP_ID":                                  30,
		"MOBILE_APP_INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL":          31,
		"NAME_REQUIRED_FOR_ASSET_TYPE":                               32,
		"LEAD_FORM_LEGACY_QUALIFYING_QUESTIONS_DISALLOWED":           33,
		"NAME_CONFLICT_FOR_ASSET_TYPE":                               34,
		"CANNOT_MODIFY_ASSET_SOURCE":                                 35,
		"CANNOT_MODIFY_AUTOMATICALLY_CREATED_ASSET":                  36,
	}
)

func (x AssetErrorEnum_AssetError) Enum() *AssetErrorEnum_AssetError {
	p := new(AssetErrorEnum_AssetError)
	*p = x
	return p
}

func (x AssetErrorEnum_AssetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetErrorEnum_AssetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v11_errors_asset_error_proto_enumTypes[0].Descriptor()
}

func (AssetErrorEnum_AssetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v11_errors_asset_error_proto_enumTypes[0]
}

func (x AssetErrorEnum_AssetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetErrorEnum_AssetError.Descriptor instead.
func (AssetErrorEnum_AssetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_errors_asset_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset errors.
type AssetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetErrorEnum) Reset() {
	*x = AssetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_errors_asset_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetErrorEnum) ProtoMessage() {}

func (x *AssetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_errors_asset_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_errors_asset_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v11_errors_asset_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_errors_asset_error_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x86, 0x0b, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xf3, 0x0a, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x55, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x04, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x49, 0x53, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x26, 0x0a, 0x22, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x58,
	0x54, 0x10, 0x08, 0x12, 0x28, 0x0a, 0x24, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x5f,
	0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x1f, 0x0a,
	0x1b, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x18,
	0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0b, 0x12, 0x2b, 0x0a, 0x27, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41,
	0x50, 0x10, 0x0e, 0x12, 0x39, 0x0a, 0x35, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x43,
	0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x46, 0x46, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x4e, 0x45,
	0x59, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x0f, 0x12, 0x3e,
	0x0a, 0x3a, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x53,
	0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x10, 0x12, 0x25,
	0x0a, 0x21, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x44, 0x45, 0x43, 0x49, 0x4d,
	0x41, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x53, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x11, 0x12, 0x2f, 0x0a, 0x2b, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x44,
	0x49, 0x46, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x12, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43,
	0x41, 0x52, 0x52, 0x49, 0x45, 0x52, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f,
	0x53, 0x48, 0x4f, 0x52, 0x54, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x13, 0x12, 0x35, 0x0a, 0x31, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x53,
	0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x14, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x15, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x16, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x10, 0x17, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x4f, 0x4d, 0x45, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x50,
	0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x10, 0x18, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x19, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x48, 0x4f, 0x4e,
	0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x52, 0x59, 0x10, 0x1a, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x45,
	0x4d, 0x49, 0x55, 0x4d, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x1b, 0x12, 0x28,
	0x0a, 0x24, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x56, 0x41, 0x4e, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x48,
	0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x1c, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x52, 0x49, 0x43,
	0x45, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x53,
	0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1d, 0x12, 0x1d,
	0x0a, 0x19, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x10, 0x1e, 0x12, 0x35, 0x0a,
	0x31, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x55,
	0x52, 0x4c, 0x10, 0x1f, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x20, 0x12, 0x34, 0x0a, 0x30, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49,
	0x46, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x44, 0x49, 0x53, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x21, 0x12, 0x20, 0x0a, 0x1c,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x22, 0x12, 0x1e,
	0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x23, 0x12, 0x2d,
	0x0a, 0x29, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f,
	0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x24, 0x42, 0xef, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x31, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_errors_asset_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_errors_asset_error_proto_rawDescData = file_google_ads_googleads_v11_errors_asset_error_proto_rawDesc
)

func file_google_ads_googleads_v11_errors_asset_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_errors_asset_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_errors_asset_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_errors_asset_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_errors_asset_error_proto_rawDescData
}

var file_google_ads_googleads_v11_errors_asset_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v11_errors_asset_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_errors_asset_error_proto_goTypes = []interface{}{
	(AssetErrorEnum_AssetError)(0), // 0: google.ads.googleads.v11.errors.AssetErrorEnum.AssetError
	(*AssetErrorEnum)(nil),         // 1: google.ads.googleads.v11.errors.AssetErrorEnum
}
var file_google_ads_googleads_v11_errors_asset_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_errors_asset_error_proto_init() }
func file_google_ads_googleads_v11_errors_asset_error_proto_init() {
	if File_google_ads_googleads_v11_errors_asset_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_errors_asset_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v11_errors_asset_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_errors_asset_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_errors_asset_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v11_errors_asset_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v11_errors_asset_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_errors_asset_error_proto = out.File
	file_google_ads_googleads_v11_errors_asset_error_proto_rawDesc = nil
	file_google_ads_googleads_v11_errors_asset_error_proto_goTypes = nil
	file_google_ads_googleads_v11_errors_asset_error_proto_depIdxs = nil
}
