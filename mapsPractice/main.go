package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	grades := map[string]int{
		"Oleg":  5,
		"Anton": 4,
		"Anna":  3,
		"Irina": 4,
	}

	delete(grades, "Irina")
	grades["Anton 2"] = 4

	grade, exists := grades["Anton"]
	if !exists {
		fmt.Println("There is no in the map.")
	} else {
		fmt.Printf("Was graded %d.\n", grade)
	}

	// this one can print maps
	bs, _ := json.Marshal(grades)
	fmt.Println(string(bs))

}
