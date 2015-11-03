package runtime

// Program data type
type Type uint8

const (
	// Empty value, nil
	NilType Type = iota
	// Boolean true/false
	BooleanType
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
	// Equality comparison
	Eq(Value) bool
	// Eval to the lowest level value
	Eval(env Env) (Value, error)
	// Eval to the lowest level value
	Type() Type
}

// Value that can be invoked
type Callable interface {
	Value
	// FIXME: *List here is too specific
	Call(args *List) (Value, error)
}

// Runtime environment
type Env interface {
	// Define a new symbol
	Def(string, Value)
	// Lookup a variable
	Get(string) (Value, bool)
	// Introduce a new scope
	Extend() Env
}
