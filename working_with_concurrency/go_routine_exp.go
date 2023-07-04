package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// Wait group
	// var wg sync.WaitGroup

	// words := []string{
	// 	"alpha",
	// 	"beta",
	// 	"delta",
	// 	"gamma",
	// 	"pi",
	// 	"zeta",
	// 	"eta", "theta", "epsilon",
	// }

	// wg.Add(len(words))
	// for i, x := range words {
	// 	go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	// }

	// // Bad idea to use sleep
	// // time.Sleep(1 * time.Second)
	// wg.Wait()

	// fmt.Println("Final Line value")

}
