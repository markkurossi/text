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
	"image/draw"
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
	tileR   = 50

	rPlus  = float64(tileR) + 0.5
	rMinus = float64(tileR) - 0.5

	aaOut    = rPlus * rPlus
	aaIn     = rMinus * rMinus
	aaBorder = aaOut - aaIn

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
	widthPx := width*tileR*2 + (width-1)*padX + 2*marginX

	height := len(cs.Schemes)
	heightPx := height*tileR*2 + (height-1)*padY + 2*marginY

	img := image.NewRGBA(image.Rect(0, 0, widthPx, heightPx))

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
		printTile(img, row, col, c, false)
	}
	if scheme.BadData != nil {
		printTile(img, row, len(scheme.Colors), scheme.BadData, true)
	}
}

func printTile(img *image.RGBA, row, col int, c *cs.Color, circ bool) {
	bg := NRGBAToRGBA(c.BG)
	fg := NRGBAToRGBA(c.FG)

	yOfs := marginY + row*(tileR*2+padY)
	xOfs := marginX + col*(tileR*2+padX)

	if circ {
		center := image.Point{
			X: xOfs + tileR,
			Y: yOfs + tileR,
		}

		draw.DrawMask(img, img.Bounds(), &image.Uniform{bg}, image.ZP,
			&circle{center, tileR}, image.ZP, draw.Over)
	} else {
		draw.Draw(img, image.Rect(xOfs, yOfs, xOfs+tileR*2, yOfs+tileR*2),
			&image.Uniform{bg}, image.ZP, draw.Src)
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
		drawString(img, xOfs+tileR, yOfs+tileR*2/4, fg2, c.Name)
		decY = yOfs + tileR*2/4*2
		hexY = yOfs + tileR*2/4*3
	} else {
		decY = yOfs + tileR*2/3*1
		hexY = yOfs + tileR*2/3*2
	}

	drawString(img, xOfs+tileR, decY, fg2,
		fmt.Sprintf("%d,%d,%d", c.BG.R, c.BG.G, c.BG.B))
	drawString(img, xOfs+tileR, hexY, fg2,
		fmt.Sprintf("%02X%02X%02X", c.BG.R, c.BG.G, c.BG.B))
}

// NRGBAToRGBA converts NRGBA color to RGBA color.
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

	point := fixed.Point26_6{
		X: fixed.I(x - wPx/2),
		Y: fixed.I(y + hPx),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(str)
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5

	p := xx*xx + yy*yy
	if p >= aaOut {
		return color.Alpha{0}
	}
	if p <= aaIn {
		return color.Alpha{255}
	}
	f := p - aaIn
	r := f / aaBorder

	return color.Alpha{255 - uint8(255*r)}
}
