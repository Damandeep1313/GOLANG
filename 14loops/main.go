package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	fmt.Println(days)

	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days{
	// 	fmt.Printf("day  %v is  %v \n",i+1,day )
	// }

	value:=1
	for value<10{

		if value==2 {
			goto lco
		}
		if value == 5{
			value++
			continue
		}
		fmt.Println("value is : ",value)
		value++
	}

	lco:
		fmt.Println("nzaare")
}
