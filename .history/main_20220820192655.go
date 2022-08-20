package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	erri := godotenv.Load(".env")

	fmt.Println(err)

}
