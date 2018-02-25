package metro2

import(
    "time"
)

// BaseSegment ...
type BaseSegment struct {
	segment
}

// NewBaseSegment creates and initialize the BaseSegment segment with the given length.
func NewBaseSegment(length int) *BaseSegment {
	b := BaseSegment{newSegment(length)}
	return &b
}

// RecordDescriptor ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 4
// Position   : (0-3)
func (b *BaseSegment) RecordDescriptor(data string) {
	b.writeNumeric(0, 4, data, "Record Descriptor")
}

// ProcessingIndicator ...
// Metro 2 Standard:
// Required   : No
// Max Length : 1
// Position   : (4)
func (b *BaseSegment) ProcessingIndicator(char byte) {
	b.writeByte(4, char)
}

// TimeStamp ...
// Metro 2 Standard:
// Required   : No
// Max Length : 14
// Position   : (5-18)
func (b *BaseSegment) TimeStamp(data string) {
	b.writeNumeric(5, 14, data, "Time Stamp")
}

// CorrectionIndicator ...
// Metro 2 Standardç
// Required   : No
// Max Length : 1
// Position   : (19)
func (b *BaseSegment) CorrectionIndicator(char byte) {
	b.writeByte(19, char)
}

// IdentificationNumber ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 20
// Position   : (20-39)
func (b *BaseSegment) IdentificationNumber(data string) {
  b.writeAlphanumeric(20, 20, data, "Identification Number", alphanumeric)
}

// CycleIdentifier ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 2
// Position   : (40-41)
func (b *BaseSegment) CycleIdentifier(data string) {
  b.writeAlphanumeric(40, 2, data, "Cycle Identifier", alphanumeric)
}

// ConsumerAccountNumber ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 30
// Position   : (42-71)
func (b *BaseSegment) ConsumerAccountNumber(data string) {
  b.writeAlphanumeric(42, 30, data, "Customer Account Number", alphanumeric)
}

// PortfolioType ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 1
// Position   : (72)
func (b *BaseSegment) PortfolioType(char byte) {
  b.writeAlphanumeric(72, 1, data, "Identification Number", alphanumeric)
}

// AccountType ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 2
// Position   : (73-74)
func (b *BaseSegment) AccountType(data string) {
  b.writeAlphanumeric(73, 2, data, "Account Type", alphanumeric)
}

// DateOpened ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 8
// Position   : (75-82)
func (b *BaseSegment) DateOpened(date time.Time) {
  b.writeDate(75, date)
}

// CreditLimit ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 9
// Position   : (83-91)
func (b *BaseSegment) CreditLimit(amount int) {
	b.writeMonetary(83, amount)
}

// OriginalLoanAmount ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 9
// Position   : (92-100)
func (b *BaseSegment) OriginalLoanAmount(amount int) {
	b.writeMonetary(92, amount)
}

// TermsDuration ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 3
// Position   : (101-103)
func (b *BaseSegment) TermsDuration(data string) {
  b.writeAlphanumeric(101, 3, data, "Terms Duation", alphanumeric)
}

// TermsFrequency ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 1
// Position   : (104)
func (b *BaseSegment) TermsFrequency(char byte) {
  b.writeAlphanumeric(104, 1, data, "Terms Frequency", alphanumeric)
}

// ScheduledMonthlyPaymentAmount ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 9
// Position   : (105-113)
func (b *BaseSegment) ScheduledMonthlyPaymentAmount(amount int) {
	b.writeMonetary(105, amount)
}

// ActualPaymentAmount ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 9
// Position   : (114-122)
func (b *BaseSegment) ActualPaymentAmount(amount int) {
	b.writeMonetary(114, amount)
}

// AccountStatus ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 2
// Position   : (123-124)
func (b *BaseSegment) AccountStatus(data string) {
  b.writeAlphanumeric(123, 2, data, "Account Status", alphanumeric)
}

// PaymentRating ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 1
// Position   : (125)
func (b *BaseSegment) PaymentRating(char byte) {
  b.writeAlphanumeric(125, 1, data, "Payment Rating", alphanumeric)
}

// PaymentHistoryProfile ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 24
// Position   : (126-149)
func (b *BaseSegment) PaymentHistoryProfile(data string) {
  b.writeAlphanumeric(126, 24, data, "Payment History Profile", alphanumeric)
}

// SpecialComment ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 2
// Position   : (150-151)
func (b *BaseSegment) SpecialComment(data string) {
  b.writeAlphanumeric(150, 2, data, "Special Comment", alphanumeric)
}
// ComplianceConditionCode ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 2
// Position   : (152-153)
func (b *BaseSegment) ComplianceConditionCode(data string) {
  b.writeAlphanumeric(152, 2, data, "Compliance Condition Code", alphanumeric)
}

// CurrentBalance ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 9
// Position   : (154-162)
func (b *BaseSegment) CurrentBalance(amount int) {
	b.writeMonetary(154, amount)
}

// AmountPastDue ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 9
// Position   : (163-171)
func (b *BaseSegment) AmountPastDue(amount int) {
	b.writeMonetary(163, amount)
}

// OriginalChargeoffAmount ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 9
// Position   : (172-180)
func (b *BaseSegment) OriginalChargeoffAmount(amount int) {
	b.writeMonetary(172, amount)
}

// DateOfAccountInformation ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 8
// Position   : (181-188)
func (b *BaseSegment) DateOfAccountInformation(date time.Time) {
  b.writeDate(181, date)
}

// DateOfFirstDelinquency ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 8
// Position   : (189-196)
func (b *BaseSegment) DateOfFirstDelinquency(date time.Time) {
  b.writeDate(189, date)
}

// DateClosed ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 8
// Position   : (197-204)
func (b *BaseSegment) DateClosed(date time.Time) {
  b.writeDate(197, date)
}

// DateOfLastPayment ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 8
// Position   : (205-212)
func (b *BaseSegment) DateOfLastPayment(data time.Time) {
  b.writeDate(205, date)
}

// InterestTypeIndicator ...
// Metro 2 Standard:
// Required   : No
// Max Length : 1
// Position   : (213)
func (b *BaseSegment) InterestTypeIndicator(char byte) {
	b.writeByte(213, char)
}

// Reserved ...
// Metro 2 Standard:
// Required   : No
// Max Length : 16
// Position   : (214-229)
func (b *BaseSegment) Reserved(data string) {
  b.writeAlphanumeric(214, 16, data, "Reserved", alphanumeric)
}

// ConsumerTransactionType ...
// Metro 2 Standard:
// Required   : No
// Max Length : 1
// Position   : (230)
func (b *BaseSegment) ConsumerTransactionType(char byte) {
	b.writeByte(230, char)
}

// Surname ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 25
// Position   : (231-255)
func (b *BaseSegment) Surname(data string) {
  b.writeAlphanumeric(231, 25, data, "Surname", alphanumericPlusDash)
}

// FirstName ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 20
// Position   : (256-275)
func (b *BaseSegment) FirstName(data string) {
  b.writeAlphanumeric(256, 20, data, "FirstName", alphanumericPlusDash)
}

// MiddleName ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 20
// Position   : (276-295)
func (b *BaseSegment) MiddleName(data string) {
  b.writeAlphanumeric(276, 20, data, "MiddleName", alphanumericPlusDash)
}

// GenerationCode ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 1
// Position   : (296)
func (b *BaseSegment) GenerationCode(char byte) {
	b.writeByte(296, char)
}

// SocialSecurityNumber ...
// Metro 2 Standard:
// Required   : Highly Recommended
// Max Length : 9
// Position   : (297-305)
func (b *BaseSegment) SocialSecurityNumber(data string) {
	b.writeNumeric(297, 9, data, "Social Security Number")
}

// DateOfBirth ...
// Metro 2 Standard:
// Required   : Highly Recommended
// Max Length : 8
// Position   : (306-313)
func (b *BaseSegment) DateOfBirth(date time.Time) {
  b.writeDate(306, date)
}

// TelephoneNumber ...
// Metro 2 Standard:
// Required   : No
// Max Length : 10
// Position   : (314-323)
func (b *BaseSegment) TelephoneNumber(data string) {
	b.writeNumeric(314, 10, data, "Telephone Number")
}

// ECOACode ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 1
// Position   : (324)
func (b *BaseSegment) ECOACode(char byte) {
	b.writeByte(324, char)
}

// ConsumerInformationIndicator ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 2
// Position   : (325-326)
func (b *BaseSegment) ConsumerInformationIndicator(data string) {
  b.writeAlphanumeric(325, 2, data, "Customer Information Indicator", alphanumeric)
}

// CountryCode ...
// Metro 2 Standard:
// Required   : No
// Max Length : 2
// Position   : (327-328)
func (b *BaseSegment) CountryCode(data string) {
  b.writeAlphanumeric(327, 2, data, "Country Code", alphanumeric)
}

// FirstLineofAddress ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 32
// Position   : (329-360)
func (b *BaseSegment) FirstLineofAddress(data string) {
  b.writeAlphanumeric(329, 32, data, "First Linee of Address", alphanumericPlusDotDashSlash)
}

// SecondLineofAddress ...
// Metro 2 Standard:
// Required   : Field is required when applicable to the account being reported
// Max Length : 32
// Position   : (361-392)
func (b *BaseSegment) SecondLineofAddress(data string) {
  b.writeAlphanumeric(361, 32, data, "Second Linee of Address", alphanumericPlusDotDashSlash)
}

// City ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 20
// Position   : (393-412)
func (b *BaseSegment) City(data string) {
  b.writeAlphanumeric(393, 20, data, "City", alphanumericPlusDotDashSlash)
}

// State ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 2
// Position   : (413-414)
func (b *BaseSegment) State(data string) {
  b.writeAlphanumeric(413, 2, data, "State", alphanumeric)
}

// ZipCode ...
// Metro 2 Standard:
// Required   : Field is always required
// Max Length : 9
// Position   : (415-423)
func (b *BaseSegment) ZipCode(data string) {
  b.writeAlphanumeric(415, 9, data, "State", alphanumeric)
}

// AddressIndicator ...
// Metro 2 Standard:
// Required   : No
// Max Length : 1
// Position   : (424)
func (b *BaseSegment) AddressIndicator(char byte) {
	b.writeByte(424, char)
}

// ResidenceCode ...
// Metro 2 Standard:
// Required   : No
// Max Length : 1
// Position   : (425)
func (b *BaseSegment) ResidenceCode(char byte) {
	b.writeByte(425, char)
}
