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
	"crypto/x509"
	"io"

	"github.com/gonvenience/ytbx"
	"github.com/sergi/go-diff/diffmatchpatch"
	yamlv3 "go.yaml.in/yaml/v3"
)

// stringWriter is the interface that wraps the WriteString method.
type stringWriter interface {
	WriteString(s string) (int, error)
}

// HumanReport is a reporter with human readable output in mind
type HumanReport struct {
	Report
	Indent                int
	UseIndentLines        bool
	MinorChangeThreshold  float64
	MultilineContextLines int
	NoTableStyle          bool
	DoNotInspectCerts     bool
	OmitHeader            bool
	UseGoPatchPaths       bool
	PrefixMultiline       bool
	Colorizer             Colorizer
}

func (report *HumanReport) colorizer() Colorizer { _ = "STUB: not implemented"; return *new(Colorizer) }

// WriteReport writes a human readable report to the provided writer
func (report *HumanReport) WriteReport(out io.Writer) error { _ = "STUB: not implemented"; return nil }

// Only show the document index if there is more than one document to show

// Show banner if enabled

// Loop over the diff and generate each report into the buffer

// Finish with one last newline so that we do not end next to the prompt

// generateHumanDiffOutput creates a human readable report of the provided diff and writes this into the given bytes buffer. There is an optional flag to indicate whether the document index (which documents of the input file) should be included in the report of the path of the difference.
func (report *HumanReport) generateHumanDiffOutput(output stringWriter, diff Diff, useGoPatchPaths bool, showPathRoot bool) error {
	_ = "STUB: not implemented"
	return nil
}

// For the use case in which only a path-less diff is suppose to be printed,
// omit the indent in this case since there is only one element to show

// generateHumanDetailOutput only serves as a dispatcher to call the correct sub function for the respective type of change
func (report *HumanReport) generateHumanDetailOutput(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (report *HumanReport) generateHumanDetailOutputAddition(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (report *HumanReport) generateHumanDetailOutputRemoval(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (report *HumanReport) generateHumanDetailOutputModification(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// delegate to special string output

func (report *HumanReport) generateHumanDetailOutputOrderchange(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (report *HumanReport) writeStringDiff(output stringWriter, from string, to string) {
	_ = "STUB: not implemented"
	return
}

// create line by line diff

// color and format each diff by type

// skip equal output if requested context is 0 or the equal text is empty

// add amount of unchanged lines as configured

// if string ends with \n we need to display one more line on the upper limit

func (report *HumanReport) highlightByLine(from, to string) string {
	_ = "STUB: not implemented"
	return ""
}

func humanReadableType(node *yamlv3.Node) string { _ = "STUB: not implemented"; return "" }

// use the YAML tag name without the exclamation marks

func (report *HumanReport) highlightRemovals(diffs []diffmatchpatch.Diff, indent int) string {
	_ = "STUB: not implemented"
	return ""
}

func (report *HumanReport) highlightAdditions(diffs []diffmatchpatch.Diff, indent int) string {
	_ = "STUB: not implemented"
	return ""
}

// LoadX509Certs tries to load the provided strings as a cert each and returns
// a textual representation of the certs, or an error if the strings are not
// X509 certs
func (report *HumanReport) LoadX509Certs(from, to string) (string, string, error) {
	_ = "STUB: not implemented"
	// Back out quickly if cert inspection is disabled
	return "", "", nil
}

// Create a YAML (hash with key/value) from a certificate to only display a few
// important fields (https://www.sslshopper.com/certificate-decoder.html):
//
//	Common Name: www.example.com
//	Organization: Company Name
//	Organization Unit: Org
//	Locality: Portland
//	State: Oregon
//	Country: US
//	Valid From: April 2, 2018
//	Valid To: April 2, 2019
//	Issuer: www.example.com, Company Name
//	Serial Number: 14581103526614300972 (0xca5a7c67490a792c)
func certificateSummaryAsYAML(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func yamlString(input interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func isMinorChange(from string, to string, minorChangeThreshold float64) bool {
	_ = "STUB: not implemented"
	return false
}

// Special case: Consider it a minor change if only two runes/characters were
// changed, which results in a default distance of four, two removals and two
// additions each.

func isMultiLine(from string, to string) bool { _ = "STUB: not implemented"; return false }

func (report *HumanReport) showWhitespaceCharacters(text string) string {
	_ = "STUB: not implemented"
	return ""
}

// createStringWithContinuousPrefix adds the defined prefix to each line of the
// objects string representation.
// The resulting string will always end with a newline.
func createStringWithContinuousPrefix(prefix string, obj interface{}, indent int) string {
	_ = "STUB: not implemented"
	return ""
}

// avoid add. additional empty newline if orig string ends with \n

// always adds a newline, even if orig string does not contain any

func createStringWithPrefix(prefix string, obj interface{}, indent int) string {
	_ = "STUB: not implemented"
	return ""
}

func plainTextLength(text string) int { _ = "STUB: not implemented"; return 0 }

func stringArrayLen(list []string) int { _ = "STUB: not implemented"; return 0 }

// writeTextBlocks writes strings into the provided buffer in either a table style (each string a column) or list style (each string a row)
func (report *HumanReport) writeTextBlocks(buf stringWriter, indent int, blocks ...string) {
	_ = "STUB: not implemented"
	return

	// Calcuclate the theoretical maximum line length if blocks would be rendered next to each other
}

// In case the line with blocks next to each other would surpass the terminal width, fall back to the no-table-style

// CreateTableStyleString takes the multi-line input strings as columns and arranges an output string to create a table-style output format with proper padding so that the text blocks can be arranged next to each other.
func CreateTableStyleString(separator string, indent int, columns ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func styledGoPatchPath(path *ytbx.Path) string { _ = "STUB: not implemented"; return "" }

func styledDotStylePath(path *ytbx.Path) string { _ = "STUB: not implemented"; return "" }
