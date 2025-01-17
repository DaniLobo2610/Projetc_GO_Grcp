// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/kombat.proto

package Kombat

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

type KombatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *KombatRequest) Reset() {
	*x = KombatRequest{}
	mi := &file_proto_kombat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KombatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KombatRequest) ProtoMessage() {}

func (x *KombatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KombatRequest.ProtoReflect.Descriptor instead.
func (*KombatRequest) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{0}
}

func (x *KombatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type KombatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Skills string `protobuf:"bytes,3,opt,name=skills,proto3" json:"skills,omitempty"`
}

func (x *KombatResponse) Reset() {
	*x = KombatResponse{}
	mi := &file_proto_kombat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KombatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KombatResponse) ProtoMessage() {}

func (x *KombatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KombatResponse.ProtoReflect.Descriptor instead.
func (*KombatResponse) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{1}
}

func (x *KombatResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KombatResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *KombatResponse) GetSkills() string {
	if x != nil {
		return x.Skills
	}
	return ""
}

type NewKombatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Skills string `protobuf:"bytes,3,opt,name=skills,proto3" json:"skills,omitempty"`
}

func (x *NewKombatRequest) Reset() {
	*x = NewKombatRequest{}
	mi := &file_proto_kombat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewKombatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewKombatRequest) ProtoMessage() {}

func (x *NewKombatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewKombatRequest.ProtoReflect.Descriptor instead.
func (*NewKombatRequest) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{2}
}

func (x *NewKombatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewKombatRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NewKombatRequest) GetSkills() string {
	if x != nil {
		return x.Skills
	}
	return ""
}

type AddKombatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddKombatResponse) Reset() {
	*x = AddKombatResponse{}
	mi := &file_proto_kombat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddKombatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKombatResponse) ProtoMessage() {}

func (x *AddKombatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKombatResponse.ProtoReflect.Descriptor instead.
func (*AddKombatResponse) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{3}
}

func (x *AddKombatResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AddKombatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_kombat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{4}
}

type KombatTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *KombatTypeRequest) Reset() {
	*x = KombatTypeRequest{}
	mi := &file_proto_kombat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KombatTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KombatTypeRequest) ProtoMessage() {}

func (x *KombatTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KombatTypeRequest.ProtoReflect.Descriptor instead.
func (*KombatTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{5}
}

func (x *KombatTypeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type KombatSkillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skills string `protobuf:"bytes,1,opt,name=skills,proto3" json:"skills,omitempty"`
}

func (x *KombatSkillsRequest) Reset() {
	*x = KombatSkillsRequest{}
	mi := &file_proto_kombat_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KombatSkillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KombatSkillsRequest) ProtoMessage() {}

func (x *KombatSkillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kombat_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KombatSkillsRequest.ProtoReflect.Descriptor instead.
func (*KombatSkillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_kombat_proto_rawDescGZIP(), []int{6}
}

func (x *KombatSkillsRequest) GetSkills() string {
	if x != nil {
		return x.Skills
	}
	return ""
}

var File_proto_kombat_proto protoreflect.FileDescriptor

var file_proto_kombat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x22, 0x23, 0x0a, 0x0d,
	0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x50, 0x0a, 0x0e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4b, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x27, 0x0a, 0x11, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x4b, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x32, 0xe6, 0x02, 0x0a, 0x0d, 0x4b, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x4b,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x4b, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x4b,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x4b, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x4b, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x4e, 0x65, 0x77,
	0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x2e,
	0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4b, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x42, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x4b, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e,
	0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x4b, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kombat_proto_rawDescOnce sync.Once
	file_proto_kombat_proto_rawDescData = file_proto_kombat_proto_rawDesc
)

func file_proto_kombat_proto_rawDescGZIP() []byte {
	file_proto_kombat_proto_rawDescOnce.Do(func() {
		file_proto_kombat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kombat_proto_rawDescData)
	})
	return file_proto_kombat_proto_rawDescData
}

var file_proto_kombat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_kombat_proto_goTypes = []any{
	(*KombatRequest)(nil),       // 0: Kombat.KombatRequest
	(*KombatResponse)(nil),      // 1: Kombat.KombatResponse
	(*NewKombatRequest)(nil),    // 2: Kombat.NewKombatRequest
	(*AddKombatResponse)(nil),   // 3: Kombat.AddKombatResponse
	(*Empty)(nil),               // 4: Kombat.Empty
	(*KombatTypeRequest)(nil),   // 5: Kombat.KombatTypeRequest
	(*KombatSkillsRequest)(nil), // 6: Kombat.KombatSkillsRequest
}
var file_proto_kombat_proto_depIdxs = []int32{
	0, // 0: Kombat.KombatService.GetKombatInfo:input_type -> Kombat.KombatRequest
	4, // 1: Kombat.KombatService.GetKombatList:input_type -> Kombat.Empty
	2, // 2: Kombat.KombatService.addKombats:input_type -> Kombat.NewKombatRequest
	5, // 3: Kombat.KombatService.GetKombatByType:input_type -> Kombat.KombatTypeRequest
	6, // 4: Kombat.KombatService.GetKombatBySkills:input_type -> Kombat.KombatSkillsRequest
	1, // 5: Kombat.KombatService.GetKombatInfo:output_type -> Kombat.KombatResponse
	1, // 6: Kombat.KombatService.GetKombatList:output_type -> Kombat.KombatResponse
	3, // 7: Kombat.KombatService.addKombats:output_type -> Kombat.AddKombatResponse
	1, // 8: Kombat.KombatService.GetKombatByType:output_type -> Kombat.KombatResponse
	1, // 9: Kombat.KombatService.GetKombatBySkills:output_type -> Kombat.KombatResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kombat_proto_init() }
func file_proto_kombat_proto_init() {
	if File_proto_kombat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_kombat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kombat_proto_goTypes,
		DependencyIndexes: file_proto_kombat_proto_depIdxs,
		MessageInfos:      file_proto_kombat_proto_msgTypes,
	}.Build()
	File_proto_kombat_proto = out.File
	file_proto_kombat_proto_rawDesc = nil
	file_proto_kombat_proto_goTypes = nil
	file_proto_kombat_proto_depIdxs = nil
}
