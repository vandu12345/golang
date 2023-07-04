package main

import "fmt"

func main() {
	// switch day := 4; day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid")
	// }

	// var value int = 2

	// Switch statement without an
	// optional statement andexpression
	// switch {
	// case value == 1:
	// 	fmt.Println("Hello")
	// case value == 2:
	// 	fmt.Println("Bonjour")
	// case value == 3:
	// 	fmt.Println("Namstay")
	// default:
	// 	fmt.Println("Invalid")
	// }

	var value interface{}
	// value = int(6)
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}
}
