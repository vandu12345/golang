package main

import "fmt"

type bot interface {
	Print() string
}
type sbot struct{}
type ebot struct{}

func main() {
	eb := ebot{}
	sb := sbot{}
	Print(eb)
	Print(sb)
	// fmt.Println(sb.Print())
	// fmt.Println(eb.Print())

}
func Print(s bot) {
	fmt.Println(s.Print())
}

// func Print(sbot sb) {

// }

func (eb ebot) Print() string {
	return "English"
}
func (sb sbot) Print() string {
	return "Spanish"
}
