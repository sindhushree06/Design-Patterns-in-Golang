package main

import "fmt"

type Observer interface {
	Update() error
}

type Subject struct {
	observers []Observer
}

type smsUpdate struct{}

func (sms smsUpdate) Update() error {
	fmt.Println("send sms")
	return nil
}

type emailUpdate struct{}

func (email emailUpdate) Update() error {
	fmt.Println("send email")
	return nil
}

func (s Subject) Subscribe(Observer Observer) {
	s.observers = append(s.observers, Observer)
}

func (s Subject) Notify() error {
	observers := s.observers
	for _, observer := range observers {
		observer.Update()
	}
	return nil
}

func main() {
	sub := Subject{}

	email := emailUpdate{}

	sub.Subscribe(&email)

}
