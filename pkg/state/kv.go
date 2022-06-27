package state

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
)

// daprStateStore points to a default statestore
const daprStateStore = "statestore"

// State is the implementation of persistence
type KV struct {
	dapr dapr.Client
}

// SaveUnique saves the short url on daprStore.
func (kv KV) SaveUnique(ctx context.Context, short, long string) error {
	return kv.dapr.SaveState(ctx, daprStateStore, short, []byte(long), nil)
}

// Retrieve retrieve the value from daprStore.
func (kv KV) Retrieve(ctx context.Context, short string) (string, error) {
	long, err := kv.dapr.GetState(ctx, daprStateStore, short, nil)

	if err != nil {
		return "", err
	}

	return string(long.Value), nil
}

// NewKV creates a new KV state using dapr.
func NewKV(dapr dapr.Client) KV {
	return KV{dapr}
}
