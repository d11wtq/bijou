package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Spawn a new process using the function f and the given args.
// The function f receives a port as the first argument.
// The port is returned to the caller for communication.
//
// Usage: (spawn f & args)
func Spawn(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	if args.Empty() {
		return nil, runtime.BadArity(1, 0)
	}

	v, args := args.Head(), args.Tail()
	fun, ok := v.(runtime.Callable)
	if ok == false {
		return nil, runtime.BadType(runtime.FuncType, v.Type())
	}

	port := GoChanPort()
	go func() {
		runtime.Apply(fun, env, runtime.Cons(port, args))
		port.Close()
	}()
	return port, nil
}
