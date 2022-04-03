package main

import (
	"fmt"
	"log"
	"test/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(3011)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day())
}
