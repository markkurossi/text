//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package color

import (
	"fmt"
	"image/color"
	"testing"
)

func TestNRGBA(t *testing.T) {
	c := color.NRGBA{
		R: 0xff,
		A: 0xff,
	}
	r, g, b, a := c.RGBA()

	fmt.Printf("R=%v, G=%v, B=%v, A=%v\n", r, g, b, a)
}

func TestBG(t *testing.T) {
	for _, scheme := range Schemes {
		testSchemeBG(t, scheme)
	}
}

func testSchemeBG(t *testing.T, scheme *Scheme) {
	fmt.Printf("Scheme %v:\n", scheme.Name)
	for idx, c := range scheme.Colors {
		l := luminance(c.BG)
		var bgName string
		if l > 128 {
			bgName = "light"
		} else {
			bgName = "dark"
		}
		fmt.Printf(" %2d: bg=%v[%v]\n", idx, bgName, l)
	}
}
