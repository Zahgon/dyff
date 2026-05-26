// Copyright © 2020 The Homeport Team
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

package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/homeport/dyff/pkg/dyff"
)

type reportConfig struct {
	Style          string `envDefault:"human"`
	UseIndentLines bool   `envDefault:"true"`

	IgnoreOrderChanges      bool `envDefault:"false"`
	IgnoreWhitespaceChanges bool `envDefault:"false"`
	IgnoreValueChanges      bool `envDefault:"false"`
	FormatStrings           bool `envDefault:"true"`
	DetectRenames           bool `envDefault:"true"`

	NoTableStyle          bool    `envDefault:"false"`
	DoNotInspectCerts     bool    `envDefault:"false"`
	UseGoPatchPaths       bool    `envDefault:"false"`
	MinorChangeThreshold  float64 `envDefault:"0.1"`
	MultilineContextLines int     `envDefault:"4"`

	KubernetesEntityDetection bool `envDefault:"true"`
	AdditionalIdentifiers     []string
	Filters                   []string
	Excludes                  []string
	FilterRegexps             []string
	ExcludeRegexps            []string

	ExitWithCode bool `envDefault:"false"`
	OmitHeader   bool `envDefault:"false"`
}

func initReportConfig() reportConfig { _ = "STUB: not implemented"; return *new(reportConfig) }

var reportOptions = initReportConfig()

func flagSet(name string, f ...func(*pflag.FlagSet)) *pflag.FlagSet {
	_ = "STUB: not implemented"
	return nil
}

func reportOptionsFlags() []*pflag.FlagSet { _ = "STUB: not implemented"; return nil }

// Deprecated

// OutputWriter encapsulates the required fields to define the look and feel of
// the output
type OutputWriter struct {
	PlainMode                  bool
	Restructure                bool
	OmitIndentHelper           bool
	EnforceDocumentStartMarker bool
	OutputStyle                string
}

func humanReadableFilename(filename string) string { _ = "STUB: not implemented"; return "" }

// WriteToStdout is a convenience function to write the content of the documents
// stored in the provided input file to the standard output
func (w *OutputWriter) WriteToStdout(filename string) error { _ = "STUB: not implemented"; return nil }

// WriteInplace writes the content of the documents stored in the provided input
// file to the file itself overwriting the content in place.
func (w *OutputWriter) WriteInplace(filename string) error { _ = "STUB: not implemented"; return nil }

// Force plain mode to make sure there are no ANSI sequences

// Write the buffered output to the provided input file (override in place)

func (w *OutputWriter) write(writer io.Writer, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeReport(cmd *cobra.Command, report dyff.Report) error {
	_ = "STUB: not implemented"
	return nil
}

// If configured, make sure `dyff` exists with an exit status
