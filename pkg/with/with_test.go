package with

import (
	"github.com/quiteclose/goatly/pkg/declare/assert"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

func TestTempDir(t *testing.T) {
	var dirPath string
	TempDir(t, func(path string) {
		dirPath = path
		assert.DirExists(t, dirPath, "TempDir must create a directory.")
	})
	assert.NotPathExists(t, dirPath, "TempDir must remove the directory.")
}
