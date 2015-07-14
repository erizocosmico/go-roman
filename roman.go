package roman

import "errors"

// RomanSymbol is the representation of one of the symbols
// used to represent a roman number
type RomanSymbol int
// RomanNumber represents a sequence of valid RomanSymbols
type RomanNumber []RomanSymbol

const (
	O = 0
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

var (
	emptyNumber        = []RomanSymbol{}
	invalidCharError   = errors.New("invalid character given")
	invalidStringError = errors.New("invalid roman number string passed")
)

// IsValid returns true only if the sequence of RomanSymbols contained in the number is a
// valid roman number
func (n RomanNumber) IsValid() bool {
	// TODO: Implement
	return true
}

// Value returns the int32 value of the roman number
func (n RomanNumber) Value() int32 {
	// TODO: Implement
	return int32(0)
}

// FromInt returns the RomanNumber value of the given number
func FromInt(i int32) (RomanNumber, error) {
	// TODO: Implement
	return emptyNumber, nil
}

// FromString returns the RomanNumber after parsing the given string
func FromString(s string) (RomanNumber, error) {
	var num = make([]RomanSymbol, len(s))

	for i, c := range s {
		sym, err := SymbolFromRune(c)
		if err != nil {
			return emptyNumber, invalidStringError
		}

		num[i] = sym
	}

	if !RomanNumber(num).IsValid() {
		return emptyNumber, invalidStringError
	}

	return num, nil
}

// SymbolFromRune returns the RomanSymbol that matches
// the given rune
func SymbolFromRune(r rune) (RomanSymbol, error) {
	switch r {
	case 73:
		return I, nil
	case 86:
		return V, nil
	case 88:
		return X, nil
	case 76:
		return L, nil
	case 67:
		return C, nil
	case 68:
		return D, nil
	case 77:
		return M, nil
	}

	return O, errors.New("invalid character given")
}

// SymbolFromString returns the RomanSymbol that matches
// the given character. Only the first character will be
// used if a string with more than one character is passed
func SymbolFromString(s string) (RomanSymbol, error) {
	for _, r := range s {
		return SymbolFromRune(r)
	}

	return O, invalidCharError
}
