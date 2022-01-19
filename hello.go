package main

import (
	"fmt"

	"github.com/JSCOP/testgit/tree/test/greetings"
)

func SayHi() {
	message := greetings.Hello("jisung")
	fmt.Println(message)
}
