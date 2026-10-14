package main

import "fmt"

type Payment interface {
	pay(amount float64)
}

type Refund interface {
	refund(amount float64)
}

type Razorpay struct{}

func (r *Razorpay) pay(amount float64) {
	fmt.Println("pay through razorpay")
}

func (r *Razorpay) refund(amount float64) {
	fmt.Println("refund through razorpay")
}

type StripePay struct{}

func (r *StripePay) pay(amount float64) {
	fmt.Println("pay through Stripe")
}

func (r *StripePay) refund(amount float64) {
	fmt.Println("refund through Stripe")
}

type objectFactory interface {
	CreatePayment() Payment
	CreateRefund() Refund
}

type RazorPayFactory struct{}

func (r *RazorPayFactory) CreatePayment() Razorpay {
	return Razorpay{}
}

func (r *RazorPayFactory) CreateRefund() Razorpay {
	return Razorpay{}
}

type StripeFactory struct{}

func (r *StripeFactory) CreatePayment() Payment {
	return &StripePay{}
}

func (r *StripeFactory) CreateRefund() Refund {
	return &StripePay{}
}

func main() {
	input := StripeFactory{}

	payment := input.CreatePayment()
	refund := input.CreateRefund()

	payment.pay(11)
	refund.refund(45)

}
