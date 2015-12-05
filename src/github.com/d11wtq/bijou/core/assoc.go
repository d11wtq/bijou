package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Lookup key in value.
// Usage: (lookup value key)
func Lookup(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var value, key runtime.Value
	if err := runtime.ReadArgs(args, &value, &key); err != nil {
		return nil, err
	}
	assoc, ok := value.(runtime.Associative)
	if ok == false {
		return nil, runtime.BadType(runtime.AssocType, value.Type())
	}

	return assoc.Lookup(key)
}
