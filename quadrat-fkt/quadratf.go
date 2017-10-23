// quadratf
//ax^2+bx+c=0
//p=b/a; q=c/a
//x^2+px+q=0
//x=-p/2+-sqrt((p^2/4)+q)
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 3
	var b float64 = 4
	var c float64 = 5
	p := b / a
	q := c / a
	d := -p / 2
	Deskriminante := d*d + q
	wurzel := math.Sqrt(Deskriminante)
	x1 := d + wurzel
	x2 := d - wurzel
	fmt.Println(x1, x2)
}
