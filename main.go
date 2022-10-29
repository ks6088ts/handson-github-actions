package main

import (
	"fmt"
	"os"
)

func main() {
	number := os.Getenv("INPUT_NUMBER")
	fmt.Printf("::set-output name=number::%s\n", number)
}
