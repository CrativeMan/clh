package main

import (
	"clh/forms"
	"clh/listastree"
	"log"
)

func main() {
	utype, err := forms.InitialForm()
	if err != nil {
		log.Fatalf("[ERROR] trying to run initial form with error:\n%v", err.Error())
	}

	switch utype {
	case 0:
		listastree.RunTree()
	default:
		log.Println("No util know like that")
	}
}
