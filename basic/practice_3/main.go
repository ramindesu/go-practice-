package main

import (
	"fmt"
	"time"
)

func main() {
	var age int
	fmt.Print("please enter youre year of birth: ")
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("somthing is wrong")
		return
	}
	rightNow := time.Now().Year()
	userAge := rightNow - age
	fmt.Printf("youre age approxmatily is %d \n", userAge)

}
