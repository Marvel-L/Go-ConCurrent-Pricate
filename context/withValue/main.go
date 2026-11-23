package main

import (
	"context"
	"fmt"
	"time"
)

func func1(ctx context.Context) {
	fmt.Printf("name is: %s", ctx.Value("name").(string))
}

func main() {
	ctx := context.WithValue(context.Background(), "name", "zhangsan")
	go func1(ctx)
	time.Sleep(time.Second)
}
