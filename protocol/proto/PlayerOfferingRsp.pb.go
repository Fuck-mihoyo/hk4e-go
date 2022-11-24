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
// source: PlayerOfferingRsp.proto

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

// CmdId: 2917
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerOfferingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList     []*ItemParam        `protobuf:"bytes,7,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	Retcode      int32               `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OfferingData *PlayerOfferingData `protobuf:"bytes,10,opt,name=offering_data,json=offeringData,proto3" json:"offering_data,omitempty"`
}

func (x *PlayerOfferingRsp) Reset() {
	*x = PlayerOfferingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerOfferingRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerOfferingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerOfferingRsp) ProtoMessage() {}

func (x *PlayerOfferingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerOfferingRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerOfferingRsp.ProtoReflect.Descriptor instead.
func (*PlayerOfferingRsp) Descriptor() ([]byte, []int) {
	return file_PlayerOfferingRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerOfferingRsp) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *PlayerOfferingRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlayerOfferingRsp) GetOfferingData() *PlayerOfferingData {
	if x != nil {
		return x.OfferingData
	}
	return nil
}

var File_PlayerOfferingRsp_proto protoreflect.FileDescriptor

var file_PlayerOfferingRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x73,
	0x70, 0x12, 0x2d, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerOfferingRsp_proto_rawDescOnce sync.Once
	file_PlayerOfferingRsp_proto_rawDescData = file_PlayerOfferingRsp_proto_rawDesc
)

func file_PlayerOfferingRsp_proto_rawDescGZIP() []byte {
	file_PlayerOfferingRsp_proto_rawDescOnce.Do(func() {
		file_PlayerOfferingRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerOfferingRsp_proto_rawDescData)
	})
	return file_PlayerOfferingRsp_proto_rawDescData
}

var file_PlayerOfferingRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerOfferingRsp_proto_goTypes = []interface{}{
	(*PlayerOfferingRsp)(nil),  // 0: proto.PlayerOfferingRsp
	(*ItemParam)(nil),          // 1: proto.ItemParam
	(*PlayerOfferingData)(nil), // 2: proto.PlayerOfferingData
}
var file_PlayerOfferingRsp_proto_depIdxs = []int32{
	1, // 0: proto.PlayerOfferingRsp.item_list:type_name -> proto.ItemParam
	2, // 1: proto.PlayerOfferingRsp.offering_data:type_name -> proto.PlayerOfferingData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerOfferingRsp_proto_init() }
func file_PlayerOfferingRsp_proto_init() {
	if File_PlayerOfferingRsp_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	file_PlayerOfferingData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerOfferingRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerOfferingRsp); i {
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
			RawDescriptor: file_PlayerOfferingRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerOfferingRsp_proto_goTypes,
		DependencyIndexes: file_PlayerOfferingRsp_proto_depIdxs,
		MessageInfos:      file_PlayerOfferingRsp_proto_msgTypes,
	}.Build()
	File_PlayerOfferingRsp_proto = out.File
	file_PlayerOfferingRsp_proto_rawDesc = nil
	file_PlayerOfferingRsp_proto_goTypes = nil
	file_PlayerOfferingRsp_proto_depIdxs = nil
}
