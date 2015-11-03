package runtime

// Boolean data type
type Boolean bool

// Representation of true
var True = (Boolean)(true)

// Representation of false
var False = (Boolean)(false)

func (v Boolean) Eq(other Value) bool {
	return v == other
}

func (v Boolean) Type() Type {
	return BooleanType
}

func (v Boolean) Eval(env Env) (Value, error) {
	return v, nil
}
