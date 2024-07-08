package is

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/quiteclose/goatly/pkg/message"
)

///////////////////////////////////////////////////////////////////////////////

// Any should return false with an error message if no item from sub-set b is in a
func TestAnyReturnsFalse(t *testing.T) {
	expected := "no item from []int{4} is in []int{1, 2, 3}"
	given, reason := Any([]int{1, 2, 3}, []int{4})
	if given {
		t.Errorf("Any must return false if no item from sub-set b is in a")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Any should return true with no error message if any item from sub-set b is in a
func TestAnyReturnsTrue(t *testing.T) {
	given, reason := Any([]int{1, 2, 3}, []int{2})
	if !given {
		t.Errorf("Any must return true if any item from sub-set b is in a")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Contains should return false with an error message if a does not contain b
func TestContainsReturnsFalse(t *testing.T) {
	expected := `"a" does not contain "b"`
	given, reason := Contains("a", "b")
	if given {
		t.Errorf("Contains must return false if a does not contain b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Contains should return true with no error message if a contains b
func TestContainsReturnsTrue(t *testing.T) {
	given, reason := Contains("a", "a")
	if !given {
		t.Errorf("Contains must return true if a contains b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// DirExists should return false with an error message if the directory does not exist
func TestDirExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	expected := `Path "spoiled" does not exist.`
	given, reason := DirExists(nonExistingPath)
	if given {
		t.Errorf("DirExists must return false if the directory does not exist")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// DirExists should return true with no error message if the directory exists
func TestDirExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	given, reason := DirExists(existingPath)
	if !given {
		t.Errorf("DirExists must return true if the directory exists")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Empty should return false with an error message if the string or slice is not empty
func TestEmptyReturnsFalse(t *testing.T) {
	expected := `"non-empty" is not empty`
	given, reason := Empty("non-empty")
	if given {
		t.Errorf("Empty must return false if the string or slice is not empty")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Empty should return true with no error message if the string or slice is empty
func TestEmptyReturnsTrue(t *testing.T) {
	given, reason := Empty("")
	if !given {
		t.Errorf("Empty must return true if the string or slice is empty")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Equal should return false with an error message if a != b
func TestEqualReturnsFalse(t *testing.T) {
	expected := `0 != 1`
	given, reason := Equal(0, 1)
	if given {
		t.Errorf("Equal must return false if a != b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Equal should return true with no error message if a == b
func TestEqualReturnsTrue(t *testing.T) {
	given, reason := Equal(1, 1)
	if !given {
		t.Errorf("Equal must return true if a == b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Error should return false with an error message if err is nil
func TestErrorReturnsFalse(t *testing.T) {
	expected := "no error occurred"
	given, reason := Error(nil)
	if given {
		t.Errorf("Error must return false if err is nil")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Error should return true with no error message if err is not nil
func TestErrorReturnsTrue(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	given, reason := Error(err)
	if !given {
		t.Errorf("Error must return true if err is not nil")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// ErrorContains should return false with an error message if the error message does not contain the substring
func TestErrorContainsReturnsFalse(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	expected := "error does not contain substring"
	given, reason := ErrorContains(err, "substring")
	if given {
		t.Errorf("ErrorContains must return false if the error message does not contain the substring")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// ErrorContains should return true with no error message if the error message contains the substring
func TestErrorContainsReturnsTrue(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	given, reason := ErrorContains(err, "error")
	if !given {
		t.Errorf("ErrorContains must return true if the error message contains the substring")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// ErrorType should return false with an error message if the error is not of the given type
func TestErrorTypeReturnsFalse(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	expected := "error is not of type *os.PathError"
	given, reason := ErrorType(err, &os.PathError{})
	if given {
		t.Errorf("ErrorType must return false if the error is not of the given type")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// ErrorType should return true with no error message if the error is of the given type
func TestErrorTypeReturnsTrue(t *testing.T) {
	err := &os.PathError{}
	given, reason := ErrorType(err, &os.PathError{})
	if !given {
		t.Errorf("ErrorType must return true if the error is of the given type")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// False should return false with an error message if a == true
func TestFalseReturnsFalse(t *testing.T) {
	expected := "true != false"
	given, reason := False(true)
	if given {
		t.Errorf("False must return false if a == true")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// False should return true with no error message if a == false
func TestFalseReturnsTrue(t *testing.T) {
	given, reason := False(false)
	if !given {
		t.Errorf("False must return true if a == false")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// FileExists should return false with an error message if the file does not exist
func TestFileExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	expected := `Path "spoiled" does not exist.`
	given, reason := FileExists(nonExistingPath)
	if given {
		t.Errorf("FileExists must return false if the file does not exist")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// FileExists should return true with no error message if the file exists
func TestFileExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	filePath := existingPath + "/testfile"
	file, _ := os.Create(filePath)
	defer file.Close()

	given, reason := FileExists(filePath)
	if !given {
		t.Errorf("FileExists must return true if the file exists")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// GreaterThan should return false with an error message if a <= b
func TestGreaterThanReturnsFalse(t *testing.T) {
	expected := "1 is not greater than 2"
	given, reason := GreaterThan(1, 2)
	if given {
		t.Errorf("GreaterThan must return false if a <= b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// GreaterThan should return true with no error message if a > b
func TestGreaterThanReturnsTrue(t *testing.T) {
	given, reason := GreaterThan(2, 1)
	if !given {
		t.Errorf("GreaterThan must return true if a > b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// GreaterThanOrEqual should return false with an error message if a < b
func TestGreaterThanOrEqualReturnsFalse(t *testing.T) {
	expected := "1 is not greater than or equal to 2"
	given, reason := GreaterThanOrEqual(1, 2)
	if given {
		t.Errorf("GreaterThanOrEqual must return false if a < b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// GreaterThanOrEqual should return true with no error message if a >= b
func TestGreaterThanOrEqualReturnsTrue(t *testing.T) {
	given, reason := GreaterThanOrEqual(2, 1)
	if !given {
		t.Errorf("GreaterThanOrEqual must return true if a >= b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// LessThan should return false with an error message if a >= b
func TestLessThanReturnsFalse(t *testing.T) {
	expected := "2 is not less than 1"
	given, reason := LessThan(2, 1)
	if given {
		t.Errorf("LessThan must return false if a >= b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// LessThan should return true with no error message if a < b
func TestLessThanReturnsTrue(t *testing.T) {
	given, reason := LessThan(1, 2)
	if !given {
		t.Errorf("LessThan must return true if a < b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// LessThanOrEqual should return false with an error message if a > b
func TestLessThanOrEqualReturnsFalse(t *testing.T) {
	expected := "2 is not less than or equal to 1"
	given, reason := LessThanOrEqual(2, 1)
	if given {
		t.Errorf("LessThanOrEqual must return false if a > b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// LessThanOrEqual should return true with no error message if a <= b
func TestLessThanOrEqualReturnsTrue(t *testing.T) {
	given, reason := LessThanOrEqual(1, 2)
	if !given {
		t.Errorf("LessThanOrEqual must return true if a <= b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// LongerThan should return false with an error message if the length of the string or slice is not longer than n
func TestLongerThanReturnsFalse(t *testing.T) {
	expected := `"a" is not longer than 2`
	given, reason := LongerThan("a", 2)
	if given {
		t.Errorf("LongerThan must return false if the length of the string or slice is not longer than n")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// LongerThan should return true with no error message if the length of the string or slice is longer than n
func TestLongerThanReturnsTrue(t *testing.T) {
	given, reason := LongerThan("abcd", 2)
	if !given {
		t.Errorf("LongerThan must return true if the length of the string or slice is longer than n")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Matches should return false with an error message if the string does not match the regex pattern
func TestMatchesReturnsFalse(t *testing.T) {
	expected := `"abc" does not match pattern "xyz"`
	given, reason := Matches("abc", "xyz")
	if given {
		t.Errorf("Matches must return false if the string does not match the regex pattern")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Matches should return true with no error message if the string matches the regex pattern
func TestMatchesReturnsTrue(t *testing.T) {
	given, reason := Matches("abc", "abc")
	if !given {
		t.Errorf("Matches must return true if the string matches the regex pattern")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Nil should return false with an error message if the value is not nil
func TestNilReturnsFalse(t *testing.T) {
	expected := `1 != <nil>`
	given, reason := Nil(1)
	if given {
		t.Errorf("Nil must return false if the value is not nil")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Nil should return true with no error message if the value is nil
func TestNilReturnsTrue(t *testing.T) {
	given, reason := Nil(nil)
	if !given {
		t.Errorf("Nil must return true if the value is nil")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotAny should return false with an error message if any item from sub-set b is in a
func TestNotAnyReturnsFalse(t *testing.T) {
	expected := "some item from []int{2} is in []int{1, 2, 3}"
	given, reason := NotAny([]int{1, 2, 3}, []int{2})
	if given {
		t.Errorf("NotAny must return false if any item from sub-set b is in a")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotAny should return true with no error message if no item from sub-set b is in a
func TestNotAnyReturnsTrue(t *testing.T) {
	given, reason := NotAny([]int{1, 2, 3}, []int{4})
	if !given {
		t.Errorf("NotAny must return true if no item from sub-set b is in a")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotContains should return false with an error message if a contains b
func TestNotContainsReturnsFalse(t *testing.T) {
	expected := `"a" contains "a"`
	given, reason := NotContains("a", "a")
	if given {
		t.Errorf("NotContains must return false if a contains b")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotContains should return true with no error message if a does not contain b
func TestNotContainsReturnsTrue(t *testing.T) {
	given, reason := NotContains("a", "b")
	if !given {
		t.Errorf("NotContains must return true if a does not contain b")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotDirExists should return false with an error message if the directory exists
func TestNotDirExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	expected := `Path "existingPath" exists.`
	given, reason := NotDirExists(existingPath)
	if given {
		t.Errorf("NotDirExists must return false if the directory exists")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotDirExists should return true with no error message if the directory does not exist
func TestNotDirExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	given, reason := NotDirExists(nonExistingPath)
	if !given {
		t.Errorf("NotDirExists must return true if the directory does not exist")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotEmpty should return false with an error message if the string or slice is empty
func TestNotEmptyReturnsFalse(t *testing.T) {
	expected := "string or slice is empty"
	given, reason := NotEmpty("")
	if given {
		t.Errorf("NotEmpty must return false if the string or slice is empty")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotEmpty should return true with no error message if the string or slice is not empty
func TestNotEmptyReturnsTrue(t *testing.T) {
	given, reason := NotEmpty("non-empty")
	if !given {
		t.Errorf("NotEmpty must return true if the string or slice is not empty")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotError should return false with an error message if err is not nil
func TestNotErrorReturnsFalse(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	expected := "error occurred: an error occurred"
	given, reason := NotError(err)
	if given {
		t.Errorf("NotError must return false if err is not nil")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotError should return true with no error message if err is nil
func TestNotErrorReturnsTrue(t *testing.T) {
	given, reason := NotError(nil)
	if !given {
		t.Errorf("NotError must return true if err is nil")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotErrorContains should return false with an error message if the error message contains the substring
func TestNotErrorContainsReturnsFalse(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	expected := "error contains error"
	given, reason := NotErrorContains(err, "error")
	if given {
		t.Errorf("NotErrorContains must return false if the error message contains the substring")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotErrorContains should return true with no error message if the error message does not contain the substring
func TestNotErrorContainsReturnsTrue(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	given, reason := NotErrorContains(err, "substring")
	if !given {
		t.Errorf("NotErrorContains must return true if the error message does not contain the substring")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotErrorType should return false with an error message if the error is of the given type
func TestNotErrorTypeReturnsFalse(t *testing.T) {
	err := &os.PathError{}
	expected := "error is of type *os.PathError"
	given, reason := NotErrorType(err, &os.PathError{})
	if given {
		t.Errorf("NotErrorType must return false if the error is of the given type")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotErrorType should return true with no error message if the error is not of the given type
func TestNotErrorTypeReturnsTrue(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	given, reason := NotErrorType(err, &os.PathError{})
	if !given {
		t.Errorf("NotErrorType must return true if the error is not of the given type")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotFileExists should return false with an error message if the file exists
func TestNotFileExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	filePath := existingPath + "/testfile"
	file, _ := os.Create(filePath)
	defer file.Close()

	expected := `Path "testfile" exists.`
	given, reason := NotFileExists(filePath)
	if given {
		t.Errorf("NotFileExists must return false if the file exists")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotFileExists should return true with no error message if the file does not exist
func TestNotFileExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	given, reason := NotFileExists(nonExistingPath)
	if !given {
		t.Errorf("NotFileExists must return true if the file does not exist")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotMatches should return false with an error message if the string matches the regex pattern
func TestNotMatchesReturnsFalse(t *testing.T) {
	expected := `"abc" matches pattern "abc"`
	given, reason := NotMatches("abc", "abc")
	if given {
		t.Errorf("NotMatches must return false if the string matches the regex pattern")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotMatches should return true with no error message if the string does not match the regex pattern
func TestNotMatchesReturnsTrue(t *testing.T) {
	given, reason := NotMatches("abc", "xyz")
	if !given {
		t.Errorf("NotMatches must return true if the string does not match the regex pattern")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotPathExists should return false with an error message if the path exists
func TestNotPathExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	expected := `Path "existingPath" exists.`
	given, reason := NotPathExists(existingPath)
	if given {
		t.Errorf("NotPathExists must return false if the path exists")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotPathExists should return true with no error message if the path does not exist
func TestNotPathExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	given, reason := NotPathExists(nonExistingPath)
	if !given {
		t.Errorf("NotPathExists must return true if the path does not exist")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotType should return false with an error message if a is of the given type
func TestNotTypeReturnsFalse(t *testing.T) {
	expected := "1 is of type int"
	given, reason := NotType(1, reflect.TypeOf(1))
	if given {
		t.Errorf("NotType must return false if a is of the given type")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotType should return true with no error message if a is not of the given type
func TestNotTypeReturnsTrue(t *testing.T) {
	given, reason := NotType(1, reflect.TypeOf(""))
	if !given {
		t.Errorf("NotType must return true if a is not of the given type")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotTimeWithin should return false with an error message if the time is within the duration
func TestNotTimeWithinReturnsFalse(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Second)
	expected := fmt.Sprintf("%v is within 1s of %v", t1, t2)
	given, reason := NotTimeWithin(t1, t2, 1*time.Second)
	if given {
		t.Errorf("NotTimeWithin must return false if the time is within the duration")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotTimeWithin should return true with no error message if the time is not within the duration
func TestNotTimeWithinReturnsTrue(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-2 * time.Second)
	given, reason := NotTimeWithin(t1, t2, 1*time.Second)
	if !given {
		t.Errorf("NotTimeWithin must return true if the time is not within the duration")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// PathExists should return false with an error message if the path does not exist
func TestPathExistsReturnsFalse(t *testing.T) {
	existingPath := t.TempDir()
	nonExistingPath := existingPath + "spoiled"
	expected := `Path "spoiled" does not exist.`
	given, reason := PathExists(nonExistingPath)
	if given {
		t.Errorf("PathExists must return false if the path does not exist")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// PathExists should return true with no error message if the path exists
func TestPathExistsReturnsTrue(t *testing.T) {
	existingPath := t.TempDir()
	given, reason := PathExists(existingPath)
	if !given {
		t.Errorf("PathExists must return true if the path exists")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// ShorterThan should return false with an error message if the length of the string or slice is not shorter than n
func TestShorterThanReturnsFalse(t *testing.T) {
	expected := `"abcd" is not shorter than 2`
	given, reason := ShorterThan("abcd", 2)
	if given {
		t.Errorf("ShorterThan must return false if the length of the string or slice is not shorter than n")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// ShorterThan should return true with no error message if the length of the string or slice is shorter than n
func TestShorterThanReturnsTrue(t *testing.T) {
	given, reason := ShorterThan("a", 2)
	if !given {
		t.Errorf("ShorterThan must return true if the length of the string or slice is shorter than n")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// TimeAfter should return false with an error message if t1 is not after t2
func TestTimeAfterReturnsFalse(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(1 * time.Second)
	expected := fmt.Sprintf("%v is not after %v", t1, t2)
	given, reason := TimeAfter(t1, t2)
	if given {
		t.Errorf("TimeAfter must return false if t1 is not after t2")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// TimeAfter should return true with no error message if t1 is after t2
func TestTimeAfterReturnsTrue(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Second)
	given, reason := TimeAfter(t1, t2)
	if !given {
		t.Errorf("TimeAfter must return true if t1 is after t2")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// TimeBefore should return false with an error message if t1 is not before t2
func TestTimeBeforeReturnsFalse(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Second)
	expected := fmt.Sprintf("%v is not before %v", t1, t2)
	given, reason := TimeBefore(t1, t2)
	if given {
		t.Errorf("TimeBefore must return false if t1 is not before t2")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// TimeBefore should return true with no error message if t1 is before t2
func TestTimeBeforeReturnsTrue(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(1 * time.Second)
	given, reason := TimeBefore(t1, t2)
	if !given {
		t.Errorf("TimeBefore must return true if t1 is before t2")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// TimeWithin should return false with an error message if the time is not within the duration
func TestTimeWithinReturnsFalse(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-2 * time.Second)
	expected := fmt.Sprintf("%v is not within 1s of %v", t1, t2)
	given, reason := TimeWithin(t1, t2, 1*time.Second)
	if given {
		t.Errorf("TimeWithin must return false if the time is not within the duration")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// TimeWithin should return true with no error message if the time is within the duration
func TestTimeWithinReturnsTrue(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Second)
	given, reason := TimeWithin(t1, t2, 2*time.Second)
	if !given {
		t.Errorf("TimeWithin must return true if the time is within the duration")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// True should return false with an error message if a == false
func TestTrueReturnsFalse(t *testing.T) {
	expected := "false != true"
	given, reason := True(false)
	if given {
		t.Errorf("True must return false if a == false")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// True should return true with no error message if a == true
func TestTrueReturnsTrue(t *testing.T) {
	given, reason := True(true)
	if !given {
		t.Errorf("True must return true if a == true")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Type should return false with an error message if a is not of the given type
func TestTypeReturnsFalse(t *testing.T) {
	expected := `"abc" is not of type int`
	given, reason := Type("abc", reflect.TypeOf(1))
	if given {
		t.Errorf("Type must return false if a is not of the given type")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Type should return true with no error message if a is of the given type
func TestTypeReturnsTrue(t *testing.T) {
	given, reason := Type("abc", reflect.TypeOf(""))
	if !given {
		t.Errorf("Type must return true if a is of the given type")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// Zero should return false with an error message if a is not zero
func TestZeroReturnsFalse(t *testing.T) {
	expected := `"abc" is not zero`
	given, reason := Zero("abc")
	if given {
		t.Errorf("Zero must return false if a is not zero")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// Zero should return true with no error message if a is zero
func TestZeroReturnsTrue(t *testing.T) {
	given, reason := Zero("")
	if !given {
		t.Errorf("Zero must return true if a is zero")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}

// NotZero should return false with an error message if a is zero
func TestNotZeroReturnsFalse(t *testing.T) {
	expected := `"" is zero`
	given, reason := NotZero("")
	if given {
		t.Errorf("NotZero must return false if a is zero")
	}
	if reason != expected {
		t.Errorf(message.MustEqual("error message", expected, reason))
	}
}

// NotZero should return true with no error message if a is not zero
func TestNotZeroReturnsTrue(t *testing.T) {
	given, reason := NotZero("abc")
	if !given {
		t.Errorf("NotZero must return true if a is not zero")
	}
	if reason != "" {
		t.Errorf("Error message must be empty if condition is met")
	}
}
