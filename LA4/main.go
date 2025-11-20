package main

import (
	"fmt"
	"la4/account"
	"la4/customer"
)

// Bank manages a collection of accounts.
type Bank struct {
	// Using a map to represent a Set of accounts as hinted.
	// Key and Value are both the interface.
	accounts map[account.Account]account.Account
}

// NewBank creates a new Bank.
func NewBank() *Bank {
	return &Bank{
		accounts: make(map[account.Account]account.Account),
	}
}

// Add adds an account to the bank.
func (b *Bank) Add(a account.Account) {
	b.accounts[a] = a
}

// Accrue applies interest to all accounts concurrently.
// It uses a channel to sum the total interest generated.
func (b *Bank) Accrue(rate float64) {
	// Create a channel to receive interest amounts
	ch := make(chan float64)
	
	// Count accounts to know how many results to expect
	count := len(b.accounts)
	
	// Launch a goroutine for each account
	for _, a := range b.accounts {
		// Capture 'a' in the loop scope or pass as argument
		go func(acc account.Account) {
			ch <- acc.Accrue(rate)
		}(a)
	}

	// Sum the interest from the channel
	var totalInterest float64
	for i := 0; i < count; i++ {
		totalInterest += <-ch
	}

	// Print the total as requested by the assignment
	fmt.Printf("Total Interest: %.2f\n", totalInterest)
}

// String returns the string representation of the bank (all accounts).
func (b *Bank) String() string {
	result := ""
	for _, a := range b.accounts {
		result += a.String() + "\n"
	}
	return result
}

func main() {
	bank := NewBank()
	c := customer.New("Ann")

	// Create and add accounts
	// Note: Java output shows specific string formats, we try to match them.
	bank.Add(account.NewCheckingAccount("01001", c, 100.00))
	bank.Add(account.NewSavingAccount("01002", c, 200.00))

	// Accrue interest
	bank.Accrue(0.02)

	// Print bank status
	fmt.Print(bank.String())
}