package log

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

// LoggerOption is a set of options to apply to the logger
type LoggerOption func(*logrus.Entry) *logrus.Entry

const (
	packageField   = "package"
	requestIDField = "requestID"
	spandIDField   = "dd.span_id"
	traceIDField   = "dd.trace_id"
)

// SpanID adds a "dd.span_id" field to the logger
// Meant to be used with gopkg.in/DataDog/dd-trace-go.v1/ddtrace SpanID
func SpanID(spanID uint64) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(spandIDField, spanID)
	}
}

// PackageName adds a "package" field to the logger
func PackageName(packageName string) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(packageField, packageName)
	}
}

// RequestID adds a "requestID" field to the logger
func RequestID(requestID uint64) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(requestIDField, requestID)
	}
}

// TraceID adds a "dd.trace_id" field to the logger
// Meant to be used with gopkg.in/DataDog/dd-trace-go.v1/ddtrace TraceID
func TraceID(traceID uint64) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(traceIDField, traceID)
	}
}

// Formatter sets the logrus formatter for the logger.
// Only one formatter maybe set at a time.
func Formatter(formatter logrus.Formatter) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		entry.Logger.SetFormatter(formatter)
		return entry
	}
}

// Field adds a field with the given key and value to the logger
func Field(key string, value interface{}) LoggerOption {
	return func(entry *logrus.Entry) *logrus.Entry {
		return entry.WithField(key, value)
	}
}

// Logger is a flavored wrapper around
// the logrus package
type Logger struct {
	*logrus.Entry
}

// DeferOnFatal alters the exit behavior of fatal logs.
// When enabled, fatal logs will wait for the current goroutine
// to terminate, before exiting. This allows all deferred statements
// in the current goroutine to execute prior to exiting.
func (l *Logger) DeferOnFatal() {
	l.Logger.ExitFunc = func(exitCode int) {
		defer os.Exit(exitCode)
		runtime.Goexit()
	}
}

// NewLogger creates a new logger with the given options
// The default format is JSON
func NewLogger(options ...LoggerOption) *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	entry := logrus.NewEntry(logger)

	for _, option := range options {
		entry = option(entry)
	}

	return &Logger{entry}
}

// FromLogger creates a new logger using the passed in logger as a base. Applies
// the given options
func FromLogger(logger *Logger, options ...LoggerOption) *Logger {
	entry := logger.Entry

	for _, option := range options {
		entry = option(entry)
	}

	return &Logger{entry}
}
