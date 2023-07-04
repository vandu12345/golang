package main

import "fmt"

func main() {
	// Creation
	// 1D
	// grades := [3]int{12, 13, 14}
	// fmt.Printf("%v\t%T", grades, grades)
	// grades := []int{1, 2, 3}
	// b := grades

	// fmt.Println(&grades[0])
	// fmt.Println(&b[0])
	// var grades [3]float32
	// fmt.Println(grades)
	// indexing
	// grades[0] = 12

	// 2D
	// var im [3][3]int = [...][3]int{[...]int{1, 0, 0}, [...]int{0, 1, 0}, [...]int{0, 0, 1}}
	// var im [3][2]int = [...][2]int{[...]int{1, 0}, [...]int{0, 0}, [...]int{0, 1}}

	// fmt.Println(im)
	// Built in Functions
	// fmt.Println(len(im[0]))
	//Working With ARray

	// a := []int{1, 2, 3}
	// // b := a
	// // b[1] = 5
	// a = append(a, 4)
	// fmt.Println(a)
	// fmt.Println(b)

	// Slice
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// b := a //Reference type
	// b[1] = 5
	// fmt.Println(&a[0])
	// fmt.Println(&b[0])
	// fmt.Println(cap(a))
	// fmt.Println(a[:2:6])

	b := make([]int, 3, 100) //underlaying array have 100 elements
	fmt.Println(len(b))
	fmt.Println(cap(b))

	//Slice like list
	// a := []int{}
	// a,1,2,3,4
	// a = append(a, []int{1, 2, 3}...)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, i)
	// 	fmt.Printf("Length - %v\t Capacity-%v\t Value - %v\n", len(a), cap(a), a[i])
	// }
	// fmt.Printf("Capacity-%v\n", cap(a))
	// }

	// //Use SLice Like Stack
	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// b := append(a, 8)
	// a = a[:len(a)-1]
	// fmt.Println(a)
	// b := append(a[:3], a[4:]...)
	// a,1,2,3,4
	// fmt.Println(a)

	// fmt.Println(cap(a))
	// fmt.Println(cap(b))

	// fmt.Println(cap(b))

}
