package main

import (
	"fmt"
	"methoddemo/helper"
)

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (T) M3() {}

type T1 struct{}

func (T1) T1M1()   { fmt.Println("T1's M1") }
func (*T1) PT1M2() { fmt.Println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { fmt.Println("T2's M1") }
func (*T2) PT2M2() { fmt.Println("PT2's M2") }

type S struct {
	T1
	*T2
}

func main() {
	var t T
	var p *T
	helper.DumpMethodSet(t)
	helper.DumpMethodSet(p)

	s := S{
		T1: T1{},
		T2: &T2{},
	}
	helper.DumpMethodSet(s)
	helper.DumpMethodSet(&s)
}
