// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: internal/protobuf/track.proto

package waves

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

type TransformerOptions_Aggregator int32

const (
	TransformerOptions_DEFAULT_AGGREGATOR TransformerOptions_Aggregator = 0
	TransformerOptions_AVG                TransformerOptions_Aggregator = 1
	TransformerOptions_RAVG               TransformerOptions_Aggregator = 2
	TransformerOptions_MAX                TransformerOptions_Aggregator = 3
	TransformerOptions_MS                 TransformerOptions_Aggregator = 4
	TransformerOptions_RMS                TransformerOptions_Aggregator = 5
)

// Enum value maps for TransformerOptions_Aggregator.
var (
	TransformerOptions_Aggregator_name = map[int32]string{
		0: "DEFAULT_AGGREGATOR",
		1: "AVG",
		2: "RAVG",
		3: "MAX",
		4: "MS",
		5: "RMS",
	}
	TransformerOptions_Aggregator_value = map[string]int32{
		"DEFAULT_AGGREGATOR": 0,
		"AVG":                1,
		"RAVG":               2,
		"MAX":                3,
		"MS":                 4,
		"RMS":                5,
	}
)

func (x TransformerOptions_Aggregator) Enum() *TransformerOptions_Aggregator {
	p := new(TransformerOptions_Aggregator)
	*p = x
	return p
}

func (x TransformerOptions_Aggregator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransformerOptions_Aggregator) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_track_proto_enumTypes[0].Descriptor()
}

func (TransformerOptions_Aggregator) Type() protoreflect.EnumType {
	return &file_internal_protobuf_track_proto_enumTypes[0]
}

func (x TransformerOptions_Aggregator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransformerOptions_Aggregator.Descriptor instead.
func (TransformerOptions_Aggregator) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{0, 0}
}

type TransformerOptions_DownsamplingMode int32

const (
	TransformerOptions_DEFAULT_DOWNSAMPLING TransformerOptions_DownsamplingMode = 0
	TransformerOptions_HEAD                 TransformerOptions_DownsamplingMode = 1
	TransformerOptions_CENTER               TransformerOptions_DownsamplingMode = 2
	TransformerOptions_TAIL                 TransformerOptions_DownsamplingMode = 3
	TransformerOptions_NONE                 TransformerOptions_DownsamplingMode = 4
)

// Enum value maps for TransformerOptions_DownsamplingMode.
var (
	TransformerOptions_DownsamplingMode_name = map[int32]string{
		0: "DEFAULT_DOWNSAMPLING",
		1: "HEAD",
		2: "CENTER",
		3: "TAIL",
		4: "NONE",
	}
	TransformerOptions_DownsamplingMode_value = map[string]int32{
		"DEFAULT_DOWNSAMPLING": 0,
		"HEAD":                 1,
		"CENTER":               2,
		"TAIL":                 3,
		"NONE":                 4,
	}
)

func (x TransformerOptions_DownsamplingMode) Enum() *TransformerOptions_DownsamplingMode {
	p := new(TransformerOptions_DownsamplingMode)
	*p = x
	return p
}

func (x TransformerOptions_DownsamplingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransformerOptions_DownsamplingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_track_proto_enumTypes[1].Descriptor()
}

func (TransformerOptions_DownsamplingMode) Type() protoreflect.EnumType {
	return &file_internal_protobuf_track_proto_enumTypes[1]
}

func (x TransformerOptions_DownsamplingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransformerOptions_DownsamplingMode.Descriptor instead.
func (TransformerOptions_DownsamplingMode) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{0, 1}
}

type BoxOptions_Alignment int32

const (
	BoxOptions_DEFAULT_ALIGNMENT BoxOptions_Alignment = 0
	BoxOptions_CENTER            BoxOptions_Alignment = 1
	BoxOptions_TOP               BoxOptions_Alignment = 2
	BoxOptions_BOTTOM            BoxOptions_Alignment = 3
)

// Enum value maps for BoxOptions_Alignment.
var (
	BoxOptions_Alignment_name = map[int32]string{
		0: "DEFAULT_ALIGNMENT",
		1: "CENTER",
		2: "TOP",
		3: "BOTTOM",
	}
	BoxOptions_Alignment_value = map[string]int32{
		"DEFAULT_ALIGNMENT": 0,
		"CENTER":            1,
		"TOP":               2,
		"BOTTOM":            3,
	}
)

func (x BoxOptions_Alignment) Enum() *BoxOptions_Alignment {
	p := new(BoxOptions_Alignment)
	*p = x
	return p
}

func (x BoxOptions_Alignment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoxOptions_Alignment) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_track_proto_enumTypes[2].Descriptor()
}

func (BoxOptions_Alignment) Type() protoreflect.EnumType {
	return &file_internal_protobuf_track_proto_enumTypes[2]
}

func (x BoxOptions_Alignment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoxOptions_Alignment.Descriptor instead.
func (BoxOptions_Alignment) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{1, 0}
}

type LineOptions_Interpolation int32

const (
	LineOptions_DEFAULT_INTERPOLATION LineOptions_Interpolation = 0
	LineOptions_FRITSCH_CARLSON       LineOptions_Interpolation = 1
	LineOptions_STEFFEN               LineOptions_Interpolation = 2
	LineOptions_NONE                  LineOptions_Interpolation = 3
)

// Enum value maps for LineOptions_Interpolation.
var (
	LineOptions_Interpolation_name = map[int32]string{
		0: "DEFAULT_INTERPOLATION",
		1: "FRITSCH_CARLSON",
		2: "STEFFEN",
		3: "NONE",
	}
	LineOptions_Interpolation_value = map[string]int32{
		"DEFAULT_INTERPOLATION": 0,
		"FRITSCH_CARLSON":       1,
		"STEFFEN":               2,
		"NONE":                  3,
	}
)

func (x LineOptions_Interpolation) Enum() *LineOptions_Interpolation {
	p := new(LineOptions_Interpolation)
	*p = x
	return p
}

func (x LineOptions_Interpolation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LineOptions_Interpolation) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_track_proto_enumTypes[3].Descriptor()
}

func (LineOptions_Interpolation) Type() protoreflect.EnumType {
	return &file_internal_protobuf_track_proto_enumTypes[3]
}

func (x LineOptions_Interpolation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LineOptions_Interpolation.Descriptor instead.
func (LineOptions_Interpolation) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{2, 0}
}

type TrackRequest_PainterType int32

const (
	TrackRequest_BOX  TrackRequest_PainterType = 0
	TrackRequest_LINE TrackRequest_PainterType = 1
)

// Enum value maps for TrackRequest_PainterType.
var (
	TrackRequest_PainterType_name = map[int32]string{
		0: "BOX",
		1: "LINE",
	}
	TrackRequest_PainterType_value = map[string]int32{
		"BOX":  0,
		"LINE": 1,
	}
)

func (x TrackRequest_PainterType) Enum() *TrackRequest_PainterType {
	p := new(TrackRequest_PainterType)
	*p = x
	return p
}

func (x TrackRequest_PainterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackRequest_PainterType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protobuf_track_proto_enumTypes[4].Descriptor()
}

func (TrackRequest_PainterType) Type() protoreflect.EnumType {
	return &file_internal_protobuf_track_proto_enumTypes[4]
}

func (x TrackRequest_PainterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackRequest_PainterType.Descriptor instead.
func (TrackRequest_PainterType) EnumDescriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{3, 0}
}

type TransformerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunks       int64                               `protobuf:"zigzag64,1,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Aggregator   TransformerOptions_Aggregator       `protobuf:"varint,2,opt,name=aggregator,proto3,enum=waves.TransformerOptions_Aggregator" json:"aggregator,omitempty"`
	Precision    int32                               `protobuf:"zigzag32,3,opt,name=precision,proto3" json:"precision,omitempty"` // TODO: this then needs range checking in the server to match the precisions supported
	Downsampling TransformerOptions_DownsamplingMode `protobuf:"varint,4,opt,name=downsampling,proto3,enum=waves.TransformerOptions_DownsamplingMode" json:"downsampling,omitempty"`
}

func (x *TransformerOptions) Reset() {
	*x = TransformerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformerOptions) ProtoMessage() {}

func (x *TransformerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformerOptions.ProtoReflect.Descriptor instead.
func (*TransformerOptions) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{0}
}

func (x *TransformerOptions) GetChunks() int64 {
	if x != nil {
		return x.Chunks
	}
	return 0
}

func (x *TransformerOptions) GetAggregator() TransformerOptions_Aggregator {
	if x != nil {
		return x.Aggregator
	}
	return TransformerOptions_DEFAULT_AGGREGATOR
}

func (x *TransformerOptions) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *TransformerOptions) GetDownsampling() TransformerOptions_DownsamplingMode {
	if x != nil {
		return x.Downsampling
	}
	return TransformerOptions_DEFAULT_DOWNSAMPLING
}

type BoxOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color     string               `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"` // fill color of the boxes
	Alignment BoxOptions_Alignment `protobuf:"varint,2,opt,name=alignment,proto3,enum=waves.BoxOptions_Alignment" json:"alignment,omitempty"`
	Height    float64              `protobuf:"fixed64,3,opt,name=height,proto3" json:"height,omitempty"`
	Width     float64              `protobuf:"fixed64,4,opt,name=width,proto3" json:"width,omitempty"`
	Rounded   float64              `protobuf:"fixed64,5,opt,name=rounded,proto3" json:"rounded,omitempty"`
	Gap       float64              `protobuf:"fixed64,6,opt,name=gap,proto3" json:"gap,omitempty"`
}

func (x *BoxOptions) Reset() {
	*x = BoxOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxOptions) ProtoMessage() {}

func (x *BoxOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxOptions.ProtoReflect.Descriptor instead.
func (*BoxOptions) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{1}
}

func (x *BoxOptions) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *BoxOptions) GetAlignment() BoxOptions_Alignment {
	if x != nil {
		return x.Alignment
	}
	return BoxOptions_DEFAULT_ALIGNMENT
}

func (x *BoxOptions) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BoxOptions) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BoxOptions) GetRounded() float64 {
	if x != nil {
		return x.Rounded
	}
	return 0
}

func (x *BoxOptions) GetGap() float64 {
	if x != nil {
		return x.Gap
	}
	return 0
}

type LineOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interpolation LineOptions_Interpolation `protobuf:"varint,1,opt,name=interpolation,proto3,enum=waves.LineOptions_Interpolation" json:"interpolation,omitempty"`
	Fill          string                    `protobuf:"bytes,2,opt,name=fill,proto3" json:"fill,omitempty"` // fill color of the path shape as string
	Stroke        *LineOptions_Stroke       `protobuf:"bytes,3,opt,name=stroke,proto3" json:"stroke,omitempty"`
	Inverted      bool                      `protobuf:"varint,4,opt,name=inverted,proto3" json:"inverted,omitempty"`
	Closed        bool                      `protobuf:"varint,5,opt,name=closed,proto3" json:"closed,omitempty"`
	Height        float64                   `protobuf:"fixed64,6,opt,name=height,proto3" json:"height,omitempty"`
	Width         float64                   `protobuf:"fixed64,7,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *LineOptions) Reset() {
	*x = LineOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineOptions) ProtoMessage() {}

func (x *LineOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineOptions.ProtoReflect.Descriptor instead.
func (*LineOptions) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{2}
}

func (x *LineOptions) GetInterpolation() LineOptions_Interpolation {
	if x != nil {
		return x.Interpolation
	}
	return LineOptions_DEFAULT_INTERPOLATION
}

func (x *LineOptions) GetFill() string {
	if x != nil {
		return x.Fill
	}
	return ""
}

func (x *LineOptions) GetStroke() *LineOptions_Stroke {
	if x != nil {
		return x.Stroke
	}
	return nil
}

func (x *LineOptions) GetInverted() bool {
	if x != nil {
		return x.Inverted
	}
	return false
}

func (x *LineOptions) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *LineOptions) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LineOptions) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

type TrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // the _id of the file inserted into a mongodb gridfs files collection
	WaveformType TrackRequest_PainterType `protobuf:"varint,3,opt,name=waveform_type,json=waveformType,proto3,enum=waves.TrackRequest_PainterType" json:"waveform_type,omitempty"`
	Transformer  *TransformerOptions      `protobuf:"bytes,4,opt,name=transformer,proto3" json:"transformer,omitempty"`
	// Types that are assignable to Painter:
	//	*TrackRequest_Line
	//	*TrackRequest_Box
	Painter isTrackRequest_Painter `protobuf_oneof:"painter"`
}

func (x *TrackRequest) Reset() {
	*x = TrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackRequest) ProtoMessage() {}

func (x *TrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackRequest.ProtoReflect.Descriptor instead.
func (*TrackRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{3}
}

func (x *TrackRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackRequest) GetWaveformType() TrackRequest_PainterType {
	if x != nil {
		return x.WaveformType
	}
	return TrackRequest_BOX
}

func (x *TrackRequest) GetTransformer() *TransformerOptions {
	if x != nil {
		return x.Transformer
	}
	return nil
}

func (m *TrackRequest) GetPainter() isTrackRequest_Painter {
	if m != nil {
		return m.Painter
	}
	return nil
}

func (x *TrackRequest) GetLine() *LineOptions {
	if x, ok := x.GetPainter().(*TrackRequest_Line); ok {
		return x.Line
	}
	return nil
}

func (x *TrackRequest) GetBox() *BoxOptions {
	if x, ok := x.GetPainter().(*TrackRequest_Box); ok {
		return x.Box
	}
	return nil
}

type isTrackRequest_Painter interface {
	isTrackRequest_Painter()
}

type TrackRequest_Line struct {
	Line *LineOptions `protobuf:"bytes,5,opt,name=line,proto3,oneof"`
}

type TrackRequest_Box struct {
	Box *BoxOptions `protobuf:"bytes,6,opt,name=box,proto3,oneof"`
}

func (*TrackRequest_Line) isTrackRequest_Painter() {}

func (*TrackRequest_Box) isTrackRequest_Painter() {}

type TrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Svg     string `protobuf:"bytes,2,opt,name=svg,proto3" json:"svg,omitempty"`
	Samples []byte `protobuf:"bytes,3,opt,name=samples,proto3" json:"samples,omitempty"`
}

func (x *TrackResponse) Reset() {
	*x = TrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackResponse) ProtoMessage() {}

func (x *TrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackResponse.ProtoReflect.Descriptor instead.
func (*TrackResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{4}
}

func (x *TrackResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackResponse) GetSvg() string {
	if x != nil {
		return x.Svg
	}
	return ""
}

func (x *TrackResponse) GetSamples() []byte {
	if x != nil {
		return x.Samples
	}
	return nil
}

type LineOptions_Stroke struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color string `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"` // stroke color
	Width string `protobuf:"bytes,2,opt,name=width,proto3" json:"width,omitempty"` // stroke width
}

func (x *LineOptions_Stroke) Reset() {
	*x = LineOptions_Stroke{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_track_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineOptions_Stroke) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineOptions_Stroke) ProtoMessage() {}

func (x *LineOptions_Stroke) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_track_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineOptions_Stroke.ProtoReflect.Descriptor instead.
func (*LineOptions_Stroke) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_track_proto_rawDescGZIP(), []int{2, 0}
}

func (x *LineOptions_Stroke) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *LineOptions_Stroke) GetWidth() string {
	if x != nil {
		return x.Width
	}
	return ""
}

var File_internal_protobuf_track_proto protoreflect.FileDescriptor

var file_internal_protobuf_track_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x77, 0x61, 0x76, 0x65, 0x73, 0x22, 0x8b, 0x03, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x77, 0x61, 0x76, 0x65,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x0a, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x56, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x41, 0x56, 0x47,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x4d,
	0x53, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4d, 0x53, 0x10, 0x05, 0x22, 0x56, 0x0a, 0x10,
	0x44, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x44, 0x4f, 0x57, 0x4e,
	0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45,
	0x41, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x04, 0x22, 0xfc, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x78, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x6c, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77,
	0x61, 0x76, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x78, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x6c, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x67, 0x61, 0x70, 0x22, 0x43,
	0x0a, 0x09, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4c, 0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x54, 0x54, 0x4f,
	0x4d, 0x10, 0x03, 0x22, 0x8c, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x6c, 0x12,
	0x31, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f,
	0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x1a, 0x34, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x56, 0x0a, 0x0d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x4f, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x52, 0x49, 0x54, 0x53, 0x43,
	0x48, 0x5f, 0x43, 0x41, 0x52, 0x4c, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x45, 0x46, 0x46, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x03, 0x22, 0x9f, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0d, 0x77, 0x61, 0x76, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x77, 0x61, 0x76,
	0x65, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x25, 0x0a, 0x03, 0x62, 0x6f, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x78, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x48, 0x00, 0x52, 0x03, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4f, 0x58, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x76, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x32, 0x40, 0x0a, 0x05, 0x57, 0x61, 0x76, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x57, 0x61,
	0x76, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x13, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x77, 0x61,
	0x76, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x6f, 0x6f, 0x6d, 0x6f, 0x69, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2d,
	0x64, 0x62, 0x2f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_track_proto_rawDescOnce sync.Once
	file_internal_protobuf_track_proto_rawDescData = file_internal_protobuf_track_proto_rawDesc
)

func file_internal_protobuf_track_proto_rawDescGZIP() []byte {
	file_internal_protobuf_track_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_track_proto_rawDescData)
	})
	return file_internal_protobuf_track_proto_rawDescData
}

var file_internal_protobuf_track_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_internal_protobuf_track_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_protobuf_track_proto_goTypes = []interface{}{
	(TransformerOptions_Aggregator)(0),       // 0: waves.TransformerOptions.Aggregator
	(TransformerOptions_DownsamplingMode)(0), // 1: waves.TransformerOptions.DownsamplingMode
	(BoxOptions_Alignment)(0),                // 2: waves.BoxOptions.Alignment
	(LineOptions_Interpolation)(0),           // 3: waves.LineOptions.Interpolation
	(TrackRequest_PainterType)(0),            // 4: waves.TrackRequest.PainterType
	(*TransformerOptions)(nil),               // 5: waves.TransformerOptions
	(*BoxOptions)(nil),                       // 6: waves.BoxOptions
	(*LineOptions)(nil),                      // 7: waves.LineOptions
	(*TrackRequest)(nil),                     // 8: waves.TrackRequest
	(*TrackResponse)(nil),                    // 9: waves.TrackResponse
	(*LineOptions_Stroke)(nil),               // 10: waves.LineOptions.Stroke
}
var file_internal_protobuf_track_proto_depIdxs = []int32{
	0,  // 0: waves.TransformerOptions.aggregator:type_name -> waves.TransformerOptions.Aggregator
	1,  // 1: waves.TransformerOptions.downsampling:type_name -> waves.TransformerOptions.DownsamplingMode
	2,  // 2: waves.BoxOptions.alignment:type_name -> waves.BoxOptions.Alignment
	3,  // 3: waves.LineOptions.interpolation:type_name -> waves.LineOptions.Interpolation
	10, // 4: waves.LineOptions.stroke:type_name -> waves.LineOptions.Stroke
	4,  // 5: waves.TrackRequest.waveform_type:type_name -> waves.TrackRequest.PainterType
	5,  // 6: waves.TrackRequest.transformer:type_name -> waves.TransformerOptions
	7,  // 7: waves.TrackRequest.line:type_name -> waves.LineOptions
	6,  // 8: waves.TrackRequest.box:type_name -> waves.BoxOptions
	8,  // 9: waves.Waves.Waveform:input_type -> waves.TrackRequest
	9,  // 10: waves.Waves.Waveform:output_type -> waves.TrackResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_internal_protobuf_track_proto_init() }
func file_internal_protobuf_track_proto_init() {
	if File_internal_protobuf_track_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_track_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformerOptions); i {
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
		file_internal_protobuf_track_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxOptions); i {
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
		file_internal_protobuf_track_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineOptions); i {
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
		file_internal_protobuf_track_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackRequest); i {
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
		file_internal_protobuf_track_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackResponse); i {
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
		file_internal_protobuf_track_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineOptions_Stroke); i {
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
	file_internal_protobuf_track_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TrackRequest_Line)(nil),
		(*TrackRequest_Box)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_protobuf_track_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_track_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_track_proto_depIdxs,
		EnumInfos:         file_internal_protobuf_track_proto_enumTypes,
		MessageInfos:      file_internal_protobuf_track_proto_msgTypes,
	}.Build()
	File_internal_protobuf_track_proto = out.File
	file_internal_protobuf_track_proto_rawDesc = nil
	file_internal_protobuf_track_proto_goTypes = nil
	file_internal_protobuf_track_proto_depIdxs = nil
}
