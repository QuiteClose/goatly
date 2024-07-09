// Re-implementation of the internal/is package generated by declare.py
// If the condition is not met then t.Errorf is called with an
// ExpectFailed message. Returns false if the test passed, otherwise true.
package expect

import (
	"testing"
	"time"

	"github.com/quiteclose/goatly/internal/is"
	//"github.com/quiteclose/goatly/pkg/run"
)

///////////////////////////////////////////////////////////////////////////////

// Any call t.Errorf unless any item from sub-set b is in a
func Any(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Any(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Contains call t.Errorf unless a contains b (sub-set b is in a)
func Contains(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Contains(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// DirExists call t.Errorf unless the directory exists
func DirExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.DirExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Empty call t.Errorf unless the string or slice is empty
func Empty(t *testing.T, a interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Empty(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Equal call t.Errorf unless a == b
func Equal(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Equal(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// False call t.Errorf unless a == false
func False(t *testing.T, a bool, message string) bool {
	t.Helper()
	conditionMet, reason := is.False(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// FileExists call t.Errorf unless the file exists
func FileExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.FileExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// GreaterThan call t.Errorf unless a > b
func GreaterThan(t *testing.T, a, b int, message string) bool {
	t.Helper()
	conditionMet, reason := is.GreaterThan(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// GreaterThanOrEqual call t.Errorf unless a >= b
func GreaterThanOrEqual(t *testing.T, a, b int, message string) bool {
	t.Helper()
	conditionMet, reason := is.GreaterThanOrEqual(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// LessThan call t.Errorf unless a < b
func LessThan(t *testing.T, a, b int, message string) bool {
	t.Helper()
	conditionMet, reason := is.LessThan(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// LessThanOrEqual call t.Errorf unless a <= b
func LessThanOrEqual(t *testing.T, a, b int, message string) bool {
	t.Helper()
	conditionMet, reason := is.LessThanOrEqual(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// LongerThan call t.Errorf unless the length of the string or slice is greater than n
func LongerThan(t *testing.T, a interface{}, n int, message string) bool {
	t.Helper()
	conditionMet, reason := is.LongerThan(a, n)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Matches call t.Errorf unless a matches the regex pattern
func Matches(t *testing.T, a, pattern string, message string) bool {
	t.Helper()
	conditionMet, reason := is.Matches(a, pattern)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Nil call t.Errorf unless a == nil
func Nil(t *testing.T, a interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Nil(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotAny call t.Errorf unless no item from sub-set b is in a
func NotAny(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotAny(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotContains call t.Errorf unless a does not contain b (sub-set b is not in a)
func NotContains(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotContains(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotDirExists call t.Errorf unless the directory does not exist
func NotDirExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotDirExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotEmpty call t.Errorf unless the string or slice is not empty
func NotEmpty(t *testing.T, a interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotEmpty(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotEqual call t.Errorf unless a != b
func NotEqual(t *testing.T, a, b interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotEqual(a, b)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotFileExists call t.Errorf unless the file does not exist
func NotFileExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotFileExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotMatches call t.Errorf unless a does not match the regex pattern
func NotMatches(t *testing.T, a, pattern string, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotMatches(a, pattern)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotPathExists call t.Errorf unless the path does not exist
func NotPathExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotPathExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotTimeWithin call t.Errorf unless the time is not within the duration
func NotTimeWithin(t *testing.T, t1, t2 time.Time, d time.Duration, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotTimeWithin(t1, t2, d)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// NotZero call t.Errorf unless a is not zero
func NotZero(t *testing.T, a interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.NotZero(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// PathExists call t.Errorf unless the path exists
func PathExists(t *testing.T, path string, message string) bool {
	t.Helper()
	conditionMet, reason := is.PathExists(path)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// ShorterThan call t.Errorf unless the length of the string or slice is less than n
func ShorterThan(t *testing.T, a interface{}, n int, message string) bool {
	t.Helper()
	conditionMet, reason := is.ShorterThan(a, n)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// TimeAfter call t.Errorf unless t1 is after t2
func TimeAfter(t *testing.T, t1, t2 time.Time, message string) bool {
	t.Helper()
	conditionMet, reason := is.TimeAfter(t1, t2)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// TimeBefore call t.Errorf unless t1 is before t2
func TimeBefore(t *testing.T, t1, t2 time.Time, message string) bool {
	t.Helper()
	conditionMet, reason := is.TimeBefore(t1, t2)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// TimeWithin call t.Errorf unless the time is within the duration
func TimeWithin(t *testing.T, t1, t2 time.Time, d time.Duration, message string) bool {
	t.Helper()
	conditionMet, reason := is.TimeWithin(t1, t2, d)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// True call t.Errorf unless a == true
func True(t *testing.T, a bool, message string) bool {
	t.Helper()
	conditionMet, reason := is.True(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}

// Zero call t.Errorf unless a is zero
func Zero(t *testing.T, a interface{}, message string) bool {
	t.Helper()
	conditionMet, reason := is.Zero(a)
	if !conditionMet {
		t.Errorf("%s\nExpectFailed: %s", message, reason)
	}
	return conditionMet
}
