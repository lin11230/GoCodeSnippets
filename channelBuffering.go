package main

import "fmt"

func main() {

	message := make(chan string, 2)

	message <- "bufferd"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
