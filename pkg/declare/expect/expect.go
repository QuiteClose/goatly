// Re-implementation of the internal/unless package generated by declare.py
// Each function calling t.Errorf with an ExpectFailed message if the
// condition is not met.
package expect

import (
	"github.com/quiteclose/goatly/internal/unless"
	"github.com/quiteclose/goatly/pkg/run"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// Contains will call t.Errorf unless a contains b
func Contains(t *testing.T, a, b string, message string) bool {
	return unless.Contains(a, b, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// DirExists will call t.Errorf unless the directory exists
func DirExists(t *testing.T, path string, message string) bool {
	return unless.DirExists(path, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// Equal will call t.Errorf unless a == b
func Equal(t *testing.T, a, b interface{}, message string) bool {
	return unless.Equal(a, b, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// False will call t.Errorf unless a == false
func False(t *testing.T, a bool, message string) bool {
	return unless.False(a, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// Nil will call t.Errorf unless a == nil
func Nil(t *testing.T, a interface{}, message string) bool {
	return unless.Nil(a, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// NotEqual will call t.Errorf unless a != b
func NotEqual(t *testing.T, a, b interface{}, message string) bool {
	return unless.NotEqual(a, b, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// NotNil will call t.Errorf unless a != nil
func NotNil(t *testing.T, a interface{}, message string) bool {
	return unless.NotNil(a, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// NotPathExists will call t.Errorf if the path exists
func NotPathExists(t *testing.T, path string, message string) bool {
	return unless.NotPathExists(path, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunError will call t.Errorf unless the run.Status has an error.
func RunError(t *testing.T, s *run.Status, message string) bool {
	return unless.RunError(s, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunExitCode will call t.Errorf unless the run.Status has the given exit code.
func RunExitCode(t *testing.T, s *run.Status, exitCode int, message string) bool {
	return unless.RunExitCode(s, exitCode, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunNoError will call t.Errorf unless the run.Status has an error.
func RunNoError(t *testing.T, s *run.Status, message string) bool {
	return unless.RunNoError(s, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunNonZeroExitCode will call t.Errorf unless the run.Status has a non-zero exit code.
func RunNonZeroExitCode(t *testing.T, s *run.Status, message string) bool {
	return unless.RunNonZeroExitCode(s, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStderr will call t.Errorf unless the run.Status.Stderr contains the given string.
func RunNotStderr(t *testing.T, s *run.Status, expected string, message string) bool {
	return unless.RunNotStderr(s, expected, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunNotStderrMatch will call t.Errorf if the run.Status.Stderr matches the given regular expression.
func RunNotStderrMatch(t *testing.T, s *run.Status, pattern regexp.Regexp, message string) bool {
	return unless.RunNotStderrMatch(s, pattern, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStdout will call t.Errorf unless the run.Status.Stdout contains the given string.
func RunNotStdout(t *testing.T, s *run.Status, expected string, message string) bool {
	return unless.RunNotStdout(s, expected, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunNotStdoutMatch will call t.Errorf if the run.Status.Stdout matches the given regular expression.
func RunNotStdoutMatch(t *testing.T, s *run.Status, pattern regexp.Regexp, message string) bool {
	return unless.RunNotStdoutMatch(s, pattern, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStderr will call t.Errorf unless the run.Status.Stderr contains the given string.
func RunStderr(t *testing.T, s *run.Status, expected string, message string) bool {
	return unless.RunStderr(s, expected, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStderrMatch will call t.Errorf unless the run.Status.Stderr matches the given regular expression.
func RunStderrMatch(t *testing.T, s *run.Status, pattern regexp.Regexp, message string) bool {
	return unless.RunStderrMatch(s, pattern, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStdout will call t.Errorf unless the run.Status.Stdout contains the given string.
func RunStdout(t *testing.T, s *run.Status, expected string, message string) bool {
	return unless.RunStdout(s, expected, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// RunStdoutMatch will call t.Errorf unless the run.Status.Stdout matches the given regular expression.
func RunStdoutMatch(t *testing.T, s *run.Status, pattern regexp.Regexp, message string) bool {
	return unless.RunStdoutMatch(s, pattern, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}

// True will call t.Errorf unless a == true
func True(t *testing.T, a bool, message string) bool {
	return unless.True(a, func(s string) {
		t.Errorf("ExpectFailed: %s\n%s", message, s)
	})
}
