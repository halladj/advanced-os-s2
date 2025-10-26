package main

import (
	"fmt"
	"time"
)

// TODO: declare the function Say(word string). It print the
//
//	the 'word' three time.
func Say(word string) {
	for i := 0; i < 3; i++ {
		fmt.Println(word)
	}
}

// TODO: Define the "flag [2]int " array.
var flag [2]int

func Task1() {
	for {
		//TODO: Complete the code.
		fmt.Println("---------Begin Section----------")
		flag[0] = 1
		for flag[1] == 1 {
		}
		fmt.Println("Critical Section")
		flag[0] = 0
		fmt.Println("---------Reminding Section----------")

		time.Sleep(time.Second * 2)
	}
}
func Task2() {
	for {
		//TODO: Complete the code.
		fmt.Println("---------Begin Section----------")
		flag[1] = 1
		for flag[0] == 1 {
		}
		fmt.Println("Critical Section")
		flag[1] = 0
		fmt.Println("---------Reminding Section----------")

		time.Sleep(time.Second * 2)
	}
}
func main() {
	//TODO: Call the function Say, with param "hello" using go keyword
	go Say("Hello")

	//TODO: Call the function Say, with param "world" without go keyword
	Say("World")

	time.Sleep(time.Second * 2)

}
