package core

import (
	"log"
)

type logger interface {
	Info(v ...any)
	Warn(v ...any)
	Debug(v ...any)
}

type debugLogger struct{}

var dl debugLogger = debugLogger{}

func (l debugLogger) Info(v ...any) {
	log.Print("INFO: ", v)
}

func (l debugLogger) Warn(v ...any) {
	log.Print("WARN: ", v)
}

func (l debugLogger) Debug(v ...any) {
	log.Print("DEBUG: ", v)
}

type prodLogger struct{}

var pl prodLogger = prodLogger{}

func (l prodLogger) Info(v ...any) {
	log.Print("INFO: ", v)
}

func (l prodLogger) Warn(v ...any) {
	log.Print("WARN: ", v)
}

func (l prodLogger) Debug(v ...any) {
	//no-op
}

func selectLogger(ctx OshawottContext) logger {
	if ctx.Debug {
		return dl
	} else {
		return pl
	}
}
