package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Typed channel for Value objects
type ValueChannel chan runtime.Value

// A Go channel for communicating with Values
type ChanPort struct {
	// The channel to communicate on
	Channel ValueChannel
	// True once this channel is closed
	Closed bool
}

// Create a wrapper around Go's chans as a runtime-compatible port.
func GoChanPort() runtime.Port {
	return &ChanPort{make(ValueChannel), false}
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

// (Value interface)
func (p *ChanPort) Eq(other runtime.Value) bool {
	return p == other
}

// (Value interface)
func (p *ChanPort) Gt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return p.Type() > other.Type()
	}
	return runtime.PtrGt(p, y)
}

// (Value interface)
func (p *ChanPort) Lt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return p.Type() < other.Type()
	}
	return runtime.PtrLt(p, y)
}

// (Port interface)
func (p *ChanPort) Write(v runtime.Value) error {
	if p.Closed {
		// writes to closed ports are silently droppped
		return nil
	}

	p.Channel <- v
	return nil
}

// (Port interface)
func (p *ChanPort) Accept() (runtime.Value, error) {
	v, ok := <-p.Channel
	if ok == false {
		return nil, p.ReadError()
	}

	return v, nil
}

// (Port interface)
func (p *ChanPort) Close() (err error) {
	if p.Closed {
		return
	}

	p.Closed = true
	p.Flush()
	close(p.Channel)
	return
}

// Consume all messages on the channel.
func (p *ChanPort) Flush() {
	for {
		_, ok := <-p.Channel
		if ok == false {
			break
		}
	}
}

// Error returned in the case we try to read from a closed port.
func (p *ChanPort) ReadError() error {
	return &runtime.RuntimeError{"Port is not open for reading"}
}
