package propertyBasedText

import (
	"fmt"
	"testing"
	"testing/quick"
)

type Case struct {
	Num   uint16
	Roman string
}

var cases = []Case{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{1984, "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Num, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Num)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToNum(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Num), func(t *testing.T) {
			got := ConvertToNum(test.Roman)
			if got != test.Num {
				t.Errorf("got %d want %d", got, test.Num)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num uint16) bool {
		roman := ConvertToRoman(num)
		fromRoman := ConvertToNum(roman)
		return fromRoman == num
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
