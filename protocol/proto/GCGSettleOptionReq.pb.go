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
// source: GCGSettleOptionReq.proto

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

// CmdId: 7124
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GCGSettleOptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option GCGSettleOption `protobuf:"varint,5,opt,name=option,proto3,enum=proto.GCGSettleOption" json:"option,omitempty"`
}

func (x *GCGSettleOptionReq) Reset() {
	*x = GCGSettleOptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGSettleOptionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGSettleOptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGSettleOptionReq) ProtoMessage() {}

func (x *GCGSettleOptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_GCGSettleOptionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGSettleOptionReq.ProtoReflect.Descriptor instead.
func (*GCGSettleOptionReq) Descriptor() ([]byte, []int) {
	return file_GCGSettleOptionReq_proto_rawDescGZIP(), []int{0}
}

func (x *GCGSettleOptionReq) GetOption() GCGSettleOption {
	if x != nil {
		return x.Option
	}
	return GCGSettleOption_GCG_SETTLE_OPTION_OPT_NONE
}

var File_GCGSettleOptionReq_proto protoreflect.FileDescriptor

var file_GCGSettleOptionReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x43, 0x47, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x47, 0x43, 0x47, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x43, 0x47, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x2e,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GCGSettleOptionReq_proto_rawDescOnce sync.Once
	file_GCGSettleOptionReq_proto_rawDescData = file_GCGSettleOptionReq_proto_rawDesc
)

func file_GCGSettleOptionReq_proto_rawDescGZIP() []byte {
	file_GCGSettleOptionReq_proto_rawDescOnce.Do(func() {
		file_GCGSettleOptionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGSettleOptionReq_proto_rawDescData)
	})
	return file_GCGSettleOptionReq_proto_rawDescData
}

var file_GCGSettleOptionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGSettleOptionReq_proto_goTypes = []interface{}{
	(*GCGSettleOptionReq)(nil), // 0: proto.GCGSettleOptionReq
	(GCGSettleOption)(0),       // 1: proto.GCGSettleOption
}
var file_GCGSettleOptionReq_proto_depIdxs = []int32{
	1, // 0: proto.GCGSettleOptionReq.option:type_name -> proto.GCGSettleOption
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGSettleOptionReq_proto_init() }
func file_GCGSettleOptionReq_proto_init() {
	if File_GCGSettleOptionReq_proto != nil {
		return
	}
	file_GCGSettleOption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGSettleOptionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGSettleOptionReq); i {
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
			RawDescriptor: file_GCGSettleOptionReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGSettleOptionReq_proto_goTypes,
		DependencyIndexes: file_GCGSettleOptionReq_proto_depIdxs,
		MessageInfos:      file_GCGSettleOptionReq_proto_msgTypes,
	}.Build()
	File_GCGSettleOptionReq_proto = out.File
	file_GCGSettleOptionReq_proto_rawDesc = nil
	file_GCGSettleOptionReq_proto_goTypes = nil
	file_GCGSettleOptionReq_proto_depIdxs = nil
}
