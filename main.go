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

func main() {
	//TODO: Call the function Say, with param "hello" using go keyword
	go Say("Hello")

	//TODO: Call the function Say, with param "world" without go keyword
	Say("World")

	time.Sleep(time.Second * 2)

}
