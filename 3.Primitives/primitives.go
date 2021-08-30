package main

import "fmt"

func main() {
	//Boolean Type
	// var n bool = true //default value false
	// fmt.Println(n)
	// n := 1 == 1
	// fmt.Println(n)

	// Numeric Types
	// 1.Integer
	// var n uint16 = 42
	// fmt.Printf("%v, %T\n", n, n)
	// var a int = 10
	// a := 8
	// var b int8 = 20
	// fmt.Println(a + b) //Error
	// fmt.Println(a + int(b))

	//Bitwise Operator
	// a := 10
	// b := 3
	// fmt.Println(a & b)
	// fmt.Println(^a) //
	// fmt.Println(a | b)
	// fmt.Println(a ^ b)
	// fmt.Println(a &^ b) // 1010  0011   1100   1001

	//Bit Shifting
	// a := 8
	// fmt.Println(a << 3) // x<<y -> x*2^y   1000000
	// fmt.Println(a >> 3) //  x>>y->  x/2^y  0001

	// 2.Floating Point

	// n := 3.14
	// n1 := 1.1e36
	// n1 := 2.1e14
	// var n1 float32
	// fmt.Println(n1)

	// 3.Complex Number - 64,128
	// var n complex64 = 1 + 2i
	// var n complex128 = complex(5, 12)
	// fmt.Println(real(n))
	// fmt.Println(imag(n))
	//operation- add,sub,mult,div

	//String - is immutable but we can do concatenation,represent UTF-8 Character
	// s := "i am lol" //actually it's kind of array
	// b := []byte(s)
	// fmt.Printf("%v,%T\n", b, b)
	// fmt.Printf("%v,%T\n", s, s)
	// fmt.Printf("%v,%T\n", string(s[2]), s[2])

	// rune UTF-32,Alias for int32,special methods normally required to process
	var r rune = 'a'
	fmt.Printf("%v,%T\n", r, r)

	//Text Types
}
