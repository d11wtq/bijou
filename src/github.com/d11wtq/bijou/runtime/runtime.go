package runtime

// Program data type
type Type uint8

const (
	// Empty value, nil
	NilType Type = iota
	// Unique symbols; variable names, function names
	SymbolType
	// Integers
	IntType
	// Strings
	StringType
	// Linked lists
	ListType
	// Functions
	FuncType
)

// Program runtime value
type Value interface {
	// Eval to the lowest level value
	Eval(env Env) (Value, error)
	// Eval to the lowest level value
	Type() Type
}

// Runtime environment
type Env interface {
	// Define a new symbol
	Def(Symbol, Value)
	// Lookup a variable
	Get(Symbol) (Value, bool)
	// Introduce a new scope
	Extend() Env
}
