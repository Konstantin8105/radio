package main

import (
	"fmt"
	"os"
)

func main() {
	err := radio.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
