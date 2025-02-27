// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: proposer.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProposerPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *ProposerPublicKeyRequest) Reset() {
	*x = ProposerPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposerPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposerPublicKeyRequest) ProtoMessage() {}

func (x *ProposerPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proposer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposerPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ProposerPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_proposer_proto_rawDescGZIP(), []int{0}
}

func (x *ProposerPublicKeyRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type ProposeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProposeResponse) Reset() {
	*x = ProposeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeResponse) ProtoMessage() {}

func (x *ProposeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proposer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeResponse.ProtoReflect.Descriptor instead.
func (*ProposeResponse) Descriptor() ([]byte, []int) {
	return file_proposer_proto_rawDescGZIP(), []int{1}
}

func (x *ProposeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proposer_proto protoreflect.FileDescriptor

var file_proposer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x63, 0x75, 0x72, 0x69, 0x65, 0x1a, 0x11, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc4,
	0x02, 0x0a, 0x11, 0x43, 0x75, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x69, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x69, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x46, 0x6f, 0x72, 0x4f, 0x47, 0x12, 0x1c, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x65, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x75, 0x72, 0x69, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x46, 0x6f, 0x72, 0x4f, 0x47, 0x1a, 0x16, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x69, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x4e, 0x47, 0x12, 0x1c, 0x2e, 0x63, 0x75, 0x72, 0x69,
	0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x75, 0x72, 0x69, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x46, 0x6f, 0x72, 0x4e, 0x47, 0x1a, 0x16, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x63, 0x75, 0x72,
	0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x75,
	0x72, 0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x66, 0x6c, 0x61, 0x67, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proposer_proto_rawDescOnce sync.Once
	file_proposer_proto_rawDescData = file_proposer_proto_rawDesc
)

func file_proposer_proto_rawDescGZIP() []byte {
	file_proposer_proto_rawDescOnce.Do(func() {
		file_proposer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proposer_proto_rawDescData)
	})
	return file_proposer_proto_rawDescData
}

var file_proposer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proposer_proto_goTypes = []interface{}{
	(*ProposerPublicKeyRequest)(nil), // 0: curie.ProposerPublicKeyRequest
	(*ProposeResponse)(nil),          // 1: curie.ProposeResponse
	(*emptypb.Empty)(nil),            // 2: google.protobuf.Empty
	(*SignedCurieBlockForOG)(nil),    // 3: curie.SignedCurieBlockForOG
	(*SignedCurieBlockForNG)(nil),    // 4: curie.SignedCurieBlockForNG
	(*CurieBlock)(nil),               // 5: curie.CurieBlock
}
var file_proposer_proto_depIdxs = []int32{
	2, // 0: curie.CurieNodeProposer.GetBlock:input_type -> google.protobuf.Empty
	3, // 1: curie.CurieNodeProposer.ProposeCurieBlockForOG:input_type -> curie.SignedCurieBlockForOG
	4, // 2: curie.CurieNodeProposer.ProposeCurieBlockForNG:input_type -> curie.SignedCurieBlockForNG
	0, // 3: curie.CurieNodeProposer.SendProposerPublicKey:input_type -> curie.ProposerPublicKeyRequest
	5, // 4: curie.CurieNodeProposer.GetBlock:output_type -> curie.CurieBlock
	1, // 5: curie.CurieNodeProposer.ProposeCurieBlockForOG:output_type -> curie.ProposeResponse
	1, // 6: curie.CurieNodeProposer.ProposeCurieBlockForNG:output_type -> curie.ProposeResponse
	1, // 7: curie.CurieNodeProposer.SendProposerPublicKey:output_type -> curie.ProposeResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proposer_proto_init() }
func file_proposer_proto_init() {
	if File_proposer_proto != nil {
		return
	}
	file_gossip_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proposer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposerPublicKeyRequest); i {
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
		file_proposer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeResponse); i {
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
			RawDescriptor: file_proposer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proposer_proto_goTypes,
		DependencyIndexes: file_proposer_proto_depIdxs,
		MessageInfos:      file_proposer_proto_msgTypes,
	}.Build()
	File_proposer_proto = out.File
	file_proposer_proto_rawDesc = nil
	file_proposer_proto_goTypes = nil
	file_proposer_proto_depIdxs = nil
}
