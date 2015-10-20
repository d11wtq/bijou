package runtime

// Signed integer data type
type Int int

func (v Int) Eval(env Env) (Value, error) {
	return v, nil
}

func (v Int) Type() Type {
	return IntType
}
