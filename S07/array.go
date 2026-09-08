package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// arr2 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// change(arr)
	// fmt.Printf("%v\n", arr)
	additem(&arr)
	fmt.Printf("%v\n", arr)
}
func change(number [8]int) {
	for i, _ := range number {
		number[i] = number[i] * 10
	}
}
func additem(add *[]int) {
	*add = append(*add, 6)

}

// slice in function if we append somting we dont have it in original he copy the slice we can use * &
// func change(number [8]int) {
//  [1 2 3 4 5 6 7 8]
// 	for i, _ := range number {
// 		number[i] = number[i] * 10
// 	}
// func change(number []int) {

// [10 20 30 40 50 60 70 80]
// 	for i, _ := range number {
// 		number[i] = number[i] * 10
// 	}
// }

// arr[]
// arr[0][1] row 0 coulmn 1
// this is copy and not grow and len == cap

//var myarr2 [3]int={1,3,4,5} //out of bound
//names := [8]string{"ali", "alice", "reza", "ahmad", "abdollah", "jafar", "neda", "thomas"}
//for index, val := range names {
//	search := "reza"
//	if val == search {
//		println("name found this index is :", index)
//		break
//	}
//}
//numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
//
//number2 := &numbers

// println(&numbers)
// println(&number2)
// number2[0] = 22
// fmt.Printf("numbers: %v\n", numbers)
// fmt.Printf("numbers: %v\n", number2)

//println(&numbers)
//println(&number2)
//println(number2)
//number2[0] = 22
//fmt.Printf("numbers: %v\n", numbers)
//fmt.Printf("numbers: %v\n", number2)
//fmt.Println("=================================")
//
//changevalue(&numbers)
//changevalue(number2)
//
//fmt.Printf("numbers: %v\n", numbers)
//fmt.Printf("numbers: %v\n", number2)
// slice pointer to arrey
// cost perfomansh

//arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
//arr3 := make([]int, 8, 16) //cap =16 len=8
//arr5 := arr[:]
// if i change slice arr change too
// if new slice create with add cap the pointer remove and make new slice
//arr5 := arr[:]
//arr5[0] = 22
//fmt.Printf("%v\n", arr5)
//fmt.Printf("%v\n", arr)

//}/

//func changevalue(arr *[8]int) {
//	arr[3] = 77
//	arr[2] = 12
// }
