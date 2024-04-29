package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)
 
func main(){
    config := tls.Config{
        MinVersion: tls.VersionTLS12,
        MaxVersion: tls.VersionTLS12,
    }
    tr := http.Transport{
        TLSClientConfig: &config,
        }
    cl := http.Client{
        Transport: &tr,
    }
    req, err := http.NewRequest(http.MethodGet,"https://distopia-a1e2.savi2w.workers.dev/",nil)
    req.Header.Set("User-Agent","Jonas")
    if err != nil{
        fmt.Println(err)
        return
    } 
    resp,err := cl.Do(req)
    if err != nil{
        fmt.Println(err)
        return
    } 
    fmt.Println("Status: ",resp.Status)
}
    
