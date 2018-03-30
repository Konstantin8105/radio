package main

import (
	"fmt"
	"os"

	"github.com/Konstantin8105/radio/radio"
)

func main() {
	err := radio.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
