// Copyright (c) 2022-2024 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: mica/partner/administration/v1/admin_service.proto

package administrationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pingv1 "github.com/1080network/golang/partner/proto/micashared/common/pingv1"
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
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2a, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf1, 0x13, 0x0a, 0x1c, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x32,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f,
	0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x93, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8e, 0x01,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac,
	0x01, 0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb2, 0x01,
	0x0a, 0x25, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x42, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x97, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5a, 0x69, 0x70, 0x12, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x69, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5a,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01, 0x0a,
	0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01, 0x0a, 0x23,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x1e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x31, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x50, 0x69, 0x6e,
	0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x26, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x60, 0x0a, 0x21,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mica_partner_administration_v1_admin_service_proto_goTypes = []interface{}{
	(*v1.GenerateMTLSCertificateRequest)(nil),                // 0: micashared.common.v1.GenerateMTLSCertificateRequest
	(*v1.UpdateMTLSCertificateRequest)(nil),                  // 1: micashared.common.v1.UpdateMTLSCertificateRequest
	(*v1.SearchMTLSCertificateRequest)(nil),                  // 2: micashared.common.v1.SearchMTLSCertificateRequest
	(*v1.GetMTLSCertificateRequest)(nil),                     // 3: micashared.common.v1.GetMTLSCertificateRequest
	(*v1.CreateSingleSignOnConsoleUserRequest)(nil),          // 4: micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	(*v1.UpdateSingleSignOnConsoleUserRequest)(nil),          // 5: micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	(*v1.SearchSingleSignOnConsoleUserRequest)(nil),          // 6: micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	(*v1.GetSingleSignOnConsoleUserRequest)(nil),             // 7: micashared.common.v1.GetSingleSignOnConsoleUserRequest
	(*v1.GetExternalClientSettingsRequest)(nil),              // 8: micashared.common.v1.GetExternalClientSettingsRequest
	(*v1.UpdateExternalClientCallBackAddressRequest)(nil),    // 9: micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	(*v1.GenerateExternalClientMTLSCertificateRequest)(nil),  // 10: micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	(*v1.GenerateQuickstartPackageZipRequest)(nil),           // 11: micashared.common.v1.GenerateQuickstartPackageZipRequest
	(*v1.UpdateExternalClientMTLSCertificateRequest)(nil),    // 12: micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	(*v1.SearchExternalClientMTLSCertificateRequest)(nil),    // 13: micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	(*v1.SearchDataExtractionRequest)(nil),                   // 14: micashared.common.v1.SearchDataExtractionRequest
	(*v1.GetDataExtractionRequest)(nil),                      // 15: micashared.common.v1.GetDataExtractionRequest
	(*pingv1.PingRequest)(nil),                               // 16: micashared.common.ping.v1.PingRequest
	(*v1.GenerateMTLSCertificateResponse)(nil),               // 17: micashared.common.v1.GenerateMTLSCertificateResponse
	(*v1.UpdateMTLSCertificateResponse)(nil),                 // 18: micashared.common.v1.UpdateMTLSCertificateResponse
	(*v1.SearchMTLSCertificateResponse)(nil),                 // 19: micashared.common.v1.SearchMTLSCertificateResponse
	(*v1.GetMTLSCertificateResponse)(nil),                    // 20: micashared.common.v1.GetMTLSCertificateResponse
	(*v1.CreateSingleSignOnConsoleUserResponse)(nil),         // 21: micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	(*v1.UpdateSingleSignOnConsoleUserResponse)(nil),         // 22: micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	(*v1.SearchSingleSignOnConsoleUserResponse)(nil),         // 23: micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	(*v1.GetSingleSignOnConsoleUserResponse)(nil),            // 24: micashared.common.v1.GetSingleSignOnConsoleUserResponse
	(*v1.GetExternalClientSettingsResponse)(nil),             // 25: micashared.common.v1.GetExternalClientSettingsResponse
	(*v1.UpdateExternalClientCallBackAddressResponse)(nil),   // 26: micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	(*v1.GenerateExternalClientMTLSCertificateResponse)(nil), // 27: micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	(*v1.GenerateQuickstartPackageZipResponse)(nil),          // 28: micashared.common.v1.GenerateQuickstartPackageZipResponse
	(*v1.UpdateExternalClientMTLSCertificateResponse)(nil),   // 29: micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	(*v1.SearchExternalClientMTLSCertificateResponse)(nil),   // 30: micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	(*v1.SearchDataExtractionResponse)(nil),                  // 31: micashared.common.v1.SearchDataExtractionResponse
	(*v1.GetDataExtractionResponse)(nil),                     // 32: micashared.common.v1.GetDataExtractionResponse
	(*pingv1.PingResponse)(nil),                              // 33: micashared.common.ping.v1.PingResponse
}
var file_mica_partner_administration_v1_admin_service_proto_depIdxs = []int32{
	0,  // 0: mica.partner.administration.v1.PartnerAdministrationService.GenerateMTLSCertificate:input_type -> micashared.common.v1.GenerateMTLSCertificateRequest
	1,  // 1: mica.partner.administration.v1.PartnerAdministrationService.UpdateMTLSCertificate:input_type -> micashared.common.v1.UpdateMTLSCertificateRequest
	2,  // 2: mica.partner.administration.v1.PartnerAdministrationService.SearchMTLSCertificate:input_type -> micashared.common.v1.SearchMTLSCertificateRequest
	3,  // 3: mica.partner.administration.v1.PartnerAdministrationService.GetMTLSCertificate:input_type -> micashared.common.v1.GetMTLSCertificateRequest
	4,  // 4: mica.partner.administration.v1.PartnerAdministrationService.CreateSingleSignOnConsoleUser:input_type -> micashared.common.v1.CreateSingleSignOnConsoleUserRequest
	5,  // 5: mica.partner.administration.v1.PartnerAdministrationService.UpdateSingleSignOnConsoleUser:input_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserRequest
	6,  // 6: mica.partner.administration.v1.PartnerAdministrationService.SearchSingleSignOnUser:input_type -> micashared.common.v1.SearchSingleSignOnConsoleUserRequest
	7,  // 7: mica.partner.administration.v1.PartnerAdministrationService.GetSingleSignOnConsoleUser:input_type -> micashared.common.v1.GetSingleSignOnConsoleUserRequest
	8,  // 8: mica.partner.administration.v1.PartnerAdministrationService.GetExternalClientSettings:input_type -> micashared.common.v1.GetExternalClientSettingsRequest
	9,  // 9: mica.partner.administration.v1.PartnerAdministrationService.UpdateExternalClientCallbackAddress:input_type -> micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	10, // 10: mica.partner.administration.v1.PartnerAdministrationService.GenerateExternalClientMTLSCertificate:input_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	11, // 11: mica.partner.administration.v1.PartnerAdministrationService.GenerateQuickstartPackageZip:input_type -> micashared.common.v1.GenerateQuickstartPackageZipRequest
	12, // 12: mica.partner.administration.v1.PartnerAdministrationService.UpdateExternalClientMTLSCertificate:input_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	13, // 13: mica.partner.administration.v1.PartnerAdministrationService.SearchExternalClientMTLSCertificate:input_type -> micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	14, // 14: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:input_type -> micashared.common.v1.SearchDataExtractionRequest
	15, // 15: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:input_type -> micashared.common.v1.GetDataExtractionRequest
	16, // 16: mica.partner.administration.v1.PartnerAdministrationService.PingExternal:input_type -> micashared.common.ping.v1.PingRequest
	17, // 17: mica.partner.administration.v1.PartnerAdministrationService.GenerateMTLSCertificate:output_type -> micashared.common.v1.GenerateMTLSCertificateResponse
	18, // 18: mica.partner.administration.v1.PartnerAdministrationService.UpdateMTLSCertificate:output_type -> micashared.common.v1.UpdateMTLSCertificateResponse
	19, // 19: mica.partner.administration.v1.PartnerAdministrationService.SearchMTLSCertificate:output_type -> micashared.common.v1.SearchMTLSCertificateResponse
	20, // 20: mica.partner.administration.v1.PartnerAdministrationService.GetMTLSCertificate:output_type -> micashared.common.v1.GetMTLSCertificateResponse
	21, // 21: mica.partner.administration.v1.PartnerAdministrationService.CreateSingleSignOnConsoleUser:output_type -> micashared.common.v1.CreateSingleSignOnConsoleUserResponse
	22, // 22: mica.partner.administration.v1.PartnerAdministrationService.UpdateSingleSignOnConsoleUser:output_type -> micashared.common.v1.UpdateSingleSignOnConsoleUserResponse
	23, // 23: mica.partner.administration.v1.PartnerAdministrationService.SearchSingleSignOnUser:output_type -> micashared.common.v1.SearchSingleSignOnConsoleUserResponse
	24, // 24: mica.partner.administration.v1.PartnerAdministrationService.GetSingleSignOnConsoleUser:output_type -> micashared.common.v1.GetSingleSignOnConsoleUserResponse
	25, // 25: mica.partner.administration.v1.PartnerAdministrationService.GetExternalClientSettings:output_type -> micashared.common.v1.GetExternalClientSettingsResponse
	26, // 26: mica.partner.administration.v1.PartnerAdministrationService.UpdateExternalClientCallbackAddress:output_type -> micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	27, // 27: mica.partner.administration.v1.PartnerAdministrationService.GenerateExternalClientMTLSCertificate:output_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	28, // 28: mica.partner.administration.v1.PartnerAdministrationService.GenerateQuickstartPackageZip:output_type -> micashared.common.v1.GenerateQuickstartPackageZipResponse
	29, // 29: mica.partner.administration.v1.PartnerAdministrationService.UpdateExternalClientMTLSCertificate:output_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	30, // 30: mica.partner.administration.v1.PartnerAdministrationService.SearchExternalClientMTLSCertificate:output_type -> micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	31, // 31: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:output_type -> micashared.common.v1.SearchDataExtractionResponse
	32, // 32: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:output_type -> micashared.common.v1.GetDataExtractionResponse
	33, // 33: mica.partner.administration.v1.PartnerAdministrationService.PingExternal:output_type -> micashared.common.ping.v1.PingResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
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
