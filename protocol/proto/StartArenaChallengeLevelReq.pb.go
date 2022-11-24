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
// source: StartArenaChallengeLevelReq.proto

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

// CmdId: 2127
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type StartArenaChallengeLevelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArenaChallengeId    uint32 `protobuf:"varint,4,opt,name=arena_challenge_id,json=arenaChallengeId,proto3" json:"arena_challenge_id,omitempty"`
	GadgetEntityId      uint32 `protobuf:"varint,5,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
	ArenaChallengeLevel uint32 `protobuf:"varint,2,opt,name=arena_challenge_level,json=arenaChallengeLevel,proto3" json:"arena_challenge_level,omitempty"`
}

func (x *StartArenaChallengeLevelReq) Reset() {
	*x = StartArenaChallengeLevelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartArenaChallengeLevelReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartArenaChallengeLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartArenaChallengeLevelReq) ProtoMessage() {}

func (x *StartArenaChallengeLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_StartArenaChallengeLevelReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartArenaChallengeLevelReq.ProtoReflect.Descriptor instead.
func (*StartArenaChallengeLevelReq) Descriptor() ([]byte, []int) {
	return file_StartArenaChallengeLevelReq_proto_rawDescGZIP(), []int{0}
}

func (x *StartArenaChallengeLevelReq) GetArenaChallengeId() uint32 {
	if x != nil {
		return x.ArenaChallengeId
	}
	return 0
}

func (x *StartArenaChallengeLevelReq) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

func (x *StartArenaChallengeLevelReq) GetArenaChallengeLevel() uint32 {
	if x != nil {
		return x.ArenaChallengeLevel
	}
	return 0
}

var File_StartArenaChallengeLevelReq_proto protoreflect.FileDescriptor

var file_StartArenaChallengeLevelReq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x1b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x72,
	0x65, 0x6e, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x13, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartArenaChallengeLevelReq_proto_rawDescOnce sync.Once
	file_StartArenaChallengeLevelReq_proto_rawDescData = file_StartArenaChallengeLevelReq_proto_rawDesc
)

func file_StartArenaChallengeLevelReq_proto_rawDescGZIP() []byte {
	file_StartArenaChallengeLevelReq_proto_rawDescOnce.Do(func() {
		file_StartArenaChallengeLevelReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartArenaChallengeLevelReq_proto_rawDescData)
	})
	return file_StartArenaChallengeLevelReq_proto_rawDescData
}

var file_StartArenaChallengeLevelReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartArenaChallengeLevelReq_proto_goTypes = []interface{}{
	(*StartArenaChallengeLevelReq)(nil), // 0: proto.StartArenaChallengeLevelReq
}
var file_StartArenaChallengeLevelReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_StartArenaChallengeLevelReq_proto_init() }
func file_StartArenaChallengeLevelReq_proto_init() {
	if File_StartArenaChallengeLevelReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_StartArenaChallengeLevelReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartArenaChallengeLevelReq); i {
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
			RawDescriptor: file_StartArenaChallengeLevelReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartArenaChallengeLevelReq_proto_goTypes,
		DependencyIndexes: file_StartArenaChallengeLevelReq_proto_depIdxs,
		MessageInfos:      file_StartArenaChallengeLevelReq_proto_msgTypes,
	}.Build()
	File_StartArenaChallengeLevelReq_proto = out.File
	file_StartArenaChallengeLevelReq_proto_rawDesc = nil
	file_StartArenaChallengeLevelReq_proto_goTypes = nil
	file_StartArenaChallengeLevelReq_proto_depIdxs = nil
}
