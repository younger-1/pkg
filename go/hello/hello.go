package main

import (
	"log"

	"github.com/younger-1/pkg/go/greetings"
)

func main() {
	// Set properties of the predefined Logger
	//   1. the log entry prefix and
	//   2. set flag to disable printing the time, source file, and line number.
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	log.SetFlags(log.Lshortfile | log.Ltime)

	// Request a greeting message.
	message, err := greetings.Hello("Younger")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	log.Println(message)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	log.Println(messages)
}
