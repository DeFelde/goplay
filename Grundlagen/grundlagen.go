package main

import "fmt"

func main() {
	var (
		a  int     = 42
		b  float64 = 3.14159
		bs string  = "3.14159"
		c  bool    = true
		d  string  = "work%d"
	)
	fmt.Printf("a int=%d,b float=%f, c bool=%v, d string='%s'\n", a, b, c, d)
	var (
		x = [3]int{9, 4, 2}
		y = [2][3]int{{1, 2, 3}, {4, 5, 6}}
		z = [2]string{"ying", "yang"}
	)
	fmt.Printf("x=%#v\n y=%#v\n  z=%#v\n", x, y, z)
	for i := range z {
		fmt.Printf("z[%d]='%s'\n", i, z[i])
	}
	for i := range x {
		fmt.Println(i, x[i])
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
