//packages aape import hunde aa naale 
//
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating please:")

	input,_ := reader.ReadString('\n')
	fmt.Println("thanks for rating us:",input)
}