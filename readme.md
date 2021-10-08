# Golang Talk - MultiCloud4u


- 1) Hello World
    - fmt is a logging library
    - Println will print to stdout
    
    ```go
    package main
    
    import "fmt"
    
    func main(){
    	fmt.Println("Hello world")
    }
    ```
    
- 2) Variables
    - The var statement declares a list of variables; as in function argument lists, the type is last.
    - Go will not compile if the declared variables are not used
    
    ```go
    package main
    
    import "fmt"
    
    func main(){
    	// `var` declares 1 or more variables.
    	var a = "initial"
    	fmt.Println(a)
    
    	// You can declare multiple variables at once.
    	var b, c int = 1, 2
    	fmt.Println(b, c)
    
    	// Go will infer the type of initialized variables.
    	var d = true
    	fmt.Println(d)
    
    	// Variables declared without a corresponding
    	// initialization are _zero-valued_. For example, the
    	// zero value for an `int` is `0`.
    	var e int
    	fmt.Println(e)
    
    	// The `:=` syntax is shorthand for declaring and
    	// initializing a variable, e.g. for
    	// `var f string = "apple"` in this case.
    	f := "apple"
    	fmt.Println(f)
    }
    ```
    
- 3) For Loops, If Else
    - For Loop helps to iterate through a group of statements multiple times
    - The Go for loop has four forms
        
        1. for <initialization>; <condition>; <increment/decrement> { }
        2. for <condition> { } like a while loop
        3. for { } an infinite while loop.
        4. for with range.
        
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    	//1. for <initialization>; <condition>; <increment/decrement> { }
    	loopWithInit()
    	//2. for <condition> { } like a while loop
    	loopWithCondition()
    	//3. for { } an infinite while loop.
    	loopInfiniteWhile()
    	//4. for with range.
    	loopWithRange()
    }
    
    func loopWithInit() {
    	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    	sum := 0
    	for i := 0; i < len(numbers); i++ {
    		sum += numbers[i]
    	}
    	fmt.Println("Sum is :: ", sum)
    }
    
    func loopWithCondition() {
    	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    	sum := 0
    	i := 0
    	n := len(numbers)
    	for i < n {
    		sum += numbers[i]
    		i++
    	}
    	fmt.Println("Sum is :: ", sum)
    }
    
    func loopInfiniteWhile() {
    	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    	n := len(numbers)
    	sum := 0
    	i := 0
    	for {
    		sum += numbers[i]
    		i++
    		if i >= n {
    			break
    		}
    	}
    	fmt.Println("Sum is :: ", sum)
    }
    
    func loopWithRange() {
    	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    	sum := 0
    	for _, val := range numbers {
    		sum += val
    	}
    	fmt.Println("Sum is :: ", sum)
    }
    ```
    
- 4) Array & Slices
    
    ## Array
    
    Arrays are mostly used as a building block in the Go programming language. In some instances, we might use an array, but mostly the recommendation is to use slices instead.
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    	// Here we create an array `a` that will hold exactly
    	// 5 `int`s. The type of elements and length are both
    	// part of the array's type. By default an array is
    	// zero-valued, which for `int`s means `0`s.
    	var a [5]int
    	fmt.Println("emp:", a)
    	
    	// We can set a value at an index using the
    	// `array[index] = value` syntax, and get a value with
    	// `array[index]`.
    	a[4] = 100
    	fmt.Println("set:", a)
    	fmt.Println("get:", a[4])
    	
    	// The builtin `len` returns the length of an array.
    	fmt.Println("len:", len(a))
    	
    	// Use this syntax to declare and initialize an array
    	// in one line.
    	b := [5]int{1, 2, 3, 4, 5}
    	fmt.Println("dcl:", b)
    	
    	// Array types are one-dimensional, but you can
    	// compose types to build multi-dimensional data
    	// structures.
    	var twoD [2][3]int
    	for i := 0; i < 2; i++ {
    		for j := 0; j < 3; j++ {
    			twoD[i][j] = i + j
    		}
    	}
    	fmt.Println("2d: ", twoD)
    }
    
    ```
    
    ### Slice - composite literal
    
    A SLICE holds VALUES of the same TYPE. If I wanted to store all of my favorite numbers, I would use a slice of int. If I wanted to store all of my favorite foods, I would use a slice of string. We will **use a COMPOSITE LITERAL to create a slice.** A composite literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.
    
    code
    
    ```go
    // In Go, an _array_ is a numbered sequence of elements of a
    // specific length.
    
    package main
    
    import "fmt"
    
    func main() {
    
    	// Here we create an array `a` that will hold exactly
    	// 5 `int`s. The type of elements and length are both
    	// part of the array's type. By default an array is
    	// zero-valued, which for `int`s means `0`s.
    	var a [5]int
    	fmt.Println("emp:", a)
    
    	// We can set a value at an index using the
    	// `array[index] = value` syntax, and get a value with
    	// `array[index]`.
    	a[4] = 100
    	fmt.Println("set:", a)
    	fmt.Println("get:", a[4])
    
    	// The builtin `len` returns the length of an array.
    	fmt.Println("len:", len(a))
    
    	// Use this syntax to declare and initialize an array
    	// in one line.
    	b := [5]int{1, 2, 3, 4, 5}
    	fmt.Println("dcl:", b)
    
    	// Array types are one-dimensional, but you can
    	// compose types to build multi-dimensional data
    	// structures.
    	var twoD [2][3]int
    	for i := 0; i < 2; i++ {
    		for j := 0; j < 3; j++ {
    			twoD[i][j] = i + j
    		}
    	}
    	fmt.Println("2d: ", twoD)
    }
    ```
    
- 5) Maps
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    
    	states := make(map[string]string)
    
    	states["New York"] = "New York City"
    	states["California"] = "Los Angeles"
    	states["Illinois"] = "Chicago"
    	states["Texas"] = "Houston"
    	states["Arizona"] = "Phoenix"
    
    	//Get Map Length
    	fmt.Println("len:", len(states))
    
    	// Delete from map
    	delete(states, "Texas")
    
    	//Print map
    	fmt.Println("map:", states)
    
    	//Check if key exists
    	_, prs := states["Arizona"]
    	fmt.Println("prs:", prs)
    
    }
    ```
    
- 6) Range
    
    We can loop over the values in a slice with the range clause. We can also access items in a slice by index position.
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
        nums := []int{2, 3, 4}
        sum := 0
        for _, num := range nums {
            sum += num
    	  }
    }
    ```
    
- 7) Structs
    
    ```go
    package main
    
    import "fmt"
    
    type person struct {
    	name string
    	age  int
    }
    
    func newPerson(name string) *person {
    	p := person{name: name}
    	p.age = 42
    	return &p
    }
    
    func main() {
    
    	fmt.Println(person{name: "Bob", age: 30})
    	fmt.Println(person{name: "Fred"})
    	fmt.Println(&person{name: "Ann", age: 40})
    	fmt.Println(newPerson("Jon"))
    
    	s := person{name: "Sean", age: 50}
    	fmt.Println(s.name)
    
    	s.age = 51
    	fmt.Println(s.age)
    }
    ```
    
- 8) Methods
    
    ```go
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
    ```
    
- 9) Pointers
    
    Go supports pointers, allowing you to pass references to values and records within your program.
    
    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	x := 0
    	fmt.Println("x befor", &x)
    	fmt.Println("x befor", x)
    	foo(&x)
    	fmt.Println("x after", &x)
    	fmt.Println("x after", x)
    }
    
    func foo(y *int) {
    	fmt.Println("y befor", y)
    	fmt.Println("y befor", *y)
    	*y = 43
    	fmt.Println("y after", y)
    	fmt.Println("y after", *y)
    }
    ```
    
- Interfaces
    
    In Go, values can be of more than one type. An interface allows a value to be of more than one type.
    
    ```go
    package main
    
    import (
    	"fmt"
    	"math"
    )
    
    // Go Interface - `Shape`
    type Shape interface {
    	Area() float64
    	Perimeter() float64
    }
    
    // Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
    type Rectangle struct {
    	Length, Width float64
    }
    
    func (r Rectangle) Area() float64 {
    	return r.Length * r.Width
    }
    
    func (r Rectangle) Perimeter() float64 {
    	return 2 * (r.Length + r.Width)
    }
    
    // Struct type `Circle` - implements the `Shape` interface by implementing all its methods.
    type Circle struct {
    	Radius float64
    }
    
    func (c Circle) Area() float64 {
    	return math.Pi * c.Radius * c.Radius
    }
    
    func (c Circle) Perimeter() float64 {
    	return 2 * math.Pi * c.Radius
    }
    
    func (c Circle) Diameter() float64 {
    	return 2 * c.Radius
    }
    
    func main() {
    	var s Shape = Circle{5.0}
    	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
    	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())
    
    	s = Rectangle{4.0, 6.0}
    	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
    	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())
    }
    ```
    
- 10) JSON
    
    **What does 'Marshaling' mean?**
    
    **marshaling** is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be moved between different parts of a computer program or from one program to another.
    
    The inverse of marshaling is called unmarshaling or demarshaling.
    