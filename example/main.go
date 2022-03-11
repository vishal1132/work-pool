package main

import (
	"fmt"

	"github.com/dunzoit/go-pool"
)

func main() {
	pool.
		SetMaxProcs(10).
		SetMaxGoRoutines(100).
		MaxTTL(100).
		MaxQueueSize(100)
	pool.Start()
	for i := 0; i < 10000; i++ {
		pool.Submit(func(args ...interface{}) {
			fmt.Println(args[0])
		}, i)
	}
	pool.Stop()
}
