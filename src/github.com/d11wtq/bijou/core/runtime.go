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
	return v.Eval(env)
}

func Inspect(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}
	return runtime.String(v.String()), nil
}
