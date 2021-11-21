package main

import "fmt"

func main() {
	var toDoList = [5]string{"wake up"}

	toDoList[1] = "brush teeth"
	toDoList[2] = "cook breakfast"
	toDoList[3] = "eat"
	toDoList[4] = "sleep"

	fmt.Printf("Number of elements in toDoList array: %d\n", len(toDoList))

	for index, item := range toDoList {
		fmt.Printf("%d. %s\n", index+1, item)
	}
}
