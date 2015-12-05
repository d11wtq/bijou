package runtime

// A Module represents a group of Values in an Env
type Module struct {
	// The environment in which the values are contained
	Env Env
}

// Wrap a given Env in a Module.
func ModuleFrom(env Env) *Module {
	return &Module{env}
}

// Lookup a value in this module, or error if not found
func (module *Module) Lookup(key Value) (Value, error) {
	if key, ok := key.(Symbol); ok == true {
		return key.Eval(module.Env)
	}

	return nil, BadType(SymbolType, key.Type())
}

// Modules are callable and return their members by name
func (module *Module) Call(site Env, args Sequence) (Value, error) {
	var key Value
	if err := ReadArgs(args, &key); err != nil {
		return nil, err
	}

	return module.Lookup(key)
}

func (module *Module) Type() Type {
	return ModuleType
}

func (module *Module) Eval(env Env) (Value, error) {
	return module, nil
}

func (module *Module) String() string {
	return "#<module>"
}
