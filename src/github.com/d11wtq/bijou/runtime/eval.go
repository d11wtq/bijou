package runtime

// Evaluate a single form and return the result.
func Eval(form Value, env Env) (Value, error) {
	res, err := form.Eval(env)
	tmp, tco := res.(TailCall)
	if tco {
		return tmp.Resolve()
	}

	return res, err
}

// Evaluate a list of forms, returning the last evaluated result.
func EvalForms(forms Sequence, env Env) (acc Value, err error) {
	acc = Nil
	for !forms.Empty() {
		acc, err = Eval(forms.Head(), env)
		forms = forms.Tail()
	}
	return acc, err
}