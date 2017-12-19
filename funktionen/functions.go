package main

import (
	"fmt"
)

func f(x float64) float64 {
	//return x * x * x
	return -x
}

//TODO: Nullstelle genauer bestimmen
func isNull(x, step float64) {
	y1 := f(x)
	y2 := f(x + step)
	if (y1 < 0 && y2 > 0) || (y1 > 0 && y2 < 0) {
		fmt.Println("Nullstelle")
	}
}

func main() {
	var i = 1
	step := 0.3
	//fmt.Printf("step %T=%v\n", step, step)
	for x := -3.0; x < 7; x = x + step {
		fmt.Println(i, "f(", x, ")=", f(x))
		isNull(x, step)
		//if f(x) < 0 && 0 < f(x+step) {
		//	fmt.Println("Nullstelle")
		//}
		i++
	}
}