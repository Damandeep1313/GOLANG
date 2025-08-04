package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content:="Waheguru Ji!"
	file,err := os.Create("./file.txt")
	checkNilErr(err)
	// if err!= nil{
	// 	panic(err)
	// }
	len,err:=io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println("length is :",len)
	defer file.Close()

	readFile("./file.txt")

}
func readFile(filename string){
	databytes,err:=os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("content inside file :\n",string(databytes))
}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}