package main

import "fmt"

func g() {
	list := []string{"apple", "banana", "grapes", "melon", "mango"}
	fmt.Println("Before:", list)

	index := 2
	list = append(list[:index], list[index+1:]...)
	fmt.Println("After :", list)
}


