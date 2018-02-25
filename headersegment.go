package metro2

// HeaderSegment ...
type HeaderSegment struct {
	segment
}

// NewHeaderSegment creates and initialize the Base segment with the given length.
func NewHeaderSegment(length int) *HeaderSegment {
	h := HeaderSegment{newSegment(length)}
	h.RecordIdentifier("HEADER")
	return &h
}

// RecordDescriptor ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 4
// Position     : (0-3)
func (h *HeaderSegment) RecordDescriptor(data string) {
	h.writeNumeric(0, 4, data, "Record Descriptor")
}

// RecordIdentifier ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 6
// Position     : (4-9)
func (h *HeaderSegment) RecordIdentifier(data string) {
	h.writeAlphanumeric(4, 9, data, "Record Identitifer", alphanumeric)
}

// CycleIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 2
// Position     : (10-11)
func (h *HeaderSegment) CycleIdentifier(data string) {
	h.writeAlphanumeric(10, 2, data, "Cycle Identitifer", alphanumeric)
}

// InnovisIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 10
// Position     : (12-21)
func (h *HeaderSegment) InnovisIdentifier(data string) {
	h.writeAlphanumeric(12, 10, data, "Innovis Identifier", alphanumeric)
}

// EquifaxIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 10
// Position     : (22-31)
func (h *HeaderSegment) EquifaxIdentifier(data string) {
	h.writeAlphanumeric(22, 10, data, "Equifax Identifier", alphanumeric)
}

// ExperianIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 5
// Position     : (32-36)
func (h *HeaderSegment) ExperianIdentifier(data string) {
	h.writeAlphanumeric(32, 5, data, "Experian Identifier", alphanumeric)
}

// TransUnionIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 10
// Position     : (37-46)
func (h *HeaderSegment) TransUnionIdentifier(data string) {
	h.writeAlphanumeric(37, 10, data, "TransUnion Identifier", alphanumeric)
}

// ActivityDate ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 8
// Position     : (47-54)
func (h *HeaderSegment) ActivityDate(data string) {
	h.writeNumeric(47, 8, data, "Activity Date")
}

// DateCreated ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 8
// Position     : (55-62)
func (h *HeaderSegment) DateCreated(data string) {
	h.writeNumeric(55, 8, data, "Date Created")
}

// ProgramDate ...
// Metro 2 standard:
// Required     : No
// Max Length   : 8
// Position     : (63-70)
func (h *HeaderSegment) ProgramDate(data string) {
	h.writeNumeric(63, 8, data, "Program Date")
}

// ProgramRevisionDate ...
// Metro 2 standard:
// Required     : No
// Max Length   : 8
// Position     : (71-78)
func (h *HeaderSegment) ProgramRevisionDate(data string) {
	h.writeNumeric(71, 8, data, "Program Revision Date")
}

// ReporterName ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 40
// Position     : (79-118)
func (h *HeaderSegment) ReporterName(data string) {
	h.writeAlphanumeric(79, 40, data, "Reporter Name", alphanumeric)
}

// ReporterAddress ...
// Metro 2 standard:
// Required     : Field is always required
// Max Length   : 96
// Position     : (119-214)
func (h *HeaderSegment) ReporterAddress(data string) {
	h.writeAlphanumeric(119, 96, data, "Reporter Address", alphanumericPlusDotDashSlash)
}

// ReporterTelephone ...
// Metro 2 standard:
// Required     : No
// Max Length   : 10
// Position     : (215-224)
func (h *HeaderSegment) ReporterTelephone(data string) {
	h.writeNumeric(215, 10, data, "Reporter Telephone")
}

// SoftwareVendorName ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 40
// Position     : (225-264)
func (h *HeaderSegment) SoftwareVendorName(data string) {
	h.writeAlphanumeric(225, 40, data, "Software Vendor Name", alphanumeric)
}

// SoftwareVersionNumber ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 5
// Position     : (265-269)
func (h *HeaderSegment) SoftwareVersionNumber(data string) {
	h.writeAlphanumeric(265, 5, data, "Reporter Name", alphanumeric)
}

// MPProgramIdentifier ...
// Metro 2 standard:
// Required     : Field is required when applicable to the file being reported
// Max Length   : 10
// Position     : (270-279)
func (h *HeaderSegment) MPProgramIdentifier(data string) {
	h.writeAlphanumeric(270, 10, data, "MicroBilt/PRBC Program Identifier", alphanumeric)
}

// Reserved ...
// Metro 2 standard:
// Required     : No
// Max Length   : 146
// Position     : (280-425)
func (h *HeaderSegment) Reserved(data string) {
	h.writeAlphanumeric(280, 146, data, "Reversedd", alphanumeric)
}
