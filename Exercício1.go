package main

import (
	"fmt"
	"net/http"
) 

func main(){
	client := http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
    },
}
    resp, err := client.Get("https://distopia.savi2w.workers.dev/")
    if err != nil {
        fmt.Println(err)
        return
    } 
    defer resp.Body.Close()
    fmt.Println("A mensagem é: ",resp.Header.Values("Distopia"))
}
