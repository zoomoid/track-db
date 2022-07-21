package runner

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"

	cmdutils "github.com/zoomoid/waveman2/cmd/utils"
	"github.com/zoomoid/waveman2/pkg/painter"
	"github.com/zoomoid/waveman2/pkg/painter/box"
	"github.com/zoomoid/waveman2/pkg/painter/line"
	"github.com/zoomoid/waveman2/pkg/svg"
	"github.com/zoomoid/waveman2/pkg/transform"
)

type RunnerOptions struct {
	Transformer *transform.ReaderOptions
	BoxPainter  *box.BoxOptions
	LinePainter *line.LineOptions
	File        io.Reader
}

type Runner struct {
	options *RunnerOptions
	// transformerOutput is a buffer capturing the float64 slice returned from the transformer
	// When no painter options are passed, the transformer can be used as a standalone, where
	// it returns the reduced samples for further processing by other means
	transformerOutput *bytes.Buffer
	// boxOutput is the output of the box painter. Note that the buffer is only allocated when
	// the box painter is configured, thus allowing for nil checks in the Output() function
	// to determine which buffer to return
	boxOutput *bytes.Buffer
	// lineOutput is the output of the line painter. Note that the buffer is only allocated when
	// the box painter is configured, thus allowing for nil checks in the Output() function
	// to determine which buffer to return
	lineOutput *bytes.Buffer
	samples    *[]float64
	// skipPainters is a flag that indicates when we need to skip painting due to an error during transformation
	skipPainters bool
	// errors is the internal list of all errors produced during running
	errors []error
}

func NewRunner(opts *RunnerOptions) *Runner {
	buf := bytes.NewBuffer([]byte{})
	return &Runner{
		options:           opts,
		transformerOutput: buf,
	}
}

func (r *Runner) Transform() *Runner {
	t, err := transform.New(r.options.Transformer, r.options.File)
	if err != nil {
		r.errors = append(r.errors, err)
		r.skipPainters = true
		return r
	}

	samples := t.Blocks()
	r.samples = &samples
	r.samplesToBytes()

	return r
}

func (r *Runner) samplesToBytes() {
	for _, s := range *r.samples {
		bits := math.Float64bits(s)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, bits)
		r.transformerOutput.Write(bytes)
	}
}

func (r *Runner) Paint() *Runner {
	if r.skipPainters {
		return r
	}

	if r.options.BoxPainter != nil {
		painter := box.New(&painter.PainterOptions{
			Data:   *r.samples,
			Height: r.options.BoxPainter.BoxHeight,
			Width:  r.options.BoxPainter.BoxWidth,
		}, r.options.BoxPainter)

		r.boxOutput = r.paint(painter)
	}

	if r.options.LinePainter != nil {
		painter := line.New(&painter.PainterOptions{
			Data:   *r.samples,
			Height: r.options.LinePainter.Amplitude,
			Width:  r.options.LinePainter.Spread,
		}, r.options.LinePainter)

		r.lineOutput = r.paint(painter)
	}

	return r
}

func (r *Runner) paint(painter painter.Painter) *bytes.Buffer {
	elements := painter.Draw()

	out, err := svg.Template(elements, painter.Width(), painter.Height(), true)
	if err != nil {
		r.errors = append(r.errors, err)
	}
	return out
}

func (r *Runner) SVG() *bytes.Buffer {
	if r.boxOutput != nil {
		return r.boxOutput
	}
	if r.lineOutput != nil {
		return r.lineOutput
	}

	return nil
}

func (r *Runner) Samples() *bytes.Buffer {
	return r.transformerOutput
}

func (r *Runner) Error() error {
	errlist := cmdutils.NewErrorList(r.errors)
	if errlist != nil {
		return errors.New(errlist.Error())
	}
	return nil
}
