package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

func main() {
	var t T
	fmt.Println(unsafe.Sizeof(t))

	var s S
	fmt.Println(unsafe.Sizeof(s))
}
