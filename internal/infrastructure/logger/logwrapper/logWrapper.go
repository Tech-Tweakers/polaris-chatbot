package logwrapper

import (
	"go.uber.org/zap"
)

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

// LoggerWrapper is a wrapper for zap.Logger + Span handling
type LoggerWrapper interface {
	TraceID(v string) LoggerWrapper
	Version(v string) LoggerWrapper
	CreateSpan() LoggerWrapper
	RemoveSpan() LoggerWrapper
	Logger // interface extends logger
}

type logWrapper struct {
	logger  Logger
	traceID string
	span    *span
	version string
}

// New returns a new logger
func New(logger Logger) LoggerWrapper {
	return &logWrapper{
		logger: logger,
	}
}

func (l *logWrapper) TraceID(v string) LoggerWrapper {
	l.traceID = v
	return l.clone()
}

func (l *logWrapper) Version(v string) LoggerWrapper {
	l.version = v
	return l.clone()
}

func (l *logWrapper) CreateSpan() LoggerWrapper {
	l.span = createSpan(l.span)
	return l
}

func (l *logWrapper) RemoveSpan() LoggerWrapper {
	if l.span != nil {
		l.span = l.span.parent
	}
	return l
}

func (l *logWrapper) Info(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Info(msg, f...)
}

func (l *logWrapper) Warn(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Warn(msg, f...)
}

func (l *logWrapper) Error(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Error(msg, f...)
}

func (l *logWrapper) Fatal(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Fatal(msg, f...)
}

func (l *logWrapper) Debug(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Debug(msg, f...)
}

func (l *logWrapper) clone() LoggerWrapper {
	return &logWrapper{
		logger:  l.logger,
		traceID: l.traceID,
		span:    l.span,
		version: l.version,
	}
}

func (l *logWrapper) mergeField(fields ...zap.Field) []zap.Field {

	parentID := ""
	spanID := ""
	if l.span != nil {
		spanID = l.span.id
		if l.span.parent != nil {
			parentID = l.span.parent.id
		}
	}
	s := []zap.Field{
		zap.String("version", l.version),
		zap.String("trace_id", l.traceID),
		zap.String("span_parent_id", parentID),
		zap.String("span_id", spanID),
	}

	s = append(s, fields...)

	return s
}
