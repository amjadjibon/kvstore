package iface

import (
	"context"
	"time"
)

type KVStore interface {
	// Get the key from store
	Get(ctx context.Context, key string, value interface{}) (bool, error)

	// Set the key value data to the store with ttl
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error

	// Delete the key
	Delete(ctx context.Context, key string) error
}