package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grades := []float64{}
	fmt.Print("please enter the greades: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("somthing happend: ", scanner.Err())
		return

	}
	userInput := scanner.Text()
	numbers := strings.Fields(userInput)
	for _, n := range numbers {
		grade, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		grades = append(grades, grade)
	}
	if len(grades) != 5 {
		fmt.Print("jsut 5 numbers \n")
		return
	}
	var all float64
	var avg float64

	for _, n := range grades {
		all += n

	}
	avg = all / float64(len(grades))
	fmt.Printf("the all is %v and the avg is %v \n", all, avg)

}
