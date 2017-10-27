// quadratf
//ax^2+bx+c=0
//p=b/a; q=c/a
//x^2+px+q=0
//x=-p/2+-sqrt((p^2/4)+q)
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//	fmt.Println(3, "3")
	//	fmt.Println(os.Args)
	fmt.Printf("os.Args=%#v\n", os.Args)
	if len(os.Args) != 4 {
		fmt.Println("Genau drei Eingabewerte vonnöten!")
		return
	}
	var params [3]float64
	for i := 0; i < 3; i++ {
		x, err := strconv.ParseFloat(os.Args[i+1], 64)
		//fmt.Printf("a=%f\n", x)
		if err != nil {
			fmt.Printf("ungültiger Wert für a(%s): %s\n", os.Args[1], err)
			return
		}
		params[i] = x
	}
	fmt.Printf("%.1f x^2 + %.1f x + %.1f = 0\n", params[0], params[1], params[2])

	//	a, err := strconv.ParseFloat(os.Args[1], 64)
	//	fmt.Printf("a=%f\n", a)
	//	if err != nil {
	//		fmt.Printf("ungültiger Wert für a(%s): %s\n", os.Args[1], err)
	//		return
	//	}
	//	b, err := strconv.ParseFloat(os.Args[2], 64)
	//	if err != nil {
	//		fmt.Println("ungültiger Wert für b")
	//		return
	//	}
	//	c, err := strconv.ParseFloat(os.Args[3], 64)
	//	if err != nil {
	//		fmt.Println("ungültiger Wert für c")
	//		return
	//	}
	p := params[1] / params[0]
	q := params[2] / params[0]
	fmt.Printf("x^2 + %.3f x + %.3f = 0\n", p, q)
	d := -p / 2
	Deskriminante := d*d + q
	if Deskriminante < 0 {
		fmt.Println("Für diese Gleichung existiert keine Lösung.")
		return
	}
	wurzel := math.Sqrt(Deskriminante)
	x1 := d + wurzel
	x2 := d - wurzel
	fmt.Printf("x1=%f, x2=%f\n", x1, x2)
	//	fmt.Println("x1=", x1, ", x2=", x2)
}
