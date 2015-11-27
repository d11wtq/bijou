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
}

// Call this procedure with the given arguments
func (proc *Proc) Call(envc Env, args Sequence) (Value, error) {
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

	return EvalDo(env, proc.Body)
}
