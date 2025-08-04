package main

import "fmt"

func main() {
	answer := adder(3,5)
	result := proAdd(3,5,9,6)
	fmt.Println("result is : ", answer)
	fmt.Println("result is : ", result)
}


func adder(val1 int, val2 int) int {
	return val1 + val2
}


func proAdd(values ... int) int{
	total := 0;
	for _,val := range values{
		total += val
	}
	return total
}