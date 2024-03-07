package main

import (
	"example.com/greetings"
	"fmt"
	"os"
)

func main() {
	// If user enter name
	if len(os.Args) > 1 {
		message := greetings.Hello(os.Args[1])
		fmt.Println(message)
	} else {
		fmt.Println(greetings.Hello("world"))
	}
}
