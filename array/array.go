package main

import "fmt"

func main() {

	// Declare string arrays to hold names.
	var names [5]string

	// Declare an array pre-populated with friend's names.
	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	// Assign the array of friends to the names array.
	names = friends

	// Display each string value and address index in names.
	for i, name := range names {
		fmt.Println(name, &names[i])
	}
}
