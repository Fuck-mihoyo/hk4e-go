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
// source: SaveMainCoopReq.proto

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

// CmdId: 1975
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type SaveMainCoopReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NormalVarMap   map[uint32]int32 `protobuf:"bytes,15,rep,name=normal_var_map,json=normalVarMap,proto3" json:"normal_var_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SelfConfidence uint32           `protobuf:"varint,2,opt,name=self_confidence,json=selfConfidence,proto3" json:"self_confidence,omitempty"`
	SavePointId    uint32           `protobuf:"varint,1,opt,name=save_point_id,json=savePointId,proto3" json:"save_point_id,omitempty"`
	TempVarMap     map[uint32]int32 `protobuf:"bytes,8,rep,name=temp_var_map,json=tempVarMap,proto3" json:"temp_var_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Id             uint32           `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveMainCoopReq) Reset() {
	*x = SaveMainCoopReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SaveMainCoopReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMainCoopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMainCoopReq) ProtoMessage() {}

func (x *SaveMainCoopReq) ProtoReflect() protoreflect.Message {
	mi := &file_SaveMainCoopReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMainCoopReq.ProtoReflect.Descriptor instead.
func (*SaveMainCoopReq) Descriptor() ([]byte, []int) {
	return file_SaveMainCoopReq_proto_rawDescGZIP(), []int{0}
}

func (x *SaveMainCoopReq) GetNormalVarMap() map[uint32]int32 {
	if x != nil {
		return x.NormalVarMap
	}
	return nil
}

func (x *SaveMainCoopReq) GetSelfConfidence() uint32 {
	if x != nil {
		return x.SelfConfidence
	}
	return 0
}

func (x *SaveMainCoopReq) GetSavePointId() uint32 {
	if x != nil {
		return x.SavePointId
	}
	return 0
}

func (x *SaveMainCoopReq) GetTempVarMap() map[uint32]int32 {
	if x != nil {
		return x.TempVarMap
	}
	return nil
}

func (x *SaveMainCoopReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_SaveMainCoopReq_proto protoreflect.FileDescriptor

var file_SaveMainCoopReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88,
	0x03, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x12, 0x4e, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x72,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x56, 0x61, 0x72, 0x4d,
	0x61, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x65, 0x6c,
	0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73,
	0x61, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x48, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x3f, 0x0a, 0x11, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x54, 0x65,
	0x6d, 0x70, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SaveMainCoopReq_proto_rawDescOnce sync.Once
	file_SaveMainCoopReq_proto_rawDescData = file_SaveMainCoopReq_proto_rawDesc
)

func file_SaveMainCoopReq_proto_rawDescGZIP() []byte {
	file_SaveMainCoopReq_proto_rawDescOnce.Do(func() {
		file_SaveMainCoopReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SaveMainCoopReq_proto_rawDescData)
	})
	return file_SaveMainCoopReq_proto_rawDescData
}

var file_SaveMainCoopReq_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_SaveMainCoopReq_proto_goTypes = []interface{}{
	(*SaveMainCoopReq)(nil), // 0: proto.SaveMainCoopReq
	nil,                     // 1: proto.SaveMainCoopReq.NormalVarMapEntry
	nil,                     // 2: proto.SaveMainCoopReq.TempVarMapEntry
}
var file_SaveMainCoopReq_proto_depIdxs = []int32{
	1, // 0: proto.SaveMainCoopReq.normal_var_map:type_name -> proto.SaveMainCoopReq.NormalVarMapEntry
	2, // 1: proto.SaveMainCoopReq.temp_var_map:type_name -> proto.SaveMainCoopReq.TempVarMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SaveMainCoopReq_proto_init() }
func file_SaveMainCoopReq_proto_init() {
	if File_SaveMainCoopReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SaveMainCoopReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMainCoopReq); i {
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
			RawDescriptor: file_SaveMainCoopReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SaveMainCoopReq_proto_goTypes,
		DependencyIndexes: file_SaveMainCoopReq_proto_depIdxs,
		MessageInfos:      file_SaveMainCoopReq_proto_msgTypes,
	}.Build()
	File_SaveMainCoopReq_proto = out.File
	file_SaveMainCoopReq_proto_rawDesc = nil
	file_SaveMainCoopReq_proto_goTypes = nil
	file_SaveMainCoopReq_proto_depIdxs = nil
}
