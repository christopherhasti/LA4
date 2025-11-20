package account

import (
	"fmt"
	"la4/customer"
)

// CheckingAccount implements the Account interface.
type CheckingAccount struct {
	Number   string
	Customer *customer.Customer
	Bal      float64
}

// NewCheckingAccount acts as the constructor.
func NewCheckingAccount(number string, customer *customer.Customer, balance float64) *CheckingAccount {
	return &CheckingAccount{
		Number:   number,
		Customer: customer,
		Bal:      balance,
	}
}

// Accrue for CheckingAccount does nothing and returns 0 interest.
func (c *CheckingAccount) Accrue(rate float64) float64 {
	return 0
}

func (c *CheckingAccount) Balance() float64 {
	return c.Bal
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.Bal += amount
}

func (c *CheckingAccount) Withdraw(amount float64) {
	c.Bal -= amount
}

func (c *CheckingAccount) String() string {
	return fmt.Sprintf("%s:%s:%.1f", c.Number, c.Customer.String(), c.Bal)
}