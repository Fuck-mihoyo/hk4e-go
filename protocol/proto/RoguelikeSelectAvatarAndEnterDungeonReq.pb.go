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
// source: RoguelikeSelectAvatarAndEnterDungeonReq.proto

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

// CmdId: 8457
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type RoguelikeSelectAvatarAndEnterDungeonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnstageAvatarGuidList   []uint64 `protobuf:"varint,14,rep,packed,name=onstage_avatar_guid_list,json=onstageAvatarGuidList,proto3" json:"onstage_avatar_guid_list,omitempty"`
	StageId                 uint32   `protobuf:"varint,4,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	BackstageAvatarGuidList []uint64 `protobuf:"varint,11,rep,packed,name=backstage_avatar_guid_list,json=backstageAvatarGuidList,proto3" json:"backstage_avatar_guid_list,omitempty"`
}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) Reset() {
	*x = RoguelikeSelectAvatarAndEnterDungeonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoguelikeSelectAvatarAndEnterDungeonReq) ProtoMessage() {}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) ProtoReflect() protoreflect.Message {
	mi := &file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoguelikeSelectAvatarAndEnterDungeonReq.ProtoReflect.Descriptor instead.
func (*RoguelikeSelectAvatarAndEnterDungeonReq) Descriptor() ([]byte, []int) {
	return file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescGZIP(), []int{0}
}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) GetOnstageAvatarGuidList() []uint64 {
	if x != nil {
		return x.OnstageAvatarGuidList
	}
	return nil
}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RoguelikeSelectAvatarAndEnterDungeonReq) GetBackstageAvatarGuidList() []uint64 {
	if x != nil {
		return x.BackstageAvatarGuidList
	}
	return nil
}

var File_RoguelikeSelectAvatarAndEnterDungeonReq_proto protoreflect.FileDescriptor

var file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x27, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x6c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x41, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x37, 0x0a, 0x18, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x15, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x04, 0x52, 0x17, 0x62, 0x61, 0x63, 0x6b,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescOnce sync.Once
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescData = file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDesc
)

func file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescGZIP() []byte {
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescOnce.Do(func() {
		file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescData)
	})
	return file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDescData
}

var file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_goTypes = []interface{}{
	(*RoguelikeSelectAvatarAndEnterDungeonReq)(nil), // 0: proto.RoguelikeSelectAvatarAndEnterDungeonReq
}
var file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_init() }
func file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_init() {
	if File_RoguelikeSelectAvatarAndEnterDungeonReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoguelikeSelectAvatarAndEnterDungeonReq); i {
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
			RawDescriptor: file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_goTypes,
		DependencyIndexes: file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_depIdxs,
		MessageInfos:      file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_msgTypes,
	}.Build()
	File_RoguelikeSelectAvatarAndEnterDungeonReq_proto = out.File
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_rawDesc = nil
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_goTypes = nil
	file_RoguelikeSelectAvatarAndEnterDungeonReq_proto_depIdxs = nil
}
