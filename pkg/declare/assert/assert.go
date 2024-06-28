package assert

import (
	"testing"
	"github.com/quiteclose/goatly/internal/unless"
)

///////////////////////////////////////////////////////////////////////////////

func Equal(t *testing.T, a, b interface{}, message string) {
	unless.Equal(a, b, func(s string) {
		t.Errorf("AssertError: %s (%s)", s, message)
	})
}
