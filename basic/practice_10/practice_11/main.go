package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var evenNumbers []int
	for _, i := range numbers {
		if i%2 == 0 {
			evenNumbers = append(evenNumbers, i)
		}
	}
	fmt.Println("even numbers: ", evenNumbers)

}
