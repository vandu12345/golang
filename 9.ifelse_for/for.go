package main

import "fmt"

func main() {
	// for i := 0; i < 4; i++ {
	// 	fmt.Printf("GeeksforGeeks\n")
	// }
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}
