// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: partner/service/v1/partner_to_mica_service.proto

package servicev1

import (
	pingv1 "github.com/1080network/golang/partner/proto/common/pingv1"
	instrumentv1 "github.com/1080network/golang/partner/proto/partner/instrumentv1"
	organizationv1 "github.com/1080network/golang/partner/proto/partner/organizationv1"
	partnerv1 "github.com/1080network/golang/partner/proto/partner/partnerv1"
	paymenttokenv1 "github.com/1080network/golang/partner/proto/partner/paymenttokenv1"
	serviceproviderv1 "github.com/1080network/golang/partner/proto/partner/serviceproviderv1"
	storev1 "github.com/1080network/golang/partner/proto/partner/storev1"
	transactionv1 "github.com/1080network/golang/partner/proto/partner/transactionv1"
	valuev1 "github.com/1080network/golang/partner/proto/partner/valuev1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_partner_service_v1_partner_to_mica_service_proto protoreflect.FileDescriptor

var file_partner_service_v1_partner_to_mica_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f,
	0x6d, 0x69, 0x63, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xab, 0x19,
	0x0a, 0x14, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x39, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f,
	0x01, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x12, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xaf, 0x01, 0x0a, 0x22, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x42, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb8, 0x01, 0x0a,
	0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x32, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x0b, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x10,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x12, 0x2a, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b, 0x0a, 0x12, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x12, 0x30, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x34, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xa1, 0x01, 0x0a, 0x1e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb3, 0x01, 0x0a, 0x24, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5c, 0x0a, 0x1a, 0x69,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x54, 0x6f, 0x4d, 0x69, 0x63, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1c, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_partner_service_v1_partner_to_mica_service_proto_goTypes = []interface{}{
	(*partnerv1.GetPartnerRequest)(nil),                                  // 0: partner.partner.v1.GetPartnerRequest
	(*partnerv1.GetPartnerLegacyConfigurationRequest)(nil),               // 1: partner.partner.v1.GetPartnerLegacyConfigurationRequest
	(*partnerv1.UpdatePartnerLegacyConfigurationRequest)(nil),            // 2: partner.partner.v1.UpdatePartnerLegacyConfigurationRequest
	(*organizationv1.GetOrganizationRequest)(nil),                        // 3: partner.organization.v1.GetOrganizationRequest
	(*organizationv1.UpdateOrganizationRequest)(nil),                     // 4: partner.organization.v1.UpdateOrganizationRequest
	(*organizationv1.SearchOrganizationRequest)(nil),                     // 5: partner.organization.v1.SearchOrganizationRequest
	(*organizationv1.GetOrganizationLegacyConfigurationRequest)(nil),     // 6: partner.organization.v1.GetOrganizationLegacyConfigurationRequest
	(*organizationv1.UpdateOrganizationLegacyConfigurationRequest)(nil),  // 7: partner.organization.v1.UpdateOrganizationLegacyConfigurationRequest
	(*storev1.CreateStoreRequest)(nil),                                   // 8: partner.store.v1.CreateStoreRequest
	(*storev1.GetStoreRequest)(nil),                                      // 9: partner.store.v1.GetStoreRequest
	(*storev1.UpdateStoreRequest)(nil),                                   // 10: partner.store.v1.UpdateStoreRequest
	(*storev1.RemoveStoreRequest)(nil),                                   // 11: partner.store.v1.RemoveStoreRequest
	(*storev1.SearchStoreRequest)(nil),                                   // 12: partner.store.v1.SearchStoreRequest
	(*serviceproviderv1.SearchServiceProviderRequest)(nil),               // 13: partner.serviceprovider.v1.SearchServiceProviderRequest
	(*paymenttokenv1.RemovePaymentTokenRequest)(nil),                     // 14: partner.paymenttoken.v1.RemovePaymentTokenRequest
	(*paymenttokenv1.ExchangePaymentTokenRequest)(nil),                   // 15: partner.paymenttoken.v1.ExchangePaymentTokenRequest
	(*paymenttokenv1.SearchPaymentTokenRequest)(nil),                     // 16: partner.paymenttoken.v1.SearchPaymentTokenRequest
	(*valuev1.ObtainValueRequest)(nil),                                   // 17: partner.value.v1.ObtainValueRequest
	(*valuev1.ReturnValueRequest)(nil),                                   // 18: partner.value.v1.ReturnValueRequest
	(*valuev1.SearchValueRequest)(nil),                                   // 19: partner.value.v1.SearchValueRequest
	(*pingv1.PingRequest)(nil),                                           // 20: common.ping.v1.PingRequest
	(*instrumentv1.InitializeWidgetRequest)(nil),                         // 21: partner.widget.v1.InitializeWidgetRequest
	(*instrumentv1.ExchangeSessionKeyRequest)(nil),                       // 22: partner.instrument.v1.ExchangeSessionKeyRequest
	(*transactionv1.SearchTransactionDataRequest)(nil),                   // 23: partner.transaction.v1.SearchTransactionDataRequest
	(*transactionv1.SearchTransactionGeographyDataRequest)(nil),          // 24: partner.transaction.v1.SearchTransactionGeographyDataRequest
	(*transactionv1.SearchTransactionServiceProviderDataRequest)(nil),    // 25: partner.transaction.v1.SearchTransactionServiceProviderDataRequest
	(*partnerv1.GetPartnerResponse)(nil),                                 // 26: partner.partner.v1.GetPartnerResponse
	(*partnerv1.GetPartnerLegacyConfigurationResponse)(nil),              // 27: partner.partner.v1.GetPartnerLegacyConfigurationResponse
	(*partnerv1.UpdatePartnerLegacyConfigurationResponse)(nil),           // 28: partner.partner.v1.UpdatePartnerLegacyConfigurationResponse
	(*organizationv1.GetOrganizationResponse)(nil),                       // 29: partner.organization.v1.GetOrganizationResponse
	(*organizationv1.UpdateOrganizationResponse)(nil),                    // 30: partner.organization.v1.UpdateOrganizationResponse
	(*organizationv1.SearchOrganizationResponse)(nil),                    // 31: partner.organization.v1.SearchOrganizationResponse
	(*organizationv1.GetOrganizationLegacyConfigurationResponse)(nil),    // 32: partner.organization.v1.GetOrganizationLegacyConfigurationResponse
	(*organizationv1.UpdateOrganizationLegacyConfigurationResponse)(nil), // 33: partner.organization.v1.UpdateOrganizationLegacyConfigurationResponse
	(*storev1.CreateStoreResponse)(nil),                                  // 34: partner.store.v1.CreateStoreResponse
	(*storev1.GetStoreResponse)(nil),                                     // 35: partner.store.v1.GetStoreResponse
	(*storev1.UpdateStoreResponse)(nil),                                  // 36: partner.store.v1.UpdateStoreResponse
	(*storev1.RemoveStoreResponse)(nil),                                  // 37: partner.store.v1.RemoveStoreResponse
	(*storev1.SearchStoreResponse)(nil),                                  // 38: partner.store.v1.SearchStoreResponse
	(*serviceproviderv1.SearchServiceProviderResponse)(nil),              // 39: partner.serviceprovider.v1.SearchServiceProviderResponse
	(*paymenttokenv1.RemovePaymentTokenResponse)(nil),                    // 40: partner.paymenttoken.v1.RemovePaymentTokenResponse
	(*paymenttokenv1.ExchangePaymentTokenResponse)(nil),                  // 41: partner.paymenttoken.v1.ExchangePaymentTokenResponse
	(*paymenttokenv1.SearchPaymentTokenResponse)(nil),                    // 42: partner.paymenttoken.v1.SearchPaymentTokenResponse
	(*valuev1.ObtainValueResponse)(nil),                                  // 43: partner.value.v1.ObtainValueResponse
	(*valuev1.ReturnValueResponse)(nil),                                  // 44: partner.value.v1.ReturnValueResponse
	(*valuev1.SearchValueResponse)(nil),                                  // 45: partner.value.v1.SearchValueResponse
	(*pingv1.PingResponse)(nil),                                          // 46: common.ping.v1.PingResponse
	(*instrumentv1.InitializeWidgetResponse)(nil),                        // 47: partner.widget.v1.InitializeWidgetResponse
	(*instrumentv1.ExchangeSessionKeyResponse)(nil),                      // 48: partner.instrument.v1.ExchangeSessionKeyResponse
	(*transactionv1.SearchTransactionDataResponse)(nil),                  // 49: partner.transaction.v1.SearchTransactionDataResponse
	(*transactionv1.SearchTransactionGeographyDataResponse)(nil),         // 50: partner.transaction.v1.SearchTransactionGeographyDataResponse
	(*transactionv1.SearchTransactionServiceProviderDataResponse)(nil),   // 51: partner.transaction.v1.SearchTransactionServiceProviderDataResponse
}
var file_partner_service_v1_partner_to_mica_service_proto_depIdxs = []int32{
	0,  // 0: partner.service.v1.PartnerToMicaService.GetPartner:input_type -> partner.partner.v1.GetPartnerRequest
	1,  // 1: partner.service.v1.PartnerToMicaService.GetPartnerLegacyConfiguration:input_type -> partner.partner.v1.GetPartnerLegacyConfigurationRequest
	2,  // 2: partner.service.v1.PartnerToMicaService.UpdatePartnerLegacyConfiguration:input_type -> partner.partner.v1.UpdatePartnerLegacyConfigurationRequest
	3,  // 3: partner.service.v1.PartnerToMicaService.GetOrganization:input_type -> partner.organization.v1.GetOrganizationRequest
	4,  // 4: partner.service.v1.PartnerToMicaService.UpdateOrganization:input_type -> partner.organization.v1.UpdateOrganizationRequest
	5,  // 5: partner.service.v1.PartnerToMicaService.SearchOrganization:input_type -> partner.organization.v1.SearchOrganizationRequest
	6,  // 6: partner.service.v1.PartnerToMicaService.GetOrganizationLegacyConfiguration:input_type -> partner.organization.v1.GetOrganizationLegacyConfigurationRequest
	7,  // 7: partner.service.v1.PartnerToMicaService.UpdateOrganizationLegacyConfiguration:input_type -> partner.organization.v1.UpdateOrganizationLegacyConfigurationRequest
	8,  // 8: partner.service.v1.PartnerToMicaService.CreateStore:input_type -> partner.store.v1.CreateStoreRequest
	9,  // 9: partner.service.v1.PartnerToMicaService.GetStore:input_type -> partner.store.v1.GetStoreRequest
	10, // 10: partner.service.v1.PartnerToMicaService.UpdateStore:input_type -> partner.store.v1.UpdateStoreRequest
	11, // 11: partner.service.v1.PartnerToMicaService.RemoveStore:input_type -> partner.store.v1.RemoveStoreRequest
	12, // 12: partner.service.v1.PartnerToMicaService.SearchStore:input_type -> partner.store.v1.SearchStoreRequest
	13, // 13: partner.service.v1.PartnerToMicaService.SearchServiceProvider:input_type -> partner.serviceprovider.v1.SearchServiceProviderRequest
	14, // 14: partner.service.v1.PartnerToMicaService.ClosePaymentToken:input_type -> partner.paymenttoken.v1.RemovePaymentTokenRequest
	15, // 15: partner.service.v1.PartnerToMicaService.ReplacePaymentToken:input_type -> partner.paymenttoken.v1.ExchangePaymentTokenRequest
	16, // 16: partner.service.v1.PartnerToMicaService.SearchPaymentToken:input_type -> partner.paymenttoken.v1.SearchPaymentTokenRequest
	17, // 17: partner.service.v1.PartnerToMicaService.ObtainValue:input_type -> partner.value.v1.ObtainValueRequest
	18, // 18: partner.service.v1.PartnerToMicaService.ReturnValue:input_type -> partner.value.v1.ReturnValueRequest
	19, // 19: partner.service.v1.PartnerToMicaService.SearchValue:input_type -> partner.value.v1.SearchValueRequest
	20, // 20: partner.service.v1.PartnerToMicaService.Ping:input_type -> common.ping.v1.PingRequest
	21, // 21: partner.service.v1.PartnerToMicaService.InitializeWidget:input_type -> partner.widget.v1.InitializeWidgetRequest
	22, // 22: partner.service.v1.PartnerToMicaService.ExchangeSessionKey:input_type -> partner.instrument.v1.ExchangeSessionKeyRequest
	23, // 23: partner.service.v1.PartnerToMicaService.SearchTransactionData:input_type -> partner.transaction.v1.SearchTransactionDataRequest
	24, // 24: partner.service.v1.PartnerToMicaService.SearchTransactionGeographyData:input_type -> partner.transaction.v1.SearchTransactionGeographyDataRequest
	25, // 25: partner.service.v1.PartnerToMicaService.SearchTransactionServiceProviderData:input_type -> partner.transaction.v1.SearchTransactionServiceProviderDataRequest
	26, // 26: partner.service.v1.PartnerToMicaService.GetPartner:output_type -> partner.partner.v1.GetPartnerResponse
	27, // 27: partner.service.v1.PartnerToMicaService.GetPartnerLegacyConfiguration:output_type -> partner.partner.v1.GetPartnerLegacyConfigurationResponse
	28, // 28: partner.service.v1.PartnerToMicaService.UpdatePartnerLegacyConfiguration:output_type -> partner.partner.v1.UpdatePartnerLegacyConfigurationResponse
	29, // 29: partner.service.v1.PartnerToMicaService.GetOrganization:output_type -> partner.organization.v1.GetOrganizationResponse
	30, // 30: partner.service.v1.PartnerToMicaService.UpdateOrganization:output_type -> partner.organization.v1.UpdateOrganizationResponse
	31, // 31: partner.service.v1.PartnerToMicaService.SearchOrganization:output_type -> partner.organization.v1.SearchOrganizationResponse
	32, // 32: partner.service.v1.PartnerToMicaService.GetOrganizationLegacyConfiguration:output_type -> partner.organization.v1.GetOrganizationLegacyConfigurationResponse
	33, // 33: partner.service.v1.PartnerToMicaService.UpdateOrganizationLegacyConfiguration:output_type -> partner.organization.v1.UpdateOrganizationLegacyConfigurationResponse
	34, // 34: partner.service.v1.PartnerToMicaService.CreateStore:output_type -> partner.store.v1.CreateStoreResponse
	35, // 35: partner.service.v1.PartnerToMicaService.GetStore:output_type -> partner.store.v1.GetStoreResponse
	36, // 36: partner.service.v1.PartnerToMicaService.UpdateStore:output_type -> partner.store.v1.UpdateStoreResponse
	37, // 37: partner.service.v1.PartnerToMicaService.RemoveStore:output_type -> partner.store.v1.RemoveStoreResponse
	38, // 38: partner.service.v1.PartnerToMicaService.SearchStore:output_type -> partner.store.v1.SearchStoreResponse
	39, // 39: partner.service.v1.PartnerToMicaService.SearchServiceProvider:output_type -> partner.serviceprovider.v1.SearchServiceProviderResponse
	40, // 40: partner.service.v1.PartnerToMicaService.ClosePaymentToken:output_type -> partner.paymenttoken.v1.RemovePaymentTokenResponse
	41, // 41: partner.service.v1.PartnerToMicaService.ReplacePaymentToken:output_type -> partner.paymenttoken.v1.ExchangePaymentTokenResponse
	42, // 42: partner.service.v1.PartnerToMicaService.SearchPaymentToken:output_type -> partner.paymenttoken.v1.SearchPaymentTokenResponse
	43, // 43: partner.service.v1.PartnerToMicaService.ObtainValue:output_type -> partner.value.v1.ObtainValueResponse
	44, // 44: partner.service.v1.PartnerToMicaService.ReturnValue:output_type -> partner.value.v1.ReturnValueResponse
	45, // 45: partner.service.v1.PartnerToMicaService.SearchValue:output_type -> partner.value.v1.SearchValueResponse
	46, // 46: partner.service.v1.PartnerToMicaService.Ping:output_type -> common.ping.v1.PingResponse
	47, // 47: partner.service.v1.PartnerToMicaService.InitializeWidget:output_type -> partner.widget.v1.InitializeWidgetResponse
	48, // 48: partner.service.v1.PartnerToMicaService.ExchangeSessionKey:output_type -> partner.instrument.v1.ExchangeSessionKeyResponse
	49, // 49: partner.service.v1.PartnerToMicaService.SearchTransactionData:output_type -> partner.transaction.v1.SearchTransactionDataResponse
	50, // 50: partner.service.v1.PartnerToMicaService.SearchTransactionGeographyData:output_type -> partner.transaction.v1.SearchTransactionGeographyDataResponse
	51, // 51: partner.service.v1.PartnerToMicaService.SearchTransactionServiceProviderData:output_type -> partner.transaction.v1.SearchTransactionServiceProviderDataResponse
	26, // [26:52] is the sub-list for method output_type
	0,  // [0:26] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_partner_service_v1_partner_to_mica_service_proto_init() }
func file_partner_service_v1_partner_to_mica_service_proto_init() {
	if File_partner_service_v1_partner_to_mica_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_partner_service_v1_partner_to_mica_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_partner_service_v1_partner_to_mica_service_proto_goTypes,
		DependencyIndexes: file_partner_service_v1_partner_to_mica_service_proto_depIdxs,
	}.Build()
	File_partner_service_v1_partner_to_mica_service_proto = out.File
	file_partner_service_v1_partner_to_mica_service_proto_rawDesc = nil
	file_partner_service_v1_partner_to_mica_service_proto_goTypes = nil
	file_partner_service_v1_partner_to_mica_service_proto_depIdxs = nil
}
