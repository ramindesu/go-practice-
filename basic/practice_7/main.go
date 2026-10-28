package main

import (
	"bufio"
	"fmt"

	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("please enter a sentence: ")

	if !scanner.Scan() {
		fmt.Println("error happend: ", scanner.Err())
		return
	}
	input := scanner.Text()
	lenCount:=0
	for _ , char := range input{
		if char != ' '{
			lenCount++
		}
	}
	fmt.Printf("len is : %d \n" , lenCount)
}
