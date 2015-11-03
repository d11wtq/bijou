package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Signature of a runtime-compatible go func
type CallableFunc (func(*runtime.List) (runtime.Value, error))

// Wrapper for a runtime-compatible go func
type FuncWrapper struct {
	Func CallableFunc
}

// Create a runtime-compatible go func
func GoFunc(fn CallableFunc) runtime.Callable {
	return &FuncWrapper{fn}
}

// Equality comparison (Value interface method)
func (w *FuncWrapper) Eq(other runtime.Value) bool {
	return w == other
}

// Return the runtime Type (Value interface method)
func (w *FuncWrapper) Type() runtime.Type {
	return runtime.FuncType
}

// Evaluate this fn to a runtime value (Value interface method)
func (w *FuncWrapper) Eval(env runtime.Env) (runtime.Value, error) {
	return w, nil
}

// Invoke this function and return a value (Callable interface method)
func (w *FuncWrapper) Call(args *runtime.List) (runtime.Value, error) {
	return w.Func(args)
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
