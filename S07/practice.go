package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter your words:")
	if !scanner.Scan() {
		return
	}
	words := scanner.Text()
	fmt.Println("=======================")

	words = strings.ToLower(words)
	arrInput := strings.Fields(words)

	wordCounts := make(map[string]int, len(arrInput))

	for _, item := range arrInput {
		wordCounts[item]++
	}

	fmt.Printf("%v\n", wordCounts)
}
