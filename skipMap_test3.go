package main

import (
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_read_all1() {
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15, 17} {
		m0.Store(strconv.Itoa(v), v+100)
	}

	m0.Range(func(key string, value int) bool {
		println("key:", key, "value:", value)
		return true
	})
}

func SkipMap_test_order1() {
	m0 := skipmap.NewFunc[int, string](func(a, b int) bool { return a < b })
	for _, v := range []int{21, 19, 17, 15} {
		m0.Store(v, strconv.Itoa(v+200))
	}
	m0.Range(func(key int, value string) bool {
		println("m0 found ", key, value)
		return true
	})
}

func SkipMap_test_order2() {
	m0 := skipmap.NewFunc[int, string](func(a, b int) bool { return a < b })
	for _, v := range []int{4, 1, 3, 2} {
		m0.Store(v, strconv.Itoa(v+200))
	}
	m0.Range(func(key int, value string) bool {
		println("m0 found ", key, value)
		return true
	})
}
