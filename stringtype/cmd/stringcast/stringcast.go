package main

import "fmt"

func main() {
	s := "中国人"

	rs := []rune(s)
	fmt.Printf("%x\n", rs)

	bs := []byte(s)
	fmt.Printf("%x\n", bs)

	s1 := string(rs)
	fmt.Println(s1)

	s2 := string(bs)
	fmt.Println(s2)
}
