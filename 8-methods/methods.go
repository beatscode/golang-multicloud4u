package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years of age", p.name, p.age)
}

func (p *person) ChangeMyName(newName string) {
	p.name = newName
}
func main() {
	s := person{name: "Bob", age: 50}
	s.SayHello()
	s.ChangeMyName("Alex")
	s.SayHello()
}
