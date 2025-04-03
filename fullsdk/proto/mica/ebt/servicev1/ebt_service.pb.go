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
// source: mica/ebt/service/v1/ebt_service.proto

package servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	listitemv1 "github.com/1080network/golang/fullsdk/proto/mica/ebt/listitemv1"
	programv1 "github.com/1080network/golang/fullsdk/proto/mica/ebt/programv1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mica_ebt_service_v1_ebt_service_proto protoreflect.FileDescriptor

var file_mica_ebt_service_v1_ebt_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x65, 0x62, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x62, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d, 0x69,
	0x63, 0x61, 0x2f, 0x65, 0x62, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x65, 0x62, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8b, 0x06, 0x0a, 0x0a, 0x45, 0x42, 0x54, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x42, 0x54, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x29, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x42, 0x54, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x42, 0x54, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x42, 0x54, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x12, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x42, 0x54, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x42, 0x54,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x64, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x28, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x65, 0x62, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x65, 0x62, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65,
	0x62, 0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62,
	0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65, 0x62, 0x74, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x44, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x65,
	0x62, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x45,
	0x62, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x12,
	0x6d, 0x69, 0x63, 0x61, 0x2f, 0x65, 0x62, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_mica_ebt_service_v1_ebt_service_proto_goTypes = []interface{}{
	(*programv1.GetEBTProgramRequest)(nil),     // 0: mica.ebt.program.v1.GetEBTProgramRequest
	(*programv1.UpdateEBTProgramRequest)(nil),  // 1: mica.ebt.program.v1.UpdateEBTProgramRequest
	(*listitemv1.CreateListItemRequest)(nil),   // 2: mica.ebt.listitem.v1.CreateListItemRequest
	(*listitemv1.GetListItemRequest)(nil),      // 3: mica.ebt.listitem.v1.GetListItemRequest
	(*listitemv1.RemoveListItemRequest)(nil),   // 4: mica.ebt.listitem.v1.RemoveListItemRequest
	(*listitemv1.UpdateListItemRequest)(nil),   // 5: mica.ebt.listitem.v1.UpdateListItemRequest
	(*listitemv1.SearchListItemRequest)(nil),   // 6: mica.ebt.listitem.v1.SearchListItemRequest
	(*programv1.GetEBTProgramResponse)(nil),    // 7: mica.ebt.program.v1.GetEBTProgramResponse
	(*programv1.UpdateEBTProgramResponse)(nil), // 8: mica.ebt.program.v1.UpdateEBTProgramResponse
	(*listitemv1.CreateListItemResponse)(nil),  // 9: mica.ebt.listitem.v1.CreateListItemResponse
	(*listitemv1.GetListItemResponse)(nil),     // 10: mica.ebt.listitem.v1.GetListItemResponse
	(*listitemv1.RemoveListItemResponse)(nil),  // 11: mica.ebt.listitem.v1.RemoveListItemResponse
	(*listitemv1.UpdateListItemResponse)(nil),  // 12: mica.ebt.listitem.v1.UpdateListItemResponse
	(*listitemv1.SearchListItemResponse)(nil),  // 13: mica.ebt.listitem.v1.SearchListItemResponse
}
var file_mica_ebt_service_v1_ebt_service_proto_depIdxs = []int32{
	0,  // 0: mica.ebt.service.v1.EBTService.GetEBTProgram:input_type -> mica.ebt.program.v1.GetEBTProgramRequest
	1,  // 1: mica.ebt.service.v1.EBTService.UpdateEBTProgram:input_type -> mica.ebt.program.v1.UpdateEBTProgramRequest
	2,  // 2: mica.ebt.service.v1.EBTService.CreateListItem:input_type -> mica.ebt.listitem.v1.CreateListItemRequest
	3,  // 3: mica.ebt.service.v1.EBTService.GetListItem:input_type -> mica.ebt.listitem.v1.GetListItemRequest
	4,  // 4: mica.ebt.service.v1.EBTService.RemoveListItem:input_type -> mica.ebt.listitem.v1.RemoveListItemRequest
	5,  // 5: mica.ebt.service.v1.EBTService.UpdateListItem:input_type -> mica.ebt.listitem.v1.UpdateListItemRequest
	6,  // 6: mica.ebt.service.v1.EBTService.SearchListItem:input_type -> mica.ebt.listitem.v1.SearchListItemRequest
	7,  // 7: mica.ebt.service.v1.EBTService.GetEBTProgram:output_type -> mica.ebt.program.v1.GetEBTProgramResponse
	8,  // 8: mica.ebt.service.v1.EBTService.UpdateEBTProgram:output_type -> mica.ebt.program.v1.UpdateEBTProgramResponse
	9,  // 9: mica.ebt.service.v1.EBTService.CreateListItem:output_type -> mica.ebt.listitem.v1.CreateListItemResponse
	10, // 10: mica.ebt.service.v1.EBTService.GetListItem:output_type -> mica.ebt.listitem.v1.GetListItemResponse
	11, // 11: mica.ebt.service.v1.EBTService.RemoveListItem:output_type -> mica.ebt.listitem.v1.RemoveListItemResponse
	12, // 12: mica.ebt.service.v1.EBTService.UpdateListItem:output_type -> mica.ebt.listitem.v1.UpdateListItemResponse
	13, // 13: mica.ebt.service.v1.EBTService.SearchListItem:output_type -> mica.ebt.listitem.v1.SearchListItemResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mica_ebt_service_v1_ebt_service_proto_init() }
func file_mica_ebt_service_v1_ebt_service_proto_init() {
	if File_mica_ebt_service_v1_ebt_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mica_ebt_service_v1_ebt_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_ebt_service_v1_ebt_service_proto_goTypes,
		DependencyIndexes: file_mica_ebt_service_v1_ebt_service_proto_depIdxs,
	}.Build()
	File_mica_ebt_service_v1_ebt_service_proto = out.File
	file_mica_ebt_service_v1_ebt_service_proto_rawDesc = nil
	file_mica_ebt_service_v1_ebt_service_proto_goTypes = nil
	file_mica_ebt_service_v1_ebt_service_proto_depIdxs = nil
}
