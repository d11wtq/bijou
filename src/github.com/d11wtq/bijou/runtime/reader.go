package runtime

import (
	"fmt"
	"strconv"
	"unicode"
)

// Unicode RangeTable definining the delimiters
var Delim = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'(', ')', 1},
	},
}

// Handle the unexpected EOF error generation
func UnexpectedEOF(s string) (Value, string, error) {
	return nil, s, &RuntimeError{"Unexpected EOF while reading"}
}

// Read an input string and convert it to an internal Value type
func Read(s string) (Value, string, error) {
	for i, r := range s {
		switch {
		case unicode.IsSpace(r):
			// ignore
		case unicode.IsDigit(r):
			return ReadInt(s[i:])
		case (r == '('):
			return ReadList(s[i:])
		case (r == '"'):
			return ReadString(s[i:])
		case (r == '\''):
			return ReadQuoted(s[i:])
		default:
			return ReadSymbol(s[i:])
		}
	}
	return UnexpectedEOF(s)
}

// Read an input string and convert it to an Int type
func ReadInt(s string) (Value, string, error) {
	atom, rem := ReadAtom(s)
	n, err := strconv.ParseInt(atom, 10, 64)
	if err != nil {
		return nil, s, &RuntimeError{fmt.Sprintf("%s: Invalid syntax", atom)}
	}
	return Int(n), rem, nil
}

// Read an input string and convert it to a Symbol type
func ReadSymbol(s string) (Value, string, error) {
	var value Value
	atom, rem := ReadAtom(s)

	switch atom {
	case "nil":
		value = Nil
	case "true":
		value = True
	case "false":
		value = False
	default:
		value = Symbol(atom)
	}

	return value, rem, nil
}

// Read an input string and convert it to a String type
func ReadString(s1 string) (Value, string, error) {
	buf := make([]rune, 0, len(s1))

	// skip over the '"'
	s2 := s1[1:]

	escaped := false

OuterLoop:
	for i, r := range s2 {
		if escaped {
			escaped = false

			switch r {
			case 'r':
				r = '\r'
			case 'n':
				r = '\n'
			case 't':
				r = '\t'
			}
		} else {
			switch r {
			case '\\':
				escaped = true
				continue OuterLoop
			case '"':
				// skip over the '"'
				return String(buf), s2[i+1:], nil
			}
		}

		buf = append(buf, r)
	}

	return UnexpectedEOF(s1)
}

// Read an input string and convert it to a *List type
func ReadList(s1 string) (Value, string, error) {
	lst := EmptyList
	// skip over the '('
	s2 := s1[1:]

OuterLoop:
	for {
		for i, r := range s2 {
			switch {
			case unicode.IsSpace(r):
				// ignore
			case (r == ')'):
				// skip over the ')'
				return lst.Reverse(), s2[i+1:], nil
			default:
				v, rem, err := Read(s2[i:])
				if err != nil {
					return nil, s1, err
				}
				lst = lst.Cons(v)
				s2 = rem
				continue OuterLoop
			}
		}

		return UnexpectedEOF(s1)
	}
}

// Read an input string and convert it to a (quote ...)
func ReadQuoted(s1 string) (Value, string, error) {
	// skip over the "'"
	s2 := s1[1:]
	v, rem, err := Read(s2)
	if err != nil {
		return nil, s1, err
	}
	return EmptyList.Cons(v).Cons(Symbol("quote")), rem, nil
}

// Read an atom string from the input string, returning it and the remainder
func ReadAtom(s string) (string, string) {
	acc, rem := s[0:], s[len(s):]
	for i, r := range s {
		if unicode.IsSpace(r) || unicode.Is(Delim, r) {
			acc, rem = s[0:i], s[i:]
			break
		}
	}

	return acc, rem
}

// Read all forms in the input string and return a list of Values
func ReadSrc(s string) (*List, error) {
	lst := EmptyList

OuterLoop:
	for s != "" {
		for i, r := range s {
			switch {
			case unicode.IsSpace(r):
				// ignore
			default:
				v, rem, err := Read(s[i:])
				if err != nil {
					return nil, err
				}
				lst = lst.Cons(v)
				s = rem
				continue OuterLoop
			}
		}
		break
	}
	return lst.Reverse(), nil
}
