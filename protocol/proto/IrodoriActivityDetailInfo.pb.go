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
// source: IrodoriActivityDetailInfo.proto

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

type IrodoriActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterLevelList []*IrodoriMasterLevelInfo `protobuf:"bytes,11,rep,name=master_level_list,json=masterLevelList,proto3" json:"master_level_list,omitempty"`
	FlowerData      *IrodoriFlowerData        `protobuf:"bytes,6,opt,name=flower_data,json=flowerData,proto3" json:"flower_data,omitempty"`
	PoetryData      *IrodoriPoetryData        `protobuf:"bytes,8,opt,name=poetry_data,json=poetryData,proto3" json:"poetry_data,omitempty"`
	ChessData       *IrodoriChessData         `protobuf:"bytes,14,opt,name=chess_data,json=chessData,proto3" json:"chess_data,omitempty"`
}

func (x *IrodoriActivityDetailInfo) Reset() {
	*x = IrodoriActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IrodoriActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IrodoriActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IrodoriActivityDetailInfo) ProtoMessage() {}

func (x *IrodoriActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_IrodoriActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IrodoriActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*IrodoriActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_IrodoriActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *IrodoriActivityDetailInfo) GetMasterLevelList() []*IrodoriMasterLevelInfo {
	if x != nil {
		return x.MasterLevelList
	}
	return nil
}

func (x *IrodoriActivityDetailInfo) GetFlowerData() *IrodoriFlowerData {
	if x != nil {
		return x.FlowerData
	}
	return nil
}

func (x *IrodoriActivityDetailInfo) GetPoetryData() *IrodoriPoetryData {
	if x != nil {
		return x.PoetryData
	}
	return nil
}

func (x *IrodoriActivityDetailInfo) GetChessData() *IrodoriChessData {
	if x != nil {
		return x.ChessData
	}
	return nil
}

var File_IrodoriActivityDetailInfo_proto protoreflect.FileDescriptor

var file_IrodoriActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72,
	0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x49, 0x72, 0x6f, 0x64, 0x6f,
	0x72, 0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69,
	0x50, 0x6f, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x02, 0x0a, 0x19, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49,
	0x0a, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x66, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x46, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0b, 0x70, 0x6f, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x50, 0x6f, 0x65, 0x74, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x6f, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x36, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x72, 0x6f, 0x64,
	0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IrodoriActivityDetailInfo_proto_rawDescOnce sync.Once
	file_IrodoriActivityDetailInfo_proto_rawDescData = file_IrodoriActivityDetailInfo_proto_rawDesc
)

func file_IrodoriActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_IrodoriActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_IrodoriActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_IrodoriActivityDetailInfo_proto_rawDescData)
	})
	return file_IrodoriActivityDetailInfo_proto_rawDescData
}

var file_IrodoriActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IrodoriActivityDetailInfo_proto_goTypes = []interface{}{
	(*IrodoriActivityDetailInfo)(nil), // 0: proto.IrodoriActivityDetailInfo
	(*IrodoriMasterLevelInfo)(nil),    // 1: proto.IrodoriMasterLevelInfo
	(*IrodoriFlowerData)(nil),         // 2: proto.IrodoriFlowerData
	(*IrodoriPoetryData)(nil),         // 3: proto.IrodoriPoetryData
	(*IrodoriChessData)(nil),          // 4: proto.IrodoriChessData
}
var file_IrodoriActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: proto.IrodoriActivityDetailInfo.master_level_list:type_name -> proto.IrodoriMasterLevelInfo
	2, // 1: proto.IrodoriActivityDetailInfo.flower_data:type_name -> proto.IrodoriFlowerData
	3, // 2: proto.IrodoriActivityDetailInfo.poetry_data:type_name -> proto.IrodoriPoetryData
	4, // 3: proto.IrodoriActivityDetailInfo.chess_data:type_name -> proto.IrodoriChessData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_IrodoriActivityDetailInfo_proto_init() }
func file_IrodoriActivityDetailInfo_proto_init() {
	if File_IrodoriActivityDetailInfo_proto != nil {
		return
	}
	file_IrodoriChessData_proto_init()
	file_IrodoriFlowerData_proto_init()
	file_IrodoriMasterLevelInfo_proto_init()
	file_IrodoriPoetryData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IrodoriActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IrodoriActivityDetailInfo); i {
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
			RawDescriptor: file_IrodoriActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IrodoriActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_IrodoriActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_IrodoriActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_IrodoriActivityDetailInfo_proto = out.File
	file_IrodoriActivityDetailInfo_proto_rawDesc = nil
	file_IrodoriActivityDetailInfo_proto_goTypes = nil
	file_IrodoriActivityDetailInfo_proto_depIdxs = nil
}