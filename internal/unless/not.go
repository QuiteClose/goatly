package unless

import "fmt"

///////////////////////////////////////////////////////////////////////////////

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

