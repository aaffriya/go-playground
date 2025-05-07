package _interface

import (
	"fmt"
)

// type Animal interface {
// 	Speak() string
// }

type Dog struct {
	Name string
}
type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Meow!", c.Name)
}
func InterfaceExample() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())
}
