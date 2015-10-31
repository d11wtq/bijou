package runtime

import (
	"fmt"
)

// Function data type
type Func struct {
	// Parameter list
	Params *List
	// Function body expressions
	Body *List
	// Closed environment
	Env Env
}

func (fn *Func) Type() Type {
	return FuncType
}

func (fn *Func) Eval(env Env) (Value, error) {
	return fn, nil
}

// Call this function with the given arguments
func (fn *Func) Call(args *List) (Value, error) {
	env := fn.Env.Extend()
	params := fn.Params
	processed := uint(0)

	for params != EmptyList && args != EmptyList {
		env.Def(string(params.Data.(Symbol)), args.Data)
		processed += 1
		params, args = params.Tail(), args.Tail()
	}

	if params != EmptyList || args != EmptyList {
		return nil, BadArity(processed, params, args)
	}

	return EvalDo(env, fn.Body)
}

// Return an ArgumentError, inspecting unprocessed arguments
func BadArity(processed uint, params *List, args *List) error {
	numParams, numArgs := processed, processed
	for params != EmptyList {
		numParams += 1
		params = params.Tail()
	}
	for args != EmptyList {
		numArgs += 1
		args = args.Tail()
	}
	return &ArgumentError{
		fmt.Sprintf(
			"wrong number of arguments (wanted %d, got %d)",
			numParams,
			numArgs,
		),
	}
}
