package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New() return nil.")
	} else {
		tracer.Trace("Hello from trace package.")
		if buf.String() != "Hello from trace package.\n" {
			t.Errorf("Output wrong strings: '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("データ")
}
