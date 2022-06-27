package shortener

import "context"

// noop is a identity shortener that returns whatever it receives as parameter.
type noop struct{}

// Short receives a from string and returns the same.
func (n noop) Short(_ context.Context, long string) (string, error) {
	return long, nil
}

// Long receives a short string and returns the same.
func (n noop) Long(_ context.Context, short string) (string, error) {
	return short, nil
}

// NewNoop creates a noop shortener engine.
func NewNoop() Engine {
	return noop{}
}
