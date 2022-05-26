package pkg3

import (
	"fmt"
)

var (
	_  = constInitCheck()
	v1 = variableInitCheck("v1")
	v2 = variableInitCheck("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("pkg3: const c1 has been initialized")
	}

	if c2 != "" {
		fmt.Println("pkg3: const c2 has been initialized")
	}
	return ""
}

func variableInitCheck(name string) string {
	fmt.Printf("pkg3: var %s has been intiailzed\n", name)
	return name
}

func init() {
	fmt.Println("pkg3: first init func invoked")
}

func init() {
	fmt.Println("pkg3: second init func invoked")
}

func main() {

}
