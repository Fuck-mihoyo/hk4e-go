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
// source: DailyDungeonEntryInfo.proto

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

type DailyDungeonEntryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonEntryConfigId      uint32            `protobuf:"varint,12,opt,name=dungeon_entry_config_id,json=dungeonEntryConfigId,proto3" json:"dungeon_entry_config_id,omitempty"`
	DungeonEntryId            uint32            `protobuf:"varint,15,opt,name=dungeon_entry_id,json=dungeonEntryId,proto3" json:"dungeon_entry_id,omitempty"`
	RecommendDungeonEntryInfo *DungeonEntryInfo `protobuf:"bytes,1,opt,name=recommend_dungeon_entry_info,json=recommendDungeonEntryInfo,proto3" json:"recommend_dungeon_entry_info,omitempty"`
	RecommendDungeonId        uint32            `protobuf:"varint,4,opt,name=recommend_dungeon_id,json=recommendDungeonId,proto3" json:"recommend_dungeon_id,omitempty"`
}

func (x *DailyDungeonEntryInfo) Reset() {
	*x = DailyDungeonEntryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyDungeonEntryInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyDungeonEntryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyDungeonEntryInfo) ProtoMessage() {}

func (x *DailyDungeonEntryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_DailyDungeonEntryInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyDungeonEntryInfo.ProtoReflect.Descriptor instead.
func (*DailyDungeonEntryInfo) Descriptor() ([]byte, []int) {
	return file_DailyDungeonEntryInfo_proto_rawDescGZIP(), []int{0}
}

func (x *DailyDungeonEntryInfo) GetDungeonEntryConfigId() uint32 {
	if x != nil {
		return x.DungeonEntryConfigId
	}
	return 0
}

func (x *DailyDungeonEntryInfo) GetDungeonEntryId() uint32 {
	if x != nil {
		return x.DungeonEntryId
	}
	return 0
}

func (x *DailyDungeonEntryInfo) GetRecommendDungeonEntryInfo() *DungeonEntryInfo {
	if x != nil {
		return x.RecommendDungeonEntryInfo
	}
	return nil
}

func (x *DailyDungeonEntryInfo) GetRecommendDungeonId() uint32 {
	if x != nil {
		return x.RecommendDungeonId
	}
	return 0
}

var File_DailyDungeonEntryInfo_proto protoreflect.FileDescriptor

var file_DailyDungeonEntryInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a,
	0x15, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x17, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x1c, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x19, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DailyDungeonEntryInfo_proto_rawDescOnce sync.Once
	file_DailyDungeonEntryInfo_proto_rawDescData = file_DailyDungeonEntryInfo_proto_rawDesc
)

func file_DailyDungeonEntryInfo_proto_rawDescGZIP() []byte {
	file_DailyDungeonEntryInfo_proto_rawDescOnce.Do(func() {
		file_DailyDungeonEntryInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyDungeonEntryInfo_proto_rawDescData)
	})
	return file_DailyDungeonEntryInfo_proto_rawDescData
}

var file_DailyDungeonEntryInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyDungeonEntryInfo_proto_goTypes = []interface{}{
	(*DailyDungeonEntryInfo)(nil), // 0: proto.DailyDungeonEntryInfo
	(*DungeonEntryInfo)(nil),      // 1: proto.DungeonEntryInfo
}
var file_DailyDungeonEntryInfo_proto_depIdxs = []int32{
	1, // 0: proto.DailyDungeonEntryInfo.recommend_dungeon_entry_info:type_name -> proto.DungeonEntryInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DailyDungeonEntryInfo_proto_init() }
func file_DailyDungeonEntryInfo_proto_init() {
	if File_DailyDungeonEntryInfo_proto != nil {
		return
	}
	file_DungeonEntryInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DailyDungeonEntryInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyDungeonEntryInfo); i {
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
			RawDescriptor: file_DailyDungeonEntryInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyDungeonEntryInfo_proto_goTypes,
		DependencyIndexes: file_DailyDungeonEntryInfo_proto_depIdxs,
		MessageInfos:      file_DailyDungeonEntryInfo_proto_msgTypes,
	}.Build()
	File_DailyDungeonEntryInfo_proto = out.File
	file_DailyDungeonEntryInfo_proto_rawDesc = nil
	file_DailyDungeonEntryInfo_proto_goTypes = nil
	file_DailyDungeonEntryInfo_proto_depIdxs = nil
}
