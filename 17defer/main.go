package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("JI")
	g()
}
func g(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
