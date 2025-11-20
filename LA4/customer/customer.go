package customer

// Customer represents a bank customer.
type Customer struct {
	Name string
}

// New creates a new Customer.
func New(name string) *Customer {
	return &Customer{Name: name}
}

// String returns the string representation of the customer.
func (c *Customer) String() string {
	return c.Name
}