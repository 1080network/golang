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
// source: mica/partner/administration/v1/admin_service.proto

package administrationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "github.com/1080network/golang/partner/proto/micashared/common/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mica_partner_administration_v1_admin_service_proto protoreflect.FileDescriptor

var file_mica_partner_administration_v1_admin_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc7, 0x0c, 0x0a, 0x1c, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x4d,
	0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x35, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x4d, 0x69,
	0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x4d, 0x69, 0x63,
	0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f,
	0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x93, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f,
	0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x89, 0x01, 0x0a, 0x1e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x65, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mica_partner_administration_v1_admin_service_proto_goTypes = []interface{}{
	(*v1.GenerateToMicaCertificateRequest)(nil),      // 0: micashared.common.v1.GenerateToMicaCertificateRequest
	(*v1.EnableToMicaCertificateRequest)(nil),        // 1: micashared.common.v1.EnableToMicaCertificateRequest
	(*v1.DisableToMicaCertificateRequest)(nil),       // 2: micashared.common.v1.DisableToMicaCertificateRequest
	(*v1.SearchToMicaCertificateRequest)(nil),        // 3: micashared.common.v1.SearchToMicaCertificateRequest
	(*v1.GetToMicaCertificateRequest)(nil),           // 4: micashared.common.v1.GetToMicaCertificateRequest
	(*v1.CreateSingleSignOnConsoleUserRequest)(nil),  // 5: micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	(*v1.UpdateSingleSignOnConsoleUserRequest)(nil),  // 6: micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	(*v1.SearchSingleSignOnConsoleUserRequest)(nil),  // 7: micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	(*v1.GetSingleSignOnConsoleUserRequest)(nil),     // 8: micashared.common.v1.GetSingleSignOnConsoleUserRequest
	(*v1.SearchDataExtractionRequest)(nil),           // 9: micashared.common.v1.SearchDataExtractionRequest
	(*v1.GetDataExtractionRequest)(nil),              // 10: micashared.common.v1.GetDataExtractionRequest
	(*v1.GenerateToMicaCertificateResponse)(nil),     // 11: micashared.common.v1.GenerateToMicaCertificateResponse
	(*v1.EnableToMicaCertificateResponse)(nil),       // 12: micashared.common.v1.EnableToMicaCertificateResponse
	(*v1.DisableToMicaCertificateResponse)(nil),      // 13: micashared.common.v1.DisableToMicaCertificateResponse
	(*v1.SearchToMicaCertificateResponse)(nil),       // 14: micashared.common.v1.SearchToMicaCertificateResponse
	(*v1.GetToMicaCertificateResponse)(nil),          // 15: micashared.common.v1.GetToMicaCertificateResponse
	(*v1.CreateSingleSignOnConsoleUserResponse)(nil), // 16: micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	(*v1.UpdateSingleSignOnConsoleUserResponse)(nil), // 17: micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	(*v1.SearchSingleSignOnConsoleUserResponse)(nil), // 18: micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	(*v1.GetSingleSignOnConsoleUserResponse)(nil),    // 19: micashared.common.v1.GetSingleSignOnConsoleUserResponse
	(*v1.SearchDataExtractionResponse)(nil),          // 20: micashared.common.v1.SearchDataExtractionResponse
	(*v1.GetDataExtractionResponse)(nil),             // 21: micashared.common.v1.GetDataExtractionResponse
}
var file_mica_partner_administration_v1_admin_service_proto_depIdxs = []int32{
	0,  // 0: mica.partner.administration.v1.PartnerAdministrationService.GenerateToMicaCertificate:input_type -> micashared.common.v1.GenerateToMicaCertificateRequest
	1,  // 1: mica.partner.administration.v1.PartnerAdministrationService.EnableToMicaCertificate:input_type -> micashared.common.v1.EnableToMicaCertificateRequest
	2,  // 2: mica.partner.administration.v1.PartnerAdministrationService.DisableToMicaCertificate:input_type -> micashared.common.v1.DisableToMicaCertificateRequest
	3,  // 3: mica.partner.administration.v1.PartnerAdministrationService.SearchToMicaCertificate:input_type -> micashared.common.v1.SearchToMicaCertificateRequest
	4,  // 4: mica.partner.administration.v1.PartnerAdministrationService.GetToMicaCertificate:input_type -> micashared.common.v1.GetToMicaCertificateRequest
	5,  // 5: mica.partner.administration.v1.PartnerAdministrationService.CreateSingleSignOnConsoleUser:input_type -> micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	6,  // 6: mica.partner.administration.v1.PartnerAdministrationService.UpdateSingleSignOnConsoleUser:input_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	7,  // 7: mica.partner.administration.v1.PartnerAdministrationService.SearchSingleSignOnUser:input_type -> micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	8,  // 8: mica.partner.administration.v1.PartnerAdministrationService.GetSingleSignOnConsoleUser:input_type -> micashared.common.v1.GetSingleSignOnConsoleUserRequest
	9,  // 9: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:input_type -> micashared.common.v1.SearchDataExtractionRequest
	10, // 10: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:input_type -> micashared.common.v1.GetDataExtractionRequest
	11, // 11: mica.partner.administration.v1.PartnerAdministrationService.GenerateToMicaCertificate:output_type -> micashared.common.v1.GenerateToMicaCertificateResponse
	12, // 12: mica.partner.administration.v1.PartnerAdministrationService.EnableToMicaCertificate:output_type -> micashared.common.v1.EnableToMicaCertificateResponse
	13, // 13: mica.partner.administration.v1.PartnerAdministrationService.DisableToMicaCertificate:output_type -> micashared.common.v1.DisableToMicaCertificateResponse
	14, // 14: mica.partner.administration.v1.PartnerAdministrationService.SearchToMicaCertificate:output_type -> micashared.common.v1.SearchToMicaCertificateResponse
	15, // 15: mica.partner.administration.v1.PartnerAdministrationService.GetToMicaCertificate:output_type -> micashared.common.v1.GetToMicaCertificateResponse
	16, // 16: mica.partner.administration.v1.PartnerAdministrationService.CreateSingleSignOnConsoleUser:output_type -> micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	17, // 17: mica.partner.administration.v1.PartnerAdministrationService.UpdateSingleSignOnConsoleUser:output_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	18, // 18: mica.partner.administration.v1.PartnerAdministrationService.SearchSingleSignOnUser:output_type -> micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	19, // 19: mica.partner.administration.v1.PartnerAdministrationService.GetSingleSignOnConsoleUser:output_type -> micashared.common.v1.GetSingleSignOnConsoleUserResponse
	20, // 20: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:output_type -> micashared.common.v1.SearchDataExtractionResponse
	21, // 21: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:output_type -> micashared.common.v1.GetDataExtractionResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mica_partner_administration_v1_admin_service_proto_init() }
func file_mica_partner_administration_v1_admin_service_proto_init() {
	if File_mica_partner_administration_v1_admin_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_partner_administration_v1_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_partner_administration_v1_admin_service_proto_goTypes,
		DependencyIndexes: file_mica_partner_administration_v1_admin_service_proto_depIdxs,
	}.Build()
	File_mica_partner_administration_v1_admin_service_proto = out.File
	file_mica_partner_administration_v1_admin_service_proto_rawDesc = nil
	file_mica_partner_administration_v1_admin_service_proto_goTypes = nil
	file_mica_partner_administration_v1_admin_service_proto_depIdxs = nil
}
