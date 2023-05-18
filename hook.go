package nexlog

// LogHook interface
type LogHook interface {
	AddHook(logLevel Level, args ...any) error
}

type DefaultSlackHook struct{}

// NewDefaultSlackHook
func NewDefaultSlackHook() LogHook {
	return &DefaultSlackHook{}
}

// AddHook
func (dsh *DefaultSlackHook) AddHook(logLevel Level, args ...any) error {
	return nil
}
