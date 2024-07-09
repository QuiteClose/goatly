package with

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/quiteclose/goatly/pkg/declare/assert"
	"github.com/quiteclose/goatly/pkg/declare/expect"
)

///////////////////////////////////////////////////////////////////////////////

// TempWorkingDir should create a temporary directory and change the working directory to it.
func TestTempWorkingDir(t *testing.T) {
	var startDir, expectedDir, actualDir, endDir string
	if currentDir, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		startDir = currentDir
	}
	t.Run("TempWorkingDirSubtest", func(st *testing.T) {
		TempWorkingDir(st, func(path string) {
			expectedDir = path
			if subtestDir, err := os.Getwd(); err != nil {
				st.Fatal(err)
			} else {
				actualDir = subtestDir
			}
			assert.DirExists(st, actualDir, "TempWorkingDir must create a directory.")
			expect.Equal(st, expectedDir, actualDir, "TempWorkingDir must provide the correct path.")
			expect.NotEqual(st, startDir, actualDir, "TempWorkingDir must change the working directory.")
		})
	})
	if currentDir, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		endDir = currentDir
	}
	expect.NotPathExists(t, actualDir, "TempWorkingDir must remove the directory.")
	expect.Equal(t, startDir, endDir, "TempWorkingDir must restore the working directory.")
}

// TempWorkingTree should behave like TempWorkingDir when not given directories or files.
func TestTempWorkingTree(t *testing.T) {
	var startDir, expectedDir, actualDir, endDir string
	directories := []string{}
	files := map[string]string{}
	if currentDir, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		startDir = currentDir
	}
	t.Run("TempWorkingTreeSubtest", func(st *testing.T) {
		TempWorkingTree(st, directories, files, func(path string) {
			expectedDir = path
			if subtestDir, err := os.Getwd(); err != nil {
				st.Fatal(err)
			} else {
				actualDir = subtestDir
			}
			assert.DirExists(st, actualDir, "TempWorkingTree must create a directory.")
			expect.Equal(st, expectedDir, actualDir, "TempWorkingTree must provide the correct path.")
			expect.NotEqual(st, startDir, actualDir, "TempWorkingTree must change the working directory.")
		})
	})
	if currentDir, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		endDir = currentDir
	}
	expect.NotPathExists(t, actualDir, "TempWorkingTree must remove the directory.")
	expect.Equal(t, startDir, endDir, "TempWorkingTree must restore the working directory.")
}

// TempWorkingTree should create directories within the working directory.
func TestTempWorkingTreeWithDirectories(t *testing.T) {
	var tempDir string
	directories := []string{"one", "two", "three"}
	files := map[string]string{}
	t.Run("TempWorkingTreeWithDirectoriesSubtest", func(st *testing.T) {
		TempWorkingTree(st, directories, files, func(path string) {
			tempDir = path
			for _, subDir := range directories {
				dirPath := filepath.Join(tempDir, subDir)
				expect.DirExists(st, dirPath, "TempWorkingTree must create sub-directories.")
			}
		})
	})
	expect.NotPathExists(t, tempDir, "TempWorkingTree must remove the directory.")
}

// TempWorkingTree should create nested sub-directories within the working directory.
func TestTempWorkingTreeWithNestedDirectories(t *testing.T) {
	var tempDir string
	directories := []string{"one", "one/two", "one/two/three", "one/four"}
	files := map[string]string{}
	t.Run("TempWorkingTreeWithNestedDirectoriesSubtest", func(st *testing.T) {
		TempWorkingTree(st, directories, files, func(path string) {
			tempDir = path
			for _, subDir := range directories {
				dirPath := filepath.Join(tempDir, subDir)
				expect.DirExists(st, dirPath, "TempWorkingTree must create nested sub-directories.")
			}
		})
	})
	expect.NotPathExists(t, tempDir, "TempWorkingTree must remove the directory.")
}

// TempWorkingTree should create files within the working directory.
func TestTempWorkingTreeWithFiles(t *testing.T) {
	var tempDir string
	directories := []string{}
	files := map[string]string{"one.txt": "one", "two.txt": "two", "three.txt": "three"}
	t.Run("TempWorkingTreeWithFilesSubtest", func(st *testing.T) {
		TempWorkingTree(st, directories, files, func(path string) {
			tempDir = path
			for fileName, fileContent := range files {
				filePath := filepath.Join(tempDir, fileName)
				expect.FileExists(st, filePath, "TempWorkingTree must create files.")
				expect.FileContent(st, filePath, fileContent, "TempWorkingTree must provide the correct content.")
			}
		})
	})
	expect.NotPathExists(t, tempDir, "TempWorkingTree must remove the directory.")
}
