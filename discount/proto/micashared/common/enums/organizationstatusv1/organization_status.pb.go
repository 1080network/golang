// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.3
// source: micashared/common/enums/organizationstatus/v1/organization_status.proto

package organizationstatusv1

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

// State transition diagram:
//
//	INACTIVE (initial state) -> ACTIVE
//	ACTIVE -> INACTIVE
type OrganizationStatus int32

const (
	OrganizationStatus_ORGANIZATION_STATUS_UNSPECIFIED OrganizationStatus = 0
	OrganizationStatus_ORGANIZATION_STATUS_ACTIVE      OrganizationStatus = 3
	OrganizationStatus_ORGANIZATION_STATUS_INACTIVE    OrganizationStatus = 5
)

// Enum value maps for OrganizationStatus.
var (
	OrganizationStatus_name = map[int32]string{
		0: "ORGANIZATION_STATUS_UNSPECIFIED",
		3: "ORGANIZATION_STATUS_ACTIVE",
		5: "ORGANIZATION_STATUS_INACTIVE",
	}
	OrganizationStatus_value = map[string]int32{
		"ORGANIZATION_STATUS_UNSPECIFIED": 0,
		"ORGANIZATION_STATUS_ACTIVE":      3,
		"ORGANIZATION_STATUS_INACTIVE":    5,
	}
)

func (x OrganizationStatus) Enum() *OrganizationStatus {
	p := new(OrganizationStatus)
	*p = x
	return p
}

func (x OrganizationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_micashared_common_enums_organizationstatus_v1_organization_status_proto_enumTypes[0].Descriptor()
}

func (OrganizationStatus) Type() protoreflect.EnumType {
	return &file_micashared_common_enums_organizationstatus_v1_organization_status_proto_enumTypes[0]
}

func (x OrganizationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationStatus.Descriptor instead.
func (OrganizationStatus) EnumDescriptor() ([]byte, []int) {
	return file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescGZIP(), []int{0}
}

var File_micashared_common_enums_organizationstatus_v1_organization_status_proto protoreflect.FileDescriptor

var file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDesc = []byte{
	0x0a, 0x47, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x8d, 0x01, 0x0a, 0x12, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x1f, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x05, 0x22, 0x04, 0x08, 0x01, 0x10, 0x01, 0x22, 0x04, 0x08, 0x02,
	0x10, 0x02, 0x22, 0x04, 0x08, 0x04, 0x10, 0x04, 0x42, 0x71, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2c, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescOnce sync.Once
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescData = file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDesc
)

func file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescGZIP() []byte {
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescOnce.Do(func() {
		file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescData)
	})
	return file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDescData
}

var file_micashared_common_enums_organizationstatus_v1_organization_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_micashared_common_enums_organizationstatus_v1_organization_status_proto_goTypes = []interface{}{
	(OrganizationStatus)(0), // 0: micashared.common.enums.organizationstatus.v1.OrganizationStatus
}
var file_micashared_common_enums_organizationstatus_v1_organization_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_micashared_common_enums_organizationstatus_v1_organization_status_proto_init() }
func file_micashared_common_enums_organizationstatus_v1_organization_status_proto_init() {
	if File_micashared_common_enums_organizationstatus_v1_organization_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_enums_organizationstatus_v1_organization_status_proto_goTypes,
		DependencyIndexes: file_micashared_common_enums_organizationstatus_v1_organization_status_proto_depIdxs,
		EnumInfos:         file_micashared_common_enums_organizationstatus_v1_organization_status_proto_enumTypes,
	}.Build()
	File_micashared_common_enums_organizationstatus_v1_organization_status_proto = out.File
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_rawDesc = nil
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_goTypes = nil
	file_micashared_common_enums_organizationstatus_v1_organization_status_proto_depIdxs = nil
}
