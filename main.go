// package main

// import (
// 	"fmt"

// 	m "github.com/go_project/golang/Packages"
// )

// func main() {
// 	xs := []float64{1, 2, 3, 4}
// 	avg := m.Average(xs)
// 	fmt.Println(avg)
// }

package main

import "fmt"

var (
	number = 10.0
	delta  = .33333
)

func main() {
	fmt.Println(squareDelta(number))
}
