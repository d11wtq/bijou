package runtime

// Signed integer data type
type Int int

func (v Int) Eq(other Value) bool {
	return v == other
}

func (v Int) Eval(env Env) (Value, error) {
	return v, nil
}

func (v Int) Type() Type {
	return IntType
}
