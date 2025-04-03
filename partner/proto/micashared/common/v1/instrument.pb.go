// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.28.3
// source: micashared/common/v1/instrument.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	instrumenttypev1 "github.com/1080network/golang/partner/proto/micashared/common/enums/instrumenttypev1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LinkedInstrumentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    instrumenttypev1.InstrumentType `protobuf:"varint,2,opt,name=type,proto3,enum=micashared.common.enums.instrumenttype.v1.InstrumentType" json:"type,omitempty"`
	IconUrl string                          `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
}

func (x *LinkedInstrumentData) Reset() {
	*x = LinkedInstrumentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_instrument_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedInstrumentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedInstrumentData) ProtoMessage() {}

func (x *LinkedInstrumentData) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_instrument_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedInstrumentData.ProtoReflect.Descriptor instead.
func (*LinkedInstrumentData) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_instrument_proto_rawDescGZIP(), []int{0}
}

func (x *LinkedInstrumentData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LinkedInstrumentData) GetType() instrumenttypev1.InstrumentType {
	if x != nil {
		return x.Type
	}
	return instrumenttypev1.InstrumentType(0)
}

func (x *LinkedInstrumentData) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

// A message that allows to exchange information between the different sides of the network
// when linking an instrument. This message is made up of a couple of generic
// claims/assertions and also structured data
type InstrumentLinkClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assertions      map[string]string `protobuf:"bytes,1,rep,name=assertions,proto3" json:"assertions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserDemographic *UserDemographic  `protobuf:"bytes,2,opt,name=user_demographic,json=userDemographic,proto3" json:"user_demographic,omitempty"`
}

func (x *InstrumentLinkClaims) Reset() {
	*x = InstrumentLinkClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_micashared_common_v1_instrument_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstrumentLinkClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentLinkClaims) ProtoMessage() {}

func (x *InstrumentLinkClaims) ProtoReflect() protoreflect.Message {
	mi := &file_micashared_common_v1_instrument_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentLinkClaims.ProtoReflect.Descriptor instead.
func (*InstrumentLinkClaims) Descriptor() ([]byte, []int) {
	return file_micashared_common_v1_instrument_proto_rawDescGZIP(), []int{1}
}

func (x *InstrumentLinkClaims) GetAssertions() map[string]string {
	if x != nil {
		return x.Assertions
	}
	return nil
}

func (x *InstrumentLinkClaims) GetUserDemographic() *UserDemographic {
	if x != nil {
		return x.UserDemographic
	}
	return nil
}

var File_micashared_common_v1_instrument_proto protoreflect.FileDescriptor

var file_micashared_common_v1_instrument_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x6d,
	0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f,
	0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x6d, 0x69, 0x63,
	0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x83, 0x02, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12,
	0x5a, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x10, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x0f, 0x75, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x1a, 0x3d, 0x0a,
	0x0f, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x50, 0x0a, 0x17,
	0x69, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x6d, 0x69, 0x63, 0x61, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x4d, 0x49, 0x43, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_micashared_common_v1_instrument_proto_rawDescOnce sync.Once
	file_micashared_common_v1_instrument_proto_rawDescData = file_micashared_common_v1_instrument_proto_rawDesc
)

func file_micashared_common_v1_instrument_proto_rawDescGZIP() []byte {
	file_micashared_common_v1_instrument_proto_rawDescOnce.Do(func() {
		file_micashared_common_v1_instrument_proto_rawDescData = protoimpl.X.CompressGZIP(file_micashared_common_v1_instrument_proto_rawDescData)
	})
	return file_micashared_common_v1_instrument_proto_rawDescData
}

var file_micashared_common_v1_instrument_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_micashared_common_v1_instrument_proto_goTypes = []interface{}{
	(*LinkedInstrumentData)(nil),         // 0: micashared.common.v1.LinkedInstrumentData
	(*InstrumentLinkClaims)(nil),         // 1: micashared.common.v1.InstrumentLinkClaims
	nil,                                  // 2: micashared.common.v1.InstrumentLinkClaims.AssertionsEntry
	(instrumenttypev1.InstrumentType)(0), // 3: micashared.common.enums.instrumenttype.v1.InstrumentType
	(*UserDemographic)(nil),              // 4: micashared.common.v1.UserDemographic
}
var file_micashared_common_v1_instrument_proto_depIdxs = []int32{
	3, // 0: micashared.common.v1.LinkedInstrumentData.type:type_name -> micashared.common.enums.instrumenttype.v1.InstrumentType
	2, // 1: micashared.common.v1.InstrumentLinkClaims.assertions:type_name -> micashared.common.v1.InstrumentLinkClaims.AssertionsEntry
	4, // 2: micashared.common.v1.InstrumentLinkClaims.user_demographic:type_name -> micashared.common.v1.UserDemographic
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_micashared_common_v1_instrument_proto_init() }
func file_micashared_common_v1_instrument_proto_init() {
	if File_micashared_common_v1_instrument_proto != nil {
		return
	}
	file_micashared_common_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_micashared_common_v1_instrument_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedInstrumentData); i {
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
		file_micashared_common_v1_instrument_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstrumentLinkClaims); i {
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
			RawDescriptor: file_micashared_common_v1_instrument_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_micashared_common_v1_instrument_proto_goTypes,
		DependencyIndexes: file_micashared_common_v1_instrument_proto_depIdxs,
		MessageInfos:      file_micashared_common_v1_instrument_proto_msgTypes,
	}.Build()
	File_micashared_common_v1_instrument_proto = out.File
	file_micashared_common_v1_instrument_proto_rawDesc = nil
	file_micashared_common_v1_instrument_proto_goTypes = nil
	file_micashared_common_v1_instrument_proto_depIdxs = nil
}
