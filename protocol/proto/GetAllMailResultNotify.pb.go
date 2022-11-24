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
// source: GetAllMailResultNotify.proto

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

// CmdId: 1481
// EnetChannelId: 0
// EnetIsReliable: true
type GetAllMailResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction    string      `protobuf:"bytes,9,opt,name=transaction,proto3" json:"transaction,omitempty"`
	MailList       []*MailData `protobuf:"bytes,5,rep,name=mail_list,json=mailList,proto3" json:"mail_list,omitempty"`
	PageIndex      uint32      `protobuf:"varint,11,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	TotalPageCount uint32      `protobuf:"varint,4,opt,name=total_page_count,json=totalPageCount,proto3" json:"total_page_count,omitempty"`
	IsCollected    bool        `protobuf:"varint,7,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"`
	Retcode        int32       `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetAllMailResultNotify) Reset() {
	*x = GetAllMailResultNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAllMailResultNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMailResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailResultNotify) ProtoMessage() {}

func (x *GetAllMailResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GetAllMailResultNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailResultNotify.ProtoReflect.Descriptor instead.
func (*GetAllMailResultNotify) Descriptor() ([]byte, []int) {
	return file_GetAllMailResultNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllMailResultNotify) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

func (x *GetAllMailResultNotify) GetMailList() []*MailData {
	if x != nil {
		return x.MailList
	}
	return nil
}

func (x *GetAllMailResultNotify) GetPageIndex() uint32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetAllMailResultNotify) GetTotalPageCount() uint32 {
	if x != nil {
		return x.TotalPageCount
	}
	return 0
}

func (x *GetAllMailResultNotify) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

func (x *GetAllMailResultNotify) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetAllMailResultNotify_proto protoreflect.FileDescriptor

var file_GetAllMailResultNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAllMailResultNotify_proto_rawDescOnce sync.Once
	file_GetAllMailResultNotify_proto_rawDescData = file_GetAllMailResultNotify_proto_rawDesc
)

func file_GetAllMailResultNotify_proto_rawDescGZIP() []byte {
	file_GetAllMailResultNotify_proto_rawDescOnce.Do(func() {
		file_GetAllMailResultNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAllMailResultNotify_proto_rawDescData)
	})
	return file_GetAllMailResultNotify_proto_rawDescData
}

var file_GetAllMailResultNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAllMailResultNotify_proto_goTypes = []interface{}{
	(*GetAllMailResultNotify)(nil), // 0: proto.GetAllMailResultNotify
	(*MailData)(nil),               // 1: proto.MailData
}
var file_GetAllMailResultNotify_proto_depIdxs = []int32{
	1, // 0: proto.GetAllMailResultNotify.mail_list:type_name -> proto.MailData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetAllMailResultNotify_proto_init() }
func file_GetAllMailResultNotify_proto_init() {
	if File_GetAllMailResultNotify_proto != nil {
		return
	}
	file_MailData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetAllMailResultNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMailResultNotify); i {
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
			RawDescriptor: file_GetAllMailResultNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAllMailResultNotify_proto_goTypes,
		DependencyIndexes: file_GetAllMailResultNotify_proto_depIdxs,
		MessageInfos:      file_GetAllMailResultNotify_proto_msgTypes,
	}.Build()
	File_GetAllMailResultNotify_proto = out.File
	file_GetAllMailResultNotify_proto_rawDesc = nil
	file_GetAllMailResultNotify_proto_goTypes = nil
	file_GetAllMailResultNotify_proto_depIdxs = nil
}
