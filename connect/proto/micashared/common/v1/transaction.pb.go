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
// source: micashared/common/v1/transaction.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Global transaction ID passed that reflects the ID of this transaction across the entire path through the network.
	// This can be used report issues and debug transactions
	GlobalTransactionId string `protobuf:"bytes,1,opt,name=global_transaction_id,json=globalTransactionId,proto3" json:"global_transaction_id,omitempty"`
	// Global timestamp passed that reflects the time when the request arrived at the network and is the same for all
	// services that are involved with processing a transaction. This field can be used when trying to reconcile data
	// between you and the network (assuming you store this timestamp).
	GlobalTransactionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=global_transaction_time,json=globalTransactionTime,proto3" json:"global_transaction_time,omitempty"`
}

func (x *TransactionIdentifier) Reset() {
	*x = TransactionIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionIdentifier) ProtoMessage() {}

func (x *TransactionIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionIdentifier.ProtoReflect.Descriptor instead.
func (*TransactionIdentifier) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionIdentifier) GetGlobalTransactionId() string {
	if x != nil {
		return x.GlobalTransactionId
	}
	return ""
}

func (x *TransactionIdentifier) GetGlobalTransactionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.GlobalTransactionTime
	}
	return nil
}

var File_micashared_common_v1_transaction_proto protoreflect.FileDescriptor

var file_micashared_common_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9f, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x52, 0x0a,
	0x17, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x51, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04,
	0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_transaction_proto_rawDescOnce sync.Once
	file_micashared_common_v1_transaction_proto_rawDescData = file_micashared_common_v1_transaction_proto_rawDesc
)

func file_micashared_common_v1_transaction_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_transaction_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_transaction_proto_rawDescData)
	})
	return file_micashared_common_v1_transaction_proto_rawDescData
}

var file_micashared_common_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_micashared_common_v1_transaction_proto_goTypes = []interface{}{
	(*TransactionIdentifier)(nil), // 0: micashared.common.v1.TransactionIdentifier
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_micashared_common_v1_transaction_proto_depIdxs = []int32{
	1, // 0: micashared.common.v1.TransactionIdentifier.global_transaction_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_transaction_proto_init() }
func file_micashared_common_v1_transaction_proto_init() {
	if File_micashared_common_v1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionIdentifier); i {
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
			RawDescriptor: file_micashared_common_v1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_transaction_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_transaction_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_transaction_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_transaction_proto = out.File
	file_micashared_common_v1_transaction_proto_rawDesc = nil
	file_micashared_common_v1_transaction_proto_goTypes = nil
	file_micashared_common_v1_transaction_proto_depIdxs = nil
}
