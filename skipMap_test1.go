package main

import (
	"fmt"
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test1() {
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15} {
		m0.Store(strconv.Itoa(v), v+100)
	}

	v, ok := m0.Load("10")
	if ok {
		fmt.Println("skipmap load 10 with value ", v)
	}
}
