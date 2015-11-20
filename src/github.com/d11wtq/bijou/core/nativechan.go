package core

import (
	"fmt"
	"github.com/d11wtq/bijou/runtime"
)

// A Go channel for communicating with Values
type ChanPort struct {
	// The channel to communicate on
	Channel chan runtime.Value
	// True once this channel is closed
	Closed bool
}

// Create a wrapper around Go's chans as a runtime-compatible port.
func GoChanPort() runtime.Port {
	return &ChanPort{make(chan runtime.Value), false}
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
	return "#<port:channel>"
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

	return fmt.Sprintf("%p", p) > fmt.Sprintf("%p", y)
}

// (Value interface)
func (p *ChanPort) Lt(other runtime.Value) bool {
	y, ok := other.(runtime.Port)
	if ok == false {
		return p.Type() < other.Type()
	}

	return fmt.Sprintf("%p", p) < fmt.Sprintf("%p", y)
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
	if p.Closed {
		return nil, &runtime.RuntimeError{"Port is not open for reading"}
	}

	return <-p.Channel, nil
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
