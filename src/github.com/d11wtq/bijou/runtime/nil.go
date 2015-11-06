package runtime

// Nil data type
type NilObj struct{}

// Representation of nil
var Nil = (*NilObj)(nil)

func (v *NilObj) Eq(other Value) bool {
	return v == other
}

func (v *NilObj) Type() Type {
	return NilType
}

func (v *NilObj) Eval(env Env) (Value, error) {
	return Nil, nil
}

func (v *NilObj) Head() Value {
	return v
}

func (v *NilObj) Tail() Sequence {
	return v
}

func (v *NilObj) Empty() bool {
	return true
}
