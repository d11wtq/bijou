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

// Abstract sequence interface
type Seq interface {
	// Seq is a first-class Value too
	Value
	// Get the value at the end of the Seq
	Head() Value
	// Pop off the value at the end of the Seq
	Tail() Seq
	// Add a new value at the end of the Seq
	Cons(v Value) Seq
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
