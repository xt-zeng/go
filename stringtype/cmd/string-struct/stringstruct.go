package main

import (
	"fmt"
	"unicode/utf8"
)

var s = "中国人"

func main() {
	fmt.Printf("the lenght of s = %d\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
	}

	fmt.Printf("\n")

	fmt.Println("the character count in s is", utf8.RuneCountInString(s))

	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}

	fmt.Printf("\n")
}
