package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("day", currentTime.Day())
	fmt.Println("month", currentTime.Month())
	fmt.Println("year", currentTime.Year())

	fmt.Println("hour", currentTime.Local().Hour())
	fmt.Println("minute", currentTime.Local().Minute())
	fmt.Println("second", currentTime.Local().Second())

	fmt.Println("formatted date", currentTime.Format("01-02-2006"))
	fmt.Println("formatted date", currentTime.Format("01-02-2006 Monday"))
	fmt.Println("formatted date", currentTime.Format("01-02-2006 Monday 15:04:05"))
}
