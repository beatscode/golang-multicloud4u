package main

import "fmt"

type person struct {
	name string
	age  int
	pets []pet
}

type pet struct {
	name string
	age  int
}

func (p person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years of age\n", p.name, p.age)
}

func (p *person) ChangeMyName(newName string) {
	p.name = newName
}
func main() {
	s := person{name: "Bob", age: 50}
	s.SayHello()
	s.ChangeMyName("Alex")
	s.SayHello()

	foo := pet{name: "foo", age: 1}
	s.pets = append(s.pets, foo)
	fmt.Println("Pets", s.pets)
}
