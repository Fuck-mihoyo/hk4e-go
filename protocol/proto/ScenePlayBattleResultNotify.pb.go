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
// source: ScenePlayBattleResultNotify.proto

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

// CmdId: 4398
// EnetChannelId: 0
// EnetIsReliable: true
type ScenePlayBattleResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWin                bool                               `protobuf:"varint,1,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	CostTime             uint32                             `protobuf:"varint,7,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	PlayType             uint32                             `protobuf:"varint,15,opt,name=play_type,json=playType,proto3" json:"play_type,omitempty"`
	PlayId               uint32                             `protobuf:"varint,11,opt,name=play_id,json=playId,proto3" json:"play_id,omitempty"`
	SettlePlayerInfoList []*ScenePlayBattleSettlePlayerInfo `protobuf:"bytes,4,rep,name=settle_player_info_list,json=settlePlayerInfoList,proto3" json:"settle_player_info_list,omitempty"`
	SettleRewardInfoList []*ScenePlayBattleSettleRewardInfo `protobuf:"bytes,14,rep,name=settle_reward_info_list,json=settleRewardInfoList,proto3" json:"settle_reward_info_list,omitempty"`
}

func (x *ScenePlayBattleResultNotify) Reset() {
	*x = ScenePlayBattleResultNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlayBattleResultNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayBattleResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayBattleResultNotify) ProtoMessage() {}

func (x *ScenePlayBattleResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlayBattleResultNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayBattleResultNotify.ProtoReflect.Descriptor instead.
func (*ScenePlayBattleResultNotify) Descriptor() ([]byte, []int) {
	return file_ScenePlayBattleResultNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayBattleResultNotify) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *ScenePlayBattleResultNotify) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *ScenePlayBattleResultNotify) GetPlayType() uint32 {
	if x != nil {
		return x.PlayType
	}
	return 0
}

func (x *ScenePlayBattleResultNotify) GetPlayId() uint32 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

func (x *ScenePlayBattleResultNotify) GetSettlePlayerInfoList() []*ScenePlayBattleSettlePlayerInfo {
	if x != nil {
		return x.SettlePlayerInfoList
	}
	return nil
}

func (x *ScenePlayBattleResultNotify) GetSettleRewardInfoList() []*ScenePlayBattleSettleRewardInfo {
	if x != nil {
		return x.SettleRewardInfoList
	}
	return nil
}

var File_ScenePlayBattleResultNotify_proto protoreflect.FileDescriptor

var file_ScenePlayBattleResultNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x1b, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x5d, 0x0a, 0x17, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x5d, 0x0a, 0x17, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlayBattleResultNotify_proto_rawDescOnce sync.Once
	file_ScenePlayBattleResultNotify_proto_rawDescData = file_ScenePlayBattleResultNotify_proto_rawDesc
)

func file_ScenePlayBattleResultNotify_proto_rawDescGZIP() []byte {
	file_ScenePlayBattleResultNotify_proto_rawDescOnce.Do(func() {
		file_ScenePlayBattleResultNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlayBattleResultNotify_proto_rawDescData)
	})
	return file_ScenePlayBattleResultNotify_proto_rawDescData
}

var file_ScenePlayBattleResultNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlayBattleResultNotify_proto_goTypes = []interface{}{
	(*ScenePlayBattleResultNotify)(nil),     // 0: proto.ScenePlayBattleResultNotify
	(*ScenePlayBattleSettlePlayerInfo)(nil), // 1: proto.ScenePlayBattleSettlePlayerInfo
	(*ScenePlayBattleSettleRewardInfo)(nil), // 2: proto.ScenePlayBattleSettleRewardInfo
}
var file_ScenePlayBattleResultNotify_proto_depIdxs = []int32{
	1, // 0: proto.ScenePlayBattleResultNotify.settle_player_info_list:type_name -> proto.ScenePlayBattleSettlePlayerInfo
	2, // 1: proto.ScenePlayBattleResultNotify.settle_reward_info_list:type_name -> proto.ScenePlayBattleSettleRewardInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ScenePlayBattleResultNotify_proto_init() }
func file_ScenePlayBattleResultNotify_proto_init() {
	if File_ScenePlayBattleResultNotify_proto != nil {
		return
	}
	file_ScenePlayBattleSettlePlayerInfo_proto_init()
	file_ScenePlayBattleSettleRewardInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ScenePlayBattleResultNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayBattleResultNotify); i {
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
			RawDescriptor: file_ScenePlayBattleResultNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlayBattleResultNotify_proto_goTypes,
		DependencyIndexes: file_ScenePlayBattleResultNotify_proto_depIdxs,
		MessageInfos:      file_ScenePlayBattleResultNotify_proto_msgTypes,
	}.Build()
	File_ScenePlayBattleResultNotify_proto = out.File
	file_ScenePlayBattleResultNotify_proto_rawDesc = nil
	file_ScenePlayBattleResultNotify_proto_goTypes = nil
	file_ScenePlayBattleResultNotify_proto_depIdxs = nil
}
