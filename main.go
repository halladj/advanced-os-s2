package main

import "time"

// TODO: Define the variable turn.
// TODO: Define the "flag [2]int " array.

func Task1() {
	for {
		//TODO: Complete the code.

		time.Sleep(time.Second * 2)
	}
}
func Task2() {
	for {
		//TODO: Complete the code.

		time.Sleep(time.Second * 2)
	}
}
func main() {
	go Task1()
	go Task2()

	for {
	}
}
