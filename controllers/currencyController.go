package controllers

import "errors"

var (
	Currencies = map[int]string{
		1: "USD", 2: "NGN", 3: "GBP", 4: "YUAN",
	}

	ExchangeRates = map[int]int{
		1: 1, 2: 2, 3: 3, 4: 3,
	}
)

func getCurrencyName(id int) string {
	return Currencies[id]
}

func getCurrencyID(name string) (error, int) {
	for i := 0; i < len(Currencies); i++ {
		if Currencies[i] == name {
			return nil, i
		}
	}
	return errors.New("currency doesn't exit"), 0
}
