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
	orderedSymbols     = []RomanSymbol{M, D, C, L, X, V, I}
	emptyNumber        = []RomanSymbol{}
	invalidCharError   = errors.New("invalid character given")
	invalidStringError = errors.New("invalid roman number string passed")
)

func symbolOrder(symbols []RomanSymbol, symbol RomanSymbol) int {
	for i, s := range symbols {
		if symbol == s {
			return i
		}
	}

	return 0
}

func getNextLowerSymbol(symbols []RomanSymbol, currentIndex int) RomanSymbol {
	var current = symbols[currentIndex]
	if current != D && current != L && current != V && current != I {
		return symbols[currentIndex+2]
	} else if current != I {
		return symbols[currentIndex+1]
	}

	return O
}

// IsValid returns true only if the sequence of RomanSymbols contained in the number is a
// valid roman number
func (n RomanNumber) IsValid() bool {
	var (
		syms = []RomanSymbol(n)
		length = len(syms)
		repeated int
		prevSym RomanSymbol = O
	)

	for i, sym := range syms {
		if prevSym != sym {
			repeated = 0
		} else {
			repeated++
		}

		// A symbol can be repeated up to 3 times except M
		isTooRepeated := repeated > 2 && sym != M
		// L, D and V can't repeate more than once
		isRepeatedAndShouldnt := repeated > 0 && (sym == L || sym == D || sym == V)
		// N cant be lower than N+1 unless N can be subtracted from N+1
		isValidSubtractor := length > i+1 && sym < syms[i+1] && repeated == 0 && getNextLowerSymbol(orderedSymbols, symbolOrder(orderedSymbols, syms[i+1])) == sym
		isSubtractor := length > i+1 && sym < syms[i+1]

		if isTooRepeated || isRepeatedAndShouldnt || (!isValidSubtractor && isSubtractor) || (!isValidSubtractor && sym < prevSym && prevSym != O && isSubtractor) || (isValidSubtractor && length > i+2 && sym == syms[i+2]) || (i-2 >= 0 && syms[i-2] < sym) {
			return false
		}

		prevSym = sym
	}

	return true
}

// Value returns the int32 value of the roman number
func (n RomanNumber) Value() int32 {
	var value int32
	syms := []RomanSymbol(n)
	length := len(syms)

	for i, sym := range syms {
		if i+1 < length && syms[i+1] > sym {
			value -= int32(sym)
		} else {
			value += int32(sym)
		}
	}

	return value
}

// FromInt returns the RomanNumber value of the given number
func FromInt(n int32) RomanNumber {
	var num []RomanSymbol

	for i, sym := range orderedSymbols {
		for n >= int32(sym) {
			n = n - int32(sym)
			num = append(num, sym)
		}

		nextLowerSym := getNextLowerSymbol(orderedSymbols, i)
		if nextLowerSym != O && n >= int32(sym)-int32(nextLowerSym) {
			n = n - int32(sym) + int32(nextLowerSym)
			num = append(num, nextLowerSym)
			num = append(num, sym)
		}
	}

	return RomanNumber(num)
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

	return O, invalidCharError
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
