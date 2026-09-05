package main

import (
	"fmt"
	"go/printer"
)

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
	var score float64
	println("enter your score")
	fmt.Scan(&score)
	switch {
	case score >= 16 && score <= 20:
		println("A")
	case score >= 10 && score <= 15:
		println("B")
	default:
		println("😢")
	}
}
