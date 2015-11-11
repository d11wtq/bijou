package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return a list of all arguments in order.
// Usage: (list & args)
func List(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	if _, ok := args.(*runtime.List); ok == true {
		return args, nil
	}

	acc := runtime.EmptyList
	for !args.Empty() {
		acc = acc.Append(args.Head())
		args = args.Tail()
	}
	return acc, nil
}

// Return cons cell, whose head is hd and whose tail is tl.
// Usage: (cons hd tl)
func Cons(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var hd, tl runtime.Value
	if err := runtime.ReadArgs(args, &hd, &tl); err != nil {
		return nil, err
	}
	seq, ok := tl.(runtime.Sequence)
	if ok == false {
		return nil, runtime.BadType(runtime.SequenceType, tl.Type())
	}

	return runtime.Cons(hd, seq), nil
}
