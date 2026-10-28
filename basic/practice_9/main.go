package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	largestNumber := numbers[0]
	smallestNumber := numbers[0]

	for _, n := range numbers {
		if largestNumber < n {
			largestNumber = n
		}
		if smallestNumber > n {
			smallestNumber = n
		}

	}
	fmt.Printf("largest number: %v and smallest is %v" , largestNumber, smallestNumber )

}
