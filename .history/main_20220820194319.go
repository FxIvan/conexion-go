package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/ola", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Estas en esta ruta /")
	})

	direccion := ":5050"

	http.ListenAndServe(direccion, nil)

}
