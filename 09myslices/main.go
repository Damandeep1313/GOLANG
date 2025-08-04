package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "grapes"}
	fmt.Println("fruit list :", fruits)
	fruits = append(fruits,"watermelon")
	fmt.Println("new list :",fruits)
	fruits = append(fruits[:2],"dragonfruit")
	fmt.Println("newest list: ",fruits)


	highscores:=make([]int,4)//still a slice dont confuse with array.
	highscores[0]=1			 //array would be ------var highscores [4]int
	highscores[1]=2			 //array has fixed size 
	highscores[2]=3
	highscores[3]=4
	// highscores[4]=5
	fmt.Println("list of scores:",highscores)

	//now if we use append it will realocate the memory earlier allocated for make
	highscores = append(highscores,6,7)
	fmt.Println("latest scores :",highscores)
	length := len(highscores)
	fmt.Println("length is :",length)
	sort.Sort(sort.Reverse(sort.IntSlice(highscores)))//🔥 
	fmt.Println(highscores)


	g()
}

