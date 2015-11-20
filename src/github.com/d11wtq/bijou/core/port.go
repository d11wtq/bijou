package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Write a value to an open port.
// Usage: (write port value)
func Write(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
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
func Accept(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
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
