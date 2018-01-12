package main

import (
	"fmt"
)

func f(x float64) float64 {
	return x*x - 4
	//	return x
}

func checkNull(epsilon, xStart, intervall float64, depht int) {
	fmt.Printf(">checkNull %2d: xStart=%.7f, intervall=%.7f\n", depht, xStart, intervall)
	xEnd := xStart + intervall
	step := intervall / 3
	for x := xStart; x < xEnd; x = x + step {
		//		fmt.Println("x=", x, "y=", f(x))
		//fmt.Println(i, "f(", x, ")=", f(x))
		x1 := x
		x2 := x + step
		y1 := f(x1)
		y2 := f(x2)
		fmt.Printf("   step= %.7f, y1=%.7f, y2=%.7f\n", step, y1, y2)
		if (y1 <= 0 && y2 > 0) || (y1 >= 0 && y2 < 0) {
			if step < epsilon {
				//				fmt.Println("Nullstelle", x1, x2)
				fmt.Printf("Nullstelle: %.2f + %.7f\n", x1, epsilon)
				fmt.Printf("   step= %.7f, y1=%.7f, y2=%.7f\n", step, y1, y2)
				fmt.Printf("<checkNull %2d: xStart=%.7f, intervall=%.7f\n", depht, xStart, intervall)
				return
			} else {
				checkNull(epsilon, x, step, depht+1)
			}
		} else {
			//			fmt.Println("x=", x, "y=", f(x))
		}
	}
	fmt.Printf("<checkNull %2d: xStart=%.7f, intervall=%.7f\n", depht, xStart, intervall)
}

func main() {
	checkNull(0.0005, -10, 20, 0)
}
