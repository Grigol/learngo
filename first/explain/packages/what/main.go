package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello!")
	hey()
	bye()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
}
