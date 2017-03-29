package main

import (
    "sync"
    "sync/atomic"
)

var counter int64
var wait sync.WaitGroup

func task() {
	for count := 0; count < 2; count++ {
        atomic.AddInt64(&counter, 1)
	}

	wait.Done()
}

func incrParallell() int64 {
	counter = 0
	wait.Add(2)
	go task()
	go task()

	wait.Wait()
	return counter
}
