package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println(os.Getenv("PORT"))
}
