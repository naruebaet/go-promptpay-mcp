package types

import "github.com/naruebaet/go-promptpay/pp"

type GeneratePromptPayRequest struct {
	AccountType   pp.AccountType `json:"accountType"`
	AccountNumber string         `json:"accountNumber"`
	Amount        *float64       `json:"amount,omitempty"`
}

type GeneratePromptPayResponse struct {
	QRCode string `json:"qrCode"`
}
