// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: loms_v1.proto

package loms_v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderPayedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderPayedRequest) Reset() {
	*x = OrderPayedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedRequest) ProtoMessage() {}

func (x *OrderPayedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedRequest.ProtoReflect.Descriptor instead.
func (*OrderPayedRequest) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{3}
}

func (x *OrderPayedRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderPayedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderPayedResponse) Reset() {
	*x = OrderPayedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedResponse) ProtoMessage() {}

func (x *OrderPayedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedResponse.ProtoReflect.Descriptor instead.
func (*OrderPayedResponse) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{4}
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CancelOrderRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_loms_v1_proto_rawDescGZIP(), []int{6}
}

var File_loms_v1_proto protoreflect.FileDescriptor

var file_loms_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x14, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a,
	0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x15, 0x0a,
	0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xeb, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x12,
	0x1a, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f,
	0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f, 0x6c,
	0x6f, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loms_v1_proto_rawDescOnce sync.Once
	file_loms_v1_proto_rawDescData = file_loms_v1_proto_rawDesc
)

func file_loms_v1_proto_rawDescGZIP() []byte {
	file_loms_v1_proto_rawDescOnce.Do(func() {
		file_loms_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_loms_v1_proto_rawDescData)
	})
	return file_loms_v1_proto_rawDescData
}

var file_loms_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_loms_v1_proto_goTypes = []interface{}{
	(*Item)(nil),                // 0: loms_v1.Item
	(*CreateOrderRequest)(nil),  // 1: loms_v1.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 2: loms_v1.CreateOrderResponse
	(*OrderPayedRequest)(nil),   // 3: loms_v1.OrderPayedRequest
	(*OrderPayedResponse)(nil),  // 4: loms_v1.OrderPayedResponse
	(*CancelOrderRequest)(nil),  // 5: loms_v1.CancelOrderRequest
	(*CancelOrderResponse)(nil), // 6: loms_v1.CancelOrderResponse
}
var file_loms_v1_proto_depIdxs = []int32{
	0, // 0: loms_v1.CreateOrderRequest.items:type_name -> loms_v1.Item
	1, // 1: loms_v1.ProductService.CreateOrder:input_type -> loms_v1.CreateOrderRequest
	3, // 2: loms_v1.ProductService.OrderPayed:input_type -> loms_v1.OrderPayedRequest
	5, // 3: loms_v1.ProductService.CancelOrder:input_type -> loms_v1.CancelOrderRequest
	2, // 4: loms_v1.ProductService.CreateOrder:output_type -> loms_v1.CreateOrderResponse
	4, // 5: loms_v1.ProductService.OrderPayed:output_type -> loms_v1.OrderPayedResponse
	6, // 6: loms_v1.ProductService.CancelOrder:output_type -> loms_v1.CancelOrderResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_loms_v1_proto_init() }
func file_loms_v1_proto_init() {
	if File_loms_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loms_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_loms_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_loms_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_loms_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedRequest); i {
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
		file_loms_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedResponse); i {
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
		file_loms_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_loms_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderResponse); i {
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
			RawDescriptor: file_loms_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loms_v1_proto_goTypes,
		DependencyIndexes: file_loms_v1_proto_depIdxs,
		MessageInfos:      file_loms_v1_proto_msgTypes,
	}.Build()
	File_loms_v1_proto = out.File
	file_loms_v1_proto_rawDesc = nil
	file_loms_v1_proto_goTypes = nil
	file_loms_v1_proto_depIdxs = nil
}
