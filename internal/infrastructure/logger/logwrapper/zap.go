package logwrapper

import (
	"go.uber.org/zap"
)

// Zap is a wrapper for zap.Logger
type Zap struct {
	Logger zap.Logger
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (z *Zap) Info(msg string, fields ...zap.Field) {
	z.Logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (z *Zap) Warn(msg string, fields ...zap.Field) {
	z.Logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (z *Zap) Error(msg string, fields ...zap.Field) {
	z.Logger.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (z *Zap) Fatal(msg string, fields ...zap.Field) {
	z.Logger.Fatal(msg, fields...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (z *Zap) Debug(msg string, fields ...zap.Field) {
	z.Logger.Debug(msg, fields...)
}
