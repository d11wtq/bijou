package runtime

// Procedure data type
type Proc struct {
	// Parameter list
	Params Sequence
	// Procedure body expressions
	Body Sequence
	// Closed environment
	Env Env
}

// Call this procedure with the given arguments
func (proc *Proc) Call(envc Env, args Sequence) (Value, error) {
	env := proc.Env.Extend()

	seen := 0
	a, p := args, proc.Params
	for !p.Empty() {
		key := p.Head().(Symbol)

		if key == Symbol("&") {
			// variadic; consume everything
			if !p.Tail().Empty() {
				key = p.Tail().Head().(Symbol)
				env.Def(string(key), a)
			}

			return EvalDo(env, proc.Body)
		}

		if a.Empty() {
			return nil, BadArity(seen+Length(p), seen)
		}

		env.Def(string(key), a.Head())
		a, p = a.Tail(), p.Tail()
		seen += 1
	}

	if !a.Empty() {
		return nil, BadArity(seen, seen+Length(a))
	}

	return EvalDo(env, proc.Body)
}
