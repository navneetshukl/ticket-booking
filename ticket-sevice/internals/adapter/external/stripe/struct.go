package stripe

import "context"

type PaymentResponse struct {
	TransactionID string
	Status        string
	Amount        int64
	Currency      string
	Message       string
}

type PaymentMethodType string

const (
	Card        PaymentMethodType = "card"
	BankAccount PaymentMethodType = "us_bank_account"
	UPI         PaymentMethodType = "upi"
)

func (p PaymentMethodType) ToString() string {
	return string(p)
}

type StripeSvc interface {
	ProcessPayment(ctx context.Context) error
}
