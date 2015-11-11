package runtime

import (
	"strconv"
)

// Signed integer data type
type Int int

func (v Int) Eq(other Value) bool {
	return v == other
}

func (v Int) Gt(other Value) bool {
	y, ok := other.(Int)
	if ok == false {
		return v.Type() > other.Type()
	}

	return v > y
}

func (v Int) Lt(other Value) bool {
	y, ok := other.(Int)
	if ok == false {
		return v.Type() < other.Type()
	}

	return v < y
}

func (v Int) Eval(env Env) (Value, error) {
	return v, nil
}

func (v Int) Type() Type {
	return IntType
}

func (v Int) String() string {
	return strconv.Itoa(int(v))
}

func (v Int) Add(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, BadType(IntType, x.Type())
	}
	return v + i, nil
}

func (v Int) Sub(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, BadType(IntType, x.Type())
	}
	return v - i, nil
}

func (v Int) Mul(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, BadType(IntType, x.Type())
	}
	return v * i, nil
}

func (v Int) Div(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, BadType(IntType, x.Type())
	}
	if i == Int(0) {
		return nil, DivisionByZero()
	}
	return v / i, nil
}
