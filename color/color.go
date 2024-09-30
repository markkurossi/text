//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

// Package color implements color schemes from the [Colour schemes and
// templates](https://personal.sron.nl/~pault/).
package color

import (
	"image/color"
)

var (
	// Black defines black (0x000000) color.
	Black = color.NRGBA{
		A: 0xff,
	}
	// White defines while (0xffffff) color.
	White = color.NRGBA{
		R: 0xff,
		G: 0xff,
		B: 0xff,
		A: 0xff,
	}
)

// Color defines a color scheme color with foreground and background
// colors and an optional name.
type Color struct {
	Name string
	FG   color.NRGBA
	BG   color.NRGBA
}

func luminance(c color.NRGBA) uint8 {
	var sum float64

	sum += 0.2126 * float64(c.R)
	sum += 0.7152 * float64(c.G)
	sum += 0.0722 * float64(c.B)
	return uint8(sum)
}

// NewColor creates a new Color for foreground fg and background
// components r, g, b. The argument name specifies an optional name
// for the color.
func NewColor(fg color.NRGBA, r, g, b uint8, name string) *Color {
	return &Color{
		Name: name,
		FG:   fg,
		BG: color.NRGBA{
			R: r,
			G: g,
			B: b,
			A: 0xff,
		},
	}
}

// Scheme defines a color scheme.
type Scheme struct {
	Name    string
	Colors  []*Color
	BadData *Color
}

var (
	// Qualitative color schemes

	// Bright qualitative color scheme that is color-blind safe.
	Bright = &Scheme{
		Name: "Bright",
		Colors: []*Color{
			NewColor(Black, 0x44, 0x77, 0xAA, "blue"),
			NewColor(Black, 0x66, 0xCC, 0xEE, "cyan"),
			NewColor(Black, 0x22, 0x88, 0x33, "green"),
			NewColor(Black, 0xCC, 0xBB, 0x44, "yellow"),
			NewColor(Black, 0xEE, 0x66, 0x77, "red"),
			NewColor(Black, 0xAA, 0x33, 0x77, "purple"),
			NewColor(Black, 0xBB, 0xBB, 0xBB, "gray"),
		},
	}
	// HighContrast qualitative color scheme, an alternative to the
	// Bright scheme that is color-blind safe and optimized for
	// contrast.
	HighContrast = &Scheme{
		Name: "HighContrast",
		Colors: []*Color{
			NewColor(Black, 0xFF, 0xFF, 0xFF, "white"),
			NewColor(Black, 0xDD, 0xAA, 0x33, "yellow"),
			NewColor(Black, 0xBB, 0x55, 0x66, "red"),
			NewColor(Black, 0x00, 0x44, 0x88, "blue"),
			NewColor(White, 0x00, 0x00, 0x00, "black"),
		},
	}
	// Vibrant qualitative color scheme, an alternative to the Bright
	// scheme that is equally color-blind safe. It has been designed
	// for data visualization framework TensorBoard, built around
	// their signature orange FF7043. That color has been replaced
	// here to make it print-friendly.
	Vibrant = &Scheme{
		Name: "Vibrant",
		Colors: []*Color{
			NewColor(Black, 0x00, 0x77, 0xBB, "blue"),
			NewColor(Black, 0x33, 0xBB, 0xEE, "cyan"),
			NewColor(Black, 0x00, 0x99, 0x88, "teal"),
			NewColor(Black, 0xEE, 0x77, 0x33, "orange"),
			NewColor(Black, 0xCC, 0x33, 0x11, "red"),
			NewColor(Black, 0xEE, 0x33, 0x77, "magenta"),
			NewColor(Black, 0xBB, 0xBB, 0xBB, "gray"),
		},
		BadData: NewColor(Black, 0xDD, 0xDD, 0xDD, "pale-gray"),
	}
	// Muted qualitative color scheme, an alternative to the Bright
	// scheme that is equally color-blind safe with more colors, but
	// lacking a clear red or medium blue. Pale grey is meant for bad
	// data in maps.
	Muted = &Scheme{
		Name: "Muted",
		Colors: []*Color{
			NewColor(White, 0x33, 0x22, 0x88, "indigo"),
			NewColor(Black, 0x88, 0xCC, 0xEE, "cyan"),
			NewColor(Black, 0x44, 0xAA, 0x99, "teal"),
			NewColor(Black, 0x11, 0x77, 0x33, "green"),
			NewColor(Black, 0x99, 0x99, 0x33, "olive"),
			NewColor(Black, 0xDD, 0xCC, 0x77, "sand"),
			NewColor(Black, 0xCC, 0x66, 0x77, "rose"),
			NewColor(White, 0x88, 0x22, 0x55, "wine"),
			NewColor(Black, 0xAA, 0x44, 0x99, "purple"),
		},
	}
	// MediumContrast qualitative color scheme, an alternative to the
	// HighContrast scheme that is color-blind safe with more
	// colors. It is also optimized for contrast to work in a
	// monochrome printout, but the differences are inevitably
	// smaller.
	MediumContrast = &Scheme{
		Name: "MediumContrast",
		Colors: []*Color{
			NewColor(Black, 0xFF, 0xFF, 0xFF, "white"),
			NewColor(Black, 0xEE, 0xCC, 0x66, "light-yellow"),
			NewColor(Black, 0xEE, 0x99, 0xAA, "light-red"),
			NewColor(Black, 0x66, 0x99, 0xCC, "light-blue"),
			NewColor(Black, 0x99, 0x77, 0x00, "dark-yellow"),
			NewColor(Black, 0x99, 0x44, 0x55, "dark-red"),
			NewColor(White, 0x00, 0x44, 0x88, "dark-blue"),
			NewColor(White, 0x00, 0x00, 0x00, "black"),
		},
	}
	// Pale and dark qualitative color schemes where the colors are
	// not very distinct in either normal or color-blind vision; they
	// are not meant for lines or maps, but for marking text. Use the
	// pale colors for the background of black text, for example to
	// highlight cells in a table. One of the dark colors can be
	// chosen for text itself on a white background, for example when
	// a large block of text has to be marked. In both cases, the text
	// remains easily readable
	Pale = &Scheme{
		Name: "Pale",
		Colors: []*Color{
			NewColor(Black, 0xBB, 0xCC, 0xEE, "pale-blue"),
			NewColor(Black, 0xCC, 0xEE, 0xFF, "pale-cyan"),
			NewColor(Black, 0xCC, 0xDD, 0xAA, "pale-green"),
			NewColor(Black, 0xEE, 0xEE, 0xBB, "pale-yellow"),
			NewColor(Black, 0xFF, 0xCC, 0xCC, "pale-red"),
			NewColor(Black, 0xDD, 0xDD, 0xDD, "pale-gray"),
		},
	}
	// Dark qualitative color scheme for marking text.
	Dark = &Scheme{
		Name: "Dark",
		Colors: []*Color{
			NewColor(White, 0x22, 0x22, 0x55, "dark-blue"),
			NewColor(White, 0x22, 0x55, 0x55, "dark-cyan"),
			NewColor(White, 0x22, 0x55, 0x22, "dark-green"),
			NewColor(White, 0x66, 0x66, 0x33, "dark-yellow"),
			NewColor(White, 0x66, 0x33, 0x33, "dark-red"),
			NewColor(White, 0x55, 0x55, 0x55, "dark-gray"),
		},
	}
	// Light qualitative color scheme that is reasonably distinct in
	// both normal and color-blind vision. It was designed to fill
	// labelled cells with more and lighter colors than contained in
	// the Bright scheme, using more distinct colors than that in the
	// Pale scheme, but keeping black labels clearly
	// readable. However, it can also be used for general qualitative
	// maps.
	Light = &Scheme{
		Name: "Light",
		Colors: []*Color{
			NewColor(Black, 0x77, 0xAA, 0xDD, "light-blue"),
			NewColor(Black, 0x99, 0xDD, 0xFF, "light-cyan"),
			NewColor(Black, 0x44, 0xBB, 0x99, "mint"),
			NewColor(Black, 0xBB, 0xCC, 0x33, "pear"),
			NewColor(Black, 0xAA, 0xAA, 0x00, "olive"),
			NewColor(Black, 0xEE, 0xDD, 0x88, "light-yellow"),
			NewColor(Black, 0xEE, 0x88, 0x66, "orange"),
			NewColor(Black, 0xFF, 0xAA, 0xBB, "pink"),
			NewColor(Black, 0xDD, 0xDD, 0xDD, "pale-gray"),
		},
	}

	// Diverging color schemes

	// Sunset defines a diverging color scheme that also works in
	// color-blind vision. The colors can be used as given or linearly
	// interpolated.
	Sunset = &Scheme{
		Name: "Sunset",
		Colors: []*Color{
			NewColor(White, 0x36, 0x4B, 0x9A, ""),
			NewColor(Black, 0x4A, 0x7B, 0xB7, ""),
			NewColor(Black, 0x6E, 0xA6, 0xCD, ""),
			NewColor(Black, 0x98, 0xCA, 0xE1, ""),
			NewColor(Black, 0xC2, 0xE4, 0xEF, ""),
			NewColor(Black, 0xEA, 0xEC, 0xCC, ""),
			NewColor(Black, 0xFE, 0xDA, 0x8B, ""),
			NewColor(Black, 0xFD, 0xB3, 0x66, ""),
			NewColor(Black, 0xF6, 0x7E, 0x4B, ""),
			NewColor(Black, 0xDD, 0x3D, 0x2D, ""),
			NewColor(White, 0xA5, 0x00, 0x26, ""),
		},
		BadData: NewColor(Black, 0xFF, 0xFF, 0xFF, "white"),
	}
	// Nightfall diverging color scheme that also works in color-blind
	// vision. The colors are linearly interpolated or every second
	// color is skipped when used as given.
	Nightfall = &Scheme{
		Name: "Nightfall",
		Colors: []*Color{
			NewColor(White, 0x12, 0x5A, 0x56, ""),
			NewColor(White, 0x00, 0x76, 0x7B, ""),
			NewColor(Black, 0x23, 0x8F, 0x9D, ""),
			NewColor(Black, 0x42, 0xA7, 0xC6, ""),
			NewColor(Black, 0x60, 0xBC, 0xE9, ""),
			NewColor(Black, 0x9D, 0xCC, 0xEF, ""),
			NewColor(Black, 0xC6, 0xDB, 0xED, ""),
			NewColor(Black, 0xDE, 0xE6, 0xE7, ""),
			NewColor(Black, 0xEC, 0xEA, 0xDA, ""),
			NewColor(Black, 0xF0, 0xE6, 0xB2, ""),
			NewColor(Black, 0xF9, 0xD5, 0x76, ""),
			NewColor(Black, 0xFF, 0xB9, 0x54, ""),
			NewColor(Black, 0xFD, 0x9A, 0x44, ""),
			NewColor(Black, 0xF5, 0x76, 0x34, ""),
			NewColor(Black, 0xE9, 0x4C, 0x1F, ""),
			NewColor(Black, 0xD1, 0x18, 0x07, ""),
			NewColor(White, 0xA0, 0x18, 0x13, ""),
		},
		BadData: NewColor(Black, 0xFF, 0xFF, 0xFF, "white"),
	}
	// BuRd diverging color scheme that also works in color-blind
	// vision. The colors can be used as given or linearly
	// interpolated.
	BuRd = &Scheme{
		Name: "BuRd",
		Colors: []*Color{
			NewColor(Black, 0x21, 0x66, 0xAC, ""),
			NewColor(Black, 0x43, 0x93, 0xC3, ""),
			NewColor(Black, 0x92, 0xC5, 0xDE, ""),
			NewColor(Black, 0xD1, 0xE5, 0xF0, ""),
			NewColor(Black, 0xF7, 0xF7, 0xF7, ""),
			NewColor(Black, 0xFD, 0xDB, 0xC7, ""),
			NewColor(Black, 0xF4, 0xA5, 0x82, ""),
			NewColor(Black, 0xD6, 0x60, 0x4D, ""),
			NewColor(White, 0xB2, 0x18, 0x2B, ""),
		},
		BadData: NewColor(Black, 0xFF, 0xEE, 0x99, ""),
	}
	// PRGn diverging color scheme that also works in color-blind
	// vision. The colors can be used as given or linearly
	// interpolated.
	PRGn = &Scheme{
		Name: "PRGn",
		Colors: []*Color{
			NewColor(White, 0x76, 0x2A, 0x83, ""),
			NewColor(Black, 0x99, 0x70, 0xAB, ""),
			NewColor(Black, 0xC2, 0xA5, 0xCF, ""),
			NewColor(Black, 0xE7, 0xD4, 0xE8, ""),
			NewColor(Black, 0xF7, 0xF7, 0xF7, ""),
			NewColor(Black, 0xD9, 0xF0, 0xD3, ""),
			NewColor(Black, 0xAC, 0xD3, 0x9E, ""),
			NewColor(Black, 0x5A, 0xAE, 0x61, ""),
			NewColor(Black, 0x1B, 0x78, 0x37, ""),
		},
		BadData: NewColor(Black, 0xFF, 0xEE, 0x99, ""),
	}

	// Sequential color schemes

	// YlOrBr sequential color scheme that also works in color-blind
	// vision. The colors can be used as given or linearly
	// interpolated.
	YlOrBr = &Scheme{
		Name: "YlOrBr",
		Colors: []*Color{
			NewColor(Black, 0xFF, 0xFF, 0xE5, ""),
			NewColor(Black, 0xFF, 0xF7, 0xBC, ""),
			NewColor(Black, 0xFE, 0xE3, 0x91, ""),
			NewColor(Black, 0xFE, 0xC4, 0x4F, ""),
			NewColor(Black, 0xFB, 0x9A, 0x29, ""),
			NewColor(Black, 0xEC, 0x70, 0x14, ""),
			NewColor(Black, 0xCC, 0x4C, 0x02, ""),
			NewColor(White, 0x99, 0x34, 0x04, ""),
			NewColor(White, 0x66, 0x25, 0x06, ""),
		},
		BadData: NewColor(Black, 0x88, 0x88, 0x88, ""),
	}
	// Iridescent sequential color scheme with a linearly varying
	// luminance that also works in color-blind vision. The colors
	// should be linearly interpolated, optionally extended towards
	// white and black.
	Iridescent = &Scheme{
		Name: "Iridescent",
		Colors: []*Color{
			NewColor(Black, 0xFE, 0xFB, 0xE9, ""),
			NewColor(Black, 0xFC, 0xF7, 0xD5, ""),
			NewColor(Black, 0xF5, 0xF3, 0xC1, ""),
			NewColor(Black, 0xEA, 0xF0, 0xB5, ""),
			NewColor(Black, 0xDD, 0xEC, 0xBF, ""),
			NewColor(Black, 0xD0, 0xE7, 0xCA, ""),
			NewColor(Black, 0xC2, 0xE3, 0xD2, ""),
			NewColor(Black, 0xB5, 0xDD, 0xD8, ""),
			NewColor(Black, 0xA8, 0xD8, 0xDC, ""),
			NewColor(Black, 0x9B, 0xD2, 0xE1, ""),
			NewColor(Black, 0x8D, 0xCB, 0xE4, ""),
			NewColor(Black, 0x81, 0xC4, 0xE7, ""),
			NewColor(Black, 0x7B, 0xBC, 0xE7, ""),
			NewColor(Black, 0x7E, 0xB2, 0xE4, ""),
			NewColor(Black, 0x88, 0xA5, 0xDD, ""),
			NewColor(Black, 0x93, 0x98, 0xD2, ""),
			NewColor(Black, 0x9B, 0x8A, 0xC4, ""),
			NewColor(Black, 0x9D, 0x7D, 0xB2, ""),
			NewColor(Black, 0x9A, 0x70, 0x9E, ""),
			NewColor(White, 0x90, 0x63, 0x88, ""),
			NewColor(White, 0x80, 0x57, 0x70, ""),
			NewColor(White, 0x68, 0x49, 0x57, ""),
			NewColor(White, 0x46, 0x35, 0x3A, ""),
		},
		BadData: NewColor(Black, 0x99, 0x99, 0x99, ""),
	}
	// Incandescent sequential color scheme with a linearly varying
	// luminance that also works in color-blind vision, but it is not
	// print-friendly. The colors should be linearly interpolated,
	// optionally extended towards white and black.
	Incandescent = &Scheme{
		Name: "Incandescent",
		Colors: []*Color{
			NewColor(Black, 0xCE, 0xFF, 0xFF, ""),
			NewColor(Black, 0xC6, 0xF7, 0xD6, ""),
			NewColor(Black, 0xA2, 0xF4, 0x9B, ""),
			NewColor(Black, 0xBB, 0xE4, 0x53, ""),
			NewColor(Black, 0xD5, 0xCE, 0x04, ""),
			NewColor(Black, 0xE7, 0xB5, 0x03, ""),
			NewColor(Black, 0xF1, 0x99, 0x03, ""),
			NewColor(Black, 0xF6, 0x79, 0x0B, ""),
			NewColor(Black, 0xF9, 0x49, 0x02, ""),
			NewColor(Black, 0xE4, 0x05, 0x15, ""),
			NewColor(White, 0xA8, 0x00, 0x03, ""),
		},
		BadData: NewColor(Black, 0x88, 0x88, 0x88, ""),
	}

	// DiscreteRainbow color scheme with 14 or 23 colors for maps. The
	// colors have to be used as given: do not interpolate.
	DiscreteRainbow = &Scheme{
		Name: "DiscreteRainbow",
		Colors: []*Color{
			NewColor(Black, 0xE8, 0xEC, 0xFB, ""),
			NewColor(Black, 0xD9, 0xCC, 0xE3, ""),
			NewColor(Black, 0xD1, 0xBB, 0xD7, ""),
			NewColor(Black, 0xCA, 0xAC, 0xCB, ""),
			NewColor(Black, 0xBA, 0x8D, 0xB4, ""),
			NewColor(Black, 0xAE, 0x76, 0xA3, ""),
			NewColor(Black, 0xAA, 0x6F, 0x9E, ""),
			NewColor(Black, 0x99, 0x4F, 0x88, ""),
			NewColor(Black, 0x88, 0x2E, 0x72, ""),
			NewColor(Black, 0x19, 0x65, 0xB0, ""),
			NewColor(Black, 0x43, 0x7D, 0xBF, ""),
			NewColor(Black, 0x52, 0x89, 0xC7, ""),
			NewColor(Black, 0x61, 0x95, 0xCF, ""),
			NewColor(Black, 0x7B, 0xAF, 0xDE, ""),
			NewColor(Black, 0x4E, 0xB2, 0x65, ""),
			NewColor(Black, 0x90, 0xC9, 0x87, ""),
			NewColor(Black, 0xCA, 0xE0, 0xAB, ""),
			NewColor(Black, 0xF7, 0xF0, 0x56, ""),
			NewColor(Black, 0xF7, 0xCB, 0x45, ""),
			NewColor(Black, 0xF6, 0xC1, 0x41, ""),
			NewColor(Black, 0xF4, 0xA7, 0x36, ""),
			NewColor(Black, 0xF1, 0x93, 0x2D, ""),
			NewColor(Black, 0xEE, 0x80, 0x26, ""),
			NewColor(Black, 0xE8, 0x60, 0x1C, ""),
			NewColor(Black, 0xE6, 0x55, 0x18, ""),
			NewColor(Black, 0xDC, 0x05, 0x0C, ""),
			NewColor(Black, 0xA5, 0x17, 0x0E, ""),
			NewColor(Black, 0x72, 0x19, 0x0E, ""),
			NewColor(Black, 0x42, 0x15, 0x0A, ""),
		},
		BadData: NewColor(Black, 0x77, 0x77, 0x77, ""),
	}
	// SmoothRainbow color scheme. The colors are meant to be linearly
	// interpolated.
	SmoothRainbow = &Scheme{
		Name: "SmoothRainbow",
		Colors: []*Color{
			NewColor(Black, 0xE8, 0xEC, 0xFB, ""),
			NewColor(Black, 0xDD, 0xD8, 0xEF, ""),
			NewColor(Black, 0xD1, 0xC1, 0xE1, ""),
			NewColor(Black, 0xC3, 0xA8, 0xD1, ""),
			NewColor(Black, 0xB5, 0x8F, 0xC2, ""),
			NewColor(Black, 0xA7, 0x78, 0xB4, ""),
			NewColor(Black, 0x9B, 0x62, 0xA7, ""),
			NewColor(Black, 0x8C, 0x4E, 0x99, ""),
			NewColor(Black, 0x6F, 0x4C, 0x9B, ""),
			NewColor(Black, 0x60, 0x59, 0xA9, ""),
			NewColor(Black, 0x55, 0x68, 0xB8, ""),
			NewColor(Black, 0x4E, 0x79, 0xC5, ""),
			NewColor(Black, 0x4D, 0x8A, 0xC6, ""),
			NewColor(Black, 0x4E, 0x96, 0xBC, ""),
			NewColor(Black, 0x54, 0x9E, 0xB3, ""),
			NewColor(Black, 0x59, 0xA5, 0xA9, ""),
			NewColor(Black, 0x60, 0xAB, 0x9E, ""),
			NewColor(Black, 0x69, 0xB1, 0x90, ""),
			NewColor(Black, 0x77, 0xB7, 0x7D, ""),
			NewColor(Black, 0x8C, 0xBC, 0x68, ""),
			NewColor(Black, 0xA6, 0xBE, 0x54, ""),
			NewColor(Black, 0xBE, 0xBC, 0x48, ""),
			NewColor(Black, 0xD1, 0xB5, 0x41, ""),
			NewColor(Black, 0xDD, 0xAA, 0x3C, ""),
			NewColor(Black, 0xE4, 0x9C, 0x39, ""),
			NewColor(Black, 0xE7, 0x8C, 0x35, ""),
			NewColor(Black, 0xE6, 0x79, 0x32, ""),
			NewColor(Black, 0xE4, 0x63, 0x2D, ""),
			NewColor(Black, 0xDF, 0x48, 0x28, ""),
			NewColor(Black, 0xDA, 0x22, 0x22, ""),
			NewColor(Black, 0xB8, 0x22, 0x1E, ""),
			NewColor(Black, 0x95, 0x21, 0x1B, ""),
			NewColor(Black, 0x72, 0x1E, 0x17, ""),
			NewColor(Black, 0x52, 0x1A, 0x13, ""),
		},
		BadData: NewColor(Black, 0x66, 0x66, 0x66, ""),
	}

	// Color scheme for ground cover.

	// AVHRR defines a color-blind safe color scheme for the AVHRR
	// global land cover classification. The colors have been numbered
	// with the data values.
	AVHRR = &Scheme{
		Name: "AVHRR",
		Colors: []*Color{
			NewColor(Black, 0x55, 0x66, 0xAA, "0"),
			NewColor(Black, 0x11, 0x77, 0x33, "1"),
			NewColor(Black, 0x44, 0xAA, 0x66, "3"),
			NewColor(Black, 0x55, 0xAA, 0x22, "5"),
			NewColor(Black, 0x66, 0x88, 0x22, "2"),
			NewColor(Black, 0x99, 0xBB, 0x55, "4"),
			NewColor(Black, 0x55, 0x88, 0x77, "6"),
			NewColor(Black, 0x88, 0xBB, 0xAA, "7"),
			NewColor(Black, 0xAA, 0xDD, 0xCC, "10"),
			NewColor(Black, 0x44, 0xAA, 0x88, "11"),
			NewColor(Black, 0xDD, 0xCC, 0x66, "8"),
			NewColor(Black, 0xFF, 0xDD, 0x44, "9"),
			NewColor(Black, 0xFF, 0xEE, 0x88, "12"),
			NewColor(Black, 0xBB, 0x00, 0x11, "13"),
		},
	}

	// Schemes list all color schemes.
	Schemes = []*Scheme{
		Bright, HighContrast, Vibrant, Muted, MediumContrast, Pale, Dark, Light,

		Sunset, Nightfall, BuRd, PRGn,

		YlOrBr, Iridescent, Incandescent, DiscreteRainbow, SmoothRainbow,

		AVHRR,
	}
)
