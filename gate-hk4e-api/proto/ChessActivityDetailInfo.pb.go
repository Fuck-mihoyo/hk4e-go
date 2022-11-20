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
// source: ChessActivityDetailInfo.proto

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

type ChessActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level                  uint32   `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	IsTeachDungeonFinished bool     `protobuf:"varint,9,opt,name=is_teach_dungeon_finished,json=isTeachDungeonFinished,proto3" json:"is_teach_dungeon_finished,omitempty"`
	ContentCloseTime       uint32   `protobuf:"varint,14,opt,name=content_close_time,json=contentCloseTime,proto3" json:"content_close_time,omitempty"`
	ObtainedExp            uint32   `protobuf:"varint,8,opt,name=obtained_exp,json=obtainedExp,proto3" json:"obtained_exp,omitempty"`
	IsContentClosed        bool     `protobuf:"varint,5,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	AvailableExp           uint32   `protobuf:"varint,2,opt,name=available_exp,json=availableExp,proto3" json:"available_exp,omitempty"`
	Exp                    uint32   `protobuf:"varint,13,opt,name=exp,proto3" json:"exp,omitempty"`
	FinishedMapIdList      []uint32 `protobuf:"varint,1,rep,packed,name=finished_map_id_list,json=finishedMapIdList,proto3" json:"finished_map_id_list,omitempty"`
	PunishOverTime         uint32   `protobuf:"varint,3,opt,name=punish_over_time,json=punishOverTime,proto3" json:"punish_over_time,omitempty"`
}

func (x *ChessActivityDetailInfo) Reset() {
	*x = ChessActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessActivityDetailInfo) ProtoMessage() {}

func (x *ChessActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*ChessActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_ChessActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessActivityDetailInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ChessActivityDetailInfo) GetIsTeachDungeonFinished() bool {
	if x != nil {
		return x.IsTeachDungeonFinished
	}
	return false
}

func (x *ChessActivityDetailInfo) GetContentCloseTime() uint32 {
	if x != nil {
		return x.ContentCloseTime
	}
	return 0
}

func (x *ChessActivityDetailInfo) GetObtainedExp() uint32 {
	if x != nil {
		return x.ObtainedExp
	}
	return 0
}

func (x *ChessActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *ChessActivityDetailInfo) GetAvailableExp() uint32 {
	if x != nil {
		return x.AvailableExp
	}
	return 0
}

func (x *ChessActivityDetailInfo) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *ChessActivityDetailInfo) GetFinishedMapIdList() []uint32 {
	if x != nil {
		return x.FinishedMapIdList
	}
	return nil
}

func (x *ChessActivityDetailInfo) GetPunishOverTime() uint32 {
	if x != nil {
		return x.PunishOverTime
	}
	return 0
}

var File_ChessActivityDetailInfo_proto protoreflect.FileDescriptor

var file_ChessActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x02, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x39, 0x0a, 0x19, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x5f, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x78, 0x70, 0x12, 0x2a, 0x0a,
	0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70,
	0x12, 0x2f, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x75, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x75, 0x6e,
	0x69, 0x73, 0x68, 0x4f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessActivityDetailInfo_proto_rawDescOnce sync.Once
	file_ChessActivityDetailInfo_proto_rawDescData = file_ChessActivityDetailInfo_proto_rawDesc
)

func file_ChessActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_ChessActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_ChessActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessActivityDetailInfo_proto_rawDescData)
	})
	return file_ChessActivityDetailInfo_proto_rawDescData
}

var file_ChessActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessActivityDetailInfo_proto_goTypes = []interface{}{
	(*ChessActivityDetailInfo)(nil), // 0: ChessActivityDetailInfo
}
var file_ChessActivityDetailInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessActivityDetailInfo_proto_init() }
func file_ChessActivityDetailInfo_proto_init() {
	if File_ChessActivityDetailInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessActivityDetailInfo); i {
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
			RawDescriptor: file_ChessActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_ChessActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_ChessActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_ChessActivityDetailInfo_proto = out.File
	file_ChessActivityDetailInfo_proto_rawDesc = nil
	file_ChessActivityDetailInfo_proto_goTypes = nil
	file_ChessActivityDetailInfo_proto_depIdxs = nil
}
