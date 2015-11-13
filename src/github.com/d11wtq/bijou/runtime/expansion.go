package runtime

// Process the elements of a macro call form
func EvalExpansion(env Env, mc Expandable, args Sequence) (Value, error) {
	v, err := mc.Expand(env, args)
	if err != nil {
		return nil, err
	}
	return v.Eval(env)
}
