package grpc

import (
	"math"

	pb "github.com/zoomoid/track-db/waves/internal/protobuf"
	"github.com/zoomoid/track-db/waves/internal/runner"
	"github.com/zoomoid/waveman2/pkg/painter/box"
	"github.com/zoomoid/waveman2/pkg/painter/line"
	"github.com/zoomoid/waveman2/pkg/transform"
)

func ConvertProtobufToWavemanSpec(in *pb.TrackRequest) *runner.RunnerOptions {
	transformerOptions := toTransformerOptions(in.GetTransformer())
	r := &runner.RunnerOptions{
		Transformer: transformerOptions,
	}

	switch p := in.Painter.(type) {
	case *pb.TrackRequest_Line:
		lp := toLinePainter(p)
		r.LinePainter = lp
	case *pb.TrackRequest_Box:
		bp := toBoxPainter(p)
		r.BoxPainter = bp
	}

	return r
}

func toTransformerOptions(opts *pb.TransformerOptions) *transform.ReaderOptions {
	dm := transform.DownsamplingEmpty
	switch opts.GetDownsampling() {
	case pb.TransformerOptions_DEFAULT_DOWNSAMPLING:
		dm = transform.DefaultDownsamplingMode
	case pb.TransformerOptions_HEAD:
		dm = transform.DownsamplingHead
	case pb.TransformerOptions_CENTER:
		dm = transform.DownsamplingCenter
	case pb.TransformerOptions_TAIL:
		dm = transform.DownsamplingTail
	case pb.TransformerOptions_NONE:
		dm = transform.DownsamplingNone
	}

	aggr := transform.AggregatorEmpty
	switch opts.GetAggregator() {
	case pb.TransformerOptions_DEFAULT_AGGREGATOR:
		aggr = transform.DefaultAggregator
	case pb.TransformerOptions_AVG:
		aggr = transform.AggregatorAverage
	case pb.TransformerOptions_MAX:
		aggr = transform.AggregatorMax
	case pb.TransformerOptions_MS:
		aggr = transform.AggregatorMeanSquare
	case pb.TransformerOptions_RAVG:
		aggr = transform.AggregatorRoundedAverage
	case pb.TransformerOptions_RMS:
		aggr = transform.AggregatorRootMeanSquare
	}

	prec := roundToNearestPrecision(opts.GetPrecision())

	return &transform.ReaderOptions{
		Chunks:       int(opts.GetChunks()),
		Downsampling: dm,
		Aggregator:   aggr,
		Precision:    prec,
	}
}

func roundToNearestPrecision(precision int32) transform.Precision {
	if precision >= int32(transform.MinimumPrecision) {
		return transform.MinimumPrecision
	}
	p := math.Max(float64(precision), float64(transform.MaximumPrecision))
	exp := math.Floor(math.Log2(p))
	return transform.Precision(math.Exp2(exp))
}

func toBoxPainter(opts *pb.TrackRequest_Box) *box.BoxOptions {
	b := opts.Box

	alignment := box.AlignmentEmpty
	switch b.GetAlignment() {
	case pb.BoxOptions_DEFAULT_ALIGNMENT:
		alignment = box.DefaultAlignment
	case pb.BoxOptions_TOP:
		alignment = box.AlignmentTop
	case pb.BoxOptions_CENTER:
		alignment = box.AlignmentCenter
	case pb.BoxOptions_BOTTOM:
		alignment = box.AlignmentBottom
	}

	return &box.BoxOptions{
		Color:     b.GetColor(),
		Alignment: alignment,
		BoxHeight: b.GetHeight(),
		BoxWidth:  b.GetWidth(),
		Rounded:   b.GetRounded(),
		Gap:       b.GetGap(),
	}
}

func toLinePainter(opts *pb.TrackRequest_Line) *line.LineOptions {
	l := opts.Line

	interpolation := line.InterpolationEmpty
	switch l.GetInterpolation() {
	case pb.LineOptions_DEFAULT_INTERPOLATION:
		interpolation = line.DefaultInterpolation
	case pb.LineOptions_FRITSCH_CARLSON:
		interpolation = line.InterpolationFritschCarlson
	case pb.LineOptions_STEFFEN:
		interpolation = line.InterpolationSteffen
	case pb.LineOptions_NONE:
		interpolation = line.InterpolationNone
	}

	stroke := &line.Stroke{
		Color: l.Stroke.GetColor(),
		Width: l.Stroke.GetWidth(),
	}

	return &line.LineOptions{
		Interpolation: interpolation,
		Stroke:        stroke,
		Fill:          l.GetFill(),
		Inverted:      l.GetInverted(),
		Closed:        l.GetClosed(),
		Spread:        l.GetWidth(),
		Amplitude:     l.GetHeight(),
	}
}
