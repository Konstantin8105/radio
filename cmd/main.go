package main

import (
	"fmt"
	"os"

	radio "github.com/Konstantin8105/Xi-radio"
)

func main() {
	err := radio.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
