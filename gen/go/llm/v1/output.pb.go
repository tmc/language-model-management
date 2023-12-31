// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: llm/v1/output.proto

package llmv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Output represents a single output from a generation.
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the output.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The actual content produced by the model.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Associated rating for this output, if available.
	Rating *OutputRating `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
	// Timestamp indicating when the output was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_output_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_output_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_llm_v1_output_proto_rawDescGZIP(), []int{0}
}

func (x *Output) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Output) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Output) GetRating() *OutputRating {
	if x != nil {
		return x.Rating
	}
	return nil
}

func (x *Output) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// OutputRating represents the evaluation or assessment of an output.
type OutputRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Numeric score representing the quality or other rating dimension.
	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	// Details about the entity providing the rating.
	Rater *Rater `protobuf:"bytes,2,opt,name=rater,proto3" json:"rater,omitempty"`
	// Additional comments or feedback provided during the rating.
	Feedback string `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	// Timestamp indicating when the rating was provided.
	RatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=rated_at,json=ratedAt,proto3" json:"rated_at,omitempty"`
}

func (x *OutputRating) Reset() {
	*x = OutputRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_v1_output_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputRating) ProtoMessage() {}

func (x *OutputRating) ProtoReflect() protoreflect.Message {
	mi := &file_llm_v1_output_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputRating.ProtoReflect.Descriptor instead.
func (*OutputRating) Descriptor() ([]byte, []int) {
	return file_llm_v1_output_proto_rawDescGZIP(), []int{1}
}

func (x *OutputRating) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *OutputRating) GetRater() *Rater {
	if x != nil {
		return x.Rater
	}
	return nil
}

func (x *OutputRating) GetFeedback() string {
	if x != nil {
		return x.Feedback
	}
	return ""
}

func (x *OutputRating) GetRatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RatedAt
	}
	return nil
}

var File_llm_v1_output_proto protoreflect.FileDescriptor

var file_llm_v1_output_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x9c, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x72, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x90, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6d, 0x63, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6c, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x4c, 0x6c, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x4c, 0x6c,
	0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4c, 0x6c, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4c, 0x6c, 0x6d, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_v1_output_proto_rawDescOnce sync.Once
	file_llm_v1_output_proto_rawDescData = file_llm_v1_output_proto_rawDesc
)

func file_llm_v1_output_proto_rawDescGZIP() []byte {
	file_llm_v1_output_proto_rawDescOnce.Do(func() {
		file_llm_v1_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_v1_output_proto_rawDescData)
	})
	return file_llm_v1_output_proto_rawDescData
}

var file_llm_v1_output_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_llm_v1_output_proto_goTypes = []interface{}{
	(*Output)(nil),                // 0: llm.v1.Output
	(*OutputRating)(nil),          // 1: llm.v1.OutputRating
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Rater)(nil),                 // 3: llm.v1.Rater
}
var file_llm_v1_output_proto_depIdxs = []int32{
	1, // 0: llm.v1.Output.rating:type_name -> llm.v1.OutputRating
	2, // 1: llm.v1.Output.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: llm.v1.OutputRating.rater:type_name -> llm.v1.Rater
	2, // 3: llm.v1.OutputRating.rated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_llm_v1_output_proto_init() }
func file_llm_v1_output_proto_init() {
	if File_llm_v1_output_proto != nil {
		return
	}
	file_llm_v1_prompt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_llm_v1_output_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_llm_v1_output_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputRating); i {
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
			RawDescriptor: file_llm_v1_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_llm_v1_output_proto_goTypes,
		DependencyIndexes: file_llm_v1_output_proto_depIdxs,
		MessageInfos:      file_llm_v1_output_proto_msgTypes,
	}.Build()
	File_llm_v1_output_proto = out.File
	file_llm_v1_output_proto_rawDesc = nil
	file_llm_v1_output_proto_goTypes = nil
	file_llm_v1_output_proto_depIdxs = nil
}
