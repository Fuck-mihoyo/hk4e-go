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
// source: GadgetPlayUidInfo.proto

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

type GadgetPlayUidInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicture  *ProfilePicture `protobuf:"bytes,2,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	BattleWatcherId uint32          `protobuf:"varint,6,opt,name=battle_watcher_id,json=battleWatcherId,proto3" json:"battle_watcher_id,omitempty"`
	Uid             uint32          `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Icon            uint32          `protobuf:"varint,14,opt,name=icon,proto3" json:"icon,omitempty"`
	Score           uint32          `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Nickname        string          `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	OnlineId        string          `protobuf:"bytes,8,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
}

func (x *GadgetPlayUidInfo) Reset() {
	*x = GadgetPlayUidInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GadgetPlayUidInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GadgetPlayUidInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GadgetPlayUidInfo) ProtoMessage() {}

func (x *GadgetPlayUidInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GadgetPlayUidInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GadgetPlayUidInfo.ProtoReflect.Descriptor instead.
func (*GadgetPlayUidInfo) Descriptor() ([]byte, []int) {
	return file_GadgetPlayUidInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GadgetPlayUidInfo) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *GadgetPlayUidInfo) GetBattleWatcherId() uint32 {
	if x != nil {
		return x.BattleWatcherId
	}
	return 0
}

func (x *GadgetPlayUidInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GadgetPlayUidInfo) GetIcon() uint32 {
	if x != nil {
		return x.Icon
	}
	return 0
}

func (x *GadgetPlayUidInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GadgetPlayUidInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GadgetPlayUidInfo) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

var File_GadgetPlayUidInfo_proto protoreflect.FileDescriptor

var file_GadgetPlayUidInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x69, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GadgetPlayUidInfo_proto_rawDescOnce sync.Once
	file_GadgetPlayUidInfo_proto_rawDescData = file_GadgetPlayUidInfo_proto_rawDesc
)

func file_GadgetPlayUidInfo_proto_rawDescGZIP() []byte {
	file_GadgetPlayUidInfo_proto_rawDescOnce.Do(func() {
		file_GadgetPlayUidInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetPlayUidInfo_proto_rawDescData)
	})
	return file_GadgetPlayUidInfo_proto_rawDescData
}

var file_GadgetPlayUidInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GadgetPlayUidInfo_proto_goTypes = []interface{}{
	(*GadgetPlayUidInfo)(nil), // 0: proto.GadgetPlayUidInfo
	(*ProfilePicture)(nil),    // 1: proto.ProfilePicture
}
var file_GadgetPlayUidInfo_proto_depIdxs = []int32{
	1, // 0: proto.GadgetPlayUidInfo.profile_picture:type_name -> proto.ProfilePicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GadgetPlayUidInfo_proto_init() }
func file_GadgetPlayUidInfo_proto_init() {
	if File_GadgetPlayUidInfo_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GadgetPlayUidInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GadgetPlayUidInfo); i {
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
			RawDescriptor: file_GadgetPlayUidInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetPlayUidInfo_proto_goTypes,
		DependencyIndexes: file_GadgetPlayUidInfo_proto_depIdxs,
		MessageInfos:      file_GadgetPlayUidInfo_proto_msgTypes,
	}.Build()
	File_GadgetPlayUidInfo_proto = out.File
	file_GadgetPlayUidInfo_proto_rawDesc = nil
	file_GadgetPlayUidInfo_proto_goTypes = nil
	file_GadgetPlayUidInfo_proto_depIdxs = nil
}
