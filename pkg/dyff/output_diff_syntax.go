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
	"io"
)

// DiffSyntaxReport is a reporter with human readable output in mind
type DiffSyntaxReport struct {
	PathPrefix            string
	RootDescriptionPrefix string
	ChangeTypePrefix      string
	HumanReport
}

// WriteReport writes a human readable report to the provided writer
func (report *DiffSyntaxReport) WriteReport(out io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Only show the document index if there is more than one document to show

// Loop over the diff and generate each report into the buffer

// Finish with one last newline so that we do not end next to the prompt

// generatedyffSyntaxDiffOutput creates a human readable report of the provided diff and writes this into the given bytes buffer. There is an optional flag to indicate whether the document index (which documents of the input file) should be included in the report of the path of the difference.
func (report *DiffSyntaxReport) generateDiffSyntaxDiffOutput(output stringWriter, diff Diff, useGoPatchPaths bool, showPathRoot bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Only @@ also needs a postfix

// Write the root description onto its own line

// For the use case in which only a path-less diff is suppose to be printed,
// omit the indent in this case since there is only one element to show

// generatedyffSyntaxDetailOutput only serves as a dispatcher to call the correct sub function for the respective type of change
func (report *DiffSyntaxReport) generateDiffSyntaxDetailOutput(detail Detail) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (report *DiffSyntaxReport) prefixChangeType(detailOutput string) string {
	_ = "STUB: not implemented"
	return ""
}

func (report *DiffSyntaxReport) prefixChangeBlock(detailOutput string, blockPrefix rune) string {
	_ = "STUB: not implemented"
	// trim newline from the end
	return ""
}

// the first line is the change type, we don't want to prefix that

// Remove the ADDITION rune from the first line
