package main

import (
	"fmt"
)

func main() {
	// & gives you the address
	// * gives you the value stored at an address when you have the address
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
