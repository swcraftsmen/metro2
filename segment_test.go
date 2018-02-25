package metro2

import (
	"github.com/stretchr/testify/assert"
	"testing"
  "time"
)

func TestNewSegment(t *testing.T) {
	s := newSegment(5)
	assert.Equal(t, s.Length(), 5)
}

func TestWriteAlphanumericWithValidInput(t *testing.T) {
	s := newSegment(5)
	s.writeAlphanumeric(0, 5, "abc12", "test", alphanumeric)
	assert.Equal(t, "ABC12", s.String())
}
func TestWriteAlphanumericLeftJustified(t *testing.T) {
	s := newSegment(5)
	s.writeAlphanumeric(0, 5, "abc", "test", alphanumeric)
	assert.Equal(t, "ABC  ", s.String())
}

func TestWriteAlphanumericWithNotAllowedChar(t *testing.T) {
	s := newSegment(5)
	s.writeAlphanumeric(0, 5, "abcd=", "test", alphanumeric)
	assert.Equal(t, 1, len(s.Errors))
	assert.Equal(t, "The field [test] contains not allowed character. [abcd=]", s.Errors[0])
}

func TestWriteNumericWithInvalidInput(t *testing.T) {
	s := newSegment(5)
	s.writeNumeric(0, 5, "12ab", "test")
	assert.Equal(t, "The field [test] contains none numeric character. [12ab]", s.Errors[0])
}

func TestWriteNumericWithValidInput(t *testing.T) {
	s := newSegment(5)
	s.writeNumeric(0, 5, "12", "test")
	assert.Equal(t, "00012", s.String())
}

func TestWriteMonetaryExcessOfBillion(t *testing.T) {
	s := newSegment(9)
	s.writeMonetary(0, 10000000000)
	assert.Equal(t, "999999999", s.String())
}

func TestWriteMonetaryLessThanBillion(t *testing.T) {
	s := newSegment(9)
	s.writeMonetary(0,100000)
	assert.Equal(t, "000100000", s.String())
}

func TestDateField(t *testing.T) {
  s := newSegment(8)
  date, _ := time.Parse(time.RFC3339, "2017-11-01T22:08:41+00:00")
  s.writeDate(0,date)
	assert.Equal(t, "11012017", s.String())
}
