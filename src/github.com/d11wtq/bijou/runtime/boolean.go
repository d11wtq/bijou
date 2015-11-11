package runtime

// Boolean data type
type Boolean bool

// Representation of true
var True = (Boolean)(true)

// Representation of false
var False = (Boolean)(false)

func (v Boolean) Eq(other Value) bool {
	return v == other
}

func (v Boolean) Gt(other Value) bool {
	y, ok := other.(Boolean)
	if ok == false {
		return v.Type() > other.Type()
	}

	return v == True && y == False
}

func (v Boolean) Lt(other Value) bool {
	y, ok := other.(Boolean)
	if ok == false {
		return v.Type() < other.Type()
	}

	return v == False && y == True
}

func (v Boolean) Type() Type {
	return BooleanType
}

func (v Boolean) Eval(env Env) (Value, error) {
	return v, nil
}

func (v Boolean) String() string {
	if bool(v) {
		return "true"
	} else {
		return "false"
	}
}
