package core

import (
	"fmt"
	"github.com/d11wtq/bijou/runtime"
	"io"
)

// A wrapper around an I/O stream
type IoPortWrapper struct {
	// The Reader end of the port, if available
	Reader io.ReadCloser
	// The writer end of the port, if available
	Writer io.WriteCloser
}

// Create a wrapper around Go's io built-ins for a runtime-compatible port.
func GoIoPort(r io.ReadCloser, w io.WriteCloser) runtime.Port {
	return &IoPortWrapper{r, w}
}

// (Value interface)
func (w *IoPortWrapper) Type() runtime.Type {
	return runtime.PortType
}

// (Value interface)
func (w *IoPortWrapper) Eval(env runtime.Env) (runtime.Value, error) {
	return w, nil
}

// (Value interface)
func (w *IoPortWrapper) String() string {
	return "#<port:io>"
}

// (Value interface)
func (w *IoPortWrapper) Eq(other runtime.Value) bool {
	return w == other
}

// (Value interface)
func (w *IoPortWrapper) Gt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return w.Type() > other.Type()
	}

	return fmt.Sprintf("%p", w) > fmt.Sprintf("%p", y)
}

// (Value interface)
func (w *IoPortWrapper) Lt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return w.Type() < other.Type()
	}

	return fmt.Sprintf("%p", w) < fmt.Sprintf("%p", y)
}

// (Port interface)
func (w *IoPortWrapper) Write(v runtime.Value) error {
	if w.Writer == nil {
		return &runtime.RuntimeError{"Port is not open for writing"}
	}

	var s string

	switch x := v.(type) {
	case runtime.String:
		s = string(x)
	default:
		s = v.String()
	}

	_, err := w.Writer.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}

// (Port interface)
func (w *IoPortWrapper) Accept() (runtime.Value, error) {
	if w.Reader == nil {
		return nil, &runtime.RuntimeError{"Port is not open for reading"}
	}

	buf := make([]byte, 1)
	if _, err := w.Reader.Read(buf); err != nil {
		return nil, err
	}

	return runtime.Int(buf[0]), nil
}

// (Port interface)
func (w *IoPortWrapper) Close() error {
	if w.Writer != nil {
		err := w.Writer.Close()
		if err != nil {
			return err
		}
	}

	if w.Reader != nil {
		err := w.Reader.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
