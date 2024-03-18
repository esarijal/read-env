package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	envFile := ".env" // Nama file default
	if len(os.Args) > 1 {
		envFile = os.Args[1]
	}

	_ = godotenv.Load(envFile)
	fmt.Println("Run")
	test := os.Getenv("TEST")
	fmt.Println(test)
}
