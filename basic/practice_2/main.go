package main

import "fmt"

func main() {
	var name string
	fmt.Print("please enter your name: ")
	_,err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("smth is not okay")
	}
	fmt.Printf("hello %s \n", name)

}
