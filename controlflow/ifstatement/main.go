package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
Usage: [username] [password]`

const (
	validUser = "grigol"
	validPwd  = "grigol"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	u, p := args[1], args[2]
	if u == validUser && p == validPwd {
		fmt.Printf("Access granted for %q.\n", u)
	} else if u != validUser {
		fmt.Printf("Access denied for %q.\n", u)
	} else {
		fmt.Printf("Invalid Password for %q.\n", u)
	}
	return
}
