package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 139.8125
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits)
}
