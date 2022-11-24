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
// source: RogueDiaryDungeonInfoNotify.proto

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

// CmdId: 8597
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type RogueDiaryDungeonInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId          uint32                `protobuf:"varint,12,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	RoundMonsterList []uint32              `protobuf:"varint,15,rep,packed,name=round_monster_list,json=roundMonsterList,proto3" json:"round_monster_list,omitempty"`
	Time             uint32                `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	CurRoom          uint32                `protobuf:"varint,5,opt,name=cur_room,json=curRoom,proto3" json:"cur_room,omitempty"`
	CurRound         uint32                `protobuf:"varint,6,opt,name=cur_round,json=curRound,proto3" json:"cur_round,omitempty"`
	Coin             uint32                `protobuf:"varint,11,opt,name=coin,proto3" json:"coin,omitempty"`
	Difficulty       uint32                `protobuf:"varint,8,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	DungeonId        uint32                `protobuf:"varint,14,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	RoomList         []*RogueDiaryRoomInfo `protobuf:"bytes,7,rep,name=room_list,json=roomList,proto3" json:"room_list,omitempty"`
	RoundCardList    []uint32              `protobuf:"varint,10,rep,packed,name=round_card_list,json=roundCardList,proto3" json:"round_card_list,omitempty"`
}

func (x *RogueDiaryDungeonInfoNotify) Reset() {
	*x = RogueDiaryDungeonInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueDiaryDungeonInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueDiaryDungeonInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueDiaryDungeonInfoNotify) ProtoMessage() {}

func (x *RogueDiaryDungeonInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RogueDiaryDungeonInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueDiaryDungeonInfoNotify.ProtoReflect.Descriptor instead.
func (*RogueDiaryDungeonInfoNotify) Descriptor() ([]byte, []int) {
	return file_RogueDiaryDungeonInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RogueDiaryDungeonInfoNotify) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetRoundMonsterList() []uint32 {
	if x != nil {
		return x.RoundMonsterList
	}
	return nil
}

func (x *RogueDiaryDungeonInfoNotify) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetCurRoom() uint32 {
	if x != nil {
		return x.CurRoom
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetCurRound() uint32 {
	if x != nil {
		return x.CurRound
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *RogueDiaryDungeonInfoNotify) GetRoomList() []*RogueDiaryRoomInfo {
	if x != nil {
		return x.RoomList
	}
	return nil
}

func (x *RogueDiaryDungeonInfoNotify) GetRoundCardList() []uint32 {
	if x != nil {
		return x.RoundCardList
	}
	return nil
}

var File_RogueDiaryDungeonInfoNotify_proto protoreflect.FileDescriptor

var file_RogueDiaryDungeonInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69,
	0x61, 0x72, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x75, 0x72, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61,
	0x72, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueDiaryDungeonInfoNotify_proto_rawDescOnce sync.Once
	file_RogueDiaryDungeonInfoNotify_proto_rawDescData = file_RogueDiaryDungeonInfoNotify_proto_rawDesc
)

func file_RogueDiaryDungeonInfoNotify_proto_rawDescGZIP() []byte {
	file_RogueDiaryDungeonInfoNotify_proto_rawDescOnce.Do(func() {
		file_RogueDiaryDungeonInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueDiaryDungeonInfoNotify_proto_rawDescData)
	})
	return file_RogueDiaryDungeonInfoNotify_proto_rawDescData
}

var file_RogueDiaryDungeonInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueDiaryDungeonInfoNotify_proto_goTypes = []interface{}{
	(*RogueDiaryDungeonInfoNotify)(nil), // 0: proto.RogueDiaryDungeonInfoNotify
	(*RogueDiaryRoomInfo)(nil),          // 1: proto.RogueDiaryRoomInfo
}
var file_RogueDiaryDungeonInfoNotify_proto_depIdxs = []int32{
	1, // 0: proto.RogueDiaryDungeonInfoNotify.room_list:type_name -> proto.RogueDiaryRoomInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueDiaryDungeonInfoNotify_proto_init() }
func file_RogueDiaryDungeonInfoNotify_proto_init() {
	if File_RogueDiaryDungeonInfoNotify_proto != nil {
		return
	}
	file_RogueDiaryRoomInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueDiaryDungeonInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueDiaryDungeonInfoNotify); i {
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
			RawDescriptor: file_RogueDiaryDungeonInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueDiaryDungeonInfoNotify_proto_goTypes,
		DependencyIndexes: file_RogueDiaryDungeonInfoNotify_proto_depIdxs,
		MessageInfos:      file_RogueDiaryDungeonInfoNotify_proto_msgTypes,
	}.Build()
	File_RogueDiaryDungeonInfoNotify_proto = out.File
	file_RogueDiaryDungeonInfoNotify_proto_rawDesc = nil
	file_RogueDiaryDungeonInfoNotify_proto_goTypes = nil
	file_RogueDiaryDungeonInfoNotify_proto_depIdxs = nil
}
