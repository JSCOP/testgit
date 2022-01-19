package SayHello

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintln("Hello ", name, "! Nice to meet you!")
	return message
}
