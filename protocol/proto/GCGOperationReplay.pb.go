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
// source: GCGOperationReplay.proto

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

type GCGOperationReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId            uint32              `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Seed              uint32              `protobuf:"varint,11,opt,name=seed,proto3" json:"seed,omitempty"`
	OperationDataList []*GCGOperationData `protobuf:"bytes,9,rep,name=operation_data_list,json=operationDataList,proto3" json:"operation_data_list,omitempty"`
}

func (x *GCGOperationReplay) Reset() {
	*x = GCGOperationReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGOperationReplay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGOperationReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGOperationReplay) ProtoMessage() {}

func (x *GCGOperationReplay) ProtoReflect() protoreflect.Message {
	mi := &file_GCGOperationReplay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGOperationReplay.ProtoReflect.Descriptor instead.
func (*GCGOperationReplay) Descriptor() ([]byte, []int) {
	return file_GCGOperationReplay_proto_rawDescGZIP(), []int{0}
}

func (x *GCGOperationReplay) GetGameId() uint32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *GCGOperationReplay) GetSeed() uint32 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *GCGOperationReplay) GetOperationDataList() []*GCGOperationData {
	if x != nil {
		return x.OperationDataList
	}
	return nil
}

var File_GCGOperationReplay_proto protoreflect.FileDescriptor

var file_GCGOperationReplay_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x43, 0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x47, 0x43, 0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x47, 0x43,
	0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x47, 0x0a,
	0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGOperationReplay_proto_rawDescOnce sync.Once
	file_GCGOperationReplay_proto_rawDescData = file_GCGOperationReplay_proto_rawDesc
)

func file_GCGOperationReplay_proto_rawDescGZIP() []byte {
	file_GCGOperationReplay_proto_rawDescOnce.Do(func() {
		file_GCGOperationReplay_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGOperationReplay_proto_rawDescData)
	})
	return file_GCGOperationReplay_proto_rawDescData
}

var file_GCGOperationReplay_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGOperationReplay_proto_goTypes = []interface{}{
	(*GCGOperationReplay)(nil), // 0: proto.GCGOperationReplay
	(*GCGOperationData)(nil),   // 1: proto.GCGOperationData
}
var file_GCGOperationReplay_proto_depIdxs = []int32{
	1, // 0: proto.GCGOperationReplay.operation_data_list:type_name -> proto.GCGOperationData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGOperationReplay_proto_init() }
func file_GCGOperationReplay_proto_init() {
	if File_GCGOperationReplay_proto != nil {
		return
	}
	file_GCGOperationData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGOperationReplay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGOperationReplay); i {
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
			RawDescriptor: file_GCGOperationReplay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGOperationReplay_proto_goTypes,
		DependencyIndexes: file_GCGOperationReplay_proto_depIdxs,
		MessageInfos:      file_GCGOperationReplay_proto_msgTypes,
	}.Build()
	File_GCGOperationReplay_proto = out.File
	file_GCGOperationReplay_proto_rawDesc = nil
	file_GCGOperationReplay_proto_goTypes = nil
	file_GCGOperationReplay_proto_depIdxs = nil
}
