// Re-implementation of the internal/unless package generated by declare.py
// Each function calling t.Fatalf with an AssertError message if the
// condition is not met.
package assert

import (
	"github.com/quiteclose/goatly/internal/unless"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

// Contains will call t.Fatalf unless a contains b
func Contains(t *testing.T, a, b string, message string) {
	unless.Contains(a, b, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}

// DirExists will call t.Fatalf unless the directory exists
func DirExists(t *testing.T, path string, message string) {
	unless.DirExists(path, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}

// Equal will call t.Fatalf unless a == b
func Equal(t *testing.T, a, b interface{}, message string) {
	unless.Equal(a, b, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}

// False will call t.Fatalf unless a == false
func False(t *testing.T, a bool, message string) {
	unless.False(a, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}

// Nil will call t.Fatalf unless a == nil
func Nil(t *testing.T, a interface{}, message string) {
	unless.Nil(a, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}

// True will call t.Fatalf unless a == true
func True(t *testing.T, a bool, message string) {
	unless.True(a, func(s string) {
		t.Fatalf("AssertError: %s (%s)", s, message)
	})
}
