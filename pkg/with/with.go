// Convenience functions for setting up environments before running tests.

package with

import (
	"os"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// TempWorkingDir creates a working directory and changes the working directory to it.
// the given funtion is called with the path of the working directory.
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

// TempWorkingTree changes the working directory to a new temporary directory
// and populates it with the given directories and files. The given function is
// then called with the path of the new working directory. The given files map
// filenames to their content and are created after the directories.
func TempWorkingTree(t *testing.T, directories []string, files map[string]string, f func(string)) {
	TempWorkingDir(t, func(path string) {
		for _, dir := range directories {
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatal(err)
			}
		}
		for filename, content := range files {
			if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
				t.Fatal(err)
			}
		}
		f(path)
	})
}
