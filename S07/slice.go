package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"hamed", "ali", "reza", "ahamad"}
	for _, item := range names {

		item = strings.ToUpper(item)
	}
	fmt.Printf("%v\n", names)
}

// in for age az itme estefade konim va mostaghim name taghir nadim tagirat ma ded nakhad shod names[i] and item make a copy of slice
// names := []string{"hamed", "ali", "reza", "ahamad"}
// for i, _ := range names {

// 	names[i] = strings.ToUpper(names[i])
// }
// fmt.Printf("%v\n", names)
