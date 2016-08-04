package main

import (
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "slice")

	//start := time.Now()

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
	fmt.Println("slice of 5:")
	for _, str := range strs {
		fmt.Println(str)
	}

	fmt.Println()
	fmt.Println("sub-slice of 2:")

	firsttwo := strs[:2]

	for _, ft := range firsttwo {
		fmt.Println(ft)
	}

	//elapsed := time.Since(start)
	//log.Printf("slice took %s", elapsed)

}
