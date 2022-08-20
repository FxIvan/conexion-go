package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResonseJSON struct {
	Message string `json:"message"`
	Status  bool   `status:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encoder(w).Encode(ResonseJSON{Message: "Estas dentro de la ruta principal", Status: true})

	})

	http.HandleFunc("/registro", func(w http.ResponseWriter, response *http.Request) {
		io.WriteString(w, "Estas dentro de registro")
	})

	direccion := ":5050"

	http.ListenAndServe(direccion, nil)

}
