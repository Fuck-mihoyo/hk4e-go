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
// source: WidgetSlotData.proto

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

type WidgetSlotData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CdOverTime uint32        `protobuf:"varint,9,opt,name=cd_over_time,json=cdOverTime,proto3" json:"cd_over_time,omitempty"`
	Tag        WidgetSlotTag `protobuf:"varint,14,opt,name=tag,proto3,enum=WidgetSlotTag" json:"tag,omitempty"`
	MaterialId uint32        `protobuf:"varint,11,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
	IsActive   bool          `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *WidgetSlotData) Reset() {
	*x = WidgetSlotData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetSlotData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetSlotData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetSlotData) ProtoMessage() {}

func (x *WidgetSlotData) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetSlotData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetSlotData.ProtoReflect.Descriptor instead.
func (*WidgetSlotData) Descriptor() ([]byte, []int) {
	return file_WidgetSlotData_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetSlotData) GetCdOverTime() uint32 {
	if x != nil {
		return x.CdOverTime
	}
	return 0
}

func (x *WidgetSlotData) GetTag() WidgetSlotTag {
	if x != nil {
		return x.Tag
	}
	return WidgetSlotTag_WIDGET_SLOT_TAG_QUICK_USE
}

func (x *WidgetSlotData) GetMaterialId() uint32 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

func (x *WidgetSlotData) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

var File_WidgetSlotData_proto protoreflect.FileDescriptor

var file_WidgetSlotData_proto_rawDesc = []byte{
	0x0a, 0x14, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c,
	0x6f, 0x74, 0x54, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0e,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x0c, 0x63, 0x64, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetSlotData_proto_rawDescOnce sync.Once
	file_WidgetSlotData_proto_rawDescData = file_WidgetSlotData_proto_rawDesc
)

func file_WidgetSlotData_proto_rawDescGZIP() []byte {
	file_WidgetSlotData_proto_rawDescOnce.Do(func() {
		file_WidgetSlotData_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetSlotData_proto_rawDescData)
	})
	return file_WidgetSlotData_proto_rawDescData
}

var file_WidgetSlotData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetSlotData_proto_goTypes = []interface{}{
	(*WidgetSlotData)(nil), // 0: WidgetSlotData
	(WidgetSlotTag)(0),     // 1: WidgetSlotTag
}
var file_WidgetSlotData_proto_depIdxs = []int32{
	1, // 0: WidgetSlotData.tag:type_name -> WidgetSlotTag
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WidgetSlotData_proto_init() }
func file_WidgetSlotData_proto_init() {
	if File_WidgetSlotData_proto != nil {
		return
	}
	file_WidgetSlotTag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetSlotData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetSlotData); i {
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
			RawDescriptor: file_WidgetSlotData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetSlotData_proto_goTypes,
		DependencyIndexes: file_WidgetSlotData_proto_depIdxs,
		MessageInfos:      file_WidgetSlotData_proto_msgTypes,
	}.Build()
	File_WidgetSlotData_proto = out.File
	file_WidgetSlotData_proto_rawDesc = nil
	file_WidgetSlotData_proto_goTypes = nil
	file_WidgetSlotData_proto_depIdxs = nil
}
