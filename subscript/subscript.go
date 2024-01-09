//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package subscript

import (
	"strconv"
)

// Rune returns the subscript version of the input rune or rune itself
// if no subscript version is defined.
func Rune(r rune) rune {
	switch r {
	case '0':
		return 0x2080
	case '1':
		return 0x2081
	case '2':
		return 0x2082
	case '3':
		return 0x2083
	case '4':
		return 0x2084
	case '5':
		return 0x2085
	case '6':
		return 0x2086
	case '7':
		return 0x2087
	case '8':
		return 0x2088
	case '9':
		return 0x2089
	case '+':
		return 0x208a
	case '-':
		return 0x208b
	case '=':
		return 0x208c
	case '(':
		return 0x208d
	case ')':
		return 0x208e
	case 'a':
		return 0x2090
	case 'e':
		return 0x2091
	case 'o':
		return 0x2092
	case 'x':
		return 0x2093
	case '\u0259':
		return 0x2094
	case 'h':
		return 0x2095
	case 'k':
		return 0x2096
	case 'l':
		return 0x2097
	case 'm':
		return 0x2098
	case 'n':
		return 0x2099
	case 'p':
		return 0x209a
	case 's':
		return 0x209b
	case 't':
		return 0x209c
	default:
		return r
	}
}

// Itoa returns a subscript string representation of the argument
// number.
func Itoa(i int) string {
	var result []rune

	input := strconv.Itoa(i)
	for _, r := range input {
		result = append(result, Rune(r))
	}
	return string(result)
}
