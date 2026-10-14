// Package main demonstrates the Adapter design pattern.
// It bridges our system's Payment interface with a third-party RazorPay API
// whose method signature (makePayment) differs from what our system expects (Pay).
package main

import "fmt"

type PaymentGateway interface {
	Pay(amount float64)
}

type paytm struct{}

// this is the adapter
type PaytmAdapter struct {
	paytm *paytm
}

func (p *paytm) SendMoney(amount float64) {
	fmt.Print("sending money through razorpay")
}

func (p *PaytmAdapter) Pay(amount float64) {
	p.paytm.SendMoney(amount)
}

func doPayment(gateway PaymentGateway, amount float64) {
	gateway.Pay(amount)
}

func main() {
	paytm := &paytm{}

	adapter := &PaytmAdapter{paytm: paytm}

	doPayment(adapter, 100.0)
}
