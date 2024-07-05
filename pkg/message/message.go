package message

import (
	"fmt"
	"regexp"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

// Prefix indents each line of the given s with the given prefix.
func Prefix(s, prefix string) string {
	return prefix + strings.ReplaceAll(s, "\n", "\n"+prefix)
}

// MustContain assumes that the expected and found values are multi-line
// text and formats a message with indented values for comparison.
func MustContain(noun, s, substring string) string {
	template := "%s must contain specific text.\nMust contain:\n%s\nFound:\n%s"
	return fmt.Sprintf(
		template,
		noun,
		Prefix(substring, "  ~ "),
		Prefix(s, "  ~ "),
	)
}

// MustEqual assumes that the expected and found values are single-line
// values and formats a message with the values for comparison.
func MustEqual(noun string, expected, found interface{}) string {
	template := "%s must equal %#v but found %#v"
	return fmt.Sprintf(template, noun, expected, found)
}

// MustMatch is like MustContain but for regular expressions.
func MustMatch(noun string, pattern regexp.Regexp, found string) string {
	template := "%s must match pattern.\nMust match: %q\nFound:\n%s"
	return fmt.Sprintf(
		template,
		noun,
		pattern.String(),
		Prefix(found, "  ~ "),
	)
}

// MustNotContain assumes that the expected and found values are multi-line
// text and formats a message with indented values for comparison.
func MustNotContain(noun, s, substring string) string {
	template := "%s must not contain specific text.\nMust not contain:\n%s\nFound:\n%s"
	return fmt.Sprintf(
		template,
		noun,
		Prefix(substring, "  ~ "),
		Prefix(s, "  ~ "),
	)
}

// MustNotEqual assumes that the expected and found values are single-line
// values and formats a message with the values for comparison.
func MustNotEqual(noun string, found interface{}) string {
	template := "%s must not equal %#v"
	return fmt.Sprintf(template, noun, found)
}

// MustMatch is like MustNotContain but for regular expressions.
func MustNotMatch(noun string, pattern regexp.Regexp, found string) string {
	template := "%s must not match pattern.\nMust not match: %q\nFound:\n%s"
	return fmt.Sprintf(
		template,
		noun,
		pattern.String(),
		Prefix(found, "  ~ "),
	)
}
