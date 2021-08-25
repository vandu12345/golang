package main

import (
	"fmt"
)

func main() {
	fmt.Println(42)

	//Variable Declaration
	// var i int
	// i = 42
	// var j int = 42
	// var k float32
	// k = 12
	// j = 13
	j := 24.0

	//Automatic Detection
	// i := 42

	fmt.Printf("%v,%T\n", j, j)
	// fmt.Println(i + j)
	// fmt.Printf("%v, %T\n", j, j)
}
