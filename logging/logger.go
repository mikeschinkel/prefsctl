package logging

import (
	"log/slog"
	"os"
	"strings"
	"sync"
)

const (
	LogLevelVar  = "LOG_LEVEL"
	LogFormatVar = "LOG_FORMAT"
)

//goland:noinspection GoUnusedConst
const (
	DebugLevel = slog.LevelDebug
	InfoLevel  = slog.LevelInfo
	WarnLevel  = slog.LevelWarn
	ErrorLevel = slog.LevelError
)

type SlogLogger = slog.Logger
type Options = slog.HandlerOptions

var logger *slog.Logger
var packageSlog *slog.Logger

func Logger() *slog.Logger {
	if !initialized {
		Initialize()
		initialized = true
	}
	return logger
}

var initialized bool
var loggerMutex sync.Mutex

func Initialize() {
	loggerMutex.Lock()
	logger = newLogger(LogFormat(), &Options{
		Level: LogLevel(),
	})
	packageSlog = logger.With(GoPackageLogArg, "logging")
	slog.SetDefault(logger)
	loggerMutex.Unlock()
}

// Debug allows calling slog.Debug() without having to create the parameters
// Useful when creating the parameters will be expensive.
func Debug(log *slog.Logger, msg string, argsFunc func() []any) {
	if slog.Default().Enabled(nil, slog.LevelDebug) {
		log.Debug(msg, argsFunc()...)
	}
}

func newLogger(format string, options *slog.HandlerOptions) *slog.Logger {
	var handler slog.Handler
	switch strings.ToLower(format) {
	case "json":
		handler = slog.NewJSONHandler(LogWriter, options)
	default:
		handler = slog.NewTextHandler(LogWriter, options)
	}
	return slog.New(handler)
}

func LogFormat() string {
	return os.Getenv(LogFormatVar)
}

func LogLevel() (level slog.Level) {
	var err error

	llText := os.Getenv(LogLevelVar)
	if llText == "" {
		slog.SetLogLoggerLevel(slog.LevelInfo)
		newLogger(LogFormat(), nil).Info("No log level set with environment variable, defaulting to 'info'",
			"environment_variable", LogLevelVar,
		)
		level = slog.LevelInfo
		goto end
	}
	err = level.UnmarshalText([]byte(llText))
	if err != nil {
		slog.SetLogLoggerLevel(slog.LevelInfo)
		newLogger(LogFormat(), nil).Warn("Invalid log level, defaulting to 'info'",
			"invalid_level", llText,
		)
		level = slog.LevelInfo
	}
end:
	return level
}
