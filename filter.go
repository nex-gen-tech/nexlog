package nexlog

// LogFilter is an interface for filtering logs.
// It exposes the Filter method that returns true if a log at the provided level with the provided arguments should be logged.
type LogFilter interface {
	Filter(logLevel Level, args ...any) bool
}

// defaultNoFilter is a type of LogFilter that does not filter any logs.
type defaultNoFilter struct{}

// NewDefaultNoFilter creates and returns a new instance of defaultNoFilter.
// This function is useful when no log filtering is required.
func NewDefaultNoFilter() LogFilter {
	return &defaultNoFilter{}
}

// Filter method for defaultNoFilter always returns true, indicating that no logs are filtered out.
func (dnf *defaultNoFilter) Filter(logLevel Level, args ...any) bool {
	return true
}
