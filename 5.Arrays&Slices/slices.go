package main
//Simple Program
import "fmt"

func main() {
	arr := []string{"diamond", print()}
	// arr  = append(arr,"jol")
	// arr = append(arr, "hihi")
	arr = append(arr, "jol")

	fmt.Println(arr)
	fmt.Println(&arr[0])
	fmt.Println(cap(arr))

}

func print() string {
	return "lol"
}
