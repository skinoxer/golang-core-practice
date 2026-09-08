package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{10, 20, 30, 40, 5, 6, 7, 8}
	count := copy(arr, arr2)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", count)
}
