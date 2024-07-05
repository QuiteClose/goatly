package run

import (
	"io"
	"testing"

	"github.com/quiteclose/goatly/pkg/declare/expect"
)

///////////////////////////////////////////////////////////////////////////////

type IoStreams struct {
	Stdin, Stdout, Stderr io.ReadWriter
}

type Status struct {
	Args     []string
	Error    *error
	ExitCode *int
	Stdin    string
	Stdout   string
	Stderr   string
	Streams  *IoStreams
}

// ExpectError expects the Status to have an error.
func (s *Status) ExpectError(t *testing.T) bool {
	return expect.NotNil(t, s.Error, "run.Status expected to have an error")
}

// ExpectNoError expects the Status to not have an error.
func (s *Status) ExpectNoError(t *testing.T) bool {
	return expect.Nil(t, s.Error, "run.Status expected to not have an error")
}

// ExpectExitCode expects the Status to have the given exit code.
func (s *Status) ExpectExitCode(t *testing.T, exitCode int) bool {
	return expect.Equal(t, exitCode, *s.ExitCode, "run.Status expected specific exit code")
}

// ExpectNonZeroExitCode expects the Status to have a non-zero exit code.
func (s *Status) ExpectNonZeroExitCode(t *testing.T) bool {
	return expect.NotEqual(t, 0, *s.ExitCode, "run.Status expected non-zero exit code")
}
