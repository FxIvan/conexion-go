package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Estas en esta ruta /")
	})

	hhtp.HandleFunc("/registro")

	direccion := ":5050"

	http.ListenAndServe(direccion, nil)

}
