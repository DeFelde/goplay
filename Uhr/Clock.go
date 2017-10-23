package main

import (
	"fmt"
	"time"
)

func main() {
	hour := 8
	minute := 30
	second := 0

	for hour < 10 {
		for minute < 60 {
			for second < 60 {
				fmt.Println(hour, minute, second)
				time.Sleep(time.Second)
				second = second + 1
			}
			minute++
			second = 0
		}
		hour++
		minute = 0
	}
}
