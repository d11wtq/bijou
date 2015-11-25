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
	if other, ok := IsList(other); ok == true {
		a, b := Sequence(cons), Sequence(other)

		for {
			if a.Empty() && b.Empty() {
				return true
			}

			if a.Empty() || b.Empty() {
				return false
			}

			if !Eq(a.Head(), b.Head()) {
				return false
			}

			a, b = a.Tail(), b.Tail()
		}
	} else {
		return false
	}
}

func (cons *ConsCell) Gt(other Value) bool {
	y, ok := IsList(other)
	if ok == false {
		return cons.Type() > other.Type()
	}

	a, b := Sequence(cons), Sequence(y)

	for {
		if a.Empty() && b.Empty() {
			return false
		}

		if b.Empty() {
			return true
		}

		if a.Empty() {
			return false
		}

		if Gt(a.Head(), b.Head()) {
			return true
		}

		if Lt(a.Head(), b.Head()) {
			return false
		}

		a, b = a.Tail(), b.Tail()
	}
}

func (cons *ConsCell) Lt(other Value) bool {
	y, ok := IsList(other)
	if ok == false {
		return cons.Type() < other.Type()
	}

	a, b := Sequence(cons), Sequence(y)

	for {
		if a.Empty() && b.Empty() {
			return false
		}

		if a.Empty() {
			return true
		}

		if b.Empty() {
			return false
		}

		if Lt(a.Head(), b.Head()) {
			return true
		}

		if Gt(a.Head(), b.Head()) {
			return false
		}

		a, b = a.Tail(), b.Tail()
	}
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
	if v, ok := IsList(value); ok == true {
		a, b := Sequence(cons), Sequence(v)

		for {
			if a.Empty() && b.Empty() {
				return nil
			}

			if a.Empty() || b.Empty() {
				return BadPattern(cons, value)
			}

			err := Bind(a.Head(), b.Head(), env)
			if err != nil {
				return err
			}

			a, b = a.Tail(), b.Tail()
		}
	} else {
		return BadPattern(cons, value)
	}
}
