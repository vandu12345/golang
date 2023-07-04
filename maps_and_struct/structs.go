package main

import "fmt"

type Doctor struct {
	n    int
	act  string
	name []string
}
type contact struct {
	email string
	zip   int
}
type person struct {
	fn      string
	ln      string
	coninfo contact
}

//Values Type - >int,float,string,bool,structs
//Reference Type -> slices,maps,channels,pointers,functions

func main() {
	per := person{"suv", "mah", contact{"lol", 23}}
	fmt.Println(per.coninfo.email)
	// suv := person{"suv", "mah", 23}
	// suv := person{fn: "suv", ln: "mah", age: 23}
	// fmt.Println(suv)

	// var declare_struct person
	// declare_struct.fn = "lol"
	// declare_struct.ln = "jol"
	// declare_struct.age = 24
	// fmt.Println(declare_struct)
	// fmt.Printf("%+v\n", declare_struct)
	//What are they

	//Create
	d1 := Doctor{
		n:    1,
		act:  "lol",
		name: []string{"lol", "lola", "jola"},
	}
	fmt.Println(d1.name[1])
	//Manipulate
}
