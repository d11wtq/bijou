package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return the head of the given sequence
func Head(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var lst runtime.Value
	if err := ReadArgs(args, &lst); err != nil {
		return nil, err
	}
	seq, ok := lst.(runtime.Sequence)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: sequence required"}
	}

	return seq.Head(), nil
}

// Return the tail of the given sequence
func Tail(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var lst runtime.Value
	if err := ReadArgs(args, &lst); err != nil {
		return nil, err
	}
	seq, ok := lst.(runtime.Sequence)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: sequence required"}
	}

	return seq.Tail(), nil
}
