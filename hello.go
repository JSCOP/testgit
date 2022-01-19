package main

import (
	"fmt"

	"SayHelloModule"
)

func SayHi() {
	message := SayHelloModule.Hello("jisung")
	fmt.Println(message)
}
