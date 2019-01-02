package main

import (
	"fmt"
	"os"
)

func sec01Ex01(myName string, bffName string) {
	fmt.Printf("My name is %q\n", myName)
	fmt.Printf("My bff's name is %q\n", bffName)
}

func sec01Ex02() {
	fmt.Println(os.Getenv("GOPATH"))
}

func main() {
	fmt.Println("Hello, Gopher")
	sec01Ex01("Diego", "Lorenzzo")
	sec01Ex02()
}
