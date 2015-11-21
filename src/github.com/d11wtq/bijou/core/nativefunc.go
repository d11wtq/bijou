package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Signature of a runtime-compatible go func
type CallableFunc (func(runtime.Env, runtime.Sequence) (runtime.Value, error))

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

// Greater than comparison (Value interface method)
func (w *FuncWrapper) Gt(other runtime.Value) bool {
	if w.Type() == other.Type() {
		return runtime.PtrGt(w, other)
	}
	return w.Type() > other.Type()
}

// Less than comparison (Value interface method)
func (w *FuncWrapper) Lt(other runtime.Value) bool {
	if w.Type() == other.Type() {
		return runtime.PtrLt(w, other)
	}
	return w.Type() < other.Type()
}

// Return the runtime Type (Value interface method)
func (w *FuncWrapper) Type() runtime.Type {
	return runtime.FuncType
}

// Evaluate this fn to a runtime value (Value interface method)
func (w *FuncWrapper) Eval(env runtime.Env) (runtime.Value, error) {
	return w, nil
}

func (w *FuncWrapper) String() string {
	return "#<function>"
}

// Invoke this function and return a value (Callable interface method)
func (w *FuncWrapper) Call(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	return w.Func(env, args)
}
