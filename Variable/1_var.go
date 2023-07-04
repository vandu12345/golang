package main

import (
	"fmt"
	"unsafe"
)

//Package Level Visible
// var (
// 	name    string
// 	company string
// 	age     int
// )

// globally visible
var I int = 42

type sample struct {
	a int
	b string
}

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
	// fmt.Printf("%v,%T", j, j)
	//Atoi(String to int)
	// var j string = "42"
	// i, err := strconv.Atoi(j)
	// fmt.Println(i, err)

	// sizeOfIntInBits := bits.UintSize
	// fmt.Printf("%d bits\n", sizeOfIntInBits)

	// var a int
	// fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	// fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	// b := 2
	// fmt.Prins is %s\n"tf("b's typ, reflect.TypeOf(b))

	s := &sample{a: 1, b: "test"}

	// fmt.Printf("Address of s: %p", s)

	fmt.Printf("check: %p\n", unsafe.Pointer(s))
	//
	// fmt.Printf("Address of s: %p", unsafe.Offsetof(s.b))

	// //Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))
	print(p)

	// //Typecasting it to a string pointer and printing the value of it
	// fmt.Println(*(*string)(p))

	// var r byte = 'a'

	// //Print Size
	// fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	// //Print Type
	// fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	// fmt.Printf("Character: %c\n", r)
	// s := "abc"

	// //This will the decimal value of byte
	// fmt.Println([]byte(s))
	fmt.Printf("%U\n", []rune("0b£"))
	//Itoa(int to string)
	// fmt.Printf("%v, %T\n", j, j)

	var r rune
	r = 'A'               // Assigning a single Unicode character (code point) to a rune variable
	fmt.Printf("%c\n", r) // Output: A

	var anotherRune rune = '😊'      // Assigning a rune literal (Unicode code point)
	fmt.Printf("%c\n", anotherRune) // Output: 😊

	//Visibility
	// lower case 1st letter for Package Scope
	//Upper Case first letter to export
	//No Private Scope
	//longer name for Longer Lives

}
