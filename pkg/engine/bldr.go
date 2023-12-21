package engine

import (
	"github.com/kon3gor/oshawott/internal/core"
	"github.com/kon3gor/oshawott/pkg/keys"
)

type engineOption func(*Engine)

func NewEngine(ctx core.OshawottContext, keyProvider keys.KeyProvider, opts ...engineOption) Engine {
	e := Engine{ctx, keyProvider}
	for _, opt := range opts {
		opt(&e)
	}

	return e
}
