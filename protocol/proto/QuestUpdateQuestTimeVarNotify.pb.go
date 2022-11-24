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
// source: QuestUpdateQuestTimeVarNotify.proto

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

// CmdId: 456
// EnetChannelId: 0
// EnetIsReliable: true
type QuestUpdateQuestTimeVarNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeVarMap    map[uint32]uint32 `protobuf:"bytes,1,rep,name=time_var_map,json=timeVarMap,proto3" json:"time_var_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ParentQuestId uint32            `protobuf:"varint,3,opt,name=parent_quest_id,json=parentQuestId,proto3" json:"parent_quest_id,omitempty"`
}

func (x *QuestUpdateQuestTimeVarNotify) Reset() {
	*x = QuestUpdateQuestTimeVarNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestUpdateQuestTimeVarNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestUpdateQuestTimeVarNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestUpdateQuestTimeVarNotify) ProtoMessage() {}

func (x *QuestUpdateQuestTimeVarNotify) ProtoReflect() protoreflect.Message {
	mi := &file_QuestUpdateQuestTimeVarNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestUpdateQuestTimeVarNotify.ProtoReflect.Descriptor instead.
func (*QuestUpdateQuestTimeVarNotify) Descriptor() ([]byte, []int) {
	return file_QuestUpdateQuestTimeVarNotify_proto_rawDescGZIP(), []int{0}
}

func (x *QuestUpdateQuestTimeVarNotify) GetTimeVarMap() map[uint32]uint32 {
	if x != nil {
		return x.TimeVarMap
	}
	return nil
}

func (x *QuestUpdateQuestTimeVarNotify) GetParentQuestId() uint32 {
	if x != nil {
		return x.ParentQuestId
	}
	return 0
}

var File_QuestUpdateQuestTimeVarNotify_proto protoreflect.FileDescriptor

var file_QuestUpdateQuestTimeVarNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a,
	0x1d, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x56,
	0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x56, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56,
	0x61, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x3d,
	0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_QuestUpdateQuestTimeVarNotify_proto_rawDescOnce sync.Once
	file_QuestUpdateQuestTimeVarNotify_proto_rawDescData = file_QuestUpdateQuestTimeVarNotify_proto_rawDesc
)

func file_QuestUpdateQuestTimeVarNotify_proto_rawDescGZIP() []byte {
	file_QuestUpdateQuestTimeVarNotify_proto_rawDescOnce.Do(func() {
		file_QuestUpdateQuestTimeVarNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestUpdateQuestTimeVarNotify_proto_rawDescData)
	})
	return file_QuestUpdateQuestTimeVarNotify_proto_rawDescData
}

var file_QuestUpdateQuestTimeVarNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_QuestUpdateQuestTimeVarNotify_proto_goTypes = []interface{}{
	(*QuestUpdateQuestTimeVarNotify)(nil), // 0: proto.QuestUpdateQuestTimeVarNotify
	nil,                                   // 1: proto.QuestUpdateQuestTimeVarNotify.TimeVarMapEntry
}
var file_QuestUpdateQuestTimeVarNotify_proto_depIdxs = []int32{
	1, // 0: proto.QuestUpdateQuestTimeVarNotify.time_var_map:type_name -> proto.QuestUpdateQuestTimeVarNotify.TimeVarMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_QuestUpdateQuestTimeVarNotify_proto_init() }
func file_QuestUpdateQuestTimeVarNotify_proto_init() {
	if File_QuestUpdateQuestTimeVarNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestUpdateQuestTimeVarNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestUpdateQuestTimeVarNotify); i {
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
			RawDescriptor: file_QuestUpdateQuestTimeVarNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestUpdateQuestTimeVarNotify_proto_goTypes,
		DependencyIndexes: file_QuestUpdateQuestTimeVarNotify_proto_depIdxs,
		MessageInfos:      file_QuestUpdateQuestTimeVarNotify_proto_msgTypes,
	}.Build()
	File_QuestUpdateQuestTimeVarNotify_proto = out.File
	file_QuestUpdateQuestTimeVarNotify_proto_rawDesc = nil
	file_QuestUpdateQuestTimeVarNotify_proto_goTypes = nil
	file_QuestUpdateQuestTimeVarNotify_proto_depIdxs = nil
}
