package runtime

// Program data type
type Type uint8

const (
	// Empty value, nil
	NilType Type = iota
	// Boolean true/false
	BooleanType
	// Integers
	IntType
	// Unique symbols; variable names, function names
	SymbolType
	// Generic sequences
	SequenceType
	// Strings
	StringType
	// Cons sequences
	ConsType
	// Linked lists
	ListType
	// Functions
	FuncType
	// Macros
	MacroType
	// Input/output ports
	PortType
)

// Program runtime value
type Value interface {
	// Eval to the lowest level value
	Eval(env Env) (Value, error)
	// Eval to the lowest level value
	Type() Type
	// Provide a representation of this value as string
	String() string
}

// Runtime environment
type Env interface {
	// Define a new symbol
	Def(string, Value) error
	// Lookup a variable
	Get(string) (Value, bool)
	// Introduce a new scope
	Extend() Env
}
