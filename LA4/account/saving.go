package account

import (
	"fmt"
	"la4/customer"
)

// SavingAccount implements the Account interface.
type SavingAccount struct {
	Number   string
	Customer *customer.Customer
	Bal      float64
	Interest float64
}

// NewSavingAccount acts as the constructor.
func NewSavingAccount(number string, customer *customer.Customer, balance float64) *SavingAccount {
	return &SavingAccount{
		Number:   number,
		Customer: customer,
		Bal:      balance,
		Interest: 0,
	}
}

// Accrue calculates interest, updates balance/interest fields, and returns the interest generated.
func (s *SavingAccount) Accrue(rate float64) float64 {
	currentInterest := s.Bal * rate
	s.Interest += currentInterest
	s.Bal += currentInterest
	return currentInterest
}

func (s *SavingAccount) Balance() float64 {
	return s.Bal
}

func (s *SavingAccount) Deposit(amount float64) {
	s.Bal += amount
}

func (s *SavingAccount) Withdraw(amount float64) {
	s.Bal -= amount
}

func (s *SavingAccount) String() string {
	// Formatting to match Java output (e.g. 200.0 instead of 200)
	return fmt.Sprintf("%s:%s:%.1f", s.Number, s.Customer.String(), s.Bal)
}