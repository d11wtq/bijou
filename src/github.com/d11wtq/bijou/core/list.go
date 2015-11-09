package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Basic list function, returns variadic args as a *List
func List(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	acc := &runtime.List{}
	for !args.Empty() {
		acc.Append(args.Head())
		args = args.Tail()
	}
	return acc, nil
}

// Return a new cons cell appending a new head to a given sequence
func Cons(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var hd, tl runtime.Value
	if err := runtime.ReadArgs(args, &hd, &tl); err != nil {
		return nil, err
	}
	seq, ok := tl.(runtime.Sequence)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: sequence required"}
	}

	return runtime.Cons(hd, seq), nil
}
