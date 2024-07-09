package run

import (
	"errors"
	"os/exec"
	"testing"
	"io"
)

type Runtime struct {
	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
}


// Any Process with a non-zero exit code should return an exec.ExitError
func TestProcessOutputNonZeroExit(t *testing.T) {
	f := func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error) {
		return 1, nil
	}
	p := EntryPoint(f)
	_, err := p.Output()
	if e, ok := err.(*exec.ExitError); ok {
		if e.ExitCode() != 1 {
			t.Errorf("Exit code should be 1, but got %d", e.ExitCode())
		}
	} else {
		t.Fatalf("ExitError should be returned, but got %T", e)
	}
}

func TestProcessOutputZeroExit(t *testing.T) {
	f := func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error) {
		return 0, nil
	}
	p := EntryPoint(f)
	output, err := p.Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			t.Errorf("nil should be returned, but got %T", e)
		case *exec.ExitError:
			t.Errorf("No ExitError expected but got %d", e.ExitCode())
		default:
			t.Errorf("nil, exec.Error or exec.ExitError should be returned, but got %T", e)
		}
	}
}

func TestProcessOutputError(t *testing.T) {
	f := func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error) {
		return 0, errors.New("test error")
	}
	p := EntryPoint(f)
	output, err := p.Output()
	switch e := err.(type) {
	case *exec.Error:
		t.Errorf("nil should be returned, but got %T", e)
	case *exec.ExitError:
		t.Errorf("No ExitError expected but got %d", e.ExitCode())
	default:
		t.Errorf("nil should be returned, but got %T", e)
	}
}

func TestFunctionWait(t *testing.T) {
	f := func(stdin, stdout, stderr io.ReadWriter, args ...string) (int, error) {
		return 0, nil
	}
	p := EntryPoint(f)
	p.Start()
	p.Wait()
	if p.Status == nil {
		t.Fatalf("Status must be set after Start is called.")
	}
	if process.Status.ExitCode != nil {
		t.Fatalf("Exit code must not be set until the routine completes.")
	}
	if process.Status.ExitCode != 0 {
		t.Fatalf("Exit code must be 0.")
	}
}
