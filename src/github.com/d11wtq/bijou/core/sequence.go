package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return the head of the given sequence
func Head(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}

	return seq.Head(), nil
}

// Return the tail of the given sequence
func Tail(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}

	return seq.Tail(), nil
}

// Return true if the sequence is empty
func Empty(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}
	return runtime.Boolean(seq.Empty()), nil
}

// Read the sequence argument from args
func ReadSequence(args *runtime.List, ptr *runtime.Sequence) error {
	var v runtime.Value
	err := ReadArgs(args, &v)
	if err != nil {
		return err
	}
	seq, ok := v.(runtime.Sequence)
	if ok == false {
		return &runtime.RuntimeError{"Bad data type: sequence required"}
	}
	*ptr = seq
	return nil
}