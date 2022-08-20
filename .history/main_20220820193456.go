package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	if err != nil {
		log.Fatal("Se requiere un PUERTO")
	}

	fmt.Println(PORT)

}
