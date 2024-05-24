package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	const key = "myPair"
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, 123)
	ctx = context.WithValue(ctx, "time", time.Now())

	pair := ctx.Value(key)
	time := ctx.Value("time")
	fmt.Printf("%d %T\n", pair, pair)
	fmt.Printf("%v %T\n", time, time)

}
