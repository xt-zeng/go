package helper

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)

	if dynType == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty\n", dynType)
		return
	}

	fmt.Printf("%s's method set:\n", dynType)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}

	fmt.Printf("\n")
}
