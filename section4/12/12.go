package main

import (
	"context"
	"fmt"

	"github.com/Msksgm/ExpertGo/section4/external"
)

func main() {
	ctx := context.Background()
	valueCtx := external.WithRequestID(ctx, 123)
	requestID, ok := external.GetRequestId(valueCtx)
	if !ok {
		fmt.Println("requestID を持ってなかった")
		return
	}
	fmt.Println(requestID)
}
