package main

import "fmt"

//There is no inheritance in golang,
// no super or parent.As there are no classes we use structs { }.

func main() {
	user1 := User{"Daman", 23, "daman@uuuu", true}
	fmt.Println(user1)
	fmt.Printf("details of user1 are :\n %+v \n",user1)
	fmt.Printf("name is %v,and email is %v :\n",user1.Name,user1.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
