package main

import (
	"fmt"
)

// Program data type
type Type uint8

const (
	NilType Type = iota
	SymbolType
	IntType
	StringType
	ListType
	FuncType
)

type Env map[Symbol]Value

// Program runtime value
type Value interface {
	// Eval to the lowest level value
	Eval(env Env) Value
	// Eval to the lowest level value
	Type() Type
}

// Symbol data type
type Symbol string

func (v Symbol) Eval(env Env) Value {
	return env[v]
}

func (v Symbol) Type() Type {
	return SymbolType
}

// Signed integer data type
type Int int

func (v Int) Eval(env Env) Value {
	return v
}

func (v Int) Type() Type {
	return IntType
}

// String data type
type String string

func (v String) Eval(env Env) Value {
	return v
}

func (v String) Type() Type {
	return StringType
}

// List data type
type List struct {
	// First element of the list
	Head Value
	// The rest of the list, or nil
	Tail *List
}

func (v *List) Eval(env Env) Value {
	if v == nil {
		return nil
	}

	switch v.Head {
	case Symbol("do"):
		return v.Tail.Do(env)
	default:
		head := v.Head.Eval(env)
		if head.Type() == FuncType {
			return head.(*Func).Apply(v.Tail.Slice())
		} else {
			// FIXME: This should be an error
			return nil
		}
	}
}

func (v *List) Type() Type {
	return ListType
}

// Convert this list into a slice
func (v *List) Slice() []Value {
	s := make([]Value, 0)
	for y := v; y != nil; y = y.Tail {
		s = append(s, y.Head)
	}
	return s
}

// Append a new element to the head of the list
func (v *List) Cons(head Value) *List {
	return &List{head, v}
}

func (v *List) Do(env Env) Value {
	var result Value
	for y := v; y != nil; y = y.Tail {
		result = y.Head.Eval(env)
	}
	return result
}

// Function data type
type Func struct {
	// Parameter names
	Params []Symbol
	// Body expressions
	Body *List
	// Closed environment
	Env Env
}

func (fn *Func) Eval(env Env) Value {
	return fn
}

func (fn *Func) Type() Type {
	return FuncType
}

// Apply this function to the given args
func (fn *Func) Apply(args []Value) Value {
	return fn.Body.Do(fn.Env)
}

func Inspect(v Value, env Env) string {
	switch y := v.Eval(env); y.Type() {
	case NilType:
		return "nil"
	case SymbolType:
		return string(y.(Symbol))
	case IntType:
		return fmt.Sprintf("%d", y.(Int))
	case StringType:
		return fmt.Sprintf(`"%s"`, y.(String))
	case ListType:
		return fmt.Sprintf("(%s)", y)
	case FuncType:
		return "<function>"
	default:
		return fmt.Sprintf("%s", y)
	}
}

var EmptyList *List = nil

func main() {
	env := make(Env)
	env[Symbol("foo")] = Int(42)
	env[Symbol("bar")] = &Func{
		make([]Symbol, 0),
		EmptyList.Cons(Symbol("foo")).Cons(Int(7)),
		env,
	}
	fmt.Println(
		Inspect(
			EmptyList.Cons(
				EmptyList.Cons(Symbol("bar")),
			).Cons(Symbol("Do")),
			env,
		),
	)
}
