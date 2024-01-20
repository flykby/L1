package main

import (
	"fmt"
	"strings"
)

func isOnlyUniqueChars(str string) bool {
	mapChars := make(map[rune]int, 0)
	for _, r := range strings.ToLower(str) {
		_, ok := mapChars[r]
		if ok {
			return false
		}
		mapChars[r] = 1
	}
	return true
}

func main() {
	fmt.Println(isOnlyUniqueChars("abcd"))
	fmt.Println(isOnlyUniqueChars("abcdA"))
	fmt.Println(isOnlyUniqueChars("абырглав"))
	fmt.Println(isOnlyUniqueChars("абырглв"))

}