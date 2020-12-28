package main

import (
	"fmt"
	"log"
	"main/greetings"
)

func main() {
	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	names := []string{"Hyunjae", "Sunmi", "Delta"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
