// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: DataResVersionNotify.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataResVersionNotify_DataResVersionOpType int32

const (
	DataResVersionNotify_DATA_RES_VERSION_OP_TYPE_NONE       DataResVersionNotify_DataResVersionOpType = 0
	DataResVersionNotify_DATA_RES_VERSION_OP_TYPE_RELOGIN    DataResVersionNotify_DataResVersionOpType = 1
	DataResVersionNotify_DATA_RES_VERSION_OP_TYPE_MP_RELOGIN DataResVersionNotify_DataResVersionOpType = 2
)

// Enum value maps for DataResVersionNotify_DataResVersionOpType.
var (
	DataResVersionNotify_DataResVersionOpType_name = map[int32]string{
		0: "DATA_RES_VERSION_OP_TYPE_NONE",
		1: "DATA_RES_VERSION_OP_TYPE_RELOGIN",
		2: "DATA_RES_VERSION_OP_TYPE_MP_RELOGIN",
	}
	DataResVersionNotify_DataResVersionOpType_value = map[string]int32{
		"DATA_RES_VERSION_OP_TYPE_NONE":       0,
		"DATA_RES_VERSION_OP_TYPE_RELOGIN":    1,
		"DATA_RES_VERSION_OP_TYPE_MP_RELOGIN": 2,
	}
)

func (x DataResVersionNotify_DataResVersionOpType) Enum() *DataResVersionNotify_DataResVersionOpType {
	p := new(DataResVersionNotify_DataResVersionOpType)
	*p = x
	return p
}

func (x DataResVersionNotify_DataResVersionOpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataResVersionNotify_DataResVersionOpType) Descriptor() protoreflect.EnumDescriptor {
	return file_DataResVersionNotify_proto_enumTypes[0].Descriptor()
}

func (DataResVersionNotify_DataResVersionOpType) Type() protoreflect.EnumType {
	return &file_DataResVersionNotify_proto_enumTypes[0]
}

func (x DataResVersionNotify_DataResVersionOpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataResVersionNotify_DataResVersionOpType.Descriptor instead.
func (DataResVersionNotify_DataResVersionOpType) EnumDescriptor() ([]byte, []int) {
	return file_DataResVersionNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 167
// EnetChannelId: 0
// EnetIsReliable: true
type DataResVersionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientSilenceMd5           string                                    `protobuf:"bytes,10,opt,name=client_silence_md5,json=clientSilenceMd5,proto3" json:"client_silence_md5,omitempty"`
	ClientSilenceVersionSuffix string                                    `protobuf:"bytes,15,opt,name=client_silence_version_suffix,json=clientSilenceVersionSuffix,proto3" json:"client_silence_version_suffix,omitempty"`
	ResVersionConfig           *ResVersionConfig                         `protobuf:"bytes,9,opt,name=res_version_config,json=resVersionConfig,proto3" json:"res_version_config,omitempty"`
	IsDataNeedRelogin          bool                                      `protobuf:"varint,7,opt,name=is_data_need_relogin,json=isDataNeedRelogin,proto3" json:"is_data_need_relogin,omitempty"`
	OpType                     DataResVersionNotify_DataResVersionOpType `protobuf:"varint,12,opt,name=op_type,json=opType,proto3,enum=DataResVersionNotify_DataResVersionOpType" json:"op_type,omitempty"`
	ClientDataVersion          uint32                                    `protobuf:"varint,2,opt,name=client_data_version,json=clientDataVersion,proto3" json:"client_data_version,omitempty"`
	ClientVersionSuffix        string                                    `protobuf:"bytes,5,opt,name=client_version_suffix,json=clientVersionSuffix,proto3" json:"client_version_suffix,omitempty"`
	ClientSilenceDataVersion   uint32                                    `protobuf:"varint,1,opt,name=client_silence_data_version,json=clientSilenceDataVersion,proto3" json:"client_silence_data_version,omitempty"`
	ClientMd5                  string                                    `protobuf:"bytes,14,opt,name=client_md5,json=clientMd5,proto3" json:"client_md5,omitempty"`
}

func (x *DataResVersionNotify) Reset() {
	*x = DataResVersionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataResVersionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResVersionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResVersionNotify) ProtoMessage() {}

func (x *DataResVersionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DataResVersionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResVersionNotify.ProtoReflect.Descriptor instead.
func (*DataResVersionNotify) Descriptor() ([]byte, []int) {
	return file_DataResVersionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DataResVersionNotify) GetClientSilenceMd5() string {
	if x != nil {
		return x.ClientSilenceMd5
	}
	return ""
}

func (x *DataResVersionNotify) GetClientSilenceVersionSuffix() string {
	if x != nil {
		return x.ClientSilenceVersionSuffix
	}
	return ""
}

func (x *DataResVersionNotify) GetResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.ResVersionConfig
	}
	return nil
}

func (x *DataResVersionNotify) GetIsDataNeedRelogin() bool {
	if x != nil {
		return x.IsDataNeedRelogin
	}
	return false
}

func (x *DataResVersionNotify) GetOpType() DataResVersionNotify_DataResVersionOpType {
	if x != nil {
		return x.OpType
	}
	return DataResVersionNotify_DATA_RES_VERSION_OP_TYPE_NONE
}

func (x *DataResVersionNotify) GetClientDataVersion() uint32 {
	if x != nil {
		return x.ClientDataVersion
	}
	return 0
}

func (x *DataResVersionNotify) GetClientVersionSuffix() string {
	if x != nil {
		return x.ClientVersionSuffix
	}
	return ""
}

func (x *DataResVersionNotify) GetClientSilenceDataVersion() uint32 {
	if x != nil {
		return x.ClientSilenceDataVersion
	}
	return 0
}

func (x *DataResVersionNotify) GetClientMd5() string {
	if x != nil {
		return x.ClientMd5
	}
	return ""
}

var File_DataResVersionNotify_proto protoreflect.FileDescriptor

var file_DataResVersionNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x05, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x6d, 0x64, 0x35, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x64, 0x35, 0x12, 0x41, 0x0a, 0x1d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x3f,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72,
	0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x65, 0x65, 0x64, 0x52, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x43, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x35, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x1d, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x52, 0x45, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x5f, 0x52, 0x45, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DataResVersionNotify_proto_rawDescOnce sync.Once
	file_DataResVersionNotify_proto_rawDescData = file_DataResVersionNotify_proto_rawDesc
)

func file_DataResVersionNotify_proto_rawDescGZIP() []byte {
	file_DataResVersionNotify_proto_rawDescOnce.Do(func() {
		file_DataResVersionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DataResVersionNotify_proto_rawDescData)
	})
	return file_DataResVersionNotify_proto_rawDescData
}

var file_DataResVersionNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DataResVersionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DataResVersionNotify_proto_goTypes = []interface{}{
	(DataResVersionNotify_DataResVersionOpType)(0), // 0: DataResVersionNotify.DataResVersionOpType
	(*DataResVersionNotify)(nil),                   // 1: DataResVersionNotify
	(*ResVersionConfig)(nil),                       // 2: ResVersionConfig
}
var file_DataResVersionNotify_proto_depIdxs = []int32{
	2, // 0: DataResVersionNotify.res_version_config:type_name -> ResVersionConfig
	0, // 1: DataResVersionNotify.op_type:type_name -> DataResVersionNotify.DataResVersionOpType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_DataResVersionNotify_proto_init() }
func file_DataResVersionNotify_proto_init() {
	if File_DataResVersionNotify_proto != nil {
		return
	}
	file_ResVersionConfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DataResVersionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResVersionNotify); i {
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
			RawDescriptor: file_DataResVersionNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DataResVersionNotify_proto_goTypes,
		DependencyIndexes: file_DataResVersionNotify_proto_depIdxs,
		EnumInfos:         file_DataResVersionNotify_proto_enumTypes,
		MessageInfos:      file_DataResVersionNotify_proto_msgTypes,
	}.Build()
	File_DataResVersionNotify_proto = out.File
	file_DataResVersionNotify_proto_rawDesc = nil
	file_DataResVersionNotify_proto_goTypes = nil
	file_DataResVersionNotify_proto_depIdxs = nil
}
