package with

import (
	"os"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// TempWorkingDir is like TempDir but also changes to that directory.
func TempWorkingDir(t *testing.T, f func(string)) {
	origin, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(origin)
	path := t.TempDir()
	if err := os.Chdir(path); err != nil {
		t.Fatal(err)
	}
	f(path)
}
