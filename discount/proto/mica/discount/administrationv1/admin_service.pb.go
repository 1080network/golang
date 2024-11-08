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
// 	protoc        v4.24.4
// source: mica/discount/administration/v1/admin_service.proto

package administrationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pingv1 "github.com/1080network/golang/discount/proto/micashared/common/pingv1"
	v1 "github.com/1080network/golang/discount/proto/micashared/common/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mica_discount_administration_v1_admin_service_proto protoreflect.FileDescriptor

var file_mica_discount_administration_v1_admin_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x0f, 0x0a, 0x1d, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x54,
	0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x54,
	0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01, 0x0a, 0x23, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb2, 0x01, 0x0a, 0x25, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x42, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01,
	0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xac, 0x01, 0x0a,
	0x23, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x40, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x94, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x35, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xa2, 0x01, 0x0a, 0x29, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x26, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x67, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x1a,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1e, 0x6d, 0x69, 0x63, 0x61,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43,
	0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mica_discount_administration_v1_admin_service_proto_goTypes = []interface{}{
	(*v1.GenerateMTLSCertificateRequest)(nil),                // 0: micashared.common.v1.GenerateMTLSCertificateRequest
	(*v1.UpdateMTLSCertificateRequest)(nil),                  // 1: micashared.common.v1.UpdateMTLSCertificateRequest
	(*v1.SearchMTLSCertificateRequest)(nil),                  // 2: micashared.common.v1.SearchMTLSCertificateRequest
	(*v1.GetMTLSCertificateRequest)(nil),                     // 3: micashared.common.v1.GetMTLSCertificateRequest
	(*v1.GetExternalClientSettingsRequest)(nil),              // 4: micashared.common.v1.GetExternalClientSettingsRequest
	(*v1.UpdateExternalClientCallBackAddressRequest)(nil),    // 5: micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	(*v1.GenerateExternalClientMTLSCertificateRequest)(nil),  // 6: micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	(*v1.UpdateExternalClientMTLSCertificateRequest)(nil),    // 7: micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	(*v1.SearchExternalClientMTLSCertificateRequest)(nil),    // 8: micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	(*v1.CreateApiTokenConfigurationRequest)(nil),            // 9: micashared.common.v1.CreateApiTokenConfigurationRequest
	(*v1.GetApiTokenConfigurationRequest)(nil),               // 10: micashared.common.v1.GetApiTokenConfigurationRequest
	(*v1.UpdateApiTokenConfigurationRequest)(nil),            // 11: micashared.common.v1.UpdateApiTokenConfigurationRequest
	(*pingv1.PingRequest)(nil),                               // 12: micashared.common.ping.v1.PingRequest
	(*v1.GenerateMTLSCertificateResponse)(nil),               // 13: micashared.common.v1.GenerateMTLSCertificateResponse
	(*v1.UpdateMTLSCertificateResponse)(nil),                 // 14: micashared.common.v1.UpdateMTLSCertificateResponse
	(*v1.SearchMTLSCertificateResponse)(nil),                 // 15: micashared.common.v1.SearchMTLSCertificateResponse
	(*v1.GetMTLSCertificateResponse)(nil),                    // 16: micashared.common.v1.GetMTLSCertificateResponse
	(*v1.GetExternalClientSettingsResponse)(nil),             // 17: micashared.common.v1.GetExternalClientSettingsResponse
	(*v1.UpdateExternalClientCallBackAddressResponse)(nil),   // 18: micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	(*v1.GenerateExternalClientMTLSCertificateResponse)(nil), // 19: micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	(*v1.UpdateExternalClientMTLSCertificateResponse)(nil),   // 20: micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	(*v1.SearchExternalClientMTLSCertificateResponse)(nil),   // 21: micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	(*v1.CreateApiTokenConfigurationResponse)(nil),           // 22: micashared.common.v1.CreateApiTokenConfigurationResponse
	(*v1.GetApiTokenConfigurationResponse)(nil),              // 23: micashared.common.v1.GetApiTokenConfigurationResponse
	(*v1.UpdateApiTokenConfigurationResponse)(nil),           // 24: micashared.common.v1.UpdateApiTokenConfigurationResponse
	(*pingv1.PingResponse)(nil),                              // 25: micashared.common.ping.v1.PingResponse
}
var file_mica_discount_administration_v1_admin_service_proto_depIdxs = []int32{
	0,  // 0: mica.discount.administration.v1.DiscountAdministrationService.GenerateMTLSCertificate:input_type -> micashared.common.v1.GenerateMTLSCertificateRequest
	1,  // 1: mica.discount.administration.v1.DiscountAdministrationService.UpdateMTLSCertificate:input_type -> micashared.common.v1.UpdateMTLSCertificateRequest
	2,  // 2: mica.discount.administration.v1.DiscountAdministrationService.SearchMTLSCertificate:input_type -> micashared.common.v1.SearchMTLSCertificateRequest
	3,  // 3: mica.discount.administration.v1.DiscountAdministrationService.GetMTLSCertificate:input_type -> micashared.common.v1.GetMTLSCertificateRequest
	4,  // 4: mica.discount.administration.v1.DiscountAdministrationService.GetExternalClientSettings:input_type -> micashared.common.v1.GetExternalClientSettingsRequest
	5,  // 5: mica.discount.administration.v1.DiscountAdministrationService.UpdateExternalClientCallbackAddress:input_type -> micashared.common.v1.UpdateExternalClientCallBackAddressRequest
	6,  // 6: mica.discount.administration.v1.DiscountAdministrationService.GenerateExternalClientMTLSCertificate:input_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateRequest
	7,  // 7: mica.discount.administration.v1.DiscountAdministrationService.UpdateExternalClientMTLSCertificate:input_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateRequest
	8,  // 8: mica.discount.administration.v1.DiscountAdministrationService.SearchExternalClientMTLSCertificate:input_type -> micashared.common.v1.SearchExternalClientMTLSCertificateRequest
	9,  // 9: mica.discount.administration.v1.DiscountAdministrationService.CreateApiTokenConfiguration:input_type -> micashared.common.v1.CreateApiTokenConfigurationRequest
	10, // 10: mica.discount.administration.v1.DiscountAdministrationService.GetApiTokenConfiguration:input_type -> micashared.common.v1.GetApiTokenConfigurationRequest
	11, // 11: mica.discount.administration.v1.DiscountAdministrationService.UpdateAPITokenAuthenticationConfiguration:input_type -> micashared.common.v1.UpdateApiTokenConfigurationRequest
	12, // 12: mica.discount.administration.v1.DiscountAdministrationService.PingExternal:input_type -> micashared.common.ping.v1.PingRequest
	13, // 13: mica.discount.administration.v1.DiscountAdministrationService.GenerateMTLSCertificate:output_type -> micashared.common.v1.GenerateMTLSCertificateResponse
	14, // 14: mica.discount.administration.v1.DiscountAdministrationService.UpdateMTLSCertificate:output_type -> micashared.common.v1.UpdateMTLSCertificateResponse
	15, // 15: mica.discount.administration.v1.DiscountAdministrationService.SearchMTLSCertificate:output_type -> micashared.common.v1.SearchMTLSCertificateResponse
	16, // 16: mica.discount.administration.v1.DiscountAdministrationService.GetMTLSCertificate:output_type -> micashared.common.v1.GetMTLSCertificateResponse
	17, // 17: mica.discount.administration.v1.DiscountAdministrationService.GetExternalClientSettings:output_type -> micashared.common.v1.GetExternalClientSettingsResponse
	18, // 18: mica.discount.administration.v1.DiscountAdministrationService.UpdateExternalClientCallbackAddress:output_type -> micashared.common.v1.UpdateExternalClientCallBackAddressResponse
	19, // 19: mica.discount.administration.v1.DiscountAdministrationService.GenerateExternalClientMTLSCertificate:output_type -> micashared.common.v1.GenerateExternalClientMTLSCertificateResponse
	20, // 20: mica.discount.administration.v1.DiscountAdministrationService.UpdateExternalClientMTLSCertificate:output_type -> micashared.common.v1.UpdateExternalClientMTLSCertificateResponse
	21, // 21: mica.discount.administration.v1.DiscountAdministrationService.SearchExternalClientMTLSCertificate:output_type -> micashared.common.v1.SearchExternalClientMTLSCertificateResponse
	22, // 22: mica.discount.administration.v1.DiscountAdministrationService.CreateApiTokenConfiguration:output_type -> micashared.common.v1.CreateApiTokenConfigurationResponse
	23, // 23: mica.discount.administration.v1.DiscountAdministrationService.GetApiTokenConfiguration:output_type -> micashared.common.v1.GetApiTokenConfigurationResponse
	24, // 24: mica.discount.administration.v1.DiscountAdministrationService.UpdateAPITokenAuthenticationConfiguration:output_type -> micashared.common.v1.UpdateApiTokenConfigurationResponse
	25, // 25: mica.discount.administration.v1.DiscountAdministrationService.PingExternal:output_type -> micashared.common.ping.v1.PingResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mica_discount_administration_v1_admin_service_proto_init() }
func file_mica_discount_administration_v1_admin_service_proto_init() {
	if File_mica_discount_administration_v1_admin_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_discount_administration_v1_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_discount_administration_v1_admin_service_proto_goTypes,
		DependencyIndexes: file_mica_discount_administration_v1_admin_service_proto_depIdxs,
	}.Build()
	File_mica_discount_administration_v1_admin_service_proto = out.File
	file_mica_discount_administration_v1_admin_service_proto_rawDesc = nil
	file_mica_discount_administration_v1_admin_service_proto_goTypes = nil
	file_mica_discount_administration_v1_admin_service_proto_depIdxs = nil
}
