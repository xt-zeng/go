package main

import "fmt"

type I interface {
	M()
}

type T struct {
}

func (T) M() {
}

func main() {
	var t T
	var i I = t
	switch i.(type) {
	case T:
		fmt.Println("it is type T")
	case int:
		fmt.Println("it is int")
	case string:
		fmt.Println("it is string")
	}

}
