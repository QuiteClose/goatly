package message

import (
	"fmt"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

// Prefix indents each line of the given s with the given prefix.
func Prefix(s, prefix string) string {
	return prefix + strings.ReplaceAll(s, "\n", "\n"+prefix)
}

// UnexpectedText assumes that the expected and found values are multi-line
// text and formats a message with indented values for comparison.
func UnexpectedText(noun, expected, found string) string {
	template := "Unexpected %s.\nExpected:\n%s\nFound:\n%s"
	return fmt.Sprintf(
		template,
		noun,
		Prefix(expected, "  ~ "),
		Prefix(found, "  ~ "),
	)
}

// UnexpectedValue assumes that the expected and found values are single-line
// values and formats a message with the values for comparison.
func UnexpectedValue(noun string, expected, found interface{}) string {
	template := "Unexpected %s. Expected: %#v, Found: %#v"
	return fmt.Sprintf(template, noun, expected, found)
}
