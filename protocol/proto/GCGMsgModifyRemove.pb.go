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
// source: GCGMsgModifyRemove.proto

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

type GCGMsgModifyRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerId  uint32    `protobuf:"varint,14,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	Reason        GCGReason `protobuf:"varint,12,opt,name=reason,proto3,enum=proto.GCGReason" json:"reason,omitempty"`
	OwnerCardGuid uint32    `protobuf:"varint,5,opt,name=owner_card_guid,json=ownerCardGuid,proto3" json:"owner_card_guid,omitempty"`
	CardGuidList  []uint32  `protobuf:"varint,4,rep,packed,name=card_guid_list,json=cardGuidList,proto3" json:"card_guid_list,omitempty"`
}

func (x *GCGMsgModifyRemove) Reset() {
	*x = GCGMsgModifyRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgModifyRemove_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgModifyRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgModifyRemove) ProtoMessage() {}

func (x *GCGMsgModifyRemove) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgModifyRemove_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgModifyRemove.ProtoReflect.Descriptor instead.
func (*GCGMsgModifyRemove) Descriptor() ([]byte, []int) {
	return file_GCGMsgModifyRemove_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgModifyRemove) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGMsgModifyRemove) GetReason() GCGReason {
	if x != nil {
		return x.Reason
	}
	return GCGReason_GCG_REASON_DEFAULT
}

func (x *GCGMsgModifyRemove) GetOwnerCardGuid() uint32 {
	if x != nil {
		return x.OwnerCardGuid
	}
	return 0
}

func (x *GCGMsgModifyRemove) GetCardGuidList() []uint32 {
	if x != nil {
		return x.CardGuidList
	}
	return nil
}

var File_GCGMsgModifyRemove_proto protoreflect.FileDescriptor

var file_GCGMsgModifyRemove_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x47, 0x43, 0x47, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x47, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x47, 0x75,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgModifyRemove_proto_rawDescOnce sync.Once
	file_GCGMsgModifyRemove_proto_rawDescData = file_GCGMsgModifyRemove_proto_rawDesc
)

func file_GCGMsgModifyRemove_proto_rawDescGZIP() []byte {
	file_GCGMsgModifyRemove_proto_rawDescOnce.Do(func() {
		file_GCGMsgModifyRemove_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgModifyRemove_proto_rawDescData)
	})
	return file_GCGMsgModifyRemove_proto_rawDescData
}

var file_GCGMsgModifyRemove_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgModifyRemove_proto_goTypes = []interface{}{
	(*GCGMsgModifyRemove)(nil), // 0: proto.GCGMsgModifyRemove
	(GCGReason)(0),             // 1: proto.GCGReason
}
var file_GCGMsgModifyRemove_proto_depIdxs = []int32{
	1, // 0: proto.GCGMsgModifyRemove.reason:type_name -> proto.GCGReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgModifyRemove_proto_init() }
func file_GCGMsgModifyRemove_proto_init() {
	if File_GCGMsgModifyRemove_proto != nil {
		return
	}
	file_GCGReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgModifyRemove_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgModifyRemove); i {
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
			RawDescriptor: file_GCGMsgModifyRemove_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgModifyRemove_proto_goTypes,
		DependencyIndexes: file_GCGMsgModifyRemove_proto_depIdxs,
		MessageInfos:      file_GCGMsgModifyRemove_proto_msgTypes,
	}.Build()
	File_GCGMsgModifyRemove_proto = out.File
	file_GCGMsgModifyRemove_proto_rawDesc = nil
	file_GCGMsgModifyRemove_proto_goTypes = nil
	file_GCGMsgModifyRemove_proto_depIdxs = nil
}
