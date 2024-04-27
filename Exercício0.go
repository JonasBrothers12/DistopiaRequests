package main
//JonasBrothers12
import (
	"fmt"
	"io"
	"net/http"
) 
func main(){
	resp,err := http.Get("https://api.ipify.org?format=json")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	body,err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("O texto do site é: ",string(body))

}
