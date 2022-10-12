package main

import "fmt"

func greet(c chan string) {
	fmt.Println("hello, " + <-c + "!")
}

func main() {
	println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "World"
	println("main() stopped")
}
