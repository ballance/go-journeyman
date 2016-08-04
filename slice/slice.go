package main

import "fmt"

func main() {

	fmt.Println("It's slice time!!1!")
	fmt.Println()

	// Declare a nil slice of integers. Create a loop that appends 10 values to the
	var nums []int

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}

	for _, num := range nums {
		fmt.Println(num)
	}

	fmt.Println()

	strs := []string{"cheese", "pepperoni", "bacon", "peppers", "sausage"}

	for _, str := range strs {
		fmt.Println(str)
	}
	// Declare a slice of five strings and initialize the slice with string literal
	// values. Display all the elements. Take a slice of index one and two
	// and display the index position and value of each element in the new slice.

}
