package main
import(
	"fmt"
	"net/http"
	"io"
)
 
const url = "https://lco.dev"
func main(){
	response,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()		//caller's responsibility to close the connection
	databytes,err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("databytes are : ",databytes)
}