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
// source: AranaraCollectionSuite.proto

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

type AranaraCollectionSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionIdStateMap map[uint32]AranaraCollectionState `protobuf:"bytes,6,rep,name=collection_id_state_map,json=collectionIdStateMap,proto3" json:"collection_id_state_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=proto.AranaraCollectionState"`
	CollectionType       uint32                            `protobuf:"varint,12,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
}

func (x *AranaraCollectionSuite) Reset() {
	*x = AranaraCollectionSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AranaraCollectionSuite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AranaraCollectionSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AranaraCollectionSuite) ProtoMessage() {}

func (x *AranaraCollectionSuite) ProtoReflect() protoreflect.Message {
	mi := &file_AranaraCollectionSuite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AranaraCollectionSuite.ProtoReflect.Descriptor instead.
func (*AranaraCollectionSuite) Descriptor() ([]byte, []int) {
	return file_AranaraCollectionSuite_proto_rawDescGZIP(), []int{0}
}

func (x *AranaraCollectionSuite) GetCollectionIdStateMap() map[uint32]AranaraCollectionState {
	if x != nil {
		return x.CollectionIdStateMap
	}
	return nil
}

func (x *AranaraCollectionSuite) GetCollectionType() uint32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

var File_AranaraCollectionSuite_proto protoreflect.FileDescriptor

var file_AranaraCollectionSuite_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x16, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x6e,
	0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x66, 0x0a, 0x19, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72,
	0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_AranaraCollectionSuite_proto_rawDescOnce sync.Once
	file_AranaraCollectionSuite_proto_rawDescData = file_AranaraCollectionSuite_proto_rawDesc
)

func file_AranaraCollectionSuite_proto_rawDescGZIP() []byte {
	file_AranaraCollectionSuite_proto_rawDescOnce.Do(func() {
		file_AranaraCollectionSuite_proto_rawDescData = protoimpl.X.CompressGZIP(file_AranaraCollectionSuite_proto_rawDescData)
	})
	return file_AranaraCollectionSuite_proto_rawDescData
}

var file_AranaraCollectionSuite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AranaraCollectionSuite_proto_goTypes = []interface{}{
	(*AranaraCollectionSuite)(nil), // 0: proto.AranaraCollectionSuite
	nil,                            // 1: proto.AranaraCollectionSuite.CollectionIdStateMapEntry
	(AranaraCollectionState)(0),    // 2: proto.AranaraCollectionState
}
var file_AranaraCollectionSuite_proto_depIdxs = []int32{
	1, // 0: proto.AranaraCollectionSuite.collection_id_state_map:type_name -> proto.AranaraCollectionSuite.CollectionIdStateMapEntry
	2, // 1: proto.AranaraCollectionSuite.CollectionIdStateMapEntry.value:type_name -> proto.AranaraCollectionState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AranaraCollectionSuite_proto_init() }
func file_AranaraCollectionSuite_proto_init() {
	if File_AranaraCollectionSuite_proto != nil {
		return
	}
	file_AranaraCollectionState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AranaraCollectionSuite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AranaraCollectionSuite); i {
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
			RawDescriptor: file_AranaraCollectionSuite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AranaraCollectionSuite_proto_goTypes,
		DependencyIndexes: file_AranaraCollectionSuite_proto_depIdxs,
		MessageInfos:      file_AranaraCollectionSuite_proto_msgTypes,
	}.Build()
	File_AranaraCollectionSuite_proto = out.File
	file_AranaraCollectionSuite_proto_rawDesc = nil
	file_AranaraCollectionSuite_proto_goTypes = nil
	file_AranaraCollectionSuite_proto_depIdxs = nil
}
