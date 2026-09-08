package main

import "fmt"

func main() {
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
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	number2 := &numbers

	// println(&numbers)
	// println(&number2)
	// number2[0] = 22
	// fmt.Printf("numbers: %v\n", numbers)
	// fmt.Printf("numbers: %v\n", number2)

	println(&numbers)
	println(&number2)
	println(number2)
	number2[0] = 22
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers: %v\n", number2)
	fmt.Println("=================================")

	changevalue(&numbers)
	changevalue(number2)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers: %v\n", number2)
}

func changevalue(arr *[8]int) {
	arr[3] = 77
	arr[2] = 12
}
