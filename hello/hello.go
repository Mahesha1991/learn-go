package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("World")
	fmt.Println(message)
}
