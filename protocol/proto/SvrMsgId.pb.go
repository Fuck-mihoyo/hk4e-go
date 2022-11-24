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
// source: SvrMsgId.proto

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

type SvrMsgId int32

const (
	SvrMsgId_SVR_MSG_ID_UNKNOWN                     SvrMsgId = 0
	SvrMsgId_SVR_MSG_ID_BLOCK_REFRESH_COUNTDOWN     SvrMsgId = 1
	SvrMsgId_SVR_MSG_ID_AVATAR_REVIVE_BY_STATUE     SvrMsgId = 2
	SvrMsgId_SVR_MSG_ID_DAILY_TASK_REWARD_MAX_NUM   SvrMsgId = 3
	SvrMsgId_SVR_MSG_ID_ROUTINE_TYPE_NOT_OPEN       SvrMsgId = 4
	SvrMsgId_SVR_MSG_ID_ROUTINE_TYPE_REWARD_MAX_NUM SvrMsgId = 5
	SvrMsgId_SVR_MSG_ID_MECHANICUS_COIN_LIMIT       SvrMsgId = 6
)

// Enum value maps for SvrMsgId.
var (
	SvrMsgId_name = map[int32]string{
		0: "SVR_MSG_ID_UNKNOWN",
		1: "SVR_MSG_ID_BLOCK_REFRESH_COUNTDOWN",
		2: "SVR_MSG_ID_AVATAR_REVIVE_BY_STATUE",
		3: "SVR_MSG_ID_DAILY_TASK_REWARD_MAX_NUM",
		4: "SVR_MSG_ID_ROUTINE_TYPE_NOT_OPEN",
		5: "SVR_MSG_ID_ROUTINE_TYPE_REWARD_MAX_NUM",
		6: "SVR_MSG_ID_MECHANICUS_COIN_LIMIT",
	}
	SvrMsgId_value = map[string]int32{
		"SVR_MSG_ID_UNKNOWN":                     0,
		"SVR_MSG_ID_BLOCK_REFRESH_COUNTDOWN":     1,
		"SVR_MSG_ID_AVATAR_REVIVE_BY_STATUE":     2,
		"SVR_MSG_ID_DAILY_TASK_REWARD_MAX_NUM":   3,
		"SVR_MSG_ID_ROUTINE_TYPE_NOT_OPEN":       4,
		"SVR_MSG_ID_ROUTINE_TYPE_REWARD_MAX_NUM": 5,
		"SVR_MSG_ID_MECHANICUS_COIN_LIMIT":       6,
	}
)

func (x SvrMsgId) Enum() *SvrMsgId {
	p := new(SvrMsgId)
	*p = x
	return p
}

func (x SvrMsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SvrMsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_SvrMsgId_proto_enumTypes[0].Descriptor()
}

func (SvrMsgId) Type() protoreflect.EnumType {
	return &file_SvrMsgId_proto_enumTypes[0]
}

func (x SvrMsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SvrMsgId.Descriptor instead.
func (SvrMsgId) EnumDescriptor() ([]byte, []int) {
	return file_SvrMsgId_proto_rawDescGZIP(), []int{0}
}

var File_SvrMsgId_proto protoreflect.FileDescriptor

var file_SvrMsgId_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x76, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x94, 0x02, 0x0a, 0x08, 0x53, 0x76, 0x72, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22,
	0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x49, 0x44, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x56, 0x45,
	0x5f, 0x42, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x45, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24,
	0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x58,
	0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53,
	0x47, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26,
	0x53, 0x56, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x49,
	0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x4d,
	0x41, 0x58, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x56, 0x52, 0x5f,
	0x4d, 0x53, 0x47, 0x5f, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x43, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x06, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SvrMsgId_proto_rawDescOnce sync.Once
	file_SvrMsgId_proto_rawDescData = file_SvrMsgId_proto_rawDesc
)

func file_SvrMsgId_proto_rawDescGZIP() []byte {
	file_SvrMsgId_proto_rawDescOnce.Do(func() {
		file_SvrMsgId_proto_rawDescData = protoimpl.X.CompressGZIP(file_SvrMsgId_proto_rawDescData)
	})
	return file_SvrMsgId_proto_rawDescData
}

var file_SvrMsgId_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SvrMsgId_proto_goTypes = []interface{}{
	(SvrMsgId)(0), // 0: proto.SvrMsgId
}
var file_SvrMsgId_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SvrMsgId_proto_init() }
func file_SvrMsgId_proto_init() {
	if File_SvrMsgId_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SvrMsgId_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SvrMsgId_proto_goTypes,
		DependencyIndexes: file_SvrMsgId_proto_depIdxs,
		EnumInfos:         file_SvrMsgId_proto_enumTypes,
	}.Build()
	File_SvrMsgId_proto = out.File
	file_SvrMsgId_proto_rawDesc = nil
	file_SvrMsgId_proto_goTypes = nil
	file_SvrMsgId_proto_depIdxs = nil
}
