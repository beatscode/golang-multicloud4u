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