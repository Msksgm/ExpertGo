package main

import (
	"context"
	"fmt"
	"time"
)

func exampleWithDeadline() {
	ctx := context.Background()
	// 指定時刻を生成
	d := time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC)
	// 指定時刻にキャンセルされるコンテキストを生成する
	timerCtx, cancel := context.WithDeadline(ctx, d)
	defer cancel()

	// 指定時刻の 1 日後の時刻を生成する
	nd := d.AddDate(0, 0, 1)
	// 時刻 nd になったときか、 timerCtx がキャンセルされたときか、どちらか先のほうが実行される
	select {
	case <-time.After(time.Until(nd)):
		fmt.Println("2022 年 12 月 19 日")
	case <-timerCtx.Done():
		fmt.Println(timerCtx.Err())
	}
}
