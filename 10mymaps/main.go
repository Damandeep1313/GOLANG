package main

import "fmt"

func main() {
	books := make(map[string]string)
	books["eng"] = "harry potter"
	books["hindi"] = "harry putr"
	books["punjabi"] = "harry puttar"
	books["urdu"] = "harry beta"

	fmt.Println(books)
	fmt.Println(books["eng"])

	delete(books, "hindi")	//deletion
	fmt.Println(books)

							//loops in go
	for key, val := range books {
		fmt.Printf("for key '%v' , the value is '%v' \n", key, val)
	}
}
