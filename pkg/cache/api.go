package cache

import "github.com/kon3gor/oshawott/pkg/keys"

type CacheProvider interface {
	GetUsedKeys() []keys.Key
}
