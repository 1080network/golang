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
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	v1 "github.com/1080network/golang/fullsdk/proto/micashared/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDataExtractionStatisticsResponse_Status int32

const (
	GetDataExtractionStatisticsResponse_STATUS_UNSPECIFIED GetDataExtractionStatisticsResponse_Status = 0
	GetDataExtractionStatisticsResponse_STATUS_SUCCESS     GetDataExtractionStatisticsResponse_Status = 1
	GetDataExtractionStatisticsResponse_STATUS_NOT_FOUND   GetDataExtractionStatisticsResponse_Status = 2
	GetDataExtractionStatisticsResponse_STATUS_ERROR       GetDataExtractionStatisticsResponse_Status = 3
)

// Enum value maps for GetDataExtractionStatisticsResponse_Status.
var (
	GetDataExtractionStatisticsResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_NOT_FOUND",
		3: "STATUS_ERROR",
	}
	GetDataExtractionStatisticsResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_NOT_FOUND":   2,
		"STATUS_ERROR":       3,
	}
)

func (x GetDataExtractionStatisticsResponse_Status) Enum() *GetDataExtractionStatisticsResponse_Status {
	p := new(GetDataExtractionStatisticsResponse_Status)
	*p = x
	return p
}

func (x GetDataExtractionStatisticsResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetDataExtractionStatisticsResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_administration_v1_admin_service_proto_enumTypes[0].Descriptor()
}

func (GetDataExtractionStatisticsResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_administration_v1_admin_service_proto_enumTypes[0]
}

func (x GetDataExtractionStatisticsResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetDataExtractionStatisticsResponse_Status.Descriptor instead.
func (GetDataExtractionStatisticsResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{2, 0}
}

type SearchDataExtractionStatisticsResponse_Status int32

const (
	SearchDataExtractionStatisticsResponse_STATUS_UNSPECIFIED SearchDataExtractionStatisticsResponse_Status = 0
	SearchDataExtractionStatisticsResponse_STATUS_SUCCESS     SearchDataExtractionStatisticsResponse_Status = 1
	SearchDataExtractionStatisticsResponse_STATUS_ERROR       SearchDataExtractionStatisticsResponse_Status = 2
)

// Enum value maps for SearchDataExtractionStatisticsResponse_Status.
var (
	SearchDataExtractionStatisticsResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_ERROR",
	}
	SearchDataExtractionStatisticsResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_ERROR":       2,
	}
)

func (x SearchDataExtractionStatisticsResponse_Status) Enum() *SearchDataExtractionStatisticsResponse_Status {
	p := new(SearchDataExtractionStatisticsResponse_Status)
	*p = x
	return p
}

func (x SearchDataExtractionStatisticsResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchDataExtractionStatisticsResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_mica_partner_administration_v1_admin_service_proto_enumTypes[1].Descriptor()
}

func (SearchDataExtractionStatisticsResponse_Status) Type() protoreflect.EnumType {
	return &file_mica_partner_administration_v1_admin_service_proto_enumTypes[1]
}

func (x SearchDataExtractionStatisticsResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchDataExtractionStatisticsResponse_Status.Descriptor instead.
func (SearchDataExtractionStatisticsResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{4, 0}
}

type DataExtraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serial number of the generated certificate
	ExtractorName     string                 `protobuf:"bytes,1,opt,name=extractor_name,json=extractorName,proto3" json:"extractor_name,omitempty"`
	WatermarkTime     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=watermark_time,json=watermarkTime,proto3" json:"watermark_time,omitempty"`
	LastExecutionTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_execution_time,json=lastExecutionTime,proto3" json:"last_execution_time,omitempty"`
	RecordsProcessed  uint32                 `protobuf:"varint,4,opt,name=records_processed,json=recordsProcessed,proto3" json:"records_processed,omitempty"`
}

func (x *DataExtraction) Reset() {
	*x = DataExtraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataExtraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataExtraction) ProtoMessage() {}

func (x *DataExtraction) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataExtraction.ProtoReflect.Descriptor instead.
func (*DataExtraction) Descriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{0}
}

func (x *DataExtraction) GetExtractorName() string {
	if x != nil {
		return x.ExtractorName
	}
	return ""
}

func (x *DataExtraction) GetWatermarkTime() *timestamppb.Timestamp {
	if x != nil {
		return x.WatermarkTime
	}
	return nil
}

func (x *DataExtraction) GetLastExecutionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastExecutionTime
	}
	return nil
}

func (x *DataExtraction) GetRecordsProcessed() uint32 {
	if x != nil {
		return x.RecordsProcessed
	}
	return 0
}

type GetDataExtractionStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtractorName string `protobuf:"bytes,1,opt,name=extractor_name,json=extractorName,proto3" json:"extractor_name,omitempty"`
}

func (x *GetDataExtractionStatisticsRequest) Reset() {
	*x = GetDataExtractionStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataExtractionStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataExtractionStatisticsRequest) ProtoMessage() {}

func (x *GetDataExtractionStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataExtractionStatisticsRequest.ProtoReflect.Descriptor instead.
func (*GetDataExtractionStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataExtractionStatisticsRequest) GetExtractorName() string {
	if x != nil {
		return x.ExtractorName
	}
	return ""
}

type GetDataExtractionStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     GetDataExtractionStatisticsResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.administration.v1.GetDataExtractionStatisticsResponse_Status" json:"status,omitempty"`
	Error      *v1.Error                                  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Extraction *DataExtraction                            `protobuf:"bytes,3,opt,name=extraction,proto3" json:"extraction,omitempty"`
}

func (x *GetDataExtractionStatisticsResponse) Reset() {
	*x = GetDataExtractionStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataExtractionStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataExtractionStatisticsResponse) ProtoMessage() {}

func (x *GetDataExtractionStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataExtractionStatisticsResponse.ProtoReflect.Descriptor instead.
func (*GetDataExtractionStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataExtractionStatisticsResponse) GetStatus() GetDataExtractionStatisticsResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetDataExtractionStatisticsResponse_STATUS_UNSPECIFIED
}

func (x *GetDataExtractionStatisticsResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetDataExtractionStatisticsResponse) GetExtraction() *DataExtraction {
	if x != nil {
		return x.Extraction
	}
	return nil
}

type SearchDataExtractionStatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchDataExtractionStatisticsRequest) Reset() {
	*x = SearchDataExtractionStatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDataExtractionStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataExtractionStatisticsRequest) ProtoMessage() {}

func (x *SearchDataExtractionStatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataExtractionStatisticsRequest.ProtoReflect.Descriptor instead.
func (*SearchDataExtractionStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{3}
}

type SearchDataExtractionStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      SearchDataExtractionStatisticsResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=mica.partner.administration.v1.SearchDataExtractionStatisticsResponse_Status" json:"status,omitempty"`
	Error       *v1.Error                                     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Extractions []*DataExtraction                             `protobuf:"bytes,3,rep,name=extractions,proto3" json:"extractions,omitempty"`
}

func (x *SearchDataExtractionStatisticsResponse) Reset() {
	*x = SearchDataExtractionStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDataExtractionStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDataExtractionStatisticsResponse) ProtoMessage() {}

func (x *SearchDataExtractionStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mica_partner_administration_v1_admin_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDataExtractionStatisticsResponse.ProtoReflect.Descriptor instead.
func (*SearchDataExtractionStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchDataExtractionStatisticsResponse) GetStatus() SearchDataExtractionStatisticsResponse_Status {
	if x != nil {
		return x.Status
	}
	return SearchDataExtractionStatisticsResponse_STATUS_UNSPECIFIED
}

func (x *SearchDataExtractionStatisticsResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchDataExtractionStatisticsResponse) GetExtractions() []*DataExtraction {
	if x != nil {
		return x.Extractions
	}
	return nil
}

var File_mica_partner_administration_v1_admin_service_proto protoreflect.FileDescriptor

var file_mica_partner_administration_v1_admin_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x02, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x0d, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4a, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x54, 0x0a, 0x22, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xea, 0x02, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4e,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x27, 0x0a, 0x25,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xdc, 0x02, 0x0a, 0x26, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x65, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4d, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x46, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x32, 0xfd, 0x02, 0x0a, 0x1c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb1, 0x01, 0x0a, 0x1e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x45, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x46, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa8, 0x01, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x42, 0x2e, 0x6d, 0x69, 0x63, 0x61,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e,
	0x6d, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x65, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mica_partner_administration_v1_admin_service_proto_rawDescOnce sync.Once
	file_mica_partner_administration_v1_admin_service_proto_rawDescData = file_mica_partner_administration_v1_admin_service_proto_rawDesc
)

func file_mica_partner_administration_v1_admin_service_proto_rawDescGZIP() []byte {
	file_mica_partner_administration_v1_admin_service_proto_rawDescOnce.Do(func() {
		file_mica_partner_administration_v1_admin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mica_partner_administration_v1_admin_service_proto_rawDescData)
	})
	return file_mica_partner_administration_v1_admin_service_proto_rawDescData
}

var file_mica_partner_administration_v1_admin_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mica_partner_administration_v1_admin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mica_partner_administration_v1_admin_service_proto_goTypes = []interface{}{
	(GetDataExtractionStatisticsResponse_Status)(0),    // 0: mica.partner.administration.v1.GetDataExtractionStatisticsResponse.Status
	(SearchDataExtractionStatisticsResponse_Status)(0), // 1: mica.partner.administration.v1.SearchDataExtractionStatisticsResponse.Status
	(*DataExtraction)(nil),                             // 2: mica.partner.administration.v1.DataExtraction
	(*GetDataExtractionStatisticsRequest)(nil),         // 3: mica.partner.administration.v1.GetDataExtractionStatisticsRequest
	(*GetDataExtractionStatisticsResponse)(nil),        // 4: mica.partner.administration.v1.GetDataExtractionStatisticsResponse
	(*SearchDataExtractionStatisticsRequest)(nil),      // 5: mica.partner.administration.v1.SearchDataExtractionStatisticsRequest
	(*SearchDataExtractionStatisticsResponse)(nil),     // 6: mica.partner.administration.v1.SearchDataExtractionStatisticsResponse
	(*timestamppb.Timestamp)(nil),                      // 7: google.protobuf.Timestamp
	(*v1.Error)(nil),                                   // 8: micashared.common.v1.Error
}
var file_mica_partner_administration_v1_admin_service_proto_depIdxs = []int32{
	7,  // 0: mica.partner.administration.v1.DataExtraction.watermark_time:type_name -> google.protobuf.Timestamp
	7,  // 1: mica.partner.administration.v1.DataExtraction.last_execution_time:type_name -> google.protobuf.Timestamp
	0,  // 2: mica.partner.administration.v1.GetDataExtractionStatisticsResponse.status:type_name -> mica.partner.administration.v1.GetDataExtractionStatisticsResponse.Status
	8,  // 3: mica.partner.administration.v1.GetDataExtractionStatisticsResponse.error:type_name -> micashared.common.v1.Error
	2,  // 4: mica.partner.administration.v1.GetDataExtractionStatisticsResponse.extraction:type_name -> mica.partner.administration.v1.DataExtraction
	1,  // 5: mica.partner.administration.v1.SearchDataExtractionStatisticsResponse.status:type_name -> mica.partner.administration.v1.SearchDataExtractionStatisticsResponse.Status
	8,  // 6: mica.partner.administration.v1.SearchDataExtractionStatisticsResponse.error:type_name -> micashared.common.v1.Error
	2,  // 7: mica.partner.administration.v1.SearchDataExtractionStatisticsResponse.extractions:type_name -> mica.partner.administration.v1.DataExtraction
	5,  // 8: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:input_type -> mica.partner.administration.v1.SearchDataExtractionStatisticsRequest
	3,  // 9: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:input_type -> mica.partner.administration.v1.GetDataExtractionStatisticsRequest
	6,  // 10: mica.partner.administration.v1.PartnerAdministrationService.SearchDataExtractionStatistics:output_type -> mica.partner.administration.v1.SearchDataExtractionStatisticsResponse
	4,  // 11: mica.partner.administration.v1.PartnerAdministrationService.GetDataExtractionStatistics:output_type -> mica.partner.administration.v1.GetDataExtractionStatisticsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mica_partner_administration_v1_admin_service_proto_init() }
func file_mica_partner_administration_v1_admin_service_proto_init() {
	if File_mica_partner_administration_v1_admin_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mica_partner_administration_v1_admin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataExtraction); i {
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
		file_mica_partner_administration_v1_admin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataExtractionStatisticsRequest); i {
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
		file_mica_partner_administration_v1_admin_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataExtractionStatisticsResponse); i {
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
		file_mica_partner_administration_v1_admin_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDataExtractionStatisticsRequest); i {
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
		file_mica_partner_administration_v1_admin_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDataExtractionStatisticsResponse); i {
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
			RawDescriptor: file_mica_partner_administration_v1_admin_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mica_partner_administration_v1_admin_service_proto_goTypes,
		DependencyIndexes: file_mica_partner_administration_v1_admin_service_proto_depIdxs,
		EnumInfos:         file_mica_partner_administration_v1_admin_service_proto_enumTypes,
		MessageInfos:      file_mica_partner_administration_v1_admin_service_proto_msgTypes,
	}.Build()
	File_mica_partner_administration_v1_admin_service_proto = out.File
	file_mica_partner_administration_v1_admin_service_proto_rawDesc = nil
	file_mica_partner_administration_v1_admin_service_proto_goTypes = nil
	file_mica_partner_administration_v1_admin_service_proto_depIdxs = nil
}
