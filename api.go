package oshawott

import (
	"context"
	"log"
)

type Key string

const NoKey Key = ""

type KeyProvider interface {
	GetKey(url string) (Key, error)
}

type Storage interface {
	GetUrl(key Key) (string, error)
	SaveUrl(key Key, url string) error
	IsUrlSaved(url string) (Key, bool)
}

type AppContext struct {
	Ctx    context.Context
	Port   int
	Debug  bool
	logger *log.Logger
}

type contextOption func(*AppContext)

func WithPort(port int) contextOption {
	return func(ctx *AppContext) {
		ctx.Port = port
	}
}

func WithLogger(logger *log.Logger) contextOption {
	return func(ctx *AppContext) {
		ctx.logger = logger
	}
}

func WithDebug(debug bool) contextOption {
	return func(ctx *AppContext) {
		ctx.Debug = debug
	}
}

func NewContext(ctx context.Context, options ...contextOption) AppContext {
	appContext := AppContext{ctx, 8080, false, log.Default()}
	for _, option := range options {
		option(&appContext)
	}
	return appContext
}
