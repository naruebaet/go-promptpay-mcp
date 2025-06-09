package promptpay

import (
	"fmt"

	"github.com/naruebaet/go-promptpay/pp"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateQRCode(accountType pp.AccountType, accountNumber string, amount *float64) (string, error) {
	if !isValidAccountType(accountType) {
		return "", fmt.Errorf("invalid account type: must be 'phone' or 'id'")
	}

	if accountNumber == "" {
		return "", fmt.Errorf("account number cannot be empty")
	}

	if amount != nil {
		return pp.GenPromptpayWithAmount(accountType, accountNumber, *amount)
	}
	return pp.GenPromptpay(accountType, accountNumber)
}

func isValidAccountType(accountType pp.AccountType) bool {
	return accountType == pp.AccountTypePhone || accountType == pp.AccountTypeID
}
