//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package text

import (
	"testing"
)

var tests = []struct {
	text *Text
	html string
}{
	{
		text: New().Plain("<Hello, world!>"),
		html: "&lt;Hello, world!&gt;",
	},
	{
		text: New().Bold("bold"),
		html: "<b>bold</b>",
	},
	{
		text: New().Oblique("oblique"),
		html: "<i>oblique</i>",
	},
	{
		text: New().BoldOblique("BoldOblique"),
		html: "<b><i>BoldOblique</i></b>",
	},
}

var htmls = []string{}

func TestHTML(t *testing.T) {
	for idx, test := range tests {
		html := HTML(test.text)
		if html != test.html {
			t.Errorf("%d HTML: got %s, expected %s", idx, html, test.html)
		}
	}
}
