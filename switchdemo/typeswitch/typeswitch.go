package main

import "fmt"

func main() {
	var x interface{} = 13
	switch v := x.(type) {
	case nil:
		fmt.Println("v is nil")
	case int:
		fmt.Println("v is int, v = ", v)
	case string:
		fmt.Println("v is string, v = ", v)
	case bool:
		fmt.Println("v is bool, v = ", v)
	default:
		fmt.Println("don't support the type")
	}
}
