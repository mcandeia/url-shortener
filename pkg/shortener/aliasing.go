package shortener

import (
	"context"
	"errors"

	"github.com/mcandeia/url-shortener/pkg/state"
)

type ContextAliasKey string

const AliasKey = ContextAliasKey("alias")

var ErrAliasIsMissing = errors.New("short url alias is missing")

// aliasing use alias to define shorten urls.
type aliasing struct {
	kv state.Store
}

// Short receives a long string and returns a short version of such string.
func (a aliasing) Short(ctx context.Context, long string) (short string, err error) {
	alias := ctx.Value(AliasKey)
	aliasStr, ok := alias.(string)

	if !ok {
		return "", ErrAliasIsMissing
	}

	return aliasStr, a.kv.SaveUnique(ctx, aliasStr, long)
}

// Long receives the short string version and returns its long version.
func (a aliasing) Long(ctx context.Context, short string) (long string, err error) {
	return a.kv.Retrieve(ctx, short)
}

func NewAliasing(kv state.Store) Engine {
	return aliasing{kv}
}
