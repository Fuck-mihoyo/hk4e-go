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
// source: RogueDiaryRepairInfoNotify.proto

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

// CmdId: 8641
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type RogueDiaryRepairInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId          uint32                `protobuf:"varint,8,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	SelectCardList   []uint32              `protobuf:"varint,14,rep,packed,name=select_card_list,json=selectCardList,proto3" json:"select_card_list,omitempty"`
	AvatarList       []*RogueDiaryAvatar   `protobuf:"bytes,13,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	RoomList         []*RogueDiaryRoomInfo `protobuf:"bytes,2,rep,name=room_list,json=roomList,proto3" json:"room_list,omitempty"`
	RandCardList     []uint32              `protobuf:"varint,10,rep,packed,name=rand_card_list,json=randCardList,proto3" json:"rand_card_list,omitempty"`
	SelectAvatarList []*RogueDiaryAvatar   `protobuf:"bytes,9,rep,name=select_avatar_list,json=selectAvatarList,proto3" json:"select_avatar_list,omitempty"`
	ChosenCardList   []uint32              `protobuf:"varint,15,rep,packed,name=chosen_card_list,json=chosenCardList,proto3" json:"chosen_card_list,omitempty"`
	TrialAvatarList  []*RogueDiaryAvatar   `protobuf:"bytes,11,rep,name=trial_avatar_list,json=trialAvatarList,proto3" json:"trial_avatar_list,omitempty"`
}

func (x *RogueDiaryRepairInfoNotify) Reset() {
	*x = RogueDiaryRepairInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueDiaryRepairInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueDiaryRepairInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueDiaryRepairInfoNotify) ProtoMessage() {}

func (x *RogueDiaryRepairInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RogueDiaryRepairInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueDiaryRepairInfoNotify.ProtoReflect.Descriptor instead.
func (*RogueDiaryRepairInfoNotify) Descriptor() ([]byte, []int) {
	return file_RogueDiaryRepairInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RogueDiaryRepairInfoNotify) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RogueDiaryRepairInfoNotify) GetSelectCardList() []uint32 {
	if x != nil {
		return x.SelectCardList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetAvatarList() []*RogueDiaryAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetRoomList() []*RogueDiaryRoomInfo {
	if x != nil {
		return x.RoomList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetRandCardList() []uint32 {
	if x != nil {
		return x.RandCardList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetSelectAvatarList() []*RogueDiaryAvatar {
	if x != nil {
		return x.SelectAvatarList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetChosenCardList() []uint32 {
	if x != nil {
		return x.ChosenCardList
	}
	return nil
}

func (x *RogueDiaryRepairInfoNotify) GetTrialAvatarList() []*RogueDiaryAvatar {
	if x != nil {
		return x.TrialAvatarList
	}
	return nil
}

var File_RogueDiaryRepairInfoNotify_proto protoreflect.FileDescriptor

var file_RogueDiaryRepairInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x1a,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x10, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e,
	0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x11, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0f, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RogueDiaryRepairInfoNotify_proto_rawDescOnce sync.Once
	file_RogueDiaryRepairInfoNotify_proto_rawDescData = file_RogueDiaryRepairInfoNotify_proto_rawDesc
)

func file_RogueDiaryRepairInfoNotify_proto_rawDescGZIP() []byte {
	file_RogueDiaryRepairInfoNotify_proto_rawDescOnce.Do(func() {
		file_RogueDiaryRepairInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueDiaryRepairInfoNotify_proto_rawDescData)
	})
	return file_RogueDiaryRepairInfoNotify_proto_rawDescData
}

var file_RogueDiaryRepairInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueDiaryRepairInfoNotify_proto_goTypes = []interface{}{
	(*RogueDiaryRepairInfoNotify)(nil), // 0: proto.RogueDiaryRepairInfoNotify
	(*RogueDiaryAvatar)(nil),           // 1: proto.RogueDiaryAvatar
	(*RogueDiaryRoomInfo)(nil),         // 2: proto.RogueDiaryRoomInfo
}
var file_RogueDiaryRepairInfoNotify_proto_depIdxs = []int32{
	1, // 0: proto.RogueDiaryRepairInfoNotify.avatar_list:type_name -> proto.RogueDiaryAvatar
	2, // 1: proto.RogueDiaryRepairInfoNotify.room_list:type_name -> proto.RogueDiaryRoomInfo
	1, // 2: proto.RogueDiaryRepairInfoNotify.select_avatar_list:type_name -> proto.RogueDiaryAvatar
	1, // 3: proto.RogueDiaryRepairInfoNotify.trial_avatar_list:type_name -> proto.RogueDiaryAvatar
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueDiaryRepairInfoNotify_proto_init() }
func file_RogueDiaryRepairInfoNotify_proto_init() {
	if File_RogueDiaryRepairInfoNotify_proto != nil {
		return
	}
	file_RogueDiaryAvatar_proto_init()
	file_RogueDiaryRoomInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueDiaryRepairInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueDiaryRepairInfoNotify); i {
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
			RawDescriptor: file_RogueDiaryRepairInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueDiaryRepairInfoNotify_proto_goTypes,
		DependencyIndexes: file_RogueDiaryRepairInfoNotify_proto_depIdxs,
		MessageInfos:      file_RogueDiaryRepairInfoNotify_proto_msgTypes,
	}.Build()
	File_RogueDiaryRepairInfoNotify_proto = out.File
	file_RogueDiaryRepairInfoNotify_proto_rawDesc = nil
	file_RogueDiaryRepairInfoNotify_proto_goTypes = nil
	file_RogueDiaryRepairInfoNotify_proto_depIdxs = nil
}
