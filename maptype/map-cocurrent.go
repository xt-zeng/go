package main

import (
	"fmt"
	"time"
)

func doInteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d, %d]", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	go func() {
		for i := 0; i < 1000; i++ {
			doInteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	time.Sleep(5 * time.Second)

	for k, v := range m {
		fmt.Printf("[k, v] = [%d, %d]\n", k, v)
	}
}
