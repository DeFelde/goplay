package main

import (
	"fmt"
)

type ()

func main() {
	var kontakte = [][]string{
		[]string{"Sarah", "99423 Weimar", "017631254433"},
		[]string{"Klaus", "Pragerstr. 42", "015134876955"},
	}
	fmt.Println(kontakte[0][0], kontakte[0][2])
}
