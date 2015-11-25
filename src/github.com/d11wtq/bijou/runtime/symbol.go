package runtime

import (
	"fmt"
)

// Symbol data type
type Symbol string

func (v Symbol) Eq(other Value) bool {
	return v == other
}

func (v Symbol) Gt(other Value) bool {
	y, ok := other.(Symbol)
	if ok == false {
		return v.Type() > other.Type()
	}

	return v > y
}

func (v Symbol) Lt(other Value) bool {
	y, ok := other.(Symbol)
	if ok == false {
		return v.Type() < other.Type()
	}

	return v < y
}

func (s Symbol) Eval(env Env) (Value, error) {
	if v, ok := env.Get(string(s)); ok {
		return v, nil
	} else {
		return v, &RuntimeError{fmt.Sprintf("Undefined variable %s", s)}
	}
}

func (s Symbol) Type() Type {
	return SymbolType
}

func (s Symbol) String() string {
	return string(s)
}

func (s Symbol) Bind(env Env, value Value) error {
	v, ok := env.Get(string(s))
	if ok == true {
		return Bind(v, value, env)
	} else {
		return env.Def(string(s), value)
	}
}
