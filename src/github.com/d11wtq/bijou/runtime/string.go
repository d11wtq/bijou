package runtime

// String data type
type String string

func (s String) Eq(other Value) bool {
	return s == other
}

func (s String) Eval(env Env) (Value, error) {
	return s, nil
}

func (s String) Type() Type {
	return StringType
}
