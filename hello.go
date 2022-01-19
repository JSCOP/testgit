package main

import (
	"fmt"

	"github.com/JSCOP/testgit/tree/test/SayHelloModule"
)

func SayHi() {
	message := SayHelloModule.Hello("jisung")
	fmt.Println(message)
}
