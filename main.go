package main

import (
	"strconv"

	"github.com/zhangyunhao116/skipmap"
)

func main() {
	print("123")
	m0 := skipmap.NewString[int]()

	for _, v := range []int{10, 12, 15} {
		m0.Store(strconv.Itoa(v), v+100)
	}
}
