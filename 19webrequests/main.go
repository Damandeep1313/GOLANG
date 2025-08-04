package main
import{
	"fmt"
}
 
const url = "https://lco.dev"
func main(){
	response,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()		//caller's responsibility to close the connection
	databytes,err := ioutil.ReadAll(response.Body)
	fmt.Println("databytes are : ",databytes)
}