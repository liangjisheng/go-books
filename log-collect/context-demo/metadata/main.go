package main

import (
	"context"
	"fmt"
)

func add(ctx context.Context, a, b int) int {
	traceID := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n", traceID)
	return a + b
}

func calc(ctx context.Context, a, b int) int {
	traceID := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n", traceID)
	// 再将ctx传入到add中
	return add(ctx, a, b)
}

func main() {
	// 将ctx传递到calc中
	ctx := context.WithValue(context.Background(),
		interface{}("trace_id"), interface{}("123456"))
	calc(ctx, 20, 30)
}
