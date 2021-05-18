package main

import (
	"context"
	"fmt"
	"time"
)

func testCtx(t int) {
	pCtx := context.Background()
	ctx, cancel := context.WithCancel(pCtx)
	defer cancel()

	go func() {
		cancel()
		time.Sleep(10 * time.Second)
		fmt.Println("Stop goroutine")
	}()

	select {
	case <-ctx.Done():
		time.Sleep(8 * time.Second)
		fmt.Println("testContext:", ctx.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("testCtx:", r)
	}
	return
}

func main() {
	testCtx(2)
}
