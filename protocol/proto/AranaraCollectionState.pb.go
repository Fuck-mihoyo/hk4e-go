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
// source: AranaraCollectionState.proto

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

type AranaraCollectionState int32

const (
	AranaraCollectionState_ARANARA_COLLECTION_STATE_NONE        AranaraCollectionState = 0
	AranaraCollectionState_ARANARA_COLLECTION_STATE_COLLECTABLE AranaraCollectionState = 1
	AranaraCollectionState_ARANARA_COLLECTION_STATE_COLLECTED   AranaraCollectionState = 2
	AranaraCollectionState_ARANARA_COLLECTION_STATE_FINISHED    AranaraCollectionState = 3
)

// Enum value maps for AranaraCollectionState.
var (
	AranaraCollectionState_name = map[int32]string{
		0: "ARANARA_COLLECTION_STATE_NONE",
		1: "ARANARA_COLLECTION_STATE_COLLECTABLE",
		2: "ARANARA_COLLECTION_STATE_COLLECTED",
		3: "ARANARA_COLLECTION_STATE_FINISHED",
	}
	AranaraCollectionState_value = map[string]int32{
		"ARANARA_COLLECTION_STATE_NONE":        0,
		"ARANARA_COLLECTION_STATE_COLLECTABLE": 1,
		"ARANARA_COLLECTION_STATE_COLLECTED":   2,
		"ARANARA_COLLECTION_STATE_FINISHED":    3,
	}
)

func (x AranaraCollectionState) Enum() *AranaraCollectionState {
	p := new(AranaraCollectionState)
	*p = x
	return p
}

func (x AranaraCollectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AranaraCollectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_AranaraCollectionState_proto_enumTypes[0].Descriptor()
}

func (AranaraCollectionState) Type() protoreflect.EnumType {
	return &file_AranaraCollectionState_proto_enumTypes[0]
}

func (x AranaraCollectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AranaraCollectionState.Descriptor instead.
func (AranaraCollectionState) EnumDescriptor() ([]byte, []int) {
	return file_AranaraCollectionState_proto_rawDescGZIP(), []int{0}
}

var File_AranaraCollectionState_proto protoreflect.FileDescriptor

var file_AranaraCollectionState_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb4, 0x01, 0x0a, 0x16, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72,
	0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x1d, 0x41, 0x52, 0x41, 0x4e, 0x41, 0x52, 0x41, 0x5f, 0x43, 0x4f, 0x4c, 0x4c,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x52, 0x41, 0x4e, 0x41, 0x52, 0x41, 0x5f, 0x43,
	0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x26, 0x0a,
	0x22, 0x41, 0x52, 0x41, 0x4e, 0x41, 0x52, 0x41, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x52, 0x41, 0x4e, 0x41, 0x52, 0x41,
	0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AranaraCollectionState_proto_rawDescOnce sync.Once
	file_AranaraCollectionState_proto_rawDescData = file_AranaraCollectionState_proto_rawDesc
)

func file_AranaraCollectionState_proto_rawDescGZIP() []byte {
	file_AranaraCollectionState_proto_rawDescOnce.Do(func() {
		file_AranaraCollectionState_proto_rawDescData = protoimpl.X.CompressGZIP(file_AranaraCollectionState_proto_rawDescData)
	})
	return file_AranaraCollectionState_proto_rawDescData
}

var file_AranaraCollectionState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AranaraCollectionState_proto_goTypes = []interface{}{
	(AranaraCollectionState)(0), // 0: proto.AranaraCollectionState
}
var file_AranaraCollectionState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AranaraCollectionState_proto_init() }
func file_AranaraCollectionState_proto_init() {
	if File_AranaraCollectionState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AranaraCollectionState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AranaraCollectionState_proto_goTypes,
		DependencyIndexes: file_AranaraCollectionState_proto_depIdxs,
		EnumInfos:         file_AranaraCollectionState_proto_enumTypes,
	}.Build()
	File_AranaraCollectionState_proto = out.File
	file_AranaraCollectionState_proto_rawDesc = nil
	file_AranaraCollectionState_proto_goTypes = nil
	file_AranaraCollectionState_proto_depIdxs = nil
}
