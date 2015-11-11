package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return the sum of all arguments.
// Usage: (+ & args)
func Add(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var acc runtime.Value = runtime.Int(0)
	var err error

	for !args.Empty() {
		v, ok := acc.(runtime.Addition)
		if ok == false {
			return nil, runtime.BadOperation("+", acc.Type())
		}
		acc, err = v.Add(args.Head())
		if err != nil {
			return nil, err
		}
		args = args.Tail()
	}

	return acc, nil
}

// Return the subtraction of all arguments
// Usage: (- & args)
func Sub(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var acc runtime.Value = runtime.Int(0)
	var err error

	if !args.Tail().Empty() {
		acc, args = args.Head(), args.Tail()
	}

	for !args.Empty() {
		v, ok := acc.(runtime.Subtraction)
		if ok == false {
			return nil, runtime.BadOperation("-", acc.Type())
		}
		acc, err = v.Sub(args.Head())
		if err != nil {
			return nil, err
		}
		args = args.Tail()
	}

	return acc, nil
}

// Return the product of all arguments
// Usage: (* & args)
func Mul(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var acc runtime.Value = runtime.Int(1)
	var err error

	for !args.Empty() {
		v, ok := acc.(runtime.Multiplication)
		if ok == false {
			return nil, runtime.BadOperation("*", acc.Type())
		}
		acc, err = v.Mul(args.Head())
		if err != nil {
			return nil, err
		}
		args = args.Tail()
	}

	return acc, nil
}

// Return the division of all arguments
// Usage: (/ & args)
func Div(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var acc runtime.Value = runtime.Int(1)
	var err error

	if !args.Tail().Empty() {
		acc, args = args.Head(), args.Tail()
	}

	for !args.Empty() {
		v, ok := acc.(runtime.Division)
		if ok == false {
			return nil, runtime.BadOperation("/", acc.Type())
		}
		acc, err = v.Div(args.Head())
		if err != nil {
			return nil, err
		}
		args = args.Tail()
	}

	return acc, nil
}
