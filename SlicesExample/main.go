package main

import "fmt"

func main() {
	toDoList := []string{
		"a",
		"b",
		"c",
		"d",
	}

	newToDoList := append(toDoList, "e", "f")

	fmt.Printf("Old slice. Lenght: %d, Capacity: %d\n", len(toDoList), cap(toDoList))
	fmt.Printf("New slice. Lenght: %d, Capacity: %d\n", len(newToDoList), cap(newToDoList))
}
