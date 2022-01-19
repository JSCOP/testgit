package main

import (
	"fmt"

	"github.com/JSCOP/testgit/tree/test/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
