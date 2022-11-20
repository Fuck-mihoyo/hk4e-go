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
// source: WidgetCreatorInfo.proto

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

type WidgetCreatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType       WidgetCreatorOpType       `protobuf:"varint,10,opt,name=op_type,json=opType,proto3,enum=WidgetCreatorOpType" json:"op_type,omitempty"`
	EntityId     uint32                    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	LocationInfo *WidgetCreateLocationInfo `protobuf:"bytes,12,opt,name=location_info,json=locationInfo,proto3" json:"location_info,omitempty"`
}

func (x *WidgetCreatorInfo) Reset() {
	*x = WidgetCreatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetCreatorInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetCreatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetCreatorInfo) ProtoMessage() {}

func (x *WidgetCreatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetCreatorInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetCreatorInfo.ProtoReflect.Descriptor instead.
func (*WidgetCreatorInfo) Descriptor() ([]byte, []int) {
	return file_WidgetCreatorInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetCreatorInfo) GetOpType() WidgetCreatorOpType {
	if x != nil {
		return x.OpType
	}
	return WidgetCreatorOpType_WIDGET_CREATOR_OP_TYPE_NONE
}

func (x *WidgetCreatorInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *WidgetCreatorInfo) GetLocationInfo() *WidgetCreateLocationInfo {
	if x != nil {
		return x.LocationInfo
	}
	return nil
}

var File_WidgetCreatorInfo_proto protoreflect.FileDescriptor

var file_WidgetCreatorInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x07, 0x6f, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetCreatorInfo_proto_rawDescOnce sync.Once
	file_WidgetCreatorInfo_proto_rawDescData = file_WidgetCreatorInfo_proto_rawDesc
)

func file_WidgetCreatorInfo_proto_rawDescGZIP() []byte {
	file_WidgetCreatorInfo_proto_rawDescOnce.Do(func() {
		file_WidgetCreatorInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetCreatorInfo_proto_rawDescData)
	})
	return file_WidgetCreatorInfo_proto_rawDescData
}

var file_WidgetCreatorInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetCreatorInfo_proto_goTypes = []interface{}{
	(*WidgetCreatorInfo)(nil),        // 0: WidgetCreatorInfo
	(WidgetCreatorOpType)(0),         // 1: WidgetCreatorOpType
	(*WidgetCreateLocationInfo)(nil), // 2: WidgetCreateLocationInfo
}
var file_WidgetCreatorInfo_proto_depIdxs = []int32{
	1, // 0: WidgetCreatorInfo.op_type:type_name -> WidgetCreatorOpType
	2, // 1: WidgetCreatorInfo.location_info:type_name -> WidgetCreateLocationInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WidgetCreatorInfo_proto_init() }
func file_WidgetCreatorInfo_proto_init() {
	if File_WidgetCreatorInfo_proto != nil {
		return
	}
	file_WidgetCreateLocationInfo_proto_init()
	file_WidgetCreatorOpType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetCreatorInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetCreatorInfo); i {
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
			RawDescriptor: file_WidgetCreatorInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetCreatorInfo_proto_goTypes,
		DependencyIndexes: file_WidgetCreatorInfo_proto_depIdxs,
		MessageInfos:      file_WidgetCreatorInfo_proto_msgTypes,
	}.Build()
	File_WidgetCreatorInfo_proto = out.File
	file_WidgetCreatorInfo_proto_rawDesc = nil
	file_WidgetCreatorInfo_proto_goTypes = nil
	file_WidgetCreatorInfo_proto_depIdxs = nil
}
