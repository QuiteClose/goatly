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

func (status *Status) Write(data string) {
	if status.Streams != nil && status.Streams.Stdin != nil {
		status.Streams.Stdin.Write([]byte(data))
	}
}

func (status *Status) ReadLine() string {
	if status.Streams != nil && status.Streams.Stdout != nil {
		buf := make([]byte, 1)
		line := ""
		for {
			n, err := status.Streams.Stdout.Read(buf)
			if n == 0 || err != nil {
				break
			}
			if buf[0] == '\n' {
				break
			}
			line += string(buf)
		}
		status.Stdout += line
		return line
	}
	return ""
}

func (status *Status) ReadErrLine() string {
	if status.Streams != nil && status.Streams.Stderr != nil {
		buf := make([]byte, 1)
		line := ""
		for {
			n, err := status.Streams.Stderr.Read(buf)
			if n == 0 || err != nil {
				break
			}
			if buf[0] == '\n' {
				break
			}
			line += string(buf)
		}
		return line
	}
	return ""
}
