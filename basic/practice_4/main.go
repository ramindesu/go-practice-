package main

import "fmt"

func main() {
	var firstNumber float64
	var secendNumber float64

	fmt.Print("enter the first number: ")
	_, err := fmt.Scanf( "%v", &firstNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("enter the secend number: ")
	_, err = fmt.Scanf( "%v", &secendNumber)

	if err != nil {
		fmt.Println(err)
		return
	}
	sub := firstNumber + secendNumber
	subtraction := firstNumber - secendNumber
	multiplicaton := firstNumber * secendNumber
	var divition float64
	if secendNumber == 0 {
		fmt.Println("zero divition error")
		return
	} else {
		divition = firstNumber / secendNumber

	}
	fmt.Printf("Sum:%g \n Subtraction: %g \n mult: %g \n, devide: %.2f \n", sub, subtraction, multiplicaton, divition)
}
