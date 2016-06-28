package main

import "fmt"

type Headlight struct {
	Brand string
	Type  string
}

func (h Headlight) TurnOff() {
	fmt.Printf("A %s's headlight turned off\n", h.Brand)
}

type Carrier interface {
	Capacity() int
}

type Car struct {
	Brand string
	Headlight
}

func (c Car) Run(speed int) {
	fmt.Printf("A %s running at %dmph\n", c.Brand, speed)
}

func (c Car) Capacity() int {
	return 4
}

func main() {
	headlight := Headlight{
		Brand: "Ford",
		Type:  "xenon",
	}
	car := Car{
		Brand:     "GMC",
		Headlight: headlight,
	}

	fmt.Printf("car: %v\n", car)
	car.Headlight.TurnOff()
	car.TurnOff()
	car.Run(65)
}
