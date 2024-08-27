package api_test

import (
	"testing"

	"fem.com/go/crypto/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetCryptoRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
