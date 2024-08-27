package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"fem.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetCryptoRate(currency string) (*datatypes.CryptoRate, error) {
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	if err != nil {
		fmt.Print("There was error")
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.CryptoRate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
