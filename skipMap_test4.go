package main

import (
	"sync"
	"time"

	"github.com/zhangyunhao116/skipmap"
)

func SkipMap_test_concurrent1() {
	m := skipmap.NewFunc[int, int](func(a, b int) bool { return a > b })

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		tmp := i
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			m.Store(tmp%10, tmp)
		}()
	}

	wg.Wait()

	m.Range(func(key int, value int) bool {
		println("key:", key, "value:", value)
		return true
	})
}
