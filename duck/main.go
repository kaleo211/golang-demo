package main

import (
	"fmt"
)

type Car struct{}

func (c Car) Run() {
	fmt.Println("I'm running by 4 wheels.")
}

type Human struct{}

func (h Human) Run() {
	fmt.Println("I'm running by 2 legs.")
}

type Runner interface {
	Run()
}

type Transformer struct {
	Runner
}

func (t Transformer) Shape() {
	switch t.Runner.(type) {
	case Car:
		fmt.Println("Now, I am a car")
	case Human:
		fmt.Println("Now, I am a human")
	}
}

func main() {
	var t Transformer

	var h Human
	t.Runner = h
	t.Shape()
	t.Run()

	var c Car
	t.Runner = c
	t.Shape()
	t.Run()
}
