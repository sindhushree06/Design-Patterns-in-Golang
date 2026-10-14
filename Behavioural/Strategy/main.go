package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) error
}

type PaymentService struct {
	strategy PaymentStrategy
}

type CreditCardPayment struct{}

type PaypalPayment struct{}

func (pt *PaymentService) SetStrategy(strategy PaymentStrategy) *PaymentService {
	pt.strategy = strategy
	return pt
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Println("payment using credit card")
	return nil
}

func (p *PaypalPayment) Pay(amount float64) error {
	fmt.Println("payment using paypal")
	return nil
}

func (n *PaymentService) Pay(amount float64) error {
	return n.strategy.Pay(amount)
}

func getPaymentFactory(paymentType string) PaymentStrategy {
	switch paymentType {
	case "creditcard":
		return &CreditCardPayment{}
	case "paypal":
		return &PaypalPayment{}
	default:
		return nil
	}
}

func main() {

	received := "creditcard"

	factory := getPaymentFactory(received)

	p := &PaymentService{}
	p.SetStrategy(factory)
	p.Pay(100.0)

}
