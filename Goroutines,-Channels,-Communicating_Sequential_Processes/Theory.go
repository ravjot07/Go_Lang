package main

import (
	"fmt"
)

//Function to compute the sum of an array slice

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     // creating a channel to communicate results
	go sum(s[:len(s)/2], c) // first half to goroutine
	go sum(s[len(s)/2:], c) // second half to goroutine

	x, y := <-c, <-c // recive from channel c

	fmt.Println("sum of first helf:", x)
	fmt.Println("sum of Second helf:", y)
	fmt.Println("Total sum:", x+y)

}
