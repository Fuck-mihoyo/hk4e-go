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
// source: OneofGatherPointDetectorDataNotify.proto

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

// CmdId: 4297
// EnetChannelId: 0
// EnetIsReliable: true
type OneofGatherPointDetectorDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OneofGatherPointDetectorDataList []*OneofGatherPointDetectorData `protobuf:"bytes,3,rep,name=oneof_gather_point_detector_data_list,json=oneofGatherPointDetectorDataList,proto3" json:"oneof_gather_point_detector_data_list,omitempty"`
}

func (x *OneofGatherPointDetectorDataNotify) Reset() {
	*x = OneofGatherPointDetectorDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OneofGatherPointDetectorDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofGatherPointDetectorDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofGatherPointDetectorDataNotify) ProtoMessage() {}

func (x *OneofGatherPointDetectorDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_OneofGatherPointDetectorDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofGatherPointDetectorDataNotify.ProtoReflect.Descriptor instead.
func (*OneofGatherPointDetectorDataNotify) Descriptor() ([]byte, []int) {
	return file_OneofGatherPointDetectorDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *OneofGatherPointDetectorDataNotify) GetOneofGatherPointDetectorDataList() []*OneofGatherPointDetectorData {
	if x != nil {
		return x.OneofGatherPointDetectorDataList
	}
	return nil
}

var File_OneofGatherPointDetectorDataNotify_proto protoreflect.FileDescriptor

var file_OneofGatherPointDetectorDataNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x22, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x74, 0x0a, 0x25,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x20, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OneofGatherPointDetectorDataNotify_proto_rawDescOnce sync.Once
	file_OneofGatherPointDetectorDataNotify_proto_rawDescData = file_OneofGatherPointDetectorDataNotify_proto_rawDesc
)

func file_OneofGatherPointDetectorDataNotify_proto_rawDescGZIP() []byte {
	file_OneofGatherPointDetectorDataNotify_proto_rawDescOnce.Do(func() {
		file_OneofGatherPointDetectorDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_OneofGatherPointDetectorDataNotify_proto_rawDescData)
	})
	return file_OneofGatherPointDetectorDataNotify_proto_rawDescData
}

var file_OneofGatherPointDetectorDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OneofGatherPointDetectorDataNotify_proto_goTypes = []interface{}{
	(*OneofGatherPointDetectorDataNotify)(nil), // 0: proto.OneofGatherPointDetectorDataNotify
	(*OneofGatherPointDetectorData)(nil),       // 1: proto.OneofGatherPointDetectorData
}
var file_OneofGatherPointDetectorDataNotify_proto_depIdxs = []int32{
	1, // 0: proto.OneofGatherPointDetectorDataNotify.oneof_gather_point_detector_data_list:type_name -> proto.OneofGatherPointDetectorData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_OneofGatherPointDetectorDataNotify_proto_init() }
func file_OneofGatherPointDetectorDataNotify_proto_init() {
	if File_OneofGatherPointDetectorDataNotify_proto != nil {
		return
	}
	file_OneofGatherPointDetectorData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OneofGatherPointDetectorDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofGatherPointDetectorDataNotify); i {
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
			RawDescriptor: file_OneofGatherPointDetectorDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OneofGatherPointDetectorDataNotify_proto_goTypes,
		DependencyIndexes: file_OneofGatherPointDetectorDataNotify_proto_depIdxs,
		MessageInfos:      file_OneofGatherPointDetectorDataNotify_proto_msgTypes,
	}.Build()
	File_OneofGatherPointDetectorDataNotify_proto = out.File
	file_OneofGatherPointDetectorDataNotify_proto_rawDesc = nil
	file_OneofGatherPointDetectorDataNotify_proto_goTypes = nil
	file_OneofGatherPointDetectorDataNotify_proto_depIdxs = nil
}
