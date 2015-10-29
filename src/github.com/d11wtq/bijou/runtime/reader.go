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
		default:
			return ReadSymbol(s[i:])
		}
	}
	return nil, s, &RuntimeError{"Unexpected EOF while reading"}
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
	atom, rem := ReadAtom(s)
	return Symbol(atom), rem, nil
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

		return nil, s1, &RuntimeError{"Unexpected EOF while reading"}
	}
}
