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
// 	protoc        v3.7.0
// source: SceneEntityMoveReq.proto

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

// CmdId: 290
// EnetChannelId: 1
// EnetIsReliable: false
// IsAllowClient: true
type SceneEntityMoveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MotionInfo  *MotionInfo `protobuf:"bytes,7,opt,name=motion_info,json=motionInfo,proto3" json:"motion_info,omitempty"`
	SceneTime   uint32      `protobuf:"varint,4,opt,name=scene_time,json=sceneTime,proto3" json:"scene_time,omitempty"`
	EntityId    uint32      `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ReliableSeq uint32      `protobuf:"varint,15,opt,name=reliable_seq,json=reliableSeq,proto3" json:"reliable_seq,omitempty"`
}

func (x *SceneEntityMoveReq) Reset() {
	*x = SceneEntityMoveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityMoveReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityMoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityMoveReq) ProtoMessage() {}

func (x *SceneEntityMoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityMoveReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityMoveReq.ProtoReflect.Descriptor instead.
func (*SceneEntityMoveReq) Descriptor() ([]byte, []int) {
	return file_SceneEntityMoveReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityMoveReq) GetMotionInfo() *MotionInfo {
	if x != nil {
		return x.MotionInfo
	}
	return nil
}

func (x *SceneEntityMoveReq) GetSceneTime() uint32 {
	if x != nil {
		return x.SceneTime
	}
	return 0
}

func (x *SceneEntityMoveReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneEntityMoveReq) GetReliableSeq() uint32 {
	if x != nil {
		return x.ReliableSeq
	}
	return 0
}

var File_SceneEntityMoveReq_proto protoreflect.FileDescriptor

var file_SceneEntityMoveReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a,
	0x12, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x0b, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x71,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityMoveReq_proto_rawDescOnce sync.Once
	file_SceneEntityMoveReq_proto_rawDescData = file_SceneEntityMoveReq_proto_rawDesc
)

func file_SceneEntityMoveReq_proto_rawDescGZIP() []byte {
	file_SceneEntityMoveReq_proto_rawDescOnce.Do(func() {
		file_SceneEntityMoveReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityMoveReq_proto_rawDescData)
	})
	return file_SceneEntityMoveReq_proto_rawDescData
}

var file_SceneEntityMoveReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityMoveReq_proto_goTypes = []interface{}{
	(*SceneEntityMoveReq)(nil), // 0: SceneEntityMoveReq
	(*MotionInfo)(nil),         // 1: MotionInfo
}
var file_SceneEntityMoveReq_proto_depIdxs = []int32{
	1, // 0: SceneEntityMoveReq.motion_info:type_name -> MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEntityMoveReq_proto_init() }
func file_SceneEntityMoveReq_proto_init() {
	if File_SceneEntityMoveReq_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityMoveReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityMoveReq); i {
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
			RawDescriptor: file_SceneEntityMoveReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityMoveReq_proto_goTypes,
		DependencyIndexes: file_SceneEntityMoveReq_proto_depIdxs,
		MessageInfos:      file_SceneEntityMoveReq_proto_msgTypes,
	}.Build()
	File_SceneEntityMoveReq_proto = out.File
	file_SceneEntityMoveReq_proto_rawDesc = nil
	file_SceneEntityMoveReq_proto_goTypes = nil
	file_SceneEntityMoveReq_proto_depIdxs = nil
}
