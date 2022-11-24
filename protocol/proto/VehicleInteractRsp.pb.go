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
// source: VehicleInteractRsp.proto

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

// CmdId: 804
// EnetChannelId: 0
// EnetIsReliable: true
type VehicleInteractRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InteractType VehicleInteractType `protobuf:"varint,15,opt,name=interact_type,json=interactType,proto3,enum=proto.VehicleInteractType" json:"interact_type,omitempty"`
	Member       *VehicleMember      `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
	EntityId     uint32              `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Retcode      int32               `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *VehicleInteractRsp) Reset() {
	*x = VehicleInteractRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VehicleInteractRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleInteractRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleInteractRsp) ProtoMessage() {}

func (x *VehicleInteractRsp) ProtoReflect() protoreflect.Message {
	mi := &file_VehicleInteractRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleInteractRsp.ProtoReflect.Descriptor instead.
func (*VehicleInteractRsp) Descriptor() ([]byte, []int) {
	return file_VehicleInteractRsp_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleInteractRsp) GetInteractType() VehicleInteractType {
	if x != nil {
		return x.InteractType
	}
	return VehicleInteractType_VEHICLE_INTERACT_TYPE_NONE
}

func (x *VehicleInteractRsp) GetMember() *VehicleMember {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *VehicleInteractRsp) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *VehicleInteractRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_VehicleInteractRsp_proto protoreflect.FileDescriptor

var file_VehicleInteractRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_VehicleInteractRsp_proto_rawDescOnce sync.Once
	file_VehicleInteractRsp_proto_rawDescData = file_VehicleInteractRsp_proto_rawDesc
)

func file_VehicleInteractRsp_proto_rawDescGZIP() []byte {
	file_VehicleInteractRsp_proto_rawDescOnce.Do(func() {
		file_VehicleInteractRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_VehicleInteractRsp_proto_rawDescData)
	})
	return file_VehicleInteractRsp_proto_rawDescData
}

var file_VehicleInteractRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_VehicleInteractRsp_proto_goTypes = []interface{}{
	(*VehicleInteractRsp)(nil), // 0: proto.VehicleInteractRsp
	(VehicleInteractType)(0),   // 1: proto.VehicleInteractType
	(*VehicleMember)(nil),      // 2: proto.VehicleMember
}
var file_VehicleInteractRsp_proto_depIdxs = []int32{
	1, // 0: proto.VehicleInteractRsp.interact_type:type_name -> proto.VehicleInteractType
	2, // 1: proto.VehicleInteractRsp.member:type_name -> proto.VehicleMember
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_VehicleInteractRsp_proto_init() }
func file_VehicleInteractRsp_proto_init() {
	if File_VehicleInteractRsp_proto != nil {
		return
	}
	file_VehicleInteractType_proto_init()
	file_VehicleMember_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VehicleInteractRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleInteractRsp); i {
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
			RawDescriptor: file_VehicleInteractRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VehicleInteractRsp_proto_goTypes,
		DependencyIndexes: file_VehicleInteractRsp_proto_depIdxs,
		MessageInfos:      file_VehicleInteractRsp_proto_msgTypes,
	}.Build()
	File_VehicleInteractRsp_proto = out.File
	file_VehicleInteractRsp_proto_rawDesc = nil
	file_VehicleInteractRsp_proto_goTypes = nil
	file_VehicleInteractRsp_proto_depIdxs = nil
}
