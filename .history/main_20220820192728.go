package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	fmt.Print(err)

}
