package _interface

import (
	"fmt"
)

type Human interface {
	Speak() 
}

type Person struct {
	Name string
	Age  int
}

type Child struct {
	Name string
	Age  int
}

func (p Person) Speak() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (c Child) Speak() {
	fmt.Printf("Hi, I'm %s and I'm just a kid!\n", c.Name)
}

func InterfaceFn(h Human) {
	h.Speak()
}


func StructExample() {
	p := Person{Name: "John", Age: 30}
	c := Child{Name: "Alice", Age: 10}
	InterfaceFn(p)
	InterfaceFn(c)
}