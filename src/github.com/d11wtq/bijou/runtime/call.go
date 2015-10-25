package runtime

// Process the elements of a function call form
func EvalCall(env Env, lst *List) (Value, error) {
	callee, err := lst.Data.Eval(env)
	if err != nil {
		return nil, err
	}

	fn, ok := callee.(*Func)
	if ok == false {
		return nil, &RuntimeError{"Attempted to call non-function"}
	}

	return fn.Apply(lst.Next)
}
