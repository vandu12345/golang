package main

import "fmt"

func main() {
	// Defer a function call to handle any panics
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	// // Simulate a panic
	// panic("Something went wrong!")
	// fmt.Println("Recovered")

	res, err := 10 / 0
	if err != nil {
		fmt.Println(err)
	}
}
