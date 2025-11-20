package account

// Account is the interface approximating the abstract class Account.
// It declares methods that all specific account types must implement.
type Account interface {
	Accrue(rate float64) float64
	Balance() float64
	Deposit(amount float64)
	Withdraw(amount float64)
	String() string
}