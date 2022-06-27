package shortener

import (
	"context"
	b64 "encoding/base64"
)

// base64 is a shortener engine that receives a string and returns its base64 encoded version.
type base64 struct {
}

// Short encode the from string to base62 version.
func (b base64) Short(_ context.Context, long string) (string, error) {
	return b64.StdEncoding.EncodeToString([]byte(long)), nil
}

// Long decode the from string to base62 version.
func (b base64) Long(_ context.Context, short string) (string, error) {
	bts, err := b64.StdEncoding.DecodeString(short)
	if err != nil {
		return "", err
	}

	return string(bts), nil
}

// NewBase64 returns a Base62 Shortener engine.
func NewBase64() Engine {
	return base64{}
}
