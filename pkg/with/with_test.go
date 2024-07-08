package with

import (
	"os"
	"testing"

	"github.com/quiteclose/goatly/pkg/declare/assert"
	"github.com/quiteclose/goatly/pkg/declare/expect"
)

///////////////////////////////////////////////////////////////////////////////

// TempWorkingDir should create a temporary directory and change the working directory to it.
func TestTempWorkingDir(t *testing.T) {
	var startDir, expectedWorkingDir, actualWorkingDir, endDir string
	startDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	TempWorkingDir(t, func(path string) {
		expectedWorkingDir = path
		actualWorkingDir, err = os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		assert.DirExists(t, actualWorkingDir, "TempDir must create a directory.")
	})
	expect.NotPathExists(t, actualWorkingDir, "TempDir must remove the directory.")
	expect.Equal(t, expectedWorkingDir, actualWorkingDir, "TempWorkingDir must change the working directory.")
	expect.NotEqual(t, startDir, actualWorkingDir, "TempWorkingDir must change the working directory.")
	expect.Equal(t, startDir, endDir, "TempWorkingDir must restore the working directory.")
	endDir, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
}
