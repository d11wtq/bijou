package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Write a value to an open port.
// Usage: (write port value)
func PortWrite(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var port, value runtime.Value
	if err := runtime.ReadArgs(args, &port, &value); err != nil {
		return nil, err
	}
	p, ok := port.(runtime.Port)
	if ok == false {
		return nil, runtime.BadType(runtime.PortType, port.Type())
	}

	err := p.Write(value)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Accept a value from an open port.
// Usage: (accept port)
func PortAccept(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var port runtime.Value
	if err := runtime.ReadArgs(args, &port); err != nil {
		return nil, err
	}
	p, ok := port.(runtime.Port)
	if ok == false {
		return nil, runtime.BadType(runtime.PortType, port.Type())
	}

	return p.Accept()
}

// Read a sequence from an open port.
// Usage: (read port n)
func PortRead(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var port, n runtime.Value
	if err := runtime.ReadArgs(args, &port, &n); err != nil {
		return nil, err
	}
	p, ok := port.(runtime.Port)
	if ok == false {
		return nil, runtime.BadType(runtime.PortType, port.Type())
	}
	n2, ok := n.(runtime.Int)
	if ok == false {
		return nil, runtime.BadType(runtime.IntType, n.Type())
	}

	return p.Read(int(n2))
}
