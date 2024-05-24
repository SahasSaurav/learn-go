package main

import (
	"context"
	"fmt"
)

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	const key = "myPair"
	ctx := context.WithValue(context.Background(), key, 123)

	pair := ctx.Value(key)
	fmt.Println(pair)

}
