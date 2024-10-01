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

	white = color.RGBA{0xff, 0xff, 0xff, 0xff}
	black = color.RGBA{0x00, 0x00, 0x00, 0xff}
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

	if false {
		for y := 0; y < heightPx; y++ {
			for x := 0; x < widthPx; x++ {
				img.SetRGBA(x, y, white)
			}
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
	for col, c := range scheme.Colors {
		printTile(img, row, col, c)
	}
	if scheme.BadData != nil {
		printTile(img, row, len(scheme.Colors), scheme.BadData)
	}
}

func printTile(img *image.RGBA, row, col int, c *cs.Color) {
	bg := NRGBAToRGBA(c.BG)
	fg := NRGBAToRGBA(c.FG)

	yOfs := marginY + row*(tileH+padY)
	xOfs := marginX + col*(tileW+padX)
	for y := 0; y < tileH; y++ {
		for x := 0; x < tileW; x++ {
			img.SetRGBA(xOfs+x, yOfs+y, bg)
		}
	}

	l := cs.Luminance(c.BG)
	var fg2 color.RGBA
	if l > 127 {
		fg2 = black
	} else {
		fg2 = white
	}
	_ = fg

	var decY, hexY int
	if len(c.Name) > 0 {
		drawString(img, xOfs+tileW/2, yOfs+tileH/4, fg2, c.Name)
		decY = yOfs + tileH/4*2
		hexY = yOfs + tileH/4*3
	} else {
		decY = yOfs + tileH/3*1
		hexY = yOfs + tileH/3*2
	}

	drawString(img, xOfs+tileW/2, decY, fg2,
		fmt.Sprintf("%d,%d,%d", c.BG.R, c.BG.G, c.BG.B))
	drawString(img, xOfs+tileW/2, hexY, fg2,
		fmt.Sprintf("%02X%02X%02X", c.BG.R, c.BG.G, c.BG.B))
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
	wPx := len(str) * basicfont.Face7x13.Advance
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
