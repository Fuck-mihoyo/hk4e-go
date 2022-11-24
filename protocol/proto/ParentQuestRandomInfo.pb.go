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
// source: ParentQuestRandomInfo.proto

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

type ParentQuestRandomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FactorList []uint32 `protobuf:"varint,1,rep,packed,name=factor_list,json=factorList,proto3" json:"factor_list,omitempty"`
	TemplateId uint32   `protobuf:"varint,8,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	EntranceId uint32   `protobuf:"varint,2,opt,name=entrance_id,json=entranceId,proto3" json:"entrance_id,omitempty"`
}

func (x *ParentQuestRandomInfo) Reset() {
	*x = ParentQuestRandomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ParentQuestRandomInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParentQuestRandomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParentQuestRandomInfo) ProtoMessage() {}

func (x *ParentQuestRandomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ParentQuestRandomInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParentQuestRandomInfo.ProtoReflect.Descriptor instead.
func (*ParentQuestRandomInfo) Descriptor() ([]byte, []int) {
	return file_ParentQuestRandomInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ParentQuestRandomInfo) GetFactorList() []uint32 {
	if x != nil {
		return x.FactorList
	}
	return nil
}

func (x *ParentQuestRandomInfo) GetTemplateId() uint32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *ParentQuestRandomInfo) GetEntranceId() uint32 {
	if x != nil {
		return x.EntranceId
	}
	return 0
}

var File_ParentQuestRandomInfo_proto protoreflect.FileDescriptor

var file_ParentQuestRandomInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x15, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ParentQuestRandomInfo_proto_rawDescOnce sync.Once
	file_ParentQuestRandomInfo_proto_rawDescData = file_ParentQuestRandomInfo_proto_rawDesc
)

func file_ParentQuestRandomInfo_proto_rawDescGZIP() []byte {
	file_ParentQuestRandomInfo_proto_rawDescOnce.Do(func() {
		file_ParentQuestRandomInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ParentQuestRandomInfo_proto_rawDescData)
	})
	return file_ParentQuestRandomInfo_proto_rawDescData
}

var file_ParentQuestRandomInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ParentQuestRandomInfo_proto_goTypes = []interface{}{
	(*ParentQuestRandomInfo)(nil), // 0: proto.ParentQuestRandomInfo
}
var file_ParentQuestRandomInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ParentQuestRandomInfo_proto_init() }
func file_ParentQuestRandomInfo_proto_init() {
	if File_ParentQuestRandomInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ParentQuestRandomInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParentQuestRandomInfo); i {
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
			RawDescriptor: file_ParentQuestRandomInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ParentQuestRandomInfo_proto_goTypes,
		DependencyIndexes: file_ParentQuestRandomInfo_proto_depIdxs,
		MessageInfos:      file_ParentQuestRandomInfo_proto_msgTypes,
	}.Build()
	File_ParentQuestRandomInfo_proto = out.File
	file_ParentQuestRandomInfo_proto_rawDesc = nil
	file_ParentQuestRandomInfo_proto_goTypes = nil
	file_ParentQuestRandomInfo_proto_depIdxs = nil
}
