package main

import (
	"strconv"
	"sync"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_concurrent1() {
	m := skipmap.NewString[int]()

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		tmp := i
		go func() {
			defer wg.Done()
			m.Store(strconv.Itoa(tmp%10), tmp)
		}()
	}

	wg.Wait()

	m.Range(func(key string, value int) bool {
		println("key:", key, "value:", value)
		return true
	})
}
