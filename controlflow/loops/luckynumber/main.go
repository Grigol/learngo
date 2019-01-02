package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Need range size and your guess")
		return
	}
	rLimit, err := strconv.Atoi(args[1])
	guess, err2 := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Provde a valid number for range limit")
		return
	}

	if err2 != nil {
		fmt.Println("Provide a valid number for your guess.")
		return
	}

	if rLimit < guess {
		fmt.Println("Your gess must be inside the range you chose.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for n := 0; n != guess; {
		n = rand.Intn(rLimit + 1)
		fmt.Printf("Showed number: %d\n", n)
	}
	fmt.Println()
}
