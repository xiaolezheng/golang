package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("publish news: ", text)
	}() // 注意这里的括号,必须调用匿名函数
}

func main() {
	Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// 等待发布新闻
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")
}
