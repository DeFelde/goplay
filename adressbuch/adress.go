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
	contact_M map[string]contact //neuer Typ (eine Map(Abbildung)), Index=string, typ=contact
)

var (
	kontakte_m = contact_M{
		"Sarah": contact{
			name:    "Sarah",
			address: "99423 Weimar",
			phone:   "017631254433",
			age:     18},
		"Klaus": contact{
			name:    "Klaus",
			address: "Pragerstr. 42",
			phone:   "015134876955",
			age:     102},
		"Mia": contact{
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
	k, found := kontakte_m[kname]
	if found {
		fmt.Println(k.name, k.address, k.phone, k.age)
	} else {
		fmt.Println("Haben wir nicht, überlegen sie sich was anderes. :D")
	}
}
