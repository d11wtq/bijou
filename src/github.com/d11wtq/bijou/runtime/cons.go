package runtime

import (
	"fmt"
	"strings"
)

// Create a new ConsCell with a given Head and Tail
func Cons(data Value, next Sequence) Sequence {
	return &ConsCell{data, next}
}

// ConsCell data type (Sequence)
type ConsCell struct {
	// Value held in this ConsCell
	Data Value
	// The Tail that this ConsCell extends
	Next Sequence
}

var EmptyCons = (*ConsCell)(nil)

func (cons *ConsCell) Eq(other Value) bool {
	return ListEq(cons, other)
}

func (cons *ConsCell) Gt(other Value) bool {
	return ListGt(cons, other)
}

func (cons *ConsCell) Lt(other Value) bool {
	return ListLt(cons, other)
}

func (cons *ConsCell) Type() Type {
	return ConsType
}

func (cons *ConsCell) Eval(env Env) (Value, error) {
	if cons.Empty() {
		return EmptyCons, nil
	}

	switch cons.Data {
	case Symbol("quote"):
		return EvalQuote(env, cons.Next)
	case Symbol("if"):
		return EvalIf(env, cons.Next)
	case Symbol("do"):
		return EvalDo(env, cons.Next)
	case Symbol("fn"):
		return EvalFn(env, cons.Next)
	case Symbol("match"):
		return EvalMatch(env, cons.Next)
	case Symbol("macro"):
		return EvalMacro(env, cons.Next)
	case Symbol("def"):
		return EvalDef(env, cons.Next)
	case Symbol("bind"):
		return EvalBind(env, cons.Next)
	default:
		return EvalProcCall(env, cons.Data, cons.Next)
	}
}

func (cons *ConsCell) String() string {
	strs := make([]string, 0, Length(cons))
	s := Sequence(cons)
	for !s.Empty() {
		strs = append(strs, s.Head().String())
		s = s.Tail()
	}
	return fmt.Sprintf("(%s)", strings.Join(strs, " "))
}

func (cons *ConsCell) Head() Value {
	if cons == EmptyCons {
		return Nil
	}

	return cons.Data
}

func (cons *ConsCell) Tail() Sequence {
	if cons == EmptyCons {
		return EmptyCons
	}

	return cons.Next
}

func (cons *ConsCell) Empty() bool {
	return cons == EmptyCons
}

func (cons *ConsCell) Put(v Value) (Sequence, error) {
	return &ConsCell{v, cons}, nil
}

func (cons *ConsCell) Bind(env Env, value Value) error {
	if cons.Head() == Symbol("quote") {
		return EqPattern(cons.Tail().Head(), value)
	}

	return ListBind(env, cons, value)
}
