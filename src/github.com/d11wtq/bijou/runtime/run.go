package runtime

// Execute the given source string in env.
func Run(src string, env Env) (Value, error) {
	forms, err := ReadSrc(src)
	if err != nil {
		return nil, err
	}
	return EvalForms(forms, env)
}
