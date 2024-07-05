package message

import (
	"testing"
)

///////////////////////////////////////////////////////////////////////////////

func TestPrefixOneLine(t *testing.T) {
	start := "abc"
	prefix := "#"
	expected := "#abc"
	found := Prefix(start, prefix)
	if found != expected {
		t.Errorf("Prefix should be added to the line: %q != %q", expected, found)
	}
}

func TestPrefixMultiLine(t *testing.T) {
	start := "abc\ndef"
	prefix := "#"
	expected := "#abc\n#def"
	found := Prefix(start, prefix)
	if found != expected {
		t.Errorf("Prefix should be added to each line: %q != %q", expected, found)
	}
}

func TestMustContain(t *testing.T) {
	noun := "value"
	a := "Long string of text."
	b := "Sub-String"
	expected := `value must contain specific text.
Must contain:
  ~ Sub-String
Found:
  ~ Long string of text.`
	found := MustContain(noun, a, b)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}

func TestMustContainMultiLine(t *testing.T) {
	noun := "value"
	a := "Long string\nof many\nLines"
	b := "Sub\nString"
	expected := `value must contain specific text.
Must contain:
  ~ Sub
  ~ String
Found:
  ~ Long string
  ~ of many
  ~ Lines`
	found := MustContain(noun, a, b)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}

func TestMustNotContain(t *testing.T) {
	noun := "value"
	a := "A"
	b := "B"
	expected := `value must not contain specific text.
Must not contain:
  ~ B
Found:
  ~ A`
	found := MustNotContain(noun, a, b)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}

func TestMustNotContainMultiLine(t *testing.T) {
	noun := "value"
	a := "Long string\nof many\nLines"
	b := "Sub\nString"
	expected := `value must not contain specific text.
Must not contain:
  ~ Sub
  ~ String
Found:
  ~ Long string
  ~ of many
  ~ Lines`
	found := MustNotContain(noun, a, b)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}

func TestMustEqual(t *testing.T) {
	noun := "value"
	a := "A"
	b := "B"
	expected := `value must equal "A" but found "B"`
	found := MustEqual(noun, a, b)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}

func TestMustNotEqual(t *testing.T) {
	noun := "value"
	a := "A"
	expected := `value must not equal "A"`
	found := MustNotEqual(noun, a)
	if found != expected {
		t.Errorf("Incorrect message:\nExpected:\n%s\nFound:\n%s", Prefix(expected, "  ~ "), Prefix(found, "  ~ "))
	}
}
