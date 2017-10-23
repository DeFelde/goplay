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
	if len(os.Args) != 4 {
		fmt.Println("Genau drei Eingabewerte vonnöten!")
		return
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("ungültiger Wert für a")
		return
	}
	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("ungültiger Wert für b")
		return
	}
	c, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("ungültiger Wert für c")
		return
	}
	//	var (
	//		a float64 = 3
	//		b float64 = 4
	//		c float64 = 17
	//	)
	p := b / a
	q := c / a
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
