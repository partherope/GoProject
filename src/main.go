package main

import "fmt"

var val int

func f(c chan int) {
	val = <-c
}
func main() {
	c := make(chan int)

	c <- 12 // write to a channel
	close(c)
	go f(c)
	fmt.Println("val:", val)
}
