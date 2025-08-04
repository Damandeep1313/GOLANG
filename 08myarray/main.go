package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	//fruits[2] = "cherry"
	fruits[3] = "melon"
	fmt.Println("fruit list is :", fruits)
	fmt.Println("fruit list is :", len(fruits))
}
