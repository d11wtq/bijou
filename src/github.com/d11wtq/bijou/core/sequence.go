package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return the head of the given sequence.
// Usage: (head seq)
func Head(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}

	return seq.Head(), nil
}

// Return the tail of the given sequence.
// Usage: (tail seq)
func Tail(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}

	return seq.Tail(), nil
}

// Create a new sequence by putting a new value at the end of the sequence.
// Usage: (put seq v)
func Put(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var seq, val runtime.Value
	if err := runtime.ReadArgs(args, &seq, &val); err != nil {
		return nil, err
	}
	seq2, ok := seq.(runtime.Sequence)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: sequence required"}
	}

	return seq2.Put(val)
}

// Return true if the sequence is empty.
// Usage: (empty? seq)
func Empty(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var seq runtime.Sequence
	if err := ReadSequence(args, &seq); err != nil {
		return nil, err
	}
	return runtime.Boolean(seq.Empty()), nil
}

// Read the sequence argument from args.
func ReadSequence(args runtime.Sequence, ptr *runtime.Sequence) error {
	var v runtime.Value
	err := runtime.ReadArgs(args, &v)
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
