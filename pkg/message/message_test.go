package message

import (
	"testing"

	"github.com/quiteclose/goatly/pkg/declare/assert"
)

///////////////////////////////////////////////////////////////////////////////

func TestPrefixOneLine(t *testing.T) {
	start := "abc"
	prefix := "#"
	expected := "#abc"
	assert.Equal(t, expected, Prefix(start, prefix), "Prefix should be added to the line.")
}

func TestPrefixMultiLine(t *testing.T) {
	start := "abc\ndef"
	prefix := "#"
	expected := "#abc\n#def"
	assert.Equal(t, expected, Prefix(start, prefix), "Prefix should be added to each line.")
}

func TestUnexpectedText(t *testing.T) {
	noun := "value"
	a := "Expected"
	b := "Found"
	expected := `Unexpected value. Expected:
  ~ Expected
Found:
  ~ Found`
	found := UnexpectedText(noun, a, b)
	assert.Equal(t, expected, found, "Unexpected should format the message correctly.")
}

func TestUnexpectedTextMultiLine(t *testing.T) {
	noun := "value"
	a := "Expected\nLines"
	b := "Found\nLines"
	expected := `Unexpected value. Expected:
  ~ Expected
  ~ Lines
Found:
  ~ Found
  ~ Lines`
	found := UnexpectedText(noun, a, b)
	assert.Equal(t, expected, found, "Unexpected should format the message correctly.")
}

func TestUnexpectedValue(t *testing.T) {
	noun := "value"
	a := "A"
	b := "B"
	expected := `Unexpected value. Expected: "A", Found: "B"`
	found := UnexpectedValue(noun, a, b)
	t.Log(expected)
	t.Log(found)
	assert.Equal(t, expected, found, "UnexpectedValue should format the message correctly.")
}
