package SayHelloModule

import (
	"fmt"
)

func SayHello(name string) string {
	message := fmt.Sprintln("Hello ", name, "! Nice to meet you!")
	return message
}
