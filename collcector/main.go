package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 容量代表最多同时运行的轮次；设为 2，方便观察名额满时跳过。
	limiter := make(chan struct{}, 2)

	for round := 1; ; round++ {
		time.Sleep(500 * time.Millisecond) // 每隔一段时间尝试启动一轮。

		select {
		case limiter <- struct{}{}: // 1. 占用一个名额。
			fmt.Printf("第 %d 轮：开始采集\n", round)

			var wg sync.WaitGroup
			wg.Add(2)

			// 2. 本轮两个任务并发执行。
			go func(round int) {
				defer wg.Done()
				collect(round, "Basegrid", time.Second)
			}(round)

			go func(round int) {
				defer wg.Done()
				collect(round, "AccountInfo", 2*time.Second)
			}(round)

			// 3. 新协程负责等待，主循环继续尝试下一轮。
			go func(round int) {
				wg.Wait() // 等本轮两个任务都结束。
				<-limiter // 4. 归还名额，后续轮次才能启动。
				fmt.Printf("第 %d 轮：全部完成，释放名额\n", round)
			}(round)

		default:
			fmt.Printf("第 %d 轮：名额已满，跳过\n", round)
		}
	}
}

func collect(round int, name string, duration time.Duration) {
	fmt.Printf("第 %d 轮：%s 开始\n", round, name)
	time.Sleep(duration) // 模拟访问交易所等耗时操作。
	fmt.Printf("第 %d 轮：%s 完成\n", round, name)
}
