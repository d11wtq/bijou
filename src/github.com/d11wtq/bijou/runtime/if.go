package runtime

// Process the elements of the 'if' special form
func EvalIf(env Env, s Sequence) (Value, error) {
	cond, body := s, s.Tail()

	if cond.Empty() {
		return nil, &RuntimeError{"Missing condition in if"}
	}
	if body.Empty() {
		return nil, &RuntimeError{"Missing body in if"}
	}

	success, err := cond.Head().Eval(env)

	if err != nil {
		return nil, err
	}

	if success != Nil && success != False {
		return body.Head().Eval(env)
	} else {
		return EvalDo(env, body.Tail())
	}
}
