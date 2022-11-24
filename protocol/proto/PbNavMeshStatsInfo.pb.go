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
// source: PbNavMeshStatsInfo.proto

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

type PbNavMeshStatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorityAiInCombat   int32 `protobuf:"varint,10,opt,name=authority_ai_in_combat,json=authorityAiInCombat,proto3" json:"authority_ai_in_combat,omitempty"`
	NoAuthorityAiInCombat int32 `protobuf:"varint,11,opt,name=no_authority_ai_in_combat,json=noAuthorityAiInCombat,proto3" json:"no_authority_ai_in_combat,omitempty"`
	TotalAuthorityAi      int32 `protobuf:"varint,8,opt,name=total_authority_ai,json=totalAuthorityAi,proto3" json:"total_authority_ai,omitempty"`
	TotalNoAuthorityAi    int32 `protobuf:"varint,13,opt,name=total_no_authority_ai,json=totalNoAuthorityAi,proto3" json:"total_no_authority_ai,omitempty"`
}

func (x *PbNavMeshStatsInfo) Reset() {
	*x = PbNavMeshStatsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PbNavMeshStatsInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbNavMeshStatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbNavMeshStatsInfo) ProtoMessage() {}

func (x *PbNavMeshStatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PbNavMeshStatsInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbNavMeshStatsInfo.ProtoReflect.Descriptor instead.
func (*PbNavMeshStatsInfo) Descriptor() ([]byte, []int) {
	return file_PbNavMeshStatsInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PbNavMeshStatsInfo) GetAuthorityAiInCombat() int32 {
	if x != nil {
		return x.AuthorityAiInCombat
	}
	return 0
}

func (x *PbNavMeshStatsInfo) GetNoAuthorityAiInCombat() int32 {
	if x != nil {
		return x.NoAuthorityAiInCombat
	}
	return 0
}

func (x *PbNavMeshStatsInfo) GetTotalAuthorityAi() int32 {
	if x != nil {
		return x.TotalAuthorityAi
	}
	return 0
}

func (x *PbNavMeshStatsInfo) GetTotalNoAuthorityAi() int32 {
	if x != nil {
		return x.TotalNoAuthorityAi
	}
	return 0
}

var File_PbNavMeshStatsInfo_proto protoreflect.FileDescriptor

var file_PbNavMeshStatsInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x62, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x50, 0x62, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x69, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x41, 0x69, 0x49, 0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x12, 0x38, 0x0a,
	0x19, 0x6e, 0x6f, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x69,
	0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x6e, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x41, 0x69, 0x49,
	0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x69, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x41, 0x69, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e,
	0x6f, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x69, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x6f, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x41, 0x69, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PbNavMeshStatsInfo_proto_rawDescOnce sync.Once
	file_PbNavMeshStatsInfo_proto_rawDescData = file_PbNavMeshStatsInfo_proto_rawDesc
)

func file_PbNavMeshStatsInfo_proto_rawDescGZIP() []byte {
	file_PbNavMeshStatsInfo_proto_rawDescOnce.Do(func() {
		file_PbNavMeshStatsInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PbNavMeshStatsInfo_proto_rawDescData)
	})
	return file_PbNavMeshStatsInfo_proto_rawDescData
}

var file_PbNavMeshStatsInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PbNavMeshStatsInfo_proto_goTypes = []interface{}{
	(*PbNavMeshStatsInfo)(nil), // 0: proto.PbNavMeshStatsInfo
}
var file_PbNavMeshStatsInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PbNavMeshStatsInfo_proto_init() }
func file_PbNavMeshStatsInfo_proto_init() {
	if File_PbNavMeshStatsInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PbNavMeshStatsInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbNavMeshStatsInfo); i {
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
			RawDescriptor: file_PbNavMeshStatsInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PbNavMeshStatsInfo_proto_goTypes,
		DependencyIndexes: file_PbNavMeshStatsInfo_proto_depIdxs,
		MessageInfos:      file_PbNavMeshStatsInfo_proto_msgTypes,
	}.Build()
	File_PbNavMeshStatsInfo_proto = out.File
	file_PbNavMeshStatsInfo_proto_rawDesc = nil
	file_PbNavMeshStatsInfo_proto_goTypes = nil
	file_PbNavMeshStatsInfo_proto_depIdxs = nil
}
