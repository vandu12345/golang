package main

import "fmt"

var area = func(l int, b int) int {
	return l * b
}

// func sum(x, y int) int {
// 	return x + y
// }
// func partialSum(x int) func(int) int {
// 	return func(y int) int {
// 		return sum(x, y)
// 	}
// }

// func squareSum(x int) func(int) func(int) int {
// 	return func(y int) func(int) int {
// 		return func(z int) int {
// 			return x*x + y*y + z*z
// 		}
// 	}
// }

type First func(int) int
type Second func(int) First

func squareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	fmt.Println(area(20, 30))

	// 	func(l int, b int) {
	// 		fmt.Println(l * b)
	// 	}(20, 30)

	// 	// Closure in go

	// 	l := 20
	// 	b := 30

	// 	func() {
	// 		var area int
	// 		area = l * b
	// 		fmt.Println(area)
	// 	}()
	// 	fmt.Println(area)

	// Higher order functions
	// partial := partialSum(3)
	// fmt.Println(partial(7))

	// Returning func from other functions
	fmt.Println(squareSum(5)(6)(7))

}
