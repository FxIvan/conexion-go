package main

import (
	"io"
	"net/http"
)

type ResonseJSON struct {
	Message string `json:"message"`
	Status  bool   `status:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Estas en esta ruta /")
	})

	http.HandleFunc("/registro", func(w http.ResponseWriter, response *http.Request) {
		io.WriteString(w, "Estas dentro de registro")
	})

	direccion := ":5050"

	http.ListenAndServe(direccion, nil)

}
