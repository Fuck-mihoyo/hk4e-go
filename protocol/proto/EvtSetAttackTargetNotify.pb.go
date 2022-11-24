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
// source: EvtSetAttackTargetNotify.proto

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

// CmdId: 399
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type EvtSetAttackTargetNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType            ForwardType             `protobuf:"varint,1,opt,name=forward_type,json=forwardType,proto3,enum=proto.ForwardType" json:"forward_type,omitempty"`
	EvtSetAttackTargetInfo *EvtSetAttackTargetInfo `protobuf:"bytes,11,opt,name=evt_set_attack_target_info,json=evtSetAttackTargetInfo,proto3" json:"evt_set_attack_target_info,omitempty"`
}

func (x *EvtSetAttackTargetNotify) Reset() {
	*x = EvtSetAttackTargetNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtSetAttackTargetNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtSetAttackTargetNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtSetAttackTargetNotify) ProtoMessage() {}

func (x *EvtSetAttackTargetNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtSetAttackTargetNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtSetAttackTargetNotify.ProtoReflect.Descriptor instead.
func (*EvtSetAttackTargetNotify) Descriptor() ([]byte, []int) {
	return file_EvtSetAttackTargetNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtSetAttackTargetNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *EvtSetAttackTargetNotify) GetEvtSetAttackTargetInfo() *EvtSetAttackTargetInfo {
	if x != nil {
		return x.EvtSetAttackTargetInfo
	}
	return nil
}

var File_EvtSetAttackTargetNotify_proto protoreflect.FileDescriptor

var file_EvtSetAttackTargetNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x45, 0x76, 0x74, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x45, 0x76, 0x74, 0x53, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x18, 0x45, 0x76, 0x74,
	0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x35, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x1a,
	0x65, 0x76, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x74, 0x53, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x16, 0x65, 0x76, 0x74, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtSetAttackTargetNotify_proto_rawDescOnce sync.Once
	file_EvtSetAttackTargetNotify_proto_rawDescData = file_EvtSetAttackTargetNotify_proto_rawDesc
)

func file_EvtSetAttackTargetNotify_proto_rawDescGZIP() []byte {
	file_EvtSetAttackTargetNotify_proto_rawDescOnce.Do(func() {
		file_EvtSetAttackTargetNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtSetAttackTargetNotify_proto_rawDescData)
	})
	return file_EvtSetAttackTargetNotify_proto_rawDescData
}

var file_EvtSetAttackTargetNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtSetAttackTargetNotify_proto_goTypes = []interface{}{
	(*EvtSetAttackTargetNotify)(nil), // 0: proto.EvtSetAttackTargetNotify
	(ForwardType)(0),                 // 1: proto.ForwardType
	(*EvtSetAttackTargetInfo)(nil),   // 2: proto.EvtSetAttackTargetInfo
}
var file_EvtSetAttackTargetNotify_proto_depIdxs = []int32{
	1, // 0: proto.EvtSetAttackTargetNotify.forward_type:type_name -> proto.ForwardType
	2, // 1: proto.EvtSetAttackTargetNotify.evt_set_attack_target_info:type_name -> proto.EvtSetAttackTargetInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtSetAttackTargetNotify_proto_init() }
func file_EvtSetAttackTargetNotify_proto_init() {
	if File_EvtSetAttackTargetNotify_proto != nil {
		return
	}
	file_EvtSetAttackTargetInfo_proto_init()
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtSetAttackTargetNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtSetAttackTargetNotify); i {
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
			RawDescriptor: file_EvtSetAttackTargetNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtSetAttackTargetNotify_proto_goTypes,
		DependencyIndexes: file_EvtSetAttackTargetNotify_proto_depIdxs,
		MessageInfos:      file_EvtSetAttackTargetNotify_proto_msgTypes,
	}.Build()
	File_EvtSetAttackTargetNotify_proto = out.File
	file_EvtSetAttackTargetNotify_proto_rawDesc = nil
	file_EvtSetAttackTargetNotify_proto_goTypes = nil
	file_EvtSetAttackTargetNotify_proto_depIdxs = nil
}
