package runtime

import (
	"fmt"
)

// Procedure data type
type Proc struct {
	// Parameter list
	Params Sequence
	// Procedure body expressions
	Body Sequence
	// Closed environment
	Env Env
	// The name to use for internal referencing
	Name string
}

// Perform a pattern match to bind args into an Env
func (proc *Proc) BindEnv(site Env, args Sequence) (Env, error) {
	env := proc.Env.Extend()
	err := Bind(proc.Params, args, env)
	if err != nil {
		top := BadPattern(proc.Params, args)

		if top.Error() != err.Error() {
			err = &PatternError{fmt.Sprintf("%s: %s", top, err)}
		}

		return nil, &ArgumentError{
			fmt.Sprintf("wrong arguments: %s", err),
		}
	}
	return env, nil
}

// Call this procedure with the given arguments
func (proc *Proc) DoBody(env Env) (Value, error) {
	return EvalDo(env, proc.Body)
}

// Call this procedure with the given environment and arguments
func (proc *Proc) Call(site Env, args Sequence) (Value, error) {
	env, err := proc.BindEnv(site, args)
	if err != nil {
		return nil, err
	}
	return proc.DoBody(env)
}
