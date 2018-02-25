package metro2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expected = "0426HEADER9 INNOVIS   EQUIFAX   EXP  TRANSUNION11092017110920171109201711092017ENOVA                                   175 W JACKSON BLVD, CHICAGO, IL 60604                                                           3125684200ENOVA METRO2 GO LIBRARY                 V1             TEST                                                                                                                                              "

func TestHeaderSegment(t *testing.T) {
	h := NewHeaderSegment(426)
	h.RecordDescriptor("426")
	h.CycleIdentifier("9")
	h.InnovisIdentifier("Innovis")
	h.EquifaxIdentifier("equifax")
	h.ExperianIdentifier("exp")
	h.TransUnionIdentifier("transunion")
	h.ActivityDate("11092017")
	h.DateCreated("11092017")
	h.ProgramDate("11092017")
	h.ProgramRevisionDate("11092017")
	h.ReporterName("Enova")
	h.ReporterAddress("175 W Jackson Blvd, Chicago, IL 60604")
	h.ReporterTelephone("3125684200")
	h.SoftwareVendorName("Enova Metro2 Go Library")
	h.SoftwareVersionNumber("v1")
	h.Reserved("test")
	assert.Equal(t, expected, h.String())
}
