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
// source: WidgetReportReq.proto

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

// CmdId: 4291
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type WidgetReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_MFEHMLBNNAG bool   `protobuf:"varint,5,opt,name=Unk2700_MFEHMLBNNAG,json=Unk2700MFEHMLBNNAG,proto3" json:"Unk2700_MFEHMLBNNAG,omitempty"`
	IsClientCollect     bool   `protobuf:"varint,14,opt,name=is_client_collect,json=isClientCollect,proto3" json:"is_client_collect,omitempty"`
	IsClearHint         bool   `protobuf:"varint,13,opt,name=is_clear_hint,json=isClearHint,proto3" json:"is_clear_hint,omitempty"`
	MaterialId          uint32 `protobuf:"varint,15,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
}

func (x *WidgetReportReq) Reset() {
	*x = WidgetReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetReportReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetReportReq) ProtoMessage() {}

func (x *WidgetReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetReportReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetReportReq.ProtoReflect.Descriptor instead.
func (*WidgetReportReq) Descriptor() ([]byte, []int) {
	return file_WidgetReportReq_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetReportReq) GetUnk2700_MFEHMLBNNAG() bool {
	if x != nil {
		return x.Unk2700_MFEHMLBNNAG
	}
	return false
}

func (x *WidgetReportReq) GetIsClientCollect() bool {
	if x != nil {
		return x.IsClientCollect
	}
	return false
}

func (x *WidgetReportReq) GetIsClearHint() bool {
	if x != nil {
		return x.IsClearHint
	}
	return false
}

func (x *WidgetReportReq) GetMaterialId() uint32 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

var File_WidgetReportReq_proto protoreflect.FileDescriptor

var file_WidgetReportReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x46, 0x45, 0x48, 0x4d, 0x4c, 0x42, 0x4e, 0x4e,
	0x41, 0x47, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x4d, 0x46, 0x45, 0x48, 0x4d, 0x4c, 0x42, 0x4e, 0x4e, 0x41, 0x47, 0x12, 0x2a, 0x0a, 0x11,
	0x69, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63,
	0x6c, 0x65, 0x61, 0x72, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_WidgetReportReq_proto_rawDescOnce sync.Once
	file_WidgetReportReq_proto_rawDescData = file_WidgetReportReq_proto_rawDesc
)

func file_WidgetReportReq_proto_rawDescGZIP() []byte {
	file_WidgetReportReq_proto_rawDescOnce.Do(func() {
		file_WidgetReportReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetReportReq_proto_rawDescData)
	})
	return file_WidgetReportReq_proto_rawDescData
}

var file_WidgetReportReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetReportReq_proto_goTypes = []interface{}{
	(*WidgetReportReq)(nil), // 0: WidgetReportReq
}
var file_WidgetReportReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WidgetReportReq_proto_init() }
func file_WidgetReportReq_proto_init() {
	if File_WidgetReportReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WidgetReportReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetReportReq); i {
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
			RawDescriptor: file_WidgetReportReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetReportReq_proto_goTypes,
		DependencyIndexes: file_WidgetReportReq_proto_depIdxs,
		MessageInfos:      file_WidgetReportReq_proto_msgTypes,
	}.Build()
	File_WidgetReportReq_proto = out.File
	file_WidgetReportReq_proto_rawDesc = nil
	file_WidgetReportReq_proto_goTypes = nil
	file_WidgetReportReq_proto_depIdxs = nil
}
