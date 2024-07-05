package unless

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/quiteclose/goatly/pkg/message"
	"github.com/quiteclose/goatly/pkg/run"
)

///////////////////////////////////////////////////////////////////////////////

// RunError will call the callback unless the run.Status has an error.
func RunError(s *run.Status, callback func(string)) bool {
	conditionMet := s.Error != nil
	if conditionMet {
		callback("run.Status.Error != nil")
	}
	return conditionMet
}

// RunNoError will call the callback unless the run.Status has an error.
func RunNoError(s *run.Status, callback func(string)) bool {
	conditionMet := s.Error == nil
	if conditionMet {
		callback("run.Status.Error == nil")
	}
	return conditionMet
}

// RunExitCode will call the callback unless the run.Status has the given exit code.
func RunExitCode(s *run.Status, exitCode int, callback func(string)) bool {
	conditionMet := *s.ExitCode != exitCode
	if conditionMet {
		callback(fmt.Sprintf("run.Status.ExitCode != %d", exitCode))
	}
	return conditionMet
}

// RunNonZeroExitCode will call the callback unless the run.Status has a non-zero exit code.
func RunNonZeroExitCode(s *run.Status, callback func(string)) bool {
	conditionMet := *s.ExitCode == 0
	if conditionMet {
		callback("run.Status.ExitCode should be non-zero")
	}
	return conditionMet
}

// RunStderr will call the callback unless the run.Status.Stderr contains the given string.
func RunStderr(s *run.Status, expected string, callback func(string)) bool {
	conditionMet := !strings.Contains(s.Stderr, expected)
	if conditionMet {
		callback(message.MustContain("run.Status.Stderr", s.Stderr, expected))
	}
	return conditionMet
}

// RunStdout will call the callback unless the run.Status.Stdout contains the given string.
func RunStdout(s *run.Status, expected string, callback func(string)) bool {
	conditionMet := !strings.Contains(s.Stdout, expected)
	if conditionMet {
		callback(message.MustContain("run.Status.Stdout", s.Stdout, expected))
	}
	return conditionMet
}

// RunStdoutMatch will call the callback unless the run.Status.Stdout matches the given regular expression.
func RunStdoutMatch(s *run.Status, pattern regexp.Regexp, callback func(string)) bool {
	conditionMet := !pattern.MatchString(s.Stdout)
	if conditionMet {
		callback(message.MustMatch("run.Status.Stdout", pattern, s.Stdout))
	}
	return conditionMet
}

// RunStderrMatch will call the callback unless the run.Status.Stderr matches the given regular expression.
func RunStderrMatch(s *run.Status, pattern regexp.Regexp, callback func(string)) bool {
	conditionMet := pattern.MatchString(s.Stderr)
	if conditionMet {
		callback(message.MustMatch("run.Status.Stderr", pattern, s.Stderr))
	}
	return conditionMet
}

// RunStderr will call the callback unless the run.Status.Stderr contains the given string.
func RunNotStderr(s *run.Status, expected string, callback func(string)) bool {
	conditionMet := strings.Contains(s.Stderr, expected)
	if conditionMet {
		callback(message.MustNotContain("run.Status.Stderr", s.Stderr, expected))
	}
	return conditionMet
}

// RunNotStderrMatch will call the callback if the run.Status.Stderr matches the given regular expression.
func RunNotStderrMatch(s *run.Status, pattern regexp.Regexp, callback func(string)) bool {
	conditionMet := pattern.MatchString(s.Stderr)
	if conditionMet {
		callback(message.MustNotMatch("run.Status.Stderr", pattern, s.Stderr))
	}
	return conditionMet
}

// RunStdout will call the callback unless the run.Status.Stdout contains the given string.
func RunNotStdout(s *run.Status, expected string, callback func(string)) bool {
	conditionMet := strings.Contains(s.Stdout, expected)
	if conditionMet {
		callback(message.MustNotContain("run.Status.Stdout", s.Stdout, expected))
	}
	return conditionMet
}

// RunNotStdoutMatch will call the callback if the run.Status.Stdout matches the given regular expression.
func RunNotStdoutMatch(s *run.Status, pattern regexp.Regexp, callback func(string)) bool {
	conditionMet := pattern.MatchString(s.Stdout)
	if conditionMet {
		callback(message.MustNotMatch("run.Status.Stdout", pattern, s.Stdout))
	}
	return conditionMet
}
