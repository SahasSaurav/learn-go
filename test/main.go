package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func Add[T Number](x T, y T) T {
	return x + y
}

func main() {
	sum := Add(1.2, 2.0)
	fmt.Println("sum", sum)
}
