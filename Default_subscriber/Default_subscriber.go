package main

import "fmt"

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Address struct {
	Street    string
	City      string
	Buildbing string
}

func printInfo(s Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Rate:", s.Rate)
	fmt.Println("Active:", s.Active)
	fmt.Println("Address", s.HomeAddress)

}

func DefaultSubscriber(name string) Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.00
	s.Active = true
	s.HomeAddress.City = "Toronto"
	return s

}

func ApplyDiscout(s *Subscriber) {
	s.Rate = 4.99
}

func main() {

	address := Address{
		Street:    "Backer",
		City:      "London",
		Buildbing: "23",
	}

	subscriber1 := DefaultSubscriber("Akzhol")
	subscriber1.HomeAddress = address
	ApplyDiscout(&subscriber1)
	printInfo(subscriber1)

}
