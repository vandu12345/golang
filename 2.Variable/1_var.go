package main

import (
	"fmt"
	"strconv"
)

//Package Level Visible
// var (
// 	name    string
// 	company string
// 	age     int
// )

//globally visible
var I int = 42

func main() {
	// fmt.Println(42)
	// name = "lol"
	// company = "lola"
	// age = 20

	//Variable Declaration
	// var i int
	// i = 42
	// var i int
	// i := 0
	// fmt.Printf("%v,%T", i, i)
	// var j int = 42
	// var k float32
	// k = 12
	// j = 13
	// j := 24.0

	//Automatic Detection
	// i := 42

	// fmt.Printf("%v,%T\n", j, j)
	// fmt.Println(i + j)
	// fmt.Printf("%v, %T\n", j, j)

	//Variable Conversion
	// var i int
	// var j float32 = 32.9
	// i = int(j)
	// fmt.Printf("%v, %T\n", i, i)

	// String it's just an alias for a stream of bytes
	// var i int = 42
	// var j string = strconv.Itoa(i) // It actually looks what uncode char is present at the value 42
	// fmt.Println(j)
	//Atoi(String to int)
	// var j string = "42"
	// i, err := strconv.Atoi(j)
	// fmt.Println(i, err)

	//Itoa(int to string)
	// fmt.Printf("%v, %T\n", j, j)

	//Visibility
	// lower case 1st letter for Package Scope
	//Upper Case first letter to export
	//No Private Scope
	//longer name for Longer Lives

}
