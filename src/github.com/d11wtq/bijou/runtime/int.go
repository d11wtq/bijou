package runtime

import (
	"strconv"
)

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

func (v Int) String() string {
	return strconv.Itoa(int(v))
}

func (v Int) Add(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, &RuntimeError{"Bad data type: int required"}
	}
	return v + i, nil
}

func (v Int) Sub(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, &RuntimeError{"Bad data type: int required"}
	}
	return v - i, nil
}

func (v Int) Div(x Value) (Value, error) {
	i, ok := x.(Int)
	if ok == false {
		return nil, &RuntimeError{"Bad data type: int required"}
	}
	if i == Int(0) {
		return nil, &ArithmeticError{"Divide by zero"}
	}
	return v / i, nil
}
