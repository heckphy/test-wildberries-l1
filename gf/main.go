package main

import (
	"fmt"
	"os"
)

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (h Human) SayHello() {
	fmt.Println(h.Greeting)
}

func main() {
	var x *os.PathError
	h := Human{Greeting: "Hello"}
	s := Speaker(h)
	h.Greeting = "Meow"
	s.SayHello()

	fmt.Printf("%T %v", x, x)
}
