package main

import "fmt"

func main() {
	var count int = 23
	var result string
	if count < 10 {
		result = "not a regular user"
	} else if count > 30 {
		result = "best user"
	}else if count == 23 {
		result = "optimal user"
	}
	fmt.Println(result)



	//this is used in webhooks where a result comes and we need to check it.

	if num:=3;num<10{
		fmt.Println("number is less than 10")
	}else{
		fmt.Println("number isnt less than 10")
	}
}
