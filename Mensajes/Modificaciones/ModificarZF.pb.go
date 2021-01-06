// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: ModificarZF.proto

package ModificarZF

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Modificacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accion        string `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
	ValorAfectado string `protobuf:"bytes,2,opt,name=valor_afectado,json=valorAfectado,proto3" json:"valor_afectado,omitempty"`
	NuevoValor    string `protobuf:"bytes,3,opt,name=nuevo_valor,json=nuevoValor,proto3" json:"nuevo_valor,omitempty"`
}

func (x *Modificacion) Reset() {
	*x = Modificacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ModificarZF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Modificacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Modificacion) ProtoMessage() {}

func (x *Modificacion) ProtoReflect() protoreflect.Message {
	mi := &file_ModificarZF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Modificacion.ProtoReflect.Descriptor instead.
func (*Modificacion) Descriptor() ([]byte, []int) {
	return file_ModificarZF_proto_rawDescGZIP(), []int{0}
}

func (x *Modificacion) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

func (x *Modificacion) GetValorAfectado() string {
	if x != nil {
		return x.ValorAfectado
	}
	return ""
}

func (x *Modificacion) GetNuevoValor() string {
	if x != nil {
		return x.NuevoValor
	}
	return ""
}

type RespuestaDNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reloj string `protobuf:"bytes,1,opt,name=reloj,proto3" json:"reloj,omitempty"`
}

func (x *RespuestaDNS) Reset() {
	*x = RespuestaDNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ModificarZF_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaDNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaDNS) ProtoMessage() {}

func (x *RespuestaDNS) ProtoReflect() protoreflect.Message {
	mi := &file_ModificarZF_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaDNS.ProtoReflect.Descriptor instead.
func (*RespuestaDNS) Descriptor() ([]byte, []int) {
	return file_ModificarZF_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaDNS) GetReloj() string {
	if x != nil {
		return x.Reloj
	}
	return ""
}

type Solicitud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dominio string `protobuf:"bytes,1,opt,name=Dominio,proto3" json:"Dominio,omitempty"`
	Reloj   string `protobuf:"bytes,2,opt,name=Reloj,proto3" json:"Reloj,omitempty"`
}

func (x *Solicitud) Reset() {
	*x = Solicitud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ModificarZF_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solicitud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solicitud) ProtoMessage() {}

func (x *Solicitud) ProtoReflect() protoreflect.Message {
	mi := &file_ModificarZF_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solicitud.ProtoReflect.Descriptor instead.
func (*Solicitud) Descriptor() ([]byte, []int) {
	return file_ModificarZF_proto_rawDescGZIP(), []int{2}
}

func (x *Solicitud) GetDominio() string {
	if x != nil {
		return x.Dominio
	}
	return ""
}

func (x *Solicitud) GetReloj() string {
	if x != nil {
		return x.Reloj
	}
	return ""
}

type RespuestaBroker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPserver string `protobuf:"bytes,1,opt,name=IPserver,proto3" json:"IPserver,omitempty"`
}

func (x *RespuestaBroker) Reset() {
	*x = RespuestaBroker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ModificarZF_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaBroker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaBroker) ProtoMessage() {}

func (x *RespuestaBroker) ProtoReflect() protoreflect.Message {
	mi := &file_ModificarZF_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaBroker.ProtoReflect.Descriptor instead.
func (*RespuestaBroker) Descriptor() ([]byte, []int) {
	return file_ModificarZF_proto_rawDescGZIP(), []int{3}
}

func (x *RespuestaBroker) GetIPserver() string {
	if x != nil {
		return x.IPserver
	}
	return ""
}

type Enviodom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dominio string `protobuf:"bytes,1,opt,name=Dominio,proto3" json:"Dominio,omitempty"`
}

func (x *Enviodom) Reset() {
	*x = Enviodom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ModificarZF_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enviodom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enviodom) ProtoMessage() {}

func (x *Enviodom) ProtoReflect() protoreflect.Message {
	mi := &file_ModificarZF_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enviodom.ProtoReflect.Descriptor instead.
func (*Enviodom) Descriptor() ([]byte, []int) {
	return file_ModificarZF_proto_rawDescGZIP(), []int{4}
}

func (x *Enviodom) GetDominio() string {
	if x != nil {
		return x.Dominio
	}
	return ""
}

var File_ModificarZF_proto protoreflect.FileDescriptor

var file_ModificarZF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46,
	0x22, 0x6e, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x5f, 0x61, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72,
	0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x44, 0x4e, 0x53,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x22, 0x3b, 0x0a, 0x09, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x75, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65,
	0x6c, 0x6f, 0x6a, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x50, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x50, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0x24, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x64, 0x6f, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x32, 0x69, 0x0a, 0x19, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x72, 0x5a, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x75, 0x64, 0x1a, 0x1c, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a,
	0x46, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x22, 0x00, 0x32, 0x65, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72,
	0x5a, 0x46, 0x45, 0x6e, 0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x45, 0x6e, 0x44,
	0x4e, 0x53, 0x12, 0x19, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46,
	0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x44, 0x4e, 0x53, 0x22, 0x00, 0x32, 0x53, 0x0a, 0x10, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x6f, 0x64,
	0x6f, 0x6d, 0x1a, 0x19, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x5a, 0x46,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x44, 0x4e, 0x53, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ModificarZF_proto_rawDescOnce sync.Once
	file_ModificarZF_proto_rawDescData = file_ModificarZF_proto_rawDesc
)

func file_ModificarZF_proto_rawDescGZIP() []byte {
	file_ModificarZF_proto_rawDescOnce.Do(func() {
		file_ModificarZF_proto_rawDescData = protoimpl.X.CompressGZIP(file_ModificarZF_proto_rawDescData)
	})
	return file_ModificarZF_proto_rawDescData
}

var file_ModificarZF_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ModificarZF_proto_goTypes = []interface{}{
	(*Modificacion)(nil),    // 0: ModificarZF.Modificacion
	(*RespuestaDNS)(nil),    // 1: ModificarZF.RespuestaDNS
	(*Solicitud)(nil),       // 2: ModificarZF.Solicitud
	(*RespuestaBroker)(nil), // 3: ModificarZF.RespuestaBroker
	(*Enviodom)(nil),        // 4: ModificarZF.Enviodom
}
var file_ModificarZF_proto_depIdxs = []int32{
	2, // 0: ModificarZF.ModificarZFRequestService.ModificarZFRequest:input_type -> ModificarZF.Solicitud
	0, // 1: ModificarZF.ModificarZFEnDNSService.ModificarZFEnDNS:input_type -> ModificarZF.Modificacion
	4, // 2: ModificarZF.VerificarService.Verificar:input_type -> ModificarZF.Enviodom
	3, // 3: ModificarZF.ModificarZFRequestService.ModificarZFRequest:output_type -> ModificarZF.RespuestaBroker
	1, // 4: ModificarZF.ModificarZFEnDNSService.ModificarZFEnDNS:output_type -> ModificarZF.RespuestaDNS
	1, // 5: ModificarZF.VerificarService.Verificar:output_type -> ModificarZF.RespuestaDNS
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ModificarZF_proto_init() }
func file_ModificarZF_proto_init() {
	if File_ModificarZF_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ModificarZF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Modificacion); i {
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
		file_ModificarZF_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaDNS); i {
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
		file_ModificarZF_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solicitud); i {
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
		file_ModificarZF_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaBroker); i {
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
		file_ModificarZF_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enviodom); i {
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
			RawDescriptor: file_ModificarZF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_ModificarZF_proto_goTypes,
		DependencyIndexes: file_ModificarZF_proto_depIdxs,
		MessageInfos:      file_ModificarZF_proto_msgTypes,
	}.Build()
	File_ModificarZF_proto = out.File
	file_ModificarZF_proto_rawDesc = nil
	file_ModificarZF_proto_goTypes = nil
	file_ModificarZF_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ModificarZFRequestServiceClient is the client API for ModificarZFRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModificarZFRequestServiceClient interface {
	ModificarZFRequest(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaBroker, error)
}

type modificarZFRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModificarZFRequestServiceClient(cc grpc.ClientConnInterface) ModificarZFRequestServiceClient {
	return &modificarZFRequestServiceClient{cc}
}

func (c *modificarZFRequestServiceClient) ModificarZFRequest(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/ModificarZF.ModificarZFRequestService/ModificarZFRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModificarZFRequestServiceServer is the server API for ModificarZFRequestService service.
type ModificarZFRequestServiceServer interface {
	ModificarZFRequest(context.Context, *Solicitud) (*RespuestaBroker, error)
}

// UnimplementedModificarZFRequestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModificarZFRequestServiceServer struct {
}

func (*UnimplementedModificarZFRequestServiceServer) ModificarZFRequest(context.Context, *Solicitud) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModificarZFRequest not implemented")
}

func RegisterModificarZFRequestServiceServer(s *grpc.Server, srv ModificarZFRequestServiceServer) {
	s.RegisterService(&_ModificarZFRequestService_serviceDesc, srv)
}

func _ModificarZFRequestService_ModificarZFRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solicitud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModificarZFRequestServiceServer).ModificarZFRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModificarZF.ModificarZFRequestService/ModificarZFRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModificarZFRequestServiceServer).ModificarZFRequest(ctx, req.(*Solicitud))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModificarZFRequestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ModificarZF.ModificarZFRequestService",
	HandlerType: (*ModificarZFRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModificarZFRequest",
			Handler:    _ModificarZFRequestService_ModificarZFRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ModificarZF.proto",
}

// ModificarZFEnDNSServiceClient is the client API for ModificarZFEnDNSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModificarZFEnDNSServiceClient interface {
	ModificarZFEnDNS(ctx context.Context, in *Modificacion, opts ...grpc.CallOption) (*RespuestaDNS, error)
}

type modificarZFEnDNSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModificarZFEnDNSServiceClient(cc grpc.ClientConnInterface) ModificarZFEnDNSServiceClient {
	return &modificarZFEnDNSServiceClient{cc}
}

func (c *modificarZFEnDNSServiceClient) ModificarZFEnDNS(ctx context.Context, in *Modificacion, opts ...grpc.CallOption) (*RespuestaDNS, error) {
	out := new(RespuestaDNS)
	err := c.cc.Invoke(ctx, "/ModificarZF.ModificarZFEnDNSService/ModificarZFEnDNS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModificarZFEnDNSServiceServer is the server API for ModificarZFEnDNSService service.
type ModificarZFEnDNSServiceServer interface {
	ModificarZFEnDNS(context.Context, *Modificacion) (*RespuestaDNS, error)
}

// UnimplementedModificarZFEnDNSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModificarZFEnDNSServiceServer struct {
}

func (*UnimplementedModificarZFEnDNSServiceServer) ModificarZFEnDNS(context.Context, *Modificacion) (*RespuestaDNS, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModificarZFEnDNS not implemented")
}

func RegisterModificarZFEnDNSServiceServer(s *grpc.Server, srv ModificarZFEnDNSServiceServer) {
	s.RegisterService(&_ModificarZFEnDNSService_serviceDesc, srv)
}

func _ModificarZFEnDNSService_ModificarZFEnDNS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Modificacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModificarZFEnDNSServiceServer).ModificarZFEnDNS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModificarZF.ModificarZFEnDNSService/ModificarZFEnDNS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModificarZFEnDNSServiceServer).ModificarZFEnDNS(ctx, req.(*Modificacion))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModificarZFEnDNSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ModificarZF.ModificarZFEnDNSService",
	HandlerType: (*ModificarZFEnDNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModificarZFEnDNS",
			Handler:    _ModificarZFEnDNSService_ModificarZFEnDNS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ModificarZF.proto",
}

// VerificarServiceClient is the client API for VerificarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerificarServiceClient interface {
	Verificar(ctx context.Context, in *Enviodom, opts ...grpc.CallOption) (*RespuestaDNS, error)
}

type verificarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificarServiceClient(cc grpc.ClientConnInterface) VerificarServiceClient {
	return &verificarServiceClient{cc}
}

func (c *verificarServiceClient) Verificar(ctx context.Context, in *Enviodom, opts ...grpc.CallOption) (*RespuestaDNS, error) {
	out := new(RespuestaDNS)
	err := c.cc.Invoke(ctx, "/ModificarZF.VerificarService/Verificar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificarServiceServer is the server API for VerificarService service.
type VerificarServiceServer interface {
	Verificar(context.Context, *Enviodom) (*RespuestaDNS, error)
}

// UnimplementedVerificarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVerificarServiceServer struct {
}

func (*UnimplementedVerificarServiceServer) Verificar(context.Context, *Enviodom) (*RespuestaDNS, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verificar not implemented")
}

func RegisterVerificarServiceServer(s *grpc.Server, srv VerificarServiceServer) {
	s.RegisterService(&_VerificarService_serviceDesc, srv)
}

func _VerificarService_Verificar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enviodom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificarServiceServer).Verificar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModificarZF.VerificarService/Verificar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificarServiceServer).Verificar(ctx, req.(*Enviodom))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerificarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ModificarZF.VerificarService",
	HandlerType: (*VerificarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verificar",
			Handler:    _VerificarService_Verificar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ModificarZF.proto",
}
