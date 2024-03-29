package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return the runtime data read by the reader for s.
// Usage: (read s)
func Read(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}

	s, ok := v.(runtime.String)
	if ok == false {
		return nil, runtime.BadType(runtime.StringType, v.Type())
	}

	v, _, err := runtime.Read(string(s))
	return v, err
}

// Evaluate runtime data as executable data.
// Usage: (eval data)
func Eval(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}
	return runtime.Eval(v, env)
}

// Dynamically apply a function with some arguments.
// Usage: (apply f args)
func Apply(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var fn, argv runtime.Value
	if err := runtime.ReadArgs(args, &fn, &argv); err != nil {
		return nil, err
	}
	fn2, ok := fn.(runtime.Callable)
	if ok == false {
		return nil, runtime.BadType(runtime.FuncType, fn.Type())
	}
	argv2, ok := argv.(runtime.Sequence)
	if ok == false {
		return nil, runtime.BadType(runtime.SequenceType, argv.Type())
	}
	return runtime.Apply(fn2, env, argv2)
}

// Return the string representation of form.
// Usage: (inspect form)
func Inspect(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}
	return runtime.String(v.String()), nil
}
