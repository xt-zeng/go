package main

import "fmt"

func foo() {
	fmt.Println("call foo")
	bar()
	fmt.Println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic: ", e)
		}
	}()
	fmt.Println("call bar")
	panic("panic occurs in bar")
	zoo()
	fmt.Println("exit zoo")
}

func zoo() {
	fmt.Println("call zoo")
	fmt.Println("exit zoo")
}

func main() {
	fmt.Println("call main")
	foo()
	fmt.Println("exit main")
}
