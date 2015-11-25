package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Typed channel for Value objects
type ValueChannel chan runtime.Value

// A Go channel for communicating with Values
type ChanPort struct {
	// The channel to read on
	Reader <-chan runtime.Value
	// The channel to write on
	Writer chan<- runtime.Value
	// True once this channel is closed
	Closed bool
}

// Return a pair of CanPorts, for bi-directional communication.
func GoChanPortPair() (runtime.Port, runtime.Port) {
	a, b := make(ValueChannel), make(ValueChannel)
	return GoChanPort(a, b), GoChanPort(b, a)
}

// Create a wrapper around Go's chans as a runtime-compatible port.
func GoChanPort(reader, writer ValueChannel) runtime.Port {
	return &ChanPort{reader, writer, false}
}

// (Value interface)
func (p *ChanPort) Type() runtime.Type {
	return runtime.PortType
}

// (Value interface)
func (p *ChanPort) Eval(env runtime.Env) (runtime.Value, error) {
	return p, nil
}

// (Value interface)
func (p *ChanPort) String() string {
	return "#<port:proc>"
}

// (Port interface)
func (p *ChanPort) Write(v runtime.Value) error {
	if p.Closed {
		return nil // dropped
	}

	p.Writer <- v
	return nil
}

// (Port interface)
func (p *ChanPort) Accept() (runtime.Value, error) {
	v, ok := <-p.Reader
	if ok == false {
		return nil, p.ReadError()
	}

	return v, nil
}

// (Port interface)
func (p *ChanPort) Read(n int) (runtime.Sequence, error) {
	var (
		val runtime.Value
		err error
	)

	acc := runtime.EmptyList
	for i := 0; i < n; i += 1 {
		val, err = p.Accept()
		if err != nil {
			return nil, err
		}
		acc = acc.Append(val)
	}
	return acc, nil
}

// (Port interface)
func (p *ChanPort) Close() (err error) {
	if p.Closed {
		return
	}

	p.Closed = true
	p.Flush()
	close(p.Writer)
	return
}

// Consume all messages on the channel.
func (p *ChanPort) Flush() {
	for range p.Reader {
		// drop
	}
}

// Error returned in the case we try to read from a closed port.
func (p *ChanPort) ReadError() error {
	return &runtime.RuntimeError{"Port is not open for reading"}
}
