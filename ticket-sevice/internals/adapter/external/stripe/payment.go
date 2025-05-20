package stripe

import (
	"context"
	"fmt"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type CardPayment struct {
	Amount          int64 // Amount in cents (e.g., 1000 = $10.00)
	Currency        string
	Description     string // Payment description
	PaymentMethodID string
}

func (c *CardPayment) ProcessPayment(ctx context.Context) (PaymentResponse, error) {
	stripe.Key = "sk_test_your_test_secret_key"

	// Validate request
	if c.Amount <= 0 {
		return PaymentResponse{}, fmt.Errorf("amount must be greater than zero")
	}
	if c.Currency == "" {
		return PaymentResponse{}, fmt.Errorf("currency is required")
	}

	// Create Payment Intent
	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(c.Amount),
		Currency:    stripe.String(c.Currency),
		Description: stripe.String(c.Description),
		PaymentMethodTypes: stripe.StringSlice([]string{
			Card.ToString(),
		}),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to create payment intent: %v", err)
	}

	paymentMethodId := "pm_card_visa"

	// Confirm Payment Intent
	confirmParams := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(paymentMethodId),
	}

	pi, err = paymentintent.Confirm(pi.ID, confirmParams)
	if err != nil {
		// Handle specific Stripe errors
		if stripeErr, ok := err.(*stripe.Error); ok {
			switch stripeErr.Code {
			case stripe.ErrorCodeCardDeclined:
				return PaymentResponse{
					TransactionID: pi.ID,
					Status:        "failed",
					Amount:        pi.Amount,
					Currency:      string(pi.Currency),
					Message:       "Card was declined: " + stripeErr.Msg,
				}, nil
			case stripe.ErrorCodeAuthenticationRequired:
				return PaymentResponse{
					TransactionID: pi.ID,
					Status:        "requires_action",
					Amount:        pi.Amount,
					Currency:      string(pi.Currency),
					Message:       "Payment requires authentication (e.g., 3D Secure for card)",
				}, nil
			default:
				return PaymentResponse{}, fmt.Errorf("failed to confirm payment: %v", err)
			}
		}
		return PaymentResponse{}, fmt.Errorf("failed to confirm payment: %v", err)
	}

	// Return successful response
	return PaymentResponse{
		TransactionID: pi.ID,
		Status:        string(pi.Status),
		Amount:        pi.Amount,
		Currency:      string(pi.Currency),
		Message:       fmt.Sprintf("%s payment processed successfully", Card.ToString()),
	}, nil
}
