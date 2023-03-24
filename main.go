package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := getEnv("MNEMONIC_KEY")
	fmt.Println(val)
}

// use this function only for string outputs
func getEnv(env string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv(env)
	return val
}
