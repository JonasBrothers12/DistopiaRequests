package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)
var key,err = rsa.GenerateKey(rand.Reader,256)

func Cript(mensagem string,chave_publica *rsa.PublicKey) (string,error)  {
    m_encrypt,err := rsa.EncryptPKCS1v15(rand.Reader,chave_publica,[]byte(mensagem))
    return base64.StdEncoding.EncodeToString(m_encrypt),err
}
func Descript(mensagemcript string ,chave_privada *rsa.PrivateKey) (string,error) {
    new_value,err := base64.StdEncoding.DecodeString(mensagemcript)
    m_decrypt,err := rsa.DecryptPKCS1v15(rand.Reader,chave_privada,new_value)
    return string(m_decrypt),err
}
func Cripthandler(w http.ResponseWriter, r *http.Request){
    value := r.URL.Query().Get("value")
    encriptado,err := Cript(value,&key.PublicKey)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Fprintln(w,encriptado)
}
func Descripthandler(w http.ResponseWriter, r *http.Request){
    value := r.URL.Query().Get("value")
    Desencriptado,err := Descript(value,key)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Fprintln(w,Desencriptado)
}

func main(){
	http.HandleFunc("/Cript", Cripthandler)
	http.HandleFunc("/Decript", Descripthandler)
	fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
