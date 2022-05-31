package main

import (
	"fmt"
	"unicode/utf8"
)

func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode character is %c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("utf-8 representation is 0x%x\b", buf)
}

func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the unicode character after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r))
}

func main() {
	encodeRune()

	fmt.Printf("\n")

	decodeRune()

}
