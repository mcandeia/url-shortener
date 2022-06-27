package shortener

import "context"

// Engine is the shortener engine.
type Engine interface {
	// Short receives a long string and returns a short version of such string.
	Short(ctx context.Context, long string) (short string, err error)
	// Long receives the short string version and returns its long version.
	Long(ctx context.Context, short string) (long string, err error)
}

// EngineID is the engine ID.
type EngineID int

const (
	// Base64 represents the engine ID of base64.
	Base64 EngineID = iota
	// Noop represents noop engine ID.
	Noop
	// Allow users to specify which shortened url should be used.
	Aliasing
)
