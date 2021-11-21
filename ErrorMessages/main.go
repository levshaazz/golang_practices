package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(-3)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Radius of circle is %d sm.\n", radius)
	fmt.Println("Formula to calculate radius is: A = pi*r^2")
	fmt.Printf("Area of circle is %f squared sm.\n\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Radius can not be negative")
	}

	return float32(radius) * float32(radius) * pi, nil
}
