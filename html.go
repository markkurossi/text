//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package text

import (
	"html"
)

// HTML creates HTML representation of the text.
func HTML(text *Text) string {
	var str string

	for _, span := range text.Spans {
		if span.Bold {
			str += "<b>"
		}
		if span.Oblique {
			str += "<i>"
		}
		str += html.EscapeString(span.Content)
		if span.Oblique {
			str += "</i>"
		}
		if span.Bold {
			str += "</b>"
		}
	}

	return str
}
