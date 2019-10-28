package c15prop_based_tests

import (
	"strings"
)

type RomanNumeral struct {
	Value uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var romanNumerals = RomanNumerals {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol  {
			return s.Value
		}
	}

	return 0
}

func ConvertToArabic(roman string) uint16 {
	total := uint16(0)

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look to next symbol if we can and the current symbol is base 10 only
		if couldBeSubtractive(i, roman, symbol) {
			// get the value of two character string
			if value := romanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++ // move to next loop
			} else {
				total += romanNumerals.ValueOf(symbol)
			}
		} else {
			total += romanNumerals.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubtractive(i int, roman string, currentSymbol uint8) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return i+1 < len(roman) && isSubtractiveSymbol
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	//for arabic > 0 {
	//	switch {
	//	case arabic > 9:
	//		result.WriteString("X")
	//		arabic -= 10
	//	case arabic > 8:
	//		result.WriteString("IX")
	//		arabic -= 9
	//	case arabic > 4:
	//		result.WriteString("V")
	//		arabic -= 5
	//	case arabic > 3:
	//		result.WriteString("IV")
	//		arabic -= 4
	//	default:
	//		result.WriteString("I")
	//		arabic--
	//	}
	//}

	//for i := arabic; i > 0; i-- {
	//	if i == 5 {
	//		result.WriteString("V")
	//		break
	//	}
	//
	//	if i == 4 {
	//		result.WriteString("IV")
	//		break
	//	}
	//
	//	result.WriteString("I")
	//}

	return result.String()
}
