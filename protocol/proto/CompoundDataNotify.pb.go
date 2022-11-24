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
// source: CompoundDataNotify.proto

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

// CmdId: 146
// EnetChannelId: 0
// EnetIsReliable: true
type CompoundDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockCompoundList  []uint32             `protobuf:"varint,1,rep,packed,name=unlock_compound_list,json=unlockCompoundList,proto3" json:"unlock_compound_list,omitempty"`
	CompoundQueDataList []*CompoundQueueData `protobuf:"bytes,9,rep,name=compound_que_data_list,json=compoundQueDataList,proto3" json:"compound_que_data_list,omitempty"`
}

func (x *CompoundDataNotify) Reset() {
	*x = CompoundDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CompoundDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompoundDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompoundDataNotify) ProtoMessage() {}

func (x *CompoundDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CompoundDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompoundDataNotify.ProtoReflect.Descriptor instead.
func (*CompoundDataNotify) Descriptor() ([]byte, []int) {
	return file_CompoundDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CompoundDataNotify) GetUnlockCompoundList() []uint32 {
	if x != nil {
		return x.UnlockCompoundList
	}
	return nil
}

func (x *CompoundDataNotify) GetCompoundQueDataList() []*CompoundQueueData {
	if x != nil {
		return x.CompoundQueDataList
	}
	return nil
}

var File_CompoundDataNotify_proto protoreflect.FileDescriptor

var file_CompoundDataNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x71, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x13, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CompoundDataNotify_proto_rawDescOnce sync.Once
	file_CompoundDataNotify_proto_rawDescData = file_CompoundDataNotify_proto_rawDesc
)

func file_CompoundDataNotify_proto_rawDescGZIP() []byte {
	file_CompoundDataNotify_proto_rawDescOnce.Do(func() {
		file_CompoundDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CompoundDataNotify_proto_rawDescData)
	})
	return file_CompoundDataNotify_proto_rawDescData
}

var file_CompoundDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CompoundDataNotify_proto_goTypes = []interface{}{
	(*CompoundDataNotify)(nil), // 0: proto.CompoundDataNotify
	(*CompoundQueueData)(nil),  // 1: proto.CompoundQueueData
}
var file_CompoundDataNotify_proto_depIdxs = []int32{
	1, // 0: proto.CompoundDataNotify.compound_que_data_list:type_name -> proto.CompoundQueueData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CompoundDataNotify_proto_init() }
func file_CompoundDataNotify_proto_init() {
	if File_CompoundDataNotify_proto != nil {
		return
	}
	file_CompoundQueueData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CompoundDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompoundDataNotify); i {
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
			RawDescriptor: file_CompoundDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CompoundDataNotify_proto_goTypes,
		DependencyIndexes: file_CompoundDataNotify_proto_depIdxs,
		MessageInfos:      file_CompoundDataNotify_proto_msgTypes,
	}.Build()
	File_CompoundDataNotify_proto = out.File
	file_CompoundDataNotify_proto_rawDesc = nil
	file_CompoundDataNotify_proto_goTypes = nil
	file_CompoundDataNotify_proto_depIdxs = nil
}
