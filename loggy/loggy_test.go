package loggy

import (
	"testing"
)

type logWriter struct {
	log []string
}

func (lw logWriter) Write(b []byte) (int, error) {
	s := string(b)
	lw.log = append(lw.log, s)
	return len(b), nil
}

func TestLoggy(t *testing.T) {
	t.Run("Set custom log writer", func(t *testing.T) { Set(logWriter{}) })
	t.Run("Info log info", func(t *testing.T) { Info("Testing 123") })
	t.Run("Warn log warning", func(t *testing.T) { Warn("Testing 456") })
	t.Run("Error log error", func(t *testing.T) { Error("Testing 789") })
}
