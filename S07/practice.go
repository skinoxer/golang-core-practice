package main

import (
	"fmt"
)

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)

// func main() {
// scanner := bufio.NewScanner(os.Stdin)
//
// fmt.Println("please enter your words:")
//
//	if !scanner.Scan() {
//		return
//	}
//
// words := scanner.Text()
// fmt.Println("=======================")
//
// words = strings.ToLower(words)
// arrInput := strings.Fields(words)
//
// wordCounts := make(map[string]int, len(arrInput))
//
//	for _, item := range arrInput {
//		wordCounts[item]++
//	}
//
// fmt.Printf("%v\n", wordCounts)
func main() {
	//arr := []int{4, 2, 5, 2, 4, 1, 8, 5}
	//obj := make(map[int]bool, len(arr))
	//keys := make([]int, len(obj))
	//for _, item := range arr {
	//	obj[item] = true
	//
	//}
	//for i, _ := range obj {
	//	keys = append(keys, i)
	//
	//}
	//
	//sort.Ints(keys)
	//fmt.Printf("%v\n", keys)

	arr := []int{4, 2, 5, 2, 4, 1, 8, 5}

	seen := make(map[int]bool, len(arr))

	result := make([]int, 0, len(arr))

	for _, item := range arr {

		if !seen[item] {
			seen[item] = true

			result = append(result, item)
		}
	}

	fmt.Printf("%v\n", result)
}
