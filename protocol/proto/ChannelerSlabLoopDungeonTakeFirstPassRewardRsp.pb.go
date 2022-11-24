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
// source: ChannelerSlabLoopDungeonTakeFirstPassRewardRsp.proto

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

// CmdId: 8539
// EnetChannelId: 0
// EnetIsReliable: true
type ChannelerSlabLoopDungeonTakeFirstPassRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      int32  `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DungeonIndex uint32 `protobuf:"varint,8,opt,name=dungeon_index,json=dungeonIndex,proto3" json:"dungeon_index,omitempty"`
}

func (x *ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) Reset() {
	*x = ChannelerSlabLoopDungeonTakeFirstPassRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) ProtoMessage() {}

func (x *ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelerSlabLoopDungeonTakeFirstPassRewardRsp.ProtoReflect.Descriptor instead.
func (*ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) Descriptor() ([]byte, []int) {
	return file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChannelerSlabLoopDungeonTakeFirstPassRewardRsp) GetDungeonIndex() uint32 {
	if x != nil {
		return x.DungeonIndex
	}
	return 0
}

var File_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto protoreflect.FileDescriptor

var file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x34, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x4c,
	0x6f, 0x6f, 0x70, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x4c, 0x6f,
	0x6f, 0x70, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x54, 0x61, 0x6b, 0x65, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescOnce sync.Once
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescData = file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDesc
)

func file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescGZIP() []byte {
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescOnce.Do(func() {
		file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescData)
	})
	return file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDescData
}

var file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_goTypes = []interface{}{
	(*ChannelerSlabLoopDungeonTakeFirstPassRewardRsp)(nil), // 0: proto.ChannelerSlabLoopDungeonTakeFirstPassRewardRsp
}
var file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_init() }
func file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_init() {
	if File_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelerSlabLoopDungeonTakeFirstPassRewardRsp); i {
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
			RawDescriptor: file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_goTypes,
		DependencyIndexes: file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_depIdxs,
		MessageInfos:      file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_msgTypes,
	}.Build()
	File_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto = out.File
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_rawDesc = nil
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_goTypes = nil
	file_ChannelerSlabLoopDungeonTakeFirstPassRewardRsp_proto_depIdxs = nil
}
