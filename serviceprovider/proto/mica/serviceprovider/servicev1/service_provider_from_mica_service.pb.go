// Copyright (c) 2022 Mica. All rights reserved. All software, including, without limitation, all source
// code and object code, is the intellectual property of Mica, Inc. and is protected by copyright, trademark and
// other intellectual property laws (collective "Mica Software"). You may not use, copy, reproduce, download, store,
// post, broadcast, transmit, modify, sell or make available to the public content from the Mica Software without a
// valid license or the prior written approval of Mica, Inc. Mica, its logos, slogans, taglines,
// products, feature names, and other trademarks are trademarks of Mica, Inc. and may not be used without
// permission.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: mica/serviceprovider/service/v1/service_provider_from_mica_service.proto

package servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	instrumentv1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/instrumentv1"
	userv1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/userv1"
	valuev1 "github.com/1080network/golang/serviceprovider/proto/mica/serviceprovider/valuev1"
	pingv1 "github.com/1080network/golang/serviceprovider/proto/micashared/common/pingv1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto protoreflect.FileDescriptor

var file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6d, 0x69, 0x63, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x33, 0x6d, 0x69, 0x63,
	0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x69, 0x63, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93, 0x06, 0x0a, 0x1e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01,
	0x0a, 0x14, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x98, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0b, 0x4f, 0x62,
	0x74, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d,
	0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x74,
	0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x76, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x0c, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x70, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x23, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x63, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1e, 0x6d, 0x69, 0x63,
	0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49,
	0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_goTypes = []interface{}{
	(*userv1.EnrollUserInstrumentRequest)(nil),       // 0: mica.serviceprovider.user.v1.EnrollUserInstrumentRequest
	(*instrumentv1.RetrieveTransactionRequest)(nil),  // 1: mica.serviceprovider.instrument.v1.RetrieveTransactionRequest
	(*valuev1.ObtainValueRequest)(nil),               // 2: mica.serviceprovider.value.v1.ObtainValueRequest
	(*valuev1.ReturnValueRequest)(nil),               // 3: mica.serviceprovider.value.v1.ReturnValueRequest
	(*valuev1.ReceiveValueRequest)(nil),              // 4: mica.serviceprovider.value.v1.ReceiveValueRequest
	(*pingv1.PingRequest)(nil),                       // 5: micashared.common.ping.v1.PingRequest
	(*userv1.EnrollUserInstrumentResponse)(nil),      // 6: mica.serviceprovider.user.v1.EnrollUserInstrumentResponse
	(*instrumentv1.RetrieveTransactionResponse)(nil), // 7: mica.serviceprovider.instrument.v1.RetrieveTransactionResponse
	(*valuev1.ObtainValueResponse)(nil),              // 8: mica.serviceprovider.value.v1.ObtainValueResponse
	(*valuev1.ReturnValueResponse)(nil),              // 9: mica.serviceprovider.value.v1.ReturnValueResponse
	(*valuev1.ReceiveValueResponse)(nil),             // 10: mica.serviceprovider.value.v1.ReceiveValueResponse
	(*pingv1.PingResponse)(nil),                      // 11: micashared.common.ping.v1.PingResponse
}
var file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_depIdxs = []int32{
	0,  // 0: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.EnrollUserInstrument:input_type -> mica.serviceprovider.user.v1.EnrollUserInstrumentRequest
	1,  // 1: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.RetrieveTransaction:input_type -> mica.serviceprovider.instrument.v1.RetrieveTransactionRequest
	2,  // 2: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ObtainValue:input_type -> mica.serviceprovider.value.v1.ObtainValueRequest
	3,  // 3: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ReturnValue:input_type -> mica.serviceprovider.value.v1.ReturnValueRequest
	4,  // 4: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ReceiveValue:input_type -> mica.serviceprovider.value.v1.ReceiveValueRequest
	5,  // 5: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.Ping:input_type -> micashared.common.ping.v1.PingRequest
	6,  // 6: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.EnrollUserInstrument:output_type -> mica.serviceprovider.user.v1.EnrollUserInstrumentResponse
	7,  // 7: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.RetrieveTransaction:output_type -> mica.serviceprovider.instrument.v1.RetrieveTransactionResponse
	8,  // 8: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ObtainValue:output_type -> mica.serviceprovider.value.v1.ObtainValueResponse
	9,  // 9: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ReturnValue:output_type -> mica.serviceprovider.value.v1.ReturnValueResponse
	10, // 10: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.ReceiveValue:output_type -> mica.serviceprovider.value.v1.ReceiveValueResponse
	11, // 11: mica.serviceprovider.service.v1.ServiceProviderFromMicaService.Ping:output_type -> micashared.common.ping.v1.PingResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_init() }
func file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_init() {
	if File_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_goTypes,
		DependencyIndexes: file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_depIdxs,
	}.Build()
	File_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto = out.File
	file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_rawDesc = nil
	file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_goTypes = nil
	file_mica_serviceprovider_service_v1_service_provider_from_mica_service_proto_depIdxs = nil
}
