package metro2

import (
	"fmt"
	"strings"
)

// Metro2 ...
type Metro2 struct {
	HeaderSegment
}

// CorrectionIndicatorLookup ...
// Metro 2 Corrrection indicator look up
// Ex.
// metro2.CorrectionIndicatorLookup("Normal") => 0
// metro2.CorrectionIndicatorLookup("Correction") => 1
// Return NULL byte and error if the indicator is not found.
func CorrectionIndicatorLookup(ci string) (byte, error) {
	switch strings.ToUpper(ci) {
	case "NORMAL":
		return byte('0'), nil
	case "CORRECTION":
		return byte('1'), nil
	default:
		return byte('\x00'), fmt.Errorf("Undefined Correction Indicator. [%s]", ci)
	}
}

// PortoflioTypeLookup ...
// Metro 2 Portfolio Type Look up
// Ex.
// metro2.PortoflioTypeLookup("intallment") => "I"
// Return NULL byte and error if the portfolio type is not found
func PortoflioTypeLookup(pType string) (byte, error) {
	switch strings.ToLower(pType) {
	case "line_of_credit":
		return byte('C'), nil
	case "installment":
		return byte('I'), nil
	case "mortgage":
		return byte('M'), nil
	case "open":
		return byte('O'), nil
	case "revolving":
		return byte('R'), nil
	default:
		return byte('\x00'), fmt.Errorf("Undefined Portoflio Type. [%s]", pType)
	}
}

// TermsFrequenccyLookup ...
// Metro 2 Terms Frequency Look up
// Ex.
// metro2.TermsFrequenccyLookup("weekly") => "W"
// Return NULL byte and error if the portfolio type is not found
func TermsFrequenccyLookup(frequency string) (byte, error) {
	switch strings.ToLower(frequency) {
	case "defferred":
		return byte('D'), nil
	case "single_payment":
		return byte('P'), nil
	case "weeklly":
		return byte('W'), nil
	case "biweekly":
		return byte('B'), nil
	case "semimonthly":
		return byte('E'), nil
	case "monthly":
		return byte('M'), nil
	case "bimonthly":
		return byte('L'), nil
	case "quarterly":
		return byte('Q'), nil
	case "tri_annually":
		return byte('T'), nil
	case "semiannually":
		return byte('S'), nil
	case "annually":
		return byte('Y'), nil
	default:
		return byte('\x00'), fmt.Errorf("Undefined Terms Frequency . [%s]", frequency)
	}
}
