package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Watch(ctx, "goroutine1")
	go Watch(ctx, "goroutine2")
	time.Sleep(6 * time.Second)
	cancel()
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "exit!")
		default:
			fmt.Println(name, "working...")
			time.Sleep(1 * time.Second)
		}
	}
}
