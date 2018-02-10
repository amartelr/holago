package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Servidor http
func Servidor(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hola Mundo!\n")
}

func main() {
	fmt.Println("http://localhost/holago:8888")
	http.HandleFunc("/holago", Servidor)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
