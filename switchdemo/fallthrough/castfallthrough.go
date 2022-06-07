package main

import "fmt"

func case1() int {
	fmt.Println("eval case1 expr")
	return 1
}

func case2() int {
	fmt.Println("eval case2 expr")
	return 2
}

func switchexpr() int {
	fmt.Println("eval switch expr")
	return 1
}

func main() {
	switch switchexpr() {
	case case1():
		fmt.Println("exec case1")
		fallthrough
	case case2():
		fmt.Println("exec case2")
		fallthrough
	default:
		fmt.Println("exec default")
	}
}
