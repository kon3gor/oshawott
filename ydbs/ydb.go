package ydbs

import (
	"errors"

	"github.com/kon3gor/oshawott"
	"github.com/yandex-cloud/ydb-go-sdk/v2/connect"
)

type Cache interface {
	GetKey(v string) (oshawott.Key, bool)
	GetValue(k oshawott.Key) (string, bool)
	SavePair(v string, k oshawott.Key)
}

type YdbStorage struct {
	ctx   oshawott.AppContext
	cache Cache
	conn  *connect.Connection
}

func NewYdbStorage(ctx oshawott.AppContext, c Cache) YdbStorage {
	conn, _ := conn(ctx)
	return YdbStorage{ctx, c, conn}
}

func (ys YdbStorage) IsUrlSaved(url string) (oshawott.Key, bool) {
	key, found := ys.cache.GetKey(url)
	if found {
		return key, true
	}

	key, found = ys.getKey(url)
	if found {
		ys.cache.SavePair(url, key)
		return key, found
	}

	return oshawott.NoKey, false
}

func (ys YdbStorage) SaveUrl(key oshawott.Key, url string) error {
	ys.cache.SavePair(url, key)
	ys.saveUrl(key, url)
	return nil
}

func (ys YdbStorage) GetUrl(key oshawott.Key) (string, error) {
	cached, found := ys.cache.GetValue(key)
	if found {
		return cached, nil
	}

	url, found := ys.getValue(key)
	if found {
		return url, nil
	} else {
		//todo: fix this shit
		return "", errors.New("damn")
	}
}

// todo: fix cache here
func (ys YdbStorage) Init() (int, error) {
	used, err := ys.getUsedKeys()
	if err != nil {
		return 0, err
	}

	return len(used), nil
}
