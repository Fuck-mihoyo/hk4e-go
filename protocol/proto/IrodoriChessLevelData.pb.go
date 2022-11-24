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
// source: IrodoriChessLevelData.proto

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

type IrodoriChessLevelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenTime      uint32               `protobuf:"varint,8,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	LevelId       uint32               `protobuf:"varint,15,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	HardMapData   *IrodoriChessMapData `protobuf:"bytes,7,opt,name=hard_map_data,json=hardMapData,proto3" json:"hard_map_data,omitempty"`
	NormalMapData *IrodoriChessMapData `protobuf:"bytes,11,opt,name=normal_map_data,json=normalMapData,proto3" json:"normal_map_data,omitempty"`
}

func (x *IrodoriChessLevelData) Reset() {
	*x = IrodoriChessLevelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IrodoriChessLevelData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IrodoriChessLevelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IrodoriChessLevelData) ProtoMessage() {}

func (x *IrodoriChessLevelData) ProtoReflect() protoreflect.Message {
	mi := &file_IrodoriChessLevelData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IrodoriChessLevelData.ProtoReflect.Descriptor instead.
func (*IrodoriChessLevelData) Descriptor() ([]byte, []int) {
	return file_IrodoriChessLevelData_proto_rawDescGZIP(), []int{0}
}

func (x *IrodoriChessLevelData) GetOpenTime() uint32 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *IrodoriChessLevelData) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *IrodoriChessLevelData) GetHardMapData() *IrodoriChessMapData {
	if x != nil {
		return x.HardMapData
	}
	return nil
}

func (x *IrodoriChessLevelData) GetNormalMapData() *IrodoriChessMapData {
	if x != nil {
		return x.NormalMapData
	}
	return nil
}

var File_IrodoriChessLevelData_proto protoreflect.FileDescriptor

var file_IrodoriChessLevelData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd3, 0x01, 0x0a, 0x15, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x42, 0x0a, 0x0f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d,
	0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4d, 0x61,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IrodoriChessLevelData_proto_rawDescOnce sync.Once
	file_IrodoriChessLevelData_proto_rawDescData = file_IrodoriChessLevelData_proto_rawDesc
)

func file_IrodoriChessLevelData_proto_rawDescGZIP() []byte {
	file_IrodoriChessLevelData_proto_rawDescOnce.Do(func() {
		file_IrodoriChessLevelData_proto_rawDescData = protoimpl.X.CompressGZIP(file_IrodoriChessLevelData_proto_rawDescData)
	})
	return file_IrodoriChessLevelData_proto_rawDescData
}

var file_IrodoriChessLevelData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IrodoriChessLevelData_proto_goTypes = []interface{}{
	(*IrodoriChessLevelData)(nil), // 0: proto.IrodoriChessLevelData
	(*IrodoriChessMapData)(nil),   // 1: proto.IrodoriChessMapData
}
var file_IrodoriChessLevelData_proto_depIdxs = []int32{
	1, // 0: proto.IrodoriChessLevelData.hard_map_data:type_name -> proto.IrodoriChessMapData
	1, // 1: proto.IrodoriChessLevelData.normal_map_data:type_name -> proto.IrodoriChessMapData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_IrodoriChessLevelData_proto_init() }
func file_IrodoriChessLevelData_proto_init() {
	if File_IrodoriChessLevelData_proto != nil {
		return
	}
	file_IrodoriChessMapData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IrodoriChessLevelData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IrodoriChessLevelData); i {
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
			RawDescriptor: file_IrodoriChessLevelData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IrodoriChessLevelData_proto_goTypes,
		DependencyIndexes: file_IrodoriChessLevelData_proto_depIdxs,
		MessageInfos:      file_IrodoriChessLevelData_proto_msgTypes,
	}.Build()
	File_IrodoriChessLevelData_proto = out.File
	file_IrodoriChessLevelData_proto_rawDesc = nil
	file_IrodoriChessLevelData_proto_goTypes = nil
	file_IrodoriChessLevelData_proto_depIdxs = nil
}
