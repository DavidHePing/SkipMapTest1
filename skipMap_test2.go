package main

import (
	"fmt"
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_delete1() {
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15} {
		m0.Store(strconv.Itoa(v), v+100)
	}

	m0.Delete("10")
	v, ok := m0.Load("10")
	if ok {
		fmt.Println("skipmap load 10 with value ", v)
	} else {
		fmt.Println("10 is not in skipmap")
	}

}

func SkipMap_test_delete2() {
	m0 := skipmap.NewString[int]()

	m0.Delete("10") // not panic
	v, ok := m0.Load("10")
	if ok {
		fmt.Println("skipmap load 10 with value ", v)
	} else {
		fmt.Println("10 is not in skipmap")
	}

}
