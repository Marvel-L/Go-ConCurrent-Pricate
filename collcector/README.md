# Collector：面试版

全部逻辑在 `main.go`，只使用 channel、WaitGroup 和 goroutine。

## 运行

```bash
go run ./collcector
```

每 500ms 尝试启动一轮，最多同时运行 2 轮；Basegrid 耗时 1 秒，AccountInfo 耗时 2 秒。观察“名额已满，跳过”和“全部完成，释放名额”，按 Ctrl+C 停止。容量改为 10 就对应原代码的轮次配置。

## 面试时这样讲

“我用有缓冲的 channel 做并发信号量。启动一轮前向 channel 放一个空结构体，表示占用名额；满了就通过 select 的 default 跳过。

每轮用 WaitGroup 管理两个并发任务。单独启动一个协程执行 Wait，避免阻塞主调度循环；两个任务完成后，从 channel 取出一个占位值，释放名额。”

记住：**channel 管名额，WaitGroup 管完成；Wait 不会自动释放 channel。**

这里限制的是轮次，不是所有 goroutine 的总数。每轮有两个采集协程和一个等待协程；如果任务还启动子协程，也必须等子协程结束后再返回，外层 WaitGroup 才能跟踪到完整任务。
