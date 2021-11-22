package main

import (
	"fmt"
	"sort"
)

// this one sorts array
func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Printf("Array: %d \n", s)

	sort.Sort(sort.IntSlice(s))
	fmt.Printf("Sorted Array: %d \n", s)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Printf("Reversed Sorted Array: %d \n", s)
}
