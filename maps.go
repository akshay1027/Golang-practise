package main

import (
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

type candidate struct {
	name       string
	interests  []string
	language   string
	experience bool
}

func main() {
	// syntax to create and initialise the map

	// just creating a map
	Mymap := make(map[string]string)

	Mymap["first"] = "Omkar"
	Mymap["second"] = "POP"

	log.Println(Mymap["first"])
	log.Println(Mymap["second"])

	// ========================================================================

	// create/declare and initiase
	myMap3 := map[string]int{
		"Age":  20,
		"Date": 30,
	}

	myMap3["Age2"] = 21

	log.Println(myMap3)

	// =========================================================================

	myMap4 := make(map[string]User)

	first := User{
		FirstName: "Omkar",
		LastName:  "Kulkarni",
	}

	myMap4["first"] = first

	// ==========================================================================

	candidates := []candidate{
		{
			name:       "ravi",
			interests:  []string{"art", "coding", "music", "travel"},
			language:   "golang",
			experience: false,
		},
	}
	fmt.Println(candidates)

	// ===========================================================================

}
