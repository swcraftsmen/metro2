package metro2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
  "time"
)

type segment struct {
	record []byte
	Errors []string
}

const (
	nines       = "999999999"
	zero        = '0'
	initiByte   = ' '
	oneBillion = 10000000000
)

var (
	alphabetic                   = regexp.MustCompile("^([[:alpha:]])+$")
	alphanumeric                 = regexp.MustCompile("^([[:alnum:]]|\\s)+$")
	alphanumericPlusDash         = regexp.MustCompile("^([[:alnum:]]|\\s|\\-)+$")
	alphanumericPlusDotDashSlash = regexp.MustCompile("^([[:alnum:]]|\\s|\\-|\\.|\\|\\/|\\,)+$")
	numberic                     = regexp.MustCompile("^\\d+\\.?\\d*$")
)

func newSegment(size int) segment {
	segment := segment{record: make([]byte, size)}
	segment.init()
	return segment
}

func (s *segment) init() {
	length := len(s.record)
	for i := 0; i < length; i++ {
		s.record[i] = initiByte
	}
}

// The total capacity of the segment record.
func (s *segment) Length() int {
	return cap(s.record)
}

//  Metro 2 Reporting Sstandard : Alphanumeric Field
//    - Every alphanumeric field is left-justified and blank filled.
//    - Every alpha field should be upper case letters.
func (s *segment) writeAlphanumeric(offset int, maxLen int, data string, name string, allowChars *regexp.Regexp) {
	lenCheck := len(data) <= maxLen
	charCheck := allowChars.MatchString(data)

	if lenCheck && charCheck {
		s.writeString(offset, strings.ToUpper(data))
	} else {
		if lenCheck == charCheck {
			m1 := fmt.Sprintf("The maximum length for the field [%s] is %d"+
				" ,and the input length is %d. [%s]", name, maxLen, len(data), data)

			m2 := fmt.Sprintf("The field [%s] contains not allowed character. [%s]", name, data)
			s.Errors = append(s.Errors, m1, m2)
		} else if lenCheck == false {
			msg := fmt.Sprintf("The maximum length for the field [%s] is %d"+
				" ,and the input length is %d. [%s]", name, maxLen, len(data), data)
			s.Errors = append(s.Errors, msg)
		} else {
			msg := fmt.Sprintf("The field [%s] contains not allowed character. [%s]", name, data)
			s.Errors = append(s.Errors, msg)
		}
	}
}

//  Metro 2 Reporting Standard : Numeric Field
//    - Every numeric field is right-justified and zero filled.
//    - If a numeric field is not available, it should be zero filled.
func (s *segment) writeNumeric(offset int, maxLen int, data string, name string) {
	if !numberic.MatchString(data) {
		msg := fmt.Sprintf("The field [%s] contains none numeric character. [%s]", name, data)
		s.Errors = append(s.Errors, msg)
		return
	}
	if len(data) < maxLen {
		// Since the numeric require to fill zero to the unused bytes,
		// we fill out with zero first and then fill out with data
		l := (maxLen - len(data)) + offset
		for i := offset; i < l; i++ {
			s.writeByte(i, zero)
		}
		s.writeString(l, data)
	} else {
		s.writeString(offset, data)
	}

}

//  Metro 2 Reporting Standard : Monetary Field
//    - Monetary fields are reported in whole dollars only. Cents
//      should be truncated.
//    - If a monetary field is not applicable, it should be zero
//      filled. Do not 9-fill these fields. A monetary field should
//      be 9-filled when the amount is in excess of $1 billion.
func (s *segment) writeMonetary(offset int, amount int) {
	if amount >= oneBillion {
		s.writeString(offset, nines)
	} else {
		sAmount := strconv.Itoa(amount)
    s.writeNumeric(offset, 9, sAmount, "Monetary Field")
	}
}

// Metro 2 Reporting Standard : Date Format
//   - Date record as numeric field but the date format be MMddyyy
func (s *segment) writeDate(offset int, date time.Time){
    sDate := fmt.Sprintf("%02d%02d%d", int(date.Month()), date.Day(), date.Year())
    s.writeString(offset, sDate)
}

// Add the field data to desired position
func (s *segment) writeString(offset int, data string) {
	copy(s.record[offset:], data)
}

func (s *segment) writeByte(p int, b byte) {
	if s.Length() > p {
		s.record[p] = b
	}
}

// String converts whole character set into the string format
func (s *segment) String() string {
	return string(s.record[:])
}
