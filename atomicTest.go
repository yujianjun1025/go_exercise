package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops int64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)
}
