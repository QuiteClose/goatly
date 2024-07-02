package run

import (
	"bytes"
	"io"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// EntryPoint creates the required Buffers and calls the given function.
// If the function returns an error, it will be reported as a test error.
func EntryPoint(t *testing.T, callable func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error), args ...string) *Status {
	var stdin, stdout, stderr bytes.Buffer
	code, err := callable(&stdin, &stdout, &stderr, args...)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	return &Status{
		Args:     args,
		Error:    &err,
		ExitCode: &code,
		Stdin:    stdin.String(),
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Streams: &IoStreams{
			Stdin:  &stdin,
			Stdout: &stdout,
			Stderr: &stderr,
		},
	}
}
