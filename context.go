package main

import (
	"context"
	"fmt"
)

func main() {
	//context.Background() same with context.TODO()

	ctx := context.Background()
	fmt.Println(ctx.Err())
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, ok)

	ctx = context.TODO()
	fmt.Println(ctx.Err())
	deadline, ok = ctx.Deadline()
	fmt.Println(deadline, ok)

	// values in context.WithValue() not stored in map, this is TREE, check debugger
	ctx = context.Background()
	ctx = context.WithValue(ctx, "1", 1)
	ctx = context.WithValue(ctx, "2", 2)

	fmt.Println(ctx.Value("1"))
	fmt.Println(ctx.Value("2"))
	fmt.Println(ctx.Value("0"))

}
