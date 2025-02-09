// LINT: LEGACY_NAMES

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: tensorflow/stream_executor/dnn.proto

package stream_executor

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

// Specifies the data type used by an operation.
type DataType int32

const (
	DataType_kFloat         DataType = 0
	DataType_kDouble        DataType = 1
	DataType_kHalf          DataType = 2
	DataType_kInt8          DataType = 3
	DataType_kInt32         DataType = 4
	DataType_kComplexFloat  DataType = 5
	DataType_kComplexDouble DataType = 6
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "kFloat",
		1: "kDouble",
		2: "kHalf",
		3: "kInt8",
		4: "kInt32",
		5: "kComplexFloat",
		6: "kComplexDouble",
	}
	DataType_value = map[string]int32{
		"kFloat":         0,
		"kDouble":        1,
		"kHalf":          2,
		"kInt8":          3,
		"kInt32":         4,
		"kComplexFloat":  5,
		"kComplexDouble": 6,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{0}
}

// Describes how a convolution input or output layer's data is formatted.
type DataLayout int32

const (
	// Naming convention:
	// Y <-> row or height
	// X <-> column or width
	// Batch <-> batch, or N
	// Depth <-> feature, or channel
	// TODO(timshen): turn them into cuDNN names, e.g. kNCHW.
	DataLayout_kYXDepthBatch  DataLayout = 0
	DataLayout_kYXBatchDepth  DataLayout = 1
	DataLayout_kBatchYXDepth  DataLayout = 2 // cuDNN's NHWC layout
	DataLayout_kBatchDepthYX  DataLayout = 3 // cuDNN's NCHW layout
	DataLayout_kBatchDepthYX4 DataLayout = 4 // cuDNN's NCHW_VECT_C layout
)

// Enum value maps for DataLayout.
var (
	DataLayout_name = map[int32]string{
		0: "kYXDepthBatch",
		1: "kYXBatchDepth",
		2: "kBatchYXDepth",
		3: "kBatchDepthYX",
		4: "kBatchDepthYX4",
	}
	DataLayout_value = map[string]int32{
		"kYXDepthBatch":  0,
		"kYXBatchDepth":  1,
		"kBatchYXDepth":  2,
		"kBatchDepthYX":  3,
		"kBatchDepthYX4": 4,
	}
)

func (x DataLayout) Enum() *DataLayout {
	p := new(DataLayout)
	*p = x
	return p
}

func (x DataLayout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataLayout) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[1].Descriptor()
}

func (DataLayout) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[1]
}

func (x DataLayout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataLayout.Descriptor instead.
func (DataLayout) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{1}
}

// Describes how a convolution filter is laid out in the memory.
type FilterLayout int32

const (
	// Naming convention:
	// Y <-> row or height
	// X <-> column or width
	// Output <-> output feature, or N
	// Input <-> input feature, or N
	// TODO(timshen): turn them into cuDNN names, e.g. kNCHW.
	FilterLayout_kOutputInputYX  FilterLayout = 0 // cuDNN's NCHW layout
	FilterLayout_kOutputYXInput  FilterLayout = 1 // cuDNN's NHWC layout
	FilterLayout_kOutputInputYX4 FilterLayout = 2 // cuDNN's NCHW_VECT_C layout
	FilterLayout_kInputYXOutput  FilterLayout = 3
	FilterLayout_kYXInputOutput  FilterLayout = 4
)

// Enum value maps for FilterLayout.
var (
	FilterLayout_name = map[int32]string{
		0: "kOutputInputYX",
		1: "kOutputYXInput",
		2: "kOutputInputYX4",
		3: "kInputYXOutput",
		4: "kYXInputOutput",
	}
	FilterLayout_value = map[string]int32{
		"kOutputInputYX":  0,
		"kOutputYXInput":  1,
		"kOutputInputYX4": 2,
		"kInputYXOutput":  3,
		"kYXInputOutput":  4,
	}
)

func (x FilterLayout) Enum() *FilterLayout {
	p := new(FilterLayout)
	*p = x
	return p
}

func (x FilterLayout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterLayout) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[2].Descriptor()
}

func (FilterLayout) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[2]
}

func (x FilterLayout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterLayout.Descriptor instead.
func (FilterLayout) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{2}
}

// Describes a kind of non-linearity (threshold-like mathematical function).
type ActivationMode int32

const (
	ActivationMode_kNone    ActivationMode = 0
	ActivationMode_kSigmoid ActivationMode = 1
	// Rectified linear activation: f(x) = x < 0 ? 0 : x
	ActivationMode_kRelu ActivationMode = 2
	// Rectified linear activation; where upper maximum is 6.0.
	ActivationMode_kRelu6 ActivationMode = 3
	// Rectified linear activation; where upper maximum specified by
	// BatchDescriptor::value_max().
	ActivationMode_kReluX ActivationMode = 4
	ActivationMode_kTanh  ActivationMode = 5
	// Like ReluX; but passes all values in the range [-X,X].
	ActivationMode_kBandPass ActivationMode = 6
)

// Enum value maps for ActivationMode.
var (
	ActivationMode_name = map[int32]string{
		0: "kNone",
		1: "kSigmoid",
		2: "kRelu",
		3: "kRelu6",
		4: "kReluX",
		5: "kTanh",
		6: "kBandPass",
	}
	ActivationMode_value = map[string]int32{
		"kNone":     0,
		"kSigmoid":  1,
		"kRelu":     2,
		"kRelu6":    3,
		"kReluX":    4,
		"kTanh":     5,
		"kBandPass": 6,
	}
)

func (x ActivationMode) Enum() *ActivationMode {
	p := new(ActivationMode)
	*p = x
	return p
}

func (x ActivationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[3].Descriptor()
}

func (ActivationMode) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[3]
}

func (x ActivationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivationMode.Descriptor instead.
func (ActivationMode) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{3}
}

// Describe the math definition for the conv op. The popular behavior is
// actually called cross-correlation in math, despite the operation is often
// referred as convolution. See cuDNN cudnnConvolutionMode_t.
type ConvolutionMode int32

const (
	ConvolutionMode_CROSS_CORRELATION ConvolutionMode = 0
	ConvolutionMode_CONVOLUTION       ConvolutionMode = 1
)

// Enum value maps for ConvolutionMode.
var (
	ConvolutionMode_name = map[int32]string{
		0: "CROSS_CORRELATION",
		1: "CONVOLUTION",
	}
	ConvolutionMode_value = map[string]int32{
		"CROSS_CORRELATION": 0,
		"CONVOLUTION":       1,
	}
)

func (x ConvolutionMode) Enum() *ConvolutionMode {
	p := new(ConvolutionMode)
	*p = x
	return p
}

func (x ConvolutionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConvolutionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[4].Descriptor()
}

func (ConvolutionMode) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[4]
}

func (x ConvolutionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConvolutionMode.Descriptor instead.
func (ConvolutionMode) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{4}
}

type ConvolutionKind int32

const (
	ConvolutionKind_INVALID                 ConvolutionKind = 0
	ConvolutionKind_FORWARD                 ConvolutionKind = 1
	ConvolutionKind_BACKWARD_FILTER         ConvolutionKind = 2
	ConvolutionKind_BACKWARD_DATA           ConvolutionKind = 3
	ConvolutionKind_FORWARD_BIAS_ACTIVATION ConvolutionKind = 4
)

// Enum value maps for ConvolutionKind.
var (
	ConvolutionKind_name = map[int32]string{
		0: "INVALID",
		1: "FORWARD",
		2: "BACKWARD_FILTER",
		3: "BACKWARD_DATA",
		4: "FORWARD_BIAS_ACTIVATION",
	}
	ConvolutionKind_value = map[string]int32{
		"INVALID":                 0,
		"FORWARD":                 1,
		"BACKWARD_FILTER":         2,
		"BACKWARD_DATA":           3,
		"FORWARD_BIAS_ACTIVATION": 4,
	}
)

func (x ConvolutionKind) Enum() *ConvolutionKind {
	p := new(ConvolutionKind)
	*p = x
	return p
}

func (x ConvolutionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConvolutionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[5].Descriptor()
}

func (ConvolutionKind) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[5]
}

func (x ConvolutionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConvolutionKind.Descriptor instead.
func (ConvolutionKind) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{5}
}

type AlgorithmProto_MathType int32

const (
	AlgorithmProto_DEFAULT_MATH AlgorithmProto_MathType = 0
	// The GPU may operate 4x4 matrix FMA.
	// See cuDNN's documentation for CUDNN_TENSOR_OP_MATH.
	AlgorithmProto_TENSOR_OP_MATH AlgorithmProto_MathType = 1
)

// Enum value maps for AlgorithmProto_MathType.
var (
	AlgorithmProto_MathType_name = map[int32]string{
		0: "DEFAULT_MATH",
		1: "TENSOR_OP_MATH",
	}
	AlgorithmProto_MathType_value = map[string]int32{
		"DEFAULT_MATH":   0,
		"TENSOR_OP_MATH": 1,
	}
)

func (x AlgorithmProto_MathType) Enum() *AlgorithmProto_MathType {
	p := new(AlgorithmProto_MathType)
	*p = x
	return p
}

func (x AlgorithmProto_MathType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlgorithmProto_MathType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_stream_executor_dnn_proto_enumTypes[6].Descriptor()
}

func (AlgorithmProto_MathType) Type() protoreflect.EnumType {
	return &file_tensorflow_stream_executor_dnn_proto_enumTypes[6]
}

func (x AlgorithmProto_MathType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlgorithmProto_MathType.Descriptor instead.
func (AlgorithmProto_MathType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{1, 0}
}

// Generic tensor representation.
type TensorDescriptorProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dimensions []int64  `protobuf:"varint,1,rep,packed,name=dimensions,proto3" json:"dimensions,omitempty"`
	DataType   DataType `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=stream_executor.dnn.DataType" json:"data_type,omitempty"`
	// Types that are assignable to LayoutOneof:
	//	*TensorDescriptorProto_DataLayout
	//	*TensorDescriptorProto_FilterLayout
	LayoutOneof isTensorDescriptorProto_LayoutOneof `protobuf_oneof:"layout_oneof"`
}

func (x *TensorDescriptorProto) Reset() {
	*x = TensorDescriptorProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorDescriptorProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorDescriptorProto) ProtoMessage() {}

func (x *TensorDescriptorProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorDescriptorProto.ProtoReflect.Descriptor instead.
func (*TensorDescriptorProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{0}
}

func (x *TensorDescriptorProto) GetDimensions() []int64 {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *TensorDescriptorProto) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_kFloat
}

func (m *TensorDescriptorProto) GetLayoutOneof() isTensorDescriptorProto_LayoutOneof {
	if m != nil {
		return m.LayoutOneof
	}
	return nil
}

func (x *TensorDescriptorProto) GetDataLayout() DataLayout {
	if x, ok := x.GetLayoutOneof().(*TensorDescriptorProto_DataLayout); ok {
		return x.DataLayout
	}
	return DataLayout_kYXDepthBatch
}

func (x *TensorDescriptorProto) GetFilterLayout() FilterLayout {
	if x, ok := x.GetLayoutOneof().(*TensorDescriptorProto_FilterLayout); ok {
		return x.FilterLayout
	}
	return FilterLayout_kOutputInputYX
}

type isTensorDescriptorProto_LayoutOneof interface {
	isTensorDescriptorProto_LayoutOneof()
}

type TensorDescriptorProto_DataLayout struct {
	DataLayout DataLayout `protobuf:"varint,3,opt,name=data_layout,json=dataLayout,proto3,enum=stream_executor.dnn.DataLayout,oneof"`
}

type TensorDescriptorProto_FilterLayout struct {
	FilterLayout FilterLayout `protobuf:"varint,4,opt,name=filter_layout,json=filterLayout,proto3,enum=stream_executor.dnn.FilterLayout,oneof"`
}

func (*TensorDescriptorProto_DataLayout) isTensorDescriptorProto_LayoutOneof() {}

func (*TensorDescriptorProto_FilterLayout) isTensorDescriptorProto_LayoutOneof() {}

// Generic algorithm representation.
type AlgorithmProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlgoId   int64                   `protobuf:"varint,1,opt,name=algo_id,json=algoId,proto3" json:"algo_id,omitempty"`
	MathType AlgorithmProto_MathType `protobuf:"varint,2,opt,name=math_type,json=mathType,proto3,enum=stream_executor.dnn.AlgorithmProto_MathType" json:"math_type,omitempty"`
	// cuDNN v8 uses a string to uniquely represent the backend plan.
	ExecPlanId string `protobuf:"bytes,3,opt,name=exec_plan_id,json=execPlanId,proto3" json:"exec_plan_id,omitempty"`
}

func (x *AlgorithmProto) Reset() {
	*x = AlgorithmProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmProto) ProtoMessage() {}

func (x *AlgorithmProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmProto.ProtoReflect.Descriptor instead.
func (*AlgorithmProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{1}
}

func (x *AlgorithmProto) GetAlgoId() int64 {
	if x != nil {
		return x.AlgoId
	}
	return 0
}

func (x *AlgorithmProto) GetMathType() AlgorithmProto_MathType {
	if x != nil {
		return x.MathType
	}
	return AlgorithmProto_DEFAULT_MATH
}

func (x *AlgorithmProto) GetExecPlanId() string {
	if x != nil {
		return x.ExecPlanId
	}
	return ""
}

// Convolution-specific parameters.
type ConvolutionDescriptorProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paddings  []int64 `protobuf:"varint,1,rep,packed,name=paddings,proto3" json:"paddings,omitempty"`
	Strides   []int64 `protobuf:"varint,2,rep,packed,name=strides,proto3" json:"strides,omitempty"`
	Dilations []int64 `protobuf:"varint,3,rep,packed,name=dilations,proto3" json:"dilations,omitempty"`
	// The "accumulator" type. For example, use F32 as an accumulator for F16
	// convolutions.
	// See cuDNN's cudnnConvolutionMode_t.
	ComputeMode DataType `protobuf:"varint,4,opt,name=compute_mode,json=computeMode,proto3,enum=stream_executor.dnn.DataType" json:"compute_mode,omitempty"`
	// See cuDNN's group count.
	GroupCount      int32           `protobuf:"varint,5,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	ConvolutionMode ConvolutionMode `protobuf:"varint,6,opt,name=convolution_mode,json=convolutionMode,proto3,enum=stream_executor.dnn.ConvolutionMode" json:"convolution_mode,omitempty"`
	// Tensorflow node name, same as in NodeDef, for debugging purposes.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConvolutionDescriptorProto) Reset() {
	*x = ConvolutionDescriptorProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvolutionDescriptorProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvolutionDescriptorProto) ProtoMessage() {}

func (x *ConvolutionDescriptorProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_stream_executor_dnn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvolutionDescriptorProto.ProtoReflect.Descriptor instead.
func (*ConvolutionDescriptorProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_stream_executor_dnn_proto_rawDescGZIP(), []int{2}
}

func (x *ConvolutionDescriptorProto) GetPaddings() []int64 {
	if x != nil {
		return x.Paddings
	}
	return nil
}

func (x *ConvolutionDescriptorProto) GetStrides() []int64 {
	if x != nil {
		return x.Strides
	}
	return nil
}

func (x *ConvolutionDescriptorProto) GetDilations() []int64 {
	if x != nil {
		return x.Dilations
	}
	return nil
}

func (x *ConvolutionDescriptorProto) GetComputeMode() DataType {
	if x != nil {
		return x.ComputeMode
	}
	return DataType_kFloat
}

func (x *ConvolutionDescriptorProto) GetGroupCount() int32 {
	if x != nil {
		return x.GroupCount
	}
	return 0
}

func (x *ConvolutionDescriptorProto) GetConvolutionMode() ConvolutionMode {
	if x != nil {
		return x.ConvolutionMode
	}
	return ConvolutionMode_CROSS_CORRELATION
}

func (x *ConvolutionDescriptorProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tensorflow_stream_executor_dnn_proto protoreflect.FileDescriptor

var file_tensorflow_stream_executor_dnn_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x6e, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x22, 0x91, 0x02, 0x0a, 0x15,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x48, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64,
	0x6e, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x42,
	0x0e, 0x0a, 0x0c, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22,
	0xc8, 0x01, 0x0a, 0x0e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x6c, 0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6c, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6d, 0x61,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x65, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f,
	0x4d, 0x41, 0x54, 0x48, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x45, 0x4e, 0x53, 0x4f, 0x52,
	0x5f, 0x4f, 0x50, 0x5f, 0x4d, 0x41, 0x54, 0x48, 0x10, 0x01, 0x22, 0xb8, 0x02, 0x0a, 0x1a, 0x43,
	0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x09, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x4f, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x6c, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x6b, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x6b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x48,
	0x61, 0x6c, 0x66, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x49, 0x6e, 0x74, 0x38, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x6b, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d,
	0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x05, 0x12,
	0x12, 0x0a, 0x0e, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x10, 0x06, 0x2a, 0x6c, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x59, 0x58, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x59, 0x58, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x59, 0x58, 0x44, 0x65, 0x70, 0x74, 0x68, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x74, 0x68, 0x59, 0x58, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x74, 0x68, 0x59, 0x58, 0x34, 0x10,
	0x04, 0x2a, 0x73, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x0e, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x59, 0x58, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x59, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x59, 0x58, 0x34, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x59, 0x58, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x6b, 0x59, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x10, 0x04, 0x2a, 0x66, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x4e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x53, 0x69, 0x67, 0x6d, 0x6f, 0x69, 0x64, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x52, 0x65, 0x6c, 0x75, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x6b, 0x52, 0x65, 0x6c, 0x75, 0x36, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x6b, 0x52, 0x65, 0x6c,
	0x75, 0x58, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x54, 0x61, 0x6e, 0x68, 0x10, 0x05, 0x12,
	0x0d, 0x0a, 0x09, 0x6b, 0x42, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x10, 0x06, 0x2a, 0x39,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x4f, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x52, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x56,
	0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x2a, 0x70, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4f, 0x52,
	0x57, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41,
	0x52, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x42,
	0x41, 0x43, 0x4b, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x12, 0x1b,
	0x0a, 0x17, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x42, 0x49, 0x41, 0x53, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_stream_executor_dnn_proto_rawDescOnce sync.Once
	file_tensorflow_stream_executor_dnn_proto_rawDescData = file_tensorflow_stream_executor_dnn_proto_rawDesc
)

func file_tensorflow_stream_executor_dnn_proto_rawDescGZIP() []byte {
	file_tensorflow_stream_executor_dnn_proto_rawDescOnce.Do(func() {
		file_tensorflow_stream_executor_dnn_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_stream_executor_dnn_proto_rawDescData)
	})
	return file_tensorflow_stream_executor_dnn_proto_rawDescData
}

var file_tensorflow_stream_executor_dnn_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_tensorflow_stream_executor_dnn_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_stream_executor_dnn_proto_goTypes = []interface{}{
	(DataType)(0),                      // 0: stream_executor.dnn.DataType
	(DataLayout)(0),                    // 1: stream_executor.dnn.DataLayout
	(FilterLayout)(0),                  // 2: stream_executor.dnn.FilterLayout
	(ActivationMode)(0),                // 3: stream_executor.dnn.ActivationMode
	(ConvolutionMode)(0),               // 4: stream_executor.dnn.ConvolutionMode
	(ConvolutionKind)(0),               // 5: stream_executor.dnn.ConvolutionKind
	(AlgorithmProto_MathType)(0),       // 6: stream_executor.dnn.AlgorithmProto.MathType
	(*TensorDescriptorProto)(nil),      // 7: stream_executor.dnn.TensorDescriptorProto
	(*AlgorithmProto)(nil),             // 8: stream_executor.dnn.AlgorithmProto
	(*ConvolutionDescriptorProto)(nil), // 9: stream_executor.dnn.ConvolutionDescriptorProto
}
var file_tensorflow_stream_executor_dnn_proto_depIdxs = []int32{
	0, // 0: stream_executor.dnn.TensorDescriptorProto.data_type:type_name -> stream_executor.dnn.DataType
	1, // 1: stream_executor.dnn.TensorDescriptorProto.data_layout:type_name -> stream_executor.dnn.DataLayout
	2, // 2: stream_executor.dnn.TensorDescriptorProto.filter_layout:type_name -> stream_executor.dnn.FilterLayout
	6, // 3: stream_executor.dnn.AlgorithmProto.math_type:type_name -> stream_executor.dnn.AlgorithmProto.MathType
	0, // 4: stream_executor.dnn.ConvolutionDescriptorProto.compute_mode:type_name -> stream_executor.dnn.DataType
	4, // 5: stream_executor.dnn.ConvolutionDescriptorProto.convolution_mode:type_name -> stream_executor.dnn.ConvolutionMode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_stream_executor_dnn_proto_init() }
func file_tensorflow_stream_executor_dnn_proto_init() {
	if File_tensorflow_stream_executor_dnn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_stream_executor_dnn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorDescriptorProto); i {
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
		file_tensorflow_stream_executor_dnn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmProto); i {
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
		file_tensorflow_stream_executor_dnn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvolutionDescriptorProto); i {
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
	file_tensorflow_stream_executor_dnn_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TensorDescriptorProto_DataLayout)(nil),
		(*TensorDescriptorProto_FilterLayout)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_stream_executor_dnn_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_stream_executor_dnn_proto_goTypes,
		DependencyIndexes: file_tensorflow_stream_executor_dnn_proto_depIdxs,
		EnumInfos:         file_tensorflow_stream_executor_dnn_proto_enumTypes,
		MessageInfos:      file_tensorflow_stream_executor_dnn_proto_msgTypes,
	}.Build()
	File_tensorflow_stream_executor_dnn_proto = out.File
	file_tensorflow_stream_executor_dnn_proto_rawDesc = nil
	file_tensorflow_stream_executor_dnn_proto_goTypes = nil
	file_tensorflow_stream_executor_dnn_proto_depIdxs = nil
}
