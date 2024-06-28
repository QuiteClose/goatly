package unless

import "fmt"

///////////////////////////////////////////////////////////////////////////////

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
