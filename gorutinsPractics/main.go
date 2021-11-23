package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // start time

	fmt.Printf("Start: %s\n", t.Format(time.RFC3339))

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go calculateSomething(1000) // calculation 1
	calculateSomething(2000)    // calculation 2

	fmt.Printf("Code execution time: %s\n", time.Since(t))
}

func calculateSomething(n int) {
	t := time.Now()
	result := 0

	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Result: %d; Time required: %s\n", result, time.Since(t))
}
