package logwrapper

import (
	"fmt"

	"go.uber.org/zap"
)

type mock struct {
}

// Mock returns a new mock logger
func Mock() LoggerWrapper {
	return &mock{}
}

func (m *mock) TraceID(v string) LoggerWrapper {
	return m
}

func (m *mock) Version(v string) LoggerWrapper {
	return m
}

func (m *mock) CreateSpan() LoggerWrapper {
	return m
}

func (m *mock) RemoveSpan() LoggerWrapper {
	return m
}

func (*mock) Info(msg string, fields ...zap.Field) {
	fmt.Printf("%s %s %v\n", "[ INFO ]", msg, fields)
}

func (*mock) Warn(msg string, fields ...zap.Field) {
	fmt.Printf("%s %s %v\n", "[ Warn ]", msg, fields)
}

func (*mock) Error(msg string, fields ...zap.Field) {
	fmt.Printf("%s %s %v\n", "[ Error ]", msg, fields)

}

func (*mock) Fatal(msg string, fields ...zap.Field) {
	fmt.Printf("%s %s %v\n", "[ Fatal ]", msg, fields)
}

func (*mock) Debug(msg string, fields ...zap.Field) {
	fmt.Printf("%s %s %v\n", "[ DEBUG ]", msg, fields)
}
