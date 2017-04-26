package main

import "fmt"

func main() {
	newMap()
}

func newMap() {
	// Maps can have keys of almost any type.
	//
	// Create a map from the names of days, eg "Monday", "Tuesday", to their
	// numbers. 1 for Monday, 2 for Tuesday.

	// uncomment and complete this declaration
	daynumber := map[string]int{
		"Monday": 1,
		"Tuesday": 2,
		"Wednesday": 3,
		"Thursday": 4,
		"Friday": 5,
		"Saturday": 6,
		"Sunday": 7,
	}

/*	daynumber["Monday"] = 1
	daynumber["Tuesday"] = 2
	daynumber["Wednesday"] = 3
	daynumber["Thursday"] = 4
	daynumber["Friday"] = 5
	daynumber["Saturday"] = 6
	daynumber["Sunday"] = 7*/
	// The correct answer should print 3 6
	fmt.Println(daynumber["Wednesday"], daynumber["Saturday"])
}
