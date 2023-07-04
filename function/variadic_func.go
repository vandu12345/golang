package main

import (
	"fmt"
	"reflect"
)

func main() {
	// variadicExample("red", "blue", "green", "yellow")
	variadicExample(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}

// func variadicExample(s ...string) {
// 	fmt.Println(s[0])
// 	fmt.Println(s[3])
// }

func variadicExample(i ...interface{}) {
	// fmt.Print(reflect.ValueOf(i).Kind())
	// fmt.Println("\n\n")
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
