package main

import (
	"io"
	"log"
	"net/http"
)

func Servidor(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hola Mundo!\n")
}

func main() {
	http.HandleFunc("/holago", Servidor)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
