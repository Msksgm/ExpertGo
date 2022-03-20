package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	if err := doSomeThingParallel(3); err != nil {
		log.Fatal(err)
	}
}

func doSomeThingParallel(workerNum int) error {
	// 必要なコンテキストを生成する
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)

	// 正常完了時にコンテキストのリソースを解放
	defer cancel()

	// 複数のゴルーチンからエラーメッセージを集約するためにチャネルを用意する
	errCh := make(chan error, workerNum)
	// workerNum 分の平行処理をおこなう
	wg := sync.WaitGroup{}

	for i := 0; i < workerNum; i++ {
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := doSomeThingWithContext(cancelCtx, num); err != nil {
				cancel()
				errCh <- err
			}
			return
		}(i)
	}

	// 平行処理の終了を待つ
	wg.Wait()

	// エラーチャネルに入ったメッセージを取り出す
	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	// エラーが発生していれば、最初のエラーを返す
	if len(errs) > 0 {
		return errs[0]
	}

	// 正常終了
	return nil
}

func doSomeThingWithContext(ctx context.Context, num int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	fmt.Println(num)
	return nil
}
