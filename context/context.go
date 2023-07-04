package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", "123")
}

func main() {
	ctx := context.Background()
	ctx = enrichContext(ctx)
	fmt.Println(ctx.Value("request_id"))
}
