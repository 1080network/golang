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
// source: mica/serviceprovider/service/v1/service_provider_to_mica_service.proto

package servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	discountv1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/discountv1"
	instrumentv1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/instrumentv1"
	serviceproviderv1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/serviceproviderv1"
	userv1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/userv1"
	uuekv1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/uuekv1"
	valuev1 "github.com/1080network/golang/fullsdk/proto/mica/serviceprovider/valuev1"
	pingv1 "github.com/1080network/golang/fullsdk/proto/micashared/common/pingv1"
	v1 "github.com/1080network/golang/fullsdk/proto/micashared/common/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto protoreflect.FileDescriptor

var file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x6d, 0x69, 0x63, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x6d, 0x69, 0x63, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3e, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x75,
	0x75, 0x65, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x69,
	0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9c, 0x14, 0x0a, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x42,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x43, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x68, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x71, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x71, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x95, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x38, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa7, 0x01, 0x0a, 0x1c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x12, 0x41, 0x2e, 0x6d, 0x69,
	0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x9e, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55,
	0x45, 0x4b, 0x12, 0x3e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9e, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55,
	0x55, 0x45, 0x4b, 0x12, 0x3e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x75, 0x65, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x55, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb9, 0x01, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x70, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x26, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6e, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x21, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x69,
	0x63, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1e,
	0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_goTypes = []interface{}{
	(*serviceproviderv1.GetServiceProviderRequest)(nil),         // 0: mica.serviceprovider.serviceprovider.v1.GetServiceProviderRequest
	(*userv1.RegisterUserRequest)(nil),                          // 1: mica.serviceprovider.user.v1.RegisterUserRequest
	(*userv1.GetUserRequest)(nil),                               // 2: mica.serviceprovider.user.v1.GetUserRequest
	(*userv1.UpdateUserRequest)(nil),                            // 3: mica.serviceprovider.user.v1.UpdateUserRequest
	(*userv1.RemoveUserRequest)(nil),                            // 4: mica.serviceprovider.user.v1.RemoveUserRequest
	(*userv1.SearchUserRequest)(nil),                            // 5: mica.serviceprovider.user.v1.SearchUserRequest
	(*instrumentv1.RegisterInstrumentRequest)(nil),              // 6: mica.serviceprovider.instrument.v1.RegisterInstrumentRequest
	(*instrumentv1.GetInstrumentRequest)(nil),                   // 7: mica.serviceprovider.instrument.v1.GetInstrumentRequest
	(*instrumentv1.RemoveInstrumentRequest)(nil),                // 8: mica.serviceprovider.instrument.v1.RemoveInstrumentRequest
	(*instrumentv1.SearchInstrumentRequest)(nil),                // 9: mica.serviceprovider.instrument.v1.SearchInstrumentRequest
	(*uuekv1.ProvisionServiceProviderUUEKRequest)(nil),          // 10: mica.serviceprovider.uuek.v1.ProvisionServiceProviderUUEKRequest
	(*uuekv1.RemoveServiceProviderUUEKRequest)(nil),             // 11: mica.serviceprovider.uuek.v1.RemoveServiceProviderUUEKRequest
	(*uuekv1.SearchServiceProviderUUEKRequest)(nil),             // 12: mica.serviceprovider.uuek.v1.SearchServiceProviderUUEKRequest
	(*instrumentv1.ProvisionInstrumentLinkingCodeRequest)(nil),  // 13: mica.serviceprovider.instrument.v1.ProvisionInstrumentLinkingCodeRequest
	(*valuev1.SendValueRequest)(nil),                            // 14: mica.serviceprovider.value.v1.SendValueRequest
	(*valuev1.GetValueRequest)(nil),                             // 15: mica.serviceprovider.value.v1.GetValueRequest
	(*discountv1.SearchUserDiscountRequest)(nil),                // 16: mica.serviceprovider.discount.v1.SearchUserDiscountRequest
	(*v1.GetReceiptRequest)(nil),                                // 17: micashared.common.v1.GetReceiptRequest
	(*pingv1.PingRequest)(nil),                                  // 18: micashared.common.ping.v1.PingRequest
	(*serviceproviderv1.GetServiceProviderResponse)(nil),        // 19: mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse
	(*userv1.RegisterUserResponse)(nil),                         // 20: mica.serviceprovider.user.v1.RegisterUserResponse
	(*userv1.GetUserResponse)(nil),                              // 21: mica.serviceprovider.user.v1.GetUserResponse
	(*userv1.UpdateUserResponse)(nil),                           // 22: mica.serviceprovider.user.v1.UpdateUserResponse
	(*userv1.RemoveUserResponse)(nil),                           // 23: mica.serviceprovider.user.v1.RemoveUserResponse
	(*userv1.SearchUserResponse)(nil),                           // 24: mica.serviceprovider.user.v1.SearchUserResponse
	(*instrumentv1.RegisterInstrumentResponse)(nil),             // 25: mica.serviceprovider.instrument.v1.RegisterInstrumentResponse
	(*instrumentv1.GetInstrumentResponse)(nil),                  // 26: mica.serviceprovider.instrument.v1.GetInstrumentResponse
	(*instrumentv1.RemoveInstrumentResponse)(nil),               // 27: mica.serviceprovider.instrument.v1.RemoveInstrumentResponse
	(*instrumentv1.SearchInstrumentResponse)(nil),               // 28: mica.serviceprovider.instrument.v1.SearchInstrumentResponse
	(*uuekv1.ProvisionServiceProviderUUEKResponse)(nil),         // 29: mica.serviceprovider.uuek.v1.ProvisionServiceProviderUUEKResponse
	(*uuekv1.RemoveServiceProviderUUEKResponse)(nil),            // 30: mica.serviceprovider.uuek.v1.RemoveServiceProviderUUEKResponse
	(*uuekv1.SearchServiceProviderUUEKResponse)(nil),            // 31: mica.serviceprovider.uuek.v1.SearchServiceProviderUUEKResponse
	(*instrumentv1.ProvisionInstrumentLinkingCodeResponse)(nil), // 32: mica.serviceprovider.instrument.v1.ProvisionInstrumentLinkingCodeResponse
	(*valuev1.SendValueResponse)(nil),                           // 33: mica.serviceprovider.value.v1.SendValueResponse
	(*valuev1.GetValueResponse)(nil),                            // 34: mica.serviceprovider.value.v1.GetValueResponse
	(*discountv1.SearchUserDiscountResponse)(nil),               // 35: mica.serviceprovider.discount.v1.SearchUserDiscountResponse
	(*v1.GetReceiptResponse)(nil),                               // 36: micashared.common.v1.GetReceiptResponse
	(*pingv1.PingResponse)(nil),                                 // 37: micashared.common.ping.v1.PingResponse
}
var file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_depIdxs = []int32{
	0,  // 0: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetServiceProvider:input_type -> mica.serviceprovider.serviceprovider.v1.GetServiceProviderRequest
	1,  // 1: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RegisterUser:input_type -> mica.serviceprovider.user.v1.RegisterUserRequest
	2,  // 2: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetUser:input_type -> mica.serviceprovider.user.v1.GetUserRequest
	3,  // 3: mica.serviceprovider.service.v1.ServiceProviderToMicaService.UpdateUser:input_type -> mica.serviceprovider.user.v1.UpdateUserRequest
	4,  // 4: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveUser:input_type -> mica.serviceprovider.user.v1.RemoveUserRequest
	5,  // 5: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchUser:input_type -> mica.serviceprovider.user.v1.SearchUserRequest
	6,  // 6: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RegisterInstrument:input_type -> mica.serviceprovider.instrument.v1.RegisterInstrumentRequest
	7,  // 7: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetInstrument:input_type -> mica.serviceprovider.instrument.v1.GetInstrumentRequest
	8,  // 8: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveInstrument:input_type -> mica.serviceprovider.instrument.v1.RemoveInstrumentRequest
	9,  // 9: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchInstrument:input_type -> mica.serviceprovider.instrument.v1.SearchInstrumentRequest
	10, // 10: mica.serviceprovider.service.v1.ServiceProviderToMicaService.ProvisionServiceProviderUUEK:input_type -> mica.serviceprovider.uuek.v1.ProvisionServiceProviderUUEKRequest
	11, // 11: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveServiceProviderUUEK:input_type -> mica.serviceprovider.uuek.v1.RemoveServiceProviderUUEKRequest
	12, // 12: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchServiceProviderUUEK:input_type -> mica.serviceprovider.uuek.v1.SearchServiceProviderUUEKRequest
	13, // 13: mica.serviceprovider.service.v1.ServiceProviderToMicaService.ProvisionInstrumentLinkingCode:input_type -> mica.serviceprovider.instrument.v1.ProvisionInstrumentLinkingCodeRequest
	14, // 14: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SendValue:input_type -> mica.serviceprovider.value.v1.SendValueRequest
	15, // 15: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetValue:input_type -> mica.serviceprovider.value.v1.GetValueRequest
	16, // 16: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchUserDiscount:input_type -> mica.serviceprovider.discount.v1.SearchUserDiscountRequest
	17, // 17: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetReceipt:input_type -> micashared.common.v1.GetReceiptRequest
	18, // 18: mica.serviceprovider.service.v1.ServiceProviderToMicaService.Ping:input_type -> micashared.common.ping.v1.PingRequest
	19, // 19: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetServiceProvider:output_type -> mica.serviceprovider.serviceprovider.v1.GetServiceProviderResponse
	20, // 20: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RegisterUser:output_type -> mica.serviceprovider.user.v1.RegisterUserResponse
	21, // 21: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetUser:output_type -> mica.serviceprovider.user.v1.GetUserResponse
	22, // 22: mica.serviceprovider.service.v1.ServiceProviderToMicaService.UpdateUser:output_type -> mica.serviceprovider.user.v1.UpdateUserResponse
	23, // 23: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveUser:output_type -> mica.serviceprovider.user.v1.RemoveUserResponse
	24, // 24: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchUser:output_type -> mica.serviceprovider.user.v1.SearchUserResponse
	25, // 25: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RegisterInstrument:output_type -> mica.serviceprovider.instrument.v1.RegisterInstrumentResponse
	26, // 26: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetInstrument:output_type -> mica.serviceprovider.instrument.v1.GetInstrumentResponse
	27, // 27: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveInstrument:output_type -> mica.serviceprovider.instrument.v1.RemoveInstrumentResponse
	28, // 28: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchInstrument:output_type -> mica.serviceprovider.instrument.v1.SearchInstrumentResponse
	29, // 29: mica.serviceprovider.service.v1.ServiceProviderToMicaService.ProvisionServiceProviderUUEK:output_type -> mica.serviceprovider.uuek.v1.ProvisionServiceProviderUUEKResponse
	30, // 30: mica.serviceprovider.service.v1.ServiceProviderToMicaService.RemoveServiceProviderUUEK:output_type -> mica.serviceprovider.uuek.v1.RemoveServiceProviderUUEKResponse
	31, // 31: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchServiceProviderUUEK:output_type -> mica.serviceprovider.uuek.v1.SearchServiceProviderUUEKResponse
	32, // 32: mica.serviceprovider.service.v1.ServiceProviderToMicaService.ProvisionInstrumentLinkingCode:output_type -> mica.serviceprovider.instrument.v1.ProvisionInstrumentLinkingCodeResponse
	33, // 33: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SendValue:output_type -> mica.serviceprovider.value.v1.SendValueResponse
	34, // 34: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetValue:output_type -> mica.serviceprovider.value.v1.GetValueResponse
	35, // 35: mica.serviceprovider.service.v1.ServiceProviderToMicaService.SearchUserDiscount:output_type -> mica.serviceprovider.discount.v1.SearchUserDiscountResponse
	36, // 36: mica.serviceprovider.service.v1.ServiceProviderToMicaService.GetReceipt:output_type -> micashared.common.v1.GetReceiptResponse
	37, // 37: mica.serviceprovider.service.v1.ServiceProviderToMicaService.Ping:output_type -> micashared.common.ping.v1.PingResponse
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_init() }
func file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_init() {
	if File_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_goTypes,
		DependencyIndexes: file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_depIdxs,
	}.Build()
	File_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto = out.File
	file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_rawDesc = nil
	file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_goTypes = nil
	file_mica_serviceprovider_service_v1_service_provider_to_mica_service_proto_depIdxs = nil
}
