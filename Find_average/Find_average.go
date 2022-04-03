package main

import (
	"fmt"
	"log"
	"test/datafile"
)

func main() {

	numbers, err := datafile.GetFloats("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {

		sum += number
	}

	count := float64(len(numbers))
	fmt.Print(sum / count)

}
