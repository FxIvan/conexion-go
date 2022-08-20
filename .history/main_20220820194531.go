package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Estas en esta ruta /")
	})

	http.HandleFunc("/registro", func(w http.ResponseWriter, rsponse http.Request) {

	})

	direccion := ":5050"

	http.ListenAndServe(direccion, nil)

}
