package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("please rate our pizza : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating :",input)

	fmt.Println("added 1 to your rating")
	newrating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to ur rating :",newrating+1)
	}
	

}
