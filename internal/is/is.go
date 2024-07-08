/* A collection of boolean functions to check if a condition is met
or generate a human-readable reason why the condition was missed.
*/
package is

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

///////////////////////////////////////////////////////////////////////////////

// Any checks if any item from sub-set b is in a
func Any(a, b interface{}) (bool, string) {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.String && aVal.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", a)
	}

	if bVal.Kind() != reflect.String && bVal.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", b)
	}

	if aVal.Kind() == reflect.String && bVal.Kind() == reflect.String {
		return Contains(a.(string), b.(string))
	}

	for i := 0; i < bVal.Len(); i++ {
		for j := 0; j < aVal.Len(); j++ {
			if reflect.DeepEqual(aVal.Index(j).Interface(), bVal.Index(i).Interface()) {
				return true, ""
			}
		}
	}
	return false, fmt.Sprintf("no item from %#v is in %#v", b, a)
}

// Contains checks if a contains b (sub-set b is in a)
func Contains(a, b interface{}) (bool, string) {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.String && aVal.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", a)
	}

	if bVal.Kind() != reflect.String && bVal.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", b)
	}

	if aVal.Kind() == reflect.String && bVal.Kind() == reflect.String {
		if strings.Contains(a.(string), b.(string)) {
			return true, ""
		}
		return false, fmt.Sprintf("%#v does not contain %#v", a, b)
	}

	for i := 0; i <= aVal.Len()-bVal.Len(); i++ {
		match := true
		for j := 0; j < bVal.Len(); j++ {
			if !reflect.DeepEqual(aVal.Index(i+j).Interface(), bVal.Index(j).Interface()) {
				match = false
				break
			}
		}
		if match {
			return true, ""
		}
	}
	return false, fmt.Sprintf("%#v does not contain %#v", a, b)
}

// DirExists checks if the directory exists
func DirExists(path string) (bool, string) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true, ""
	}
	return false, fmt.Sprintf("Path \"%s\" does not exist.", path)
}

// Empty checks if the string or slice is empty
func Empty(a interface{}) (bool, string) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", a)
	}
	if v.Len() == 0 {
		return true, ""
	}
	return false, fmt.Sprintf("%#v is not empty", a)
}

// Equal checks if a == b
func Equal(a, b interface{}) (bool, string) {
	if a == b {
		return true, ""
	}
	return false, fmt.Sprintf("%#v != %#v", a, b)
}

// False checks if a == false
func False(a bool) (bool, string) {
	if !a {
		return true, ""
	}
	return false, "true != false"
}

// FileExists checks if the file exists
func FileExists(path string) (bool, string) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Sprintf("Path \"%s\" does not exist.", path)
		}
		return false, fmt.Sprintf("Error occurred: %v", err)
	}
	if info.IsDir() {
		return false, fmt.Sprintf("Path %q is a directory", path)
	}
	return true, ""
}

// GreaterThan checks if a > b
func GreaterThan(a, b int) (bool, string) {
	if a > b {
		return true, ""
	}
	return false, fmt.Sprintf("%d is not greater than %d", a, b)
}

// GreaterThanOrEqual checks if a >= b
func GreaterThanOrEqual(a, b int) (bool, string) {
	if a >= b {
		return true, ""
	}
	return false, fmt.Sprintf("%d is not greater than or equal to %d", a, b)
}

// LessThan checks if a < b
func LessThan(a, b int) (bool, string) {
	if a < b {
		return true, ""
	}
	return false, fmt.Sprintf("%d is not less than %d", a, b)
}

// LessThanOrEqual checks if a <= b
func LessThanOrEqual(a, b int) (bool, string) {
	if a <= b {
		return true, ""
	}
	return false, fmt.Sprintf("%d is not less than or equal to %d", a, b)
}

// LongerThan checks if the length of the string or slice is greater than n
func LongerThan(a interface{}, n int) (bool, string) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", a)
	}
	if v.Len() > n {
		return true, ""
	}
	return false, fmt.Sprintf("%#v is not longer than %d", a, n)
}

// Matches checks if a matches the regex pattern
func Matches(a, pattern string) (bool, string) {
	matched, err := regexp.MatchString(pattern, a)
	if err != nil {
		return false, fmt.Sprintf("error occurred: %v", err)
	}
	if matched {
		return true, ""
	}
	return false, fmt.Sprintf("%#v does not match pattern %#v", a, pattern)
}

// Nil checks if a == nil
func Nil(a interface{}) (bool, string) {
	return Equal(a, nil)
}

// NotAny checks if no item from sub-set b is in a
func NotAny(a, b interface{}) (bool, string) {
	if ok, _ := Any(a, b); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("some item from %#v is in %#v", b, a)
}

// NotContains checks if a does not contain b (sub-set b is not in a)
func NotContains(a, b interface{}) (bool, string) {
	if ok, _ := Contains(a, b); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("%#v contains %#v", a, b)
}

// NotDirExists checks if the directory does not exist
func NotDirExists(path string) (bool, string) {
	if ok, _ := DirExists(path); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("Path \"%s\" exists.", path)
}

// NotEmpty checks if the string or slice is not empty
func NotEmpty(a interface{}) (bool, string) {
	if ok, _ := Empty(a); !ok {
		return true, ""
	}
	return false, "string or slice is empty"
}

// NotEqual checks if a != b
func NotEqual(a, b interface{}) (bool, string) {
	if ok, _ := Equal(a, b); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("%#v == %#v", a, b)
}

// NotFileExists checks if the file does not exist
func NotFileExists(path string) (bool, string) {
	if ok, _ := FileExists(path); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("Path \"%s\" exists.", path)
}

// NotMatches checks if a does not match the regex pattern
func NotMatches(a, pattern string) (bool, string) {
	matched, err := regexp.MatchString(pattern, a)
	if err != nil {
		return false, fmt.Sprintf("error occurred: %v", err)
	}
	if !matched {
		return true, ""
	}
	return false, fmt.Sprintf("%#v matches pattern %#v", a, pattern)
}

// NotPathExists checks if the path does not exist
func NotPathExists(path string) (bool, string) {
	if ok, _ := PathExists(path); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("Path \"%s\" exists.", path)
}

// NotTimeWithin checks if the time is not within the duration
func NotTimeWithin(t1, t2 time.Time, d time.Duration) (bool, string) {
	if ok, _ := TimeWithin(t1, t2, d); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("%v is within %v of %v", t1, d, t2)
}

// PathExists checks if the path exists
func PathExists(path string) (bool, string) {
	_, err := os.Stat(path)
	if err == nil {
		return true, ""
	}
	return false, fmt.Sprintf("Path \"%s\" does not exist.", path)
}

// ShorterThan checks if the length of the string or slice is less than n
func ShorterThan(a interface{}, n int) (bool, string) {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		return false, fmt.Sprintf("%#v is not a string or slice", a)
	}
	if v.Len() < n {
		return true, ""
	}
	return false, fmt.Sprintf("%#v is not shorter than %d", a, n)
}

// TimeAfter checks if t1 is after t2
func TimeAfter(t1, t2 time.Time) (bool, string) {
	if t1.After(t2) {
		return true, ""
	}
	return false, fmt.Sprintf("%v is not after %v", t1, t2)
}

// TimeBefore checks if t1 is before t2
func TimeBefore(t1, t2 time.Time) (bool, string) {
	if t1.Before(t2) {
		return true, ""
	}
	return false, fmt.Sprintf("%v is not before %v", t1, t2)
}

// TimeWithin checks if the time is within the duration
func TimeWithin(t1, t2 time.Time, d time.Duration) (bool, string) {
	if t1.Sub(t2) <= d {
		return true, ""
	}
	return false, fmt.Sprintf("%v is not within %v of %v", t1, d, t2)
}

// True checks if a == true
func True(a bool) (bool, string) {
	if a {
		return true, ""
	}
	return false, "false != true"
}

// Zero checks if a is zero
func Zero(a interface{}) (bool, string) {
	if reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface()) {
		return true, ""
	}
	return false, fmt.Sprintf("%#v is not zero", a)
}

// NotZero checks if a is not zero
func NotZero(a interface{}) (bool, string) {
	if ok, _ := Zero(a); !ok {
		return true, ""
	}
	return false, fmt.Sprintf("%#v is zero", a)
}

