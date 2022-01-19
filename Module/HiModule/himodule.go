package hi

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintln("Nice to meet you!", name, ".")

	return message
}
