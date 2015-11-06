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
func (fn *Func) Call(envc Env, args *List) (Value, error) {
	env := fn.Env.Extend()

	seen := 0
	for params := fn.Params; params != EmptyList; params = params.Next {
		key := params.Data.(Symbol)

		if key == Symbol("&") {
			// variadic; consume everything
			if params.Next != EmptyList {
				key = params.Next.Data.(Symbol)
				env.Def(string(key), args)
			}

			return EvalDo(env, fn.Body)
		}

		if args == EmptyList {
			return nil, BadArity(seen+params.Length(), seen)
		}

		env.Def(string(key), args.Data)
		args = args.Next
		seen += 1
	}

	if args != EmptyList {
		return nil, BadArity(seen, seen+args.Length())
	}

	return EvalDo(env, fn.Body)
}
