package main

import "fmt"

func main() {
	//var salary float64
	//
	//var minSalary float64 = 5_600_000
	//
	//var tax float64 = 0
	//
	//fmt.Print("Enter your salary: ")

	//fmt.Scanln(&salary)

	//if salary <= minSalary {
	//	tax = 0
	//} else if salary <= minSalary*2 {
	//	tax = 0.05
	//} else if salary <= minSalary*3 {
	//	tax = 0.07
	//} else if salary <= minSalary*4 {
	//
	//	tax = 0.1
	//
	//} else {
	//
	//	tax = 0.15
	//}
	//fmt.Printf("your tax is:%.2f\n", tax*salary)
	//fmt.Printf("your salary is:%.2f\n", salary-(tax*salary))

	//var score float64
	//println("enter your score")
	//fmt.Scan(&score)
	//switch score {
	//
	//case 15, 14, 13, 12, 11:
	//	println("B")
	//default:
	//	println("😢")
	//}
	// var score float64
	// println("enter your score")
	// fmt.Scan(&score)
	// switch {
	// case score >= 16 && score <= 20:
	// 	println("A")
	// case score >= 10 && score <= 15:
	// 	println("B")
	// default:
	// 	println("😢")
	// }
	// ********** break *********
	var month int
	println("please enter your month")
	fmt.Scanln(&month)

	var totaldays int = 0

	switch month {
	case 12:
		totaldays += 29
		fallthrough
	case 11:
		totaldays += 30
		fallthrough
	case 10:
		totaldays += 30
		fallthrough
	case 9:
		totaldays += 30
		fallthrough
	case 8:
		totaldays += 30
		fallthrough
	case 7:
		totaldays += 30
		fallthrough
	case 6:
		totaldays += 31
		fallthrough
	case 5:
		totaldays += 31
		fallthrough
	case 4:
		totaldays += 31
		fallthrough
	case 3:
		totaldays += 31
		fallthrough
	case 2:
		totaldays += 31
		fallthrough
	case 1:
		totaldays += 31

	}
}
