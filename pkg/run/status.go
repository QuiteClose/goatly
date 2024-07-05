package run

import (
	"io"
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
