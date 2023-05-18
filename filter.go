package nexlog

// LogFilter interface
type LogFilter interface {
	Filter(logLevel Level, args ...any) bool
}

type defaultNoFilter struct{}

// NewDefaultNoFilter
func NewDefaultNoFilter() LogFilter {
	return &defaultNoFilter{}
}

// Filter
func (dnf *defaultNoFilter) Filter(logLevel Level, args ...any) bool {
	return true
}
