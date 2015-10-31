package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Signature of a runtime-compatible go func
type GoFunc (func(*runtime.List) (runtime.Value, error))

// Return the runtime Type (Value interface method)
func (fn GoFunc) Type() runtime.Type {
	return runtime.FuncType
}

// Evaluate this fn to a runtime value (Value interface method)
func (fn GoFunc) Eval(env runtime.Env) (runtime.Value, error) {
	return fn, nil
}

// Invoke this function and return a value (Callable interface method)
func (fn GoFunc) Call(args *runtime.List) (runtime.Value, error) {
	return fn(args)
}

// Read the expected number of arguments into the slice of pointers ptrs
func ReadArgs(args *runtime.List, ptrs ...*runtime.Value) error {
	for n, ptr := range ptrs {
		if args == runtime.EmptyList {
			return runtime.BadArity(len(ptrs), n)
		}
		*ptr = args.Data
		args = args.Next
	}

	if args != runtime.EmptyList {
		n := len(ptrs)
		for args != runtime.EmptyList {
			n += 1
			args = args.Next
		}
		return runtime.BadArity(len(ptrs), n)
	}

	return nil
}
