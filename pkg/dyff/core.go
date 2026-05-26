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
	"github.com/gonvenience/ytbx"

	yamlv3 "go.yaml.in/yaml/v3"
)

// CompareOption sets a specific compare setting for the object comparison
type CompareOption func(*compareSettings)

type compareSettings struct {
	NonStandardIdentifierGuessCountThreshold int
	IgnoreOrderChanges                       bool
	IgnoreWhitespaceChanges                  bool
	KubernetesEntityDetection                bool
	DetectRenames                            bool
	FormatStrings                            bool
	AdditionalIdentifiers                    []string
}

type compare struct {
	settings compareSettings
}

// AdditionalIdentifiers specifies additional identifiers that will be
// used as the key for matching maps from source to target.
func AdditionalIdentifiers(fieldNames ...string) CompareOption {
	_ = "STUB: not implemented"
	return *new(CompareOption)
}

// NonStandardIdentifierGuessCountThreshold specifies how many list entries are
// needed for the guess-the-identifier function to actually consider the key
// name. Or in short, if the lists only contain two entries each, there are more
// possibilities to find unique enough key, which might not qualify as such.
func NonStandardIdentifierGuessCountThreshold(nonStandardIdentifierGuessCountThreshold int) CompareOption {
	_ = "STUB: not implemented"
	return *new(CompareOption)
}

// IgnoreOrderChanges disables the detection for changes of the order in lists
func IgnoreOrderChanges(value bool) CompareOption {
	_ = "STUB: not implemented"
	return *new(CompareOption)
}

// IgnoreWhitespaceChanges disables the detection for whitespace only changes
func IgnoreWhitespaceChanges(value bool) CompareOption {
	_ = "STUB: not implemented"
	return *new(CompareOption)
}

// KubernetesEntityDetection enabled detecting entity identifiers from Kubernetes "kind:" and "metadata:" fields.
func KubernetesEntityDetection(value bool) CompareOption {
	_ = "STUB: not implemented"
	return *new(CompareOption)
}

// DetectRenames enabled detection of renames so that it correlates two entries based on their identifier field
func DetectRenames(value bool) CompareOption { _ = "STUB: not implemented"; return *new(CompareOption) }

func FormatStrings(value bool) CompareOption { _ = "STUB: not implemented"; return *new(CompareOption) }

// CompareInputFiles is one of the convenience main entry points for comparing
// objects. In this case the representation of an input file, which might
// contain multiple documents. It returns a report with the list of differences.
func CompareInputFiles(from ytbx.InputFile, to ytbx.InputFile, compareOptions ...CompareOption) (Report, error) {
	_ = "STUB: not implemented"
	// initialize the comparator with the tool defaults
	return *new(Report), nil
}

// apply the optional compare options provided to this function call

// in case Kubernetes mode is enabled, try to compare documents in the YAML
// file by their names rather than just by the order of the documents

// when the look-up of a name for each document in each file worked out, it
// means that the documents are most likely Kubernetes resources, so a comparison
// using the names can be done, otherwise, leave and continue with default behavior

// Reset the docs and names based on the collected details

// Compare the document nodes

func (compare *compare) objects(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (compare *compare) nonNilSameKindNodes(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore different ways to define a null value

func (compare *compare) documentNodes(from, to ytbx.InputFile) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// `from` and `to` contain the same `key` -> require comparison

// `from` contain the `key`, but `to` does not -> removal

// `to` contains a `key` that `from` does not have -> addition

// Push rename detection results

// Exclude from order change calculation

func (compare *compare) mappingNodes(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// `from` and `to` contain the same `key` -> require comparison

// `from` contain the `key`, but `to` does not -> removal

// `to` contains a `key` that `from` does not have -> addition

func (compare *compare) sequenceNodes(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	// Bail out quickly if there is nothing to check
	return nil, nil
}

// check if a known identifier (e.g. name, or id) can be used

// check if there is a field in all entries that could serve as an identifier

// check if Kubernetes resource fields can be used to identify items

// in any other case, compare lists as simple lists by relying on hashes

func (compare *compare) simpleLists(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Special case if both lists only contain one entry, then directly compare
// the two entries with each other

// Fill two lists with the hashes of the entries of each list

// `from` entry does not exist in `to` list

// `from` entry exists in `to` list, but there are duplicates and
// the number of duplicates is smaller

// `to` entry does not exist in `from` list

// `to` entry exists in `from` list, but there are duplicates and
// the number of duplicates is increased

func (compare *compare) namedEntryLists(path ytbx.Path, identifier listItemIdentifier, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fill two lists with the names of the entries that are common in both lists

// Find entries that are common to both lists to compare them separately, and
// find entries that are only in from, but not to and are therefore removed

// `from` and `to` have the same entry identified by identifier and name -> require comparison

// `from` has an entry (identified by identifier and name), but `to` does not -> removal

// Find entries that are only in to, but not from and are therefore added

// `to` and `from` have the same entry identified by identifier and name (comparison already covered by previous range)

// `to` has an entry (identified by identifier and name), but `from` does not -> addition

func (compare *compare) nodeValues(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// leave and don't report any differences if ignore whitespaces changes is
// configured and it is really only a whitespace only change between the strings

func (compare *compare) boolValues(path ytbx.Path, from *yamlv3.Node, to *yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this uses the various values mentioned in https://yaml.org/type/bool.html
var trueValues = [...]string{"y", "Y", "yes", "Yes", "YES", "true", "True", "TRUE", "on", "On", "ON"}
var falseValues = [...]string{"n", "N", "no", "No", "NO", "false", "False", "FALSE", "off", "Off", "OFF"}

func toBool(input string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (compare *compare) findOrderChangesInSimpleList(fromCommon, toCommon []*yamlv3.Node) []Detail {
	_ = "STUB: not implemented"
	// Try to find order changes ...
	return nil
}

// hasEntry returns whether the given node is in the provided list. Not exactly
// a fast or efficient way to verify that a node is already in a list, but
// given that this should rarely be used it is ok for now.
func (compare *compare) hasEntry(list []*yamlv3.Node, searchEntry *yamlv3.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// AsSequenceNode translates a string list into a SequenceNode
func AsSequenceNode(list ...string) *yamlv3.Node { _ = "STUB: not implemented"; return nil }

func findOrderChangesInNamedEntryLists(fromNames, toNames []string) []Detail {
	_ = "STUB: not implemented"
	return nil
}

// Try to find order changes ...

func packChangesAndAddToResult(list []Diff, path ytbx.Path, orderchanges []Detail, additions, removals []*yamlv3.Node) ([]Diff, error) {
	_ = "STUB: not implemented"
	// Prepare a diff for this path to added to the result set (if there are changes)
	return nil, nil
}

// If there were changes added to the details list, we can safely add it to
// the result set. Otherwise it the result set will be returned as-is.

func followAlias(node *yamlv3.Node) *yamlv3.Node { _ = "STUB: not implemented"; return nil }

func findValueByKey(mappingNode *yamlv3.Node, key string) (*yamlv3.Node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (compare *compare) listItemIdentifierCandidates() []string {
	_ = "STUB: not implemented"
	// Set default candidates that are most widly used
	return nil
}

// Add user supplied additional candidates (taking precedence over defaults)

// Add Kubernetes specific extra candidate

func (compare *compare) getIdentifierFromNamedLists(listA, listB *yamlv3.Node) (listItemIdentifier, error) {
	_ = "STUB: not implemented"
	return *new(listItemIdentifier), nil
}

// Check for the usual suspects: name, key, and id

func (compare *compare) getNonStandardIdentifierFromNamedLists(listA, listB *yamlv3.Node) listItemIdentifier {
	_ = "STUB: not implemented"
	return *new(listItemIdentifier)
}

// Sort the keys to ensure deterministic order

// getIdentifierFromKubernetesEntityList returns 'metadata.name' as a field identifier if the provided objects all have the key.
func (compare *compare) getIdentifierFromKubernetesEntityList(listA, listB *yamlv3.Node) (listItemIdentifier, error) {
	_ = "STUB: not implemented"
	return *new(listItemIdentifier), nil
}

// isEmptyDocument returns true in case the given YAML node is an empty document
func isEmptyDocument(node *yamlv3.Node) bool { _ = "STUB: not implemented"; return false }

// special case: content is just null (scalar)

func (compare *compare) basicType(node *yamlv3.Node) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (compare *compare) calcNodeHash(node *yamlv3.Node) (hash uint64) {
	_ = "STUB: not implemented"
	return 0
}

func sortNode(node *yamlv3.Node) { _ = "STUB: not implemented"; return }

func isList(node *yamlv3.Node) bool { _ = "STUB: not implemented"; return false }

// ChangeRoot changes the root of an input file to a position inside its
// document based on the given path. Input files with more than one document are
// not supported, since they could have multiple elements with that path.
func ChangeRoot(inputFile *ytbx.InputFile, path string, useGoPatchPaths bool, translateListToDocuments bool) error {
	_ = "STUB: not implemented"
	return nil
}

// For reference reasons, keep the original root level

// Find the object at the given path

// Change root of input file main document to a new list of documents based on the the list that was found

// Change root of input file main document to the object that was found

// Parse path string and create nicely formatted output path

func pathToString(path *ytbx.Path, useGoPatchPaths bool, showPathRoot bool) string {
	_ = "STUB: not implemented"
	return ""
}

func grab(node *yamlv3.Node, pathString string) (*yamlv3.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isWhitespaceOnlyChange(from string, to string) bool { _ = "STUB: not implemented"; return false }
