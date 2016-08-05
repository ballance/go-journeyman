// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the match method for hockey, call into the match method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
	//	Append(itemToAppend hockey) matcher
}

type []matcher interface {
	//match(searchTerm string) bool
	Append(itemToAppend hockey) matcher
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// match checks the value for the specified term.
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

type hockey struct {
	sport
	country string
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.

// match checks the value for the specified term.
func (h hockey) match(searchTerm string) bool {
	return h.sport.match(searchTerm) || strings.Contains(h.country, searchTerm)
}

// Append an additional hockey item to an existing hockey[] slice
//func (h matchers) Append(slice []matcher, element matcher) []matcher {
//	n := len(slice)
//	slice = slice[0 : n+1]
//	slice[n] = element
//	return slice
//}

func main() {

	defer timeTrack(time.Now(), "main")

	// Define the term to match.
	term := "Hurricanes"

	// Create a slice of matcher values and assign values
	// of the concrete hockey type.
	matchers := []matcher{
		hockey{sport{"Ranchers", "Dallas"}, "USA"},
		hockey{sport{"Cattle", "Houston"}, "USA"},
		hockey{sport{"Canadiens", "Montreal"}, "Canada"},
		hockey{sport{"Hurly Burlies", "Charlotte"}, "USA"},
	}

	matchers = []matchers.Append(matchers, hockey{sport{"Hurricanes", "Raleigh"}, "USA"})

	// Display what we are trying to match.
	fmt.Println("Matching For:", term)
	fmt.Println()

	// Range of each matcher value and check the term.
	for _, m := range matchers {
		if m.match(term) {
			fmt.Println("FOUND: %+v", m)

		}
	}
}
