package main

import (
	"fmt"
	"sort"
)

func main() {

	//func() {

	//}()

	//myx := func() {

	//}
	//myx()
	//println(func(number ...int) (total int) {
	//	for _, item := range number {
	//		total += item
	//	}

	//	return
	//}(1, 2, 3))

	//sum := func(number ...int) (total int) {
	//	for _, item := range number {
	//		total += item
	//	}

	//	return
	//}
	//println(sum(1, 2, 3))
	number := []int{12, 5, 6, 162, 63, 7, 11}
	fmt.Printf("%v\n", number)
	sort.Slice(number, func(i, j int) bool {
		return number[i] < number[j]
		// az kochej be bozorg
	})
	fmt.Printf("%v", number)

}

// name ...int
// name ...interface{}
//name int (sum int ,negtive int)//defult is 0 false '' ""
// interface -> fmt. (%v)
// anonymos fuctnion
//  funciont is a type like int string
// funciont in fuciotn ya bayad bdi ve varible ya call cony() bdon esm
