package main

import (
	"fmt"
	"os"
	"strconv"
)

const feetInMeter = 0.3048

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a number")
		return
	}

	feets, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Printf("ERROR: %q is not a number!\n", args[1])
		return
	}

	fmt.Printf("SUCCESS: %g feet is equal to %g meters.\n", feets, feets*feetInMeter)
}
