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
// source: GetMapAreaReq.proto

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

// CmdId: 3108
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GetMapAreaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMapAreaReq) Reset() {
	*x = GetMapAreaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMapAreaReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapAreaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapAreaReq) ProtoMessage() {}

func (x *GetMapAreaReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetMapAreaReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapAreaReq.ProtoReflect.Descriptor instead.
func (*GetMapAreaReq) Descriptor() ([]byte, []int) {
	return file_GetMapAreaReq_proto_rawDescGZIP(), []int{0}
}

var File_GetMapAreaReq_proto protoreflect.FileDescriptor

var file_GetMapAreaReq_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GetMapAreaReq_proto_rawDescOnce sync.Once
	file_GetMapAreaReq_proto_rawDescData = file_GetMapAreaReq_proto_rawDesc
)

func file_GetMapAreaReq_proto_rawDescGZIP() []byte {
	file_GetMapAreaReq_proto_rawDescOnce.Do(func() {
		file_GetMapAreaReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMapAreaReq_proto_rawDescData)
	})
	return file_GetMapAreaReq_proto_rawDescData
}

var file_GetMapAreaReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMapAreaReq_proto_goTypes = []interface{}{
	(*GetMapAreaReq)(nil), // 0: proto.GetMapAreaReq
}
var file_GetMapAreaReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetMapAreaReq_proto_init() }
func file_GetMapAreaReq_proto_init() {
	if File_GetMapAreaReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetMapAreaReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapAreaReq); i {
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
			RawDescriptor: file_GetMapAreaReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMapAreaReq_proto_goTypes,
		DependencyIndexes: file_GetMapAreaReq_proto_depIdxs,
		MessageInfos:      file_GetMapAreaReq_proto_msgTypes,
	}.Build()
	File_GetMapAreaReq_proto = out.File
	file_GetMapAreaReq_proto_rawDesc = nil
	file_GetMapAreaReq_proto_goTypes = nil
	file_GetMapAreaReq_proto_depIdxs = nil
}
