package main

import (
	"fmt"
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_set_read() {
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15} {
		m0.Store(strconv.Itoa(v), v+100)
	}

	v, ok := m0.Load("10")
	if ok {
		fmt.Println("skipmap load 10 with value ", v)
	}

	_, ok = m0.Load("11")
	if !ok {
		fmt.Println("11 is not in skipmap")
	}
}
