package metro2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseSegment(t *testing.T){
    b := NewBaseSegment(426)
    b.RecordDescriptor("426")
    b.ProcessingIndicator(0)
    b.TimeStamp("02242018204221")
    b.CorrectionIndicator(0)
    b.IdentificationNumber("12345678901234567890")
    b.ConsumerAccountNumber("123456789")
    b.PortfolioType("I")
    b.AccountType("01")
    b.DateOpened("02232018")
    b.CreditLimit(0)
    b.OriginalLoanAmount(1000000000)

}
