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
// source: GCGStartChallengeReq.proto

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

// CmdId: 7595
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GCGStartChallengeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelType GCGLevelType `protobuf:"varint,5,opt,name=level_type,json=levelType,proto3,enum=proto.GCGLevelType" json:"level_type,omitempty"`
	ConfigId  uint32       `protobuf:"varint,13,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	LevelId   uint32       `protobuf:"varint,12,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *GCGStartChallengeReq) Reset() {
	*x = GCGStartChallengeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGStartChallengeReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGStartChallengeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGStartChallengeReq) ProtoMessage() {}

func (x *GCGStartChallengeReq) ProtoReflect() protoreflect.Message {
	mi := &file_GCGStartChallengeReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGStartChallengeReq.ProtoReflect.Descriptor instead.
func (*GCGStartChallengeReq) Descriptor() ([]byte, []int) {
	return file_GCGStartChallengeReq_proto_rawDescGZIP(), []int{0}
}

func (x *GCGStartChallengeReq) GetLevelType() GCGLevelType {
	if x != nil {
		return x.LevelType
	}
	return GCGLevelType_GCG_LEVEL_TYPE_NONE
}

func (x *GCGStartChallengeReq) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *GCGStartChallengeReq) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_GCGStartChallengeReq_proto protoreflect.FileDescriptor

var file_GCGStartChallengeReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x43, 0x47, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x43, 0x47, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x32, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGStartChallengeReq_proto_rawDescOnce sync.Once
	file_GCGStartChallengeReq_proto_rawDescData = file_GCGStartChallengeReq_proto_rawDesc
)

func file_GCGStartChallengeReq_proto_rawDescGZIP() []byte {
	file_GCGStartChallengeReq_proto_rawDescOnce.Do(func() {
		file_GCGStartChallengeReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGStartChallengeReq_proto_rawDescData)
	})
	return file_GCGStartChallengeReq_proto_rawDescData
}

var file_GCGStartChallengeReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGStartChallengeReq_proto_goTypes = []interface{}{
	(*GCGStartChallengeReq)(nil), // 0: proto.GCGStartChallengeReq
	(GCGLevelType)(0),            // 1: proto.GCGLevelType
}
var file_GCGStartChallengeReq_proto_depIdxs = []int32{
	1, // 0: proto.GCGStartChallengeReq.level_type:type_name -> proto.GCGLevelType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGStartChallengeReq_proto_init() }
func file_GCGStartChallengeReq_proto_init() {
	if File_GCGStartChallengeReq_proto != nil {
		return
	}
	file_GCGLevelType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGStartChallengeReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGStartChallengeReq); i {
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
			RawDescriptor: file_GCGStartChallengeReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGStartChallengeReq_proto_goTypes,
		DependencyIndexes: file_GCGStartChallengeReq_proto_depIdxs,
		MessageInfos:      file_GCGStartChallengeReq_proto_msgTypes,
	}.Build()
	File_GCGStartChallengeReq_proto = out.File
	file_GCGStartChallengeReq_proto_rawDesc = nil
	file_GCGStartChallengeReq_proto_goTypes = nil
	file_GCGStartChallengeReq_proto_depIdxs = nil
}
