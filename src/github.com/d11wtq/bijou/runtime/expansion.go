package runtime

// Process the elements of a macro call form
func EvalExpansion(env Env, mc Expandable, args Sequence) (Value, error) {
	res, err := mc.Expand(env, args)
	t, ok := res.(TailCall)
	if ok == true {
		res, err = t.Return()
	}

	if err != nil {
		return nil, err
	}

	return res.Eval(env)
}
