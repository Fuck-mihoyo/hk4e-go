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
// source: FungusFighterTrainingSettleInfo.proto

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

type FungusFighterTrainingSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsedTime    uint32            `protobuf:"varint,15,opt,name=used_time,json=usedTime,proto3" json:"used_time,omitempty"`
	Reason      GalleryStopReason `protobuf:"varint,3,opt,name=reason,proto3,enum=proto.GalleryStopReason" json:"reason,omitempty"`
	Transaction string            `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *FungusFighterTrainingSettleInfo) Reset() {
	*x = FungusFighterTrainingSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusFighterTrainingSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterTrainingSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterTrainingSettleInfo) ProtoMessage() {}

func (x *FungusFighterTrainingSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FungusFighterTrainingSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterTrainingSettleInfo.ProtoReflect.Descriptor instead.
func (*FungusFighterTrainingSettleInfo) Descriptor() ([]byte, []int) {
	return file_FungusFighterTrainingSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterTrainingSettleInfo) GetUsedTime() uint32 {
	if x != nil {
		return x.UsedTime
	}
	return 0
}

func (x *FungusFighterTrainingSettleInfo) GetReason() GalleryStopReason {
	if x != nil {
		return x.Reason
	}
	return GalleryStopReason_GALLERY_STOP_REASON_NONE
}

func (x *FungusFighterTrainingSettleInfo) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

var File_FungusFighterTrainingSettleInfo_proto protoreflect.FileDescriptor

var file_FungusFighterTrainingSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x25, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x1f, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FungusFighterTrainingSettleInfo_proto_rawDescOnce sync.Once
	file_FungusFighterTrainingSettleInfo_proto_rawDescData = file_FungusFighterTrainingSettleInfo_proto_rawDesc
)

func file_FungusFighterTrainingSettleInfo_proto_rawDescGZIP() []byte {
	file_FungusFighterTrainingSettleInfo_proto_rawDescOnce.Do(func() {
		file_FungusFighterTrainingSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusFighterTrainingSettleInfo_proto_rawDescData)
	})
	return file_FungusFighterTrainingSettleInfo_proto_rawDescData
}

var file_FungusFighterTrainingSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusFighterTrainingSettleInfo_proto_goTypes = []interface{}{
	(*FungusFighterTrainingSettleInfo)(nil), // 0: proto.FungusFighterTrainingSettleInfo
	(GalleryStopReason)(0),                  // 1: proto.GalleryStopReason
}
var file_FungusFighterTrainingSettleInfo_proto_depIdxs = []int32{
	1, // 0: proto.FungusFighterTrainingSettleInfo.reason:type_name -> proto.GalleryStopReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FungusFighterTrainingSettleInfo_proto_init() }
func file_FungusFighterTrainingSettleInfo_proto_init() {
	if File_FungusFighterTrainingSettleInfo_proto != nil {
		return
	}
	file_GalleryStopReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FungusFighterTrainingSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterTrainingSettleInfo); i {
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
			RawDescriptor: file_FungusFighterTrainingSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusFighterTrainingSettleInfo_proto_goTypes,
		DependencyIndexes: file_FungusFighterTrainingSettleInfo_proto_depIdxs,
		MessageInfos:      file_FungusFighterTrainingSettleInfo_proto_msgTypes,
	}.Build()
	File_FungusFighterTrainingSettleInfo_proto = out.File
	file_FungusFighterTrainingSettleInfo_proto_rawDesc = nil
	file_FungusFighterTrainingSettleInfo_proto_goTypes = nil
	file_FungusFighterTrainingSettleInfo_proto_depIdxs = nil
}
