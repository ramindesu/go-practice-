package main

import "fmt"


func main(){
	names := [4]string{"ram" , "ghaz" , "sani" , "fifi"}

	for  _,n := range names{
		fmt.Println(n)

	}
	println("all names are: " , len(names))
}