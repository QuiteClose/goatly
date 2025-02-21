// Re-implementation of the internal/unless package generated by declare.py
// Each function calling t.Errorf with an ExpectError message if the
// condition is not met.
package expect

import (
	"github.com/quiteclose/goatly/internal/unless"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// DirExists will call t.Errorf unless the directory exists
func DirExists(t *testing.T, path string, message string) {
	unless.DirExists(path, func(s string) {
		t.Errorf("ExpectError: %s (%s)", s, message)
	})
}

// Equal will call t.Errorf unless a == b
func Equal(t *testing.T, a, b interface{}, message string) {
	unless.Equal(a, b, func(s string) {
		t.Errorf("ExpectError: %s (%s)", s, message)
	})
}

// False will call t.Errorf unless a == false
func False(t *testing.T, a bool, message string) {
	unless.False(a, func(s string) {
		t.Errorf("ExpectError: %s (%s)", s, message)
	})
}

// Nil will call t.Errorf unless a == nil
func Nil(t *testing.T, a interface{}, message string) {
	unless.Nil(a, func(s string) {
		t.Errorf("ExpectError: %s (%s)", s, message)
	})
}

// True will call t.Errorf unless a == true
func True(t *testing.T, a bool, message string) {
	unless.True(a, func(s string) {
		t.Errorf("ExpectError: %s (%s)", s, message)
	})
}
