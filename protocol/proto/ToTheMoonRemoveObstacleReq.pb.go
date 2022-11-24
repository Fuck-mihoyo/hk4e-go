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
// source: ToTheMoonRemoveObstacleReq.proto

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

type ToTheMoonRemoveObstacleReq_ObstacleType int32

const (
	ToTheMoonRemoveObstacleReq_OBSTACLE_TYPE_BOX     ToTheMoonRemoveObstacleReq_ObstacleType = 0
	ToTheMoonRemoveObstacleReq_OBSTACLE_TYPE_CAPSULE ToTheMoonRemoveObstacleReq_ObstacleType = 1
)

// Enum value maps for ToTheMoonRemoveObstacleReq_ObstacleType.
var (
	ToTheMoonRemoveObstacleReq_ObstacleType_name = map[int32]string{
		0: "OBSTACLE_TYPE_BOX",
		1: "OBSTACLE_TYPE_CAPSULE",
	}
	ToTheMoonRemoveObstacleReq_ObstacleType_value = map[string]int32{
		"OBSTACLE_TYPE_BOX":     0,
		"OBSTACLE_TYPE_CAPSULE": 1,
	}
)

func (x ToTheMoonRemoveObstacleReq_ObstacleType) Enum() *ToTheMoonRemoveObstacleReq_ObstacleType {
	p := new(ToTheMoonRemoveObstacleReq_ObstacleType)
	*p = x
	return p
}

func (x ToTheMoonRemoveObstacleReq_ObstacleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonRemoveObstacleReq_ObstacleType) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonRemoveObstacleReq_proto_enumTypes[0].Descriptor()
}

func (ToTheMoonRemoveObstacleReq_ObstacleType) Type() protoreflect.EnumType {
	return &file_ToTheMoonRemoveObstacleReq_proto_enumTypes[0]
}

func (x ToTheMoonRemoveObstacleReq_ObstacleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonRemoveObstacleReq_ObstacleType.Descriptor instead.
func (ToTheMoonRemoveObstacleReq_ObstacleType) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonRemoveObstacleReq_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 6190
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ToTheMoonRemoveObstacleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handle  int32  `protobuf:"varint,12,opt,name=handle,proto3" json:"handle,omitempty"`
	SceneId uint32 `protobuf:"varint,10,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	QueryId int32  `protobuf:"varint,11,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *ToTheMoonRemoveObstacleReq) Reset() {
	*x = ToTheMoonRemoveObstacleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToTheMoonRemoveObstacleReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToTheMoonRemoveObstacleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToTheMoonRemoveObstacleReq) ProtoMessage() {}

func (x *ToTheMoonRemoveObstacleReq) ProtoReflect() protoreflect.Message {
	mi := &file_ToTheMoonRemoveObstacleReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToTheMoonRemoveObstacleReq.ProtoReflect.Descriptor instead.
func (*ToTheMoonRemoveObstacleReq) Descriptor() ([]byte, []int) {
	return file_ToTheMoonRemoveObstacleReq_proto_rawDescGZIP(), []int{0}
}

func (x *ToTheMoonRemoveObstacleReq) GetHandle() int32 {
	if x != nil {
		return x.Handle
	}
	return 0
}

func (x *ToTheMoonRemoveObstacleReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *ToTheMoonRemoveObstacleReq) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

var File_ToTheMoonRemoveObstacleReq_proto protoreflect.FileDescriptor

var file_ToTheMoonRemoveObstacleReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x1a, 0x54, 0x6f,
	0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x62, 0x73,
	0x74, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x0c, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x42, 0x53, 0x54, 0x41, 0x43,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x58, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x4f, 0x42, 0x53, 0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x41, 0x50, 0x53, 0x55, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ToTheMoonRemoveObstacleReq_proto_rawDescOnce sync.Once
	file_ToTheMoonRemoveObstacleReq_proto_rawDescData = file_ToTheMoonRemoveObstacleReq_proto_rawDesc
)

func file_ToTheMoonRemoveObstacleReq_proto_rawDescGZIP() []byte {
	file_ToTheMoonRemoveObstacleReq_proto_rawDescOnce.Do(func() {
		file_ToTheMoonRemoveObstacleReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ToTheMoonRemoveObstacleReq_proto_rawDescData)
	})
	return file_ToTheMoonRemoveObstacleReq_proto_rawDescData
}

var file_ToTheMoonRemoveObstacleReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ToTheMoonRemoveObstacleReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ToTheMoonRemoveObstacleReq_proto_goTypes = []interface{}{
	(ToTheMoonRemoveObstacleReq_ObstacleType)(0), // 0: proto.ToTheMoonRemoveObstacleReq.ObstacleType
	(*ToTheMoonRemoveObstacleReq)(nil),           // 1: proto.ToTheMoonRemoveObstacleReq
}
var file_ToTheMoonRemoveObstacleReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ToTheMoonRemoveObstacleReq_proto_init() }
func file_ToTheMoonRemoveObstacleReq_proto_init() {
	if File_ToTheMoonRemoveObstacleReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ToTheMoonRemoveObstacleReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToTheMoonRemoveObstacleReq); i {
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
			RawDescriptor: file_ToTheMoonRemoveObstacleReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ToTheMoonRemoveObstacleReq_proto_goTypes,
		DependencyIndexes: file_ToTheMoonRemoveObstacleReq_proto_depIdxs,
		EnumInfos:         file_ToTheMoonRemoveObstacleReq_proto_enumTypes,
		MessageInfos:      file_ToTheMoonRemoveObstacleReq_proto_msgTypes,
	}.Build()
	File_ToTheMoonRemoveObstacleReq_proto = out.File
	file_ToTheMoonRemoveObstacleReq_proto_rawDesc = nil
	file_ToTheMoonRemoveObstacleReq_proto_goTypes = nil
	file_ToTheMoonRemoveObstacleReq_proto_depIdxs = nil
}
