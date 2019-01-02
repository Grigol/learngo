package main

import (
	"fmt"
	"time"
)

func main() {
	switch h := time.Now().Hour(); {
	case h < 12 && h >= 6:
		fmt.Println("Good Morning")
	case h < 18:
		fmt.Println("Good Afternoon")
	case h >= 18:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}

	fmt.Printf("Current hour is %d.\n", time.Now().Hour())
}
