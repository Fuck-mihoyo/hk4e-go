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
// source: EffigyChallengeInfoNotify.proto

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

// CmdId: 2090
// EnetChannelId: 0
// EnetIsReliable: true
type EffigyChallengeInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DifficultyId    uint32   `protobuf:"varint,9,opt,name=difficulty_id,json=difficultyId,proto3" json:"difficulty_id,omitempty"`
	ConditionIdList []uint32 `protobuf:"varint,11,rep,packed,name=condition_id_list,json=conditionIdList,proto3" json:"condition_id_list,omitempty"`
	ChallengeScore  uint32   `protobuf:"varint,14,opt,name=challenge_score,json=challengeScore,proto3" json:"challenge_score,omitempty"`
	ChallengeId     uint32   `protobuf:"varint,8,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
}

func (x *EffigyChallengeInfoNotify) Reset() {
	*x = EffigyChallengeInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EffigyChallengeInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffigyChallengeInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffigyChallengeInfoNotify) ProtoMessage() {}

func (x *EffigyChallengeInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EffigyChallengeInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffigyChallengeInfoNotify.ProtoReflect.Descriptor instead.
func (*EffigyChallengeInfoNotify) Descriptor() ([]byte, []int) {
	return file_EffigyChallengeInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EffigyChallengeInfoNotify) GetDifficultyId() uint32 {
	if x != nil {
		return x.DifficultyId
	}
	return 0
}

func (x *EffigyChallengeInfoNotify) GetConditionIdList() []uint32 {
	if x != nil {
		return x.ConditionIdList
	}
	return nil
}

func (x *EffigyChallengeInfoNotify) GetChallengeScore() uint32 {
	if x != nil {
		return x.ChallengeScore
	}
	return 0
}

func (x *EffigyChallengeInfoNotify) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

var File_EffigyChallengeInfoNotify_proto protoreflect.FileDescriptor

var file_EffigyChallengeInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x66, 0x66, 0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x19, 0x45, 0x66, 0x66,
	0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EffigyChallengeInfoNotify_proto_rawDescOnce sync.Once
	file_EffigyChallengeInfoNotify_proto_rawDescData = file_EffigyChallengeInfoNotify_proto_rawDesc
)

func file_EffigyChallengeInfoNotify_proto_rawDescGZIP() []byte {
	file_EffigyChallengeInfoNotify_proto_rawDescOnce.Do(func() {
		file_EffigyChallengeInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EffigyChallengeInfoNotify_proto_rawDescData)
	})
	return file_EffigyChallengeInfoNotify_proto_rawDescData
}

var file_EffigyChallengeInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EffigyChallengeInfoNotify_proto_goTypes = []interface{}{
	(*EffigyChallengeInfoNotify)(nil), // 0: proto.EffigyChallengeInfoNotify
}
var file_EffigyChallengeInfoNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EffigyChallengeInfoNotify_proto_init() }
func file_EffigyChallengeInfoNotify_proto_init() {
	if File_EffigyChallengeInfoNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EffigyChallengeInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffigyChallengeInfoNotify); i {
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
			RawDescriptor: file_EffigyChallengeInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EffigyChallengeInfoNotify_proto_goTypes,
		DependencyIndexes: file_EffigyChallengeInfoNotify_proto_depIdxs,
		MessageInfos:      file_EffigyChallengeInfoNotify_proto_msgTypes,
	}.Build()
	File_EffigyChallengeInfoNotify_proto = out.File
	file_EffigyChallengeInfoNotify_proto_rawDesc = nil
	file_EffigyChallengeInfoNotify_proto_goTypes = nil
	file_EffigyChallengeInfoNotify_proto_depIdxs = nil
}
