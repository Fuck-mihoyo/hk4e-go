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
// source: EnterFungusFighterTrainingDungeonReq.proto

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

// CmdId: 23860
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type EnterFungusFighterTrainingDungeonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonId uint32 `protobuf:"varint,3,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
}

func (x *EnterFungusFighterTrainingDungeonReq) Reset() {
	*x = EnterFungusFighterTrainingDungeonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterFungusFighterTrainingDungeonReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterFungusFighterTrainingDungeonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterFungusFighterTrainingDungeonReq) ProtoMessage() {}

func (x *EnterFungusFighterTrainingDungeonReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterFungusFighterTrainingDungeonReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterFungusFighterTrainingDungeonReq.ProtoReflect.Descriptor instead.
func (*EnterFungusFighterTrainingDungeonReq) Descriptor() ([]byte, []int) {
	return file_EnterFungusFighterTrainingDungeonReq_proto_rawDescGZIP(), []int{0}
}

func (x *EnterFungusFighterTrainingDungeonReq) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

var File_EnterFungusFighterTrainingDungeonReq_proto protoreflect.FileDescriptor

var file_EnterFungusFighterTrainingDungeonReq_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x24, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterFungusFighterTrainingDungeonReq_proto_rawDescOnce sync.Once
	file_EnterFungusFighterTrainingDungeonReq_proto_rawDescData = file_EnterFungusFighterTrainingDungeonReq_proto_rawDesc
)

func file_EnterFungusFighterTrainingDungeonReq_proto_rawDescGZIP() []byte {
	file_EnterFungusFighterTrainingDungeonReq_proto_rawDescOnce.Do(func() {
		file_EnterFungusFighterTrainingDungeonReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterFungusFighterTrainingDungeonReq_proto_rawDescData)
	})
	return file_EnterFungusFighterTrainingDungeonReq_proto_rawDescData
}

var file_EnterFungusFighterTrainingDungeonReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterFungusFighterTrainingDungeonReq_proto_goTypes = []interface{}{
	(*EnterFungusFighterTrainingDungeonReq)(nil), // 0: proto.EnterFungusFighterTrainingDungeonReq
}
var file_EnterFungusFighterTrainingDungeonReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterFungusFighterTrainingDungeonReq_proto_init() }
func file_EnterFungusFighterTrainingDungeonReq_proto_init() {
	if File_EnterFungusFighterTrainingDungeonReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterFungusFighterTrainingDungeonReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterFungusFighterTrainingDungeonReq); i {
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
			RawDescriptor: file_EnterFungusFighterTrainingDungeonReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterFungusFighterTrainingDungeonReq_proto_goTypes,
		DependencyIndexes: file_EnterFungusFighterTrainingDungeonReq_proto_depIdxs,
		MessageInfos:      file_EnterFungusFighterTrainingDungeonReq_proto_msgTypes,
	}.Build()
	File_EnterFungusFighterTrainingDungeonReq_proto = out.File
	file_EnterFungusFighterTrainingDungeonReq_proto_rawDesc = nil
	file_EnterFungusFighterTrainingDungeonReq_proto_goTypes = nil
	file_EnterFungusFighterTrainingDungeonReq_proto_depIdxs = nil
}
