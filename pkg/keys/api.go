package keys

import "github.com/kon3gor/oshawott/internal/core"

type Key string

const NoKey Key = ""

type KeyProvider interface {
	// Provides a unique key for the given url
	GetKey(ctx core.OshawottContext, url string) (Key, error)
}
