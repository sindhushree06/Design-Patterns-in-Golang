package main

// centralise the creation of objects and hide the logic of object creation from the client
type PaymentGateway interface {
	Pay(amount float64)
}

type upi struct{}

type card struct{}

func (u *upi) Pay(amount float64) {
	println("paying through UPI")
}

func (c *card) Pay(amount float64) {
	println("paying through card")
}

type PaymentStrategy struct {
	method PaymentGateway
}

func (p *PaymentStrategy) SetMethod(method PaymentGateway) {
	p.method = method
}

func (p *PaymentStrategy) Pay(amount float64) {
	p.method.Pay(amount)
}

func PaymentMethodFactory(method string) PaymentGateway {
	switch method {
	case "upi":
		upi := &upi{}
		return upi
	case "card":
		card := &card{}
		return card
	}
	return nil
}

func main() {
	paymentStrategy := &PaymentStrategy{}
	paymentStrategy.SetMethod(PaymentMethodFactory("upi"))
	paymentStrategy.Pay(100.0)
}
