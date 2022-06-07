package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"tom", 13}
	switch p {
	case person{"tony", 33}:
		fmt.Println("match tony")
	case person{"tom", 13}:
		fmt.Println("match tom")
	case person{"lucy", 23}:
		fmt.Println("match lucy")
	default:
		fmt.Println("no match")
	}
}
