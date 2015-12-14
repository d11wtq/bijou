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
	VectorType:   "vector",
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
