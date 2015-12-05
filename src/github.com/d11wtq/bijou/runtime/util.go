package runtime

import (
	"fmt"
)

// Map of Type -> name associations
var TypeNameMap = map[Type]string{
	NilType:      "nil",
	BooleanType:  "boolean",
	IntType:      "integer",
	SymbolType:   "symbol",
	SequenceType: "sequence",
	AssocType:    "associative",
	StringType:   "string",
	ConsType:     "cons",
	ListType:     "list",
	FuncType:     "function",
	MacroType:    "macro",
	PortType:     "port",
	ModuleType:   "module",
}

// Get the name of a given Type
func TypeName(t Type) string {
	s, ok := TypeNameMap[t]
	if ok == false {
		panic(fmt.Sprintf("Unknown type: %d", t))
	}
	return s
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

// Check if the address if a > b.
func PtrGt(a, b interface{}) bool {
	return fmt.Sprintf("%p", a) > fmt.Sprintf("%p", b)
}

// Check if the address if a < b.
func PtrLt(a, b interface{}) bool {
	return fmt.Sprintf("%p", a) < fmt.Sprintf("%p", b)
}
