// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: reader.proto

package pbtest

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

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{0}
}

func (x *GetTestRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type TestGetManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *TestGetManyRequest) Reset() {
	*x = TestGetManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestGetManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestGetManyRequest) ProtoMessage() {}

func (x *TestGetManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestGetManyRequest.ProtoReflect.Descriptor instead.
func (*TestGetManyRequest) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{1}
}

func (x *TestGetManyRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type TestPrefixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TestPrefixRequest) Reset() {
	*x = TestPrefixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPrefixRequest) ProtoMessage() {}

func (x *TestPrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPrefixRequest.ProtoReflect.Descriptor instead.
func (*TestPrefixRequest) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{2}
}

func (x *TestPrefixRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *TestPrefixRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TestScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start        string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	ExclusiveEnd string `protobuf:"bytes,2,opt,name=exclusive_end,json=exclusiveEnd,proto3" json:"exclusive_end,omitempty"`
	Limit        int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TestScanRequest) Reset() {
	*x = TestScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestScanRequest) ProtoMessage() {}

func (x *TestScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestScanRequest.ProtoReflect.Descriptor instead.
func (*TestScanRequest) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{3}
}

func (x *TestScanRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *TestScanRequest) GetExclusiveEnd() string {
	if x != nil {
		return x.ExclusiveEnd
	}
	return ""
}

func (x *TestScanRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Tuples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*Tuple `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *Tuples) Reset() {
	*x = Tuples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuples) ProtoMessage() {}

func (x *Tuples) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuples.ProtoReflect.Descriptor instead.
func (*Tuples) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{4}
}

func (x *Tuples) GetPairs() []*Tuple {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type OptionalTuples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Pairs []*Tuple `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *OptionalTuples) Reset() {
	*x = OptionalTuples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalTuples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalTuples) ProtoMessage() {}

func (x *OptionalTuples) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalTuples.ProtoReflect.Descriptor instead.
func (*OptionalTuples) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{5}
}

func (x *OptionalTuples) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *OptionalTuples) GetPairs() []*Tuple {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type Tuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tuple) Reset() {
	*x = Tuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reader_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuple) ProtoMessage() {}

func (x *Tuple) ProtoReflect() protoreflect.Message {
	mi := &file_reader_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuple.ProtoReflect.Descriptor instead.
func (*Tuple) Descriptor() ([]byte, []int) {
	return file_reader_proto_rawDescGZIP(), []int{6}
}

func (x *Tuple) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tuple) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_reader_proto protoreflect.FileDescriptor

var file_reader_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28,
	0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x41, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x62, 0x0a, 0x0f, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x31, 0x0a, 0x06, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x66, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x22, 0x4f, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x61,
	0x69, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x66, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x70, 0x61,
	0x69, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x05, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x98, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x12,
	0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x66,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x30, 0x01,
	0x12, 0x4b, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12,
	0x1e, 0x2e, 0x73, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x30, 0x01, 0x12, 0x41, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x2e, 0x73, 0x66,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x66, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x30, 0x01,
	0x12, 0x3d, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x1b, 0x2e, 0x73,
	0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x66, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x30, 0x01, 0x42,
	0x19, 0x5a, 0x17, 0x73, 0x69, 0x6e, 0x6b, 0x2d, 0x6b, 0x76, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_reader_proto_rawDescOnce sync.Once
	file_reader_proto_rawDescData = file_reader_proto_rawDesc
)

func file_reader_proto_rawDescGZIP() []byte {
	file_reader_proto_rawDescOnce.Do(func() {
		file_reader_proto_rawDescData = protoimpl.X.CompressGZIP(file_reader_proto_rawDescData)
	})
	return file_reader_proto_rawDescData
}

var file_reader_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_reader_proto_goTypes = []interface{}{
	(*GetTestRequest)(nil),     // 0: sf.test.v1.GetTestRequest
	(*TestGetManyRequest)(nil), // 1: sf.test.v1.TestGetManyRequest
	(*TestPrefixRequest)(nil),  // 2: sf.test.v1.TestPrefixRequest
	(*TestScanRequest)(nil),    // 3: sf.test.v1.TestScanRequest
	(*Tuples)(nil),             // 4: sf.test.v1.Tuples
	(*OptionalTuples)(nil),     // 5: sf.test.v1.OptionalTuples
	(*Tuple)(nil),              // 6: sf.test.v1.Tuple
}
var file_reader_proto_depIdxs = []int32{
	6, // 0: sf.test.v1.Tuples.pairs:type_name -> sf.test.v1.Tuple
	6, // 1: sf.test.v1.OptionalTuples.pairs:type_name -> sf.test.v1.Tuple
	0, // 2: sf.test.v1.TestService.TestGet:input_type -> sf.test.v1.GetTestRequest
	1, // 3: sf.test.v1.TestService.TestGetMany:input_type -> sf.test.v1.TestGetManyRequest
	2, // 4: sf.test.v1.TestService.TestPrefix:input_type -> sf.test.v1.TestPrefixRequest
	3, // 5: sf.test.v1.TestService.TestScan:input_type -> sf.test.v1.TestScanRequest
	6, // 6: sf.test.v1.TestService.TestGet:output_type -> sf.test.v1.Tuple
	5, // 7: sf.test.v1.TestService.TestGetMany:output_type -> sf.test.v1.OptionalTuples
	4, // 8: sf.test.v1.TestService.TestPrefix:output_type -> sf.test.v1.Tuples
	4, // 9: sf.test.v1.TestService.TestScan:output_type -> sf.test.v1.Tuples
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reader_proto_init() }
func file_reader_proto_init() {
	if File_reader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
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
		file_reader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestGetManyRequest); i {
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
		file_reader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPrefixRequest); i {
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
		file_reader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestScanRequest); i {
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
		file_reader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuples); i {
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
		file_reader_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionalTuples); i {
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
		file_reader_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuple); i {
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
			RawDescriptor: file_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reader_proto_goTypes,
		DependencyIndexes: file_reader_proto_depIdxs,
		MessageInfos:      file_reader_proto_msgTypes,
	}.Build()
	File_reader_proto = out.File
	file_reader_proto_rawDesc = nil
	file_reader_proto_goTypes = nil
	file_reader_proto_depIdxs = nil
}
