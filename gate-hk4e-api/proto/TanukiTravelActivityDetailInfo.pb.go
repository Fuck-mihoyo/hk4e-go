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
// source: TanukiTravelActivityDetailInfo.proto

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

type TanukiTravelActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_JBPFIDDPGME []*Unk2700_BIFNFOGBPNM `protobuf:"bytes,4,rep,name=Unk2700_JBPFIDDPGME,json=Unk2700JBPFIDDPGME,proto3" json:"Unk2700_JBPFIDDPGME,omitempty"`
	IsContentClosed     bool                   `protobuf:"varint,11,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	Unk2700_BHHCNOLMCJM uint32                 `protobuf:"varint,10,opt,name=Unk2700_BHHCNOLMCJM,json=Unk2700BHHCNOLMCJM,proto3" json:"Unk2700_BHHCNOLMCJM,omitempty"`
}

func (x *TanukiTravelActivityDetailInfo) Reset() {
	*x = TanukiTravelActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TanukiTravelActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TanukiTravelActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TanukiTravelActivityDetailInfo) ProtoMessage() {}

func (x *TanukiTravelActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TanukiTravelActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TanukiTravelActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*TanukiTravelActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_TanukiTravelActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TanukiTravelActivityDetailInfo) GetUnk2700_JBPFIDDPGME() []*Unk2700_BIFNFOGBPNM {
	if x != nil {
		return x.Unk2700_JBPFIDDPGME
	}
	return nil
}

func (x *TanukiTravelActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *TanukiTravelActivityDetailInfo) GetUnk2700_BHHCNOLMCJM() uint32 {
	if x != nil {
		return x.Unk2700_BHHCNOLMCJM
	}
	return 0
}

var File_TanukiTravelActivityDetailInfo_proto protoreflect.FileDescriptor

var file_TanukiTravelActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x24, 0x54, 0x61, 0x6e, 0x75, 0x6b, 0x69, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x42, 0x49, 0x46, 0x4e, 0x46, 0x4f, 0x47, 0x42, 0x50, 0x4e, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x1e, 0x54, 0x61, 0x6e, 0x75, 0x6b, 0x69, 0x54, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x4a, 0x42, 0x50, 0x46, 0x49, 0x44, 0x44, 0x50, 0x47, 0x4d, 0x45, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x42, 0x49, 0x46, 0x4e,
	0x46, 0x4f, 0x47, 0x42, 0x50, 0x4e, 0x4d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x4a, 0x42, 0x50, 0x46, 0x49, 0x44, 0x44, 0x50, 0x47, 0x4d, 0x45, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x42, 0x48, 0x48, 0x43, 0x4e, 0x4f, 0x4c, 0x4d, 0x43, 0x4a, 0x4d, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x42, 0x48, 0x48,
	0x43, 0x4e, 0x4f, 0x4c, 0x4d, 0x43, 0x4a, 0x4d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TanukiTravelActivityDetailInfo_proto_rawDescOnce sync.Once
	file_TanukiTravelActivityDetailInfo_proto_rawDescData = file_TanukiTravelActivityDetailInfo_proto_rawDesc
)

func file_TanukiTravelActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_TanukiTravelActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_TanukiTravelActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TanukiTravelActivityDetailInfo_proto_rawDescData)
	})
	return file_TanukiTravelActivityDetailInfo_proto_rawDescData
}

var file_TanukiTravelActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TanukiTravelActivityDetailInfo_proto_goTypes = []interface{}{
	(*TanukiTravelActivityDetailInfo)(nil), // 0: TanukiTravelActivityDetailInfo
	(*Unk2700_BIFNFOGBPNM)(nil),            // 1: Unk2700_BIFNFOGBPNM
}
var file_TanukiTravelActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: TanukiTravelActivityDetailInfo.Unk2700_JBPFIDDPGME:type_name -> Unk2700_BIFNFOGBPNM
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TanukiTravelActivityDetailInfo_proto_init() }
func file_TanukiTravelActivityDetailInfo_proto_init() {
	if File_TanukiTravelActivityDetailInfo_proto != nil {
		return
	}
	file_Unk2700_BIFNFOGBPNM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TanukiTravelActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TanukiTravelActivityDetailInfo); i {
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
			RawDescriptor: file_TanukiTravelActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TanukiTravelActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_TanukiTravelActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_TanukiTravelActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_TanukiTravelActivityDetailInfo_proto = out.File
	file_TanukiTravelActivityDetailInfo_proto_rawDesc = nil
	file_TanukiTravelActivityDetailInfo_proto_goTypes = nil
	file_TanukiTravelActivityDetailInfo_proto_depIdxs = nil
}
