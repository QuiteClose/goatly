package with

import (
	"os"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// TempDir creates a temporary directory and calls f with its path.
func TempDir(t *testing.T, f func(string)) {
	f(t.TempDir())
}

// TempWorkingDir is like TempDir but also changes to that directory.
func TempWorkingDir(t *testing.T, f func(string)) {
	origin, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(origin)
	TempDir(t, func(path string) {
		if err := os.Chdir(path); err != nil {
			t.Fatal(err)
		}
		f(path)
	})
}
