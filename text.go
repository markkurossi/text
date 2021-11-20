//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package text

import (
	"fmt"
)

// Text represents a text as a collection of formatted spans with
// optional formatting information.
type Text struct {
	Spans []Span
}

// New creates a new text instance.
func New() *Text {
	return &Text{}
}

// Append appends the argument text to the text object.
func (text *Text) Append(t *Text) *Text {
	text.Spans = append(text.Spans, t.Spans...)
	return text
}

// Plain appends a plain text span to the text object.
func (text *Text) Plain(content string) *Text {
	text.Spans = append(text.Spans, Span{
		Content: content,
	})
	return text
}

// Bold appends a bold text span to the text object.
func (text *Text) Bold(content string) *Text {
	text.Spans = append(text.Spans, Span{
		Bold:    true,
		Content: content,
	})
	return text
}

// Oblique appends an oblique text span to the text object.
func (text *Text) Oblique(content string) *Text {
	text.Spans = append(text.Spans, Span{
		Oblique: true,
		Content: content,
	})
	return text
}

// BoldOblique appends a bold oblique text span to the text object.
func (text *Text) BoldOblique(content string) *Text {
	text.Spans = append(text.Spans, Span{
		Bold:    true,
		Oblique: true,
		Content: content,
	})
	return text
}

// Plainf appends a plain formatted text span to the text object.
func (text *Text) Plainf(format string, a ...interface{}) *Text {
	return text.Plain(fmt.Sprintf(format, a...))
}

// Boldf appends a bold formatted text span to the text object.
func (text *Text) Boldf(format string, a ...interface{}) *Text {
	return text.Bold(fmt.Sprintf(format, a...))
}

// Obliquef appends an oblique formatted text span to the text object.
func (text *Text) Obliquef(format string, a ...interface{}) *Text {
	return text.Oblique(fmt.Sprintf(format, a...))
}

// BoldObliquef appends a bold oblique text span to the text object.
func (text *Text) BoldObliquef(format string, a ...interface{}) *Text {
	return text.BoldOblique(fmt.Sprintf(format, a...))
}

// Link appends a hyperlink to the text object.
func (text *Text) Link(url string, link *Text) *Text {
	text.Spans = append(text.Spans, Span{
		Content: url,
		Link:    link,
	})
	return text
}

// Span implements a text span with formatting options.
type Span struct {
	Bold    bool
	Oblique bool
	Content string
	Link    *Text
}
