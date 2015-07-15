package roman

import (
	"reflect"
	"testing"
)

func TestSymbolFromString(t *testing.T) {
	var cases = []struct {
		char   string
		result RomanSymbol
	}{
		{"I", I},
		{"V", V},
		{"X", X},
		{"L", L},
		{"C", C},
		{"D", D},
		{"M", M},
	}

	for _, _case := range cases {
		if s, err := SymbolFromString(_case.char); s != _case.result || err != nil {
			t.Errorf("expected %d to be %d", s, _case.result)
		}
	}

	if s, err := SymbolFromString("Z"); s != O || err == nil {
		t.Errorf("expected error not to be nil and result to be 0")
	}
}

func TestSymbolFromRune(t *testing.T) {
	var cases = []struct {
		char   rune
		result RomanSymbol
	}{
		{73, I},
		{86, V},
		{88, X},
		{76, L},
		{67, C},
		{68, D},
		{77, M},
	}

	for _, _case := range cases {
		if s, err := SymbolFromRune(_case.char); s != _case.result || err != nil {
			t.Errorf("expected %d to be %d", s, _case.result)
		}
	}

	if s, err := SymbolFromRune(90); s != O || err == nil {
		t.Errorf("expected error not to be nil and result to be 0")
	}
}

func TestFromString(t *testing.T) {
	var cases = []struct {
		input  string
		result RomanNumber
	}{
		{"I", []RomanSymbol{I}},
		{"MDCLXVI", []RomanSymbol{M, D, C, L, X, V, I}},
		{"VIII", []RomanSymbol{V, I, I, I}},
		{"LIX", []RomanSymbol{L, I, X}},
		{"KMLX", []RomanSymbol{}},
	}

	for _, _case := range cases {
		if n, _ := FromString(_case.input); !reflect.DeepEqual(n, _case.result) {
			t.Errorf("expected %v to be %v", n, _case.result)
		}
	}
}

func rn(s string) RomanNumber {
	r, _ := FromString(s)
	return r
}

func TestIsValid(t *testing.T) {
	var cases = []struct {
		num   RomanNumber
		valid bool
	}{
		{rn("I"), true},
		{rn("II"), true},
		{rn("III"), true},
		{rn("IIII"), false},
		{rn("IV"), true},
		{rn("V"), true},
		{rn("VV"), false},
		{rn("IVI"), false},
		{rn("VI"), true},
		{rn("VII"), true},
		{rn("VIII"), true},
		{rn("VIIII"), true},
		{rn("IX"), true},
		{rn("IXI"), false},
		{rn("XVI"), true},
		{rn("VX"), false},
		{rn("L"), true},
		{rn("XL"), true},
		{rn("XXL"), false},
		{rn("VL"), false},
		{rn("LL"), false},
		{rn("C"), true},
		{rn("XC"), true},
		{rn("LC"), false},
		{rn("IC"), false},
		{rn("CC"), true},
		{rn("CCC"), true},
		{rn("CD"), true},
		{rn("LD"), false},
		{rn("XD"), false},
		{rn("VD"), false},
		{rn("ID"), false},
		{rn("DD"), false},
		{rn("M"), true},
		{rn("MM"), true},
		{rn("MMM"), true},
		{rn("DM"), false},
		{rn("CM"), true},
		{rn("LM"), false},
		{rn("XM"), false},
		{rn("VM"), false},
		{rn("IM"), false},
		{rn("IMVX"), false},
		{rn("XXI"), false},
		{rn("MCMXLV"), true},
	}

	for _, c := range cases {
		if c.num.IsValid() != c.valid {
			if c.valid {
				t.Errorf("expected num %v to be valid", c.num)
			} else {
				t.Errorf("expected num %v to be invalid", c.num)
			}
		}
	}
}

func TestValue(t *testing.T) {
	var cases = []struct {
		num RomanNumber
		val int32
	}{
		{rn("I"), int32(1)},
		{rn("II"), int32(2)},
		{rn("III"), int32(3)},
		{rn("IV"), int32(4)},
		{rn("V"), int32(5)},
		{rn("VI"), int32(6)},
		{rn("VII"), int32(7)},
		{rn("VIII"), int32(8)},
		{rn("IX"), int32(9)},
		{rn("X"), int32(10)},
		{rn("XVI"), int32(16)},
		{rn("L"), int32(50)},
		{rn("XL"), int32(40)},
		{rn("C"), int32(100)},
		{rn("XC"), int32(90)},
		{rn("CC"), int32(200)},
		{rn("CCC"), int32(300)},
		{rn("D"), int32(5)},
		{rn("M"), int32(1000)},
		{rn("MM"), int32(2000)},
		{rn("MMM"), int32(3000)},
		{rn("CM"), int32(900)},
		{rn("XXI"), int32(21)},
		{rn("MCMXLV"), int32(1945)},
	}

	for _, c := range cases {
		val := c.num.Value()
		if val != c.val {
			t.Errorf("expected %d to be %d", val, c.val)
		}
	}
}

func TestFromInt(t *testing.T) {
	var cases = []struct {
		result RomanNumber
		num    int32
	}{
		{rn("I"), int32(1)},
		{rn("II"), int32(2)},
		{rn("III"), int32(3)},
		{rn("IV"), int32(4)},
		{rn("V"), int32(5)},
		{rn("VI"), int32(6)},
		{rn("VII"), int32(7)},
		{rn("VIII"), int32(8)},
		{rn("IX"), int32(9)},
		{rn("X"), int32(10)},
		{rn("XVI"), int32(16)},
		{rn("L"), int32(50)},
		{rn("XL"), int32(40)},
		{rn("C"), int32(100)},
		{rn("XC"), int32(90)},
		{rn("CC"), int32(200)},
		{rn("CCC"), int32(300)},
		{rn("D"), int32(500)},
		{rn("M"), int32(1000)},
		{rn("MM"), int32(2000)},
		{rn("MMM"), int32(3000)},
		{rn("CM"), int32(900)},
		{rn("XXI"), int32(21)},
		{rn("MCMXLV"), int32(1945)},
	}

	for _, c := range cases {
		n := FromInt(c.num)
		if !reflect.DeepEqual(n, c.result) {
			t.Errorf("expected %v to be %v", n, c.result)
		}
	}
}
