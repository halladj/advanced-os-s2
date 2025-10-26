package main

import (
	"fmt"
	"time"
)

// TODO: Define the variable turn.
//   - we assume Task-1 has the turn initially.
var turn int = 1

func Task1() {
	for {
		//TODO: Complete the code.
		fmt.Println("-------Begin Section---------")
		for turn != 1 {
		}
		fmt.Println("-------Critical Section---------")
		turn = 2
		fmt.Println("-------Reminder Section---------")

		time.Sleep(time.Second * 2)
	}
}
func Task2() {
	for {
		//TODO: Complete the code.
		fmt.Println("-------Begin Section---------")
		for turn != 2 {
		}
		fmt.Println("-------Critical Section---------")
		turn = 1
		fmt.Println("-------Reminder Section---------")

		time.Sleep(time.Second * 2)
	}
}
func main() {
	for i := 0; i < 8; i++ {
		go Task1()
		go Task2()
	}
	for {
	}
}
