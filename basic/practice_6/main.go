package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("please enter your name: ")
	scanner.Scan()
	if !scanner.Scan() {
		fmt.Println("input error:", scanner.Err())
		return
	}
	userData := scanner.Text()
	_, err := strconv.Atoi(userData)
	if err == nil {
		fmt.Println("please eneter a text not a nunber")
		return
	}

	fmt.Printf("hello and welcome %v \n", userData)

}
