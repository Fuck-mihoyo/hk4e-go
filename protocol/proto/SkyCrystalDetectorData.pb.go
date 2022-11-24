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
// source: SkyCrystalDetectorData.proto

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

type SkyCrystalDetectorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHintValid   bool    `protobuf:"varint,3,opt,name=is_hint_valid,json=isHintValid,proto3" json:"is_hint_valid,omitempty"`
	HintCenterPos *Vector `protobuf:"bytes,8,opt,name=hint_center_pos,json=hintCenterPos,proto3" json:"hint_center_pos,omitempty"`
	GroupId       uint32  `protobuf:"varint,6,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ConfigId      uint32  `protobuf:"varint,9,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *SkyCrystalDetectorData) Reset() {
	*x = SkyCrystalDetectorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SkyCrystalDetectorData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkyCrystalDetectorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkyCrystalDetectorData) ProtoMessage() {}

func (x *SkyCrystalDetectorData) ProtoReflect() protoreflect.Message {
	mi := &file_SkyCrystalDetectorData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkyCrystalDetectorData.ProtoReflect.Descriptor instead.
func (*SkyCrystalDetectorData) Descriptor() ([]byte, []int) {
	return file_SkyCrystalDetectorData_proto_rawDescGZIP(), []int{0}
}

func (x *SkyCrystalDetectorData) GetIsHintValid() bool {
	if x != nil {
		return x.IsHintValid
	}
	return false
}

func (x *SkyCrystalDetectorData) GetHintCenterPos() *Vector {
	if x != nil {
		return x.HintCenterPos
	}
	return nil
}

func (x *SkyCrystalDetectorData) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SkyCrystalDetectorData) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

var File_SkyCrystalDetectorData_proto protoreflect.FileDescriptor

var file_SkyCrystalDetectorData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x53, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73, 0x74,
	0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x48, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x35, 0x0a, 0x0f, 0x68, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x68, 0x69, 0x6e, 0x74,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SkyCrystalDetectorData_proto_rawDescOnce sync.Once
	file_SkyCrystalDetectorData_proto_rawDescData = file_SkyCrystalDetectorData_proto_rawDesc
)

func file_SkyCrystalDetectorData_proto_rawDescGZIP() []byte {
	file_SkyCrystalDetectorData_proto_rawDescOnce.Do(func() {
		file_SkyCrystalDetectorData_proto_rawDescData = protoimpl.X.CompressGZIP(file_SkyCrystalDetectorData_proto_rawDescData)
	})
	return file_SkyCrystalDetectorData_proto_rawDescData
}

var file_SkyCrystalDetectorData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SkyCrystalDetectorData_proto_goTypes = []interface{}{
	(*SkyCrystalDetectorData)(nil), // 0: proto.SkyCrystalDetectorData
	(*Vector)(nil),                 // 1: proto.Vector
}
var file_SkyCrystalDetectorData_proto_depIdxs = []int32{
	1, // 0: proto.SkyCrystalDetectorData.hint_center_pos:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SkyCrystalDetectorData_proto_init() }
func file_SkyCrystalDetectorData_proto_init() {
	if File_SkyCrystalDetectorData_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SkyCrystalDetectorData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkyCrystalDetectorData); i {
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
			RawDescriptor: file_SkyCrystalDetectorData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SkyCrystalDetectorData_proto_goTypes,
		DependencyIndexes: file_SkyCrystalDetectorData_proto_depIdxs,
		MessageInfos:      file_SkyCrystalDetectorData_proto_msgTypes,
	}.Build()
	File_SkyCrystalDetectorData_proto = out.File
	file_SkyCrystalDetectorData_proto_rawDesc = nil
	file_SkyCrystalDetectorData_proto_goTypes = nil
	file_SkyCrystalDetectorData_proto_depIdxs = nil
}
