package runtime

import (
	"fmt"
)

// Symbol data type
type Symbol string

func (v Symbol) Eq(other Value) bool {
	return v == other
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
