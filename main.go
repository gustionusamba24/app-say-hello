package main

import (
	"fmt"
	go_say_hello "github.com/gustionusamba24/go-say-hello"
)

func main() {
	greeting := go_say_hello.SayHello()
	fmt.Println(greeting)
}
