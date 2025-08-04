package services

import "fmt"

type BankAdapter interface {
	Auth(clientID, clientSecret string) error
	Charge(transactionID string, amount float64) error
	Refund(transactionID string, amount float64) error
}

type MockBank struct {
	Authorized bool
}

func (m *MockBank) Auth(clientID, clientSecret string) error {
	if clientID == "" || clientSecret == "" {
		return fmt.Errorf("auth failed: missing credentials")
	}
	m.Authorized = true
	return nil
}

func (m *MockBank) Charge(transactionID string, amount float64) error {
	if !m.Authorized {
		return fmt.Errorf("charge failed: not authorized")
	}
	fmt.Printf("MockBank: Charged %.2f for transaction %s\n", amount, transactionID)
	return nil
}

func (m *MockBank) Refund(transactionID string, amount float64) error {
	if !m.Authorized {
		return fmt.Errorf("refund failed: not authorized")
	}
	fmt.Printf("MockBank: Refunded %.2f for transaction %s\n", amount, transactionID)
	return nil
}
