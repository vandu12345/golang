package main

import "fmt"

func main() {
	//What are they? like key value

	//Create
	sp := map[rune]int{
		'a': 97,
		'b': 98,
	}
	// m := map[[3]int]string{}
	// sp := make(map[string]int)
	// fmt.Println(m)
	sp['c'] = 99
	sp1 := sp
	delete(sp1, 'a')
	fmt.Println(sp)
	fmt.Println(sp1)
	// val, ok := sp['a']
	// fmt.Println(val, ok)
	// fmt.Println(len(sp))
	//Naming Convention
	//EMbedding
	//Tags
	//Manipulate
}
