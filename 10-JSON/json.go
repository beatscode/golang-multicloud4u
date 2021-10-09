package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  32,
	}

	p2 := person{
		Name: "Jackie Chan",
		Age:  27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s := `[{"Name":"James Bond","Age":32},{"Name":"Jackie Chan","Age":27}]`
	bs = []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.Name, v.Age)
	}

}
