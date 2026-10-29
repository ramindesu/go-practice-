package main

import "fmt"

func main() {
	numbers := make([]int, 0, 5)

	for i := 0; i < 7; i++ {
		numbers = append(numbers, i)
		fmt.Println(cap(numbers))
		fmt.Println(len(numbers))
	}
}
