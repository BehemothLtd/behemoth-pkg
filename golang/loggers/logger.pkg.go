package loggers

import (
	"io"
	"strings"

	"github.com/BehemothLtd/behemoth-pkg/golang/constants"
	"github.com/BehemothLtd/behemoth-pkg/golang/utils"

	"github.com/rs/zerolog"
)

type InfoDebugWriter struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (w *InfoDebugWriter) Write(p []byte) (n int, err error) {
	return w.Stderr.Write(p)
}

func (w *InfoDebugWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level == zerolog.DebugLevel || level == zerolog.InfoLevel {
		return w.Stdout.Write(p)
	}
	return w.Stderr.Write(p)
}

func GetLogLevel() zerolog.Level {
	switch strings.ToLower(utils.GetEnv("LOGGER_LEVEL", "info")) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()

	if requestID, ok := ctx.Value(constants.CtxRequestID).(string); ok {
		e.Str("requestId", requestID)
	}

	if remoteIp, ok := ctx.Value(constants.CtxRemoteIP).(string); ok {
		e.Str("remoteIp", remoteIp)
	}
}
