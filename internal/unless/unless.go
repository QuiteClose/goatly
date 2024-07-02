package unless

import (
	"fmt"
	"os"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

// Contains will call the callback unless a contains b
func Contains(a, b string, callback func(string)) {
	if !strings.Contains(a, b) {
		callback(fmt.Sprintf("%#v does not contain %#v", a, b))
	}
}

// DirExists will call the callback unless the directory exists
func DirExists(path string, callback func(string)) {
	info, err := os.Stat(path)
	if err == nil {
		if !info.IsDir() {
			callback(fmt.Sprintf("%s is not a directory", path))
		}
	} else {
		if os.IsNotExist(err) {
			callback(fmt.Sprintf("%s is not an existing path", path))
		}
	}
}

// Equal will call the callback unless a == b
func Equal(a, b interface{}, callback func(string)) {
	if a != b {
		callback(fmt.Sprintf("%#v != %#v", a, b))
	}
}

// False will call the callback unless a == false
func False(a bool, callback func(string)) {
	if a {
		callback("true != false")
	}
}

// Nil will call the callback unless a == nil
func Nil(a interface{}, callback func(string)) {
	Equal(a, nil, callback)
}

// True will call the callback unless a == true
func True(a bool, callback func(string)) {
	if !a {
		callback("false != true")
	}
}
