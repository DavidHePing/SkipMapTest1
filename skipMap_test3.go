package main

import (
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_read_all() {
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15, 17} {
		m0.Store(strconv.Itoa(v), v+100)
	}

	m0.Range(func(key string, value int) bool {
		println("key:", key, "value:", value)
		return true
	})
}
