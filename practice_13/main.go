package main

import "fmt"

func main() {

	phoneBook := make(map[string]int)
	phoneBook["ram"] = 989919404389
	phoneBook["ghaz"] = 989392563698
	var input string
	fmt.Print("please enter a name: ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		print(err)
		return
	}
	phone, exists := phoneBook[input]
	if exists{
		fmt.Println("here is the phone:" ,phone)
	}else{
		fmt.Println("sorry")
	}
}
