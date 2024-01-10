//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package symbols

import (
	"fmt"
	"testing"
)

var symbols = []rune{
	LAMBDA,
	Lambda,
}

func TestSymbols(t *testing.T) {
	for _, sym := range symbols {
		fmt.Printf("symbol: %c\n", sym)
	}
}
