//
// Copyright (c) 2021-2024 Markku Rossi
//
// All rights reserved.
//

package sparkline

import (
	"fmt"
	"testing"
)

func TestSparkline(t *testing.T) {
	values := []int{83, 61, 33, 25, 12, 1, 0, 75}

	fmt.Printf("%s\n", New(nil))
	fmt.Printf("%s\n", New(values))
	fmt.Printf("%s\n", Range(0, 100, values))
	fmt.Printf("%s\n", Range(-100, 100, values))
	fmt.Printf("%s\n", Range(50, 50, values))
	fmt.Printf("%s\n", Range(25, 65, values))
	fmt.Printf("%s\n", Range(65, 25, values))
}
