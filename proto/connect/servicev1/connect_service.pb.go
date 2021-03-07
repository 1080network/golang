// Copyright (c) 2021 1080 Network, Inc. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of 1080 Network, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "1080 Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the 1080 Software without a
// valid license or the prior written approval of 1080 Network, Inc. 1080 Network, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of 1080 Network, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: connect/service/v1/connect_service.proto

package servicev1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pingv1 "ten80/proto/common/pingv1"
	instrumentv1 "ten80/proto/connect/instrumentv1"
	serviceproviderv1 "ten80/proto/connect/serviceproviderv1"
	staticdatav1 "ten80/proto/connect/staticdatav1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_connect_service_v1_connect_service_proto protoreflect.FileDescriptor

var file_connect_service_v1_connect_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x26,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8f, 0x08,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8e,
	0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0xa5, 0x01, 0x0a, 0x20, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xae, 0x01, 0x0a, 0x23, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x41, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x20, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3e, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xb7, 0x01, 0x0a, 0x26, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x69, 0x62, 0x62, 0x69,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x69, 0x62, 0x62, 0x69, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x45, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x52, 0x69, 0x62, 0x62, 0x69, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1f, 0x5a, 0x1d, 0x74, 0x65, 0x6e, 0x38, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_connect_service_v1_connect_service_proto_goTypes = []interface{}{
	(*staticdatav1.GetStaticDataRequest)(nil),                           // 0: connect.staticdata.v1.GetStaticDataRequest
	(*serviceproviderv1.SearchServiceProviderRequest)(nil),              // 1: connect.serviceprovider.v1.SearchServiceProviderRequest
	(*instrumentv1.WidgetRegisterInstrumentInitiateRequest)(nil),        // 2: connect.instrument.v1.WidgetRegisterInstrumentInitiateRequest
	(*instrumentv1.WidgetRegisterInstrumentWithAccountRequest)(nil),     // 3: connect.instrument.v1.WidgetRegisterInstrumentWithAccountRequest
	(*instrumentv1.WidgetRegisterInstrumentCompleteRequest)(nil),        // 4: connect.instrument.v1.WidgetRegisterInstrumentCompleteRequest
	(*instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest)(nil),  // 5: connect.instrument.v1.WidgetRibbitRegisterInstrumentInitiateRequest
	(*pingv1.PingRequest)(nil),                                          // 6: common.ping.v1.PingRequest
	(*staticdatav1.GetStaticDataResponse)(nil),                          // 7: connect.staticdata.v1.GetStaticDataResponse
	(*serviceproviderv1.SearchServiceProviderResponse)(nil),             // 8: connect.serviceprovider.v1.SearchServiceProviderResponse
	(*instrumentv1.WidgetRegisterInstrumentInitiateResponse)(nil),       // 9: connect.instrument.v1.WidgetRegisterInstrumentInitiateResponse
	(*instrumentv1.WidgetRegisterInstrumentWithAccountResponse)(nil),    // 10: connect.instrument.v1.WidgetRegisterInstrumentWithAccountResponse
	(*instrumentv1.WidgetRegisterInstrumentCompleteResponse)(nil),       // 11: connect.instrument.v1.WidgetRegisterInstrumentCompleteResponse
	(*instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse)(nil), // 12: connect.instrument.v1.WidgetRibbitRegisterInstrumentInitiateResponse
	(*pingv1.PingResponse)(nil),                                         // 13: common.ping.v1.PingResponse
}
var file_connect_service_v1_connect_service_proto_depIdxs = []int32{
	0,  // 0: connect.service.v1.ConnectService.GetStaticData:input_type -> connect.staticdata.v1.GetStaticDataRequest
	1,  // 1: connect.service.v1.ConnectService.SearchServiceProvider:input_type -> connect.serviceprovider.v1.SearchServiceProviderRequest
	2,  // 2: connect.service.v1.ConnectService.WidgetRegisterInstrumentInitiate:input_type -> connect.instrument.v1.WidgetRegisterInstrumentInitiateRequest
	3,  // 3: connect.service.v1.ConnectService.WidgetRegisterInstrumentWithAccount:input_type -> connect.instrument.v1.WidgetRegisterInstrumentWithAccountRequest
	4,  // 4: connect.service.v1.ConnectService.WidgetRegisterInstrumentComplete:input_type -> connect.instrument.v1.WidgetRegisterInstrumentCompleteRequest
	5,  // 5: connect.service.v1.ConnectService.WidgetRibbitRegisterInstrumentInitiate:input_type -> connect.instrument.v1.WidgetRibbitRegisterInstrumentInitiateRequest
	6,  // 6: connect.service.v1.ConnectService.Ping:input_type -> common.ping.v1.PingRequest
	7,  // 7: connect.service.v1.ConnectService.GetStaticData:output_type -> connect.staticdata.v1.GetStaticDataResponse
	8,  // 8: connect.service.v1.ConnectService.SearchServiceProvider:output_type -> connect.serviceprovider.v1.SearchServiceProviderResponse
	9,  // 9: connect.service.v1.ConnectService.WidgetRegisterInstrumentInitiate:output_type -> connect.instrument.v1.WidgetRegisterInstrumentInitiateResponse
	10, // 10: connect.service.v1.ConnectService.WidgetRegisterInstrumentWithAccount:output_type -> connect.instrument.v1.WidgetRegisterInstrumentWithAccountResponse
	11, // 11: connect.service.v1.ConnectService.WidgetRegisterInstrumentComplete:output_type -> connect.instrument.v1.WidgetRegisterInstrumentCompleteResponse
	12, // 12: connect.service.v1.ConnectService.WidgetRibbitRegisterInstrumentInitiate:output_type -> connect.instrument.v1.WidgetRibbitRegisterInstrumentInitiateResponse
	13, // 13: connect.service.v1.ConnectService.Ping:output_type -> common.ping.v1.PingResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_connect_service_v1_connect_service_proto_init() }
func file_connect_service_v1_connect_service_proto_init() {
	if File_connect_service_v1_connect_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connect_service_v1_connect_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_service_v1_connect_service_proto_goTypes,
		DependencyIndexes: file_connect_service_v1_connect_service_proto_depIdxs,
	}.Build()
	File_connect_service_v1_connect_service_proto = out.File
	file_connect_service_v1_connect_service_proto_rawDesc = nil
	file_connect_service_v1_connect_service_proto_goTypes = nil
	file_connect_service_v1_connect_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectServiceClient is the client API for ConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectServiceClient interface {
	GetStaticData(ctx context.Context, in *staticdatav1.GetStaticDataRequest, opts ...grpc.CallOption) (*staticdatav1.GetStaticDataResponse, error)
	SearchServiceProvider(ctx context.Context, in *serviceproviderv1.SearchServiceProviderRequest, opts ...grpc.CallOption) (*serviceproviderv1.SearchServiceProviderResponse, error)
	WidgetRegisterInstrumentInitiate(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentInitiateRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentInitiateResponse, error)
	// optional call required when user uses username/password and they have more than one account at the service
	// provider
	WidgetRegisterInstrumentWithAccount(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentWithAccountRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentWithAccountResponse, error)
	WidgetRegisterInstrumentComplete(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentCompleteRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentCompleteResponse, error)
	WidgetRibbitRegisterInstrumentInitiate(ctx context.Context, in *instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse, error)
	Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error)
}

type connectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectServiceClient(cc grpc.ClientConnInterface) ConnectServiceClient {
	return &connectServiceClient{cc}
}

func (c *connectServiceClient) GetStaticData(ctx context.Context, in *staticdatav1.GetStaticDataRequest, opts ...grpc.CallOption) (*staticdatav1.GetStaticDataResponse, error) {
	out := new(staticdatav1.GetStaticDataResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/GetStaticData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) SearchServiceProvider(ctx context.Context, in *serviceproviderv1.SearchServiceProviderRequest, opts ...grpc.CallOption) (*serviceproviderv1.SearchServiceProviderResponse, error) {
	out := new(serviceproviderv1.SearchServiceProviderResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/SearchServiceProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) WidgetRegisterInstrumentInitiate(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentInitiateRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentInitiateResponse, error) {
	out := new(instrumentv1.WidgetRegisterInstrumentInitiateResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/WidgetRegisterInstrumentInitiate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) WidgetRegisterInstrumentWithAccount(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentWithAccountRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentWithAccountResponse, error) {
	out := new(instrumentv1.WidgetRegisterInstrumentWithAccountResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/WidgetRegisterInstrumentWithAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) WidgetRegisterInstrumentComplete(ctx context.Context, in *instrumentv1.WidgetRegisterInstrumentCompleteRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRegisterInstrumentCompleteResponse, error) {
	out := new(instrumentv1.WidgetRegisterInstrumentCompleteResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/WidgetRegisterInstrumentComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) WidgetRibbitRegisterInstrumentInitiate(ctx context.Context, in *instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest, opts ...grpc.CallOption) (*instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse, error) {
	out := new(instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/WidgetRibbitRegisterInstrumentInitiate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) Ping(ctx context.Context, in *pingv1.PingRequest, opts ...grpc.CallOption) (*pingv1.PingResponse, error) {
	out := new(pingv1.PingResponse)
	err := c.cc.Invoke(ctx, "/connect.service.v1.ConnectService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServiceServer is the server API for ConnectService service.
type ConnectServiceServer interface {
	GetStaticData(context.Context, *staticdatav1.GetStaticDataRequest) (*staticdatav1.GetStaticDataResponse, error)
	SearchServiceProvider(context.Context, *serviceproviderv1.SearchServiceProviderRequest) (*serviceproviderv1.SearchServiceProviderResponse, error)
	WidgetRegisterInstrumentInitiate(context.Context, *instrumentv1.WidgetRegisterInstrumentInitiateRequest) (*instrumentv1.WidgetRegisterInstrumentInitiateResponse, error)
	// optional call required when user uses username/password and they have more than one account at the service
	// provider
	WidgetRegisterInstrumentWithAccount(context.Context, *instrumentv1.WidgetRegisterInstrumentWithAccountRequest) (*instrumentv1.WidgetRegisterInstrumentWithAccountResponse, error)
	WidgetRegisterInstrumentComplete(context.Context, *instrumentv1.WidgetRegisterInstrumentCompleteRequest) (*instrumentv1.WidgetRegisterInstrumentCompleteResponse, error)
	WidgetRibbitRegisterInstrumentInitiate(context.Context, *instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest) (*instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse, error)
	Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error)
}

// UnimplementedConnectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConnectServiceServer struct {
}

func (*UnimplementedConnectServiceServer) GetStaticData(context.Context, *staticdatav1.GetStaticDataRequest) (*staticdatav1.GetStaticDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaticData not implemented")
}
func (*UnimplementedConnectServiceServer) SearchServiceProvider(context.Context, *serviceproviderv1.SearchServiceProviderRequest) (*serviceproviderv1.SearchServiceProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchServiceProvider not implemented")
}
func (*UnimplementedConnectServiceServer) WidgetRegisterInstrumentInitiate(context.Context, *instrumentv1.WidgetRegisterInstrumentInitiateRequest) (*instrumentv1.WidgetRegisterInstrumentInitiateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WidgetRegisterInstrumentInitiate not implemented")
}
func (*UnimplementedConnectServiceServer) WidgetRegisterInstrumentWithAccount(context.Context, *instrumentv1.WidgetRegisterInstrumentWithAccountRequest) (*instrumentv1.WidgetRegisterInstrumentWithAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WidgetRegisterInstrumentWithAccount not implemented")
}
func (*UnimplementedConnectServiceServer) WidgetRegisterInstrumentComplete(context.Context, *instrumentv1.WidgetRegisterInstrumentCompleteRequest) (*instrumentv1.WidgetRegisterInstrumentCompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WidgetRegisterInstrumentComplete not implemented")
}
func (*UnimplementedConnectServiceServer) WidgetRibbitRegisterInstrumentInitiate(context.Context, *instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest) (*instrumentv1.WidgetRibbitRegisterInstrumentInitiateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WidgetRibbitRegisterInstrumentInitiate not implemented")
}
func (*UnimplementedConnectServiceServer) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterConnectServiceServer(s *grpc.Server, srv ConnectServiceServer) {
	s.RegisterService(&_ConnectService_serviceDesc, srv)
}

func _ConnectService_GetStaticData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(staticdatav1.GetStaticDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).GetStaticData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/GetStaticData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).GetStaticData(ctx, req.(*staticdatav1.GetStaticDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_SearchServiceProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serviceproviderv1.SearchServiceProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).SearchServiceProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/SearchServiceProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).SearchServiceProvider(ctx, req.(*serviceproviderv1.SearchServiceProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_WidgetRegisterInstrumentInitiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.WidgetRegisterInstrumentInitiateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentInitiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/WidgetRegisterInstrumentInitiate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentInitiate(ctx, req.(*instrumentv1.WidgetRegisterInstrumentInitiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_WidgetRegisterInstrumentWithAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.WidgetRegisterInstrumentWithAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentWithAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/WidgetRegisterInstrumentWithAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentWithAccount(ctx, req.(*instrumentv1.WidgetRegisterInstrumentWithAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_WidgetRegisterInstrumentComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.WidgetRegisterInstrumentCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/WidgetRegisterInstrumentComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).WidgetRegisterInstrumentComplete(ctx, req.(*instrumentv1.WidgetRegisterInstrumentCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_WidgetRibbitRegisterInstrumentInitiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).WidgetRibbitRegisterInstrumentInitiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/WidgetRibbitRegisterInstrumentInitiate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).WidgetRibbitRegisterInstrumentInitiate(ctx, req.(*instrumentv1.WidgetRibbitRegisterInstrumentInitiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pingv1.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.service.v1.ConnectService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).Ping(ctx, req.(*pingv1.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connect.service.v1.ConnectService",
	HandlerType: (*ConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaticData",
			Handler:    _ConnectService_GetStaticData_Handler,
		},
		{
			MethodName: "SearchServiceProvider",
			Handler:    _ConnectService_SearchServiceProvider_Handler,
		},
		{
			MethodName: "WidgetRegisterInstrumentInitiate",
			Handler:    _ConnectService_WidgetRegisterInstrumentInitiate_Handler,
		},
		{
			MethodName: "WidgetRegisterInstrumentWithAccount",
			Handler:    _ConnectService_WidgetRegisterInstrumentWithAccount_Handler,
		},
		{
			MethodName: "WidgetRegisterInstrumentComplete",
			Handler:    _ConnectService_WidgetRegisterInstrumentComplete_Handler,
		},
		{
			MethodName: "WidgetRibbitRegisterInstrumentInitiate",
			Handler:    _ConnectService_WidgetRibbitRegisterInstrumentInitiate_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ConnectService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connect/service/v1/connect_service.proto",
}