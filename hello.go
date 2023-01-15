// hello.go
package main

import (
	"example.com/greetings"

	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings module error: ")
	log.SetFlags(0)

	names := []string{"Andrew", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
