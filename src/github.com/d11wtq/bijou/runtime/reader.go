package runtime

import (
	"fmt"
	"strconv"
	"unicode"
)

// Unicode RangeTable definining the delimiters
var Delim = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'"', '"', 1},
		{'(', ')', 1},
		{';', ';', 1},
	},
}

// Handle the unexpected EOF error generation
func UnexpectedEOF(s string) (Value, string, error) {
	return nil, s, &RuntimeError{"Unexpected EOF while reading"}
}

// Handle the unexpected delimiter error generation
func UnexpectedChar(s string) (Value, string, error) {
	return nil, s, &RuntimeError{
		fmt.Sprintf("Unexpected %c while reading", s[0]),
	}
}

// Read an input string and convert it to an internal Value type
func Read(s string) (Value, string, error) {
	s = SkipBlank(s)
	for i, r := range s {
		switch {
		case (r == '+'), (r == '-'):
			return ReadMaybeInt(s[i:])
		case unicode.IsDigit(r):
			return ReadInt(s[i:])
		case (r == '('):
			return ReadList(s[i:])
		case (r == '"'):
			return ReadString(s[i:])
		case (r == '\''):
			return ReadQuoted(s[i:])
		case unicode.Is(Delim, r):
			return UnexpectedChar(s)
		default:
			return ReadSymbol(s[i:])
		}
	}
	return UnexpectedEOF(s)
}

// Read an input string prefixed '+' or '-' and convert it to Int or Symbol
func ReadMaybeInt(s string) (Value, string, error) {
	for _, r := range s[1:] {
		if unicode.IsDigit(r) {
			return ReadInt(s)
		}
		break
	}

	return ReadSymbol(s)
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

Loop:
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
				continue Loop
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
	acc := EmptyList
	// skip over the '('
	s2 := s1[1:]

OuterLoop:
	for {
		s2 = SkipBlank(s2)
		for i, r := range s2 {
			switch {
			case (r == ')'):
				// skip over the ')'
				return acc, s2[i+1:], nil
			default:
				v, rem, err := Read(s2[i:])
				if err != nil {
					return nil, s1, err
				}
				acc = acc.Append(v)
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
	return EmptyList.Append(Symbol("quote")).Append(v), rem, nil
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

// Skip over a comment, returning the remainder
func SkipComment(s string) string {
	rem := s[len(s):]
	for i, r := range s {
		if r == '\n' || r == '\r' {
			rem = s[i:]
			break
		}
	}
	return rem
}

// Skip over whitespace, returning the remainder
func SkipSpace(s string) string {
	rem := s[len(s):]
	for i, r := range s {
		if !unicode.IsSpace(r) {
			rem = s[i:]
			break
		}
	}
	return rem
}

// Skip over all whitespace and comments
func SkipBlank(s string) string {
	for _, r := range s {
		switch {
		case unicode.IsSpace(r):
			return SkipBlank(SkipSpace(s))
		case (r == ';'):
			return SkipBlank(SkipComment(s))
		}
		break
	}
	return s
}

// Read all forms in the input string and return a list of Values
func ReadSrc(s string) (*List, error) {
	acc := EmptyList

Loop:
	for {
		s = SkipBlank(s)
		if s == "" {
			break
		}

		v, rem, err := Read(s)
		if err != nil {
			return nil, err
		}
		acc = acc.Append(v)
		s = rem
		continue Loop
	}

	return acc, nil
}
