package unless

import (
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////

// NotPathExists will call the callback if the path exists
func NotPathExists(path string, callback func(string)) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		callback(fmt.Sprintf("%s is not an existing path", path))
	}
}

// NotEqual will call the callback unless a != b
func NotEqual(a, b interface{}, callback func(string)) {
	if a == b {
		callback(fmt.Sprintf("%v == %v", a, b))
	}
}

// NotNil will call the callback unless a != nil
func NotNil(a interface{}, callback func(string)) {
	NotEqual(a, nil, callback)
}
