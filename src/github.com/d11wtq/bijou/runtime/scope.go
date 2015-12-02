package runtime

import (
	"fmt"
)

func AssertDefinable(env Env, k string, v Value) error {
	x, ok := env.Get(k)
	if ok && !Eq(x, v) {
		return &RuntimeError{
			fmt.Sprintf("Attempted to def %s more than once", k),
		}
	}
	return nil
}

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
	if err := AssertDefinable(s, k, v); err != nil {
		return err
	}

	s.Values[k] = v
	return nil
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

// Create a new transaction in this scope
func (s *Scope) Transaction() EnvTransaction {
	return &ScopeTransaction{
		Scope:  s,
		Staged: make(map[string]Value),
	}
}

// A transaction in Scope
type ScopeTransaction struct {
	*Scope
	// Staged values defined during this transaction
	Staged map[string]Value
}

// Define a value, only visible inside this transaction
func (tx *ScopeTransaction) Def(k string, v Value) error {
	if err := AssertDefinable(tx, k, v); err != nil {
		return err
	}

	tx.Staged[k] = v
	return nil
}

// Get a value either from this transaction, or its scope
func (tx *ScopeTransaction) Get(k string) (Value, bool) {
	v, ok := tx.Staged[k]
	if ok == false {
		v, ok = tx.Scope.Get(k)
	}
	return v, ok
}

// Resolve a value recursively in this environment and its scope
func (tx *ScopeTransaction) Resolve(k string) (Value, bool) {
	v, ok := tx.Staged[k]
	if !ok {
		return tx.Scope.Resolve(k)
	}
	return v, ok
}

// Atomically commit all values in this transaction to the scope
func (tx *ScopeTransaction) Commit() {
	for k, v := range tx.Scope.Values {
		tx.Staged[k] = v
	}
	tx.Scope.Values = tx.Staged
}
