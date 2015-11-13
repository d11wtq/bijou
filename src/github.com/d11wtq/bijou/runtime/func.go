package runtime

import (
	"fmt"
)

// Function data type
type Func struct {
	// Parameter list
	Params Sequence
	// Function body expressions
	Body Sequence
	// Closed environment
	Env Env
}

func (fn *Func) Eq(other Value) bool {
	return fn == other
}

func (fn *Func) Gt(other Value) bool {
	y, ok := other.(*Func)
	if ok == false {
		return fn.Type() > other.Type()
	}

	return fmt.Sprintf("%p", fn) > fmt.Sprintf("%p", y)
}

func (fn *Func) Lt(other Value) bool {
	y, ok := other.(*Func)
	if ok == false {
		return fn.Type() < other.Type()
	}

	return fmt.Sprintf("%p", fn) < fmt.Sprintf("%p", y)
}

func (fn *Func) Type() Type {
	return FuncType
}

func (fn *Func) Eval(env Env) (Value, error) {
	return fn, nil
}

func (fn *Func) String() string {
	return "#<function>"
}

// Call this function with the given arguments
func (fn *Func) Call(envc Env, args Sequence) (Value, error) {
	env := fn.Env.Extend()

	seen := 0
	a, p := args, fn.Params
	for !p.Empty() {
		key := p.Head().(Symbol)

		if key == Symbol("&") {
			// variadic; consume everything
			if !p.Tail().Empty() {
				key = p.Tail().Head().(Symbol)
				env.Def(string(key), a)
			}

			return EvalDo(env, fn.Body)
		}

		if a.Empty() {
			return nil, BadArity(seen+Length(p), seen)
		}

		env.Def(string(key), a.Head())
		a, p = a.Tail(), p.Tail()
		seen += 1
	}

	if !a.Empty() {
		return nil, BadArity(seen, seen+Length(a))
	}

	return EvalDo(env, fn.Body)
}
