package main

import (
	"fmt"
)

type (
	contact struct {
		name    string
		address string
		phone   string
		age     int
	}
)

func main() {
	var kontakte = []contact{
		contact{
			name:    "Sarah",
			address: "99423 Weimar",
			phone:   "017631254433",
			age:     18},
		contact{
			name:    "Klaus",
			address: "Pragerstr. 42",
			phone:   "015134876955",
			age:     102},
	}

	for i := 0; i < 2; i++ {
		k := kontakte[i]
		if k.name == "Sarah" {
			fmt.Println(k.name, k.phone, k.age)
		}
	}
}
