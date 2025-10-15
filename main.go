package main

import (
	"fmt"
	"time"
)

const N = 3 // number of processes

var choosing [N]bool
var number [N]int

// helper function to get the maximum ticket number
// Do Not modify, meant to just be used.
func maxNumber() int {
	max := 0
	for i := 0; i < N; i++ {
		if number[i] > max {
			max = number[i]
		}
	}
	return max
}

func Task(id int) {
	for {
		fmt.Printf("Task-%d: Begin Section\n", id)

		// Entry Section
		// TODO: 1. Set choosing[id]
		// TODO: 2. Assign number[id]
		// TODO: 3. Set choosing[id]

		// TODO: 4. Wait for all other processes (follow Bakery algorithm conditions)

		// Critical Section
		fmt.Printf(">>> Task-%d: Critical Section <<<\n", id)
		time.Sleep(time.Second)

		// Exit Section
		// TODO: 5. Set number[id]

		fmt.Printf("Task-%d: Remainder Section\n", id)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	for i := 0; i < N; i++ {
		go Task(i)
	}
	select {} // keeps the program running
}
