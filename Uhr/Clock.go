package main

import (
	"fmt"
	"time"
)

func main() {
	//hour := 8
	minute := 30
	//	second := 0

	for hour := 8; hour < 10; hour++ {
		for minute < 60 {
			for second := 0; second < 60; second++ {
				fmt.Println(hour, minute, second)
				time.Sleep(time.Second)
				//				second = second + 1
			}
			minute++
			//			second = 0
		}
		//		hour++
		minute = 0
	}
}
