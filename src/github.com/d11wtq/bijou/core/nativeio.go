package core

import (
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
	return runtime.PtrGt(w, y)
}

// (Value interface)
func (w *IoPortWrapper) Lt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return w.Type() < other.Type()
	}
	return runtime.PtrLt(w, y)
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
		return nil, w.ReadError()
	}

	buf := make([]byte, 1)
	_, err := w.Reader.Read(buf)

	switch err {
	case nil:
		return runtime.Int(buf[0]), nil
	case io.EOF:
		return runtime.Nil, nil
	default:
		return nil, err
	}
}

// (Port interface)
func (w *IoPortWrapper) Read(n int) (runtime.Sequence, error) {
	if w.Reader == nil {
		return nil, w.ReadError()
	}

	var (
		got int
		pos int
		err error
	)

	buf := make([]byte, n)

Loop:
	for {
		got, err = w.Reader.Read(buf[pos:])
		pos += got

		if got < n {
			switch err {
			case nil:
				// no error, get some more
				n -= got
				continue Loop
			case io.EOF:
				// ok, done
			default:
				return nil, err
			}
		}
		break
	}

	return runtime.String(buf[0:pos]), nil
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

// Error returned in the case we try to read from a non-read port.
func (w *IoPortWrapper) ReadError() error {
	return &runtime.RuntimeError{"Port is not open for reading"}
}
