package runtime

import (
	"fmt"
)

// Concrete implementation of Env
type Scope struct {
	// Parent scope, if not the root scope
	Parent Env
	// List of defined symbols in the current scope
	Values map[string]Value
}

// Create a new Scope, extending the parent env.
func NewScope(parent Env) Env {
	return &Scope{
		Parent: parent,
		Values: make(map[string]Value),
	}
}

// Define a new symbol in the current scope
func (s *Scope) Def(k string, v Value) error {
	x, ok := s.Values[k]
	if ok && !Eq(x, v) {
		return &RuntimeError{
			fmt.Sprintf("Attempted to def %s more than once", k),
		}
	}

	s.Values[k] = v
	return nil
}

// Perform a pattern matching bind
func (s *Scope) Bind(binding, data Value) error {
	return Bind(binding, data, s)
}

// Lookup this symbol in the current scope (non recursive)
func (s *Scope) Get(k string) (Value, bool) {
	v, ok := s.Values[k]
	return v, ok
}

// Resolve a symbol within the current scope, or any parent scope
func (s *Scope) Resolve(k string) (Value, bool) {
	v, ok := s.Values[k]
	if !ok {
		if s.Parent != nil {
			return s.Parent.Resolve(k)
		}
	}
	return v, ok
}

// Create a new child scope
func (s *Scope) Extend() Env {
	return NewScope(s)
}
