package unless

import (
	"fmt"
	"os"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

// Contains will call the callback unless a contains b
func Contains(a, b string, callback func(string)) bool {
	conditionMet := !strings.Contains(a, b)
	if conditionMet {
		callback(fmt.Sprintf("%#v does not contain %#v", a, b))
	}
	return conditionMet
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

// Equal will call the callback unless a == b
func Equal(a, b interface{}, callback func(string)) bool {
	conditionMet := a != b
	if conditionMet {
		callback(fmt.Sprintf("%#v != %#v", a, b))
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

// Nil will call the callback unless a == nil
func Nil(a interface{}, callback func(string)) bool {
	return Equal(a, nil, callback)
}

// True will call the callback unless a == true
func True(a bool, callback func(string)) bool {
	if !a {
		callback("false != true")
	}
	return !a
}
