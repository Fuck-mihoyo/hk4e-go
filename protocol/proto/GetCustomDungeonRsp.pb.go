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
// 	protoc        v3.14.0
// source: GetCustomDungeonRsp.proto

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

// CmdId: 6227
// EnetChannelId: 0
// EnetIsReliable: true
type GetCustomDungeonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32                 `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BanInfo   *CustomDungeonBanInfo `protobuf:"bytes,14,opt,name=ban_info,json=banInfo,proto3" json:"ban_info,omitempty"`
	BriefList []*CustomDungeonBrief `protobuf:"bytes,5,rep,name=brief_list,json=briefList,proto3" json:"brief_list,omitempty"`
}

func (x *GetCustomDungeonRsp) Reset() {
	*x = GetCustomDungeonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCustomDungeonRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomDungeonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomDungeonRsp) ProtoMessage() {}

func (x *GetCustomDungeonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCustomDungeonRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomDungeonRsp.ProtoReflect.Descriptor instead.
func (*GetCustomDungeonRsp) Descriptor() ([]byte, []int) {
	return file_GetCustomDungeonRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomDungeonRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCustomDungeonRsp) GetBanInfo() *CustomDungeonBanInfo {
	if x != nil {
		return x.BanInfo
	}
	return nil
}

func (x *GetCustomDungeonRsp) GetBriefList() []*CustomDungeonBrief {
	if x != nil {
		return x.BriefList
	}
	return nil
}

var File_GetCustomDungeonRsp_proto protoreflect.FileDescriptor

var file_GetCustomDungeonRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x42, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x61,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x42, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x62, 0x61, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x65,
	0x66, 0x52, 0x09, 0x62, 0x72, 0x69, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCustomDungeonRsp_proto_rawDescOnce sync.Once
	file_GetCustomDungeonRsp_proto_rawDescData = file_GetCustomDungeonRsp_proto_rawDesc
)

func file_GetCustomDungeonRsp_proto_rawDescGZIP() []byte {
	file_GetCustomDungeonRsp_proto_rawDescOnce.Do(func() {
		file_GetCustomDungeonRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCustomDungeonRsp_proto_rawDescData)
	})
	return file_GetCustomDungeonRsp_proto_rawDescData
}

var file_GetCustomDungeonRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCustomDungeonRsp_proto_goTypes = []interface{}{
	(*GetCustomDungeonRsp)(nil),  // 0: proto.GetCustomDungeonRsp
	(*CustomDungeonBanInfo)(nil), // 1: proto.CustomDungeonBanInfo
	(*CustomDungeonBrief)(nil),   // 2: proto.CustomDungeonBrief
}
var file_GetCustomDungeonRsp_proto_depIdxs = []int32{
	1, // 0: proto.GetCustomDungeonRsp.ban_info:type_name -> proto.CustomDungeonBanInfo
	2, // 1: proto.GetCustomDungeonRsp.brief_list:type_name -> proto.CustomDungeonBrief
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetCustomDungeonRsp_proto_init() }
func file_GetCustomDungeonRsp_proto_init() {
	if File_GetCustomDungeonRsp_proto != nil {
		return
	}
	file_CustomDungeonBanInfo_proto_init()
	file_CustomDungeonBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCustomDungeonRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomDungeonRsp); i {
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
			RawDescriptor: file_GetCustomDungeonRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCustomDungeonRsp_proto_goTypes,
		DependencyIndexes: file_GetCustomDungeonRsp_proto_depIdxs,
		MessageInfos:      file_GetCustomDungeonRsp_proto_msgTypes,
	}.Build()
	File_GetCustomDungeonRsp_proto = out.File
	file_GetCustomDungeonRsp_proto_rawDesc = nil
	file_GetCustomDungeonRsp_proto_goTypes = nil
	file_GetCustomDungeonRsp_proto_depIdxs = nil
}
