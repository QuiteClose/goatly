package unless

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

///////////////////////////////////////////////////////////////////////////////

// Any will call the callback unless any item from sub-set b is in a
func Any(a, b interface{}, callback func(string)) bool {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.String && aVal.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}

	if bVal.Kind() != reflect.String && bVal.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", b))
		return true
	}

	if aVal.Kind() == reflect.String && bVal.Kind() == reflect.String {
		return Contains(a.(string), b.(string), callback)
	}

	for i := 0; i < bVal.Len(); i++ {
		for j := 0; j < aVal.Len(); j++ {
			if reflect.DeepEqual(aVal.Index(j).Interface(), bVal.Index(i).Interface()) {
				return false
			}
		}
	}
	callback(fmt.Sprintf("no item from %#v is in %#v", b, a))
	return true
}

// Contains will call the callback unless a contains b (sub-set b is in a)
func Contains(a, b interface{}, callback func(string)) bool {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.String && aVal.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}

	if bVal.Kind() != reflect.String && bVal.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", b))
		return true
	}

	if aVal.Kind() == reflect.String && bVal.Kind() == reflect.String {
		if !strings.Contains(a.(string), b.(string)) {
			callback(fmt.Sprintf("%#v does not contain %#v", a, b))
			return true
		}
		return false
	}

	contains := false
	for i := 0; i <= aVal.Len()-bVal.Len(); i++ {
		match := true
		for j := 0; j < bVal.Len(); j++ {
			if !reflect.DeepEqual(aVal.Index(i+j).Interface(), bVal.Index(j).Interface()) {
				match = false
				break
			}
		}
		if match {
			contains = true
			break
		}
	}
	if !contains {
		callback(fmt.Sprintf("%#v does not contain %#v", a, b))
	}
	return !contains
}

// DirExists will call the callback unless the directory exists
func DirExists(path string, callback func(string)) bool {
	conditionMet := false
	info, err := os.Stat(path)
	if err == nil {
		conditionMet = !info.IsDir()
	} else {
		conditionMet = os.IsNotExist(err)
	}
	if conditionMet {
		callback(fmt.Sprintf("%s is not a directory", path))
	}
	return conditionMet
}

// Empty will call the callback unless the string or slice is empty
func Empty(a interface{}, callback func(string)) bool {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}
	conditionMet := v.Len() != 0
	if conditionMet {
		callback(fmt.Sprintf("%#v is not empty", a))
	}
	return conditionMet
}

// Equal will call the callback unless a == b
func Equal(a, b interface{}, callback func(string)) bool {
	conditionMet := a != b
	if conditionMet {
		callback(fmt.Sprintf("%#v != %#v", a, b))
	}
	return conditionMet
}

// Error will call the callback unless err is not nil
func Error(err error, callback func(string)) bool {
	conditionMet := err == nil
	if conditionMet {
		callback("no error occurred")
	}
	return conditionMet
}

// ErrorContains will call the callback unless the error message contains the substring
func ErrorContains(err error, substr string, callback func(string)) bool {
	conditionMet := err == nil || !strings.Contains(err.Error(), substr)
	if conditionMet {
		callback(fmt.Sprintf("error does not contain %s", substr))
	}
	return conditionMet
}

// ErrorType will call the callback unless the error is of the given type
func ErrorType(err error, errType interface{}, callback func(string)) bool {
	conditionMet := err == nil || reflect.TypeOf(err) != reflect.TypeOf(errType)
	if conditionMet {
		callback(fmt.Sprintf("error is not of type %T", errType))
	}
	return conditionMet
}

// False will call the callback unless a == false
func False(a bool, callback func(string)) bool {
	if a {
		callback("true != false")
	}
	return a
}

// FileExists will call the callback unless the file exists
func FileExists(path string, callback func(string)) bool {
	conditionMet := false
	info, err := os.Stat(path)
	if err == nil {
		conditionMet = info.IsDir()
	} else {
		conditionMet = os.IsNotExist(err)
	}
	if conditionMet {
		callback(fmt.Sprintf("%s is not a file", path))
	}
	return conditionMet
}

// GreaterThan will call the callback unless a > b
func GreaterThan(a, b int, callback func(string)) bool {
	conditionMet := a <= b
	if conditionMet {
		callback(fmt.Sprintf("%d is not greater than %d", a, b))
	}
	return conditionMet
}

// GreaterThanOrEqual will call the callback unless a >= b
func GreaterThanOrEqual(a, b int, callback func(string)) bool {
	conditionMet := a < b
	if conditionMet {
		callback(fmt.Sprintf("%d is not greater than or equal to %d", a, b))
	}
	return conditionMet
}

// LessThan will call the callback unless a < b
func LessThan(a, b int, callback func(string)) bool {
	conditionMet := a >= b
	if conditionMet {
		callback(fmt.Sprintf("%d is not less than %d", a, b))
	}
	return conditionMet
}

// LessThanOrEqual will call the callback unless a <= b
func LessThanOrEqual(a, b int, callback func(string)) bool {
	conditionMet := a > b
	if conditionMet {
		callback(fmt.Sprintf("%d is not less than or equal to %d", a, b))
	}
	return conditionMet
}

// LongerThan will call the callback unless the length of the string or slice is greater than n
func LongerThan(a interface{}, n int, callback func(string)) bool {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}
	conditionMet := v.Len() <= n
	if conditionMet {
		callback(fmt.Sprintf("%#v is not longer than %d", a, n))
	}
	return conditionMet
}

// Matches will call the callback unless a matches the regex pattern
func Matches(a, pattern string, callback func(string)) bool {
	matched, err := regexp.MatchString(pattern, a)
	conditionMet := err != nil || !matched
	if conditionMet {
		callback(fmt.Sprintf("%#v does not match pattern %#v", a, pattern))
	}
	return conditionMet
}

// Nil will call the callback unless a == nil
func Nil(a interface{}, callback func(string)) bool {
	return Equal(a, nil, callback)
}

// NotAny will call the callback unless no item from sub-set b is in a
func NotAny(a, b interface{}, callback func(string)) bool {
	if !Any(a, b, callback) {
		return false
	}
	callback(fmt.Sprintf("some item from %#v is in %#v", b, a))
	return true
}

// NotContains will call the callback unless a does not contain b (sub-set b is not in a)
func NotContains(a, b interface{}, callback func(string)) bool {
	if !Contains(a, b, callback) {
		return false
	}
	callback(fmt.Sprintf("%#v contains %#v", a, b))
	return true
}

// NotDirExists will call the callback unless the directory does not exist
func NotDirExists(path string, callback func(string)) bool {
	conditionMet := false
	info, err := os.Stat(path)
	if err == nil {
		conditionMet = info.IsDir()
	}
	if conditionMet {
		callback(fmt.Sprintf("%s is a directory", path))
	}
	return conditionMet
}

// NotEmpty will call the callback unless the string or slice is not empty
func NotEmpty(a interface{}, callback func(string)) bool {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}
	conditionMet := v.Len() == 0
	if conditionMet {
		callback("slice or string is empty")
	}
	return conditionMet
}

// NotEqual will call the callback unless a != b
func NotEqual(a, b interface{}, callback func(string)) bool {
	conditionMet := a == b
	if conditionMet {
		callback(fmt.Sprintf("%v == %v", a, b))
	}
	return conditionMet
}

// NotError will call the callback unless err is nil
func NotError(err error, callback func(string)) bool {
	conditionMet := err != nil
	if conditionMet {
		callback(fmt.Sprintf("error occurred: %v", err))
	}
	return conditionMet
}

// NotErrorContains will call the callback unless the error message does not contain the substring
func NotErrorContains(err error, substr string, callback func(string)) bool {
	conditionMet := err != nil && strings.Contains(err.Error(), substr)
	if conditionMet {
		callback(fmt.Sprintf("error contains %s", substr))
	}
	return conditionMet
}

// NotErrorType will call the callback unless the error is not of the given type
func NotErrorType(err error, errType interface{}, callback func(string)) bool {
	conditionMet := err != nil && reflect.TypeOf(err) == reflect.TypeOf(errType)
	if conditionMet {
		callback(fmt.Sprintf("error is of type %T", errType))
	}
	return conditionMet
}

// NotFileExists will call the callback unless the file does not exist
func NotFileExists(path string, callback func(string)) bool {
	conditionMet := false
	info, err := os.Stat(path)
	if err == nil {
		conditionMet = !info.IsDir()
	} else {
		conditionMet = os.IsNotExist(err)
	}
	if conditionMet {
		callback(fmt.Sprintf("%s is a file", path))
	}
	return conditionMet
}

// NotMatches will call the callback unless a does not match the regex pattern
func NotMatches(a, pattern string, callback func(string)) bool {
	matched, err := regexp.MatchString(pattern, a)
	conditionMet := err == nil && matched
	if conditionMet {
		callback(fmt.Sprintf("%#v matches pattern %#v", a, pattern))
	}
	return conditionMet
}

// NotNil will call the callback unless a != nil
func NotNil(a interface{}, callback func(string)) bool {
	return NotEqual(a, nil, callback)
}

// NotPathExists will call the callback if the path exists
func NotPathExists(path string, callback func(string)) bool {
	_, err := os.Stat(path)
	conditionMet := os.IsNotExist(err)
	if conditionMet {
		callback(fmt.Sprintf("%s is not an existing path", path))
	}
	return conditionMet
}

// NotType will call the callback unless a is not of the given type
func NotType(a interface{}, b reflect.Type, callback func(string)) bool {
	conditionMet := reflect.TypeOf(a) == b
	if conditionMet {
		callback(fmt.Sprintf("%#v is of type %T", a, b))
	}
	return conditionMet
}

// NotTimeWithin will call the callback unless the time is not within the duration
func NotTimeWithin(t1, t2 time.Time, d time.Duration, callback func(string)) bool {
	conditionMet := t1.Sub(t2) <= d
	if conditionMet {
		callback(fmt.Sprintf("%v is within %v of %v", t1, d, t2))
	}
	return conditionMet
}

// PathExists will call the callback unless the path exists
func PathExists(path string, callback func(string)) bool {
	_, err := os.Stat(path)
	conditionMet := os.IsNotExist(err)
	if conditionMet {
		callback(fmt.Sprintf("%s does not exist", path))
	}
	return conditionMet
}

// ShorterThan will call the callback unless the length of the string or slice is less than n
func ShorterThan(a interface{}, n int, callback func(string)) bool {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.String && v.Kind() != reflect.Slice {
		callback(fmt.Sprintf("%#v is not a string or slice", a))
		return true
	}
	conditionMet := v.Len() >= n
	if conditionMet {
		callback(fmt.Sprintf("%#v is not shorter than %d", a, n))
	}
	return conditionMet
}

// TimeAfter will call the callback unless t1 is after t2
func TimeAfter(t1, t2 time.Time, callback func(string)) bool {
	conditionMet := t1.Before(t2)
	if conditionMet {
		callback(fmt.Sprintf("%v is not after %v", t1, t2))
	}
	return conditionMet
}

// TimeBefore will call the callback unless t1 is before t2
func TimeBefore(t1, t2 time.Time, callback func(string)) bool {
	conditionMet := t1.After(t2)
	if conditionMet {
		callback(fmt.Sprintf("%v is not before %v", t1, t2))
	}
	return conditionMet
}

// TimeWithin will call the callback unless the time is within the duration
func TimeWithin(t1, t2 time.Time, d time.Duration, callback func(string)) bool {
	conditionMet := t1.Sub(t2) > d
	if conditionMet {
		callback(fmt.Sprintf("%v is not within %v of %v", t1, d, t2))
	}
	return conditionMet
}

// True will call the callback unless a == true
func True(a bool, callback func(string)) bool {
	if !a {
		callback("false != true")
	}
	return !a
}

// Type will call the callback unless a is of the given type
func Type(a interface{}, b reflect.Type, callback func(string)) bool {
	conditionMet := reflect.TypeOf(a) != b
	if conditionMet {
		callback(fmt.Sprintf("%#v is not of type %T", a, b))
	}
	return conditionMet
}

// Zero will call the callback unless a is zero
func Zero(a interface{}, callback func(string)) bool {
	conditionMet := !reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface())
	if conditionMet {
		callback(fmt.Sprintf("%#v is not zero", a))
	}
	return conditionMet
}

// NotZero will call the callback unless a is not zero
func NotZero(a interface{}, callback func(string)) bool {
	conditionMet := reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface())
	if conditionMet {
		callback(fmt.Sprintf("%#v is zero", a))
	}
	return conditionMet
}
