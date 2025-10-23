package main

import "fmt"

// TODO: declare the shared variable x.
var x int

// TODO: define the increament() function.
//
//	increaments x by 1 + print the current x value.
func increment() {
	x++
	fmt.Println(x)
}

func main() {

	// TODO: call increament() multiple times, with/wihtout
	//       'go'keyword.
	for i := 0; i < 6; i++ {
		go increment()
		increment()
	}

	for {
	}
}
