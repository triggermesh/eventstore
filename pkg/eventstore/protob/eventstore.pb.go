//
//Copyright (c) 2020 TriggerMesh Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.8.0
// source: pkg/eventstore/protob/eventstore.proto

package protob

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ScopeChoice int32

const (
	ScopeChoice_Instance ScopeChoice = 0
	ScopeChoice_Bridge   ScopeChoice = 1
	ScopeChoice_Global   ScopeChoice = 2
)

// Enum value maps for ScopeChoice.
var (
	ScopeChoice_name = map[int32]string{
		0: "Instance",
		1: "Bridge",
		2: "Global",
	}
	ScopeChoice_value = map[string]int32{
		"Instance": 0,
		"Bridge":   1,
		"Global":   2,
	}
)

func (x ScopeChoice) Enum() *ScopeChoice {
	p := new(ScopeChoice)
	*p = x
	return p
}

func (x ScopeChoice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScopeChoice) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_eventstore_protob_eventstore_proto_enumTypes[0].Descriptor()
}

func (ScopeChoice) Type() protoreflect.EnumType {
	return &file_pkg_eventstore_protob_eventstore_proto_enumTypes[0]
}

func (x ScopeChoice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScopeChoice.Descriptor instead.
func (ScopeChoice) EnumDescriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{0}
}

type ScopeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     ScopeChoice `protobuf:"varint,1,opt,name=type,proto3,enum=protob.ScopeChoice" json:"type,omitempty"`
	Bridge   string      `protobuf:"bytes,2,opt,name=bridge,proto3" json:"bridge,omitempty"`
	Instance string      `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *ScopeType) Reset() {
	*x = ScopeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeType) ProtoMessage() {}

func (x *ScopeType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeType.ProtoReflect.Descriptor instead.
func (*ScopeType) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{0}
}

func (x *ScopeType) GetType() ScopeChoice {
	if x != nil {
		return x.Type
	}
	return ScopeChoice_Instance
}

func (x *ScopeType) GetBridge() string {
	if x != nil {
		return x.Bridge
	}
	return ""
}

func (x *ScopeType) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type LocationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope *ScopeType `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	Key   string     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *LocationType) Reset() {
	*x = LocationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationType) ProtoMessage() {}

func (x *LocationType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationType.ProtoReflect.Descriptor instead.
func (*LocationType) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{1}
}

func (x *LocationType) GetScope() *ScopeType {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *LocationType) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LocationType `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Ttl      int32         `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Value    []byte        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{2}
}

func (x *SaveRequest) GetLocation() *LocationType {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *SaveRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *SaveRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{3}
}

type LoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LocationType `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *LoadRequest) Reset() {
	*x = LoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRequest) ProtoMessage() {}

func (x *LoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRequest.ProtoReflect.Descriptor instead.
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{4}
}

func (x *LoadRequest) GetLocation() *LocationType {
	if x != nil {
		return x.Location
	}
	return nil
}

type LoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LoadResponse) Reset() {
	*x = LoadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResponse) ProtoMessage() {}

func (x *LoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResponse.ProtoReflect.Descriptor instead.
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{5}
}

func (x *LoadResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LocationType `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetLocation() *LocationType {
	if x != nil {
		return x.Location
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_eventstore_protob_eventstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP(), []int{7}
}

var File_pkg_eventstore_protob_eventstore_proto protoreflect.FileDescriptor

var file_pkg_eventstore_protob_eventstore_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x22, 0x68, 0x0a, 0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x0c, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x67, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f,
	0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x24, 0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x33, 0x0a, 0x0b, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x10, 0x02, 0x32,
	0xb1, 0x01, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_eventstore_protob_eventstore_proto_rawDescOnce sync.Once
	file_pkg_eventstore_protob_eventstore_proto_rawDescData = file_pkg_eventstore_protob_eventstore_proto_rawDesc
)

func file_pkg_eventstore_protob_eventstore_proto_rawDescGZIP() []byte {
	file_pkg_eventstore_protob_eventstore_proto_rawDescOnce.Do(func() {
		file_pkg_eventstore_protob_eventstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_eventstore_protob_eventstore_proto_rawDescData)
	})
	return file_pkg_eventstore_protob_eventstore_proto_rawDescData
}

var file_pkg_eventstore_protob_eventstore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_eventstore_protob_eventstore_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_eventstore_protob_eventstore_proto_goTypes = []interface{}{
	(ScopeChoice)(0),       // 0: protob.ScopeChoice
	(*ScopeType)(nil),      // 1: protob.ScopeType
	(*LocationType)(nil),   // 2: protob.LocationType
	(*SaveRequest)(nil),    // 3: protob.SaveRequest
	(*SaveResponse)(nil),   // 4: protob.SaveResponse
	(*LoadRequest)(nil),    // 5: protob.LoadRequest
	(*LoadResponse)(nil),   // 6: protob.LoadResponse
	(*DeleteRequest)(nil),  // 7: protob.DeleteRequest
	(*DeleteResponse)(nil), // 8: protob.DeleteResponse
}
var file_pkg_eventstore_protob_eventstore_proto_depIdxs = []int32{
	0, // 0: protob.ScopeType.type:type_name -> protob.ScopeChoice
	1, // 1: protob.LocationType.scope:type_name -> protob.ScopeType
	2, // 2: protob.SaveRequest.location:type_name -> protob.LocationType
	2, // 3: protob.LoadRequest.location:type_name -> protob.LocationType
	2, // 4: protob.DeleteRequest.location:type_name -> protob.LocationType
	3, // 5: protob.EventStore.Save:input_type -> protob.SaveRequest
	5, // 6: protob.EventStore.Load:input_type -> protob.LoadRequest
	7, // 7: protob.EventStore.Delete:input_type -> protob.DeleteRequest
	4, // 8: protob.EventStore.Save:output_type -> protob.SaveResponse
	6, // 9: protob.EventStore.Load:output_type -> protob.LoadResponse
	8, // 10: protob.EventStore.Delete:output_type -> protob.DeleteResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_eventstore_protob_eventstore_proto_init() }
func file_pkg_eventstore_protob_eventstore_proto_init() {
	if File_pkg_eventstore_protob_eventstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeType); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationType); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadRequest); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadResponse); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_pkg_eventstore_protob_eventstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_pkg_eventstore_protob_eventstore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_eventstore_protob_eventstore_proto_goTypes,
		DependencyIndexes: file_pkg_eventstore_protob_eventstore_proto_depIdxs,
		EnumInfos:         file_pkg_eventstore_protob_eventstore_proto_enumTypes,
		MessageInfos:      file_pkg_eventstore_protob_eventstore_proto_msgTypes,
	}.Build()
	File_pkg_eventstore_protob_eventstore_proto = out.File
	file_pkg_eventstore_protob_eventstore_proto_rawDesc = nil
	file_pkg_eventstore_protob_eventstore_proto_goTypes = nil
	file_pkg_eventstore_protob_eventstore_proto_depIdxs = nil
}