package kv

import (
	"context"
	"errors"
)

var (
	ErrNotFound          = errors.New("key not found")
	ErrPageTokenNotFound = errors.New("page token not found")
	ErrNoNamespace       = errors.New("no namespace")
)

type Pair struct {
	Key   string
	Value []byte
}

type Query struct {
	Prefix    string
	Pairs     []*Pair
	NextToken string
}

type KV interface {
	// Get returns the contents of a file from the storage backend.
	// Key must have a namespace prefix.
	// Example: /kernels/entries/123, where kernels is the namespace and the rest is the range key.
	Get(ctx context.Context, key string) (*Pair, error)
	Set(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	RangePrefix(ctx context.Context, prefix string, pageSize int32, pageToken string) (*Query, error)
}
