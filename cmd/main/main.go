package main

import (
	"context"

	"github.com/kon3gor/oshawott"
	"github.com/kon3gor/oshawott/file"
	"github.com/kon3gor/oshawott/internal/server"
	"github.com/kon3gor/oshawott/ydbs"
)

type MapCahce map[string]string

func (mc MapCahce) GetValue(k oshawott.Key) (string, bool) {
	v, f := mc[string(k)]
	return v, f
}

func (mc MapCahce) GetKey(v string) (oshawott.Key, bool) {
	k, f := mc[v]
	return oshawott.Key(k), f
}

func (mc MapCahce) SavePair(v string, k oshawott.Key) {
	sk := string(k)
	mc[v] = sk
	mc[sk] = v
}

func main() {
	ctx := oshawott.NewContext(context.Background(), oshawott.WithDebug(true))
	ydbs := ydbs.NewYdbStorage(ctx, make(MapCahce))
	count, _ := ydbs.Init()
	fkp := file.NewFileKeyProvider("keys.txt", ',', count)
	e := oshawott.NewEngine(fkp, ydbs)
	server.Start(ctx, e)
}
