package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("value of ptr is :", ptr)

	myAge := 23
	var ptr = &myAge
	fmt.Println("value of ptr is: ",ptr)
	fmt.Println("value of ptr is: ",*ptr)

	*ptr = *ptr * 2
	fmt.Println("new value : ",*ptr)
}
