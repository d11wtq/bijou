package runtime

// Nil data type
type NilObj struct{}

// Representation of nil
var Nil = (*NilObj)(nil)

func (*NilObj) Type() Type {
	return NilType
}

func (*NilObj) Eval(env Env) (Value, error) {
	return Nil, nil
}
