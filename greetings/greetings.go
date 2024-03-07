package greetings

import "strings"

func Hello(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	message := "Hi, " + strings.Join(names, ", ") + ". Welcome!"
	return message
}
