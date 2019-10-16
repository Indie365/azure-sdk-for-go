// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

// LogClassification is used to group entries.  Each group can be toggled on or off.
type LogClassification string

const (
	// LogError
	LogError        LogClassification = "Error"
	LogRequest      LogClassification = "Request"
	LogResponse     LogClassification = "Response"
	LogSlowResponse LogClassification = "SlowResponse"
)

// Listener is the function signature invoked when writing log entries.
// A Listener is required to perform its own synchronization if it's
// expected to be called from multiple Go routines.
type Listener func(LogClassification, string)

// Logger controls which classifications to log and writing to the underlying log.
type Logger struct {
	cls []LogClassification
	lst Listener
}

// SetClassifications is used to control which classifications are written to
// the log.  By default all log classifications are written.
func (l *Logger) SetClassifications(cls ...LogClassification) {
	l.cls = cls
}

// SetListener will set the Logger to write to the specified Listener.
func (l *Logger) SetListener(lst Listener) {
	l.lst = lst
}

// Should returns true if the specified log classification should be written to the log.
// TODO: explain why you would want to call this
func (l *Logger) Should(cls LogClassification) bool {
	if l.cls == nil || len(l.cls) == 0 {
		return true
	}
	for _, c := range l.cls {
		if c == cls {
			return true
		}
	}
	return false
}

// Write invokes the underlying Listener with the specified classification and message.
// If the classification shouldn't be logged or there is no listener then Write does nothing.
func (l *Logger) Write(cls LogClassification, message string) {
	if l.lst == nil || !l.Should(cls) {
		return
	}
	l.lst(cls, message)
}

var log Logger

// Log returns the process-wide logger.
func Log() *Logger {
	return &log
}
