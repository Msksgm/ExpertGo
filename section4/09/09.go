package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	exampleWithDeadline()
}

func exampleWithDeadline() {
	ctx := context.Background()

	// WithTimeout で 15 秒後にキャンセルされる処理を生成
	d := 15 * time.Second
	timerCtx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	// 10 秒経ったときか、timerCtx がキャンセルされたときか、どちらか先のほうが実行される
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("10秒経ちました")
	case <-timerCtx.Done():
		fmt.Println(timerCtx.Err())
	}

	// さらに 10 秒たったときか、timerCtx がキャンセルされたときか
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("10 秒が経ちました")
	case <-timerCtx.Done():
		fmt.Println(timerCtx.Err())
	}
}
