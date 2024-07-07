package unless

import (
	"testing"

	"github.com/quiteclose/goatly/pkg/message"
)

///////////////////////////////////////////////////////////////////////////////

// Any should return false unless the callback is called
func TestAnyReturnsConditionMet(t *testing.T) {
	called := false
	given := Any([]int{1, 2, 3}, []int{4}, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must not be called when the condition is met!")
	}
	if given != called {
		t.Errorf("Any must return true/false according to whether the callback was called!")
	}
}

// Any should return true if the callback is called
func TestAnyReturnsConditionMiss(t *testing.T) {
	called := false
	given := Any([]int{1, 2, 3}, []int{2}, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Any must return true/false according to whether the callback was called!")
	}
}

// Any should not call the callback if any item from sub-set b is in a
func TestAnyConditionMet(t *testing.T) {
	called := false
	Any([]int{1, 2, 3}, []int{2}, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Any should call the callback if no item from sub-set b is in a
func TestAnyConditionMiss(t *testing.T) {
	called := false
	Any([]int{1, 2, 3}, []int{4}, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

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

// Empty should return false unless the callback is called
func TestEmptyReturnsConditionMet(t *testing.T) {
	called := false
	given := Empty("", func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Empty must only return true when the callback is called!")
	}
}

// Empty should return true if the callback is called
func TestEmptyReturnsConditionMiss(t *testing.T) {
	called := false
	given := Empty("non-empty", func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Empty must return true/false according to whether the callback was called!")
	}
}

// Empty should not call the callback if the string or slice is empty
func TestEmptyConditionMet(t *testing.T) {
	called := false
	Empty("", func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Empty should call the callback if the string or slice is not empty
func TestEmptyConditionMiss(t *testing.T) {
	called := false
	Empty("non-empty", func(s string) {
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

// LongerThan should return false unless the callback is called
func TestLongerThanReturnsConditionMet(t *testing.T) {
	called := false
	given := LongerThan("abcd", 2, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("LongerThan must return true/false according to whether the callback was called!")
	}
}

// LongerThan should return true if the callback is called
func TestLongerThanReturnsConditionMiss(t *testing.T) {
	called := false
	given := LongerThan("a", 2, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("LongerThan must return true/false according to whether the callback was called!")
	}
}

// LongerThan should not call the callback if the string or slice is longer than n
func TestLongerThanConditionMet(t *testing.T) {
	called := false
	LongerThan("abcd", 2, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// LongerThan should call the callback if the string or slice is not longer than n
func TestLongerThanConditionMiss(t *testing.T) {
	called := false
	LongerThan("a", 2, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// Matches should return false unless the callback is called
func TestMatchesReturnsConditionMet(t *testing.T) {
	called := false
	given := Matches("abc", "abc", func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Matches must return true/false according to whether the callback was called!")
	}
}

// Matches should return true if the callback is called
func TestMatchesReturnsConditionMiss(t *testing.T) {
	called := false
	given := Matches("abc", "xyz", func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Matches must return true/false according to whether the callback was called!")
	}
}

// Matches should not call the callback if the string matches the regex pattern
func TestMatchesConditionMet(t *testing.T) {
	called := false
	Matches("abc", "abc", func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Matches should call the callback if the string does not match the regex pattern
func TestMatchesConditionMiss(t *testing.T) {
	called := false
	Matches("abc", "xyz", func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// Nil should return false unless the callback is called
func TestNilReturnsConditionMet(t *testing.T) {
	called := false
	given := Nil(nil, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Nil must only return true when the callback is called!")
	}
}

// Nil should return true if the callback is called
func TestNilReturnsConditionMiss(t *testing.T) {
	called := false
	given := Nil(1, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("Nil must return true/false according to whether the callback was called!")
	}
}

// Nil should not call the callback if the value is nil
func TestNilConditionMet(t *testing.T) {
	called := false
	Nil(nil, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// Nil should call the callback if the value is not nil
func TestNilConditionMiss(t *testing.T) {
	called := false
	Nil(1, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// NotAny should return false unless the callback is called
func TestNotAnyReturnsConditionMet(t *testing.T) {
	called := false
	given := NotAny([]int{1, 2, 3}, []int{2}, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotAny must return true/false according to whether the callback was called!")
	}
}

// NotAny should return true if the callback is called
func TestNotAnyReturnsConditionMiss(t *testing.T) {
	called := false
	given := NotAny([]int{1, 2, 3}, []int{4}, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotAny must return true/false according to whether the callback was called!")
	}
}

// NotAny should not call the callback if no item from sub-set b is in a
func TestNotAnyConditionMet(t *testing.T) {
	called := false
	NotAny([]int{1, 2, 3}, []int{4}, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// NotAny should call the callback if any item from sub-set b is in a
func TestNotAnyConditionMiss(t *testing.T) {
	called := false
	NotAny([]int{1, 2, 3}, []int{2}, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// NotContains should return false unless the callback is called
func TestNotContainsReturnsConditionMet(t *testing.T) {
	called := false
	given := NotContains("a", "b", func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotContains must return true/false according to whether the callback was called!")
	}
}

// NotContains should return true if the callback is called
func TestNotContainsReturnsConditionMiss(t *testing.T) {
	called := false
	given := NotContains("a", "a", func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotContains must return true/false according to whether the callback was called!")
	}
}

// NotContains should not call the callback if the string or slice does not contain the sub-set
func TestNotContainsConditionMet(t *testing.T) {
	called := false
	NotContains("a", "b", func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// NotContains should call the callback if the string or slice contains the sub-set
func TestNotContainsConditionMiss(t *testing.T) {
	called := false
	NotContains("a", "a", func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// NotDirExists should return false unless the callback is called
func TestNotDirExistsReturnsConditionMet(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	given := NotDirExists(nonExistingPath, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotDirExists must only return true when the callback is called!")
	}
}

// NotDirExists should return true if the callback is called
func TestNotDirExistsReturnsConditionMiss(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	given := NotDirExists(existingPath, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotDirExists must return true/false according to whether the callback was called!")
	}
}

// NotDirExists should not call the callback if the directory does not exist
func TestNotDirExistsConditionMet(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	NotDirExists(nonExistingPath, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// NotDirExists should call the callback if the directory exists
func TestNotDirExistsConditionMiss(t *testing.T) {
	called := false
	existingPath := t.TempDir()
	NotDirExists(existingPath, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}

// NotEmpty should return false unless the callback is called
func TestNotEmptyReturnsConditionMet(t *testing.T) {
	called := false
	given := NotEmpty("non-empty", func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotEmpty must only return true when the callback is called!")
	}
}

// NotEmpty should return true if the callback is called
func TestNotEmptyReturnsConditionMiss(t *testing.T) {
	called := false
	given := NotEmpty("", func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("NotEmpty must return true/false according to whether the callback was called!")
	}
}

// NotEmpty should not call the callback if the string or slice is not empty
func TestNotEmptyConditionMet(t *testing.T) {
	called := false
	NotEmpty("non-empty", func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// NotEmpty should call the callback if the string or slice is empty
func TestNotEmptyConditionMiss(t *testing.T) {
	called := false
	NotEmpty("", func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
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
		t.Errorf("True must only return true when the callback is called!")
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

// LongerThan should return false unless the callback is called
func TestShorterThanReturnsConditionMet(t *testing.T) {
	called := false
	given := ShorterThan("a", 2, func(s string) {
		called = true
	})
	if called {
		t.Errorf("Callback must only be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("ShorterThan must return true/false according to whether the callback was called!")
	}
}

// ShorterThan should return true if the callback is called
func TestShorterThanReturnsConditionMiss(t *testing.T) {
	called := false
	given := ShorterThan("abcd", 2, func(s string) {
		called = true
	})
	if !called {
		t.Errorf("Callback must be called when the condition is missed!")
	}
	if given != called {
		t.Errorf("ShorterThan must return true/false according to whether the callback was called!")
	}
}

// ShorterThan should not call the callback if the string or slice is shorter than n
func TestShorterThanConditionMet(t *testing.T) {
	called := false
	ShorterThan("a", 2, func(s string) {
		called = true
	})
	if called {
		t.Error("Callback must not be called when the condition is met!")
	}
}

// ShorterThan should call the callback if the string or slice is not shorter than n
func TestShorterThanConditionMiss(t *testing.T) {
	called := false
	ShorterThan("abcd", 2, func(s string) {
		called = true
	})
	if !called {
		t.Error("Callback must be called when the condition is missed!")
	}
}
