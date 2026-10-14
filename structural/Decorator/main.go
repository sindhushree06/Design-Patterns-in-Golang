//skip.New() - that is usually written in protected api to skip auth in case headers present in request, but we want to add some extra functionality to it, we can use decorator pattern to add extra functionality to the existing code without modifying it

// add new behaviour to the existing object without modifying the original code
package main

import "fmt"

// Component
type Payment interface {
	Pay(amount float64)
}

// Original object
type BasicPayment struct{}

func (b *BasicPayment) Pay(amount float64) {
	fmt.Println("Processing payment:", amount)
}

// Logging Decorator
type LoggingPayment struct {
	payment Payment
}

func (l *LoggingPayment) Pay(amount float64) {
	fmt.Println("Logging: payment started")

	l.payment.Pay(amount)

	fmt.Println("Logging: payment completed")
}

// Authentication Decorator
type AuthPayment struct {
	payment Payment
}

func (a *AuthPayment) Pay(amount float64) {
	fmt.Println("Checking authentication")

	a.payment.Pay(amount)
}

func main() {

	// Original payment
	var payment Payment = &BasicPayment{}

	// Add logging
	payment = &LoggingPayment{
		payment: payment,
	}

	// Add authentication
	payment = &AuthPayment{
		payment: payment,
	}

	payment.Pay(1000)
}
