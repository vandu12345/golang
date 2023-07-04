package main

import "fmt"

func main() {
	//Naming Convention
	//Camel Case
	// const myVar

	//Typed Constant
	// const a int = 12
	// var b int16 = 13
	// fmt.Println(a + b)//Error

	//Untyped Constant
	const a = 12
	fmt.Printf("%v,%T\n", a, a)
	var b int16 = 13
	fmt.Println(a + b)
	fmt.Printf("%v,%T\n", b, b)

	//Enumerated Constants

	//Enumeration Expression
}
