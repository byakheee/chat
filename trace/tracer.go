package trace

import (
	"fmt"
	"io"
)

// Tracer is interface that record the object for event in code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

// New return Tracer object.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
