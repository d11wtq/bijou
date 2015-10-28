package runtime

import (
	"strconv"
	"unicode"
)

// Read an input string and convert it to an internal Value type
func Read(s string) (Value, string, error) {
Loop:
	for i, r := range s {
		switch {
		case unicode.IsSpace(r):
			// ignore
		case unicode.IsDigit(r):
			return ReadInt(s[i:])
		case (r == '('):
			return ReadList(s[i:])
		default:
			break Loop
		}
	}
	return Nil, s[:0], nil
}

// Read an input string and convert it to an Int type
func ReadInt(s string) (Value, string, error) {
	acc, rem := s[0:], s[len(s):]
	for i, r := range s {
		if !unicode.IsDigit(r) {
			acc, rem = s[0:i], s[i:]
			break
		}
	}

	n, err := strconv.ParseInt(acc, 10, 64)
	if err != nil {
		return nil, s, err
	}
	return Int(n), rem, nil
}

func ReadList(s1 string) (Value, string, error) {
	lst := EmptyList
	s2 := s1[1:]

	for {
	InnerLoop:
		for i, r := range s2 { // skip over the '('
			switch {
			case unicode.IsSpace(r):
				// ignore
			case (r == ')'):
				// skip over the ')'
				return lst.Reverse(), s2[i+1:], nil // FIXME: Optimize
			default:
				v, s, err := Read(s2[i:])
				if err != nil {
					return nil, s1, err
				}
				lst = lst.Cons(v)
				s2 = s
				break InnerLoop
			}
		}
	}

	return Nil, s1, nil
}
