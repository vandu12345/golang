package main

import "fmt"

func main() {
	defer fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("End")
}
