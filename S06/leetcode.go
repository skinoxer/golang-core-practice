package main

import (
	"strings"
)

func main() {
	text := "allice and bob love leetcode"
	newtext := strings.Split(text, " ")
	max := 0
	for _, value := range newtext {
		if len(value) > max {
			max = len(value)
		}

	}
	print(max)

}
