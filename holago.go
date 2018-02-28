package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Usuario http
func Usuario(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Usuario</h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {

	msg := mensaje{
		msg: "Hola Nuevo Mundo",
	}
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	mux.HandleFunc("/user", Usuario)
	mux.Handle("/hola", msg)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando ...")
	log.Fatal(server.ListenAndServe())
}
