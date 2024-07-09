package run

import (
	"bytes"
	"io"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////

type ProcessStatus struct {
	StartTime *int64
	EndTime *int64
	ExitCode *int
	Error error
}

type Process struct {
	EntryPoint *func (stdin, stdout, stderr io.ReadWriter, args ...string) (int, error)
	Command *string
	Args []string
	Env map[string]string
	WorkingDir string
	Status *ProcessStatus
}


func (p *Process) Start() {
	startTime := time.Now().UnixNano()
	p.Status = &ProcessStatus{
		StartTime: &startTime,
	}
}

///////////////////////////////////////////////////////////////////////////////

func EntryPoint(callable func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error)) *Process {
	return &Process{
		EntryPoint: &callable,
	}
}

// EntryPoint creates the required Buffers and calls the given function.
// If the function returns an error, it will be reported as a test error.
func OldEntryPoint(t *testing.T, callable func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error), args ...string) *Status {
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
