//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package superscript

import (
	"strconv"
)

// Rune returns the superscript version of the input rune or rune
// itself if no superscript version is defined.
func Rune(r rune) rune {
	switch r {
	case '0':
		return 0x2070
	case '1':
		return 0x0b9
	case '2':
		return 0x00b2
	case '3':
		return 0x00b3
	case '4':
		return 0x2074
	case '5':
		return 0x2075
	case '6':
		return 0x2076
	case '7':
		return 0x2077
	case '8':
		return 0x2078
	case '9':
		return 0x2079
	case '+':
		return 0x207a
	case '-':
		return 0x207b
	case '=':
		return 0x207c
	case '(':
		return 0x207d
	case ')':
		return 0x207e
	case 'n':
		return 0x207f
	case 'i':
		return 0x2071
	default:
		return r
	}
}

// Itoa returns a superscript string representation of the argument
// number.
func Itoa(i int) string {
	var result []rune

	input := strconv.Itoa(i)
	for _, r := range input {
		result = append(result, Rune(r))
	}
	return string(result)
}
