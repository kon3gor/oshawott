package main

import (
	"context"

	"github.com/kon3gor/oshawott/internal/core"
	"github.com/kon3gor/oshawott/pkg/engine"
	"github.com/kon3gor/oshawott/pkg/keys"
)

func main() {
	ctx := core.NewContext(context.Background(), core.WithDebug(true))
	e := engine.NewEngine(ctx, keys.NewFileKeyProvider("keys.txt", ","))
	e.Start()
}
