package main

import (
	"context"
	"fmt"
)

func main() {
	exampleWithValue()
}

func exampleWithValue() {
	// string 型の"key1"と"value1"が格納されたvalueCtxを生成
	ctx := context.Background()
	valueCtx := context.WithValue(ctx, "key1", "value1")
	// valueCtx から保存した値を取得して、元の型にキャストする
	fmt.Println(valueCtx.Value("key1").(string))
}
