package slogutil

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

const (
	LogFormatVar = "MACPREFS_LOG_FORMAT"
	LogLevelVar  = "MACPREFS_LOG_LEVEL"
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
	packageSlog = logger.With(logargs.GoPackageLogArg, "slogutil")
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

func LogLevel() (level slog.Level) {
	var llText string
	var err error
	levelText := os.Getenv(LogLevelVar)

	if levelText == "" {
		slog.SetLogLoggerLevel(slog.LevelInfo)
		newLogger(LogFormat(), nil).Info("No log level set with environment variable, defaulting to 'info'",
			"environment_variable", LogLevelVar,
		)
		level = slog.LevelInfo
		goto end
	}
	llText = levelText
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

func LogFormat() string {
	return os.Getenv(LogFormatVar)
}
