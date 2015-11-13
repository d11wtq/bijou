package runtime

// Process the elements of a procedure call
func EvalProcCall(env Env, proc Value, args Sequence) (Value, error) {
	proc, err := proc.Eval(env)
	if err != nil {
		return nil, err
	}

	switch proc.(type) {
	case Callable:
		return EvalCall(env, proc.(Callable), args)
	case Expandable:
		return EvalExpansion(env, proc.(Expandable), args)
	default:
		return nil, BadOperation("call", proc.Type())
	}
}
