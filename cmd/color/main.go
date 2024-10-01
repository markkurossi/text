//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	cs "github.com/markkurossi/text/color"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	marginX = 10
	marginY = 10
	padX    = 0
	padY    = 10
	tileW   = 100
	tileH   = 100
)

func main() {
	var width int

	for _, scheme := range cs.Schemes {
		w := len(scheme.Colors)
		if scheme.BadData != nil {
			w++
		}
		if w > width {
			width = w
		}
	}
	widthPx := width*tileW + (width-1)*padX + 2*marginX

	height := len(cs.Schemes)
	heightPx := height*tileH + (height-1)*padY + 2*marginY

	img := image.NewRGBA(image.Rect(0, 0, widthPx, heightPx))

	white := color.RGBA{
		R: 0xff,
		G: 0xff,
		B: 0xff,
		A: 0xff,
	}

	for y := 0; y < heightPx; y++ {
		for x := 0; x < widthPx; x++ {
			img.SetRGBA(x, y, white)
		}
	}

	for row, scheme := range cs.Schemes {
		drawScheme(img, row, scheme)
	}

	f, err := os.Create("schemes.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}

func drawScheme(img *image.RGBA, row int, scheme *cs.Scheme) {
	yOfs := marginY + row*(tileH+padY)
	for col, c := range scheme.Colors {
		bg := NRGBAToRGBA(c.BG)
		fg := NRGBAToRGBA(c.FG)

		xOfs := marginX + col*(tileW+padX)
		for y := 0; y < tileH; y++ {
			for x := 0; x < tileW; x++ {
				img.SetRGBA(xOfs+x, yOfs+y, bg)
			}
		}

		drawString(img, xOfs+tileW/2, yOfs+tileH/2, fg, c.Name)
	}
}

func NRGBAToRGBA(c color.NRGBA) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

func drawString(img *image.RGBA, x, y int, c color.RGBA, str string) {
	wPx := len(str) * 7
	hPx := basicfont.Face7x13.Ascent / 2

	point := fixed.Point26_6{fixed.I(x - wPx/2), fixed.I(y + hPx)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(str)
}
