package unless

import (
	"testing"

	"github.com/quiteclose/goatly/pkg/message"
)

///////////////////////////////////////////////////////////////////////////////

// Equal should not call the callback unless a == b
func TestEqualConditionMet(t *testing.T) {
	called := false
	Equal(1, 1, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Equal should call the callback if a != b
func TestEqualConditionMiss(t *testing.T) {
	called := false
	Equal(0, 1, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// Equal should correctly represent unequal integers, e.g. 0 != 1
func TestEqualIntMessage(t *testing.T) {
	var expected, found string
	Equal(0, 1, func(s string) {
		found = s
	})
	expected = `0 != 1`
	if expected != found {
		t.Errorf(message.UnexpectedText("message", expected, found))
	}
}

// Equal should correctly represent unequal strings, e.g. "a" != "b"
func TestEqualStringMessage(t *testing.T) {
	var expected, found string
	Equal("a", "b", func(s string) {
		found = s
	})
	expected = `"a" != "b"`
	if found != expected {
		t.Errorf(message.UnexpectedText("message", expected, found))
	}
}

func TestFalseConditionMet(t *testing.T) {
	called := false
	False(false, func(s string) {
		called = true
	})
	if called {
		t.Error("False executed the callback when value was false!")
	}
}

func TestFalseConditionMiss(t *testing.T) {
	called := false
	False(true, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

func TestFalseMessage(t *testing.T) {
	var expected, found string
	False(true, func(s string) {
		found = s
	})
	expected = `true != false`
	if found != expected {
		t.Errorf(message.UnexpectedValue("found", expected, found))
	}
}
