package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()
	for {

		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10],
			digits[hour%10],
			separator,
			digits[min/10],
			digits[min%10],
			separator,
			digits[sec/10],
			digits[sec%10],
		}

		fmt.Println()
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]

				if digit == separator && sec%2 == 0 {
					next = "   "
				}

				fmt.Printf("%s  ", next)
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
