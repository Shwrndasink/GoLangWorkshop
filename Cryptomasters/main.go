package main

import (
	"fmt"
	"sync"

	"fem.com/go/crypto/api"
)

func main() {
	var wg sync.WaitGroup // Wait acts as a counter
	currencies := []string{"BTC", "ETH", "BCH"}

	// Loop through slice and getCurrencyData for each currency
	for _, currency := range currencies {
		wg.Add(1)
		go func(c string) {
			getCurrencyData(c)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetCryptoRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
