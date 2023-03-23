package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	//Make the channel buffered so we do not wait
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
