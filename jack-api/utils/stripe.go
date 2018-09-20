package utils

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

var StripKey = "sk_live_0FeoUv8ixC4nzyUndaRqyEC8"

func ChargeCustomer(amount int, currency stripe.Currency, customerId string, description string) (*stripe.Charge, error)  {
	stripe.Key = StripKey

	ammount := int64(50)
	//ammount := int64(amount * 100)

	// Charge the Customer instead of the card:
	chargeParams := &stripe.ChargeParams{
		Amount: &ammount,
		Currency: stripe.String(string(currency)),
		Customer: &customerId,
		Description: &description,
	}

	return charge.New(chargeParams)
}
