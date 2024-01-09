//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package superscript

import (
	"fmt"
	"testing"
)

func TestNumbers(t *testing.T) {
	for i := -9; i < 10; i++ {
		fmt.Printf("%v:\tnumber%s\n", i, Itoa(i))
	}
}

func TestSymbols(t *testing.T) {
	for _, r := range []rune{
		'+', '-', '=', '(', ')',
		'a', 'e', 'o', 'x', '\u0259', 'h', 'k', 'l', 'm', 'n', 'p', 's', 't'} {
		fmt.Printf("%v:\tsymbol%c\n", r, Rune(r))
	}
}
