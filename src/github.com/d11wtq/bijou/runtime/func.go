package runtime

// Function data type
type Func struct {
	// Parameter list
	Params *List
	// Function body expressions
	Body *List
	// Closed environment
	Env Env
}

func (fn *Func) Eq(other Value) bool {
	return fn == other
}

func (fn *Func) Type() Type {
	return FuncType
}

func (fn *Func) Eval(env Env) (Value, error) {
	return fn, nil
}

func (fn *Func) IsMacro() bool {
	return false
}

// Call this function with the given arguments
func (fn *Func) Call(args *List) (Value, error) {
	env := fn.Env.Extend()
	params := fn.Params
	processed := 0

	for params != EmptyList && args != EmptyList {
		env.Def(string(params.Data.(Symbol)), args.Data)
		processed += 1
		params, args = params.Tail(), args.Tail()
	}

	if params != EmptyList || args != EmptyList {
		return nil, HandleBadArity(processed, params, args)
	}

	return EvalDo(env, fn.Body)
}

// Return an ArgumentError, inspecting unprocessed arguments
func HandleBadArity(processed int, params *List, args *List) error {
	numParams, numArgs := processed, processed
	for params != EmptyList {
		numParams += 1
		params = params.Next
	}
	for args != EmptyList {
		numArgs += 1
		args = args.Next
	}
	return BadArity(numParams, numArgs)
}
