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
			phone:   "017631254433"},
		contact{
			name:    "Klaus",
			address: "Pragerstr. 42",
			phone:   "015134876955"},
	}
	fmt.Println(kontakte[0].name, kontakte[0].phone)
}
