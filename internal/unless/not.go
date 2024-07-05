package unless

import (
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////

// NotPathExists will call the callback if the path exists
func NotPathExists(path string, callback func(string)) bool {
	_, err := os.Stat(path)
	conditionMet := os.IsNotExist(err)
	if conditionMet {
		callback(fmt.Sprintf("%s is not an existing path", path))
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

// NotNil will call the callback unless a != nil
func NotNil(a interface{}, callback func(string)) bool {
	return NotEqual(a, nil, callback)
}
