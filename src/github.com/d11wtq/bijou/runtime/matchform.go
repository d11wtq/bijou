package runtime

// Process the elements of the 'match' special form
func EvalMatch(env Env, s Sequence) (Value, error) {
	if s.Empty() {
		return nil, &ArgumentError{"missing cases in match"}
	}

	head, err := Eval(s.Head(), env)
	if err != nil {
		return nil, err
	}

	switch head.(type) {
	case *Func:
		return EvalMatchFunc(env, Cons(head, s.Tail()))
	default:
		return EvalMatchCase(env, head, s.Tail())
	}
}
