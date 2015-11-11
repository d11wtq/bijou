package runtime

import (
	"fmt"
)

// Get the name of a given Type
func TypeName(t Type) string {
	switch t {
	case NilType:
		return "nil"
	case BooleanType:
		return "boolean"
	case IntType:
		return "integer"
	case SymbolType:
		return "symbol"
	case SequenceType:
		return "sequence"
	case StringType:
		return "string"
	case ConsType:
		return "cons"
	case ListType:
		return "list"
	case FuncType:
		return "function"
	case MacroType:
		return "macro"
	default:
		panic(fmt.Sprintf("Unknown type: %d", t))
	}
}

// Predicate to check if a Value is a List or a Cons
//
// Lists and Cons should be considered the same conceptually, thus parts of the
// runtime check for either of the two.
func IsList(v Value) (Sequence, bool) {
	var (
		s  Sequence
		ok bool
	)
	s, ok = v.(*ConsCell)
	if ok == false {
		s, ok = v.(*List)
	}
	return s, ok
}

// Return the length of a given Sequence.
func Length(s Sequence) int {
	acc := 0
	for !s.Empty() {
		acc += 1
		s = s.Tail()
	}
	return acc
}
