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
// source: PlayerGeneralMatchDismissNotify.proto

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

// CmdId: 4191
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerGeneralMatchDismissNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UidList []uint32    `protobuf:"varint,3,rep,packed,name=uid_list,json=uidList,proto3" json:"uid_list,omitempty"`
	Reason  MatchReason `protobuf:"varint,13,opt,name=reason,proto3,enum=proto.MatchReason" json:"reason,omitempty"`
	MatchId uint32      `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
}

func (x *PlayerGeneralMatchDismissNotify) Reset() {
	*x = PlayerGeneralMatchDismissNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerGeneralMatchDismissNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerGeneralMatchDismissNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerGeneralMatchDismissNotify) ProtoMessage() {}

func (x *PlayerGeneralMatchDismissNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerGeneralMatchDismissNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerGeneralMatchDismissNotify.ProtoReflect.Descriptor instead.
func (*PlayerGeneralMatchDismissNotify) Descriptor() ([]byte, []int) {
	return file_PlayerGeneralMatchDismissNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerGeneralMatchDismissNotify) GetUidList() []uint32 {
	if x != nil {
		return x.UidList
	}
	return nil
}

func (x *PlayerGeneralMatchDismissNotify) GetReason() MatchReason {
	if x != nil {
		return x.Reason
	}
	return MatchReason_MATCH_REASON_NONE
}

func (x *PlayerGeneralMatchDismissNotify) GetMatchId() uint32 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

var File_PlayerGeneralMatchDismissNotify_proto protoreflect.FileDescriptor

var file_PlayerGeneralMatchDismissNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x83, 0x01, 0x0a, 0x1f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerGeneralMatchDismissNotify_proto_rawDescOnce sync.Once
	file_PlayerGeneralMatchDismissNotify_proto_rawDescData = file_PlayerGeneralMatchDismissNotify_proto_rawDesc
)

func file_PlayerGeneralMatchDismissNotify_proto_rawDescGZIP() []byte {
	file_PlayerGeneralMatchDismissNotify_proto_rawDescOnce.Do(func() {
		file_PlayerGeneralMatchDismissNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerGeneralMatchDismissNotify_proto_rawDescData)
	})
	return file_PlayerGeneralMatchDismissNotify_proto_rawDescData
}

var file_PlayerGeneralMatchDismissNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerGeneralMatchDismissNotify_proto_goTypes = []interface{}{
	(*PlayerGeneralMatchDismissNotify)(nil), // 0: proto.PlayerGeneralMatchDismissNotify
	(MatchReason)(0),                        // 1: proto.MatchReason
}
var file_PlayerGeneralMatchDismissNotify_proto_depIdxs = []int32{
	1, // 0: proto.PlayerGeneralMatchDismissNotify.reason:type_name -> proto.MatchReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerGeneralMatchDismissNotify_proto_init() }
func file_PlayerGeneralMatchDismissNotify_proto_init() {
	if File_PlayerGeneralMatchDismissNotify_proto != nil {
		return
	}
	file_MatchReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerGeneralMatchDismissNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerGeneralMatchDismissNotify); i {
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
			RawDescriptor: file_PlayerGeneralMatchDismissNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerGeneralMatchDismissNotify_proto_goTypes,
		DependencyIndexes: file_PlayerGeneralMatchDismissNotify_proto_depIdxs,
		MessageInfos:      file_PlayerGeneralMatchDismissNotify_proto_msgTypes,
	}.Build()
	File_PlayerGeneralMatchDismissNotify_proto = out.File
	file_PlayerGeneralMatchDismissNotify_proto_rawDesc = nil
	file_PlayerGeneralMatchDismissNotify_proto_goTypes = nil
	file_PlayerGeneralMatchDismissNotify_proto_depIdxs = nil
}
