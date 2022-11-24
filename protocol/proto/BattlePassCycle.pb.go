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
// source: BattlePassCycle.proto

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

type BattlePassCycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CycleIdx  uint32 `protobuf:"varint,3,opt,name=cycle_idx,json=cycleIdx,proto3" json:"cycle_idx,omitempty"`
	EndTime   uint32 `protobuf:"varint,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	BeginTime uint32 `protobuf:"varint,13,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *BattlePassCycle) Reset() {
	*x = BattlePassCycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassCycle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassCycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassCycle) ProtoMessage() {}

func (x *BattlePassCycle) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassCycle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassCycle.ProtoReflect.Descriptor instead.
func (*BattlePassCycle) Descriptor() ([]byte, []int) {
	return file_BattlePassCycle_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassCycle) GetCycleIdx() uint32 {
	if x != nil {
		return x.CycleIdx
	}
	return 0
}

func (x *BattlePassCycle) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *BattlePassCycle) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_BattlePassCycle_proto protoreflect.FileDescriptor

var file_BattlePassCycle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68,
	0x0a, 0x0f, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x78, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassCycle_proto_rawDescOnce sync.Once
	file_BattlePassCycle_proto_rawDescData = file_BattlePassCycle_proto_rawDesc
)

func file_BattlePassCycle_proto_rawDescGZIP() []byte {
	file_BattlePassCycle_proto_rawDescOnce.Do(func() {
		file_BattlePassCycle_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassCycle_proto_rawDescData)
	})
	return file_BattlePassCycle_proto_rawDescData
}

var file_BattlePassCycle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassCycle_proto_goTypes = []interface{}{
	(*BattlePassCycle)(nil), // 0: proto.BattlePassCycle
}
var file_BattlePassCycle_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BattlePassCycle_proto_init() }
func file_BattlePassCycle_proto_init() {
	if File_BattlePassCycle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BattlePassCycle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassCycle); i {
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
			RawDescriptor: file_BattlePassCycle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassCycle_proto_goTypes,
		DependencyIndexes: file_BattlePassCycle_proto_depIdxs,
		MessageInfos:      file_BattlePassCycle_proto_msgTypes,
	}.Build()
	File_BattlePassCycle_proto = out.File
	file_BattlePassCycle_proto_rawDesc = nil
	file_BattlePassCycle_proto_goTypes = nil
	file_BattlePassCycle_proto_depIdxs = nil
}
