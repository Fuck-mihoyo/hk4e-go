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
// source: SummerTimeV2DetailInfo.proto

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

type SummerTimeV2DetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2800_PNBLCPIBKPO []*Unk2800_CGODFDDALAG `protobuf:"bytes,13,rep,name=Unk2800_PNBLCPIBKPO,json=Unk2800PNBLCPIBKPO,proto3" json:"Unk2800_PNBLCPIBKPO,omitempty"`
	Unk2800_HDEFJKGDNEH uint32                 `protobuf:"varint,10,opt,name=Unk2800_HDEFJKGDNEH,json=Unk2800HDEFJKGDNEH,proto3" json:"Unk2800_HDEFJKGDNEH,omitempty"`
	IsContentClosed     bool                   `protobuf:"varint,4,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	Unk2800_ELHBCNPKOJG uint32                 `protobuf:"varint,5,opt,name=Unk2800_ELHBCNPKOJG,json=Unk2800ELHBCNPKOJG,proto3" json:"Unk2800_ELHBCNPKOJG,omitempty"`
	Unk2800_MPKLJJIEHIB []*Unk2800_CGPNLBNMPCM `protobuf:"bytes,15,rep,name=Unk2800_MPKLJJIEHIB,json=Unk2800MPKLJJIEHIB,proto3" json:"Unk2800_MPKLJJIEHIB,omitempty"`
}

func (x *SummerTimeV2DetailInfo) Reset() {
	*x = SummerTimeV2DetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SummerTimeV2DetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummerTimeV2DetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummerTimeV2DetailInfo) ProtoMessage() {}

func (x *SummerTimeV2DetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SummerTimeV2DetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummerTimeV2DetailInfo.ProtoReflect.Descriptor instead.
func (*SummerTimeV2DetailInfo) Descriptor() ([]byte, []int) {
	return file_SummerTimeV2DetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SummerTimeV2DetailInfo) GetUnk2800_PNBLCPIBKPO() []*Unk2800_CGODFDDALAG {
	if x != nil {
		return x.Unk2800_PNBLCPIBKPO
	}
	return nil
}

func (x *SummerTimeV2DetailInfo) GetUnk2800_HDEFJKGDNEH() uint32 {
	if x != nil {
		return x.Unk2800_HDEFJKGDNEH
	}
	return 0
}

func (x *SummerTimeV2DetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *SummerTimeV2DetailInfo) GetUnk2800_ELHBCNPKOJG() uint32 {
	if x != nil {
		return x.Unk2800_ELHBCNPKOJG
	}
	return 0
}

func (x *SummerTimeV2DetailInfo) GetUnk2800_MPKLJJIEHIB() []*Unk2800_CGPNLBNMPCM {
	if x != nil {
		return x.Unk2800_MPKLJJIEHIB
	}
	return nil
}

var File_SummerTimeV2DetailInfo_proto protoreflect.FileDescriptor

var file_SummerTimeV2DetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x32, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x43, 0x47, 0x4f, 0x44, 0x46, 0x44, 0x44, 0x41,
	0x4c, 0x41, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38,
	0x30, 0x30, 0x5f, 0x43, 0x47, 0x50, 0x4e, 0x4c, 0x42, 0x4e, 0x4d, 0x50, 0x43, 0x4d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x16, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x56, 0x32, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x50, 0x4e, 0x42, 0x4c, 0x43,
	0x50, 0x49, 0x42, 0x4b, 0x50, 0x4f, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55,
	0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x43, 0x47, 0x4f, 0x44, 0x46, 0x44, 0x44, 0x41, 0x4c,
	0x41, 0x47, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x50, 0x4e, 0x42, 0x4c, 0x43,
	0x50, 0x49, 0x42, 0x4b, 0x50, 0x4f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30,
	0x30, 0x5f, 0x48, 0x44, 0x45, 0x46, 0x4a, 0x4b, 0x47, 0x44, 0x4e, 0x45, 0x48, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x48, 0x44, 0x45, 0x46,
	0x4a, 0x4b, 0x47, 0x44, 0x4e, 0x45, 0x48, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x45,
	0x4c, 0x48, 0x42, 0x43, 0x4e, 0x50, 0x4b, 0x4f, 0x4a, 0x47, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x45, 0x4c, 0x48, 0x42, 0x43, 0x4e, 0x50,
	0x4b, 0x4f, 0x4a, 0x47, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f,
	0x4d, 0x50, 0x4b, 0x4c, 0x4a, 0x4a, 0x49, 0x45, 0x48, 0x49, 0x42, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x43, 0x47, 0x50, 0x4e,
	0x4c, 0x42, 0x4e, 0x4d, 0x50, 0x43, 0x4d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30,
	0x4d, 0x50, 0x4b, 0x4c, 0x4a, 0x4a, 0x49, 0x45, 0x48, 0x49, 0x42, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SummerTimeV2DetailInfo_proto_rawDescOnce sync.Once
	file_SummerTimeV2DetailInfo_proto_rawDescData = file_SummerTimeV2DetailInfo_proto_rawDesc
)

func file_SummerTimeV2DetailInfo_proto_rawDescGZIP() []byte {
	file_SummerTimeV2DetailInfo_proto_rawDescOnce.Do(func() {
		file_SummerTimeV2DetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SummerTimeV2DetailInfo_proto_rawDescData)
	})
	return file_SummerTimeV2DetailInfo_proto_rawDescData
}

var file_SummerTimeV2DetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SummerTimeV2DetailInfo_proto_goTypes = []interface{}{
	(*SummerTimeV2DetailInfo)(nil), // 0: SummerTimeV2DetailInfo
	(*Unk2800_CGODFDDALAG)(nil),    // 1: Unk2800_CGODFDDALAG
	(*Unk2800_CGPNLBNMPCM)(nil),    // 2: Unk2800_CGPNLBNMPCM
}
var file_SummerTimeV2DetailInfo_proto_depIdxs = []int32{
	1, // 0: SummerTimeV2DetailInfo.Unk2800_PNBLCPIBKPO:type_name -> Unk2800_CGODFDDALAG
	2, // 1: SummerTimeV2DetailInfo.Unk2800_MPKLJJIEHIB:type_name -> Unk2800_CGPNLBNMPCM
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SummerTimeV2DetailInfo_proto_init() }
func file_SummerTimeV2DetailInfo_proto_init() {
	if File_SummerTimeV2DetailInfo_proto != nil {
		return
	}
	file_Unk2800_CGODFDDALAG_proto_init()
	file_Unk2800_CGPNLBNMPCM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SummerTimeV2DetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummerTimeV2DetailInfo); i {
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
			RawDescriptor: file_SummerTimeV2DetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SummerTimeV2DetailInfo_proto_goTypes,
		DependencyIndexes: file_SummerTimeV2DetailInfo_proto_depIdxs,
		MessageInfos:      file_SummerTimeV2DetailInfo_proto_msgTypes,
	}.Build()
	File_SummerTimeV2DetailInfo_proto = out.File
	file_SummerTimeV2DetailInfo_proto_rawDesc = nil
	file_SummerTimeV2DetailInfo_proto_goTypes = nil
	file_SummerTimeV2DetailInfo_proto_depIdxs = nil
}
