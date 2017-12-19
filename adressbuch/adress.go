package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	contact struct {
		Name    string
		Address string
		Phone   string
		Age     int
	}
	contact_M map[string]contact //neuer Typ (eine Map(Abbildung)), Index=string, typ=contact
)

var (
	kontakte_m = contact_M{
		"Sarah": contact{
			Name:    "Sarah",
			Address: "99423 Weimar",
			Phone:   "017631254433",
			Age:     18},
		"Klaus": contact{
			Name:    "Klaus",
			Address: "Pragerstr. 42",
			Phone:   "015134876955",
			Age:     102},
		"Mia": contact{
			Name:    "Mia",
			Address: "Pragerstr. 42",
			Phone:   "015134876955",
			Age:     102},
	}
)

func main() {
	kontakte_b, err := ioutil.ReadFile("adressbuch.json")
	if err != nil {
		handleFileReadError(err)
		return
	}
	err = json.Unmarshal(kontakte_b, &kontakte_m)
	if err != nil {
		fmt.Println("Beim Interpretieren von adressbuch.json ist ein Fehler aufgetreten:", err)
		return
	}
	var kname string
	fmt.Print("Gebt bitte einen Namen ein: ")
	fmt.Scanln(&kname)
	k, found := kontakte_m[kname]
	if found {
		fmt.Println(k.Name, k.Address, k.Phone, k.Age)
	} else {
		fmt.Println("Haben wir nicht, überlegen sie sich was anderes. :D")
	}
}

func handleFileReadError(err error) {
	if os.IsNotExist(err) {
		kontakte_b, err := json.MarshalIndent(kontakte_m, "", "  ")
		if err != nil {
			fmt.Println("Beim Konvertieren der Beispielkontakte ist ein Fehler aufgetreten:", err)
			return
		}
		err = ioutil.WriteFile("adressbuch.json", kontakte_b, 0777)
		if err != nil {
			fmt.Println("Beim Schreiben der Beispielkontakte in das File adressbuch.json ist ein Fehler aufgetreten:", err)
			return
		}
		fmt.Println("File adressuch.json hat nicht existiert, es wurde ein Beispieltext erzeugt.")
		return
	}
	fmt.Println("Beim Lesen von adressbuch.json ist ein Fehler aufgetreten:", err)
}
