package main

import (
	"fmt"

)

func main() {
	var width float64
	var height float64
	var area float64
	var perimeter float64

	fmt.Print("please enter the width: ")
	_, err := fmt.Scanf("%f", &width)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("please enter the height: ")
	_, err = fmt.Scanf("%f", &height)
	if err != nil {
		fmt.Print(err)
		return
	}
	if width <= 0 || height <=0{
		fmt.Print("cant be zero")
		return
	}
	area = width * height
	perimeter = (width + height)*2
	fmt.Printf("Area: %g \nPerimeter: %g \n", area , perimeter)
}
