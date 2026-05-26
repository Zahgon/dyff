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

package cmd

import (
	"os"
	"path/filepath"

	"github.com/gonvenience/bunt"
	"github.com/gonvenience/term"
	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
)

var name = func() string {
	ep, err := os.Executable()
	if err != nil {
		return "dyff"
	}

	return filepath.Base(ep)
}()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           name,
	SilenceErrors: true,
	SilenceUsage:  true,
	Long: `
δyƒƒ /ˈdʏf/ - a diff tool for YAML files, and sometimes JSON. Also, It
can transform YAML to JSON, and vice versa. The order of keys in hashes
is preserved during the conversion.
`,
}

// NewRootCmd returns the root command (for generating documentation)
func NewRootCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// ResetSettings resets command settings to default. This is only required by
	// the test suite to make sure that the flag parsing works correctly.
	return nil
}

func ResetSettings() { _ = "STUB: not implemented"; return }

// rearrange will rearrange the OS args to match `dyff between --flags from to`
// to mitigate an issue in `kubectl`, which puts the `from` and `to` at the
// second and third position in the command arguments.
func rearrange() []string { _ = "STUB: not implemented"; return nil }

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	_ = "STUB: not implemented"
	// In case `KUBECTL_EXTERNAL_DIFF` is set with `dyff`, it is very likely
	// that `kubectl` intends to use `dyff` for its `diff` command. Therefore,
	// enable Kubernetes specific entity detection and fix the order issue.
	return nil
}

// Make sure the OS args are in a supported order

// Enable Kubernetes specific entity detection implicitly

// Add implicit exclude for metadata.managedFields as this cannot
// be configured via a command-line flag using KUBECTL_EXTERNAL_DIFF
// due to an bug/feature in kubectl that ignore command-line flags
// in the diff environment variable with non alphanumeric characters

// Special case ExitCode, which means that we will exit immediately
// with the given exit code

// In any other case, create a default ExitCode with `error` value

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	rootCmd.PersistentFlags().VarP(&bunt.ColorSetting, "color", "c", "specify color usage: on, off, or auto")
	rootCmd.PersistentFlags().VarP(&bunt.TrueColorSetting, "truecolor", "t", "specify true color usage: on, off, or auto")
	rootCmd.PersistentFlags().IntVarP(&term.FixedTerminalWidth, "fixed-width", "w", -1, "disable terminal width detection and use provided fixed value")
	rootCmd.PersistentFlags().BoolVarP(&ytbx.PreserveKeyOrderInJSON, "preserve-key-order-in-json", "k", false, "use ordered keys during JSON decoding (non standard behavior)")
}
