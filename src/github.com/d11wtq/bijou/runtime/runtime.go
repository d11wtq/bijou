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
	ConsType
	// Linked lists
	ListType
	// Functions
	FuncType
	// Macros
	MacroType
)

// Program runtime value
type Value interface {
	// Equality comparison
	Eq(Value) bool
	// Eval to the lowest level value
	Eval(env Env) (Value, error)
	// Eval to the lowest level value
	Type() Type
	// Provide a representation of this value as string
	String() string
}

// Value that can be invoked
type Callable interface {
	Value
	// Invoke this value with the given arguments
	Call(env Env, args Sequence) (Value, error)
	// True if arguments should not be evaluated
	IsMacro() bool
}

// Data structures that can be looped over
//
// The semantics of where Head(), Tail() and Put() operate are subject to the
// underlying data structures.
type Sequence interface {
	Value
	// Take the first element of the sequence
	Head() Value
	// Get everything except the first element
	Tail() Sequence
	// Add a new value to the sequence
	Put(Value) (Sequence, error)
	// True if there are no elements in the sequence
	Empty() bool
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
