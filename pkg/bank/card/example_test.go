package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:  "5058_xxxx_xxxx_8888",
			Active: true,
			Balance: 10_000,
		},
		{
			PAN:  "5058_xxxx_xxxx_4444",
			Active: true,
			Balance: 15_000,
		},
		{
			PAN:  "5058_xxxx_xxxx_3333",
			Active: false,
			Balance: 15_000,
		},
		{
			PAN:  "5058_xxxx_xxxx_2222",
			Active: true,
			Balance: 0,
		},
	}
	paymentSources := PaymentSources(cards)

	for _, paymentSource := range paymentSources {
		fmt.Println(paymentSource.Number)
	}
	// Output:
	//5058_xxxx_xxxx_8888
	//5058_xxxx_xxxx_4444
}