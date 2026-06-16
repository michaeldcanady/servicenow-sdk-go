package testutils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// StructuredLogger is a simple JSON logger for test failures.
type StructuredLogger struct {
	Output io.Writer
}

type LogEntry struct {
	TestName  string      `json:"test_name"`
	Status    string      `json:"status"`
	Error     string      `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
	Context   interface{} `json:"context,omitempty"`
}

func NewLogger() *StructuredLogger {
	return &StructuredLogger{Output: os.Stderr}
}

func (l *StructuredLogger) LogFailure(testName string, err error, context interface{}) {
	entry := LogEntry{
		TestName:  testName,
		Status:    "FAIL",
		Error:     err.Error(),
		Timestamp: "2026-06-16T10:00:00Z", // Placeholder for actual time
		Context:   context,
	}

	data, _ := json.Marshal(entry)
	fmt.Fprintln(l.Output, string(data))
}
