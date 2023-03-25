package main

import (
	"fmt"
)

func main() {
	val := GetEnv("MNEMONIC_KEY")
	fmt.Println(val)
}
