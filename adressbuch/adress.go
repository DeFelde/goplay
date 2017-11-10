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

var (
	kontakte = []contact{
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
		contact{
			name:    "Mia",
			address: "Pragerstr. 42",
			phone:   "015134876955",
			age:     102},
	}
)

func main() {
	var kname string
	fmt.Print("Gebt bitte einen Namen ein: ")
	fmt.Scanln(&kname)
	//func Scanln(a ...interface{}) (n int, err error)
	for i := 0; i < len(kontakte); i++ {
		k := kontakte[i]
		if k.name == kname {
			fmt.Println(k.name, k.address, k.phone, k.age)
		}
	}
}
