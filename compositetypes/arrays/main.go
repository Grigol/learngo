package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nMoods = 5
)

func compareArray() {
	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	data := [5]uint8{}

	fmt.Println(data == image)
}

func assignArray() {
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	books := [4]string{}

	for i, v := range prev {
		books[i] += v + " 2nd Ed."
	}

	books[3] = "Awesomeness"

	fmt.Printf("Last year:\n%#v\n", prev)
	fmt.Printf("Now:\n%#v\n", books)

}

func fix() {
	fmt.Println("Fix Exercise on arrays")

	names := [...]string{"Einstein", "Shepard", "Tesla"}
	books := [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

func mArray() {
	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}
	var sum float64

	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	avg := sum / float64(len(students)*len(students[0]))

	fmt.Printf("AVG: %#v\n", avg)
}

func moodChallenge(args []string) {
	const (
		nMoods    = 4
		msg       = "%s is feeling %s.\n"
		noMoodMsg = "No mood for you now. Provide Positive or negative as status :)"
		usage     = "Provide a name and a mood status."
	)

	var i int

	if len(args) < 2 {
		fmt.Println(usage)
		return
	}

	name, moodStatus := args[0], args[1]

	moods := [...][nMoods]string{
		{"great", "good", "happy", "awesome"},
		{"bad", "sad", "terrible", "mad"},
	}

	rand.Seed(time.Now().UnixNano())
	i = rand.Intn(nMoods)

	if moodStatus == "positive" {
		fmt.Printf(msg, name, moods[0][i])
	} else if moodStatus == "negative" {
		fmt.Printf(msg, name, moods[1][i])
	} else {
		fmt.Println(noMoodMsg)
	}

}

func keyedIndex() {
	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,  // Ethereum
		WAN: 120.5, // wainchain
	}

	fmt.Printf("1 BTC is %g ETH.\n", rates[ETH])
}

func main() {
	// compareArray()
	// assignArray()
	// fix()
	// mArray()
	// moodChallenge(os.Args[1:])
	keyedIndex()
}
