// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dyff

import (
	"github.com/lucasb-eyer/go-colorful"
)

// Colorizer is the interface for applying colors to diff output. Implement this
// to replace the default bunt-based coloring with an alternative (e.g. ANSI).
type Colorizer interface {
	Green(text string) string
	Greenf(format string, a ...interface{}) string
	Red(text string) string
	Redf(format string, a ...interface{}) string
	Yellowf(format string, a ...interface{}) string
	LightGreen(text string) string
	LightRed(text string) string
	DimGray(text string) string
	Bold(text string) string
	Italic(text string) string
	BoldGreen(text string) string
	BoldRed(text string) string
	StylizeHeader(header string) string
	YAMLInRedishColors(input interface{}, useIndentLines bool) (string, error)
	YAMLInGreenishColors(input interface{}, useIndentLines bool) (string, error)
}

// BuntColorizer is the default Colorizer using bunt and go-colorful.
type BuntColorizer struct{}

var (
	additionGreen      = hexColor("#58BF38")
	modificationYellow = hexColor("#C7C43F")
	removalRed         = hexColor("#B9311B")
)

func hexColor(hex string) colorful.Color { _ = "STUB: not implemented"; return *new(colorful.Color) }

func render(format string, a ...interface{}) string { _ = "STUB: not implemented"; return "" }

func colored(color colorful.Color, text string) string { _ = "STUB: not implemented"; return "" }

func coloredf(color colorful.Color, format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func defaultColorizer() Colorizer { _ = "STUB: not implemented"; return *new(Colorizer) }

func (c *BuntColorizer) Green(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) Greenf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *BuntColorizer) Red(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) Redf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *BuntColorizer) Yellowf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *BuntColorizer) LightGreen(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) LightRed(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) DimGray(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) Bold(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) Italic(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) BoldGreen(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) BoldRed(text string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) StylizeHeader(header string) string { _ = "STUB: not implemented"; return "" }

func (c *BuntColorizer) YAMLInRedishColors(input interface{}, useIndentLines bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *BuntColorizer) YAMLInGreenishColors(input interface{}, useIndentLines bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
