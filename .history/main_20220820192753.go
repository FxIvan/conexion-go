package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	PORT := os.Getenv("PORT")

}
