package unless

import (
	"testing"

	"github.com/quiteclose/goatly/pkg/message"
)

///////////////////////////////////////////////////////////////////////////////

// Contains should return false unless the callback is called
func TestContainsReturnsConditionMet(t *testing.T) {
	called := false
	given := Contains("a", "a", func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Contains must return true/false according to whether the callback was called!")
	}
}

// Contains should return true if the callback is called
func TestContainsReturnsConditionMiss(t *testing.T) {
	called := false
	given := Contains("a", "b", func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Contains must return true/false according to whether the callback was called!")
	}
}

// Contains should not call the callback if the string contains the substring
func TestContainsConditionMet(t *testing.T) {
	called := false
	Contains("a", "a", func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Contains should call the callback if the string does not contain the substring
func TestContainsConditionMiss(t *testing.T) {
	called := false
	Contains("a", "b", func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// DirExists should return false unless the callback is called
func TestDirExistsReturnsConditionMet(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	given := DirExists(existingPath, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("DirExists must only return true when the callback is called!")
	}
}

// DirExists should return true if the callback is called
func TestDirExistsReturnsConditionMiss(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	given := DirExists(nonExistingPath, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("DirExists must return true/false according to whether the callback was called!")
	}
}

// DirExists should not call the callback if the directory exists
func TestDirExistsConditionMet(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	DirExists(existingPath, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// DirExists should call the callback if the directory does not exist
func TestDirExistsConditionMiss(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	DirExists(nonExistingPath, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// Equal should return false unless the callback is called
func TestEqualReturnsConditionMet(t *testing.T) {
	called := false
	given := Equal(1, 1, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Equal must only return true when the callback is called!")
	}
}

// Equal should return true if the callback is called
func TestEqualReturnsConditionMiss(t *testing.T) {
	called := false
	given := Equal(1, 2, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Equal must return true/false according to whether the callback was called!")
	}
}

// Equal should not call the callback if a == b
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
		t.Errorf(message.MustContain("message", found, expected))
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
		t.Errorf(message.MustContain("message", found, expected))
	}
}

// False should return false unless the callback is called
func TestFalseReturnsConditionMet(t *testing.T) {
	called := false
	given := False(false, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("False must only return true when the callback is called!")
	}
}

// False should return true if the callback is called
func TestFalseReturnsConditionMiss(t *testing.T) {
	called := false
	given := False(true, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("False must return true/false according to whether the callback was called!")
	}
}

func TestFalseMessage(t *testing.T) {
	var expected, found string
	False(true, func(s string) {
		found = s
	})
	expected = `true != false`
	if found != expected {
		t.Errorf(message.MustContain("found", found, expected))
	}
}

// True should return false unless the callback is called
func TestTrueReturnsConditionMet(t *testing.T) {
	called := false
	given := True(true, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("False must only return true when the callback is called!")
	}
}

// True should return true if the callback is called
func TestTrueReturnsConditionMiss(t *testing.T) {
	called := false
	given := True(false, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("True must return true/false according to whether the callback was called!")
	}
}

func TestTrueMessage(t *testing.T) {
	var expected, found string
	True(false, func(s string) {
		found = s
	})
	expected = `false != true`
	if found != expected {
		t.Errorf(message.MustEqual("found", expected, found))
	}
}
