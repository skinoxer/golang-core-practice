package main

import "fmt"

// func main() {
//*** practice one ***
//list1 := []string{"Ali", "Reza"}
//list2 := []string{"Sara", "Neda"}
//singleList := append(list1, list2...)
//fmt.Printf("%v\n", singleList)
//newuser := "VIP_Hassan"
//targetIndex := 0
//
//singleList = append(singleList, "")
//copy(singleList[targetIndex+1:], singleList[targetIndex:])
//singleList[targetIndex] = newuser
//fmt.Printf("%v\n", singleList)
//for _, item := range singleList {
//	if item == "Reza" {
//		println("رضا در لیست است!")
//	} else {
//		println("خوش آمدید")
//	}
//}
//*** practice two ***
//grades := []int{15, 9, 20, 12, 18}
//for i, item := range grades {
//	item = item + 2 //grades not change item make copy
//	grades[i] = grades[i] + 2
//}
//fmt.Printf("%v\n", grades)
//for _, item := range grades {
//	switch {
//	case item >= 18:
//		fmt.Println("A")
//	case item >= 15 && item <= 17:
//		fmt.Println("B")
//	case item < 15:
//		fmt.Println("C")
//	}
//}

// *** practice three ***
func main() {
	// cart := []string{"Apple", "Banana", "Cherry", "Mango"}
	//
	// targetAfter := 2
	// targetBefor := 0
	// cart = removeItem(targetBefor, targetAfter, cart)
	//
	// fmt.Printf("%v\n", cart)
	// cart = cart[:0]
	//	bigData := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//	smalldata := make([]int, 3, 3)
	//	copy(smalldata, bigData[3:6])
	//	bigData[0] = 22
	//	copy(bigData, smalldata[0:3])
	//	fmt.Printf("%v\n", smalldata)
	//	fmt.Printf("%v\n", bigData)
}

//func removeItem(targetBefor int, targetAfter int, cart []string) []string {
//	cart = append(cart[targetBefor:targetAfter], cart[targetAfter+1:]...)
//	return cart

//}
