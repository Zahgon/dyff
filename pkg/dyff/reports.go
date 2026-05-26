package dyff

import (
	"github.com/gonvenience/ytbx"
)

func (r Report) filter(hasPath func(*ytbx.Path) bool) (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}

// Filter accepts YAML paths as input and returns a new report with differences for those paths only
func (r Report) Filter(paths ...string) (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}

// Exclude accepts YAML paths as input and returns a new report with differences without those paths
func (r Report) Exclude(paths ...string) (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}

// FilterRegexp accepts regular expressions as input and returns a new report with differences for matching those patterns
func (r Report) FilterRegexp(pattern ...string) (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}

// ExcludeRegexp accepts regular expressions as input and returns a new report with differences for not matching those patterns
func (r Report) ExcludeRegexp(pattern ...string) (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}

func (r Report) IgnoreValueChanges() (result Report) {
	_ = "STUB: not implemented"
	return *new(Report)
}
