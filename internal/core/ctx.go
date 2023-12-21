package core

import "context"

type OshawottContext struct {
	ctx    context.Context
	Port   int
	Debug  bool
	logger logger
}

type ContextOption func(*OshawottContext)

func WithPort(port int) ContextOption {
	return func(ctx *OshawottContext) {
		ctx.Port = port
	}
}

func WithDebug(debug bool) ContextOption {
	return func(ctx *OshawottContext) {
		ctx.Debug = debug
		ctx.logger = selectLogger(*ctx)
	}
}

func NewContext(ctx context.Context, options ...ContextOption) OshawottContext {
	appContext := OshawottContext{ctx, 8080, false, pl}
	for _, option := range options {
		option(&appContext)
	}
	return appContext
}
