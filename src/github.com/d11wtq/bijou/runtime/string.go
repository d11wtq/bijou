package runtime

import "fmt"

// String data type
type String string

func (s String) Eq(other Value) bool {
	return s == other
}

func (s String) Gt(other Value) bool {
	s2, ok := other.(String)
	if ok == false {
		return s.Type() > other.Type()
	}

	return s > s2
}

func (s String) Lt(other Value) bool {
	s2, ok := other.(String)
	if ok == false {
		return s.Type() < other.Type()
	}

	return s < s2
}

func (s String) Eval(env Env) (Value, error) {
	return s, nil
}

func (s String) Type() Type {
	return StringType
}

func (s String) String() string {
	acc := make([]rune, 0, len(s)*2)
	for _, r := range s {
		switch r {
		case '\r', '\n', '\t', '\\', '"':
			acc = append(acc, '\\')
		}

		switch r {
		case '\r':
			r = 'r'
		case '\n':
			r = 'n'
		case '\t':
			r = 't'
		}

		acc = append(acc, r)
	}

	return fmt.Sprintf(`"%s"`, string(acc))
}

func (s String) Head() Value {
	if s.Empty() {
		return Nil
	}

	return Int(s[0])
}

func (s String) Tail() Sequence {
	if s.Empty() {
		return s
	}

	return s[1:]
}

func (s String) Put(v Value) (Sequence, error) {
	s2, ok := v.(Int)
	if ok == false {
		return nil, &RuntimeError{"Can only put integer values to string"}
	}
	return String(append([]rune(string(s)), rune(s2))), nil
}

func (s String) Empty() bool {
	return len(s) == 0
}
