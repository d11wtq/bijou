package core

import (
	"github.com/d11wtq/bijou/runtime"
)

func Read(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}

	s, ok := v.(runtime.String)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: string required"}
	}

	v, _, err := runtime.Read(string(s))
	return v, err
}

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
