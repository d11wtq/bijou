package runtime

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

			if !a.Head().Eq(b.Head()) {
				return false
			}

			a, b = a.Tail(), b.Tail()
		}
	} else {
		return false
	}
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
	default:
		return EvalCall(env, cons)
	}
}

func (cons *ConsCell) Type() Type {
	return ConsType
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