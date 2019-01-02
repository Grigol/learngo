package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("Element #%-2d: %q.\n", i, v)
	}
}
